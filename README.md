# Readwise-like API in Go - **EDUCATIONAL PROJECT**

**This is a LEARNING PROJECT designed to teach Go web development concepts.**

A Go-based API service for managing and storing Kindle highlights, inspired by Readwise. This project demonstrates fundamental web development concepts including HTTP servers, database operations, file handling, and RESTful API design.

**Perfect for:**
- Learning Go web development
- Understanding HTTP APIs
- Database integration with Go
- File upload and processing
- RESTful service architecture

## **CREDIT & ATTRIBUTION**

**This project is a LEARNING IMPLEMENTATION based on the excellent work by [sikozonpc](https://github.com/sikozonpc/readwise-in-go).**

- **Original Project**: [sikozonpc/readwise-in-go](https://github.com/sikozonpc/readwise-in-go)
- **YouTube Tutorial**: [Readwise in Go Tutorial](https://youtu.be/Q3TT1WH7Dl8?si=84KIqMtY8JgryqXC)

**Purpose**: This repository exists purely for educational purposes to help developers learn Go web development concepts.

## **Learning Objectives**

This project demonstrates key Go web development concepts:

- **HTTP Server Creation**: Building a web server from scratch
- **Database Operations**: MySQL integration with Go
- **File Processing**: Handling file uploads and JSON parsing
- **API Design**: RESTful endpoint design and implementation
- **Error Handling**: Proper error management in Go
- **Project Structure**: Organizing Go code into logical packages
- **Environment Configuration**: Managing configuration safely

**What You'll Learn:**
- How to create HTTP servers in Go
- Database connectivity and operations
- File upload handling
- JSON processing
- API endpoint design
- Go project organization
- Security best practices

## Features

- **File Upload**: Accept Kindle highlight JSON files via HTTP POST
- **JSON Parsing**: Parse Kindle highlight files and extract book information
- **Database Storage**: Store books and highlights in MySQL database
- **RESTful API**: Clean HTTP endpoints for file processing
- **User Management**: Support for multiple users with unique IDs
- **Health Monitoring**: Built-in health check endpoints

## Architecture

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   Postman   │───▶│  Go Server  │───▶│   Service   │───▶│   MySQL    │
│   Client    │    │   (Port 8080)│    │   Layer     │    │  Database  │
└─────────────┘    └─────────────┘    └─────────────┘    └─────────────┘
```

## Project Structure

```
readwise/
├── api.go          # HTTP server and routing
├── db.go           # Database operations and MySQL store
├── main.go         # Application entry point
├── service.go      # Business logic and file processing
├── store.go        # Storage interface definitions
├── types.go        # Data structures and types
├── utils.go        # Utility functions
├── .envrc          # Environment variables template
├── go.mod          # Go module dependencies
├── README.md       # This file
└── tests/          # Test utilities and sample data
    ├── run_tests.go    # Test runner for populating sample data
    └── README.md       # Test documentation and safety notes
```

## Technologies Used

- **Go 1.24+**: Core programming language
- **Gorilla Mux**: HTTP router and URL matcher
- **MySQL**: Database for storing books and highlights
- **XAMPP**: Local development environment (optional)

## Prerequisites

- Go 1.24 or higher
- MySQL server (or XAMPP for local development)
- Git

## Installation & Setup

### 1. Clone the Repository
```bash
git clone <your-repo-url>
cd readwise
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Set Up Environment Variables
Copy the `.envrc` file and configure your database settings:
```bash
cp .envrc .envrc.local
# Edit .envrc.local with your database credentials
```

### 4. Start MySQL Server
```bash
# Using XAMPP (recommended for development)
# Start MySQL from XAMPP Control Panel

# Or using command line
& "C:\xampp\mysql\bin\mysqld.exe" --defaults-file="C:\xampp\mysql\bin\my.ini" --standalone
```

### 5. Build and Run
```bash
go build .
.\readwise-go.exe
```

### 6. (Optional) Run Test Instructions
```bash
# This will show you how to test the API
go run tests/run_tests.go
```

## API Endpoints

### Health Check
```
GET /health
Response: {"status": "healthy", "service": "readwise-api"}
```

### Home Page
```
GET /
Response: Welcome message with available endpoints
```

### Parse Kindle File
```
POST /api/v1/users/{user_id}/parse-kindle-file
Content-Type: multipart/form-data
Body: file=highlights.json
Response: {"message": "Highlights saved successfully"}
```

### Daily Insights (Placeholder)
```
GET /api/v1/cloud/send-daily-insights
Response: {"message": "Daily insights feature coming soon", "status": "not implemented yet"}
```

## Database Schema

### Books Table
```sql
CREATE TABLE books (
    id VARCHAR(255) PRIMARY KEY,
    asin VARCHAR(255) NOT NULL UNIQUE,
    title TEXT NOT NULL,
    authors TEXT NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### Highlights Table
```sql
CREATE TABLE highlights (
    id VARCHAR(255) PRIMARY KEY,
    book_id VARCHAR(255) NOT NULL,
    highlight TEXT NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (book_id) REFERENCES books(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

### Users Table
```sql
CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

## Testing with Postman

1. **Start the server** (see Installation steps above)
2. **Open Postman** and create a new request
3. **Set method to POST** and URL to: `http://localhost:8080/api/v1/users/user_001/parse-kindle-file`
4. **Set Body to form-data** with key `file` and type `File`
5. **Select your highlights.json file** and click Send

## Security Features

- **Environment Variables**: Database credentials stored in environment variables
- **Input Validation**: File upload validation and error handling
- **SQL Injection Protection**: Parameterized queries for database operations
- **User Isolation**: Highlights are stored per user ID

## Kindle File Format

The API expects Kindle highlight files in this JSON format:
```json
{
  "asin": "B0DXC331LY",
  "title": "Book Title",
  "authors": "Author Name",
  "highlights": [
    {
      "text": "Highlighted text content",
      "isNoteOnly": false,
      "location": {
        "url": "kindle://book?action=open&asin=B0DXC331LY&location=24",
        "value": 24
      },
      "note": null
    }
  ]
}
```

## **Learning & Contributing**

This is an **EDUCATIONAL PROJECT** designed for learning Go web development. Feel free to:

- **Study the code** to understand Go web development patterns
- **Fork and experiment** with different approaches
- **Modify and extend** to practice your skills
- **Report issues** to help improve the learning experience
- **Share your learnings** and improvements
- **Use as a reference** for your own Go projects

**Remember**: This is for learning purposes - feel free to break it, fix it, and learn from it!

## **Educational License**

**This project is for EDUCATIONAL PURPOSES ONLY.**

- **Use**: Feel free to use this code for learning Go web development
- **Modify**: Experiment with the code to understand how it works
- **Share**: Share your learnings and improvements

**Not for production use** without proper security review and testing.

## **Learning Resources**

- **Go Documentation**: [golang.org](https://golang.org) - Official Go language reference
- **Gorilla Mux**: [github.com/gorilla/mux](https://github.com/gorilla/mux) - HTTP router documentation
- **Inspiration**: Readwise service for highlighting and note-taking

## **Learning Support**

**For Learning Questions:**
- Open an issue in this repository for questions about the Go implementation
- Ask about Go web development concepts and patterns
- Request clarification on any part of the code

**For Original Project Questions:**
- Refer to the [original repository](https://github.com/sikozonpc/readwise-in-go) or [YouTube tutorial](https://youtu.be/Q3TT1WH7Dl8?si=84KIqMtY8JgryqXC)

---

## **Quick Start for Learners**

1. **Clone and explore** the code structure
2. **Follow the tutorial** to understand the concepts
3. **Experiment** with modifications
4. **Build your own** Go web services using these patterns
5. **Share** what you learn with the community

**This is your playground for learning Go web development!**
