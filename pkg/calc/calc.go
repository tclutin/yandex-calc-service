package calc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrValidationError = errors.New("validation error")
	ErrDivisionByZero  = errors.New("division by zero")
)

type Calc struct {
}

func NewCalc() *Calc {
	return &Calc{}
}

func (c *Calc) precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		return 0
	}
}

func (c *Calc) applyOperation(a, b float64, op rune) (float64, error) {
	switch op {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		if b != 0 {
			return a / b, nil
		}
		return 0, ErrDivisionByZero
	}
	return 0, nil
}

func (c *Calc) isOperator(op rune) bool {
	return op == '+' || op == '-' || op == '*' || op == '/'
}

func (c *Calc) isDigit(char string) bool {
	_, err := strconv.ParseFloat(char, 64)
	return err == nil
}

func (c *Calc) transformToPostfix(expression string) string {
	var result strings.Builder
	var operators []rune
	var wasOperator = true

	for i := 0; i < len(expression); i++ {
		char := rune(expression[i])

		if unicode.IsDigit(char) || char == '.' {
			result.WriteRune(char)
			wasOperator = false
			continue
		}

		if char == '+' && wasOperator {
			continue
		}

		if char == '-' && wasOperator {
			result.WriteString("-")
			wasOperator = true
			continue
		}

		if char == '(' {
			operators = append(operators, char)
			wasOperator = true
			continue
		}

		if char == ')' {
			for len(operators) > 0 && operators[len(operators)-1] != '(' {
				result.WriteRune(' ')
				result.WriteRune(operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}

			if len(operators) > 0 && operators[len(operators)-1] == '(' {
				operators = operators[:len(operators)-1]
			}
			wasOperator = false
		}

		if c.isOperator(char) {
			result.WriteRune(' ')
			for len(operators) > 0 && c.precedence(operators[len(operators)-1]) >= c.precedence(char) {
				result.WriteRune(operators[len(operators)-1])
				result.WriteRune(' ')
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, char)
			wasOperator = true
		}
	}

	for len(operators) > 0 {
		result.WriteRune(' ')
		result.WriteRune(operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return result.String()
}

func (c *Calc) evaluatePostfix(expression string) (float64, error) {
	var elements = strings.Split(expression, " ")
	var values []float64

	for _, element := range elements {
		if c.isDigit(element) {
			value, err := strconv.ParseFloat(element, 64)
			if err != nil {
				return 0, fmt.Errorf("%w: %v", ErrValidationError, err)
			}
			values = append(values, value)
			continue
		}

		if c.isOperator(rune(element[0])) {
			if len(values) < 2 {
				return 0, fmt.Errorf("%w: not enough operands - '%s'", ErrValidationError, expression)
			}

			b := values[len(values)-1]
			a := values[len(values)-2]
			values = values[:len(values)-2]

			result, err := c.applyOperation(a, b, rune(element[0]))
			if err != nil {
				return 0, err
			}
			values = append(values, result)
		}

	}

	if len(values) != 1 {
		return 0, fmt.Errorf("%w: invalid expression - '%s'", ErrValidationError, expression)
	}

	return values[0], nil
}

func (c *Calc) Calc(expression string) (float64, error) {
	if err := c.validate(expression); err != nil {
		return 0, err
	}
	return c.evaluatePostfix(c.transformToPostfix(expression))
}

func (c *Calc) validate(expression string) error {
	var brackets []rune
	var wasOperator = false
	var wasDot = false
	var wasDigit = false

	expression = strings.TrimSpace(expression)

	if len(expression) == 0 {
		return fmt.Errorf("%w: empty expression", ErrValidationError)
	}

	for i := 0; i < len(expression); i++ {
		char := rune(expression[i])

		if !unicode.IsDigit(char) && !c.isOperator(char) && char != '(' && char != ')' && char != '.' {
			return fmt.Errorf("%w: invalid char - '%s'", ErrValidationError, string(char))
		}

		if char == '(' {
			brackets = append(brackets, char)
			wasOperator = true
			continue
		}

		if char == ')' {
			if len(brackets) == 0 {
				return fmt.Errorf("%w: unmatched brackets", ErrValidationError)
			}

			if wasOperator {
				return fmt.Errorf("%w: operator before closing bracket", ErrValidationError)
			}

			brackets = brackets[:len(brackets)-1]
			wasOperator = false
			continue
		}

		if c.isOperator(char) {
			if wasOperator {
				return fmt.Errorf("%w: sequence operators - '%s'", ErrValidationError, string(char))
			}
			wasOperator = true
			wasDigit = false
			wasDot = false
			continue
		}

		if char == '.' {
			if wasDot {
				return fmt.Errorf("%w: sequence dots - '%s'", ErrValidationError, string(char))
			}

			if !wasDigit {
				return fmt.Errorf("%w: missing digit before dot", ErrValidationError)
			}
			wasDot = true
			wasOperator = false
			continue
		}

		if unicode.IsDigit(char) {
			wasDigit = true
			wasOperator = false
			wasDot = false
			continue
		}
	}

	if len(brackets) > 0 {
		return fmt.Errorf("%w: unmatched brackets", ErrValidationError)
	}

	if wasOperator {
		return fmt.Errorf("%w: last character was operator", ErrValidationError)
	}

	return nil
}
