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
		data.Attributes = tmp
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

func testResult(k int, arr []knnGo.KNNData, inputAttributes []float64, distance string, p int) {
	fmt.Println()
	result, err := knnGo.KNNClassification(k, arr, inputAttributes, distance, p)
	fmt.Print(distance)
	if distance == strings.ToLower("minkowski") {
		fmt.Print(" order value : ", p)
	}
	fmt.Println()
	if err != nil {
		fmt.Println(err)
		return
	}
	result.PrintKNNData()
	fmt.Println(result.Classification)
}

func main() {
	arr := loadExcel("car.csv")

	//knnGo.PrintKNNDataset(arr)
	fmt.Println()

	input := []string{
		"v-high", "high", "4", "4", "med", "high",
	}
	inputfloat := inputProcess(input)
	fmt.Println(input)
	fmt.Println(inputfloat)
	testResult(5, arr, inputfloat, "euclidean", 420)
	testResult(5, arr, inputfloat, "manhattan", 30)
	testResult(5, arr, inputfloat, "supremum", 40)
	testResult(5, arr, inputfloat, "chebyshev", 20)
	testResult(5, arr, inputfloat, "cosine", 42)
	testResult(5, arr, inputfloat, "jaccard", 4202)
	testResult(5, arr, inputfloat, "minkowski", 4)
	testResult(5, arr, inputfloat, "minkowski", 2)
	testResult(5, arr, inputfloat, "minkowski", 0)
	testResult(5, arr, inputfloat, "sinus", 0)

	fmt.Println()
	fmt.Println("Error Cases : ")
	input = []string{
		"v-high", "high", "4", "4", "med",
	}
	inputfloat = inputProcess(input)
	fmt.Println(input)
	fmt.Println(inputfloat)
	testResult(5, arr, inputfloat, "euclidean", 420)
	testResult(5, arr, inputfloat, "manhattan", 30)
	testResult(5, arr, inputfloat, "supremum", 40)
	testResult(5, arr, inputfloat, "chebyshev", 20)
	testResult(5, arr, inputfloat, "cosine", 42)
	testResult(5, arr, inputfloat, "jaccard", 4202)
	testResult(5, arr, inputfloat, "minkowski", 4)
	testResult(5, arr, inputfloat, "minkowski", 2)
	testResult(5, arr, inputfloat, "minkowski", 0)
	testResult(5, arr, inputfloat, "sinus", 0)

}
