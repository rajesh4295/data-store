package util

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rajesh4295/data-store/lib"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Could not load env file", err)
	}
	return os.Getenv(key)
}

func ReadFile(path string) []byte {
	body, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal("Unable to read file", err)
	}
	return body
}

func WriteFile(data []byte, path string) {
	perm, err := strconv.Atoi(GetEnv("FILE_PERM"))

	if err != nil {
		log.Fatal("Could not convert string to int")
	}

	err = ioutil.WriteFile(path, data, fs.FileMode(perm))
	if err != nil {
		log.Fatal("Error writing file ", err)
	}
}

func PrintResult(data lib.StoreData) {
	fmt.Printf("%v\n", data)
}
