package main

import (
	"fmt"
)

func main() {
	// expression := "|(&(t,f,t),!(t))" // Answer: false
	// expression := "|(f,t)" // Answer: true
	// expression := "!(f)" // Answer: true
	// expression := "&(t,t,t)" // Answer: true
	expression := "|(f,&(t,t))" // Answer: true
	fmt.Println(parseBoolExpr(expression))
}

func findFirstSubexpIndices(exp string) (int, int) {
	var lastOpenIndex int // last open BEFORE the first close
	var firstCloseIndex int

	for i, c := range exp {
		if c == '(' {
			lastOpenIndex = i
		} else if c == ')' {
			firstCloseIndex = i
			break
		}
	}

	return lastOpenIndex, firstCloseIndex
}

func parseBoolExpr(expression string) bool {
	exp := expression

	for len(exp) > 1 {
		// fmt.Println(exp)

		lastOpenIndex, firstCloseIndex := findFirstSubexpIndices(exp)

		// Separate logical operator from the expression to be evaluated
		subexp := exp[lastOpenIndex+1 : firstCloseIndex]
		logical := exp[lastOpenIndex-1]

		// Count the number of 't', 'f', and ',' in the expression to be
		// evaluated
		numTrue := 0
		numFalse := 0
		numCommas := 0
		for i := range subexp {
			switch subexp[i] {
			case 't':
				numTrue++
			case 'f':
				numFalse++
			case ',':
				numCommas++
			}
		}

		// Evaluate the expression
		var eval bool
		numC := len(subexp) - numCommas
		if logical == '&' && numTrue == numC ||
			logical == '|' && numFalse != numC ||
			logical == '!' && numFalse == 1 {
			eval = true
		}

		evalStr := "f"
		if eval {
			evalStr = "t"
		}

		// Update exp with eval
		switch {
		case lastOpenIndex-2 >= 0:
			exp = exp[:lastOpenIndex-1] + evalStr + exp[firstCloseIndex+1:]
		case firstCloseIndex != len(exp)-1:
			exp = evalStr + exp[firstCloseIndex+1:]
		default:
			exp = evalStr
		}
	}

	if exp[0] == 't' {
		return true
	}
	return false
}
