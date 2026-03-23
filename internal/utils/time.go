package utils

import (
	"fmt"
	"time"
)

func MeasureTimeErr(name string, fn func() error) error {
	start := time.Now()
	err := fn()
	fmt.Printf("[%s] took %s\n", name, time.Since(start))
	return err
}
