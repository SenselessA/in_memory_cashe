package main

import (
	"fmt"
	"log"
	"time"

	"github.com/SenselessA/in-memory-cache/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42, time.Second*5)
	userId, err := cache.Get("userId")
	if err != nil { // err == nil
		log.Fatal(err)
	}
	fmt.Println(userId) // Output: 42

	time.Sleep(time.Second * 6) // прошло 5 секунд

	userId, err = cache.Get("userId")
	if err != nil { // err != nil
		log.Fatal(err) // сработает этот код
	}

	fmt.Println(userId) // не сработает
}
