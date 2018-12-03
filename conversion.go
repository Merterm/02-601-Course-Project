package main

import (
	"fmt"
	"strconv"
)

//TreeTraversal visit every node in the tree and return the *Vesicle
//Create a vesicle for each iteration
//Base case: the last second level. Process the children into vesicle, and add
//           this vesicle to the vesicle.
//Recurssive case: call the recurssion for the child, add the vesicle to the vesicle
func TreeTraversal(parseTree ParseTree) *Vesicle {
	vesicle := new(Vesicle) //the new function return a pointer to the object
	(*vesicle).vesicles = make([]*Vesicle, 0)

	if parseTree != nil {
		if parseTree.children != nil {
			for num, child := range parseTree.children {
				//if it's == nil means now we are on the second last level
				if child.children == nil {
					//append the vesicle we created to the vesicle in this level
					//fmt.Println("reach the leaves the current node is: ", (*parseTree).name, child.name)
					vesicle.vesicles = append(vesicle.vesicles, GenerateProtein(parseTree))
				} else {
					//if it it not the second last level, then call TreeTraversal
					//fmt.Println("not the leaves the current node is: ", (*parseTree).name, child.name)
					childvesicle := TreeTraversal(child)

					//Assign the order of vesicle operation on receptor's LocSignalRec
					//in every level
					//fmt.Println("current node is ", parseTree.name, ", child is ", child.name)
					//fmt.Println("the receptor name is ", (*childvesicle).receptorList[0])
					/******************
					Solve this problem, the locsignalRec doesn't work
					******************/
					for i := 0; i < len((*childvesicle).receptorList); i++ {
						(*childvesicle).receptorList[i].locSignalRec = num
						fmt.Println((*childvesicle).receptorList[i].locSignalRec)
					}

					vesicle.vesicles = append(vesicle.vesicles, childvesicle)
					if parseTree.name == "IF" {
						ifKinase := new(IfKinase)
						//the receptorName will be the "code"
						ifKinase.receptorName = parseTree.children[1].name
						//the recognizeVesicleName will be the first one(condition)
						ifKinase.recognizeVesicleName = parseTree.children[0].name
						vesicle.ifKinase = ifKinase
					} else if parseTree.name == "WHILE" {
						ifKinase := new(IfKinase)
						//the receptorName will be the "code"
						ifKinase.receptorName = parseTree.children[1].name
						//the recognizeVesicleName will be the first one(condition)
						ifKinase.recognizeVesicleName = parseTree.children[0].name
						vesicle.ifKinase = ifKinase
						//// TODO: how to add label for the while loop???????
					}

				}
			}

		}
	}
	return vesicle
}

//GenerateProtein create the basic vesicle in the last two level
func GenerateProtein(parseTree ParseTree) *Vesicle {
	receptor := new(Receptor)
	vesicle := new(Vesicle)
	vesicle.InitializeVesicle()

	for i := 0; i < len((*parseTree).children); i++ {

		if parseTree.name == "COND" {
			checkerKinase := new(CheckerKinase)
			receptor2 := new(Receptor)

			(*receptor).name = (*parseTree.children[0]).name
			(*checkerKinase).name = (*parseTree.children[1]).name
			(*receptor2).name = (*parseTree.children[2]).name

			//vesicle.proteinList = append(vesicle.proteinList, checkerKinase)
			vesicle.receptorList = make([]*Receptor, 0)
			vesicle.receptorList = append(vesicle.receptorList, receptor)
			vesicle.receptorList = append(vesicle.receptorList, receptor2)

		} else if parseTree.name == "ARTH" {
			glucotrans := new(Glucotrans)

			(*receptor).name = (*parseTree.children[0]).name
			number, _ := strconv.Atoi((*parseTree.children[2]).name)
			if (*parseTree.children[1]).name == "+" {
				(*glucotrans).glucoCount = number
			} else if (*parseTree.children[1]).name == "-" {
				(*glucotrans).glucoCount = -number
			}
			vesicle.receptorList = make([]*Receptor, 0)
			vesicle.glucoTrans = glucotrans
			vesicle.receptorList = append(vesicle.receptorList, receptor)

		} else if parseTree.name == "ASSIGN" {
			glucotrans := new(Glucotrans)
			(*receptor).name = (*parseTree.children[1]).name
			number, _ := strconv.Atoi((*parseTree.children[0]).name)
			(*glucotrans).glucoCount = number
			vesicle.glucoTrans = glucotrans
			vesicle.receptorList = append(vesicle.receptorList, receptor)
		}
	}
	return vesicle
}
