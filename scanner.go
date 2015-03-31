package main

import (
	"fmt"
	"github.com/kraii/id3v1"
	"github.com/kraii/id3v2"
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
		defer file.Close()
		tag := readATag(file)
		newPath := fmt.Sprintf("./%s/%s/%s.mp3", tag.Artist(), tag.Album(), tag.Title())
		fmt.Println(normalise(newPath))
	}
}

func normalise(path string) string {
	r := strings.NewReplacer(" ", "_")
	return r.Replace(path)
}

func readATag(file *os.File) Tag {
	tagV2, _ := id3v2.ReadTag(file)
	if tagV2.Artist() != "" {
		return &tagV2
	} else {
		tagV1, _ := id3v1.ReadTag(file)
		return &tagV1
	}
}

type Tag interface {
	Artist() string
	Album() string
	Title() string
}
