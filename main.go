package main

import(
	"fmt"
	"os"

	"github.com/abema/go-mp4"
)
//"io/ioutil"

func main() {
	// Handle flags
	args := os.Args[1:]
	var cmd string

	if len(args) > 0 {
		cmd, args = args[0], args[1:]
	}

	switch cmd {
	  case "-f", "--file":
			loadVideo(args[0])
	  case "", "-h", "--help":
		  usage()
		default:
			fmt.Printf("Unknown command %s\n", cmd)
			fmt.Errorf("Unknown command %s\n", cmd)
	}
	// TODO: Split video file into screenshots
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

// TODO: Load video file
func loadVideo(filename string) {
	fmt.Printf("Loading video %s ...\n", filename)
  file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error %v", err)
		panic(err)
	}
	defer file.Close()

	_, h_err := mp4.ReadBoxStructure(file, func(h *mp4.ReadHandle) (interface{}, error) {
		fmt.Println("Size", h.BoxInfo.Size)
		return nil, nil
	})
	if h_err != nil {
    fmt.Printf("Error %v", err)
		panic(err)
	}

	file.Close()
}
