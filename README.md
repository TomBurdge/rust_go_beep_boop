# Rust Go Beep Boop

[Rust](https://www.rust-lang.org/) and [go](https://go.dev/) are both really nice.

Both are memory safe. Go is super quick to iterate and get running. Rust is super quick at runtime (and go isn't too shabby).

Since they are both great, but are both a little better at two different things

Wouldn't it be nice if we could use them together?

It could look something like:
* Go for the higher level code glue, iteration, and scripts.
* Rust for performance critical operations.

## This is possible!
Rust and go can inter-op via the [C FFI](https://doc.rust-lang.org/nomicon/ffi.html).

This project involves a basic example of passing strings from go to rust, to call a rust function.

This project is an implementation of the pattern from [this blog post](https://blog.arcjet.com/calling-rust-ffi-libraries-from-go/). There are excellent resources linked within the blog post, plus [this mdbook](https://github.com/Michael-F-Bryan/rust-ffi-guide) and [this](https://github.com/thebracket/RustNations2025) are particularly good for understanding the rust C interface.

This project uses [purego](https://github.com/ebitengine/purego), rather than [cgo](https://go.dev/wiki/cgo). purego performs syscall magic so that we can use the C FFI interface without requiring libc.

[WASM](https://blog.arcjet.com/webassembly-on-the-server-compiling-rust-to-wasm-and-executing-it-from-go/) for the interface also looks exciting, but beyond the scope of this project.

## Running
There is mask.md file that can be used to re-create and run this example.

Mask is a command line or script runner, with *some* similarities to make, which also serves as accessible documentation by being in markdown.

The task runner will build the rust crate for: Linux x86, Linux ARM, and macOS ARM.

You must have the below requirements.

Requirements:
* Linux x86, Linux ARM, macOS ARM device. (Windows would be possible, it's just not implemented here.)
* Cargo installed.
* Zig installed.
* Cargo zigbuild installed. You can install with: `cargo install cargo-zigbuild`
* [mask installed](https://github.com/jacobdeichert/mask#installation).

Then, to run the example, run:
`mask build-run`

The code will compile, and you will see printed to console:
`beep boop I got your string. here it is: Hello from go!`
Rust has been printed this to console, after receiving and parsing a 'Hello from go' string via the C FFI. 


