package main
/*
	Given the meal price (base cost of a meal), 
	tip percent (the percentage of the meal price being added as tip), 
	and tax percent (the percentage of the meal price being added as tax) for a meal, 
	find and print the meal's total cost. Round the result to the nearest integer.
*/

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
	"math"
)

func solve(meal_cost float64, tip_percent int32, tax_percent int32) {
    totalCost := ((meal_cost * tip_percent)/100) + ((meal_cost * tax_percent)/100) + meal_cost
	fmt.Print(math.Round(totalCost))
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    meal_cost, err := strconv.ParseFloat(strings.TrimSpace(readLine(reader)), 64)
    checkError(err)

    tip_percentTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    tip_percent := int32(tip_percentTemp)

    tax_percentTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    tax_percent := int32(tax_percentTemp)

    solve(meal_cost, tip_percent, tax_percent)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}