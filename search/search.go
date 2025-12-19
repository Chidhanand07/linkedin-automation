package search

import (
	"strings"
	"time"

	"github.com/go-rod/rod"
)

func FindProfiles(page *rod.Page) ([]string, error) {
	page.MustNavigate("https://www.linkedin.com/search/results/people/")
	time.Sleep(6 * time.Second)

	var profiles []string
	seen := make(map[string]bool)

	links, err := page.Elements("a.app-aware-link")
	if err != nil {
		return nil, err
	}

	for _, link := range links {
		href, err := link.Attribute("href")
		if err != nil || href == nil {
			continue
		}

		url := *href
		if strings.Contains(url, "/in/") && !strings.Contains(url, "overlay") {
			if !seen[url] {
				seen[url] = true
				profiles = append(profiles, url)
			}
		}
	}

	return profiles, nil
}
