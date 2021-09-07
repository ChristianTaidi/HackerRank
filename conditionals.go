package main
/*
	Given an integer, n, perform the following conditional actions:	
	If n is odd, print Weird
	If n is even and in the inclusive range of 2 to 5, print Not Weird
	If n is even and in the inclusive range of 6 to 20, print Weird
	If n is even and greater than 20, print Not Weird
	Complete the stub code provided in your editor to print whether or not n is weird.

	1<= n <= 100
*/

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func checkWeird(number int32){
	if number%2 == 1 {
		fmt.Print("Weird")
		return
	}
	if number >= 2 && number <= 5{
		fmt.Print("Not Weird")
	}
	if number >= 6 && number <= 20 {
		fmt.Print("Weird")
	}
	if number > 20 {
		fmt.Print("Not Weird")
	}
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    N := int32(NTemp)
	checkWeird(N)
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
