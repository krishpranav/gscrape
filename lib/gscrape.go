package lib

import (
	"errors"
	"net/url"
	"strings"
)

func root(u *url.URL) string {
	return u.Scheme + "://" + u.Host
}

func fixURL(u *url.URL) *url.URL {
	if u.Scheme == "" {
		s := strings.Split(u.String(), "/")
		u.Scheme = "http"
		u.Host = s[0]
		u.Path = u.Path[len(s[0]):]
	}
	return u
}

type gscrape struct {
	url           string
	scheme        string
	host          string
	rootUrl       string
	title         string
	language      string
	author        string
	description   string
	generator     string
	feed          string
	charset       string
	links         []string
	images        []string
	keywords      []string
	compatibility map[string]string
}

func New(uri string) (*gscrape, error) {
	if uri == "" {
		return nil, errors.New("Url is empty!")
	}

	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	scraper, err := newScraper(fixURL(u), 20)
	if err != nil {
		return nil, err
	}

	return &gscrape{uri,
		u.Scheme,
		u.Host,
		root(u),
		scraper.title,
		scraper.language,
		scraper.author,
		scraper.description,
		scraper.generator,
		scraper.feed,
		scraper.charset,
		scraper.links,
		scraper.images,
		scraper.keywords,
		scraper.compatibility}, nil
}

func (m gscrape) Url() string {
	return m.url
}
func (m gscrape) Scheme() string {
	return m.scheme
}

func (m gscrape) Host() string {
	return m.host
}

func (m gscrape) RootURL() string {
	return m.rootUrl
}

func (m gscrape) Title() string {
	return m.title
}

func (m gscrape) Language() string {
	return m.language
}

func (m gscrape) Author() string {
	return m.author
}

func (m gscrape) Description() string {
	return m.description
}

func (m gscrape) Generator() string {
	return m.generator
}

func (m gscrape) Feed() string {
	return m.feed
}

func (m gscrape) Charset() string {
	return m.charset
}

func (m gscrape) Links() []string {
	return m.links
}

func (m gscrape) Images() []string {
	return m.images
}

func (m gscrape) Keywords() []string {
	return m.keywords
}

func (m gscrape) Compatibility() map[string]string {
	return m.compatibility
}
