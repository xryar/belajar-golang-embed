package belajargolangembed

import (
	_ "embed"
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
