package main

import (
//	"fmt"
	"solid_design/dsa/graph"
	"solid_design/dsa/sorting"
//	"solid_design/dsa/tree"

)

func Sort(a []int, strategy sorting.SortingStrategy) []int {
	s := sorting.Sorter{}
	s.SetSortingStrategy(strategy)
	return s.Strategy.Sort(a)

}

func getEdges() [][]int { 
	return [][]int {
		{1,2},
		{1,3},
		{2,4},
		{2,5},
		{3,6},
		{4,5},
		{5,6},
	}
}

/**
lets say graph be like 


		1
	    /\
	    2 3
	   /\ \
      4-5-6


**/

func getWeightedEdges() []graph.EdgeWeighted { 

	return []graph.EdgeWeighted {
		graph.NewEdgeWeighted(1,2,4),
		graph.NewEdgeWeighted(1,3,8),
		graph.NewEdgeWeighted(2, 4,  1),
		graph.NewEdgeWeighted(2, 5, 5),
		graph.NewEdgeWeighted(3, 6, 4),
		graph.NewEdgeWeighted(4, 5, 9),
		graph.NewEdgeWeighted(5, 6, 1),
	}
}


func main() {
	graph.Kruskal(getWeightedEdges(),6)
   
}
