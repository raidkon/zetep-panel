package daemon

import (
	"context"
	"log"
	"time"

	"z-panel/internal/settings"
)

// Loop periodically reloads config and is the hook for reconciling kernel state
// with saved state (TODO: compare state/*.json vs ip rule/route and re-apply if needed).
func Loop(ctx context.Context) {
	log.SetPrefix("z-panel: ")
	t := time.NewTicker(30 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			if err := settings.Load(); err != nil {
				log.Printf("watch: reload config: %v", err)
				continue
			}
			// Future: enumerate state files, verify policy routing matches, optionally repair.
		}
	}
}
