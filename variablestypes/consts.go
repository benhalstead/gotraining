package main

//Consts are immutable values that can be declared in various scopes
//They are generally decared at the package level with either package scope (lowercase first character)

const private = 2

// Or public (uppercase first character)

const Public = "HELLO"

// consts have the same rules of type inference as other variables - you can either specify the type or let the compiler work
// it out

const PUB_STRING string = "PS"

// You can ONLY use bools, strings and numbers as consts

// You can define multiple consts in a single block

const (
	SOME_CONST = 1
	SOME_OTHER = 3
)

// Go has no explicit support for enums, but these can be simulated with this pattern (type aliasing is discussed elsewhere)
type MyString string

const (
	FIRST  MyString = "1"
	SECOND MyString = "2"
)

// For integer based 'enums' there is another trick to make the implementation cleaner
const (
	NUM1 = iota
	NUM2
	NUM3
)

// The first member of the 'enum' will have the value 0, the next 1 and so on

func main() {

}
