package main

import (
	"github.com/fatih/color"
	"strings"
	"testing"
)

func calculatorTester(c calculator, x, y, res float64) bool {
	if c.x == x && c.y == y && c.calc(c.x, c.y) == res {
		return true
	}

	return false
}

func TestParse(t *testing.T) {
	type tester struct {
		args []string
		x    float64
		y    float64
		res  float64
	}

	testers := []tester{
		{
			args: []string{"1.0", "+", "99"},
			x:    1,
			y:    99,
			res:  100,
		},
		{
			args: []string{"99.0", "-", "1.0"},
			x:    99,
			y:    1,
			res:  98,
		},
		{
			args: []string{"5", "*", "5"},
			x:    5,
			y:    5,
			res:  25,
		},
		{
			args: []string{"100", "/", "25.0"},
			x:    100,
			y:    25,
			res:  4,
		},
	}

	for _, tester := range testers {
		for _, args := range [][]string{tester.args, {strings.Join(tester.args, "")}} {
			c, err := parse(args)

			if err != nil {
				t.Error(err)
				continue
			}

			if !calculatorTester(c, tester.x, tester.y, tester.res) {
				t.Errorf("calculator does not match expected x=%v y=%v res=%v, got x=%v y=%v res=%v", tester.x, tester.y, tester.res, c.x, c.y, c.calc(c.x, c.y))
			}
		}
	}
}

func TestParseInvalid(t *testing.T) {
	testers := [][]string{
		{},
		{
			"11",
		},
		{
			"a", "a", "a",
		},
		{
			"1", "2", "3",
		},
		{
			"1", "+", "1", "1",
		},
		{
			"2", "^", "2",
		},
	}

	for _, tester := range testers {
		_, err := parse(tester)

		if err == nil {
			t.Errorf("expects expression to be invalid, got valid: %v", tester)
		}
	}
}

func TestFormat(t *testing.T) {
	calcFun := color.New(color.Italic).SprintFunc()
	resFun := color.New(color.Bold, color.Underline).SprintFunc()

	formattedCalculator := func(a, b string) string {
		return calcFun(a) + " " + resFun(b)
	}

	testers := map[string]calculator{
		formattedCalculator("1 + 1 =", "2"): {
			x:        1,
			y:        1,
			operator: "+",
			calc:     operators["+"],
		},
		formattedCalculator("100.12 - 20 =", "80.12"): {
			x:        100.12,
			y:        20,
			operator: "-",
			calc:     operators["-"],
		},
		formattedCalculator("1337 * 42 =", "56154"): {
			x:        1337,
			y:        42,
			operator: "*",
			calc:     operators["*"],
		},
		formattedCalculator("100 / 125 =", "0.8"): {
			x:        100,
			y:        125,
			operator: "/",
			calc:     operators["/"],
		},
	}

	for str, tester := range testers {
		if str != format(tester) {
			t.Errorf("failed asserting that \"%v\" is \"%v\"", format(tester), str)
		}
	}
}
