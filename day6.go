package main
/*
Given a string, S, of length N that is indexed from 0 to N-1,
print its even-indexed and odd-indexed characters as 2 space-separated strings on a single line 
(see the Sample below for more detail).

Note:  is considered to be an even index.
*/
import "fmt"
func main() {
    var T int

    fmt.Scan(&T)

    for i := 0; i < T; i++ {
        var str string
        fmt.Scan(&str)
        processAndPrint(str)
    }
}

func processAndPrint(str string){
    length := len(str)
    var oddString string
    var evenString string
    for i := 0; i < length; i ++{
        if i % 2 == 0{
            evenString += string(str[i])
        }else{
            oddString += string(str[i])
        }
    }

    fmt.Printf("%s %s\n",evenString,oddString)
}