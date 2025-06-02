# MyLang Interpreter

A programming language interpreter built with Go and ANTLR4, developed using Test-Driven Development (TDD) methodology.

## Features

- Complete interpreter for a custom programming language (MyLang)
- Support for variables, expressions, control flow, functions, and more
- ANTLR4-based grammar and parser
- Comprehensive test suite with TDD approach
- CLI with file execution and REPL mode

## Project Structure

```
my-lang-interpreter/
├── go.mod            # Go module definition
├── go.sum            # Go module checksums
├── Makefile          # Build automation
├── README.md         # This file
├── grammar/          # ANTLR grammar definition
│   └── MyLang.g4     # Grammar file for MyLang
├── parser/           # Generated ANTLR parser files
├── ast/              # Abstract Syntax Tree definitions
│   ├── nodes.go      # AST node types
│   └── visitor.go    # Visitor pattern implementation
├── interpreter/      # Interpreter implementation
│   ├── interpreter.go  # Main interpreter logic
│   ├── environment.go  # Scoping and environments
│   └── values.go       # Value representation
├── lexer/            # Lexical analysis
│   └── lexer.go      # Custom lexer functionality
├── cmd/              # Command line application
│   └── main.go       # Entry point
├── examples/         # Example MyLang programs
│   ├── basic.mylang
│   ├── conditionals.mylang
│   ├── loops.mylang
│   └── functions.mylang
└── tests/            # Test suite
    ├── parser_test.go
    ├── interpreter_test.go
    ├── statements_test.go
    ├── expressions_test.go
    └── integration_test.go
```

## Language Features

### 1. Basic Statements
- Variable declarations and assignments
- Expression statements
- Print/output statements

### 2. Control Flow
- If-else statements (nested)
- While loops
- Do-while loops
- For loops
- Switch-case statements (with default)
- Break and continue statements

### 3. Expressions
- Arithmetic operations (+, -, *, /, %)
- Comparison operations (==, !=, <, >, <=, >=)
- Logical operations (&&, ||, !)
- Parenthesized expressions

### 4. Data Types
- Numbers (integers and floats)
- Strings
- Booleans
- Arrays/Lists

### 5. Functions
- Function declarations
- Function calls
- Return statements
- Local scope handling

## Installation

### Prerequisites

- Go 1.21 or higher
- Java Runtime Environment (for ANTLR4)
- ANTLR4 tool (downloaded automatically by the Makefile)

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/my-lang-interpreter.git
   cd my-lang-interpreter
   ```

2. Generate the parser code from the grammar:
   ```bash
   make generate
   ```

3. Build the interpreter:
   ```bash
   make build
   ```

## Usage

### Running the Interpreter

#### File Execution
```bash
./bin/mylang path/to/yourfile.mylang
```

#### REPL Mode
```bash
./bin/mylang
```

#### Debug Mode
```bash
./bin/mylang --debug path/to/yourfile.mylang
```

### Example Programs

Run an example program:
```bash
make run-example EXAMPLE=basic
```

## Development

### Testing

Run all tests:
```bash
make test
```

Generate test coverage report:
```bash
make coverage
```

### Clean Build

Clean generated files and rebuild:
```bash
make clean build
```

## Testing Strategy

This project follows Test-Driven Development principles:

1. **Parser Tests**: Validate syntax parsing and AST construction
2. **Statement Tests**: Test execution of different statement types
3. **Expression Tests**: Verify correct evaluation of expressions
4. **Interpreter Tests**: Test the interpreter's execution environment
5. **Integration Tests**: End-to-end tests with complete programs

## License

MIT

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request
