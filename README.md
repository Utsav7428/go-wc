# wc - Go Command Line Word Counter

A simple Go implementation of the classic `wc` (word count) command-line tool.

This tool can count the number of lines, words, and bytes in a given text file or standard input.

---

## 📦 Features

- Count lines (`-l`)
- Count words (`-w`)
- Count bytes (`-c`)
- Support input from files or stdin

---

## 🚀 Getting Started

### Prerequisites
- Go 1.18 or later

### Installation

You can install directly from source or build the executable.

#### A. Run from source

```bash
go run wc.go -l -w -c test.txt
```

#### B. Build executable

```bash
go build -o wc.exe wc.go
```

Then add `wc.exe` to your system PATH to use it globally.
