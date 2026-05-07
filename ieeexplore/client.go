package ieeexplore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/PuerkitoBio/goquery"
)

type Client struct {
	http *http.Client
}

func NewClient() *Client {
	return &Client{
		http: http.DefaultClient,
	}
}

type searchRequest struct {
	NewSearch    bool     `json:"newsearch"`
	QueryText    string   `json:"queryText"`
	Highlight    bool     `json:"highlight"`
	ReturnFacets []string `json:"returnFacets"`
	ReturnType   string   `json:"returnType"`
	MatchPubs    bool     `json:"matchPubs"`
	RowsPerPage  int      `json:"rowsPerPage"`
}

func (c *Client) Search(query string) (*SearchResult, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, fmt.Errorf("query must not be empty")
	}

	payload, err := json.Marshal(searchRequest{
		NewSearch:    true,
		QueryText:    query,
		Highlight:    false,
		ReturnFacets: []string{"ALL"},
		ReturnType:   "SEARCH",
		MatchPubs:    true,
		RowsPerPage:  100,
	})
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://ieeexplore.ieee.org/rest/search", bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	setSearchHeaders(req)
	req.Header.Set("content-type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("search IEEE Xplore: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 10<<20))
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("IEEE Xplore returned %s: %s", resp.Status, strings.TrimSpace(string(body)))
	}

	var result SearchResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	return &result, nil
}

type Article struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Publication string   `json:"publication"`
	Year        string   `json:"year"`
	DOI         string   `json:"doi"`
	Abstract    string   `json:"abstract"`
	Content     string   `json:"content"`
}

func (c *Client) GetArticle(id string) (*Article, error) {
	endpoint := fmt.Sprintf("https://ieeexplore.ieee.org/document/%s", id)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	setSearchHeaders(req)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get article: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parse article page: %w", err)
	}

	var metadataJSON string
	const prefix = "xplGlobal.document.metadata="
	doc.Find("script").EachWithBreak(func(i int, s *goquery.Selection) bool {
		scriptContent := s.Text()
		for line := range strings.Lines(scriptContent) {
			line = strings.TrimSpace(line)
			if content, ok := strings.CutPrefix(line, prefix); ok {
				if content, ok = strings.CutSuffix(content, ";"); ok {
					metadataJSON = content
				}
			}
		}
		return true
	})

	if metadataJSON == "" {
		return nil, fmt.Errorf("metadata not found in article page")
	}

	var metadata XPLGlobal
	if err = json.Unmarshal([]byte(metadataJSON), &metadata); err != nil {
		return nil, fmt.Errorf("decode metadata JSON: %w", err)
	}

	authors := make([]string, 0, len(metadata.Authors))
	for _, author := range metadata.Authors {
		authors = append(authors, author.Name)
	}

	content, err := c.fetchArticleContent(id)
	if err != nil {
		return nil, fmt.Errorf("fetch article content: %w", err)
	}

	return &Article{
		ID:          id,
		Title:       metadata.Title,
		Authors:     authors,
		Publication: metadata.PublicationTitle,
		Year:        metadata.PublicationYear,
		DOI:         metadata.DOI,
		Abstract:    metadata.Abstract,
		Content:     content,
	}, nil
}

func (c *Client) fetchArticleContent(id string) (string, error) {
	endpoint := fmt.Sprintf("https://ieeexplore.ieee.org/rest/document/%s", id)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	setSearchHeaders(req)

	resp, err := c.http.Do(req)
	if err != nil {
		return "", fmt.Errorf("fetch article content: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("parse article content: %w", err)
	}

	// unwrap anchor tags with href="javascript:void()"
	doc.Find("a[href='javascript:void()']").Each(func(i int, s *goquery.Selection) {
		s.ReplaceWithHtml(s.Text())
	})

	html, err := doc.Html()
	if err != nil {
		return "", fmt.Errorf("get article content HTML: %w", err)
	}

	markdown, err := htmltomarkdown.ConvertString(html)
	if err != nil {
		return "", fmt.Errorf("convert HTML to markdown: %w", err)
	}

	return string(markdown), nil
}

func setSearchHeaders(req *http.Request) {
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "en-US,en;q=0.9,ja;q=0.8")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("origin", "https://ieeexplore.ieee.org")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://ieeexplore.ieee.org")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/147.0.0.0 Safari/537.36")
	req.Header.Set("x-security-request", "required")
}
