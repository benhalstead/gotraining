package unittests

import "testing"

/*
	Files containing unit tests must have the suffix _test.go

	By convention, the files have the same name as the file containing the code under test so this package
	contains the files

	code.go
	code_test.go

	Any code in files with the _test.go suffix will not be included in your final executable (likewise code that
	is only imported and used by your tests)


	Go includes a testing and benchmarking suite in the testing package:

	https://golang.org/pkg/testing/

	But it is very raw. The vast majority of projects use a third-party unit testing framework which builds on Go's
	internal framework.

	This article is a decent overview of the currently popular test frameworks

	https://bmuschko.com/blog/go-testing-frameworks/

	This file shows some of the features available in the basic testing package

*/

//Individual tests must have the method name func TestXxx(*testing.T)
//The testing package's support function are provided in the *testing.T object passed into the test
func TestSquare(t *testing.T) {

	m := Math{}

	result := m.Square(4)

	//There are no built-in assert methods!

	if result != 16 {

		//t.Fail() marks this test failed without a message and continues execution
		//t.FailNow() marks this test failed and exits this test immediately
		//t.Fatalf calls FailNow _and_ logs an error message

		t.Errorf("Expected 16 but found %d", result) //Calls Fail() and logs an error message
	}

}

func TestBrokenSquare(t *testing.T) {

	m := Math{}

	result := m.BrokenSquare(1)
	expected := 1

	if result != expected {
		t.Fatalf("Expected %d, found %d", expected, result)
	}

}
