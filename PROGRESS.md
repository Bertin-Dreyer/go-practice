# Go Practice Progress Tracker

## ch1 — Hello World (Learning Go book, Chapter 1)

- [x] `ch1/go.mod` — module `hello-world`, Go 1.27.3
- [x] `ch1/hello.go` — prints `Hello, World!`
- [x] `ch1/makefile` — fmt / vet / build targets

**Mistakes log:** No mistakes today.

## ch2 — Exercises (Learning Go book, Chapter 2)

- [x] `ex1-int-to-float/main.go` — `i := 20`, `f := float64(i)`, print both
- [x] `ex2-untyped-constant/main.go` — `const value = 20`, assign to int `i` and float `f`, print both
- [x] `ex3-max-overflow/main.go` — `b` (byte=255), `smallI` (int32=max), `bigI` (uint64=max), add 1 to each, print (overflow)

**Mistakes log:** No mistakes today.

## ch3 — Composite Types (Learning Go book, Chapter 3)

- [x] `ex1-greetings-slices/main.go` — slice of five greetings, three subslices, all printed
- [x] `ex2-fourth-rune/main.go` — string with male/female emojis, print fourth rune as character
- [x] `ex3-employee-struct/main.go` — Employee struct, three instance styles, all printed

**Mistakes log:** No mistakes today.

## ch4 — Control Structures (Learning Go book, Chapter 4)

- [x] `ex1-random-slice/main.go` — 100 random numbers in [0, 100] into an int slice
- [x] `ex2-divisibility-rules/main.go` — divisibility by 2, 3, both (Six!), or neither
- [x] `ex3-total-shadowing/main.go` — variable shadowing with `:=` inside a loop

**Mistakes log:** No mistakes today.

## ch5 — Functions (Learning Go book, Chapter 5)

- [ ] `ex1-errors-division/main.go` — math operations return (int, error), div by zero returns error
- [ ] `ex2-file-len/main.go` — fileLen(filename) (int, error), uses defer to close file
- [ ] `ex3-prefixer/main.go` — prefixer(string) returns func(string) string

**Mistakes log:** Not started yet.

## ch6 — Pointers (Learning Go book, Chapter 6)

- [ ] `ex1-person-pointers/main.go` — Person struct, MakePerson and MakePersonPointer functions
- [ ] `ex2-slice-pointers/main.go` — UpdateSlice and GrowSlice functions, print before/after
- [ ] `ex3-gc-tuning/main.go` — []Person with 10M entries, test GOGC effects

**Mistakes log:** Not started yet.

## ch7 — Types, Methods, and Interfaces (Learning Go book, Chapter 7)

- [ ] `ex1-basketball-league/main.go` — Team and League types
- [ ] `ex2-league-methods/main.go` — MatchResult and Ranking methods on League
- [ ] `ex3-ranker-interface/main.go` — Ranker interface and RankPrinter function

**Mistakes log:** Not started yet.

## Projects

- [x] `projects/text-analyzer` — word frequency counter (reads bible.txt, top 20, gofmt'd)
- [ ] `projects/todo-app` — CLI todo with JSON storage (ch4–13, skeleton done)
- [ ] `projects/cli-calc` — CLI integer calculator (ch1–5, scaffold done)

## Repo

- [x] Git initialized
- [x] Remote: `origin` → `https://github.com/Bertin-Dreyer/go-practice.git`
- [x] Pushed to GitHub
