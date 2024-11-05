package tools

import (
	"fmt"
	"os"
	"strings"
)

func GetNextAvailableFilename(filename string) string {
	if _, err := os.Stat(filename); err != nil {
		return filename
	}

	base := strings.TrimSuffix(filename, ".txt")
	ext := ".txt"
	counter := 1
	for {
		newFilename := fmt.Sprintf("%s%d%s", base, counter, ext)
		if _, err := os.Stat(newFilename); os.IsNotExist(err) {
			return newFilename
		}
		counter++
	}
}
