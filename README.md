# go-practice

Learning Go in depth.

This repo contains my Go practice exercises, organized by chapter, as I work through learning the language thoroughly.

## Structure

```
ch1/    - Hello World basics (done)
ch2/    - Types and variables exercises (done)
ch3/    - Composite types (tomorrow)
projects/
  text-analyzer/  - Word frequency counter (planned, ch1–3)
  todo-app/       - CLI todo manager (planned, ch4–13)
```

## Getting Started

Each chapter is its own Go module.

```bash
cd ch1
make fmt
make vet
make build
go run hello.go
```

- `make fmt` — format code
- `make vet` — run `go vet`
- `make build` — build the binary

## License

MIT
