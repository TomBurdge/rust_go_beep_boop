package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/ebitengine/purego"
)

func main() {
	var target string

	goos := runtime.GOOS
	goarch := runtime.GOARCH

	switch {
	case goos == "linux" && goarch == "amd64":
		target = "x86_64-unknown-linux-gnu"
	case goos == "linux" && goarch == "arm64":
		target = "aarch64-unknown-linux-gnu"
	case goos == "darwin" && goarch == "arm64":
		target = "aarch64-apple-darwin"
	default:
		panic(fmt.Sprintf("unsupported platform: %s/%s. Only Linux ARM/AMD64 and darwin are supported", goos, goarch))
	}

	libPath := fmt.Sprintf("lib/rust_ffi/target/%s/release/libhello.so", target)

	rustlib, err := purego.Dlopen(libPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		log.Fatalf("failed to load library: %v", err)
	}
	defer purego.Dlclose(rustlib)

	var speak func(*byte) *byte
	purego.RegisterLibFunc(&speak, rustlib, "speak")

	input := []byte("Hello from go\x00")

	_ = speak(&input[0])
}
