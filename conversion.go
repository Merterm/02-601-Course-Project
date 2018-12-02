package main

import "strconv"

//TreeTraversal visit every node in the tree and return the *Vesicle
//Create a vesicle for each iteration
//Base case: the last second level. Process the children into vesicle, and add
//           this vesicle to the vesicle.
//Recurssive case: call the recurssion for the child, add the vesicle to the vesicle
func TreeTraversal(parseTree ParseTree) *Vesicle {
	var vesicle *Vesicle
	vesicle.vesicles = make([]*Vesicle, 0)

	if parseTree != nil {
		if parseTree.children != nil {
			for num, child := range parseTree.children {
				//if it's == nil means now we are on the second last level
				if child.children == nil {
					//append the vesicle we created to the vesicle in this level
					vesicle.vesicles = append(vesicle.vesicles, GenerateProtein(parseTree))
				} else {
					//if it it not the second last level, then call TreeTraversal

					childvesicle := TreeTraversal(child)

					//Assign the order of vesicle operation on receptor's LocSignalRec
					//in every level
					(*childvesicle).receptorList[num].locSignalRec = num
					vesicle.vesicles = append(vesicle.vesicles, childvesicle)

					if parseTree.name == "IF" {
						var ifKinase *IfKinase
						//the receptorName will be the "code"
						ifKinase.receptorName = parseTree.children[1].name
						//the recognizeVesicleName will be the first one(condition)
						ifKinase.recognizeVesicleName = parseTree.children[0].name
						//vesicle.proteinList = append(vesicle.proteinList, ifKinase)
					} else if parseTree.name == "WHILE" {
						var ifKinase *IfKinase
						//the receptorName will be the "code"
						ifKinase.receptorName = parseTree.children[1].name
						//the recognizeVesicleName will be the first one(condition)
						ifKinase.recognizeVesicleName = parseTree.children[0].name
						//vesicle.proteinList = append(vesicle.proteinList, ifKinase)
						//// TODO: how to add value to subtrate? label it go back and forth
					}
				}
			}

		}
	}
	return vesicle
}

//GenerateProtein create the basic vesicle in the last two level
func GenerateProtein(parseTree ParseTree) *Vesicle {
	var receptor *Receptor
	var vesicle *Vesicle
	vesicle.InitializeVesicle()

	for i := 0; i < len((*parseTree).children); i++ {

		if parseTree.name == "COND" {
			var checkerKinase *CheckerKinase
			var receptor2 *Receptor

			(*receptor).name = (*parseTree.children[0]).name
			(*checkerKinase).name = (*parseTree.children[1]).name
			(*receptor2).name = (*parseTree.children[2]).name

			//vesicle.proteinList = append(vesicle.proteinList, checkerKinase)
			//vesicle.receptorList = append(vesicle.receptorList, receptor)
			//vesicle.receptorList = append(vesicle.recpeotrList, receptor)

		} else if parseTree.name == "ARTH" {
			var glucotrans *Glucotrans

			(*receptor).name = (*parseTree.children[0]).name
			number, _ := strconv.Atoi((*parseTree.children[2]).name)
			if (*parseTree.children[1]).name == "+" {
				(*glucotrans).glucoCount = number
			} else if (*parseTree.children[1]).name == "-" {
				(*glucotrans).glucoCount = -number
			}
			//vesicle.proteinList = append(vesicle.proteinList, glucotrans)
			//vesicle.receptorList = append(vesicle.receptorList, receptor)

		} else if parseTree.name == "ASSIGN" {
			var glucotrans *Glucotrans

			(*receptor).name = (*parseTree.children[0]).name
			number, _ := strconv.Atoi((*parseTree.children[1]).name)
			(*glucotrans).glucoCount = number
			//vesicle.proteinList = append(vesicle.proteinList, glucotrans)
			//vesicle.receptorList = append(vesicle.receptorList, receptor)
		}
	}
	return vesicle
}
