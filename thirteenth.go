// file using os package
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() // close the file at last
	fmt.Println("File name is ", file.Name())
	fileInfo, err := file.Stat()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Is dir? ", fileInfo.IsDir(), "Modified time: ", fileInfo.ModTime(), "File size: ", fileInfo.Size())

	// write some text
	file.WriteString("Test")

	data := make([]byte, 16)

	file.Seek(0, 0) // seek to 0 from 0

	n, err := file.Read(data)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("No. of bytes read: ", n)

	fmt.Println("File content:", string(data))
}
