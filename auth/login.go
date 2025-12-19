package auth

import (
	"errors"
	"os"
	"time"

	"github.com/go-rod/rod"
)

func Login(page *rod.Page) error {
	page.MustNavigate("https://www.linkedin.com/login")
	time.Sleep(4 * time.Second)

	email := os.Getenv("LINKEDIN_EMAIL")
	password := os.Getenv("LINKEDIN_PASSWORD")

	if email == "" || password == "" {
		return errors.New("credentials not found in environment variables")
	}

	emailInput, err := page.Element("#username")
	if err != nil {
		return errors.New("email input not found")
	}

	passwordInput, err := page.Element("#password")
	if err != nil {
		return errors.New("password input not found")
	}

	emailInput.MustInput(email)
	time.Sleep(time.Second)

	passwordInput.MustInput(password)
	time.Sleep(time.Second)

	submitBtn, err := page.Element("button[type=submit]")
	if err != nil {
		return errors.New("login button not found")
	}

	submitBtn.MustClick()
	time.Sleep(6 * time.Second)

	if page.MustHas("input[name=pin]") || page.MustHas("iframe[src*='captcha']") {
		return errors.New("security checkpoint detected")
	}

	return nil
}
