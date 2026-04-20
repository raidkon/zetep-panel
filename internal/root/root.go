package root

import (
	"fmt"
	"os"

	"z-panel/internal/i18n"
)

func Require() error {
	if os.Geteuid() != 0 {
		return fmt.Errorf("%s", i18n.T("root.need_root"))
	}
	return nil
}
