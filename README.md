# GoTask — CLI Task Manager 🥉

GoTask is a simple command-line task manager built with Go. This is Project B-01 from my Bronze Tier Backend Engineering Roadmap.

author : sakar zaidan

## Features

- Add tasks with optional descriptions
- List tasks in a table or JSON format
- Mark tasks as completed
- Delete tasks
- Atomic file writes for data integrity
- Secure file permissions (0600)
- Configurable storage path via environment variable `GOTASK_FILE`

## Installation

```bash
make build
# or
go install ./cmd/gotask
```

## Usage

```bash
./gotask add "Buy kdd milk" -d "Need 2L of whole milk"
./gotask list
./gotask list --json
./gotask done 1
./gotask delete 1
```

## Configuration

By default, tasks are stored in `$HOME/.gotask.json`. You can override this by setting the `GOTASK_FILE` environment variable:

```bash
export GOTASK_FILE=/path/to/your/tasks.json
```

## Implementation Details

- **CLI Framework:** [Cobra](https://github.com/spf13/cobra)
- **Storage:** JSON file with atomic write pattern (temp file + rename)
- **Structure:** Follows Go idiomatic project layout (`cmd/`, `internal/`)
