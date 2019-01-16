# Go Training

Training materials for beginners to Go. Each package represents a general topic area and each source
file is runnable (has it's own main function)

## Important

This repo should not be used as an example of how to organise your project or structure your
source files. For a better example of organisation and structure, clone the [Granitic source
repository](https://github.com/graniticio/granitic)

The following is the suggested reading order:

## Foundations
  1. [Getting output from your code](output/output.go)
 
## Types, variables and control structures

  1. [Declaring and intialising variables](variablestypes/variables.go)
  1. [Basic types](variablestypes/builtin.go)
  1. [Constants](variablestypes/builtin.go)
  1. [if-else statements](controlstructures/ifelse.go)
  1. [for statements](controlstructures/forloop.go)
  1. [switch statements](controlstructures/switch.go)
  1. [Functions](functions/basics.go)
 
## Data structures
 
  1. [Slices and arrays](structures/slices.go)
  1. [Maps](structures/maps.go)
  1. [Pointers](structures/pointers.go)
  1. [Structs](structures/structs.go)
  1. [Methods and receivers](structures/methods.go)
  1. [Interfaces](variablestypes/interfaces.go)
  1. [Tags](structures/tags.go)
  
## Error handling
 
  1. [Errors](errorhandling/errors.go)
  1. [Panic](errorhandling/panic.go)
  
## Testing

  1. [Unit tests](unittests/code_test.go)
  
## Advanced functions
 
  1. [Variadic functions](functions/variadic.go)
  1. [Defer](functions/defer.go)
  1. [Functions as types](functions/types.go)
  1. [Closures and anonymous functions](functions/closures.go)
  
## Concurrency
   
  1. [goroutines](concurrency/goroutines.go)
  1. [Channels](concurrency/channels.go)
  1. [Mutex and atomic](concurrency/mutex.go)
  1. [Context](concurrency/context.go)      
  
## Essential packages

These are some of the packages you'll find yourself using most frequently.


  1. [bufio](essential/bufio.go) ([Godoc](https://golang.org/pkg/bufio))
  1. [encoding/json](essential/json.go) ([Godoc](https://golang.org/pkg/encoding/json))
  1. [net/http](essential/http.go) ([Godoc](https://golang.org/pkg/net/http))  
  1. [regexp](essential/regexp.go) ([Godoc](https://golang.org/pkg/regexp))  
  1. [strconv](essential/strconv.go) ([Godoc](https://golang.org/pkg/strconv))
  1. [strings](essential/strings.go) ([Godoc](https://golang.org/pkg/strings))
  1. [time](essential/time.go) ([Godoc](https://golang.org/pkg/time))  

