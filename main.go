package main

import (
	"fmt"
	"strconv"
	"os"
	"log"
	"strings"
	"regexp"
	"github.com/fatih/color"
)

type calculate func(float64, float64) float64

type calculator struct {
	x        float64
	y        float64
	operator string
	calc     calculate
}

var operators = map[string]calculate{
	"+": func(x float64, y float64) float64 { return x + y },
	"-": func(x float64, y float64) float64 { return x - y },
	"*": func(x float64, y float64) float64 { return x * y },
	"/": func(x float64, y float64) float64 { return x / y },
}

func main() {
	c, err := parse(os.Args[1:])

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(format(c))
}

func parse(args []string) (c calculator, err error) {
	if len(args) < 1 {
		err = fmt.Errorf("you must provide at least 1 argument, got %v", len(args))
		return
	}
	if len(args) > 3 {
		err = fmt.Errorf("you can not provide more than 3 arguments, got %v", len(args))
		return
	}

	q := strings.Join(args, "")
	nr := "\\d+(?:\\.\\d+)?"
	re := regexp.MustCompile(fmt.Sprintf("^(%s)(\\%s)(%s)$", nr, strings.Join(keysOfStringCalculateMap(operators), "|\\"), nr))

	matches := re.FindStringSubmatch(q)

	if len(matches) == 0 {
		err = fmt.Errorf("invalid calculation provided: \"%v\"", q)
		return
	}

	x, _ := strconv.ParseFloat(matches[1], 64)
	y, _ := strconv.ParseFloat(matches[3], 64)
	calc, _ := operators[matches[2]]

	c.x = x
	c.y = y
	c.calc = calc
	c.operator = matches[2]

	return
}

func format(c calculator) string {
	calc := color.New(color.Italic).Sprintf("%v %s %v =", c.x, c.operator, c.y)
	res := color.New(color.Bold, color.Underline).Sprint(c.calc(c.x, c.y))

	return fmt.Sprintf("%v %v", calc, res)
}

func keysOfStringCalculateMap(slice map[string]calculate) []string {
	r := make([]string, 0, len(slice))

	for k := range slice {
		r = append(r, k)
	}

	return r
}
