// Package settings loads /etc/z-panel/config.toml and defaults.
package settings

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"

	"z-panel/internal/config"
	"z-panel/internal/executil"
	"z-panel/internal/i18n"
	"z-panel/internal/root"
	"z-panel/internal/state"
)

// Cfg is stored in config.toml.
type Cfg struct {
	Table              string `toml:"table"`
	SystemdNetworkDir  string `toml:"systemd_network_dir"`
	UfwMarker          string `toml:"ufw_marker"`
	XrayTunManagedMark string `toml:"xray_tun_managed_mark"`
	DefaultLANCIDR     string `toml:"default_lan_cidr"`
	DefaultLANDev      string `toml:"default_lan_dev"`
	DefaultXrayAddr    string `toml:"default_xray_addr"`
	DefaultXrayPeer    string `toml:"default_xray_peer"`
	// SocketPath: Unix socket for daemon IPC (empty → default in /run/z-panel/).
	SocketPath string `toml:"socket_path"`
	// PidPath: daemon PID file (empty → config.DefaultPidPath). Not used for non-daemon commands.
	PidPath string `toml:"pid_path"`
	// XrayRedirect: persisted snapshot per interface for xray-redirect (single-file config principle).
	XrayRedirect map[string]state.File `toml:"xray_redirect,omitempty"`
	// Daemon: 1 = when the daemon is running, run subcommands via it; 0 = always run locally.
	Daemon int `toml:"daemon"`
	// NoBanner: omit the first stderr line (z-panel version) on each run.
	NoBanner bool `toml:"no_banner"`
	// SSHNoMultiplex: do not start SSH ControlMaster for z-panel --ssh=… (one TCP connection per ssh).
	SSHNoMultiplex bool `toml:"ssh_no_multiplex"`
	// SSHNoTTY: omit ssh -t for remote sudo (e.g. NOPASSWD).
	SSHNoTTY bool `toml:"ssh_no_tty"`
	// Language: auto (follow system LANG / LANGUAGE / LC_*), or fixed en, ru, zh, …
	Language string `toml:"language"`
	// SchemaVersion is written by the program; 0 in file means legacy (treated as 1).
	SchemaVersion int `toml:"schema_version"`
}

// C is the loaded configuration; non-nil after Load().
var C *Cfg

// schemaJustAutoMigrated is set when Load() upgraded schema_version and saved config in this process.
var schemaJustAutoMigrated bool

// SchemaJustAutoMigrated reports whether the last Load() applied a schema upgrade and persisted it.
func SchemaJustAutoMigrated() bool {
	return schemaJustAutoMigrated
}

func defaults() Cfg {
	return Cfg{
		Table:              "51843",
		SystemdNetworkDir:  "/etc/systemd/network",
		UfwMarker:          "z-panel",
		XrayTunManagedMark: "# z-panel-managed",
		DefaultLANCIDR:     "192.168.0.0/16",
		DefaultLANDev:      "lan0",
		DefaultXrayAddr:    "10.252.0.1/30",
		DefaultXrayPeer:    "10.252.0.2/30",
		SocketPath:         config.DefaultSocketPath,
		PidPath:            config.DefaultPidPath,
		Daemon:             0,
		Language:           "auto",
		SchemaVersion:      0,
	}
}

// readConfigFromDisk loads config.toml or defaults when the file is missing.
func readConfigFromDisk() (c Cfg, fileExists bool, err error) {
	path := config.ConfigFile
	c = defaults()
	b, readErr := os.ReadFile(path)
	if readErr != nil {
		if os.IsNotExist(readErr) {
			normalize(&c)
			mergeLegacyStateFiles(&c)
			return c, false, nil
		}
		return c, false, fmt.Errorf(i18n.T("settings.err.read"), path, readErr)
	}
	if err := toml.Unmarshal(b, &c); err != nil {
		return c, true, fmt.Errorf(i18n.T("settings.err.parse"), path, err)
	}
	normalize(&c)
	mergeLegacyStateFiles(&c)
	return c, true, nil
}

// reloadFromDiskIntoC refreshes [C] from disk without running schema migration (used after Write).
func reloadFromDiskIntoC() error {
	c, _, err := readConfigFromDisk()
	if err != nil {
		return err
	}
	C = &c
	return nil
}

// persistConfigToDisk writes canonical config.toml (requires root). Does not reload [C].
func persistConfigToDisk(c Cfg) error {
	if err := root.Require(); err != nil {
		return err
	}
	normalize(&c)
	if err := os.MkdirAll(config.ConfigDir, 0o755); err != nil {
		return fmt.Errorf(i18n.T("settings.err.mkdir"), config.ConfigDir, err)
	}
	body, err := toml.Marshal(&c)
	if err != nil {
		return err
	}
	hdr := []byte(i18n.T("settings.config_hdr"))
	path := config.ConfigFile
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, append(hdr, body...), 0o644); err != nil {
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		_ = os.Remove(tmp)
		return fmt.Errorf(i18n.T("settings.err.write"), path, err)
	}
	return nil
}

// Load reads config.toml; if missing, only defaults apply.
// If the file exists and schema_version is older than [CurrentSchemaVersion], applies all registered
// upgrade steps, persists immediately, and reloads from disk (requires root for the write).
func Load() error {
	schemaJustAutoMigrated = false
	c, fileExists, err := readConfigFromDisk()
	if err != nil {
		return err
	}
	if !fileExists {
		C = &c
		return nil
	}
	stored := EffectiveStoredSchema(c.SchemaVersion)
	if stored < CurrentSchemaVersion {
		if err := applySchemaUpgradesAuto(&c, stored); err != nil {
			return err
		}
		c.SchemaVersion = CurrentSchemaVersion
		normalize(&c)
		// Persist immediately when privileged; otherwise keep migrated view in memory (disk updated on next sudo/root run).
		if err := root.Require(); err != nil {
			C = &c
			if !c.NoBanner {
				fmt.Fprint(os.Stderr, i18n.T("settings.migrate_deferred", stored, CurrentSchemaVersion, config.ConfigFile))
			}
			return nil
		}
		if executil.RemoteSSHHost() != "" && os.Geteuid() != 0 {
			C = &c
			if !c.NoBanner {
				fmt.Fprint(os.Stderr, i18n.T("settings.migrate_deferred", stored, CurrentSchemaVersion, config.ConfigFile))
			}
			return nil
		}
		if err := persistConfigToDisk(c); err != nil {
			return fmt.Errorf("%s: %w", i18n.T("settings.err.migrate_persist"), err)
		}
		schemaJustAutoMigrated = true
		if !c.NoBanner {
			fmt.Fprint(os.Stderr, i18n.T("settings.migrate_auto_stderr", stored, CurrentSchemaVersion, config.ConfigFile))
		}
		return reloadFromDiskIntoC()
	}
	C = &c
	return nil
}

func normalize(c *Cfg) {
	d := defaults()
	if strings.TrimSpace(c.Table) == "" {
		c.Table = d.Table
	}
	if strings.TrimSpace(c.SystemdNetworkDir) == "" {
		c.SystemdNetworkDir = d.SystemdNetworkDir
	}
	if strings.TrimSpace(c.UfwMarker) == "" {
		c.UfwMarker = d.UfwMarker
	}
	if strings.TrimSpace(c.XrayTunManagedMark) == "" {
		c.XrayTunManagedMark = d.XrayTunManagedMark
	}
	if strings.TrimSpace(c.DefaultLANCIDR) == "" {
		c.DefaultLANCIDR = d.DefaultLANCIDR
	}
	if strings.TrimSpace(c.DefaultLANDev) == "" {
		c.DefaultLANDev = d.DefaultLANDev
	}
	if strings.TrimSpace(c.DefaultXrayAddr) == "" {
		c.DefaultXrayAddr = d.DefaultXrayAddr
	}
	if strings.TrimSpace(c.DefaultXrayPeer) == "" {
		c.DefaultXrayPeer = d.DefaultXrayPeer
	}
	if strings.TrimSpace(c.SocketPath) == "" {
		c.SocketPath = d.SocketPath
	}
	if strings.TrimSpace(c.PidPath) == "" {
		c.PidPath = d.PidPath
	}
	if c.Daemon != 0 {
		c.Daemon = 1
	}
	normalizeLanguage(c)
}

// DaemonEnabled reports whether config.toml requests routing commands through the daemon (when it is up).
func (c *Cfg) DaemonEnabled() bool {
	return c != nil && c.Daemon != 0
}

// ApplyDefaults sets C to built-in defaults without reading config.toml (used for z-panel --ssh=… / --ssh-connect=… when config may be unreadable locally).
func ApplyDefaults() {
	c := defaults()
	C = &c
}

func normalizeLanguage(c *Cfg) {
	s := strings.ToLower(strings.TrimSpace(c.Language))
	if s == "" || s == "auto" {
		c.Language = "auto"
		return
	}
	if l, ok := i18n.CanonicalLanguage(s); ok {
		c.Language = string(l)
		return
	}
	c.Language = "auto"
}

// Write saves config to disk (requires root).
func Write(c Cfg) error {
	if err := persistConfigToDisk(c); err != nil {
		return err
	}
	return reloadFromDiskIntoC()
}

// InitInteractive prompts and writes config.toml.
// If force=false and the file exists: Load() applies schema migration if needed, then prints init_exists.
func InitInteractive(stdin io.Reader, stdout io.Writer, force bool) error {
	if err := root.Require(); err != nil {
		return err
	}
	path := config.ConfigFile
	if !force {
		if _, err := os.Stat(path); err == nil {
			if err := Load(); err != nil {
				return err
			}
			fmt.Fprintf(stdout, i18n.T("settings.init_exists"), path)
			i18n.ApplyFromConfig(C.Language)
			return nil
		}
	}
	if err := Load(); err != nil {
		return err
	}
	d := *C
	r := bufio.NewReader(stdin)

	fmt.Fprintln(stdout, i18n.T("settings.init_intro"))
	fmt.Fprintln(stdout)

	d.Language = prompt(r, stdout, i18n.T("settings.prompt.language"), d.Language)
	normalize(&d)
	i18n.ApplyFromConfig(d.Language)

	d.Table = prompt(r, stdout, i18n.T("settings.prompt.table"), d.Table)
	d.SystemdNetworkDir = prompt(r, stdout, i18n.T("settings.prompt.systemd_network"), d.SystemdNetworkDir)
	d.DefaultLANCIDR = prompt(r, stdout, i18n.T("settings.prompt.lan_cidr"), d.DefaultLANCIDR)
	d.DefaultLANDev = prompt(r, stdout, i18n.T("settings.prompt.lan_dev"), d.DefaultLANDev)
	d.DefaultXrayAddr = prompt(r, stdout, i18n.T("settings.prompt.xray_addr"), d.DefaultXrayAddr)
	d.DefaultXrayPeer = prompt(r, stdout, i18n.T("settings.prompt.xray_peer"), d.DefaultXrayPeer)
	d.UfwMarker = prompt(r, stdout, i18n.T("settings.prompt.ufw_marker"), d.UfwMarker)
	d.XrayTunManagedMark = prompt(r, stdout, i18n.T("settings.prompt.xray_mark"), d.XrayTunManagedMark)

	d.SchemaVersion = CurrentSchemaVersion
	normalize(&d)
	if err := Write(d); err != nil {
		return err
	}
	i18n.ApplyFromConfig(C.Language)
	fmt.Fprintf(stdout, i18n.T("settings.saved"), path)
	return nil
}

func prompt(r *bufio.Reader, out io.Writer, label, def string) string {
	fmt.Fprintf(out, "%s [%s]: ", label, def)
	line, err := r.ReadString('\n')
	if err != nil {
		return def
	}
	line = strings.TrimSpace(line)
	if line == "" {
		return def
	}
	return line
}

// ConfigPath returns the config path for display.
func ConfigPath() string {
	return filepath.Clean(config.ConfigFile)
}
