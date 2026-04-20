package ufw

import (
	"testing"
)

func TestRun_help(t *testing.T) {
	c := New()
	if err := c.Run([]string{"help"}); err != nil {
		t.Fatal(err)
	}
}
