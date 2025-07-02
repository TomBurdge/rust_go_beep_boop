# Comparing Rust and Go deep dive
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

I think that, being able to inter-op them with go as your script/high level and rust as your super safe, super close to the metal system-level glue is a great combination.
