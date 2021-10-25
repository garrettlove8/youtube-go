// 3. Our package follow the same convention
// 4. Our package name is "names" even though it is imported as a path to its source location
package names

// 2. Package names are one word, lowercased

// If a package only exports one struct, the constructor function should be "New()" not "NewObj()"
type Person struct {
	name string
}

// 1. lower case for private, uppercase for public
// Favor documentation comments over long, descriptive function names.
// Go can be verbose, don't unnecessarily make a mess

// New returns a new instance of a Person struct.
func New(name string) *Person {
	return &Person{
		name: name,
	}
}

// 5. Go doesn't have a built in concept of getters or setters.
// A getter should be named after the data it's getting, no need for the word "Get"
// But you should use the word "Set" if you need setters
func (p *Person) Name() string {
	return p.name
}

// 5. DO NOT use . notation when importing packages
//		This is useful for testing purposes, but will cause confusion if used in source code
//		This can cause collisions in with different packages being used in your code
