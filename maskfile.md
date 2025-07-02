# Tasks to build and the project

## build
> Builds the rust & go code


```sh
echo 'building rust library'
cd lib/rust_ffi/src
echo 'building for AMD64 linux'
rustup target add x86_64-unknown-linux-gnu
cargo zigbuild --release --target=x86_64-unknown-linux-gnu
echo 'building for Linux ARM'
rustup target add aarch64-unknown-linux-gnu
cargo zigbuild --release --target=aarch64-unknown-linux-gnu
echo 'building for macOS ARM'
rustup target add aarch64-apple-darwin
cargo zigbuild --target=aarch64-apple-darwin
cd ../../..
```

## run
> Runs the go code which calls the built go library
```sh
go get github.com/ebitengine/purego
echo 'running the go code...'
go run main.go
```

## build-run
> Build the rust library & go crate, and run the go code which calls rust.
```sh
mask build && mask run
```
