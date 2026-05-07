package settings

import (
	"fmt"
	"io"
	"os"

	"z-panel/internal/config"
	"z-panel/internal/i18n"
	"z-panel/internal/root"
)

// CurrentSchemaVersion must be incremented whenever new user-facing keys are added
// to [Cfg] that should be prompted on upgrade. Register the matching step in schemaUpgradesAuto.
const CurrentSchemaVersion = 4

// EffectiveStoredSchema maps on-disk schema_version: 0 / missing → 1 (legacy configs).
func EffectiveStoredSchema(v int) int {
	if v < 1 {
		return 1
	}
	return v
}

// schemaUpgradesAuto must contain an entry for every version V with 1 < V <= CurrentSchemaVersion.
// Each step mutates [Cfg] in place; interactive prompts are not used — migration runs on Load().
var schemaUpgradesAuto = map[int]func(*Cfg) error{
	2: schemaUpgradeAutoV2,
	3: schemaUpgradeAutoV3,
	4: schemaUpgradeAutoV4,
}

func schemaUpgradeAutoV2(*Cfg) error {
	// Previously interactive language prompt; existing file values are kept as-is.
	return nil
}

func schemaUpgradeAutoV3(*Cfg) error {
	// xray-redirect state: legacy JSON under /etc/z-panel/state/ is merged in mergeLegacyStateFiles on read.
	return nil
}

func schemaUpgradeAutoV4(*Cfg) error {
	// no_banner, ssh_no_multiplex, ssh_no_tty — defaults applied in normalize; z-panel env vars removed.
	return nil
}

func applySchemaUpgradesAuto(c *Cfg, storedEffective int) error {
	for v := storedEffective + 1; v <= CurrentSchemaVersion; v++ {
		fn, ok := schemaUpgradesAuto[v]
		if !ok {
			return fmt.Errorf("z-panel: missing schema upgrade step for version %d (program bug)", v)
		}
		if err := fn(c); err != nil {
			return err
		}
	}
	return nil
}

// MigrateInteractive ensures config.toml is loaded; schema upgrades are applied automatically in Load().
func MigrateInteractive(_ io.Reader, stdout io.Writer) error {
	if err := root.Require(); err != nil {
		return err
	}
	path := config.ConfigFile
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("%s", i18n.T("settings.migrate_no_file", path))
	}
	if err := Load(); err != nil {
		return err
	}
	if SchemaJustAutoMigrated() {
		fmt.Fprint(stdout, i18n.T("settings.migrate_completed", CurrentSchemaVersion))
	} else {
		fmt.Fprintln(stdout, i18n.T("settings.migrate_uptodate"))
	}
	i18n.ApplyFromConfig(C.Language)
	return nil
}
