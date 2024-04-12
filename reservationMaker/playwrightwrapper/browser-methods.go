// page-methods.go: browser methods for the PlaywrightWrapper struct
package playwrightwrapper

import (
	"fmt"
)

// SwitchPage switches the CurrPage in PlaywrightWrapper to the Page the given index in the browser context's list of pages.
func (pw *PlaywrightWrapper) SwitchPage(idx int) error {
	pages := pw.BrowserContext.Pages()
	if idx < 0 || idx >= len(pages) {
		return fmt.Errorf("invalid page index: %d. Pages length: %d", idx, len(pages))
	}
	pw.CurrPage = pages[idx]
	return nil
}
