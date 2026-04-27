# 🐹 Go Language Basics for Beginners

If you are new to Go (Golang), this guide explains the syntax and concepts we used in the **GoTask** project.

## 1. Packages and Imports
Every file in Go starts with `package`. 
- `package main`: Tells Go this is a program that can be RUN.
- `package task`: Tells Go this is a "Library" or "Module" that provides features to other parts of the app.
- `import`: Brings in code from the Standard Library (like `fmt` for printing) or external libraries (like `cobra`).

## 2. Variables and Types
Go is "Statically Typed," meaning every variable has a specific type.
- `string`: Text.
- `int`: Whole numbers.
- `bool`: True/False.
- `time.Time`: Date and time.

### The Walrus Operator `:=`
Inside functions, we use `:=` to create and initialize a variable in one step. Go "guesses" the type for you.
```go
name := "John" // Go knows this is a string
```

## 3. Structs (The "Template")
A `struct` is a collection of fields. Think of it like a row in a spreadsheet.
```go
type User struct {
    Name string
    Age  int
}
```

## 4. Slices (The "Dynamic List")
A `slice` is like an array but it can grow or shrink. We write it as `[]Type`.
- `[]Task` is a list of Task objects.
- We use `append(list, item)` to add things to it.

## 5. Pointers (`*` and `&`)
A pointer points to the memory address of a value instead of the value itself.
- `&item`: Get the address (The "Link").
- `*Type`: The type is a pointer (The "Receiver").
**Why?** If you pass a big object to a function, Go makes a copy. If you pass a pointer, Go uses the original. This is faster and allows you to change the original.

## 6. Error Handling
In most languages, you use "Try/Catch". In Go, functions return the error as a second value.
```go
data, err := os.ReadFile("file.txt")
if err != nil {
    // Something went wrong! Handle it here.
    return err
}
```
This makes code very explicit and easy to debug.

## 7. Interfaces (The "Contract")
An interface defines BEHAVIOR.
If you have a `Store` interface with a `Save()` function, it doesn't matter if you save to a File, a Database, or the Cloud. As long as the code has a `Save()` function, it "satisfies" the interface.

## 8. JSON Tags
The little text in backticks `` `json:"id"` `` tells Go how to rename fields when turning them into JSON text for files or APIs.
