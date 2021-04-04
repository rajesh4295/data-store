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
		if len(cmdArr) > 0 {
			baseCmd = cmdArr[0]
		} else {
			continue
		}

		switch baseCmd {
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
	fmt.Println(store.Get(cmdArr[1]))
}

func exist(cmdArr []string, store *lib.Store) {
	fmt.Println(store.Search(cmdArr[1]))
}

func delete(cmdArr []string, store *lib.Store) {
	fmt.Println(store.Delete(cmdArr[1]))
}

func size(store *lib.Store) {
	fmt.Println(store.GetLoad(), "%")
}

func export(store *lib.Store) {
	store.Export()
}
func handleUnsupportedInp(cmd string) {
	fmt.Println(cmd, "not recognized")
}
