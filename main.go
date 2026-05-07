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
	Query string `json:"query" jsonschema:"search query for IEEE Xplore"`
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
	TotalEntries int                 `json:"totalEntries"`
	Entries      []searchResultEntry `json:"entries"`
}

type getArticleInput struct {
	ID string `json:"id" jsonschema:"article ID"`
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

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "ieeexplore-search",
		Title:   "IEEE Xplore Search MCP",
		Version: "0.1",
	}, &mcp.ServerOptions{
		Logger: logger,
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
		Description: "Search IEEE Xplore with a query and return a list of articles.",
		Annotations: &mcp.ToolAnnotations{
			ReadOnlyHint:    true,
			DestructiveHint: new(false),
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest, input searchInput) (*mcp.CallToolResult, searchResult, error) {
		result, err := client.Search(input.Query)
		if err != nil {
			return nil, searchResult{}, err
		}
		var res searchResult
		res.TotalEntries = result.TotalRecords
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
			res.Entries = append(res.Entries, entry)
		}

		return nil, res, nil
	})

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_article",
		Description: "Get single article and its content by article ID.",
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
