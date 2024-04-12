// playwrightwrapper package: wrapper for the Playwright library. Playwright is a software for easily controlling automated browser. This wrapper class abstracts away some of the details for simpler use cases.
//
// playwrightwrapper.go: contains PlaywrightWrapper struct and initialization and closing functions
package playwrightwrapper

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

type PlaywrightWrapper struct {
	Playwright     *playwright.Playwright
	Browser        playwright.Browser
	BrowserContext playwright.BrowserContext
	CurrPage       playwright.Page
}

// Initialize() initializes the PlaywrightWrapper struct. It starts Playwright, launches a browser, creates a browser context, and creates a new page on that browser (assigned to CurrPage).
//
// headlessMode: if true, the browser will be launched in headless mode (no GUI)
func (pw *PlaywrightWrapper) Initialize(headlessMode bool) error {
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

	// create a new browser context
	pw.BrowserContext, err = pw.Browser.NewContext()
	if err != nil {
		return fmt.Errorf("could not create browser context: %v", err)
	}

	// create new page
	pw.CurrPage, err = pw.BrowserContext.NewPage()
	if err != nil {
		return fmt.Errorf("could not create new page: %v", err)
	}

	return nil
}

// Close() closes the browser context, the browser, and stops Playwright
func (pw *PlaywrightWrapper) Close() error {
	if err := pw.BrowserContext.Close(); err != nil {
		return fmt.Errorf("could not close browser context: %v", err)
	}
	if err := pw.Browser.Close(); err != nil {
		return fmt.Errorf("could not close browser: %v", err)
	}
	if err := pw.Playwright.Stop(); err != nil {
		return fmt.Errorf("could not stop Playwright: %v", err)
	}
	return nil
}
