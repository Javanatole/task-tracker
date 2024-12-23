# Task CLI

[Task trackers](https://roadmap.sh/projects/task-tracker)

A simple Command-Line Interface (CLI) tool for managing tasks efficiently. With task-cli, you can add, update, delete, and manage the status of your tasks directly from the terminal.

---

## Usage

### Adding a New Task
Create a new task with a description.
```bash
task-cli add "Buy groceries"
```
**Output**:  
`Task added successfully (ID: 1)`

---

### Updating a Task
Update the description of an existing task by its ID.
```bash
task-cli update 1 "Buy groceries and cook dinner"
```

---

### Deleting a Task
Remove a task permanently by its ID.
```bash
task-cli delete 1
```

---

### Marking Task Status
Change the status of a task by its ID.

- Mark a task as **in progress**:
  ```bash
  task-cli mark-in-progress 1
  ```

- Mark a task as **done**:
  ```bash
  task-cli mark-done 1
  ```

---

### Listing Tasks
View tasks in your list.

- List **all tasks**:
  ```bash
  task-cli list
  ```

- List tasks by **status**:
  ```bash
  task-cli list done
  task-cli list todo
  task-cli list in-progress
  ```

---

## Examples

1. **Add and manage tasks**:
   ```bash
   task-cli add "Read a book"
   task-cli mark-in-progress 1
   task-cli mark-done 1
   ```

2. **List and filter tasks**:
   ```bash
   task-cli list
   task-cli list done
   ```
