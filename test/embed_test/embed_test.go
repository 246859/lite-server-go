package embed_test

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"os"
	"testing"
)

//go:embed defaultConfig.yml
var defaultConfigString string

func TestEmbedString(t *testing.T) {
	dist, err := os.Create("./config2.yml")
	fmt.Println(err)
	io.Copy(dist, bytes.NewBufferString(defaultConfigString))
}

//go:embed defaultConfig.yml
var defaultConfigBytes []byte

func TestEmbedBytes(t *testing.T) {
	dist, err := os.Create("./config2.yml")
	fmt.Println(err)
	io.Copy(dist, bytes.NewReader(defaultConfigBytes))
}

//go:embed defaultConfig.yml
var defaultConfigFS embed.FS

func TestEmbedFS(t *testing.T) {
	dist, err := os.Create("./config2.yml")
	fmt.Println(err)
	file, err := defaultConfigFS.Open("defaultConfig.yml")
	fmt.Println(err)
	io.Copy(dist, file)
}
