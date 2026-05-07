package settings

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"z-panel/internal/config"
	"z-panel/internal/i18n"
	"z-panel/internal/root"
)

// CurrentSchemaVersion must be incremented whenever new user-facing keys are added
// to [Cfg] that should be prompted on upgrade. Register the matching step in schemaUpgrades.
const CurrentSchemaVersion = 3

// EffectiveStoredSchema maps on-disk schema_version: 0 / missing → 1 (legacy configs).
func EffectiveStoredSchema(v int) int {
	if v < 1 {
		return 1
	}
	return v
}

// schemaUpgrades must contain an entry for every version V with 1 < V <= CurrentSchemaVersion.
var schemaUpgrades = map[int]func(*Cfg, *bufio.Reader, io.Writer){
	2: schemaUpgradeV2,
	3: schemaUpgradeV3,
}

func schemaUpgradeV2(c *Cfg, r *bufio.Reader, w io.Writer) {
	c.Language = prompt(r, w, i18n.T("settings.prompt.language"), c.Language)
}

func schemaUpgradeV3(_ *Cfg, _ *bufio.Reader, _ io.Writer) {
	// xray-redirect state moved from state_dir JSON files into config.toml [xray_redirect]; merge on Load.
}

func runSchemaMigration(stdin io.Reader, stdout io.Writer, path string, withIntro bool) error {
	d := *C
	r := bufio.NewReader(stdin)
	if withIntro {
		fmt.Fprintln(stdout, i18n.T("settings.migrate_intro"))
		fmt.Fprintln(stdout)
	}
	stored := EffectiveStoredSchema(d.SchemaVersion)
	for v := stored + 1; v <= CurrentSchemaVersion; v++ {
		fn, ok := schemaUpgrades[v]
		if !ok {
			return fmt.Errorf("z-panel: missing schema upgrade step for version %d (program bug)", v)
		}
		fn(&d, r, stdout)
	}
	d.SchemaVersion = CurrentSchemaVersion
	normalize(&d)
	if err := Write(d); err != nil {
		return err
	}
	i18n.ApplyFromConfig(C.Language)
	fmt.Fprintf(stdout, i18n.T("settings.saved"), path)
	return nil
}

// MigrateInteractive updates config.toml when schema_version is older than [CurrentSchemaVersion].
func MigrateInteractive(stdin io.Reader, stdout io.Writer) error {
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
	stored := EffectiveStoredSchema(C.SchemaVersion)
	if stored >= CurrentSchemaVersion {
		fmt.Fprintln(stdout, i18n.T("settings.migrate_uptodate"))
		i18n.ApplyFromConfig(C.Language)
		return nil
	}
	return runSchemaMigration(stdin, stdout, path, true)
}
