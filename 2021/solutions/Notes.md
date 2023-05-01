## Day 1 learnings
* std::fs::File for opening file
* std::io::Read for reading file
* std::path::Path for path verification
* v! macro for creating a vector
* .collect() for collecting into a vector
* `for i in 0..vec.len()` for iterating over a vector
* use references to avoid copying (and allow borrowing)
* use std::env::current_dir to get current dir
* -> notation for return type in fsignature


## Day 3 learnings
* u32::from_str_radix for parsing non-decimal strings to numbers
* string::parse<u64> assumed base 10 string number
* use vec::retain to filter vector items based on a condition
* use &[u32] for slice of u32, can be an array or vector. Read but cannot modify the vector.

## Day 4 learnings
* `iter` borrows each element of the collection and leaves the collection untouched
* `into_iter` consumes (moves) the collection, so it cannot be used afterwards
* `iter_mut` borrows the elements of the collection mutably, it can be used afterwards
