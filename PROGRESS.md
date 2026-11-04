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

- [x] `ex1-errors-division/main.go` — math operations return (int, error), div by zero returns error
- [x] `ex2-file-len/main.go` — fileLen(filename) (int, error), uses defer to close file
- [x] `ex3-prefixer/main.go` — prefixer(string) returns func(string) string

**Mistakes log:** Confused by functions returning other functions (closures) in ex3-prefixer — understood after stepping through signature and return type.

## ch6 — Pointers (Learning Go book, Chapter 6)

- [x] `ex1-person-pointers/main.go` — Person struct, MakePerson and MakePersonPointer functions
- [x] `ex2-slice-pointers/main.go` — UpdateSlice and GrowSlice functions, print before/after
- [x] `ex3-gc-tuning/main.go` — []Person with 10M entries, test GOGC effects

**Mistakes log:** Initially confused why GrowSlice didn't persist to caller — learned slice headers are copied by value; append reallocation means caller must reassign. GOGC=500 fastest (fewer GCs), GOGC=10 slowest (most GCs).

## ch7 — Types, Methods, and Interfaces (Learning Go book, Chapter 7)

- [x] `ex1-basketball-league/main.go` — Team and League types, 4 teams populated, printed
- [x] `ex2-league-methods/main.go` — MatchResult and Ranking methods on League, both working
- [ ] `ex3-ranker-interface/main.go` — Ranker interface and RankPrinter function

**Mistakes log:** Need to check how to do ex2 again with the confusing slices etc.

## DSA — Data Structures & Algorithms

- [x] `twosum/main.go` — `twoSum(nums []int, target int) []int`, map[int]int one-pass hash map
- [x] `contains-duplicate/main.go` — `containsDuplicate(nums []int) bool`, map[int]bool set
- [x] `valid-anagram/main.go` — `isAnagram(s, t string) bool`, map[rune]int frequency count
- [x] `first-unique-character/main.go` — `firstUniqChar(s string) int`, map[rune]int + second pass over s

**Mistakes log:** Removed `strings.ToLower()` from `firstUniqChar` — it changed the answer by collapsing case (`aA` returned `-1` instead of `0`). Then re-added `ToLower` so `'a'` and `'A'` are treated as same character.

## Projects

- [x] `projects/text-analyzer` — word frequency counter (reads any file via -file flag, top N via -top flag, gofmt'd)
- [ ] `projects/todo-app` — CLI todo with JSON storage (ch4–13, skeleton done; task.go has Task struct with Complete() and TaskList with Delete(id) methods)
- [x] `projects/cli-calc` — CLI integer calculator (done, ch1–5)

## Next session (Monday)

First task: Implement `ch7/ex3-ranker-interface/main.go` — Ranker interface and RankPrinter function. Then continue with todo-app CLI commands (add, list, done, delete).

## Repo

- [x] Git initialized
- [x] Remote: `origin` → `https://github.com/Bertin-Dreyer/go-practice.git`
- [x] Pushed to GitHub
