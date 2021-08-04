package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fileInfo, err := os.Stat("mee.txt") // produces an os.fileStat value
	// fmt.Println(reflect.TypeOf(fileInfo))
	// fmt.Println(reflect.TypeOf(err))

	if err != nil {
		log.Fatal(err) // prints the error that exist if any
	}

	fmt.Println(fileInfo.Size()) // produces an integer value of the filesize of mee.txt in bytes

}
