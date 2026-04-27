package task

// Store is an "Interface". 
// Think of an interface as a "Contract". 
// It says: "Any piece of code that wants to be a 'Store' MUST provide these 5 functions."
//
// Why do we do this?
// Today we store tasks in a JSON file. Tomorrow we might want to use a Database (PostgreSQL).
// Because we use an interface, the rest of our app doesn't care WHERE the data goes,
// as long as the "Contract" is followed.
type Store interface {
	Load() ([]Task, error) // Get all tasks from storage
	Save([]Task) error     // Save a full list of tasks
	Add(Task) error        // Add one new task
	Update(Task) error     // Update an existing task
	Delete(int) error      // Delete a task by its ID number
}
