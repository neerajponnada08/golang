package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files section !!")
	content := "This needs to go in a file"

	file, err := os.Create("./mylegofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is: ",length)
	defer file.Close()

	readFile("./mylegofile.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)

	checkNilErr(err)

	fmt.Println("Text data inside the file is: \n",string(dataByte))

}

func checkNilErr (err error) {
	if err!= nil {
		panic(err)
	}
}