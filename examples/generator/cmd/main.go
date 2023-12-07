package main

import (
	"fmt"
	"github.com/sourcefellows/mongo-query/examples/generator/filter"
)

func main() {

	//filter embedded types
	fmt.Println(filter.ListingAndReview.Address.Location.Type.Equals("test"))

	//"Specify a Query Condition on a Field Embedded in an Array of Documents"
	//take the second element and query on that
	fmt.Println(filter.Reviews.ElementNo(2).ReviewerName.Equals("Milo"))

}
