package main

/* ----------------------------------------------------------------------------
ProLANG Project
Author: Ian Lee
Date: 30 Nov 2018
Description:
-----------------------------------------------------------------------------*/

/******************************************************************************
																VESICLE OBJECT
******************************************************************************/

//Vesicle is the object of our function
type Vesicle struct {
	name         string
	vesicleType  string
	proteinList  []*Protein
	receptorList []*Receptor
	vesicles     []*Vesicle
}

func (vesicle *Vesicle) InitializeVesicle() {
	vesicle = new(Vesicle)
	vesicle.proteinList = make([]*Protein, 0)
	vesicle.receptorList = make([]*Receptor, 0)
	vesicle.vesicles = make([]*Vesicle, 0)
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
func (vesicle *Vesicle) TakeInProtein(protein Protein) {
	exist := false
	if vesicle.receptorList != nil {
		for _, otherProtein := range vesicle.receptorList {
			exist = (*otherProtein).CheckProtein(protein.name)
		}
		if exist {
			if vesicle.proteinList != nil {
				vesicle.proteinList = append(vesicle.proteinList, &protein)
			} else {
				vesicle.proteinList = make([]*Protein, 0)
				vesicle.proteinList = append(vesicle.proteinList, &protein)
			}
		}
	} else {
		panic("receptor list is empty!!")
	}
}

//PumpOutProtein remove the protein from proteinList
func (vesicle *Vesicle) PumpOutProtein(protein Protein) {
	if vesicle.proteinList != nil {
		for number, theProtein := range vesicle.proteinList {
			if theProtein.name == protein.name {
				vesicle.RemoveFromProteinList(number)
			}
		}
	} else {
		panic("protein list is empty")
	}
}

//RemoveFromProteinList delete the element from proteinList
func (vesicle *Vesicle) RemoveFromProteinList(number int) {
	vesicle.proteinList = append(vesicle.proteinList[:number],
		vesicle.proteinList[number+1:]...)
}

//DoReactionInside  ****SKIPPED****
func (vesicle *Vesicle) DoReactionInside(number int) {
	if vesicle.vesicleType == "WholeCell" {
		//Haven't decide what to do here
	} else if vesicle.vesicleType == "IfWhile" {

	} else if vesicle.vesicleType == "Assignment" {

	} else if vesicle.vesicleType == "Condition" {

	} else if vesicle.vesicleType == "Boolean" {

	}
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
																IFKINASE OBJECT
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
															CHECKERKINASE OBJECT
******************************************************************************/

//CheckerKinase is a kinase check the glucoCount of two input protein
//if it's the same, phosphoStatus is true
type CheckerKinase struct { //IVSK
	Kinase
	phosphoStatus bool
}

//CheckGluCount check the number of glucose on substrate
//if the # of glucose == # of intput, return true
func (checkerKinase *CheckerKinase) CheckGluCount(number int,
	substrate *Substrate) bool {
	if (*substrate).glucoCount == number {
		return true
	}
	return false
}

//AutophosphorylateStatus set the phosphostatus of substrate to true if input
//true
func (checkerKinase *CheckerKinase) AutophosphorylateStatus(substrate *Substrate,
	status bool) {
	if status {
		(*substrate).phosphoStatus = true
	} else {
		(*substrate).phosphoStatus = false
	}
}

/******************************************************************************
																GLUCOTRANS OBJECT
******************************************************************************/

//Glucotrans is a protein transfer glucose to the protein
//arithmetic and assignment use this object
type Glucotrans struct {
	Protein
}

//TransferGlucose assign the number of substrate.glucoCount to number
func (glucotrans *Glucotrans) TransferGlucose(substrate *Substrate, number int) {
	(*substrate).glucoCount = number
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

//CheckProtein check the name of protein and the receptor, if it matches to each
//other, then return true, if not, return false
func (receptor *Receptor) CheckProtein(proteinName string) bool {
	if proteinName == (*receptor).name {
		return true
	}
	return false
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

//IncreaseGlu Increase the number of substrate.glucoCount to the number
func (substrate *Substrate) IncreaseGlu(number int) {
	for i := 0; i < number; i++ {
		(*substrate).glucoCount++
	}
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
