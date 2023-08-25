package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Arrs struct {
	Name string
	Id   int
}
type Arr struct {
	Name string
	Id   int
	Data []Arrs
}

func main() {
	_, err := net.Dial("tcp", "localhost:80")
	if err == nil {
		fmt.Println("Connection successful")
	} else {
		fmt.Println(err)
	}

	var Olddatas = Arrs{Name: "网吧3", Id: 3}
	a := []Arr{
		Arr{Name: "adas", Id: 0, Data: []Arrs{
			Olddatas,
		}},
	}
	a = append(a, Arr{Name: "网吧dd", Id: 0})

	for index, value := range a {
		a[index].Data = []Arrs{
			Arrs{Name: "adas", Id: 20 + index},
		}
		fmt.Println(index, value.Data)
		for key, vs := range value.Data {
			a[index].Data[key].Name = "6666"
			fmt.Println("二级打印")
			fmt.Println(key, vs)
			// vs.Data = []Arrs{
			// 	Olddatas,
			// }
			// fmt.Println(vs.Data)
		}
	}
	fmt.Println(a)
	bytes, _ := json.Marshal(a)
	stringData := string(bytes)
	fmt.Println(stringData)
	// fmt.Println(a[0])
	// {

	// 	{"user1", 10,
	// 	[]Arr{
	// 		"user2", 20,}
	// },
	// 	// {"user2", 20, []Arr{"user1", 10}},
	// }
	// for key, v := range a {
	// 	fmt.Println(v, key)
	// }
	// userInfo := map[string]string{
	// 	"username": "pprof.cn",
	// 	"password": "123456",
	// }
	//
}
