package main

import(
	"fmt"
	"os"
)

func main() {
	// Handle flags
	args := os.Args[1:]
	var cmd string

	if len(args) > 0 {
		cmd, args = args[0], args[1:]
	}

	switch cmd {
	  case "-f", "--file":
			framesFromVideo(args[0])
	  case "", "-h", "--help":
		  usage()
		default:
			fmt.Printf("Unknown command %s\n", cmd)
			fmt.Errorf("Unknown command %s\n", cmd)
	}
	// TODO: return/output video file information
}

func usage() {
	fmt.Println(`
Show Screenshot Search

Usage:
  $ go build main.go
  $ main <command> [arguments]

Commands:
  -h, --help: prints this message
  -f, --file: the video file to run on
	`)
}

