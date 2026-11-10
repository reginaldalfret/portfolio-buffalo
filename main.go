package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/plush/v5"
)

//go:embed templates/*
var templateFS embed.FS

//go:embed public/*
var publicFS embed.FS

func renderHTML() (string, error) {
	tmplBytes, err := templateFS.ReadFile("templates/index.plush.html")
	if err != nil {
		return "", fmt.Errorf("reading template: %w", err)
	}

	data := GetPortfolioData()
	ctx := plush.NewContextWith(data)

	out, err := plush.Render(string(tmplBytes), ctx)
	if err != nil {
		return "", fmt.Errorf("rendering plush: %w", err)
	}
	return out, nil
}

func main() {
	exportFlag := flag.Bool("export", false, "Export static HTML to target directory")
	exportDir := flag.String("out", `E:\Portfolio\public\stacks\buffalo`, "Output directory for export")
	flag.Parse()

	if *exportFlag {
		fmt.Printf("Rendering static Buffalo portfolio to %s...\n", *exportDir)
		html, err := renderHTML()
		if err != nil {
			fmt.Printf("Error rendering: %v\n", err)
			os.Exit(1)
		}

		if err := os.MkdirAll(*exportDir, 0755); err != nil {
			fmt.Printf("Error creating dir: %v\n", err)
			os.Exit(1)
		}

		// Write index.html
		indexPath := filepath.Join(*exportDir, "index.html")
		if err := os.WriteFile(indexPath, []byte(html), 0644); err != nil {
			fmt.Printf("Error writing index.html: %v\n", err)
			os.Exit(1)
		}

		// Copy static assets (pdf, favicon)
		for _, name := range []string{"Reginald_Alfret_Resume.pdf", "favicon.ico"} {
			b, err := publicFS.ReadFile("public/" + name)
			if err == nil {
				os.WriteFile(filepath.Join(*exportDir, name), b, 0644)
			}
		}

		fmt.Printf("Successfully exported %d bytes to %s\n", len(html), indexPath)
		return
	}

	port := envy.Get("PORT", "8087")
	host := envy.Get("HOST", "0.0.0.0")
	app := buffalo.New(buffalo.Options{
		Env:  "development",
		Addr: host + ":" + port,
	})

	r := render.New(render.Options{
		TemplatesFS: templateFS,
		HTMLLayout:  "",
	})

	handler := func(c buffalo.Context) error {
		html, err := renderHTML()
		if err != nil {
			return c.Error(http.StatusInternalServerError, err)
		}
		c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
		return c.Render(http.StatusOK, r.String(html))
	}

	app.GET("/", handler)
	app.GET("/stacks/buffalo", handler)

	// Serve static files from embedded FS or disk
	subFS, err := fs.Sub(publicFS, "public")
	if err == nil {
		app.ServeFiles("/", http.FS(subFS))
	} else {
		app.ServeFiles("/", http.Dir("public"))
	}

	fmt.Printf("🦬 Starting Buffalo Portfolio Server on http://127.0.0.1:%s\n", port)
	if err := app.Serve(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
