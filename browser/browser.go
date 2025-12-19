package browser

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

var userAgents = []string{
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64)",
	"Mozilla/5.0 (X11; Linux x86_64)",
}

func NewBrowser() (*rod.Browser, *rod.Page) {
	rand.Seed(time.Now().UnixNano())

	u := launcher.New().
		Bin("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome").
		Headless(false).
		Set("disable-blink-features", "AutomationControlled").
		MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()
	page := browser.MustPage()

	page.MustSetUserAgent(&proto.NetworkSetUserAgentOverride{
		UserAgent: userAgents[rand.Intn(len(userAgents))],
	})

	page.MustEval(`() => {
		Object.defineProperty(navigator, 'webdriver', { get: () => undefined })
		Object.defineProperty(navigator, 'languages', { get: () => ['en-US', 'en'] })
		Object.defineProperty(navigator, 'platform', { get: () => 'MacIntel' })
	}`)

	return browser, page
}
