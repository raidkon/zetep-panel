package ufw

import "os"

func stdoutIsTTY() bool {
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

const (
	ansiReset  = "\033[0m"
	ansiGreen  = "\033[32m"
	ansiYellow = "\033[33m"
	ansiRed    = "\033[31m"
)

func colorGreen(s string) string {
	if !stdoutIsTTY() {
		return s
	}
	return ansiGreen + s + ansiReset
}

func colorYellow(s string) string {
	if !stdoutIsTTY() {
		return s
	}
	return ansiYellow + s + ansiReset
}

func colorRed(s string) string {
	if !stdoutIsTTY() {
		return s
	}
	return ansiRed + s + ansiReset
}
