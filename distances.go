package knnGo

import (
	"errors"
	"math"
)

func MinkowskiDistance(dataset []KNNData, inputAttributes []float64, p int) ([]KNNData, error) {

	arr := dataset

	pfloat := float64(p)

	for i := 0; i < len(arr); i++ {
		if len(arr[i].Attributes) != len(inputAttributes) {
			return nil, errors.New("input attribute length not is matching with the data attribute length")
		}
		if p == 0 {
			return nil, errors.New("order cannot be 0")
		}
		sumAttribute := float64(0)
		for j := 0; j < len(arr[i].Attributes); j++ {
			margin := arr[i].Attributes[j] - inputAttributes[j]
			tmp := math.Pow(margin, pfloat) //margin * margin
			if tmp < 0 {
				tmp = -1 * tmp
			}
			sumAttribute += tmp
		}

		Distance := math.Pow(sumAttribute, 1/pfloat)
		arr[i].Distance = Distance
	}
	return arr, nil
}

func EuclideanDistance(dataset []KNNData, inputAttributes []float64) ([]KNNData, error) {

	return MinkowskiDistance(dataset, inputAttributes, 2)
}

func ManhattanDistance(dataset []KNNData, inputAttributes []float64) ([]KNNData, error) {

	return MinkowskiDistance(dataset, inputAttributes, 1)
}

func ChebyshevDistance(dataset []KNNData, inputAttributes []float64) ([]KNNData, error) {

	arr := dataset

	for i := 0; i < len(arr); i++ {
		if len(arr[i].Attributes) != len(inputAttributes) {
			return nil, errors.New("input attribute length not is matching with the data attribute length")
		}

		maxMargin := float64(0)
		for j := 0; j < len(arr[i].Attributes); j++ {
			margin := arr[i].Attributes[j] - inputAttributes[j]
			if margin < 0 {
				margin = -1 * margin
			}
			if margin >= maxMargin {
				maxMargin = margin
			}
		}
		arr[i].Distance = maxMargin
	}
	return arr, nil
}

func SupremumDistance(dataset []KNNData, inputAttributes []float64) ([]KNNData, error) {
	return ChebyshevDistance(dataset, inputAttributes)
}

func cosineSimilarity(dataAttributes []float64, inputAttributes []float64) float64 {
	arr := []float64{0, 0, 0}
	for j := 0; j < len(dataAttributes); j++ {
		subarr := []float64{
			dataAttributes[j] * inputAttributes[j],
			dataAttributes[j] * dataAttributes[j],
			inputAttributes[j] * inputAttributes[j],
		}
		for i := 0; i < len(arr); i++ {
			arr[i] += subarr[i]
		}
	}
	dataPM := math.Pow(arr[1], 0.5)
	inputPM := math.Pow(arr[2], 0.5)
	return (arr[0] / (dataPM * inputPM))
}

func CosineDistance(dataset []KNNData, inputAttributes []float64) ([]KNNData, error) {

	arr := dataset

	for i := 0; i < len(arr); i++ {
		if len(arr[i].Attributes) != len(inputAttributes) {
			return nil, errors.New("input attribute length not is matching with the data attribute length")
		}
		cosine := cosineSimilarity(arr[i].Attributes, inputAttributes)
		arr[i].Distance = 1 - cosine
	}

	return arr, nil
}

func jaccardCoefficient(dataAttributes []float64, inputAttributes []float64) float64 {
	sumMin := float64(0)
	sumMax := float64(0)
	for i := 0; i < len(dataAttributes); i++ {
		if dataAttributes[i] >= inputAttributes[i] {
			sumMin += inputAttributes[i]
			sumMax += dataAttributes[i]
		} else {
			sumMax += inputAttributes[i]
			sumMin += dataAttributes[i]
		}
	}
	return sumMin / sumMax
}

func JaccardDistance(dataset []KNNData, inputAttributes []float64) ([]KNNData, error) {

	arr := dataset
	//1 - (SUM(MIN) / SUM(MAX))
	for i := 0; i < len(arr); i++ {
		if len(arr[i].Attributes) != len(inputAttributes) {
			return nil, errors.New("input attribute length not is matching with the data attribute length")
		}
		jaccardCoef := jaccardCoefficient(arr[i].Attributes, inputAttributes)
		arr[i].Distance = 1 - jaccardCoef
	}
	return arr, nil
}
