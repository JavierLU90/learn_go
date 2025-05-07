package main

// Generics:
/* Go does not support classes. 
For a long time, that meant that Go code couldn't easily be reused in many cases. 
For example, imagine some code that splits a slice into 2 equal parts. 
The code that splits the slice doesn't care about the types of values stored in the slice. 
Before generics, we needed to write the same code for each type, which is a very un-DRY thing to do.

In Go 1.18 however, support for generics was released, effectively solving this problem!*/


// Type Parameters:
// Put simply, generics allow us to use variables to refer to specific types. 
// This is an amazing feature because it allows us to write abstract functions that drastically reduce code duplication.
func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
// In the example above, T is the name of the type parameter for the splitAnySlice function, 
// and we've said that it must match the any constraint, which means it can be anything. 
// This makes sense because the body of the function doesn't care about the types of things stored in the slice.
firstInts, secondInts := splitAnySlice([]int{0, 1, 2, 3})
fmt.Println(firstInts, secondInts)

// Creating a variable that's the zero value of a type is easy:
var myZeroInt int
// It's the same with generics, we just have a variable that represents the type:
var myZero T


// Constraints:
// Constraints are just interfaces that allow us to write generics that only operate 
// within the constraints of a given interface type. 
// In the example above, the any constraint is the same as the empty interface because it means the type in question can be anything.
type stringer interface {
    String() string
}

func concat[T stringer](vals []T) string {
    result := ""
    for _, val := range vals {
        // this is where the .String() method
        // is used. That's why we need a more specific
        // constraint instead of the any constraint
        result += val.String()
    }
    return result
}


// Interface Type Lists
// When generics were released, a new way of writing interfaces was also released at the same time!
// We can now simply list a bunch of types to get a new interface/constraint.
type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
        ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
        ~float32 | ~float64 |
        ~string
}
// Ordered is a type constraint that matches any ordered type.
// An ordered type is one that supports the <, <=, >, and >= operators.


// Parametric Constraints: 
// Your interface definitions, which can later be used as constraints, can accept type parameters as well.
type store[P product] interface { // The store interface represents a store that sells products.
	Sell(P) // It takes a type parameter P that represents the type of products the store sells.
}

type product interface {
	Price() float64
	Name() string
}

type book struct {
	title  string
	author string
	price  float64
}

func (b book) Price() float64 {
	return b.price
}

func (b book) Name() string {
	return fmt.Sprintf("%s by %s", b.title, b.author)
}

type toy struct {
	name  string
	price float64
}

func (t toy) Price() float64 {
	return t.price
}

func (t toy) Name() string {
	return t.name
}

// The bookStore struct represents a store that sells books.
type bookStore struct {
	booksSold []book
}

// Sell adds a book to the bookStore's inventory.
func (bs *bookStore) Sell(b book) {
	bs.booksSold = append(bs.booksSold, b)
}

// The toyStore struct represents a store that sells toys.
type toyStore struct {
	toysSold []toy
}

// Sell adds a toy to the toyStore's inventory.
func (ts *toyStore) Sell(t toy) {
	ts.toysSold = append(ts.toysSold, t)
}

// sellProducts takes a store and a slice of products and sells
// each product one by one.
func sellProducts[P product](s store[P], products []P) {
	for _, p := range products {
		s.Sell(p)
	}
}

func main() {
	bs := bookStore{
		booksSold: []book{},
	}

    // By passing in "book" as a type parameter, we can use the sellProducts function to sell books in a bookStore
	sellProducts[book](&bs, []book{
		{
			title:  "The Hobbit",
			author: "J.R.R. Tolkien",
			price:  10.0,
		},
		{
			title:  "The Lord of the Rings",
			author: "J.R.R. Tolkien",
			price:  20.0,
		},
	})
	fmt.Println(bs.booksSold)

    // We can then do the same for toys
	ts := toyStore{
		toysSold: []toy{},
	}
	sellProducts[toy](&ts, []toy{
		{
			name:  "Lego",
			price: 10.0,
		},
		{
			name:  "Barbie",
			price: 20.0,
		},
	})
	fmt.Println(ts.toysSold)
}
