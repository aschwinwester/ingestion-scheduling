package main

import (
	"os"
	"fmt"
	"github.com/aschwinwester/ingestion-scheduling/rest"
)	

func isStartServer(args []string) bool {
	start := false
	server := false
	for i := 0; i < len(args); i++ {
		if args[i] == "start" {
			start = true
		}
		if args[i] == "server" {
			server = true
		}
	}
	return start && server
}
func main() {
	args := os.Args
	if isStartServer(args) {
		rest.StartServer()
	} else {
		fmt.Println("Nothing to do")
	}
	fmt.Println("Done!")
}
