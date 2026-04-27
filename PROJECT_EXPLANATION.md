# Project Architecture & Thought Process

This document explains how we built **GoTask** and why we made certain engineering decisions.

## 1. Project Structure (The "Standard Layout")
We used a structure common in professional Go projects:
- `cmd/gotask/`: The "Entry Point". It should be very small. Its only job is to start the app.
- `internal/`: Code that belongs ONLY to this project. Other people cannot "import" it.
    - `task/`: The **Domain Logic**. This is the heart of the app. It defines what a "Task" is and how it is stored.
    - `cli/`: The **Presentation Layer**. This handles how the user interacts with the app (commands, flags, printing text).

## 2. The Thought Process (Step-by-Step)

### Step 1: Define the Data (The Noun)
Before writing code, we thought: *What is a Task?* 
It needs an ID, a Title, a Status, and Timestamps. We defined this in `task.go`.

### Step 2: Define the Action (The Verb)
We needed a way to Save and Load tasks. Instead of just writing a file-saver, we created an **Interface** called `Store`. 
**Thinking:** "If I want to change from a JSON file to a Database later, I don't want to rewrite the whole app. I'll make a contract (Interface) so the CLI doesn't care about the storage details."

### Step 3: Implement Persistence (The "How")
We implemented `JSONStore`. 
**Engineering Challenge:** What if the program crashes while writing the file? 
**Solution:** The **Atomic Write Pattern**. 
1. Write to a temporary file (`tasks.123.tmp`).
2. If that works, Rename it to the real file (`tasks.json`).
Renaming is "Instant" at the OS level, so the file never gets corrupted.

### Step 4: Build the Interface (The CLI)
We used **Cobra** to create subcommands like `add` and `list`.
**UX Thinking:** "Users like pretty output." 
**Solution:** We used `text/tabwriter` to ensure columns align perfectly, regardless of how long the task titles are.

### Step 5: Safety First
We used `os.Chmod` to set file permissions to `0600`.
**Thinking:** "Task managers might contain private info. Only the user who owns the computer should be allowed to read the file."

## 3. Key Backend Concepts Used
- **Serialization:** Turning a Go Object into a String (JSON) so it can be saved.
- **Dependency Injection:** We "inject" the Store into the CLI commands.
- **Idempotency:** Designing commands so that running them twice doesn't cause errors (e.g., completing an already complete task).
- **Graceful Error Handling:** Never letting the program "Crash" (panic). We always return an error and print a nice message to the user.
