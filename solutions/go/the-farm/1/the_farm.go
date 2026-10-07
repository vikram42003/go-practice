package thefarm

import (
    "errors"
    "fmt"
)

var ErrInvalidCowCount = errors.New("invalid number of cows")

type InvalidCowsError struct {
    cows int
    message string
}

func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, cowCount int) (float64, error) {
    amount, err := fc.FodderAmount(cowCount)
    if err != nil {
        return 0.0, err
    }
    factor, err := fc.FatteningFactor()
    if err != nil {
        return 0.0, err
    }
    return (amount * factor) / float64(cowCount), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cowCount int) (float64, error) {
    if cowCount > 0 {
        return DivideFood(fc, cowCount)
    } else {
        return 0.0, ErrInvalidCowCount
    }
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cowCount int) error {
    if cowCount < 0 {
        return &InvalidCowsError{ cowCount, "there are no negative cows" }
    } else if cowCount == 0 {
        return &InvalidCowsError{ cowCount, "no cows don't need food" }
    } else {
        return nil
    }
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
