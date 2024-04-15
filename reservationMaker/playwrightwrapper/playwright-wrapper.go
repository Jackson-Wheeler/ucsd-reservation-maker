// playwrightwrapper package: wrapper for the Playwright library. Playwright is a software for easily controlling automated browser. This wrapper class abstracts away some of the details for simpler use cases.
//
// playwrightwrapper.go: contains PlaywrightWrapper struct and initialization and closing functions
package playwrightwrapper

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
// browserContextStrict: if false, will disable strict for browserContext
func (pw *PlaywrightWrapper) Initialize(headlessMode bool, browserContextStrict bool) error {
	var err error
	// start Playwright
	pw.Playwright, err = playwright.Run()
	if err != nil {
		// if error, assume browser and OS dependencies not installed
		// -> install browser and OS dependencies
		err = installPlaywright()
		if err != nil {
			return err
		}

		// start Playwright
		pw.Playwright, err = playwright.Run()
		if err != nil {
			return fmt.Errorf("error starting Playwright: %v", err)
		}
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
		err = installPlaywright()
		if err != nil {
			return err
		}

		// try to launch browser again
		pw.Browser, err = pw.Playwright.Chromium.Launch(browserConfig)
		if err != nil {
			return fmt.Errorf("error launching browser: %v", err)
		}
	}

	// create a new browser context
	browserContextConfig := playwright.BrowserNewContextOptions{
		StrictSelectors: playwright.Bool(browserContextStrict),
	}
	pw.BrowserContext, err = pw.Browser.NewContext(browserContextConfig)
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

// InstallPlaywright() installs the Playwright browser and OS dependencies, asking the
// user for confirmation before doing so
func installPlaywright() error {
	// inform user of installation
	fmt.Println("Automated browser dependencies not detected. They need to be installed for the UCSD Reservation Maker program to run.")
	// ask for confirmation
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Are you ok with the dependencies being installed (this only needs to be done once)? (y/n): ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	// if user confirms, install dependencies
	if text == "y" || text == "Y" {
		err := playwright.Install()
		if err != nil {
			return fmt.Errorf("error installing Playwright browser and OS dependencies: %v", err)
		}
	} else {
		fmt.Println("Exiting program. Dependencies not installed.")
		os.Exit(0)
	}

	return nil
}
