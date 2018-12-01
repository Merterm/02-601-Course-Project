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
	//Look for substrate proteins in all vesicles

	//
}

func UpdateVesicle(vesicle *Vesicle) {

}
