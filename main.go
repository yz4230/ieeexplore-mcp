package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yz4230/ieeexplore-mcp/ieeexplore"
)

type searchInput struct {
	Query           string `json:"query" jsonschema:"Short IEEE Xplore search terms, usually 2-6 technical keywords"`
	Page            int    `json:"page,omitempty" jsonschema:"Pagination page number, starting at 1"`
	ArticlesPerPage int    `json:"articlesPerPage,omitempty" jsonschema:"How many articles to return per page, default 25"`
}

type searchResultEntry struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Publication string   `json:"publication"`
	Year        string   `json:"year"`
	DOI         string   `json:"doi"`
	Abstract    string   `json:"abstract"`
}

type searchResult struct {
	CurrentPage   int                 `json:"current_page"`
	TotalPages    int                 `json:"total_pages"`
	TotalArticles int                 `json:"total_articles"`
	Articles      []searchResultEntry `json:"articles"`
}

type getArticleInput struct {
	ID string `json:"id" jsonschema:"IEEE Xplore document number from search results, not a DOI"`
}

type getArticleResult struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Publication string   `json:"publication"`
	Year        string   `json:"year"`
	DOI         string   `json:"doi"`
	Abstract    string   `json:"abstract"`
	Content     string   `json:"content"`
}

const (
	searchToolDescription = "Search IEEE Xplore for papers and return concise metadata. " +
		"Use this first for discovery, related-work lookups, and title matching. " +
		"Keep queries short and technical, usually 2-6 essential keywords rather than full natural-language questions. " +
		"If the user knows only a title or topic, search first to identify the IEEE document number. " +
		"Use get_article with a returned id when you need the abstract, metadata, or available article page content."
	getArticleToolDescription = "Fetch metadata and available article page content as Markdown for an IEEE Xplore document number. " +
		"Use the id returned by search, not a DOI. " +
		"Prefer this after search when the user wants to read or summarize a specific paper."
	serverInstructions = "Use search for paper discovery and get_article for full article details.\n" +
		"Search queries should be short, technical keyword phrases rather than long questions.\n" +
		"When a user provides only a title or topic, search first to resolve the IEEE document number.\n" +
		"The get_article id is the IEEE document/article number returned by search, not a DOI.\n" +
		"When summarizing results, preserve the article id, DOI, publication, and year."
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "ieeexplore-search",
		Title:   "IEEE Xplore Search MCP",
		Version: "0.1",
	}, &mcp.ServerOptions{
		Instructions: serverInstructions,
		Logger:       logger,
	})

	server.AddReceivingMiddleware(func(mh mcp.MethodHandler) mcp.MethodHandler {
		return func(ctx context.Context, method string, req mcp.Request) (result mcp.Result, err error) {
			if ctr, ok := req.(*mcp.CallToolRequest); ok {
				logger.Info("Tool called", "method", ctr.Params.Name, "input", string(ctr.Params.Arguments))
			}
			return mh(ctx, method, req)
		}
	})

	client := ieeexplore.NewClient()

	mcp.AddTool(server, &mcp.Tool{
		Name:        "search",
		Title:       "IEEE Xplore Search",
		Description: searchToolDescription,
		Annotations: &mcp.ToolAnnotations{
			ReadOnlyHint:    true,
			DestructiveHint: new(false),
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest, input searchInput) (*mcp.CallToolResult, searchResult, error) {
		if input.Page == 0 {
			input.Page = 1
		}
		if input.ArticlesPerPage == 0 {
			input.ArticlesPerPage = ieeexplore.DefaultArticlesPerPage
		}
		result, err := client.Search(input.Query, input.Page, input.ArticlesPerPage)
		if err != nil {
			return nil, searchResult{}, err
		}
		var res searchResult
		res.Articles = make([]searchResultEntry, 0, len(result.Records))
		res.CurrentPage = input.Page
		res.TotalPages = result.TotalPages
		res.TotalArticles = result.TotalRecords
		for _, record := range result.Records {
			authors := make([]string, 0, len(record.Authors))
			for _, author := range record.Authors {
				authors = append(authors, author.PreferredName)
			}
			entry := searchResultEntry{
				ID:          record.ArticleNumber,
				Title:       record.ArticleTitle,
				Authors:     authors,
				Publication: record.PublicationTitle,
				Year:        record.PublicationYear,
				DOI:         record.DOI,
				Abstract:    record.Abstract,
			}
			res.Articles = append(res.Articles, entry)
		}

		return nil, res, nil
	})

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_article",
		Title:       "Get IEEE Xplore Article",
		Description: getArticleToolDescription,
		Annotations: &mcp.ToolAnnotations{
			ReadOnlyHint:    true,
			DestructiveHint: new(false),
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest, input getArticleInput) (*mcp.CallToolResult, getArticleResult, error) {
		article, err := client.GetArticle(input.ID)
		if err != nil {
			return nil, getArticleResult{}, err
		}
		res := getArticleResult{
			ID:          article.ID,
			Title:       article.Title,
			Authors:     article.Authors,
			Publication: article.Publication,
			Year:        article.Year,
			DOI:         article.DOI,
			Abstract:    article.Abstract,
			Content:     article.Content,
		}
		return nil, res, nil
	})

	addr := ":8080"
	port := os.Getenv("PORT")
	if port != "" {
		addr = fmt.Sprintf(":%s", port)
	}

	logger.Info("Starting server", "addr", addr)

	handler := mcp.NewStreamableHTTPHandler(func(r *http.Request) *mcp.Server { return server }, nil)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
