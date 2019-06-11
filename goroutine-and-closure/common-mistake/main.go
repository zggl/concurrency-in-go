package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}


/*
output:
good day
good day
good day

因为使用了go关键字，for循环里面的逻辑执行是非阻碍的，for循环瞬间执行完了，在三个go routine执行fmt.Println时，三个goroutine取到的都是同一个值，"good day"
*/
