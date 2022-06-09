package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func checkExpressionValidity(input string) (string, error) {
	trimmed := strings.Replace(input, " ", "", -1)

	// check for whitespaces only
	if len(trimmed) == 0 {
		err := fmt.Errorf("bad token nil. %w", errorEmptyInput)
		return "", err
	}

	// check chars
	charsRe := regexp.MustCompile("[A-Za-z]+")
	values := charsRe.FindAllString(input, -1)
	if len(values) != 0 {
		err := fmt.Errorf("chars in input string. %w", errorNotTwoOperands)
		return "", err
	}

	return trimmed, nil
}

func sumStringValues(values []string) (int64, error) {
	var result int64 = 0
	for _, v := range values {
		value, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, err
		}
		result += value
	}
	return result, nil
}

func StringSum(input string) (output string, err error) {
	validated, err := checkExpressionValidity(input)
	if err != nil {
		return "", err
	}

	var result int64
	allValues := regexp.MustCompile("\\d+").FindAllString(validated, -1)
	negativeValues := regexp.MustCompile("(-)\\d+").FindAllString(validated, -1)

	counted, err := sumStringValues(allValues)
	if err != nil {
		return "", err
	}

	result = counted

	counted, err = sumStringValues(negativeValues)
	if err != nil {
		return "", err
	}

	result += counted * 2
	return strconv.FormatInt(result, 10), nil
}
