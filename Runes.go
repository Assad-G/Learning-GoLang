package main

import ( "fmt"
 "reflect")

func main(){
	Runes()
}
func Runes(){
rune1 := 'b'
rune2 := 'g'
rune3 := '\a'
fmt.Printf("Rune1 : %c; Unicode : %U; Type: %s", rune1, rune1, reflect.TypeOf(rune1))
fmt.Printf("\nRune 2: %c; Unicode: %U; Type: %s", rune2,  rune2, reflect.TypeOf(rune2))
fmt.Printf("\nRune 3: Unicode: %U; Type: %s", rune3, reflect.TypeOf(rune3))



rune4 := 'a'
rune5 := '🙂'
rune6 := '\a'

fmt.Printf("rune4: %c Unicode: %U Type: %s\n", rune4, rune4, reflect.TypeOf(rune4))
fmt.Printf("rune5: %c Unicode: %U Type: %s\n", rune5, rune5, reflect.TypeOf(rune5))
fmt.Printf("rune6: %c Unicode: %U Type: %s\n", rune6, rune6, reflect.TypeOf(rune6))
}
