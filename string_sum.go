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

func StringSum(input string) (output string, err error) {
	if input == "" {
		err := fmt.Errorf("bad token nil. %w", errorEmptyInput)
		fmt.Println(err.Error())
		return "", err
	}
	input = strings.ReplaceAll(input, " ", "")
	re := regexp.MustCompile(`[\+\-]*[0-9]+`)
	n := re.FindAllString(input, -1)
	res := regexp.MustCompile(`[\^+\^-]*[^0-9]+`)
	fmt.Println(res.FindAllString(input, -1))
	n0 := res.FindAllString(input, -1)
	res0 := regexp.MustCompile(`[\+\-]+`)
	fmt.Println(res0.FindAllString(input, -1))
	n1 := res0.FindAllString(input, -1)
	res1 := regexp.MustCompile(`[a-zA-Zа-яА-Я]*[0-9]+[a-zA-Zа-яА-Я]+`)
	fmt.Println("n2=", res1.FindAllString(input, -1))
	n2 := res1.FindAllString(input, -1)
	if len(n) == 0 && len(n0) > 0 {
		err := fmt.Errorf("\n Ошибка пустое значение: %w", errorEmptyInput)
		fmt.Println(err.Error())
		return "", err
	}
	if len(n) == 1 {
		err := fmt.Errorf("\n Ошибка одно значение: %w", errorNotTwoOperands)
		if err != nil {
			fmt.Println(err.Error())
			return "", err
		}

	}
	if len(n) > 2 {
		err := fmt.Errorf("\n Ошибка введено больше двух значений: %w", errorNotTwoOperands)
		fmt.Println(err.Error())
		return "", err
	}
	if len(n1) > 2 {
		err := fmt.Errorf("\n Ошибка введено больше двух значений: %w", errorNotTwoOperands)
		fmt.Println(err.Error())
		return "", err
	}
	if len(n2) > 0 {
		_, e := strconv.Atoi(n2[0])
		err := fmt.Errorf("bad token. %w", e)
		if err != nil {
			fmt.Println(err.Error())
			return "", err
		}
	}
	var y int = 0
	for _, r := range n {
		y0, err := strconv.Atoi(r)
		if err != nil {
			err = fmt.Errorf("%w", err)
			fmt.Println(err.Error())
			return "", err
		}
		y = y + y0
	}
	output = strconv.Itoa(y)
	return output, nil

}
