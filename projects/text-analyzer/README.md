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

## Sample output (top 5 from bible.txt)

```
the 64113
and 51696
of 34720
to 13598
that 12912
```

## Limits

- **Joined words:** `well—known` (em dash, no spaces) stays one token — `TrimFunc` only trims edges.
- **Digits:** numbers count as words (e.g. `"123"` → `"123"`).
- **Gutenberg boilerplate:** the Project Gutenberg header and licence text are included in the count, so words from that front matter appear alongside the actual book text.
- **Verse labels:** tokens like `1:1` survive trimming because the digits and colon are treated as word characters in places, producing one-off tokens.
- **Ties:** words with equal counts appear in no guaranteed order.

## Usage

```bash
cd projects/text-analyzer
go run main.go
```

Reads `bible.txt` from your current working directory (run from `projects/text-analyzer/`). To analyze a different file, either replace `bible.txt` in this folder or change the filename in `main.go` — command-line args are planned for a later chapter.

