package knnGo

import "math"

func MinskowskiDistance(dataset []KNNData, inputAttributes []float64, p int) []KNNData {

	arr := dataset

	pfloat := float64(p)

	for i := 0; i < len(arr); i++ {
		if len(arr[i].Attributes) != len(inputAttributes) {
			return nil
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
	return arr
}

func EuclideanDistance(dataset []KNNData, inputAttributes []float64) []KNNData {

	return MinskowskiDistance(dataset, inputAttributes, 2)
}

func ManhattanDistance(dataset []KNNData, inputAttributes []float64) []KNNData {

	return MinskowskiDistance(dataset, inputAttributes, 1)
}

func ChebyshevDistance(dataset []KNNData, inputAttributes []float64) []KNNData {

	arr := dataset

	for i := 0; i < len(arr); i++ {
		if len(arr[i].Attributes) != len(inputAttributes) {
			return nil
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
	return arr
}

func SupremumDistance(dataset []KNNData, inputAttributes []float64) []KNNData {
	return ChebyshevDistance(dataset, inputAttributes)
}
