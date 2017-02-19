package search

import (
	"math/rand"
	"net/http"
	"time"
)

var r *rand.Rand

var ua = []string{
	"Mozilla/5.0 (Windows NT 5.1; rv:5.0.1) Gecko/20100101 Firefox/5.0.1",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.85 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
}

func randomUserAgent() string {
	return ua[r.Intn(len(ua))]
}

func (c Client) get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("user-agent", randomUserAgent())

	return c.httpClient.Do(req)
}

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}
