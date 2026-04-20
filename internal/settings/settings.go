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
	"z-panel/internal/i18n"
	"z-panel/internal/root"
)

// Cfg is stored in config.toml.
type Cfg struct {
	Table              string `toml:"table"`
	StateDir           string `toml:"state_dir"`
	SystemdNetworkDir  string `toml:"systemd_network_dir"`
	UfwMarker          string `toml:"ufw_marker"`
	XrayTunManagedMark string `toml:"xray_tun_managed_mark"`
	DefaultLANCIDR     string `toml:"default_lan_cidr"`
	DefaultLANDev      string `toml:"default_lan_dev"`
	DefaultXrayAddr    string `toml:"default_xray_addr"`
	DefaultXrayPeer    string `toml:"default_xray_peer"`
	// Language: auto (use locale / Z_PANEL_LANG), en, or ru.
	Language string `toml:"language"`
	// SchemaVersion is written by the program; 0 in file means legacy (treated as 1).
	SchemaVersion int `toml:"schema_version"`
}

// C is the loaded configuration; non-nil after Load().
var C *Cfg

func defaults() Cfg {
	return Cfg{
		Table:              "51843",
		StateDir:           filepath.Join(config.ConfigDir, "state"),
		SystemdNetworkDir:  "/etc/systemd/network",
		UfwMarker:          "z-panel",
		XrayTunManagedMark: "# z-panel-managed",
		DefaultLANCIDR:     "192.168.0.0/16",
		DefaultLANDev:      "lan0",
		DefaultXrayAddr:    "10.252.0.1/30",
		DefaultXrayPeer:    "10.252.0.2/30",
		Language:           "auto",
		SchemaVersion:      0,
	}
}

// Load reads config.toml; if missing, only defaults apply.
func Load() error {
	path := config.ConfigFile
	base := defaults()
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			c := base
			C = &c
			return nil
		}
		return fmt.Errorf(i18n.T("settings.err.read"), path, err)
	}
	if err := toml.Unmarshal(b, &base); err != nil {
		return fmt.Errorf(i18n.T("settings.err.parse"), path, err)
	}
	normalize(&base)
	C = &base
	return nil
}

func normalize(c *Cfg) {
	d := defaults()
	if strings.TrimSpace(c.Table) == "" {
		c.Table = d.Table
	}
	if strings.TrimSpace(c.StateDir) == "" {
		c.StateDir = d.StateDir
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
	normalizeLanguage(c)
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
		os.Remove(tmp)
		return fmt.Errorf(i18n.T("settings.err.write"), path, err)
	}
	return Load()
}

// InitInteractive prompts and writes config.toml.
// If force=false and the file exists: runs schema migration when needed, otherwise prints init_exists.
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
			stored := EffectiveStoredSchema(C.SchemaVersion)
			if stored < CurrentSchemaVersion {
				return runSchemaMigration(stdin, stdout, path, true)
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
	d.StateDir = prompt(r, stdout, i18n.T("settings.prompt.state_dir"), d.StateDir)
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
