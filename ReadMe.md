# lastused

### Created by Vedant Gupta

A simple, minimal CLI tool written in Go to list the most recently modified files in the current directory and subdirectories. It’s fast, no-nonsense, and easy to use — built for developers who just want answers.

---

### 🔧 Features

- Lists all files (recursively) sorted by **last modified time**
- Use `-n` to limit how many results you want to see
- Use `--ext` to filter files by extension (like `.go`, `.py`, `.cpp`)

---

### 🚀 Usage

```bash
lastused                # List all files sorted by last modified time
lastused -n 5           # Show only the 5 most recently modified files
lastused --ext .go      # Show only recently modified .go files