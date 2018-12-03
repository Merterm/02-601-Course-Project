package main

/* ----------------------------------------------------------------------------
ProLANG Project
Author: Ian Lee & Mert Inan
Date: 30 Nov 2018
Description:
-----------------------------------------------------------------------------*/

/******************************************************************************
																VESICLE OBJECT
******************************************************************************/

//Vesicle is the object of our function
type Vesicle struct {
	name          string
	vesicleType   string
	substrateList []*Substrate
	receptorList  []*Receptor
	ifKinase      *IfKinase
	checkerKinase *CheckerKinase
	glucoTrans    *Glucotrans
	vesicles      []*Vesicle
}

func InitializeVesicle(vesicle *Vesicle) *Vesicle {
	vesicle = new(Vesicle)
	vesicle.substrateList = make([]*Substrate, 0)
	vesicle.receptorList = make([]*Receptor, 0)
	vesicle.vesicles = make([]*Vesicle, 0)
	return vesicle
}

func (vesicle *Vesicle) CopyVesicle(copiedVesicle *Vesicle) {
	//Copy the trivial parameters
	vesicle.name = copiedVesicle.name
	vesicle.vesicleType = copiedVesicle.vesicleType

	//Deep copy the lists !!!!!TODO!!!!!!!

	//Call the same function for children vesicles !!!!!TODO!!!!!!!
}

//TakeInProtein add protein to proteinList if it could be recognized by the
//receptor
func (vesicle *Vesicle) TakeInProtein(substrate *Substrate) {
	exist := false
	if vesicle.receptorList != nil {
		for _, otherProtein := range vesicle.receptorList {
			//Check that the receptor is active
			if otherProtein.CheckPhosphoStatus() {
				exist = (*otherProtein).CheckProtein(substrate.name, substrate.locSignal)
			}
		}
		if exist {
			if vesicle.substrateList != nil {
				vesicle.substrateList = append(vesicle.substrateList, substrate)
			} else {
				vesicle.substrateList = make([]*Substrate, 0)
				vesicle.substrateList = append(vesicle.substrateList, substrate)
			}
		}
	} else {
		panic("receptor list is empty!!")
	}
}

//AddReceptor adds the given receptor pointer to the vesicle's receptors
func (vesicle *Vesicle) AddReceptor(receptor *Receptor) {
	vesicle.receptorList = append(vesicle.receptorList, receptor)
}

//!!!!!!!!!!! TODO !!!! Need to return a pointer to protein and increment its locSignal
//PumpOutProtein remove the protein from proteinList
func (vesicle *Vesicle) PumpOutProtein(substrate *Substrate) {
	if vesicle.substrateList != nil {
		for number, theProtein := range vesicle.substrateList {
			if theProtein.name == substrate.name {
				vesicle.RemoveFromProteinList(number)
			}
		}
	} else {
		panic("protein list is empty")
	}
}

//RemoveFromProteinList delete the element from proteinList
func (vesicle *Vesicle) RemoveFromProteinList(number int) {
	vesicle.substrateList = append(vesicle.substrateList[:number],
		vesicle.substrateList[number+1:]...)
}

//DoReactionInside
func (vesicle *Vesicle) DoReactionInside() {
	if vesicle.vesicleType == IF {
		vesicle.DoIfReaction()
	} else if vesicle.vesicleType == WHILE {
		vesicle.DoWhileReaction()
	} else if vesicle.vesicleType == ASSIGN {
		vesicle.DoAssignmentReaction()
	} else if vesicle.vesicleType == COND {
		vesicle.DoConditionalReaction()
	} else if vesicle.vesicleType == ARTH {
		vesicle.DoArithmeticReaction()
	}

	//else if vesicle.vesicleType == "Boolean" { !!!!!! TODO !!!!!!!!
	//
	//}
}

func (vesicle *Vesicle) DoIfReaction() {
	//Check that vesicle has if kinase
	if vesicle.ifKinase != nil {
		//Check that there are two inner vesicles
		if len(vesicle.vesicles) == 2 {
			//Check that one of the inner vesicle is a conditional vesicle
			var condVesicle *Vesicle
			var nonCondVesicle *Vesicle
			for _, inrVesicle := range vesicle.vesicles {
				if inrVesicle.name == COND {
					condVesicle = inrVesicle
				} else {
					nonCondVesicle = inrVesicle
				}
			}

			if condVesicle != nil {
				//Inactivate the receptors of non-conditional vesicle
				for _, receptor := range nonCondVesicle.receptorList {
					receptor.Inactivate()
				}

				//If phosphorylated, then activate the receptors of the other vesicle.
				if condVesicle.checkerKinase.GetPhosphoStatus() {
					//Inactivate the receptors of non-conditional vesicle
					for _, receptor := range nonCondVesicle.receptorList {
						receptor.Activate()
					}
				}
			}
			panic("There is no conditional vesicle inside")
		}
		panic("There are not enough vesicles inside!")
	}
	panic("There is no if kinase to carry out if reaction!")
}

func (vesicle *Vesicle) DoWhileReaction() {
	//Check that vesicle has if kinase
	if vesicle.ifKinase != nil {
		//Check that there are two inner vesicles
		if len(vesicle.vesicles) == 2 {
			//Check that one of the inner vesicle is a conditional vesicle
			var condVesicle *Vesicle
			var nonCondVesicle *Vesicle
			for _, inrVesicle := range vesicle.vesicles {
				if inrVesicle.name == COND {
					condVesicle = inrVesicle
				} else {
					nonCondVesicle = inrVesicle
				}
			}

			if condVesicle != nil {
				//Inactivate the receptors of non-conditional vesicle
				for _, receptor := range nonCondVesicle.receptorList {
					receptor.Inactivate()
				}

				//If phosphorylated, then activate the receptors of the other vesicle.
				if condVesicle.checkerKinase.GetPhosphoStatus() {
					//Inactivate the receptors of non-conditional vesicle
					for _, receptor := range nonCondVesicle.receptorList {
						receptor.Activate()
					}
				} // !!!!!!!! TODO !!!! YOU HAVE TO CHANGE THE LOCSIGNAL (DECREMENT MAYBE)
			}
			panic("There is no conditional vesicle inside")
		}
		panic("There are not enough vesicles inside!")
	}
	panic("There is no if kinase to carry out while reaction!")
}

func (vesicle *Vesicle) DoAssignmentReaction() {
	//Check that there is a glucosyltransferase
	if vesicle.glucoTrans != nil {
		//for each substrate, call transferGlucose
		for _, substrate := range vesicle.substrateList {
			vesicle.glucoTrans.TransferGlucose(substrate)
		}
	}
	panic("No glucosyltransferase to do assignment reaction!")
}

func (vesicle *Vesicle) DoConditionalReaction() {
	//Check that there is a checker kinase
	if vesicle.checkerKinase != nil {
		//Check that there are two substrates in the list. Otherwise no comparison
		//can be done
		if len(vesicle.substrateList) == 2 {
			//Initialize the two substrates
			substrate1 := vesicle.substrateList[0]
			substrate2 := vesicle.substrateList[1]
			//Check the glucoCount
			phospho := vesicle.checkerKinase.CheckGlucoCount(substrate1, substrate2)
			//Autophosphorylate
			if phospho {
				vesicle.checkerKinase.SetPhosphoStatus(true)
			} else {
				vesicle.checkerKinase.SetPhosphoStatus(false)
			}
		}
	}
}

func (vesicle *Vesicle) DoArithmeticReaction() {
	//Check that there is a glucosyltransferase
	if vesicle.glucoTrans != nil {
		//for each substrate, call transferGlucose
		for _, substrate := range vesicle.substrateList {
			vesicle.glucoTrans.TransferGlucose(substrate)
		}
	}
	panic("No glucosyltransferase to do arithmetic reaction!")
}

/******************************************************************************
																PROTEIN OBJECT
******************************************************************************/

//Protein is abstract parent object for all types of proteins
type Protein struct {
	name string
}

/******************************************************************************
																KINASE OBJECT
******************************************************************************/

//Kinase is a protein, and an abstract parent object for
//IfKinase and CheckerKinase
type Kinase struct {
	Protein
}

/******************************************************************************
																IF-KINASE OBJECT
******************************************************************************/

//IfKinase is a kinase check the recognizeVesicle
//if autophosphorylationstatus then activate the receptor
type IfKinase struct { //EVPSAK
	Kinase
	receptorName         string
	recognizeVesicleName string
}

//CheckPhosphoStatus check the phosphoStatus of input substrate
//if phosphorylated then return true
func (ifKinase *IfKinase) CheckPhosphoStatus(substrate *Substrate) bool {
	if (*substrate).phosphoStatus == true {
		return true
	}
	return false
}

//ActivateReceptor set receptor.phosphoStatus true
func (ifKinase *IfKinase) ActivateReceptor(receptor *Receptor) {
	(*receptor).phosphoStatus = true
}

/******************************************************************************
															CHECKER-KINASE OBJECT
******************************************************************************/

//CheckerKinase is a kinase that checks the glucoCount of two input protein
//if it's the same, phosphoStatus is true
type CheckerKinase struct { //IVSK
	Kinase
	checkerType   string
	phosphoStatus bool
}

//CheckGluCount checks the number of glucose on substrate
//if the # of glucose == # of input, return true
func (checkerKinase *CheckerKinase) CheckGlucoCount(substrate1 *Substrate,
	substrate2 *Substrate) bool {
	//Check the type of the transferase first and return tru accordingly
	if checkerKinase.checkerType == ">" {
		if substrate1.glucoCount > substrate2.glucoCount {
			return true
		}
		return false
	} else if checkerKinase.checkerType == "<" {
		if substrate1.glucoCount < substrate2.glucoCount {
			return true
		}
		return false
	} else if checkerKinase.checkerType == "<=" {
		if substrate1.glucoCount <= substrate2.glucoCount {
			return true
		}
		return false
	} else if checkerKinase.checkerType == ">=" {
		if substrate1.glucoCount >= substrate2.glucoCount {
			return true
		}
		return false
	} else if checkerKinase.checkerType == "!=" {
		if substrate1.glucoCount != substrate2.glucoCount {
			return true
		}
		return false
	} else if checkerKinase.checkerType == "==" {
		if substrate1.glucoCount == substrate2.glucoCount {
			return true
		}
		return false
	}
	return false
}

//AutophosphorylateStatus set the phosphostatus of substrate to true if input
//true
func (checkerKinase *CheckerKinase) SetPhosphoStatus(status bool) {
	if status {
		checkerKinase.phosphoStatus = true
	} else {
		checkerKinase.phosphoStatus = false
	}
}

func (checkerKinase *CheckerKinase) GetPhosphoStatus() bool {
	return checkerKinase.phosphoStatus
}

/******************************************************************************
																GLUCOTRANS OBJECT
******************************************************************************/

//Glucotrans is a protein transfer glucose to the protein
//arithmetic and assignment use this object
type Glucotrans struct {
	Protein
	glucoCount int
}

//TransferGlucose assign the number of substrate.glucoCount to the internal
//number that glucotrans has
func (glucotrans *Glucotrans) TransferGlucose(substrate *Substrate) {
	substrate.IncreaseGlu(glucotrans.glucoCount)
}

/******************************************************************************
																RECEPTOR OBJECT
******************************************************************************/

//Receptor is a protein recognize the input protein and decide
//whether it should be engulfed into the vesicle or not
type Receptor struct {
	Protein
	phosphoStatus bool
	locSignalRec  int //to recognize the localization signal on the substrate
}

func InitializeReceptor(receptor *Receptor, name string) *Receptor {
	receptor = new(Receptor)
	receptor.name = name
	receptor.phosphoStatus = true
	return receptor
}

func (receptor *Receptor) CheckPhosphoStatus() bool {
	return receptor.phosphoStatus
}

//CheckProtein check the name of protein and the receptor, if it matches to each
//other, then return true, if not, return false
func (receptor *Receptor) CheckProtein(proteinName string, locSignal int) bool {
	if proteinName == (*receptor).name && locSignal == receptor.locSignalRec {
		return true
	}
	return false
}

func (receptor *Receptor) Inactivate() {
	receptor.phosphoStatus = false
}

func (receptor *Receptor) Activate() {
	receptor.phosphoStatus = true
}

/******************************************************************************
																SUBSTRATE OBJECT
******************************************************************************/

//Substrate is a input protein, which is tagged with glucose and phospho group
type Substrate struct {
	Protein
	glucoCount    int
	phosphoStatus bool
	locSignal     int //this is to localize a protein to a specific vesicle
}

func InitializeSubstrate(substrate *Substrate, name string) *Substrate {
	substrate = new(Substrate)
	substrate.name = name
	substrate.locSignal = 0
	substrate.glucoCount = 0
	return substrate
}

//IncreaseGlu Increase the number of substrate.glucoCount to the number
func (substrate *Substrate) IncreaseGlu(number int) {
	(*substrate).glucoCount += number
}

//Phosphorylate assign the bool value true to the substrate.phosphoStatus
func (substrate *Substrate) Phosphorylate() {
	(*substrate).phosphoStatus = true
}

//DePhosphorylate assign the bool value false to the substrate.phosphoStatus
func (substrate *Substrate) DePhosphorylate() {
	(*substrate).phosphoStatus = false
}

//CheckPhosphoStatus return whether the substrate is phosphorylated,
//which is true
func (substrate *Substrate) CheckPhosphoStatus() bool {
	if (*substrate).phosphoStatus == true {
		return true
	}
	return false
}

//CheckGlucoNumber check the number of glucose on substrate
func (substrate *Substrate) CheckGlucoNumber() int {
	return (*substrate).glucoCount
}
