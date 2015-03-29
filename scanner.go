package main

import (
	"fmt"
	"github.com/kraii/id3v1"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Printf("Using %q", os.Args[1])
	filepath.Walk(os.Args[1], walky)
}

func walky(path string, info os.FileInfo, err error) error {
	if strings.HasSuffix(path, ".mp3") {
		readTag(path)
	}
	return nil
}

func readTag(path string) {
	file, err := os.Open(path)
	if err == nil {
		tag, err := id3v1.ReadTag(file)
		if err == nil {
			fmt.Printf("artist=%q\n", tag.Artist())
		}
	}
}
