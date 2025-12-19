package stealth

import "github.com/go-rod/rod"

func ApplyFingerprint(page *rod.Page) {
	page.MustEval(`
		Object.defineProperty(navigator, 'webdriver', { get: () => undefined });
		Object.defineProperty(navigator, 'languages', { get: () => ['en-US', 'en'] });
		Object.defineProperty(navigator, 'platform', { get: () => 'Win32' });
	`)
}
