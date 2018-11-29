package main

//Vesicle is the object of our function
type Vesicle struct {
	name         string
	proteinList  []*Protein
	receptorList []*Receptor
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

func (vesicle *Vesicle) TakeInProtein() {

}

func (vesicle *Vesicle) PumpOutProtein() {

}

func (vesicle *Vesicle) DoReactionInside() {

}

func (glucotrans *Glucotrans) TrnasferGlucose() {

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
