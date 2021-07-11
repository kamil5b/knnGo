package knnGo

import (
	"errors"
	"fmt"
	"strings"
)

type KNNData struct {
	Classification string
	Attributes     []float64
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
	for i := 0; i < len(d.Attributes); i++ {
		fmt.Print(d.Attributes[i], " ")
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

func KNNClassification(k int, dataset []KNNData, inputAttributes []float64, distance string, p int) (KNNData, error) {

	var d KNNData
	var err error = nil
	classes := make([]string, 0)
	arr := dataset
	if strings.ToLower(distance) == "euclidean" {
		arr, err = EuclideanDistance(dataset, inputAttributes)
	} else if strings.ToLower(distance) == "manhattan" {
		arr, err = ManhattanDistance(dataset, inputAttributes)
	} else if strings.ToLower(distance) == "supremum" {
		arr, err = SupremumDistance(dataset, inputAttributes)
	} else if strings.ToLower(distance) == "chebyshev" {
		arr, err = ChebyshevDistance(dataset, inputAttributes)
	} else if strings.ToLower(distance) == "cosine" {
		arr, err = CosineDistance(dataset, inputAttributes)
	} else if strings.ToLower(distance) == "jaccard" {
		arr, err = JaccardDistance(dataset, inputAttributes)
	} else if strings.ToLower(distance) == "minkowski" {
		arr, err = MinkowskiDistance(dataset, inputAttributes, p)
	} else {
		return KNNData{}, errors.New("invalid distance metrics or knnGo haven't implemented yet")
	}
	if err != nil {
		return KNNData{}, err
	}
	arr = KNNDataSort(arr, true)
	arr = arr[:k]
	for i := 0; i < len(arr); i++ {
		classes = append(classes, arr[i].Classification)
	}
	d.Classification = ValueVote(classes)
	d.Attributes = inputAttributes
	return d, nil

}
