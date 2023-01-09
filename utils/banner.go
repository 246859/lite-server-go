package utils

import (
	"os"

	"github.com/dimiro1/banner"
)

func LogBanner() {
	reader, err := os.Open("./banner.txt")
	if err != nil {
		return
	}
	banner.Init(os.Stdout, true, true, reader)
}
