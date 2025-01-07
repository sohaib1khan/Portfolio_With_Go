package main

import (
	"fmt"
	"html/template" // Provides support for HTML templates
	"io/ioutil"     // Used to read files, such as Markdown content
	"log"           // Provides logging functionality
	"net/http"      // Used to create a web server
	"os"            // Provides functions to read environment variables
	"path/filepath" // For constructing file paths

	"github.com/gomarkdown/markdown" // Converts Markdown content to HTML
)

// Function to render templates with Markdown content
func renderTemplate(w http.ResponseWriter, templateName string, data map[string]interface{}) {
	// Parse the specified HTML template file
	tmpl, err := template.ParseFiles(filepath.Join("templates", templateName+".html"))
	if err != nil {
		log.Printf("Error rendering template %s: %v", templateName, err)
		// Send a 500 Internal Server Error response if the template cannot be loaded
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// If a Markdown file is specified in the data, read and convert it to HTML
	if markdownFile, ok := data["MarkdownFile"].(string); ok {
		// Read the Markdown file content
		content, err := ioutil.ReadFile(filepath.Join("content", markdownFile))
		if err != nil {
			log.Printf("Error reading markdown file %s: %v", markdownFile, err)
			// If there is an error reading the file, provide a fallback HTML message
			data["MarkdownHTML"] = template.HTML("<p>Error loading content</p>")
		} else {
			// Convert the Markdown content to HTML
			htmlContent := markdown.ToHTML(content, nil, nil)
			data["MarkdownHTML"] = template.HTML(htmlContent)
		}
	}

	// Execute the template, injecting the provided data
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template %s: %v", templateName, err)
		// Send a 500 Internal Server Error response if template execution fails
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Handlers for each page
// Each handler corresponds to a specific route and uses renderTemplate to serve the page

// Handler for the Home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home", map[string]interface{}{
		"Title":           "Home",                           // Title of the page
		"MarkdownFile":    "home.md",                        // Markdown file to render
		"UserProfileImage": "/static/images/profile.jpg",    // Example data: Profile image URL
	})
}

// Handler for the Projects page
func projectsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "projects", map[string]interface{}{
		"Title":        "Projects",      // Title of the page
		"MarkdownFile": "projects.md",   // Markdown file to render
	})
}

// Handler for the TechStack page
func TechStackHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "TechStack", map[string]interface{}{
		"Title":        "TechStack",        // Title of the page
		"MarkdownFile": "TechStack.md",     // Markdown file to render
	})
}

// Handler for the Experience page
func experienceHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "experience", map[string]interface{}{
		"Title":        "Experience",    // Title of the page
		"MarkdownFile": "experience.md", // Markdown file to render
	})
}

// Handler for the Contact page
func contactHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact", map[string]interface{}{
		"Title":        "Contact Me",    // Title of the page
		"MarkdownFile": "contact.md",    // Markdown file to render
	})
}

// Main function: Entry point of the application
func main() {
	// Read the port number from the environment variable "PORT", default to "8989"
	port := os.Getenv("PORT")
	if port == "" {
		port = "8989"
	}

	// Read the environment (e.g., "development" or "production") from the environment variable "ENVIRONMENT", default to "development"
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	// Print a banner message when the application starts
	banner := `
=============================================================
Portfolio App
=============================================================
Environment: %s
Server running on: http://localhost:%s
=============================================================
`
	fmt.Printf(banner, env, port)

	// Define routes and link them to handlers
	http.HandleFunc("/", homeHandler)                // Home page
	http.HandleFunc("/projects", projectsHandler)    // Projects page
	http.HandleFunc("/TechStack", TechStackHandler)        // TechStack page
	http.HandleFunc("/experience", experienceHandler)// Experience page
	http.HandleFunc("/contact", contactHandler)      // Contact page

	// Serve static files (e.g., CSS, JS, images) from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	// Start the HTTP server and listen on the specified port
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
		// Log an error and terminate if the server fails to start
		log.Fatalf("Failed to start server: %v", err)
	}
}
