package search

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/bcho/dmm"
)

// QueryOpts sets query options.
type QueryOpts struct {
	Keyword string
}

// Query performs query to dmm.co.jp
func (c Client) Query(queryOpts *QueryOpts) (ps []dmm.Product, err error) {
	queryURL := c.queryURLBuilder(queryOpts.Keyword)
	resp, err := c.get(queryURL)
	if err != nil {
		return
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return
	}

	doc.Find("#list > li").Each(func(_ int, s *goquery.Selection) {
		detailLink := s.Find(".tmb a").First()
		detailUrl := detailLink.AttrOr("href", "")

		if detailUrl == "" {
			return
		}

		thumbnail := s.Find(".tmb .img img").First()
		title := thumbnail.AttrOr("alt", "")
		thumbnailUrl := thumbnail.AttrOr("src", "")

		if title == "" || thumbnailUrl == "" {
			return
		}

		ps = append(ps, dmm.Product{
			ContentID: dmm.ParseContentID(detailUrl),
			Title:     title,
			DetailURL: detailUrl,
			CoverURL:  dmm.ParseLargeCoverURL(thumbnailUrl),
		})
	})

	return
}

// Query with default client.
func Query(queryOpts *QueryOpts) ([]dmm.Product, error) {
	return defaultClient.Query(queryOpts)
}
