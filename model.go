package dmm

import (
	"fmt"
	"regexp"
)

// Product represents a product entity.
type Product struct {
	ContentID string
	Title     string
	DetailURL string
	CoverURL  string
	Artworks  []string
}

func (p Product) String() string {
	return fmt.Sprintf("%s %s", p.ContentID, p.CoverURL)
}

var contentIDRegex = regexp.MustCompile("cid=([^/]+)")

// ParseContentID extract (possible) content id from string.
func ParseContentID(s string) string {
	p := contentIDRegex.FindStringSubmatch(s)
	if len(p) >= 0 {
		return p[1]
	}
	return ""
}

var (
	coverURLRegex = regexp.MustCompile("p[tsl]\\.")
	schemeRegex   = regexp.MustCompile("^//")
)

// ParseLargeCoverURL convert a url to large size.
func ParseLargeCoverURL(url string) (u string) {
	u = url
	u = coverURLRegex.ReplaceAllString(u, "pl.")
	u = schemeRegex.ReplaceAllString(u, "http://")

	return
}
