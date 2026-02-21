# Task Tracker CLI

A simple command-line task tracker built in Go.

The application allows you to manage tasks directly from the terminal and stores them in a JSON file in the current directory.

---

## Features

- Add tasks
- Update tasks
- Delete tasks
- Mark tasks as:
  - `todo`
  - `in-progress`
  - `done`
- List all tasks
- List tasks filtered by status
- Automatic JSON file creation
- No external libraries (standard library only)

---

## Project Structure

task-tracker/
├── cmd/
│   └── task-cli/
│       └── main.go
├── internal/
│   └── task/
│       ├── model.go
│       ├── service.go
│       └── storage.go
├── go.mod
└── README.md

### Architecture

The project follows a layered structure:

- **CLI Layer** (`cmd/task-cli`) — handles command-line arguments.
- **Service Layer** (`internal/task/service.go`) — contains business logic.
- **Storage Layer** (`internal/task/storage.go`) — handles file system operations and JSON persistence.
- **Model Layer** (`internal/task/model.go`) — defines task structure and statuses.

---

## Task Model

Each task contains:

- `id` — unique integer identifier
- `description` — task description
- `status` — `todo`, `in-progress`, or `done`
- `createdAt` — creation timestamp
- `updatedAt` — last update timestamp

Tasks are stored in `tasks.json` in the current directory.

---

## Build

From the project root:

```bash
go build -o task-cli ./cmd/task-cli

This creates an executable file:
./task-cli

Usage
Add a task
./task-cli add "Buy groceries"

Output:

Task added successfully (ID: 1)

Update a task
./task-cli update 1 "Buy groceries and cook dinner"

Delete a task
./task-cli delete 1

Mark task status
./task-cli mark-in-progress 1
./task-cli mark-done 1

List tasks

List all:
./task-cli list

List by status:
`./task-cli list done
./task-cli list todo
./task-cli list in-progress`

Error Handling
	•	Invalid commands are handled gracefully.
	•	Missing arguments return informative messages.
	•	Non-existent task IDs return an error.
	•	The JSON file is automatically created if it does not exist.

⸻

Implementation Constraints
	•	Standard library only
	•	Positional CLI arguments
	•	Native file system operations
	•	JSON file persistence

⸻

Future Improvements
	•	Input validation improvements
	•	Sorting by date
	•	Concurrency safety
	•	Unit tests
	•	Configurable storage path
	•	Support for flags