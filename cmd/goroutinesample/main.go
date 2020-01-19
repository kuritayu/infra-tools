package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	log.Println("started.")

	// クロージャの配列
	funcs := []func(){
		func() {
			// 1秒かかる
			log.Println("sleep1 started.")
			time.Sleep(1 * time.Second)
			log.Println("sleep1 ended.")
		},

		func() {
			// 2秒かかる
			log.Println("sleep2 started.")
			time.Sleep(2 * time.Second)
			log.Println("sleep2 ended.")
		},

		func() {
			// 3秒かかる
			log.Println("sleep3 started.")
			time.Sleep(3 * time.Second)
			log.Println("sleep3 ended.")
		},
	}

	var wg sync.WaitGroup
	for _, sleep := range funcs {
		wg.Add(1)

		go func(function func()) {
			defer wg.Done()
			function()
		}(sleep)

	}

	// 待ち
	wg.Wait()

	log.Println("all process ended.")
}
