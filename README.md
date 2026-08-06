# mapcipher

A text encryption/decryption tool built around a 10×10 coordinate code table. Every character maps to a two-digit coordinate code `xy` (x = row, y = column) — an extended take on the Polybius square, with digits and common punctuation squeezed into the table as well.

## Build

You need Go 1.26.5 or newer. From the project directory:

```bash
go build .
```

This produces a single executable named `mapcrypt` in the current directory. If you just want to try it out without building, `go run .` compiles and runs in one step — the commands below work either way.

## Usage

```bash
# Encrypt: replace every character in a file with its two-digit code, write the result to <file>.enc
go run . encrypt <file>

# Decrypt: restore the original text from the codes, write the result to <file>.txt
go run . decrypt <file>.enc
```

Once built, you can swap `go run .` for `./mapcrypt`, e.g. `./mapcrypt encrypt <file>`.

Example:

```bash
echo -n "I LOVE YOU" > demo.txt
go run . encrypt demo.txt
cat demo.txt.enc
# ["93","82","96","18","60","94","82","02","18","68"]

go run . decrypt demo.txt.enc
cat demo.txt.txt
# I LOVE YOU
```

## Files

| File | Description |
| --- | --- |
| `main.go` | The encryption/decryption program |
| `code.table` | The code table definition (JSON: character → two-digit coordinate code) |
| `go.mod` | Go module configuration |

## Code table

A code `xy` means row `x`, column `y`. The mapping itself lives in `code.table` — it's yours to design (a 10×10 grid works well, e.g. letters, digits, space, and newline). Every character you want to support needs exactly one code, and each code should map to exactly one character.

## Notes

- Supported characters: digits `0-9`, upper/lowercase letters, space, newline, and `! ( ) , .`
- Any character not in the table causes an error and aborts — nothing is silently dropped
- The mapping in `code.table` acts as the key of this cipher, so keep it private if you use this for anything real
