# CLI Task Manager

A simple command-line task manager built using **Go**. It allows users to add, view, complete, and delete tasks directly from the terminal.

Tasks are stored locally in a JSON file, ensuring that data persists between program executions.

## Features

- **Add Tasks** — Create new tasks with unique IDs.
- **List Tasks** — View all tasks along with their completion status.
- **Complete Tasks** — Mark existing tasks as completed.
- **Delete Tasks** — Remove tasks using their IDs.
- **JSON Storage** — Automatically save and load tasks from a local JSON file.
- **Error Handling** — Handle invalid task IDs, empty titles, and file errors.

## Tech Stack

- **Language:** Go
- **Storage:** JSON file
- **Libraries:** Go Standard Library (`fmt`, `os`, `strconv`, `strings`, `encoding/json`)

## Project Structure

```text
CLI-Task-Manager/
├── main.go
├── go.mod
├── internal/
│   └── task/
│       ├── task.go
│       └── storage.go
├── data/
│   └── tasks.json
└── README.md
```

## Getting Started

### Prerequisites

Make sure Go is installed on your system.

Check your installation:

```bash
go version
```

### Installation

Clone the repository:

```bash
git clone https://github.com/Srinidhi444/CLI-Task-Manager.git
cd CLI-Task-Manager
```

### Run the Application

Run the application using:

```bash
go run . <command> [arguments]
```

Run all commands from the project root directory.

## Usage

### 1. Add a Task

```bash
go run . add "Learn Go"
go run . add "Build CLI Project"
```

Output:

```text
Task added successfully!
```

### 2. List All Tasks

```bash
go run . list
```

Example output:

```text
Your Tasks:
1. Learn Go [Pending]
2. Build CLI Project [Pending]
```

### 3. Complete a Task

Mark a task as completed using its ID:

```bash
go run . complete 1
```

Output:

```text
Task marked as completed!
```

Running `go run . list` now displays:

```text
Your Tasks:
1. Learn Go [Completed]
2. Build CLI Project [Pending]
```

### 4. Delete a Task

Delete a task using its ID:

```bash
go run . delete 2
```

Output:

```text
Task deleted successfully!
```

### 5. Handle Invalid Input

The application handles invalid commands and task IDs.

For example:

```bash
go run . complete abc
```

Output:

```text
Invalid task ID: please enter a positive integer
```

## Data Persistence

Tasks are stored in `data/tasks.json`.

Example:

```json
[
  {
    "id": 1,
    "title": "Learn Go",
    "completed": true
  },
  {
    "id": 2,
    "title": "Build CLI Project",
    "completed": false
  }
]
```

The application reads existing tasks when it starts and saves changes after adding, completing, or deleting tasks.

**Note:** Create the `data` directory before using the application. The `tasks.json` file can be initialized with `[]`. If the file is missing, the application starts with an empty task list.

## How It Works

1. The application reads command-line arguments using `os.Args`.
2. Existing tasks are loaded from the JSON file using `os.ReadFile` and `json.Unmarshal`.
3. The requested operation is executed.
4. Changes are converted into JSON using `json.MarshalIndent`.
5. The updated tasks are written to the file using `os.WriteFile`.



## What I Learned

This project helped me practice:

- Go structs and slices
- Functions and error handling
- Working with command-line arguments
- JSON marshaling and unmarshaling
- File handling using the `os` package
- Organizing Go code into packages
- Building a command-line application using Go's standard library

## License

No license has been specified yet.