package main
import "fmt"

func main() {
	var u int64 = 50600099
	var intArray = make([]int, u+1)
	fmt.Println(len(intArray))
	//var j int64 =5


	// Schleife, von 2 bis kleiner zahl b
	for y:=2;y<5;y++{

		intArray[y] =y
	}
 var f int64=10000
 var b = make([]int, f)
 // gibt die Länge des arrays zurück len(b)
 fmt.Println("len exper." ,len(b))


}
