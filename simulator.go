package main

import "fmt"

/* ----------------------------------------------------------------------------
ProLANG Project
Author: Mert Inan
Date: 01 Dec 2018
Description:	This is the simulator code for the ProLANG project.
              Main function receives a static cell object and then runs
              Monte Carlo simulation on it. The function checks recursively
              every protein and vesicle to
-----------------------------------------------------------------------------*/

//CellBoard is a slice of Vesicle objects to store each iteration of the
//simulation.
type CellBoard []Vesicle

//SimulateCell simulates a given cell for numIter iterations. It returns every
//iteration on a cell array.
func SimulateCell(cell *Vesicle, numIter int) CellBoard {
	fmt.Println("beginning simulation!")
	//Create a cellBoard
	cellBoard := make(CellBoard, numIter)

	fmt.Println("before add substrates")
	//Add substrates to cell
	AddSubstrates(cell)

	fmt.Println("after add substrates")
	fmt.Println(cell.substrateList[0])

	//Monte Carlo Simulation Loop
	for i := 0; i < numIter; i++ {
		fmt.Println("before initialize vesicle!")
		//Initialize new Cell
		var newCell *Vesicle
		newCell = InitializeVesicle(newCell)

		fmt.Println("after initialize vesicle!")

		//Deep Copy new cell
		newCell.CopyVesicle(cell)
		fmt.Println("after deep copy!")

		//Update cell
		UpdateVesicle(newCell)
		fmt.Println("after update vesicle!")
	}

	//Return the cellboard
	return cellBoard
}

//AddSubstrates checks the receptor in all the nested vesicles and then adds
//found substrates from the top level cell.
func AddSubstrates(cell *Vesicle) {
	//initialize arrays
	substrates := make([]string, 0)
	vesicles := make([]*Vesicle, 0)
	vesicles = append(vesicles, cell)

	//Look for substrate proteins in all vesicles
	for len(vesicles) != 0 {
		var vesicle *Vesicle
		vesicle, vesicles = PopVesicle(vesicles)
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
		substrateProtein = InitializeSubstrate(substrateProtein, substrate)

		//Put the substrate into cell
		cell.substrateList = append(cell.substrateList, substrateProtein)
	}
}

//UpdateVesicle is a recursive function that updates the vesicle by following
//three steps: take in protein, do reaction and pupm out.
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
			vesicle.substrateList = append(vesicle.substrateList, inrVesicle.PumpOutProtein(substrate))
		}

		//Do it recursively for the children vesicles
		UpdateVesicle(inrVesicle)
	}
}

//PopVesicle deletes the initial element of the arr and returns it
func PopVesicle(vesicles []*Vesicle) (*Vesicle, []*Vesicle) {
	vesicle := vesicles[0]
	vesicles = append(vesicles[:0], vesicles[1:]...)
	return vesicle, vesicles
}

//InList checks whether the given name is in the given array arr.
func InList(name string, arr []string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == name {
			return true
		}
	}
	return false
}
