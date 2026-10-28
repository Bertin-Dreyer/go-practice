# Text Analyzer

Counts word frequencies in a text file and prints the top 20 most used words.

## How it works

1. `os.ReadFile("bible.txt")` reads the whole file into a byte slice.
2. `string(data)` converts bytes to a string.
3. `strings.Fields(strings.ToLower(text))` splits on whitespace and lowercases.
4. `strings.TrimFunc` strips leading/trailing non-letter/non-digit characters from each word
   (e.g. `"town!"` → `"town"`, `"state-of-the-art"` stays intact).
5. `map[string]int` counts each cleaned word.
6. Results are dumped into a `[]wordCount` slice, sorted by count (descending) with `sort.Slice`,
   and the top 20 are printed.

## Limits

- **Joined words:** `well—known` (em dash, no spaces) stays one token — `TrimFunc` only trims edges.
- **Digits:** numbers count as words (e.g. `"123"` → `"123"`).

## Usage

```bash
cd projects/text-analyzer
go run main.go
```

Reads `bible.txt` from the same directory. Replace `bible.txt` with any plain-text file to analyze it.
