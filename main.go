package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/rajesh4295/data-store/lib"
)

const (
	helpBanner = `
		List of available commands
		help
			- Prints help
			- Usage: help
		set
			- Add key value to store.
			- Usage: set <key> <value>
			- Example: set username rajesh4295
		get
			- Get value by key from store.
			- Usage: get <key>
			- Example: get username	
		exist
			- Check if key exists in store.
			- Usage: exist <key>
			- Example: exist username
		delete
			- Delete value by key from store.
			- Usage: delete <key>
			- Example: delete username
		size
			- Returns current load in the store.
			- Usage: size
		export
			- Exports complete data in store to data.json file in the current folder.
			- Usage: export
		clear
			- Clears the terminal
			- Usage: clear
		quit/exit
			- Exits from the data-store cli shell.
			- Usage: quit or exit
	`
)

func main() {
	store := lib.InitStore()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		cmd, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		cmdArr := strings.Fields(cmd)
		baseCmd := ""
		if len(cmdArr) < 1 {
			continue
		}
		baseCmd = cmdArr[0]

		switch baseCmd {
		case "help":
			fmt.Println(helpBanner)
		case "set":
			isValid := precheck(baseCmd, cmdArr)
			if isValid {
				set(cmdArr, store)
			}
		case "get":
			isValid := precheck(baseCmd, cmdArr)
			if isValid {
				get(cmdArr, store)
			}
		case "exist":
			isValid := precheck(baseCmd, cmdArr)
			if isValid {
				exist(cmdArr, store)
			}
		case "delete":
			isValid := precheck(baseCmd, cmdArr)
			if isValid {
				delete(cmdArr, store)
			}
		case "export":
			export(store)
		case "size":
			size(store)
		case "exit", "quit":
			os.Exit(0)
		case "clear":
			c := exec.Command("clear") // works in bash shell not in windows
			c.Stderr = os.Stderr
			c.Stdout = os.Stdout
			err := c.Run()
			if err != nil {
				handleUnsupportedInp(baseCmd)
			}
		default:
			handleUnsupportedInp(baseCmd)
		}
	}
}

func precheck(action string, cmdArr []string) bool {

	switch action {
	case "set":
		if len(cmdArr) < 3 {
			fmt.Println(cmdArr[0], "expects a key and a value. Example:$ set name jhon")
			return false
		}
	case "get", "exist", "delete":
		if len(cmdArr) < 2 {
			fmt.Println(cmdArr[0], "expects a key. Example:$", cmdArr[0], "name")
			return false
		}
	}

	return true
}

func set(cmdArr []string, store *lib.Store) {
	data := lib.StoreData{Key: cmdArr[1], Value: cmdArr[2]}
	store.Insert(data)
}

func get(cmdArr []string, store *lib.Store) {
	res, err := store.Get(cmdArr[1])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func exist(cmdArr []string, store *lib.Store) {
	res, _ := store.Search(cmdArr[1])
	fmt.Println(res)
}

func delete(cmdArr []string, store *lib.Store) {
	res, err := store.Delete(cmdArr[1])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func size(store *lib.Store) {
	fmt.Println(store.GetLoad(), "%")
}

func export(store *lib.Store) {
	err := store.Export()
	if err != nil {
		fmt.Println(err)
	}
}

func handleUnsupportedInp(cmd string) {
	fmt.Println(cmd, "not recognized")
}
