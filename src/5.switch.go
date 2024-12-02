package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 2:
		fmt.Println(i)
	}
	weekday := func() {
		fmt.Println(time.Now().Weekday())
	}
	weekday()

}
