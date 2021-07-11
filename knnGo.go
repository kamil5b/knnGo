package knnGo

import (
	"fmt"
	"strings"
)

type KNNData struct {
	Classification string
	Factors        []float64
	Distance       float64
}

func KNNDataSort(dataset []KNNData, ascending bool) []KNNData {
	//insertionSort
	arr := dataset
	var n = len(arr)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if ascending {
				if arr[j-1].Distance > arr[j].Distance {
					arr[j-1], arr[j] = arr[j], arr[j-1]
				}
			} else {
				if arr[j-1].Distance < arr[j].Distance {
					arr[j-1], arr[j] = arr[j], arr[j-1]
				}
			}
			j = j - 1
		}
	}
	return arr
}

func (d KNNData) PrintKNNData() {
	fmt.Print(d.Classification, " ")
	for i := 0; i < len(d.Factors); i++ {
		fmt.Print(d.Factors[i], " ")
	}
	if d.Distance >= 0 {
		fmt.Print(d.Distance)
	}
	fmt.Println()
}

func PrintKNNDataset(arr []KNNData) {
	for i := 0; i < len(arr); i++ {
		arr[i].PrintKNNData()
	}
}

func ValueVote(arr []string) string {

	dict := make(map[string]int)
	maxval := 0
	votewin := ""
	for _, nama := range arr {
		dict[nama] = dict[nama] + 1
	}

	for k := range dict {
		if dict[k] > maxval {
			maxval = dict[k]
			votewin = k
		}
	}
	return votewin
}

/*
func MinMaxNormalize(arr []KNNData) {

}*/

func KNNClassification(k int, dataset []KNNData, inputFactors []float64, distance string, p int) KNNData {

	var d KNNData

	classes := make([]string, 0)
	arr := dataset
	if strings.ToLower(distance) == "euclidean" {
		arr = EuclideanDistance(dataset, inputFactors)
	} else if strings.ToLower(distance) == "manhattan" {
		arr = ManhattanDistance(dataset, inputFactors)
	} else if strings.ToLower(distance) == "supremum" {
		arr = SupremumDistance(dataset, inputFactors)
	} else if strings.ToLower(distance) == "minskowski" {
		arr = MinskowskiDistance(dataset, inputFactors, p)
	} else {
		return d
	}
	arr = KNNDataSort(arr, true)
	arr = arr[:k]
	for i := 0; i < len(arr); i++ {
		classes = append(classes, arr[i].Classification)
	}
	d.Classification = ValueVote(classes)
	d.Factors = inputFactors
	return d

}
