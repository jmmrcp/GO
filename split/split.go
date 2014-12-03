package main

import "fmt"
import "strings"

func main() {
	str := "the quick brown fox jumps over the lazy dog"
	s1 := strings.Fields(str)
	fmt.Println(str)
	fmt.Println(s1)
	fmt.Println(s1[3])
	LetraNif := "T R W A G M Y F P D X B N J Z S Q V H L C K E"
	Nif := 22311845
	Nif = Nif & 23
	fmt.Println(Nif)
	fmt.Println(LetraNif[1])
}