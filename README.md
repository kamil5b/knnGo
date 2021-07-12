# KNN Classification Algorithm using Go

A basic K-Nearest Neighbors Classification Algorithm using Go by kamil5b

### TABLE OF CONTENT
<!--
- [How it work](#how-it-work)
  * [Processing Dataset](#processing-dataset)
  * [Processing Data Input](#processing-data-input)
    + [Euclidean Distance](#euclidean-distance)
    + [Mantahattan Distance](#mantahattan-distance)
    + [Minkowski Distance](#minkowski-distance)
    + [Chebyshev/Supremum Distance](#chebyshev-or-supremum-distance)
    + [Cosine Distance](#cosine-distance)
    + [Jaccard Distance](#jaccard-distance)
- [How to use](#how-to-use)-->
- [How it work](#how-it-work)
  * [Processing Dataset](#processing-dataset)
  * [Processing Data Input](#processing-data-input)
    + [Euclidean Distance](#euclidean-distance)
    + [Manhattan Distance](#manhattan-distance)
    + [Minkowski Distance](#minkowski-distance)
    + [Chebyshev Or Supremum Distance](#chebyshev-or-supremum-distance)
    + [Cosine Distance](#cosine-distance)
    + [Jaccard Distance](#jaccard-distance)
- [How to use](#how-to-use)
  * [knnGo.go types and functions](#knngogo-types-and-functions)
    + [type KNNData struct](#type-knndata-struct)
    + [func KNNClassification](#func-knnclassification)
    + [func (d KNNData) PrintKNNData](#func--d-knndata--printknndata)
    + [func ValueVote](#func-valuevote)
  * [distance.go functions](#distancego-functions)
    + [func EuclideanDistance](#func-euclideandistance)
    + [func ManhattanDistance](#func-manhattandistance)
    + [func MinkowskiDistance](#func-minkowskidistance)
    + [func ChebyshevDistance](#func-chebyshevdistance)
    + [func SupremumDistance](#func-supremumdistance)
    + [func CosineDistance](#func-cosinedistance)
    + [func JaccardDistance](#func-jaccarddistance)

# How it work

The inputs for this package are a dataset that contain a classification column and the attributes columns and an unclassified data input based on the dataset cases

## Processing Dataset
```Go
type KNNData struct {
	Classification string
	Attributes     []float64
	Distance       float64
}
```

For each data row in a dataset there will be a classification column and the attributes columns and they are casted to the KNNData type.
Each attributes **have** to be casted to float64 before cast it to the KNNData Attributes.

Each data row will be stored in the KNNData array.

## Processing Data Input

The input is an unclassified data containing attributes that similar with or based on the dataset that we've worked on.

Calculate the distance of the input data with the rest of data in the dataset using :

- Euclidean Distance
- Manhattan Distance
- Minkowski Distance
- Chebyshev/Supremum Distance
- Cosine Distance
- Jaccard Distace

### Euclidean Distance

Euclidean distance is 

### Mantahattan Distance

Mantahattan distance is

### Minkowski Distance

Minkowski distance is

### Chebyshev Or Supremum Distance

Chebyshev a.k.a Supremum distance is

### Cosine Distance

Cosine distance is

### Jaccard Distance

Jaccard distance is

# How to use

To use this package, you have to download and install this repository.

Write this to your terminal :
```sh
go get github.com/kamil5b/knnGo
```
If you want to use this package to your code you also need to import it
```Go
import github.com/kamil5b/knnGo
```

In order this package work, the dataset you are using **must** be converted to an array of [KNNData](#type-knndata-struct) where the classification is a string and the attributes are float64. **(if your attributes data is ranked enum data, you may convert it to an integer array that contains the ranking of the enum which then convert it to discrete float64)**

## knnGo.go types and functions
### type KNNData struct
```Go
type KNNData struct {
	Classification string
	Attributes     []float64
	Distance       float64
}
```
### func KNNClassification
```Go
func KNNClassification(k int, dataset []KNNData, inputAttributes []float64, distance string, p int) (KNNData, error)
```
This is the main function for this package. This function returning a KNNData containing the classification of the input data and for the attributes are the input data attribute itself. 

- k is a value for how much neighbor we observe
- dataset in this parameter is an array of KNNData which is converted from dataset that we had
- inputAttributes is an array of float64 that contains the input attributes
- distance is an enum that converted into non-case-sensitive string. The available distances are :
  * [euclidean](#func-euclideandistance)
  * [manhattan](#func-manhattandistance)
  * [minkowski](#func-minkowskidistance)
  * [chebyshev](#func-chebyshevdistance)
  * [supremum](#func-supremumdistance)
  * [cosine](#func-cosinedistance)
  * [jaccard](#func-jaccarddistance)
- p is an integer for minkowski's distance order or the p value. the value doesn't matter if you are not choosing minkowski distance

This function will calculate the distance of the input data with every data in the duplicated dataset and sort ascending based on the distances and then pick the top-k of it and then the classification for the data input is the most populated class among the top-k.

### func (d KNNData) PrintKNNData
```Go
func (d KNNData) PrintKNNData()
```
This function is for printing the KNNData

### func ValueVote
```Go
func ValueVote(arr []string) string
```

## distance.go functions
Functions in this file containing the metric distances that can be used in this package, it return an array of KNNData that each data containing the result of the distances.

### func [EuclideanDistance](#euclidean-distance)
```Go
func EuclideanDistance(dataset []KNNData, inputAttributes []float64) ([]KNNData, error)
```

### func [ManhattanDistance](#manhattan-distance)
```Go
func ManhattanDistance(dataset []KNNData, inputAttributes []float64) ([]KNNData, error)
```

### func [MinkowskiDistance](#minkowski-distance)
```Go
func MinkowskiDistance(dataset []KNNData, inputAttributes []float64, p int) ([]KNNData, error)
```

### func [ChebyshevDistance](#chebyshev-or-supremum-distance)
```Go
func MinkowskiDistance(dataset []KNNData, inputAttributes []float64) ([]KNNData, error)
```

### func [SupremumDistance](#chebyshev-or-supremum-distance)
```Go
func MinkowskiDistance(dataset []KNNData, inputAttributes []float64) ([]KNNData, error)
```

### func [CosineDistance](#cosine-distance)
```Go
func MinkowskiDistance(dataset []KNNData, inputAttributes []float64) ([]KNNData, error)
```

### func [JaccardDistance](#jaccard-distance)
```Go
func MinkowskiDistance(dataset []KNNData, inputAttributes []float64) ([]KNNData, error)
```
