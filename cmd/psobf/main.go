package main

import (
	"fmt"
	"os"

	"github.com/Doom-tech/psobf/internal/obfuscator"
)

func main() {
	opts := obfuscator.ParseFlags()
	if err := obfuscator.Run(opts); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
