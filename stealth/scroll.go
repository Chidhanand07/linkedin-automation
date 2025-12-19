package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func HumanScroll(page *rod.Page) {
	scroll := rand.Intn(400) + 200
	page.Mouse.Scroll(0, float64(scroll))
	time.Sleep(time.Duration(rand.Intn(1200)+800) * time.Millisecond)
}
