package main
import "fmt"
func main(){
	Boolens()
}
func Boolens(){
bool1 := "Assadiscool"
bool2 := "assadiscoollle"
bool3 := "Assadiscool"

Result1:= bool1 == bool2
Result2:= bool1 == bool3

fmt.Println(Result1)
fmt.Println(Result2)

fmt.Printf("The type of Result1 is : %T and "+ "The type of Result2 is : %T", Result1,Result2 )


}


