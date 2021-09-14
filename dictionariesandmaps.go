package main
/*
Given n names and phone numbers, assemble a phone book that maps friends' names to their respective phone numbers. 
You will then be given an unknown number of names to query your phone book for. For each name queried, 
print the associated entry from your phone book on a new line in the form name=phoneNumber; 
if an entry for name is not found, print Not found instead.
*/
import "fmt"
import "strings"
import "bufio"
import "os"
import "io"

func main(){
    var n int
    phoneBook := make(map[string]string)
    fmt.Scan(&n)
    reader := bufio.NewReaderSize(os.Stdin,16*1024*1024)
    for i:= 0; i<n; i++{
        var name string
        var phone string
        arrTmp := strings.Split(strings.TrimSpace(readLine(reader))," ")
        name = arrTmp[0]
        phone = arrTmp[1]
        phoneBook[name] = phone
    }
    // queries
    for{
        
        name := readLine(reader)
        if name == ""{
            break
        }
        number,ok := phoneBook[name]
        if !ok{
            fmt.Printf("Not found\n")
            continue
        }
        fmt.Printf("%s=%s\n",name,number)
    }

}

func readLine( reader *bufio.Reader) string{
    str, _, err := reader.ReadLine()
    if err == io.EOF{
        return ""
    }
    return strings.TrimRight(string(str),"\r\n")
}