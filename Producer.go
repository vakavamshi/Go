package main
import (
	"fmt"
	"time"
)

func Producer(a int) {
	for b := 0; b < 3; b++ {
		i := []int{}
		i = append(i, a)
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go Producer(1)
	Producer(2)

}

