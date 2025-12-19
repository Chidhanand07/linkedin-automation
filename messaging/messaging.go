package messaging

import (
	"database/sql"
	"errors"
	"time"

	"github.com/go-rod/rod"
)

func SendFollowUps(page *rod.Page, db *sql.DB) error {
	page.MustNavigate("https://www.linkedin.com/mynetwork/invite-connect/connections/")
	time.Sleep(6 * time.Second)

	profiles, err := page.Elements("a.app-aware-link")
	if err != nil {
		return err
	}

	for _, profile := range profiles {
		href, err := profile.Attribute("href")
		if err != nil || href == nil {
			continue
		}

		page.MustNavigate(*href)
		time.Sleep(5 * time.Second)

		msgBtn, err := page.ElementR("button", "Message")
		if err != nil {
			continue
		}

		msgBtn.MustClick()
		time.Sleep(2 * time.Second)

		box, err := page.Element("div[contenteditable='true']")
		if err != nil {
			return errors.New("message input not found")
		}

		message := "Hi, thanks for connecting. Looking forward to staying in touch."
		box.MustInput(message)
		time.Sleep(time.Second)

		sendBtn, err := page.ElementR("button", "Send")
		if err == nil {
			sendBtn.MustClick()
			time.Sleep(3 * time.Second)
		}
	}

	return nil
}
