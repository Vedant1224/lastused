# lastused

**`lastused`** is a fast, minimal command-line tool written in Go that helps you quickly list the most recently modified files in your current directory (and all subfolders). It’s cleaner and more focused than `ls -lt`, especially when you want to track recent changes in your project.

---

##  Features

- Recursively scans directories and subfolders  
- Sorts files by **last modified time** - Optional `-n` flag to show only the top N files  
- Optional `--ext` flag to filter by file extension (for example: `.go`, `.py`, `.md`)

---

##  Installation (From Source)

```bash
git clone [https://github.com/Vedant1224/lastused.git](https://github.com/Vedant1224/lastused.git)
cd lastused
go build -o lastused
sudo mv lastused /usr/local/bin/
```

## Usage
```bash
lastused              # Show all files sorted by last modified time
lastused -n 5         # Show the 5 most recently modified files
lastused --ext .go    # Show only recently modified .go files

```

## Example Output
```bash
main.go                                — modified on 2025-06-22 14:35:02
README.md                              — modified on 2025-06-22 14:33:10
go.mod                                 — modified on 2025-06-22 14:30:01

```