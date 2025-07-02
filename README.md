# Rust Go Beep Boop

Rust and go are both really nice.

Both are memory safe. Go is super iterate and get running. Rust is super quick at runtime. 

Since they are both great, but are both a little better at two different things

Wouldn't it be nice if we could use them together?

It could look something like:
* Go for the higher level glue iteration.
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

## Comparing Rust and Go deep dive
Below assertions, which are true at a very high level, and include some opinion:

* Rust is a super quick language, because of: 
  * Type strictness.
  * A unique memory model, with no garbage collector.
  * Optimisations made by a compiler.
* Go is very quick language, because of:
  * Type strictness.
  * Optimisations made by a compiler.

* Rust is a super safe language at runtime, because of:
  * Type strictness.
  * Memory strictness (memory leaks are basically your fault).
  * Errors as values.
* Go is a very safe language at runtime, because of:
  * Type strictness.
  * A garbage collector (memory leaks are almost impossible with FFI).
  * Errors as values.

* Go is super quick to iterate, because:
  * Go has deliberately limited features. Don't abstract, just ship.
  * Compile times are quick, but you can still be confident it will run.
* With rust, one *can* iterate quickly. This requires either:
  * Initial compromises in either: safety, strictness, soundness, or runtime speed. This is, broadly what is recommended [here](https://corrode.dev/blog/prototyping/)
  * High competence in rust. (A principal engineer at a hedge fund once told me that he would never reach for python to write a script, and can iterate quicker in rust.)

