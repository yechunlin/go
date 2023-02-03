package main

import (
	"fmt"
	"script/lib"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		db := lib.GetInstance()
		sql := "select id,nickname from `users` where id = 33"
		var id int
		var nickname string
		err := db.QueryRow(sql).Scan(&id, &nickname)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second * 2)
		fmt.Println(id, nickname)
	}()

	go func() {
		wg.Done()
		db := lib.GetInstance()
		sql := "select id,nickname from `users` where id = 37"
		var id int
		var nickname string
		err := db.QueryRow(sql).Scan(&id, &nickname)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
		fmt.Println(id, nickname)
	}()
	wg.Wait()
}
