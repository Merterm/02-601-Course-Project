package main

/* ----------------------------------------------------------------------------
ProLANG Project
Author: Mert Inan
Date: 01 Dec 2018
Description:	This is the parser code for the ProLANG project.
              This program reads a specific code in ProLANG and generates the
              parse tree for it. Highest level function is the Parse(filename)
              function. It takes the filename of the ProLANG file.
-----------------------------------------------------------------------------*/

type CellBoard []Vesicle

func SimulateCell(cell *Vesicle, numIter int) CellBoard {
	//Create a cellBoard
	cellBoard := make(CellBoard, numIter)

	//Add substrates to cell
	AddSubstrates(cell)

	//Monte Carlo Simulation Loop
	for i := 0; i < numIter; i++ {
		//Initialize new Cell
		var newCell *Vesicle
		newCell.InitializeVesicle()

		//Deep Copy new cell
		newCell.CopyVesicle(cell)

		//Update cell
		UpdateVesicle(newCell)
	}

	//Return the cellboard
	return cellBoard
}

func AddSubstrates(cell *Vesicle) {
	//initialize arrays
	substrates := make([]string, 0)
	vesicles := make([]*Vesicle, 0)
	vesicles = append(vesicles, cell)

	//Look for substrate proteins in all vesicles
	for len(vesicles) != 0 {
		vesicle := PopVesicle(vesicles)
		for _, receptor := range vesicle.receptorList {
			name := receptor.name
			if !InList(name, substrates) {
				substrates = append(substrates, name)
			}
		}
		for _, inrVesicle := range vesicle.vesicles {
			vesicles = append(vesicles, inrVesicle)
		}
	}

	//Add substrates to cell
	for _, substrate := range substrates {
		//Create new substrate object
		var substrateProtein *Substrate
		substrateProtein.InitializeSubstrate(substrate)

		//Put the substrate into cell
		//cell.proteinList = append(cell.proteinList, substrateProtein) HEEEEEEEERRRRRRRREEEEEEEE
	}
}

func UpdateVesicle(vesicle *Vesicle) {
	for _, inrVesicle := range vesicle.vesicles {
		for _, substrate := range vesicle.substrateList {
			//Take in the protein if possible
			inrVesicle.TakeInProtein(substrate)
		}
		//Do the reaction
		inrVesicle.DoReactionInside()

		//Get the processed protein out
		for _, substrate := range inrVesicle.substrateList {
			inrVesicle.PumpOutProtein(substrate)
		}

		//Do it recursively for the children vesicles
		UpdateVesicle(inrVesicle)
	}
}

//Pop deletes the initial element of the arr and returns it
func PopVesicle(vesicles []*Vesicle) *Vesicle {
	vesicle := vesicles[0]
	vesicles = append(vesicles[:0], vesicles[1:]...)
	return vesicle
}

func InList(name string, arr []string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == name {
			return true
		}
	}
	return false
}
