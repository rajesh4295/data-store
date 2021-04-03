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
		baseCmd := cmdArr[0]

		switch baseCmd {
		case "set":
			isValid := precheck(baseCmd, cmdArr)
			if isValid {
				handleSet(cmdArr, store)
			}
		case "get":
			isValid := precheck(baseCmd, cmdArr)
			if isValid {
				handleGet(cmdArr, store)
			}
		case "size":
			handleSize(store)
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
	case "get":
		if len(cmdArr) < 2 {
			fmt.Println(cmdArr[0], "expects a key. Example:$ get name")
			return false
		}
	}

	return true
}

func handleSet(cmdArr []string, store *lib.Store) {
	data := lib.StoreData{Key: cmdArr[1], Value: cmdArr[2]}
	store.Insert(data)
}

func handleGet(cmdArr []string, store *lib.Store) {
	fmt.Println(store.Get(cmdArr[1]))
}

func handleSize(store *lib.Store) {
	fmt.Println(store.GetLoad())
}

func handleUnsupportedInp(cmd string) {
	fmt.Println(cmd, "not recognized")
}
