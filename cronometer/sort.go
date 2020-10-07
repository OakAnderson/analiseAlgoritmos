package cronometer

import (
	"fmt"
	"time"

	"github.com/OakAnderson/analiseAlgoritmos/random"
)

// Sort is an object that make tests of a sort algorithm and rank their results
type Sort struct {
	f            func([]int)
	hasParameter bool
	ArrSize      int
	LastResults  []time.Duration
	LastResult   time.Duration
}

// SetFunction is a method that set a new function to the Sort struct
func (s *Sort) SetFunction(function func([]int)) {
	s.hasParameter = false
	s.f = function
}

// SetFunctionWithReturn set a function that returns the sorted array to the Sort object
func (s *Sort) SetFunctionWithReturn(function func([]int)[]int) {
	s.hasParameter = false
	s.f = func(arr []int) {
		function(arr)
	}
}

// SetArrSize set the size of the array to make tests
func (s *Sort) SetArrSize(n int) {
	s.ArrSize = n
}

// MultipleTests make multiple tests for the sort algorithm defined
func (s *Sort) MultipleTests(tests int) ([]time.Duration, error) {
	if s.ArrSize == 0 {
		s.SetArrSize(100)
	}
	if s.f == nil {
		return nil, fmt.Errorf("Sort function is undefined")
	}

	results := make([]time.Duration, tests)
	for i := 0; i < tests; i++ {
		results[i] = s.SingleTest()
	}
	s.LastResults = results

	return s.LastResults, nil
}

// MultipleTestsMean return the mean of time duration that the sort algorithm takes to sort an array
func (s *Sort) MultipleTestsMean(tests int) (time.Duration, error) {
	if s.ArrSize == 0 {
		s.SetArrSize(100)
	}
	if s.f == nil {
		return time.Duration(0), fmt.Errorf("Sort function is undefined")
	}

	var sum time.Duration
	for i := 0; i < tests; i++ {
		sum += s.SingleTest()
	}
	sum = sum / time.Duration(tests)

	return sum, nil
}

// Mean return the mean of the results
func (s *Sort) Mean() time.Duration {
	var sum time.Duration
	if s.LastResults == nil || len(s.LastResults) == 1 {
		return s.LastResult
	}
	for _, v := range s.LastResults {
		sum += v
	}
	return sum / time.Duration(len(s.LastResults))
}

// SingleTest make a single test for the sort function
func (s *Sort) SingleTest() time.Duration {
	arr := random.Ints(s.ArrSize)
	init := time.Now()
	s.f(arr)
	s.LastResult = time.Since(init)
	return s.LastResult
}
