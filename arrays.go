package main
/*
Given an array, A, of N integers, print A s elements in reverse order as a single line of space-separated numbers.
*/

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



func main(){
    reader := bufio.NewReaderSize(os.Stdin,16*1024*1024)
    n,err := strconv.ParseInt(strings.TrimSpace(readLine(reader)),10,64)
    //Read string as array of chars
    arrTmp := strings.Split(strings.TrimSpace(readLine(reader))," ")
    var arr []int
    checkError(err)
    //Convert string to array of ints
    for i:=0 ; i<int(n); i++{
        arrItem, err := strconv.ParseInt(arrTmp[i],10,64)
        checkError(err)
        arr = append(arr,int(arrItem))
    }
    reverse(int(n), arr)
}

func reverse(n int, arr []int ){
    for i:= n-1; i>=0; i--{
        fmt.Printf("%d ",arr[i])
    }
    fmt.Printf("\n")
}

func readLine( reader *bufio.Reader) string{
    str, _, err := reader.ReadLine()
    if err == io.EOF{
        return ""
    }
    return strings.TrimRight(string(str),"\r\n")
}
func checkError(err error){
    if err != nil{
        panic(err)
    }
}

