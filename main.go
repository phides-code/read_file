/*Write a program which reads information from a file and represents it in a slice of 
structs. Assume that there is a text file which contains a series of names. Each line of the 
text file has a first name and a last name, in that order, separated by a single space on the 
line. 

Your program will define a name struct which has two fields, fname for the first name, and 
lname for the last name.

Your program should prompt the user for the name of the text file. Your program will 
successively read each line of the text file and create a struct which contains the first and 
last names found in the file. Each struct created will be added to a slice, and after all 
lines have been read from the file, your program will have a slice containing one struct for 
each line in the file. After reading all lines from the file, your program should iterate 
through your slice of structs and print the first and last names found in each struct.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// define a name struct w/ 2 fields, fname and lname
	type Name struct {
		fname string
		lname string
	}

	var mySlice = make([]Name, 0)
	var newName Name
	var textFileName string

	// prompt the user for the name of the text file
	fmt.Printf("Please enter the name of the text file: ")
	fmt.Scan(&textFileName)

	// read each line of text from the file
	file, _ := os.Open(textFileName)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		thisLine := strings.Fields(scanner.Text())

		// create a struct containing the first and last names
		newName.fname = thisLine[0]
		newName.lname = thisLine[1]

		// each struct will be added to a slice
		mySlice = append(mySlice, newName)
	}

	file.Close()

	//iterate through the slice and print the first and last names
	for _, n := range mySlice {
		println(n.fname, n.lname)
	}
}
