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
ch7/    - Types, methods, interfaces exercises (ex1, ex2 done; ex3 skeleton)
DSA/    - Data structures & algorithms (twosum, contains-duplicate, valid-anagram, first-unique-character done)
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
