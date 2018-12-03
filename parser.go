package main

/* ----------------------------------------------------------------------------
ProLANG Project
Author: Mert Inan
Date: 28 Nov 2018
Description:	This is the parser code for the ProLANG project.
              This program reads a specific code in ProLANG and generates the
              parse tree for it. Highest level function is the Parse(filename)
              function. It takes the filename of the ProLANG file.
-----------------------------------------------------------------------------*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/******************************************************************************
															OBJECT DECLARATIONS
******************************************************************************/

//CODE is a top level constant
const CODE = "CODE"

//ParseTree is the header pointer to the parse tree
type ParseTree *Node

//Node is an object on the parse tree. It has children and a name.
//Its children are in parsing order.
type Node struct {
	name     string
	children []*Node
}

/******************************************************************************
																PARSER FUNCTIONS
******************************************************************************/

//Parse is the highest level function. It reads a ProLANG file and creates a
// parse tree out of it and returns it.
func Parse(filename string) ParseTree {
	//1. Read the file into a string

	// open file
	parseFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: file cannot opened!")
		os.Exit(1)
	}

	// hold the whole code in a single string
	var code string

	//Scan the whole file into the string
	scanner := bufio.NewScanner(parseFile)
	for scanner.Scan() {
		code += strings.TrimSpace(scanner.Text()) + " "
	}
	// check the scanner
	if scanner.Err() != nil {
		fmt.Println("Error: Something wrong with the scanner!")
		os.Exit(1)
	}
	// close file
	parseFile.Close()

	//2.Split the string by spaces
	var codeArr []string = strings.Split(code, " ")

	//Remove the last space
	codeArr = codeArr[:len(codeArr)-1]

	//Initialize the parse tree
	var parseTree ParseTree = new(Node)

	//3.Pass it to ParseCode function.
	var err2 bool
	codeArr, err2 = ParseCode(codeArr, parseTree)

	if err2 {
		panic("There was an error while parsing!")
	}

	return parseTree
}

//ParseCode takes a split code array and calls the relevant parsing function
//depending on the first keyword in the array. It adds a code node to the tree
// and passes down the tree. Returns true if there is an error while parsing.
func ParseCode(codeArr []string, ptr ParseTree) ([]string, bool) {

	//Add the code node to the pointer's children
	if ptr == nil {
		//Create a Code node
		codeNode := new(Node)
		codeNode.name = CODE
		ptr = codeNode
	} else if ptr != nil && ptr.children == nil {
		ptr.name = CODE
	} else if ptr.name != CODE {
		//Create a Code node
		codeNode := new(Node)
		codeNode.name = CODE
		ptr.children = append(ptr.children, codeNode)
		ptr = codeNode
	}

	//Check whether the given array is empty or not
	if len(codeArr) == 0 {
		return codeArr, true //there is an error
	}

	//Check the first element and call the relevant function
	err := false
	if codeArr[0] == "while" {
		codeArr, err = ParseWhile(codeArr, ptr)
	} else if codeArr[0] == "if" {
		codeArr, err = ParseIf(codeArr, ptr)
	} else if codeArr[1] == "=" {
		codeArr, err = ParseArithmetic(codeArr, ptr)
	} else if codeArr[1] == ":=" {
		codeArr, err = ParseAssign(codeArr, ptr)
	} else {
		panic("Not parseable!")
	}

	//Call the code function again on the remaining code
	if len(codeArr) != 0 && !err {
		codeArr, _ = ParseCode(codeArr, ptr)
	}

	//If there is error, return true
	if err {
		return codeArr, true
	}
	return codeArr, false
}

//ParseWhile takes codeArr as input and parses a while loop and puts into parse
//tree. Returns true if there is an error while parsing.
func ParseWhile(codeArr []string, ptr *Node) ([]string, bool) {
	//Create a while node & add to tree
	whileNode := new(Node)
	whileNode.name = "WHILE"
	ptr.children = append(ptr.children, whileNode)

	//Remove while from array
	codeArr, _ = Pop(codeArr)

	//Read until you see a { and put it into tmp
	i := 0
	cond := make([]string, 0)
	for codeArr[i] != "{" {
		var val string
		codeArr, val = Pop(codeArr)
		cond = append(cond, val)
	}

	//Pop the bracket
	codeArr, _ = Pop(codeArr)

	//Parse the conditional statement
	var err bool
	_, err = ParseConditional(cond, whileNode)

	//Read until you see a } and put it into tmp
	j := 0
	stmt := make([]string, 0)
	bracketCnt := 1
	for (codeArr[j] != "}" || bracketCnt != 1) && j < len(codeArr) {
		if codeArr[j] == "{" {
			bracketCnt++
		}
		if codeArr[j] == "}" {
			bracketCnt--
		}

		var val string
		codeArr, val = Pop(codeArr)
		stmt = append(stmt, val)
	}

	fmt.Println(stmt)
	//Pop the bracket
	codeArr, _ = Pop(codeArr)

	//Parse the statement
	var err2 bool
	_, err2 = ParseCode(stmt, whileNode)

	if err || err2 {
		return codeArr, true
	}
	return codeArr, false
}

//ParseConditional takes codeArr as input and parses a conditional statement and
//puts it into tree. Returns true if there is an error while parsing.
func ParseConditional(codeArr []string, ptr *Node) ([]string, bool) {
	//Parse if the conditional is of the form <VAR> <OP> <VAR>
	if len(codeArr) == 3 {
		//Create a conditional node & add to tree
		condNode := new(Node)
		condNode.name = "COND"
		ptr.children = append(ptr.children, condNode)

		//TODO: YOU MAY WANT TO ADD PARANTHESIZED EXPRESSIONS, AND OR EXPRESSIONS
		//Put the variables to the tree
		//Create a variable node & add to tree
		var1Node := new(Node)
		var val string
		codeArr, val = Pop(codeArr)
		var1Node.name = val
		condNode.children = append(condNode.children, var1Node)

		//Create an operation node & add to tree
		opNode := new(Node)
		var val2 string
		codeArr, val2 = Pop(codeArr)
		opNode.name = val2
		condNode.children = append(condNode.children, opNode)

		//Create a variable node & add to tree
		var2Node := new(Node)
		var val3 string
		codeArr, val3 = Pop(codeArr)
		var2Node.name = val3
		condNode.children = append(condNode.children, var2Node)

		return codeArr, false //no error
	}
	return codeArr, true //there is an error
}

//ParseIf takes codeArr as input and parses an if statement and puts it in to
//parse tree. Returns true if there is an error while parsing.
func ParseIf(codeArr []string, ptr *Node) ([]string, bool) {
	//Create a while node & add to tree
	ifNode := new(Node)
	ifNode.name = "IF"
	ptr.children = append(ptr.children, ifNode)

	//Remove if from array
	codeArr, _ = Pop(codeArr)

	//Read until you see a { and put it into cond
	i := 0
	cond := make([]string, 0)
	for codeArr[i] != "{" {
		var val string
		codeArr, val = Pop(codeArr)
		cond = append(cond, val)
	}

	//Pop the bracket
	codeArr, _ = Pop(codeArr)

	//Parse the conditional statement
	var err bool
	_, err = ParseConditional(cond, ifNode)

	//Read until you see a } and put it into stmt
	j := 0
	stmt := make([]string, 0)
	bracketCnt := 1
	for (codeArr[j] != "}" || bracketCnt != 1) && j < len(codeArr) {
		if codeArr[j] == "{" {
			bracketCnt++
		}
		if codeArr[j] == "}" {
			bracketCnt--
		}
		var val string
		codeArr, val = Pop(codeArr)
		stmt = append(stmt, val)
	}

	//Pop the bracket
	codeArr, _ = Pop(codeArr)

	//Parse the statement
	var err2 bool
	_, err2 = ParseCode(stmt, ifNode)

	if err || err2 {
		return codeArr, true
	}
	return codeArr, false
}

//ParseArithmetic takes codeArr as input and parses an arithmetic statement and
//puts it into parse tree. Returns true if there is an error while parsing.
func ParseArithmetic(codeArr []string, ptr *Node) ([]string, bool) {
	//Parse if the conditional is of the form <VAR> <=> <VAR> <OP> <NUM> or
	//<VAR> <=> <NUM> <OP> <VAR>
	fmt.Println(codeArr)
	if codeArr[1] == "=" {
		if codeArr[0] == codeArr[2] || codeArr[0] == codeArr[4] {

			//Create an arithmetic node & add to tree
			arthNode := new(Node)
			arthNode.name = "ARTH"
			ptr.children = append(ptr.children, arthNode)

			//Put the variable to the tree
			//Create a variable node & add to tree
			varNode := new(Node)
			var val string
			codeArr, val = Pop(codeArr)
			varNode.name = val
			arthNode.children = append(arthNode.children, varNode)

			//Pop the elements in-between
			codeArr, _ = Pop(codeArr)
			codeArr, _ = Pop(codeArr)

			//Create an operation node & add to tree
			opNode := new(Node)
			var val2 string
			codeArr, val2 = Pop(codeArr)
			opNode.name = val2
			arthNode.children = append(arthNode.children, opNode)

			//Create a number node & add to tree
			//Check that it is an actual number
			var val3 string
			codeArr, val3 = Pop(codeArr)
			numNode := new(Node)
			numNode.name = val3
			arthNode.children = append(arthNode.children, numNode)

			return codeArr, false //no error
		}
	}
	return codeArr, true //there is an error
}

//ParseAssign takes codeArr as input and parses an assignment statement. Returns
//true if there is an error while parsing.
func ParseAssign(codeArr []string, ptr *Node) ([]string, bool) {
	//Parse if the assignment is of the form <VAR> <:=> <NUM> or <VAR> <:=> <BOOL>
	//Check that it is an actual number
	if codeArr[2] != "true" && codeArr[2] != "false" {
		_, err := strconv.Atoi(codeArr[2])
		if err != nil {
			return codeArr, true //there is an error
		}
	}

	//Create an assignment node & add to tree
	assignNode := new(Node)
	assignNode.name = "ASSIGN"
	ptr.children = append(ptr.children, assignNode)

	//Put the variable to the tree
	//Create a variable node & add to tree
	varNode := new(Node)
	var val string
	codeArr, val = Pop(codeArr)
	varNode.name = val
	assignNode.children = append(assignNode.children, varNode)

	//Pop the assignment sign
	codeArr, _ = Pop(codeArr)

	//Create a number or boolean node & add to tree
	numBoolNode := new(Node)
	var val2 string
	codeArr, val2 = Pop(codeArr)
	numBoolNode.name = val2
	assignNode.children = append(assignNode.children, numBoolNode)

	return codeArr, false //no error
}

/******************************************************************************
																HELPER FUNCTIONS
******************************************************************************/

//Pop deletes the initial element of the arr and returns it
func Pop(arr []string) ([]string, string) {
	val := arr[0]
	arr = arr[1:]
	return arr, val
}

//PrintParseTree prints the parseTree in depth first search order.
func PrintParseTree(parseTree ParseTree) {
	if parseTree != nil {
		//Print the main node and call the print function for all the children
		fmt.Println("Node: ", parseTree.name)
		fmt.Print("Children of ", parseTree.name, ": [")
		for _, child := range parseTree.children {
			fmt.Print(child.name, " ")
		}
		fmt.Println("]")
		for _, child := range parseTree.children {
			PrintParseTree(child)
		}
	}
}

//End of program
