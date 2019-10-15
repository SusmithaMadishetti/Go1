package main

import (
	"fmt"
	"time"
	//"math/rand"
)

// Create the Player stuct that implements the Playbacker interface.

// Playbacker is an interface that defines functionality to add items to a queue and play those items back.
type Playbacker interface {
	// Set the playback delay in seconds.
	SetInterval(int)
	// Adds a single value of any type to the queue.
	Add(interface{})
	// Prints each item in the queue, in the order they were added (fifo), with a delay between each item.
	Play() error
}

// Player implements the Playbacker interface
type Player struct {
	Stut  []interface{}
	Delay int
}

func (p *Player) SetInterval(a int) {
	p.Delay = a
}

func (p *Player) Add(i interface{}) {
	p.Stut = append(p.Stut, i)

}
func (p Player) Play() {
	for _, v := range p.Stut {
		fmt.Println(v)
		//time.Sleep(time.Duration(rand.Intn(p.Delay)) * time.Second)
		time.Sleep(time.Duration(p.Delay) * time.Second)

	}
}
func main() {
	p := Player{}
	p.SetInterval(3)

	p.Add("nine")
	p.Add(9)
	p.Add("this is the last item")

	p.Play()
}
