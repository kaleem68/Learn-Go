package main
import (
    "fmt"
    "bufio"
    "os"
    "log"
)
var pl = fmt.Println

func main(){
    variables()
}

func variables(){
    var vName string = "Kaleem"
    var v1, v2 = 1.2, 3.4
    var v3 = "hello"
    v4 := 2.4
    v4 = 5.4
    pl(vName, v1, v2, v3, v4)
}

func readInputAndDisplayOutput(){
    pl("What is your name ?")  
    reader := bufio.NewReader(os.Stdin)
    name, err := reader.ReadString('\n')

    if err == nil {
        pl("Hello",name)
    } else {
        log.Fatal(err)
    }
}