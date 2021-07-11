# KNN Classification Algorithm using Go

A basic K-Nearest Neighbors Classification Algorithm using Go by kamil5b

### TABLE OF CONTENT
- [How it work](#how-it-work)
  * [Processing Dataset](#processing-dataset)
  * [Processing Data Input](#processing-data-input)
    + [Euclidean Distance](#euclidean-distance)
    + [Mantahattan Distance](#mantahattan-distance)
    + [Minkowski Distance](#minkowski-distance)
    + [Chebyshev/Supremum Distance](#chebyshev-or-supremum-distance)
    + [Cosine Distance](#cosine-distance)
    + [Jaccard Distance](#jaccard-distance)
- [How to use](#how-to-use)

# How it work

The inputs for this package are a dataset that contain a classification column and the attributes columns and an unclassified data input based on the dataset cases

## Processing Dataset
```Go
type KNNData struct {
	Classification string
	Attributes        []float64
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
- Mantahattan Distance
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
