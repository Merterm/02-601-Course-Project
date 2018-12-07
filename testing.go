package main

import "fmt"

//This is the main function that tests the parser, converter and the simulator
func main() {
	filename := "hello.txt"
	parseTree := Parse(filename)
	PrintParseTree(parseTree)

	//Testing Object Functions
	fmt.Println("*************** Testing Initialize Vesicle *****************")
	var tmpVesicle *Vesicle
	tmpVesicle = InitializeVesicle(tmpVesicle)
	fmt.Println(tmpVesicle)

	fmt.Println("*************** Testing Initialize Substrate *****************")
	var tmpSubstrate *Substrate
	tmpSubstrate = InitializeSubstrate(tmpSubstrate, "a")
	fmt.Println(tmpSubstrate)

	fmt.Println("*************** Testing TakeInProtein *****************")
	fmt.Println("No Receptor, so should not take in")
	tmpVesicle.TakeInProtein(tmpSubstrate)
	fmt.Println(tmpVesicle)

	fmt.Println("*************** Testing TakeInProtein After Receptor Addition *****************")
	fmt.Println("It should take it now!")
	var tmpReceptor *Receptor
	tmpReceptor = InitializeReceptor(tmpReceptor, "a")
	fmt.Println(tmpReceptor)
	tmpVesicle.AddReceptor(tmpReceptor)
	tmpVesicle.TakeInProtein(tmpSubstrate)
	fmt.Println(tmpVesicle)

	fmt.Println("*************** Testing an Assignment Vesicle *****************")
	var assignVesicle *Vesicle
	assignVesicle = InitializeVesicle(assignVesicle)
	fmt.Println(assignVesicle)
	assignVesicle.glucoTrans = new(Glucotrans)
	assignVesicle.glucoTrans.name = "a"
	assignVesicle.glucoTrans.glucoCount = 1
	fmt.Println(assignVesicle.glucoTrans)
	assignVesicle.AddReceptor(tmpReceptor)
	fmt.Println(assignVesicle)
	assignVesicle.TakeInProtein(tmpSubstrate)
	fmt.Println(tmpSubstrate.name, ": ", tmpSubstrate.glucoCount)
	fmt.Println(assignVesicle)
	assignVesicle.DoAssignmentReaction()
	fmt.Println(assignVesicle.substrateList[0].name, ": ", assignVesicle.substrateList[0].glucoCount)

	fmt.Println("*************** Testing Simulation on Assignment Vesicle *****************")
	var assignVesicle2 *Vesicle
	assignVesicle2 = InitializeVesicle(assignVesicle2)
	fmt.Println("Assignment vesicle: ", assignVesicle2)
	assignVesicle2.glucoTrans = new(Glucotrans)
	assignVesicle2.glucoTrans.name = "a"
	assignVesicle2.glucoTrans.glucoCount = 1
	assignVesicle2.vesicleType = ASSIGN
	fmt.Println("Glucosyltransferase effector: ", assignVesicle2.glucoTrans.name, " count : ", assignVesicle2.glucoTrans.glucoCount)
	assignVesicle2.AddReceptor(tmpReceptor)
	fmt.Println("Assignment vesicle after receptor addition: ", assignVesicle2.name,
		"receptor:", assignVesicle2.receptorList[0].name, "Glucosyltransferase: ",
		assignVesicle2.glucoTrans.name)
	var tmpCell *Vesicle
	tmpCell = InitializeVesicle(tmpCell)
	tmpCell.vesicles = append(tmpCell.vesicles, assignVesicle2)
	SimulateCell(tmpCell, 2)
	fmt.Println(tmpCell)
	fmt.Println("Simulation successful.")

	fmt.Println("*************** Testing Conversion *****************")
	fmt.Println("Creating a parse tree that contains two assignment statements and a if statement.")
	var code, assign, assignName, assignValue, assign2, assign2Name, assign2Value Node
	var code2, ifNode, condition, conditionName, conditionSymbol, conditionValue Node
	var arth, arthName, arthSymbol, arthValue Node
	var parseTree2 ParseTree
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

	parseTree2 = &code
	vesicle := TreeTraversal(parseTree2)
	fmt.Println("assignment parse tree", condition.children[0].name, condition.children[1].name, condition.children[2].name)
	fmt.Println("assignment vesicle 1: ", vesicle.vesicles[0].vesicleType)
	fmt.Println("assignment vesicle 2: ", vesicle.vesicles[1].vesicleType)
	fmt.Println("if vesicle: ", vesicle.vesicles[2].vesicles[0].vesicleType)
	fmt.Println("Conversion successful!")
	fmt.Println("Testing complete.")
}
