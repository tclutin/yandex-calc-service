package calc

import "testing"

func TestCalc(t *testing.T) {
	calc := New()

	testCaseSuccess := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "addition",
			expression:     "1+1",
			expectedResult: 2,
		},
		{
			name:           "priority with brackets",
			expression:     "(2+2)*2",
			expectedResult: 8,
		},
		{
			name:           "multiplication precedence",
			expression:     "2+2*2",
			expectedResult: 6,
		},
		{
			name:           "division",
			expression:     "1/2",
			expectedResult: 0.5,
		},
		{
			name:           "subtraction",
			expression:     "5-3",
			expectedResult: 2,
		},
		{
			name:           "complex mixed operations",
			expression:     "2+3*(4-1)/3",
			expectedResult: 5,
		},
		{
			name:           "nested brackets",
			expression:     "((1+2)*(3+4))/2",
			expectedResult: 10.5,
		},
		{
			name:           "decimal numbers",
			expression:     "2.5+3.5",
			expectedResult: 6.0,
		},
		{
			name:           "unary minus",
			expression:     "-2+3",
			expectedResult: 1,
		},
		{
			name:           "negative result",
			expression:     "1-2-3",
			expectedResult: -4,
		},
		{
			name:           "zero in expression",
			expression:     "0+3*2",
			expectedResult: 6,
		},
		{
			name:           "division by one",
			expression:     "5/1",
			expectedResult: 5,
		},
		{
			name:           "division by a fraction",
			expression:     "1/0.5",
			expectedResult: 2,
		},
		{
			name:           "multiplication by zero",
			expression:     "10*0",
			expectedResult: 0,
		},
		{
			name:           "addition of zero",
			expression:     "0+0",
			expectedResult: 0,
		},
		{
			name:           "large numbers",
			expression:     "1000000+2000000",
			expectedResult: 3000000,
		},
		{
			name:           "unary plus",
			expression:     "+3-4*(2+2)",
			expectedResult: -13,
		},
	}

	for _, testCase := range testCaseSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := calc.Calc(testCase.expression)
			if err != nil {
				t.Fatalf("successful case %s returns error", testCase.expression)
			}

			if result != testCase.expectedResult {
				t.Fatalf("%f should be equal %f", result, testCase.expectedResult)
			}
		})
	}

	testCasesFail := []struct {
		name        string
		expression  string
		expectedErr error
	}{
		{
			name:       "operator after expression",
			expression: "1+1*",
		},
		{
			name:       "double operator",
			expression: "2+2**2",
		},
		{
			name:       "invalid brackets and operator",
			expression: "((2+2-*(2",
		},
		{
			name:       "empty expression",
			expression: "",
		},
		{
			name:       "unmatched opening brackets",
			expression: "(2+2",
		},
		{
			name:       "unmatched closing brackets",
			expression: "2+2)",
		},
		{
			name:       "operator at end",
			expression: "2+2-",
		},
		{
			name:       "only operator",
			expression: "+",
		},
		{
			name:       "invalid characters",
			expression: "2+2a",
		},
		{
			name:       "multiple invalid characters",
			expression: "2+@3&5",
		},
		{
			name:       "division by zero",
			expression: "5/0",
		},
		{
			name:       "division by zero with brackets",
			expression: "(2+3)/(1-1)",
		},
		{
			name:       "division by zero with complex expression",
			expression: "10/(5-5)+3",
		},
		{
			name:       "whitespace-only",
			expression: "   ",
		},
		{
			name:       "multiple spaces between numbers",
			expression: "2   +   3",
		},
		{
			name:       "missing operand",
			expression: "2+",
		},
		{
			name:       "negative without operand",
			expression: "-",
		},
		{
			name:       "multiple operators together",
			expression: "2++2",
		},
		{
			name:       "nested invalid brackets",
			expression: "((2+2)*(3+))-1",
		},
		{
			name:       "floating point without digits",
			expression: "2. + 3",
		},
		{
			name:       "division by zero in mixed operators",
			expression: "1+2/(3-3)*4",
		},
	}

	for _, testCase := range testCasesFail {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := calc.Calc(testCase.expression)
			if err == nil {
				t.Fatalf("expression %s is invalid but result  %f was obtained", testCase.expression, result)
			}
		})
	}
}
