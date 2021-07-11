package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/kamil5b/knnGo"
)

var attributeValue = map[string]float64{
	"low":    0,
	"med":    1,
	"high":   2,
	"v-high": 3,
	"2":      0,
	"3":      1,
	"4":      2,
	"5-more": 3,
	"small":  0,
	"big":    2,
}

var personArr = []string{"2", "4", "more"}

func loadExcel(filename string) []knnGo.KNNData {
	var (
		arr []knnGo.KNNData
	)

	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	csvdata := make([][]string, 0)
	for _, heythere := range csvLines {
		s := strings.Split(heythere[0], ";")
		csvdata = append(csvdata, s)
	}

	for _, line := range csvdata {
		data := knnGo.KNNData{Distance: 0}
		tmp := []float64{}
		for _, factor := range line {
			if factor == line[len(line)-1] {
				data.Classification = line[len(line)-1] //acc unacc
			} else if factor == line[3] {
				person := func() float64 {
					a := int(0)
					for _, r := range personArr {
						if string(factor) == r {
							return float64(a)
						} else {
							a++
						}
					}
					return float64(len(personArr) - 1)
				}
				tmp = append(tmp, person())
			} else {
				tmp = append(tmp, attributeValue[string(factor)])
			}
		}
		data.Factors = tmp
		arr = append(arr, data)
	}
	return arr
}

func inputProcess(line []string) []float64 {
	tmp := []float64{}
	for _, factor := range line {
		if factor == line[3] {
			person := func() float64 {
				a := int(0)
				for _, r := range personArr {
					if string(factor) == r {
						return float64(a)
					} else {
						a++
					}
				}
				return float64(len(personArr) - 1)
			}
			tmp = append(tmp, person())
		} else {
			tmp = append(tmp, attributeValue[string(factor)])
		}
	}
	return tmp
}

func main() {
	arr := loadExcel("car.csv")

	knnGo.PrintKNNDataset(arr)
	fmt.Println()
	input := []string{
		"v-high", "high", "4", "4", "med", "high",
	}
	inputfloat := inputProcess(input)
	fmt.Println(input)
	fmt.Println(inputfloat)
	/*
		knn := knnGo.EuclideanDistance(arr, inputfloat)
		knnGo.PrintKNNDataset(knn)
		fmt.Println()
		knnsorted := knnGo.KNNDataSort(arr, true)
		knnGo.PrintKNNDataset(knnsorted)
		fmt.Println()
		knnsorted[0].PrintKNNData()*/
	result := knnGo.KNNClassification(5, arr, inputfloat, "euclidean", 0)
	result.PrintKNNData()
	print(result.Classification)
}
