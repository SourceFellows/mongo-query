package filter

import (
	"fmt"
)

func ExampleExpression_String_simpleEquals() {

	Filter := struct {
		Size struct {
			Uom Field
		}
	}{
		Size: struct {
			Uom Field
		}{
			Uom: Field("size.uom"),
		},
	}

	f1 := Filter.Size.Uom.Equals("in")

	fmt.Println(f1)
	// Output: bson.D{{"size.uom", "in"}}
}

func ExampleExpression_String_simpleLowerThan() {

	Filter := struct {
		Size struct {
			H Field
		}
	}{
		Size: struct {
			H Field
		}{
			H: Field("size.h"),
		},
	}

	f1 := Filter.Size.H.Lt(15)

	fmt.Println(f1)
	// Output: bson.D{{"size.h", bson.D{{"$lt", 15}}}}
}

func ExampleExpression_String_simpleAnd() {

	Filter := struct {
		Status Field
		Size   struct {
			H   Field
			Uom Field
		}
	}{
		Status: Field("status"),
		Size: struct {
			H   Field
			Uom Field
		}{
			Uom: Field("size.uom"),
			H:   Field("size.h"),
		},
	}

	f1 := Filter.Size.H.Lt(15).And(Filter.Size.Uom.Equals("in"))

	fmt.Println(f1)
	// Output: bson.D{{"$and", []bson.D{bson.D{{"size.h", bson.D{{"$lt", 15}}}}, bson.D{{"size.uom", "in"}}}}}

}

func ExampleExpression_String_simpleAndWithThreeParts() {

	Filter := struct {
		Status Field
		Size   struct {
			H   Field
			Uom Field
		}
	}{
		Status: Field("status"),
		Size: struct {
			H   Field
			Uom Field
		}{
			Uom: Field("size.uom"),
			H:   Field("size.h"),
		},
	}

	f1 := Filter.Size.H.Lt(15).And(Filter.Size.Uom.Equals("in"), Filter.Status.Equals("D"))

	fmt.Println(f1)
	// Output: bson.D{{"$and", []bson.D{bson.D{{"size.h", bson.D{{"$lt", 15}}}}, bson.D{{"size.uom", "in"}}, bson.D{{"status", "D"}}}}}

}

func ExampleExpression_String_simpleOr() {

	Filter := struct {
		Status Field
		Size   struct {
			H   Field
			Uom Field
		}
	}{
		Status: Field("status"),
		Size: struct {
			H   Field
			Uom Field
		}{
			Uom: Field("size.uom"),
			H:   Field("size.h"),
		},
	}

	f1 := Filter.Size.H.Lt(15).Or(Filter.Size.Uom.Equals("in"))

	fmt.Println(f1)
	// Output: bson.D{{"$or", []bson.D{bson.D{{"size.h", bson.D{{"$lt", 15}}}}, bson.D{{"size.uom", "in"}}}}}

}
