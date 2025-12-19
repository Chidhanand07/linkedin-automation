package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func HumanMove(page *rod.Page, x, y float64) {
	steps := rand.Intn(20) + 20
	cx, cy := rand.Float64()*50, rand.Float64()*50

	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)
		nx := (1-t)*(1-t)*cx + 2*(1-t)*t*x + t*t*x
		ny := (1-t)*(1-t)*cy + 2*(1-t)*t*y + t*t*y
		page.Mouse.Move(nx, ny)
		time.Sleep(time.Duration(rand.Intn(15)+5) * time.Millisecond)
	}
}
