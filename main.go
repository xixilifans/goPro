package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix() > time.Now().AddDate(0, -1, 0).Unix())
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now())

	format := "2006-01-02 15:04:05"
	a, _ := time.Parse(format, "2019-03-10 11:00:00")

	b, _ := time.Parse(format, "2015-03-10 16:00:00")

	fmt.Println("\na:", a.Format(format), "\nb:", b.Format(format))

	fmt.Println("a  b   After:", a.After(b))

	fmt.Println("a  b   Before:", a.Before(b))

}
