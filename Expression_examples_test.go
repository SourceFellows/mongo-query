/**
 * MIT License
 *
 * Copyright (c) 2023 Source Fellows GmbH (https://www.source-fellows.com)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

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