package main

import "fmt"

/*
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//SUM is ARTH operation
const SUM string = "SUM"

//SUB is ARTH operation
const SUB string = "SUB"

//Protein is a structure for variable in the function
type Protein struct {
	label  string
	id     int
	chTree int
	phos   bool
	interX []*Protein
}

//Cell is an environment of protein protein interaction
type Cell struct {
	proteinList []Protein
}
*/
func main() {

	var code, assign, assignName, assignValue, assign2, assign2Name, assign2Value Node
	var code2, ifNode, condition, conditionName, conditionSymbol, conditionValue Node
	var arth, arthName, arthSymbol, arthValue Node
	var parseTree ParseTree
	code.name = "CODE"
	assign.name = "ASSIGN"
	assignName.name = "a"
	assignValue.name = "2"
	assign2.name = "ASSIGN"
	assign2Name.name = "b"
	assign2Value.name = "3"
	code2.name = "CODE"
	ifNode.name = "IF"
	condition.name = "COND"
	conditionName.name = "a"
	conditionSymbol.name = "=="
	conditionValue.name = "2"
	arth.name = "ARTH"
	arthName.name = "a"
	arthSymbol.name = "+"
	arthValue.name = "1"
	code.children = make([]*Node, 0)
	assign.children = make([]*Node, 0)
	assign2.children = make([]*Node, 0)
	code2.children = make([]*Node, 0)
	ifNode.children = make([]*Node, 0)
	condition.children = make([]*Node, 0)

	arth.children = make([]*Node, 0)

	code.children = append(code.children, &assign, &assign2, &code2)
	assign.children = append(assign.children, &assignName, &assignValue)
	assign2.children = append(assign2.children, &assign2Name, &assign2Value)
	code2.children = append(code2.children, &ifNode)
	ifNode.children = append(ifNode.children, &condition, &arth)
	condition.children = append(condition.children, &conditionName, &conditionSymbol, &conditionValue)
	arth.children = append(arth.children, &arthName, &arthSymbol, &arthValue)

	parseTree = &code
	vesicle := TreeTraversal(parseTree)
	//fmt.Println("main ", condition.children[0].name, condition.children[1].name, condition.children[2].name)
	fmt.Println("main ", vesicle.vesicles[0].vesicles)
	/*
		//Creating empty cell structure
		cell := CreateCell()

		//reading the code and create protein
		var filename = "field.txt"
		code := ReadCode(filename)
		fmt.Println(ConvertToProtein(cell, code))
	*/
}

/*
//CreateCell create the Cell structure, this will be the space where protein interacts.
func CreateCell() Cell {
	var cell Cell
	cell.proteinList = make([]Protein, 0)
	return cell
}

//ReadCode open file, read it into str, and then call convert protein
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

//ConvertToProtein call detect keyword, create object and return protein object
func ConvertToProtein(cell Cell, code string) (Protein, Cell) {
	dataType, name, number, proteinName := DetectOperator(code)
	var protein Protein
	if dataType == "number" {
		protein = CreateProtein(name, number)
	} else if dataType == "protein" {
		protein, cell = BindToProtein(name, proteinName, cell)
	} else {
		panic("The dataType in ConvertToProtein is wrong!")
	}
	return protein, cell
}

func DetectOperator(code string) (string, string, int, string) {
	var splitCode []string = strings.Split(code, " ")

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

	//convert the splitCode[2] to integer, if it is number, then return the numbe
	//if it is not number, then return the string of the protein name
	dataType := ""
	vrbl := ""
	//// TODO: Not sure whether the value of num when splitCode[2] is not an integer
	num, err := strconv.Atoi(splitCode[2])
	if err != nil {
		vrbl = splitCode[2]
		dataType = "protein"
	} else {
		dataType = "number"
	}
	return dataType, splitCode[0], num, vrbl
}

//CreateProtein use the name and number to create the protein with carbohydrate labling
func CreateProtein(name string, number int) Protein {
	var p Protein
	p.label = name
	p.chTree = number

	return p
}

//BindToProtein add the proteinName into the interX list of protein name, and update the cell
func BindToProtein(name, proteinName string, cell Cell) (Protein, Cell) {
	num := 0
	var tmpProtein, addProtein Protein
	for i := 0; i < len(cell.proteinList); i++ {
		if cell.proteinList[i].label == "name" {
			tmpProtein = cell.proteinList[i]
			num = i
		} else if cell.proteinList[i].label == "proteinName" {
			addProtein = cell.proteinList[i]
		}
	}
	tmpProtein.interX = append(tmpProtein.interX, &addProtein)
	cell.proteinList[num] = tmpProtein
	return tmpProtein, cell
}
*/
