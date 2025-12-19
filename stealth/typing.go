package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func HumanType(el *rod.Element, text string) {
	for _, c := range text {
		el.MustInput(string(c))
		time.Sleep(time.Duration(rand.Intn(120)+60) * time.Millisecond)
	}
}
