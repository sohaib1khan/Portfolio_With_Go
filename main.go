package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gomarkdown/markdown"
)

// Function to render templates with Markdown content
func renderTemplate(w http.ResponseWriter, templateName string, data map[string]interface{}) {
	// Parse template
	tmpl, err := template.ParseFiles(filepath.Join("templates", templateName+".html"))
	if err != nil {
		log.Printf("Error rendering template %s: %v", templateName, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Convert Markdown content (if present) to HTML
	if markdownFile, ok := data["MarkdownFile"].(string); ok {
		content, err := ioutil.ReadFile(filepath.Join("content", markdownFile))
		if err != nil {
			log.Printf("Error reading markdown file %s: %v", markdownFile, err)
			data["MarkdownHTML"] = template.HTML("<p>Error loading content</p>")
		} else {
			htmlContent := markdown.ToHTML(content, nil, nil)
			data["MarkdownHTML"] = template.HTML(htmlContent)
		}
	}

	// Execute template with data
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template %s: %v", templateName, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Handlers for each page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home", map[string]interface{}{
		"Title":           "Home",
		"MarkdownFile":    "home.md",
		"UserProfileImage": "/static/images/profile.jpg",
	})
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "projects", map[string]interface{}{
		"Title":        "Projects",
		"MarkdownFile": "projects.md",
	})
}

func skillsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "skills", map[string]interface{}{
		"Title":        "Skills",
		"MarkdownFile": "skills.md",
	})
}

func experienceHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "experience", map[string]interface{}{
		"Title":        "Experience",
		"MarkdownFile": "experience.md",
	})
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact", map[string]interface{}{
		"Title":        "Contact Me",
		"MarkdownFile": "contact.md",
	})
}

func main() {
	// Define port and environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8989"
	}

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	// Banner message
	banner := `
=============================================================
Portfolio App
=============================================================
Environment: %s
Server running on: http://localhost:%s
=============================================================
`
	fmt.Printf(banner, env, port)

	// Define routes and handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/projects", projectsHandler)
	http.HandleFunc("/skills", skillsHandler)
	http.HandleFunc("/experience", experienceHandler)
	http.HandleFunc("/contact", contactHandler)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	// Start the server
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
