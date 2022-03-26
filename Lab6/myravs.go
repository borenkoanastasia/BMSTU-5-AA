package main

const PHEROMON_MIN = 0.5

func (acolony ant_colony_t)make_feramons(size int, p float64, ){
	var m matrix_t = make_empty_matrix(size, size)
	for i:=0;i<size;i++{
		for j:=0;j<size;j++{
			m.elem[i][j] = PHEROMON_MIN
		}
	}
	acolony.pheromonths = m
}

func myravs(){}