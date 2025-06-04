module main

// Import statement
import os

// Constants
const (
    pi = 3.14159
    app_name = 'V Language Interpreter Example'
)

// Simple function
fn add(x int, y int) int {
    return x + y
}

// Function with option type return
fn find_user(id int) ?string {
    if id == 1 {
        return 'admin'
    }
    return none
}

// Struct definition
struct User {
    id int
    name string
mut:
    is_active bool
}

// Struct method
fn (u User) display() string {
    status := if u.is_active { 'active' } else { 'inactive' }
    return '${u.name} (${u.id}) - ${status}'
}

// Enum definition
enum Color {
    red
    green
    blue
    custom = 100
}

// Main function
fn main() {
    // Variable declaration
    mut count := 0
    name := 'V Language'
    
    // Print statement
    println('Hello from ${name}!')
    println('Constants: pi=${pi}, app_name=${app_name}')
    
    // Using functions
    sum := add(5, 3)
    println('5 + 3 = ${sum}')
    
    // For loop with range
    for i := 0; i < 3; i++ {
        println('Loop ${i}')
        count++
    }
    
    // For-in loop with array
    numbers := [1, 2, 3, 4, 5]
    for num in numbers {
        println('Number: ${num}')
    }
    
    // Match statement (V's switch)
    match count {
        0 { println('Count is zero') }
        1 { println('Count is one') }
        3 { println('Count is three') }
        else { println('Count is something else: ${count}') }
    }
    
    // Using option types
    user := find_user(1) or {
        println('User not found')
        return
    }
    println('Found user: ${user}')
    
    // Map
    mut scores := map[string]int{}
    scores['Alice'] = 95
    scores['Bob'] = 88
    scores['Charlie'] = 92
    
    // For-in with map
    for name, score in scores {
        println('${name}: ${score}')
    }
    
    // Struct usage
    mut admin := User{
        id: 1
        name: 'Admin'
        is_active: true
    }
    
    println(admin.display())
    
    // Using enum
    color := Color.blue
    println('Selected color: ${color}')
}
