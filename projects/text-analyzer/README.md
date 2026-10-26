# Text Analyzer

Counts word frequencies in a text file and prints the top N most used words.

## Chapters used

- **Ch1**: `go mod init`, `os.ReadFile`, `fmt`, `os.Exit`, basic CLI structure
- **Ch2**: `string`, `int`, `bool`, type conversion (`string(data)`), `:=`, `var`, zero values
- **Ch3**: `map[string]int`, `[]pair` slice, `append`, `len`, `range`, sub-slicing via loop bounds

## Usage

```bash
go run main.go bible.txt 10
```

## Notes

- Punctuation is trimmed from word edges with `strings.Trim`.
- Comparison is case-insensitive (`strings.ToLower`).
- Top-N selection uses a simple selection sort loop (no `sort` package yet — that comes later).
