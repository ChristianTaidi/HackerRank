package main

/*
Given a 6x6 2D Array, A:

1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0
We define an hourglass in A to be a subset of values with indices falling in this pattern in A's graphical representation:

a b c
  d
e f g
There are 16 hourglasses in A, and an hourglass sum is the sum of an hourglass' values.

Calculate the hourglass sum for every hourglass in A, then print the maximum hourglass sum.
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

  var arr [][]int32
  for i := 0; i < 6; i++ {
      arrRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

      var arrRow []int32
      for _, arrRowItem := range arrRowTemp {
          arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
          checkError(err)
          arrItem := int32(arrItemTemp)
          arrRow = append(arrRow, arrItem)
      }

      if len(arrRow) != 6 {
          panic("Bad input")
      }

      arr = append(arr, arrRow)
  }
}

func computeHourglass(array [][]int32){
  maxHourglass := math.MinInt32
  for i:= 0 i < len(array) - 2 ; i++{
    for j := 0 j < len(array[i]) - 2; j++{
      hourglassTop := array[i][j] + array[i][j+1] + array[i][j+2]
      hourglassMid := array[i+1][j+1]
      hourglassBot := array[i+2][j] + array[i+2][j+1] + array[i+2][j+2]
      total := hourglassTop + hourglassMid + hourglassBot
      if total > maxHourglass{
        maxHourglass = total
      }

    }
  }
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