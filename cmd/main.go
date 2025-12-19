package main

import (
	"log"
	"os"
	"time"

	"linkedin-automation/auth"
	"linkedin-automation/browser"
	"linkedin-automation/connect"
	"linkedin-automation/messaging"
	"linkedin-automation/search"
	"linkedin-automation/storage"
)

func main() {
	email := os.Getenv("LINKEDIN_EMAIL")
	password := os.Getenv("LINKEDIN_PASSWORD")

	if email == "" || password == "" {
		log.Fatal("Environment variables LINKEDIN_EMAIL or LINKEDIN_PASSWORD not set")
	}

	log.Println("Starting LinkedIn Automation POC")

	db, err := storage.InitDB("linkedin_state.db")
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}
	defer db.Close()

	browserInstance, page := browser.NewBrowser()
	defer browserInstance.Close()

	err = auth.Login(page)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	time.Sleep(5 * time.Second)

	profiles, err := search.FindProfiles(page)
	if err != nil {
		log.Fatalf("Profile search failed: %v", err)
	}

	for _, profileURL := range profiles {
		if storage.IsProfileProcessed(db, profileURL) {
			continue
		}

		err := connect.SendConnectionRequest(page, profileURL)
		if err != nil {
			log.Printf("Failed to send request to %s: %v", profileURL, err)
			continue
		}

		storage.MarkProfileProcessed(db, profileURL)
		time.Sleep(20 * time.Second)
	}

	err = messaging.SendFollowUps(page, db)
	if err != nil {
		log.Printf("Messaging error: %v", err)
	}

	log.Println("Automation run completed successfully")
}
