// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example2

// Sample program to create a dataframe and subsequently filter
// and subset the dataframe.
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kniren/gota/data-frame"
)

func main() {

	// Pull in the CSV data.
	irisData, err := ioutil.ReadFile("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Create a dataframe from the CSV string.
	// The types of the columns will be inferred.
	irisDF := df.ReadCSV(string(irisData))

	// Create a filter for the dataframe.
	filter := df.F{
		Colname:    "species",
		Comparator: "==",
		Comparando: "Iris-versicolor",
	}

	// Filter the dataframe to see only the rows where
	// the iris species is "Iris-versicolor".
	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err() != nil {
		log.Fatal(versicolorDF.Err())
	}

	// Output the results to standard out.
	fmt.Println(versicolorDF)

	// Filter the dataframe again, but only select out the
	// sepal_width and species columns.
	versicolorDF = irisDF.Filter(filter).Select("sepal_width", "species")
	fmt.Println(versicolorDF)

	// Filter and select the dataframe again, but only display
	// the first three results.
	versicolorDF = irisDF.Filter(filter).Select("sepal_width", "species").Subset([]int{0, 1, 2})
	fmt.Println(versicolorDF)

}
