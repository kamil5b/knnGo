package knnGo

import "math"

func MinskowskiDistance(dataset []KNNData, inputFactors []float64, p int) []KNNData {

	arr := dataset

	pfloat := float64(p)

	for i := 0; i < len(arr); i++ {
		if len(arr[i].Factors) != len(inputFactors) {
			return nil
		}
		sumfactor := float64(0)
		for j := 0; j < len(arr[i].Factors); j++ {
			margin := arr[i].Factors[j] - inputFactors[j]
			tmp := math.Pow(margin, pfloat) //margin * margin
			if tmp < 0 {
				tmp = -1 * tmp
			}
			sumfactor += tmp
		}

		Distance := math.Pow(sumfactor, 1/pfloat)
		arr[i].Distance = Distance
	}
	return arr
}

func EuclideanDistance(dataset []KNNData, inputFactors []float64) []KNNData {

	return MinskowskiDistance(dataset, inputFactors, 2)
}

func ManhattanDistance(dataset []KNNData, inputFactors []float64) []KNNData {

	return MinskowskiDistance(dataset, inputFactors, 1)
}

func SupremumDistance(dataset []KNNData, inputFactors []float64) []KNNData {

	arr := dataset

	for i := 0; i < len(arr); i++ {
		if len(arr[i].Factors) != len(inputFactors) {
			return nil
		}

		maxMargin := float64(0)
		for j := 0; j < len(arr[i].Factors); j++ {
			margin := arr[i].Factors[j] - inputFactors[j]
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
