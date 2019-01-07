package main

import (
	"fmt"
	"time"
)

type Rocket struct {
	Num int
}

func (r *Rocket) Launch() {
	fmt.Println(r.Num)
}

func main() {
	r := new(Rocket)
	r.Launch()
	time.AfterFunc(5*time.Second, r.Launch)
}
