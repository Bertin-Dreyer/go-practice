# CLI Calculator

A command-line integer calculator for `+`, `-`, `*`, and `/`.
Demonstrates Go functions as first-class values, multiple return values,
and error handling — chapters 1–5 of Learning Go.

## Features

- Add, subtract, multiply, divide two integers
- Division returns an error on divide-by-zero
- Operations stored in a `map[string]opFunc`, so adding `%` or `^` is one
  function plus one map entry

## Usage

```bash
cd projects/cli-calc
go run . 6 "*" 3   # prints 18
go run . 10 / 2    # prints 5
go run . 7 / 0     # prints "division by zero" to stderr, exits 1
go run . 7 % 2     # prints "unsupported operator: %" to stderr, exits 1
go run . abc + 1   # prints parse error to stderr, exits 1
```

> **Note:** Quote `*` in Git Bash (or your shell) so it isn't expanded as a glob.

## Design

Operations live in a `map[string]opFunc`. To add a new operator:
1. Write a function with signature `func(int, int) (int, error)`.
2. Add one entry to `opMap`.

## Limits

- Integers only: `7 / 2` gives `3`, not `3.5`.
- Negative numbers work fine.
- No floating-point or decimal support.

## What it demonstrates (ch1–5)

- **Ch1–2:** variables, constants, basic types, type conversion (`strconv.Atoi`)
- **Ch3:** slices (`os.Args`), maps (`opMap`)
- **Ch4:** control structures (`if/else`)
- **Ch5:** functions, multiple returns `(int, error)`, function types (`opFunc`), functions as values in a map
