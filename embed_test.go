package belajargolangembed

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

func TestEmbedString(t *testing.T) {
	fmt.Println(version)
}

//go:embed kucing_tengok_lah.jpg
var image []byte

func TestEmbedImage(t *testing.T) {
	err := os.WriteFile("kucing_new.jpg", image, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/*.txt
var files embed.FS

func TestMulipleEmbed(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := files.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := files.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
