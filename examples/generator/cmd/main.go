package main

import (
	"fmt"
	"github.com/sourcefellows/mongo-query/examples/generator/filter"
)

func main() {

	//filter embedded types
	fmt.Println(filter.ListingAndReview.Address.Location.Type.Equals("test"))

}
