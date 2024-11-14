package vara

// constructor is a function that takes any number of dependencies
// as its parameters and build then returns an arbitrary number of values of
// one or more type and may optionally return an error to indicate that it failed to build the value(s).
//
// Any arguments that the constructor has are treated as its dependencies. The dependencies are instantiated
// in an unspecified order along with any dependencies that they might have.
type constructor any
