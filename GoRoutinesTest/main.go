// package main

// import (
// 	"fmt"
// 	"time"
// )

// func printMessage(message string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(message)
// 		time.Sleep(2 * time.Second)
// 	}
// }

// func main() {
// 	go printMessage("Hello from goroutines")
// 	printMessage("Hello from main")

// 	fmt.Println("Main function")
// 	time.Sleep(11 * time.Second)
// }

// package main

// import "fmt"

// func worker(ch chan string) {
// 	ch <- "Work Done"
// }

// func main() {
// 	ch := make(chan string)

// 	go worker(ch)
// 	message := <-ch
// 	fmt.Println(message)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// // --- fake API calls (simulate network requests) ---
// func getWeather(ch chan string) {
// 	time.Sleep(2 * time.Second)
// 	ch <- "Weather: Sunny 30°C"
// }

// func getNews(ch chan string) {
// 	time.Sleep(1 * time.Second)
// 	ch <- "News: Go language concurrency is amazing!"
// }

// func getStock(ch chan string) {
// 	time.Sleep(4 * time.Second)
// 	ch <- "Stock: GOOG +2.3%"
// }

// // --- main ---
// func main() {
// 	start := time.Now()

// 	// one channel to collect results
// 	ch := make(chan string)

// 	// run all three API calls concurrently
// 	go getWeather(ch)
// 	go getNews(ch)
// 	go getStock(ch)

// 	// receive 3 results (in any order)
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(<-ch)
// 	}

// 	fmt.Println("All services done in:", time.Since(start))
// }

// go rountines eg without channel
// package main

// import (
// 	"fmt"
// 	"time"
// )

// var count = 0 // shared variable

// func increment() {
// 	for i := 0; i < 100000; i++ {
// 		count++ // many goroutines writing this at the same time
// 	}
// }

// func main() {
// 	for i := 0; i < 5; i++ { // 5 goroutines
// 		go increment()
// 	}

// 	time.Sleep(1 * time.Second) // wait for them to finish
// 	fmt.Println("Final count:", count)
// }

package main

import (
	"fmt"
)

func increment(ch chan int) {
	for i := 0; i < 1000000; i++ {
		ch <- 1
	}
}

func main() {

	ch := make(chan int)
	count := 0
	numGoroutines := 5
	numPerGoroutine := 1000000
	expected := numGoroutines * numPerGoroutine

	for i := 0; i < numGoroutines; i++ {
		go increment(ch)
	}

	for i := 0; i < expected; i++ {
		count += <-ch
	}

	fmt.Println("final count: ", count)
}
