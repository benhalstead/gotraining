package sample

func NewContactDetails(wll, wm, pm string) ContactDetails {

	return ContactDetails{
		WorkMobile:     wm,
		WorkLandline:   wll,
		personalMobile: pm,
	}

}

//An exported struct (because first char is a capital)
type ContactDetails struct {
	WorkLandline   string // This member is public (because first char is a capital)
	WorkMobile     string
	personalMobile string // This member is private (but code in the same package can access and modify it)
}

// An unexported struct
type bankDetails struct {
	// Members can still be exported or unexported - this is subtle, but relevant for things like JSON encoding and decoding where
	// structs are created using reflection (way out of scope for these tutorials!)
	pin           string
	AccountNumber string
}

// A struct with nested struct definitions - clever and neat, but likely to cause problems with initialisation
type PersonNested struct {
	Name         string
	BirthDetails struct {
		Date string
		City string
	}
}

// A struct with members that are also structs
type Person struct {
	Name           string
	ContactDetails ContactDetails //It's possible (and common) to use the name of the type as the name of the member
}

// If you have a complex data structure and might want to pass parts of it around independently, you will want to define
// members that are structs as pointers
type PersonHolder struct {
	Name           string
	ContactDetails *ContactDetails
}
