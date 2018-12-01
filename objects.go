package main

//Vesicle is the object of our function
type Vesicle struct {
	name         string
	vesicleType  string
	proteinList  []*Protein
	receptorList []*Protein
}

//Protein is abstract parent object for all types of proteins
type Protein struct {
	name string
}

//Kinase is a protein, and an abstract parent object for
//IfKinase and CheckerKinase
type Kinase struct {
	Protein
}

//IfKinase is a kinase check the recognizeVesicle
//if autophosphorylationstatus then activate the receptor
type IfKinase struct { //EVPSAK
	Kinase
	recetporName         string
	recognizeVesicleName string
}

//CheckerKinase is a kinase check the glucoCount of two input protein
//if it's the same, phosphoStatus is true
type CheckerKinase struct { //IVSK
	Kinase
	phosphoStatus bool
}

//Glucotrans is a protein transfer glucose to the protein
//arithmetic and assignment use this object
type Glucotrans struct {
	Protein
	name string
}

//Receptor is a protein recognize the input protein and decide
//whether it should be engulfed into the vesicle or not
type Receptor struct {
	Protein
	name          string
	phosphoStatus bool
}

//Substrate is a input protein, which is tagged with glucose and phospho group
type Substrate struct {
	Protein
	name          string
	glucoCount    int
	phosphoStatus bool
}

//TakeInProtein add protein to proteinList if it could be recognized by the receptor
func (vesicle *Vesicle) TakeInProtein(protein Protein) {
	exist := false
	if vesicle.receptorList != nil {
		for _, otherProtein := range vesicle.receptorList {
			if protein.name == (*otherProtein).name {
				exist = true
			}
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
	vesicle.proteinList = append(vesicle.proteinList[:number], vesicle.proteinList[number+1:]...)
}

//DoReactionInside
func (vesicle *Vesicle) DoReactionInside(number int) {
	if vesicle.vesicleType == "WholeCell" {
		//Haven't decide what to do here
	} else if vesicle.vesicleType == "IfWhile" {

	} else if vesicle.vesicleType == "Assignment" {
		TransferGlucose(number)
	} else if vesicle.vesicleType == "Condition" {

	} else if vesicle.vesicleType == "Boolean" {

	}
}

func (glucotrans *Glucotrans) TransferGlucose(number int) {

}

func (receptor *Receptor) CheckProtein() {

}

func (substrate *Substrate) IncreaseGlu() {

}

func (substrate *Substrate) Phosphorylate() {

}

func (substrate *Substrate) CheckPhosphoStatus() {

}

func (substrate *Substrate) CheckGlucoNumber() {

}

func (ifKinase *IfKinase) CheckPhosphoStatus() {

}

func (ifKinase *IfKinase) ActivateReceptor() {

}

func (checkerKinase *CheckerKinase) CheckGluCount() {

}

func (checkerKinase *CheckerKinase) AutophosphorylateStatue() {

}
