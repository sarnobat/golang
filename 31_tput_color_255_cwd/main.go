package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(0)
		return
	}

	// Normalize so the same directory reached via symlink/relative path stays stable.
	if abs, err := filepath.Abs(wd); err == nil {
		wd = abs
	}
	if real, err := filepath.EvalSymlinks(wd); err == nil {
		wd = real
	}

	sum := sha256.Sum256([]byte(wd))
	fmt.Println(uint8(sum[0])) // 0..255
}