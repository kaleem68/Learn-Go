package main
import (
    "fmt"
    "bufio"
    "os"
    "log"
	"reflect"
)
var pl = fmt.Println

func main(){
    dataTypes()
}

func dataTypes(){
    //int, float64, bool, string, rune
    //Default type 0, 0.0, false, ""
    pl(reflect.TypeOf(35))
    pl(reflect.TypeOf(35.43))
    pl(reflect.TypeOf(false))
    pl(reflect.TypeOf("35"))
    pl(reflect.TypeOf('🐵'))
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