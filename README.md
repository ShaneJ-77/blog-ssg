# Static Site Generator 
This is a static site generator I created for the purpose of converting markdown files into html blog posts for my vps.

## Project Design
### Markdown Files With YAML Metadata
### Go through the steps of what the program does
### How does it fit into my vps 

## Project Layout
```
cmd
│   └── main.go
│
content
│   └── test.md
internal
│   ├── generator
│   │   └── html.go
│   ├── parser
│   │   ├── frontmatter.go
│   │   └── markdown.go
public
│   ├── test-post
│   │   └── index.html
│   └── next-post
│       └── index.html
templates
│  └── default.html
.gitignore
go.mod
go.sum
README.md
```
## Tech Stack
- go 1.24.0
- html/template
- github.com/yuin/goldmark v1.8.6
- gopkg.in/yaml.v3

## Usage
- Clone repo
- go get dependencies
- place markdown files into /content
- place html templates into /templates
- go run ./cmd

