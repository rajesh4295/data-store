# Data-store
## First project in "Golang"


Data-store is a simple key-value storage cli shell written in Go.

## Features
- Set data
- Get data
- Search data
- Delete data
- Check curent load in store
- Export data

## Installation

Data-store is written with [Go](https://golang.org/dl/) v1.16.

Running the cli

```sh
go run main.go
$ execute commands in the shell
```
## CLI commands and usage


| Cmd | Description | Syntax | Usage | Output
| ------ |------ | ------ | ------ | ------ | 
| set | Add key value | set [key] [value] | set username rajesh4295 |
| get | Get value by key | get [key] | get username | rajesh4295
| exist | Check if key exists | exist [key] | exist username | true/false
| delete | Delete value by key | delete [key] | delete username | true/false
| size | Returns current load in the store. | size | size |  `0% 10% 80%` ...
| export | Exports complete data to `data.json` in the current folder | export | export | 
| exit/quit | Exits from the cli shell | exit/quit | exit/quit | 

## Implementation
- Store is implemented with hash table where each index (calculated using hash method) holds a bucket (Singly linkedlist) and it's bucket node's (linkedlist node)

## TODO
- Grow/ shrink store based on load factor
- Preload store from json
- Support value with spaces

## License

MIT
