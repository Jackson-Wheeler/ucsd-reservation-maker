package myplaywright

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

type MyPlaywright struct {
	Playwright *playwright.Playwright
	Browser    playwright.Browser
}

func (pw *MyPlaywright) Initialize(headlessMode bool) error {
	var err error

	// start Playwright
	pw.Playwright, err = playwright.Run()
	if err != nil {
		return fmt.Errorf("could not start Playwright: %v", err)
	}

	// set browser configuration
	browserConfig := playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(headlessMode),
	}

	// attempt to launch browser (download dependencies if necessary)
	pw.Browser, err = pw.Playwright.Chromium.Launch(browserConfig)
	if err != nil {
		// if error, assume browser and OS dependencies not installed
		// -> install browser and OS dependencies
		err = playwright.Install()
		if err != nil {
			return fmt.Errorf("error installing Playwright browser and OS dependencies: %v", err)
		}

		// try to launch browser again
		pw.Browser, err = pw.Playwright.Chromium.Launch(browserConfig)
		if err != nil {
			return fmt.Errorf("error launching browser: %v", err)
		}
	}

	return nil
}

// Close() closes the browser and stops Playwright
func (pw *MyPlaywright) Close() error {
	err := pw.Browser.Close()
	err = pw.Playwright.Stop()
	return err
}

// USAGE:
// page, err := browser.NewPage()
// if err != nil {
// 	log.Fatalf("could not create page: %v", err)
// }
// if _, err = page.Goto("https://www.google.com"); err != nil {
// 	log.Fatalf("could not goto: %v", err)
// }
// time.Sleep(5 * time.Second)
// fmt.Println("End")
// -- end --
