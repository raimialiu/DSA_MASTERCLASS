package assomptotic_notations

import (
	"fmt"
	"reflect"
	"time"
)

type RuntimeComplexity struct{}

func New() *RuntimeComplexity {
	return &RuntimeComplexity{}
}

// HowLong /**
/**
This function accept a function and variable args len as parameter and also an optional v (verbose) params to log more detailed information to the console.
Measure how long it take the function to execute
This function does not take into account async functions (i.e. functions that use channels to communicate)



REQUIREMENT: ->
1)	function as parameter (denotes what function to execute)
2) 	variable length of argument, that can be passed to the function. The function in 1 above can accept parameter  or not
4)  optional verbose flag to denote if more details should be logged to the console
3)	return type can be void or has a particular return type

EXPECTED OUTCOME
1)	int32 to denote how long (how many seconds it took before the function complete)
2)	console log output to show more detailed information. This will be controlled by passing verbose as argument to the function
*/
func (r *RuntimeComplexity) HowLong(f interface{}, v bool, args ...interface{}) int {
	startTime := time.Now()

	value := reflect.ValueOf(f)

	var valueIn = make([]reflect.Value, len(args))

	for i := 0; i < len(args); i++ {
		valueIn[i] = reflect.ValueOf(args[i])
	}

	_ = value.Call(valueIn)

	endTime := startTime.Sub(time.Now())

	if v {
		fmt.Printf("The function took %d seconds to execute.\n", endTime.Seconds())
	}

	return (int)(endTime.Seconds())
}
