package open

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Akagi201/spy/crawler"
	"github.com/skratchdot/open-golang/open"
)

// Open - Opens crawl response in browser.
func Open(resp *crawler.Response) error {
	fname := filepath.Join(os.TempDir(), fmt.Sprintf("%s.html", randFileName()))
	if err := crawler.WriteResponseFile(resp, fname); err != nil {
		return err
	}
	return open.Start(fmt.Sprintf("file://%s", fname))
}

// Handler - Crawl handler that opens crawl response in browser.
func Handler(_ context.Context, resp *crawler.Response) error {
	return Open(resp)
}

func randFileName() string {
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		return "temp-file"
	}
	return fmt.Sprintf("%x", b)
}
