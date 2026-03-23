package utils

import (
	"fmt"
	"os"
	"unicode"
)

func CheckASCIIPath() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	for _, r := range dir {
		if r > unicode.MaxASCII {
			return fmt.Errorf("path contains non-ASCII characters: %s", dir)
		}
	}

	return nil
}
