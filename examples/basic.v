module main

import math
import { println, eprintln }

pub fn main() {
    x := 42
    y := add(x, 10)
    println(y)
}

fn add(a int, b int) int {
    return a + b
}
