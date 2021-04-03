package main

import (
	"fmt"

	"github.com/rajesh4295/data-store/lib"
)

// "bufio"

// "log"
// "os"
// "os/exec"
// "strings"

func main() {
	store := lib.InitStore()

	store.Insert(lib.StoreData{Key: "name", Value: "rajesh"})
	store.Insert(lib.StoreData{Key: "dob", Value: "feb"})
	store.Insert(lib.StoreData{Key: "company", Value: "tibco"})

	fmt.Println(store.GetLoad())

	fmt.Println(store.Search("name"))

	// reader := bufio.NewReader(os.Stdin)

	/* for {
		fmt.Print("$ ")
		cmd, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		cmdArr := strings.Fields(cmd)
		baseCmd := cmdArr[0]

		switch baseCmd {
		case "set":
			isValid := precheckSet(cmdArr)
			if isValid {

			}
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
	} */
}

// func precheckSet(cmdArr []string) bool {
// 	if len(cmdArr) == 1 {
// 		fmt.Println(cmdArr[0], "cmd expects a key and a value")
// 		return false
// 	}
// 	return true
// }
// func handleUnsupportedInp(cmd string) {
// 	fmt.Println(cmd, "not recognized")
// }
