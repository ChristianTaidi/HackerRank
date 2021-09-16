package main

/*
Given a base-10 integer, n, convert it to binary (base-2). 
Then find and print the base-10 integer denoting the maximum number of consecutive 1's in n's binary representation. 
When working with different bases, it is common to show the base as a subscript.

Example: The binary representation of 125 is 1111101. In base 2, there are 5 and 1 consecutive ones in two groups. Print the maximum, 5.


*/

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int(nTemp)
    result := maxConsecutive1s(n)
    fmt.Printf("%d\n",result)
}

func maxConsecutive1s(bin int) int {
    count := 0
    for bin!=0{
        bin = (bin & ( bin <<1 ))
        count ++
    }
    return count
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