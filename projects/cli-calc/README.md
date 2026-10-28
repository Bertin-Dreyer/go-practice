# CLI Calculator

A command-line integer calculator that demonstrates Go functions as first-class values,
multiple return values, and error handling — chapters 1–5 of Learning Go.

## Features

- `add`, `subtract`, `multiply`, `divide` two integers
- Division returns an error on divide-by-zero
- Operators stored in a `map[string]opFunc` (function values in a map)

## Usage

```bash
cd projects/cli-calc
go run main.go 10 + 5   # prints 15
go run main.go 10 / 0   # prints "division by zero" and exits 1
go run main.go 7 % 3    # prints "unsupported operator: %" and exits 1
```

## What it demonstrates (ch1–5)

- **Ch1–2:** variables, constants, basic types, type conversion (`strconv.Atoi`)
- **Ch3:** slices (`os.Args`), maps (`opMap`)
- **Ch4:** control structures (`if/else`, `for` implied by os.Args handling)
- **Ch5:** functions, multiple returns `(int, error)`, function types (`opFunc`),
  functions as values in a map, `defer`-ready error handling pattern

## Status

Implemented. Ready for review.
