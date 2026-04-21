package root

import (
	"fmt"
	"os"

	"z-panel/internal/executil"
	"z-panel/internal/i18n"
)

func Require() error {
	if executil.RemoteSSHHost() != "" {
		return nil
	}
	if os.Geteuid() != 0 {
		return fmt.Errorf("%s", i18n.T("root.need_root"))
	}
	return nil
}
