package utils

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func MeasureTimeErr(name string, fn func() error) error {
	start := time.Now()
	err := fn()
	prefix := fmt.Sprintf("%s ended for:", name)
	coloredPrefix := color.CyanString(prefix)
	fmt.Printf("%s %s\n", coloredPrefix, time.Since(start))
	return err
}
