package main
import "fmt"
import "bufio"
import "os"

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("Hello, World.\n")
    text,_ := reader.ReadString('\n')
    fmt.Printf("%s\n",text)
}