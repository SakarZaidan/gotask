/*
This is the "Entry Point" of the application.
When you run "go run main.go" or execute the compiled binary, 
this is the first piece of code the computer looks for.
*/
package main

import "github.com/horux/gotask/internal/cli"

func main() {
	/*
	   NOTE ON NETWORKING:
	   At this "Bronze Tier" stage (Project B-01), there is NO network or API.
	   The program communicates directly with the computer's File System.
	   In Project B-02, we will replace this CLI logic with a Web Server
	   using the "net/http" package to handle requests over the internet.
	*/
	cli.Execute()
}
