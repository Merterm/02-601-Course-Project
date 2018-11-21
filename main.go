package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//ARTH
const SUM string = "SUM"
const SUB string = "SUB"
const MUL string = "MUL"
const DIV string = "DIV"
const MOD string = "MOD"

type Protein struct {
	label  string
	id     int
	chTree int
	phos   bool
	interX []*Protein
}

func main() {
	var filename = "field.txt"
	code := ReadCode(filename)
	fmt.Println(ConvertToProtein(code))

}

//Open file, read it into str, and then call convert protein
func ReadCode(filename string) string {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: something went wrong opening the file.")
		fmt.Println("Probably you gave the wrong filename.")
	}
	defer file.Close()

	//reading the file into str
	// TODO: Remember to read multiline code.!!!!!!!!!
	//var lines []string = make([]string, 0)
	code := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		code = code + scanner.Text() //append(lines, scanner.Text())
		//fmt.Println("This line is: ", scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Sorry: there was some kind of error during the file reading")
		os.Exit(1)
	}
	return code
}

//call detect keyword, create object and return protein object
func ConvertToProtein(code string) Protein {
	name, number := DetectOperator(code)

	protein := CreateProtein(name, number)
	return protein
}

func DetectOperator(code string) (string, int) {
	var splitCode []string = strings.Split(code, " ")
	/*
		  if splitCode[1] == "=" {

					    switch splitCode[3] {
							case "+":
								return SUM
							case "-":
								return SUB
							case "*":
								return MUL
							case "/":
								return DIV
							case "%":
								return MOD
							default:
								return "sorry"
							}

			}
			panic("Panic?")
	*/
	num, _ := strconv.Atoi(splitCode[2])
	return splitCode[0], num
}

func CreateProtein(name string, number int) Protein {
	var p Protein
	p.label = name
	p.chTree = number
	return p
}
