package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{"Matthew","Sarah","Augustus","Heidi",
		"Emilie","peter","Ginana","Adriano","Aaron","Elizabeth"}
	distribution = make(map[string]int,len(users))
)
func ContainStr(name string){
	strMap := map[string]int{"eE":1,"iI":2,"oO":3,"uU":4}
	for k,v :=range strMap{
		if strings.ContainsAny(name,k){
			distribution[name] +=v
			coins -=v
		}
	}

}


func dispatchCoin() int {
	for _,v := range users{
		ContainStr(v)
	}
	return coins
}

func main() {
	//studentMap := make(map[string]map[string]int,10)
	////初始化内层的map
	//studentMap["豪杰"] = make(map[string]int,3)
	//studentMap["豪杰"]["id"] = 1
	//studentMap["豪杰"]["age"] = 18
	//studentMap["豪杰"]["score"] = 90
	//
	//studentMap["哪吒"] = make(map[string]int,3)
	//studentMap["哪吒"]["id"] = 1
	//studentMap["哪吒"]["age"] = 28
	//studentMap["哪吒"]["score"] = 100
	//
	//for k,v := range studentMap{
	//	fmt.Println(k)
	//	for k2, v2 :=range v{
	//		fmt.Println(k2,v2)
	//	}
	//}

	//------------------------------------------------------
	left := dispatchCoin()
	fmt.Println("剩下:",left)
	fmt.Println(distribution)
}
