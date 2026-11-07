# go-practice

Learning Go in depth.

This repo contains my Go practice exercises, organized by chapter, as I work through learning the language thoroughly.

## Structure

```
ch1/    - Hello World basics (done)
ch2/    - Types and variables exercises (done)
ch3/    - Composite types exercises (done)
ch4/    - Blocks, shadows, control structures exercises (done)
ch5/    - Functions exercises (done)
ch6/    - Pointers exercises (done)
ch7/    - Types, methods, interfaces exercises (done)
ch8/    - Generics exercises (done)
ch9/    - Errors exercises (done)
ch10/   - Modules, packages, imports exercises (done)
ch11/   - Go tooling exercises (done)
ch12/   - Concurrency exercises (done)
ch13/   - Standard library exercises (done)
ch14/   - Context exercises (done)
ch15/   - Writing tests exercises (done)
ch16/   - Reflect, unsafe, cgo exercises (done)
projects/
  text-analyzer/  - Word frequency counter (done, reads any file, top N via CLI flags)
  todo-app/       - CLI todo manager (skeleton; task.go has Task struct with Complete() and TaskList.Delete() methods)
  cli-calc/       - CLI integer calculator (done, ch1–5)
```

## Getting Started

Each chapter and project is its own Go module.

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
