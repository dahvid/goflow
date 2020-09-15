package main

import (
	"fmt"
	"time"
	"math/rand"
	//"github.com/dahvid/goflow"
)


// RandomNumberGenerator Listen for ssh connection
type RandomNumberGenerator struct {
	Generator *rand.Rand
	info map[string]string
	Out chan<- int //output
}


// Process listen
func (c *RandomNumberGenerator) Process() {
	for {
		i := c.Generator.Intn(100)
		fmt.Println("Generating",i)
		c.Out <- i
	}
}

func (c *RandomNumberGenerator) Info() map[string]string {
	return c.info
}



//Plug1 something
func Plug1() (interface{}, error) {
	seed := rand.NewSource(time.Now().UnixNano())
    gen := rand.New(seed)
	r  := new(RandomNumberGenerator)
	r.Generator = gen
	r.info["How are you?"] = "I'm OK"
	return r,nil
}