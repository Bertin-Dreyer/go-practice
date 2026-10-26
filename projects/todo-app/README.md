# CLI Todo App

A command-line todo manager that persists tasks to a JSON file.

## Planned chapters

- **Ch4**: pointers, methods, interfaces, structs with receivers
- **Ch5**: functions, multiple returns, `defer`
- **Ch6**: pointers (needed to mutate tasks in place)
- **Ch9**: custom errors (`ErrTaskNotFound`)
- **Ch13**: `encoding/json`, `os.Open`/`os.Create`, `bufio`, `flag`

## Planned features

- `add <task>` — append a todo
- `list` — print all todos with done/pending status
- `done <id>` — mark a todo complete
- `delete <id>` — remove a todo
- Storage: `todos.json` in the working directory

## Portfolio upgrades (after MVP works)

- Table-driven unit tests for every function (`*_test.go`)
- `encoding/json` persistence with atomic writes (write temp file, then `os.Rename`)
- `flag` package for subcommand parsing instead of raw `os.Args`
- Custom error type: `ErrTaskNotFound`
- `Task` type with methods, e.g. `func (t *Task) MarkDone()`

## Status

Not implemented yet. Build the in-memory version after Ch4–Ch6. Add JSON persistence after Ch13.
