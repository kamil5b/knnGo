# KNN Classification Algorithm using Go

A basic K-Nearest Neighbors Classification Algorithm using Go by kamil5b

## How it work

The inputs for this package are a dataset that contain a classification column and the factors columns and an unclassified data input based on the dataset cases

### Processing the dataset
```Go
type KNNData struct {
	Classification string
	Factors        []float64
	Distance       float64
}
```

For each data row in a dataset there will be a classification column and the factors columns and they are casted to the KNNData type.
Each factors **have** to be casted to float64 before cast it to the KNNData Factors.

Each data row will be stored in the KNNData array.

###Processing the input
The input is an unclassified data containing factors that similar to or based on the dataset that we've worked on.

Calculate the distance of the input data with the rest of data in the dataset using 

## How to use :

1. Have a function for fitnessing the chromosomes 
make sure the parameters are (float32,float32) and returning float32:
```Go
equation := func(x,y float32) float32 {
  return ((x*y)+y)/(x+y)
}
```
1.1. This part is optional but if you don't want to hardcode it you could add some information about how many evolutions you want to observe, how much individuals within a population (population size), how many mating processes, what are the chance of the chromosome mutates, maximum and minimum value of x, maximum and minimum of y. As in right now version 1.0 maximum for x and y is 9.999 and the minimum is -9.999
```Go
var (
  evolution, totalPopulation, matingProcess int
  mutationChance, Xmax, Xmin, Ymax, Ymin float32
 )

```

2. Have an object/unallocated variable with basicGoGA.Generation data type :
```Go
var generation basicGoGA.Generation
```
2.1 Create the first generation with the resources that you had
```Go
generation = basicGoGA.createGeneration(totalPopulation, matingProcess, mutationChance, Xmax, Xmin, Ymax, Ymin, equation)
```
2.2 generate a generation based on how many evolutions you've inputted
```Go
generation = basicGoGA.generateGenerations(evolution, totalPopulation, matingProcess, mutationChance, Xmax, Xmin, Ymax, Ymin, equation)
```
