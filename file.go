package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Create a new file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write data to the file
	data := []byte("Hello, world!")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data written to file successfully.")

	// Read data from the file
	readData, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Data read from file:", string(readData))
}
