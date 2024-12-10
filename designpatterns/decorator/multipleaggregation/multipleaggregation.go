package main

import "fmt"

// Aged interface defines methods for getting and setting age
type Aged interface {
	Age() int
	SetAge(age int)
}

// Bird struct implements Aged and has a Fly method
type Bird struct {
	age int
}

func (b *Bird) Age() int       { return b.age }
func (b *Bird) SetAge(age int)  { b.age = age }

func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying")
	}
}

// Lizard struct implements Aged and has a Crawl method
type Lizard struct {
	age int
}

func (l *Lizard) Age() int      { return l.age }
func (l *Lizard) SetAge(age int) { l.age = age }

func (l *Lizard) Crawl() {
	if l.age < 10 {
		fmt.Println("Crawling")
	}
}

// Dragon struct composes Bird and Lizard and implements Aged
type Dragon struct {
	bird   Bird
	lizard Lizard
}

func (d *Dragon) Age() int { return d.bird.age }

func (d *Dragon) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizard.SetAge(age)
}

func (d *Dragon) Fly() {
	d.bird.Fly()
}

func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}

// NewDragon is a constructor function for creating a Dragon instance
func NewDragon() *Dragon {
	return &Dragon{bird: Bird{}, lizard: Lizard{}}
}

func main() {
	d := NewDragon()
	d.SetAge(10)
	d.Fly()
	d.Crawl()
}
