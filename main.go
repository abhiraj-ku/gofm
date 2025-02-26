package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, helpMess)
	}
	log.SetFlags(0)

	// Parse the args to be read by flag
	flag.Parse()

	err := execCommand()
	if err != nil {
		log.Fatal(err)
	}
}

// execCommand - one function to set all the commands

func execCommand() error {
	command := getArg(0, helpMess)

	// switch-case command based parsing based on provided commands
	switch command {
	case "info":
		path := getArg(1, "Enter path")
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		fmt.Printf("Name: %s\nSize: %d\nIsDir: %t\nModification: %s\n",
			fileInfo.Name(),
			fileInfo.Size(),
			fileInfo.Mode(),
			fileInfo.IsDir(),
			fileInfo.ModTime(),
		)
	case "ls":
		files, err := os.ReadDir(".")
		if err != nil {
			return err
		}
		for _, file := range files {
			fmt.Printf("%s ", file.Name())
		}
		fmt.Println()

	case "rm":
		path := getArg(1, "Enter path")
		return os.RemoveAll(path)

	case "pwd":
		currPath, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Println(currPath)
	case "mv":
		path1 := getArg(1, "Enter path 1")
		path2 := getArg(2, "Enter path 2")
		return os.Rename(path1, path2)

	case "mkdir":
		path := getArg(1, "Enter dir name:")
		return os.Mkdir(path, 0777)

	case "cat":
		path := getArg(1, "Enter path")
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		os.Stdout.Write(data)

	case "touch":
		path := getArg(1, "Enter path")
		_, err := os.Create(path)
		return err

	case "cp":
		path1 := getArg(1, "Enter path 1")
		path2 := getArg(2, "Enter path 2")

		data, err := os.ReadFile(path1)
		if err != nil {
			return err
		}

		return os.WriteFile(path2, data, 0777)

	case "echo":
		content := getArg(1, "Enter content")
		path := getArg(2, "enter path")
		return os.WriteFile(path, []byte(content), 0777)

	default:
		fmt.Print(helpMess)
	}
	return nil

}

// getArg functio to get the command line args
func getArg(indexArg int, message string) string {
	val := flag.Arg(indexArg)
	if val == "" {
		log.Fatal(message)
	}
	return val

}
