package main

import (
	"fmt"
	"time"

	MTK "github.com/kryast/lumoshive-golang-day-11a"
)

func main() {
	ch := make(chan []MTK.Data)
	JumlahData := 100
	go MTK.GetId(JumlahData, ch)

	for i := 1; i <= JumlahData; i++ {
		data := <-ch
		fmt.Println(data)
	}
	time.Sleep(1 * time.Second)

}
