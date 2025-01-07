# Portfolio with Go

This is a personal portfolio web application built using Go (Golang). It uses a simple and lightweight structure to serve HTML templates, Markdown content, and static files. The application can be deployed locally or on a containerized environment using Docker and Kubernetes.

---

## Project Structure
```
.
├── content                # Markdown files for each page
│   ├── contact.md
│   ├── experience.md
│   ├── home.md
│   ├── projects.md
│   └── skills.md
├── docker-compose.yaml    # Docker Compose configuration
├── Dockerfile             # Dockerfile for building the Go app container
├── go.mod                 # Go module dependencies
├── go.sum                 # Go dependency checksum file
├── k8s                    # Kubernetes manifests
│   ├── deployment.yaml    # Deployment configuration
│   ├── service.yaml       # Service configuration
│   └── skaffold.yaml      # Skaffold configuration
├── main.go                # Main application logic
├── README.md              # Project documentation (this file)
├── skaffold.yaml          # Skaffold pipeline for local development
├── static                 # Static assets (CSS, JS, Images)
│   ├── css
│   │   ├── navigation.css
│   │   └── styles.css
│   ├── images
│   │   └── profile.jpg
│   └── js
│       └── stars.js
└── templates              # HTML templates
    ├── about.html
    ├── contact.html
    ├── experience.html
    ├── home.html
    ├── projects.html
    └── skills.html
```

---

## Prerequisites

- **Go** (1.19 or later)
- **Docker** (if containerization is needed)
- **Kubernetes** (optional, for cluster deployment)
- **Skaffold** (optional, for local Kubernetes development)

---

## Getting Started

### 1. Clone the Repository
```bash
git clone <repository-url>
cd Portfolio_With_Go
```

### 2. Run Locally

#### Using Go
```bash
go run main.go
```
The application will be available at `http://localhost:8989` by default.

#### Using Docker
1. Build the Docker image:
   ```bash
   docker build -t portfolio-with-go .
   ```
2. Run the container:
   ```bash
   docker run -p 8989:8989 portfolio-with-go
   ```

### 3. Kubernetes Deployment
1. Apply the manifests:
   ```bash
   kubectl apply -f k8s/
   ```
2. Access the application using the Kubernetes service's external IP or load balancer.

### 4. Skaffold Development
Use Skaffold for continuous local development:
```bash
skaffold dev
```

---

## Features

- **Markdown Integration**: Render Markdown files as HTML content for pages like "Home," "Projects," and "Contact."
- **Dynamic Templates**: Serve HTML templates with injected data.
- **Static Files**: Serve CSS, JavaScript, and images.
- **Containerization**: Easily deployable via Docker or Kubernetes.
- **Aurora Effect and Shooting Stars**: Visually stunning effects for the background and animations.

---

## Configuration

### Environment Variables
- `PORT`: The port on which the server runs (default: `8989`).
- `ENVIRONMENT`: The environment type (`development` or `production`).

### File Locations
- **Templates**: Located in the `templates/` directory.
- **Markdown Content**: Located in the `content/` directory.
- **Static Assets**: Located in the `static/` directory.

