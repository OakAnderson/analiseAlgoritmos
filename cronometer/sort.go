package cronometer

import (
	"fmt"
	"time"

	"github.com/OakAnderson/analiseAlgoritmos/random"
)

type sortFunction func([]int)
type sortFucntionWithReturn func([]int) []int
type sortFunctionWithBool func([]int, bool) ([]int, error)

// Sort is an object that make tests of a sort algorithm and rank their results
type Sort struct {
	f            sortFunction
	hasParameter bool
	ArrSize      int
	LastResults  []time.Duration
	LastResult   time.Duration
}

// SetSortFunction is a method that set a new function to the Sort struct
func (s *Sort) SetSortFunction(function sortFunction) {
	s.hasParameter = false
	s.f = function
}

// SetSortFunctionWithReturn set a function that returns the sorted array to the Sort object
func (s *Sort) SetSortFunctionWithReturn(function sortFucntionWithReturn) {
	s.hasParameter = false
	s.f = func(arr []int) {
		function(arr)
	}
}

// SetSortFunctionWithBool set a function that wait for a bool parameter to the Sort object
func (s *Sort) SetSortFunctionWithBool(function sortFunctionWithBool, parameter bool) {
	s.hasParameter = true
	s.f = func(arr []int) {
		function(arr, parameter)
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
		results[i] = s.UnitTest()
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
		sum += s.UnitTest()
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

// UnitTest make a single test for the sort function
func (s *Sort) UnitTest() time.Duration {
	arr := random.Ints(s.ArrSize)
	init := time.Now()
	s.f(arr)
	s.LastResult = time.Since(init)
	return s.LastResult
}
