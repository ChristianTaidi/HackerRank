package main


import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
  
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)
    // Declare second integer, double, and String variables.
    var j int 
    scanner.Scan()
    j,_ = strconv.Atoi(string(scanner.Bytes()))
    var f float64
    scanner.Scan()
    f,_ = strconv.ParseFloat(string(scanner.Bytes()),64)
    var s2 string    
    scanner.Scan()
    s2 = scanner.Text()
    fmt.Printf("%v\n",i+uint64(j))
    fmt.Printf("%.1f\n",d+f)
    fmt.Printf("%s",s+s2)
}