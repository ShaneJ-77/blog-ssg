# Static Site Generator 
This is a static site generator created for the purpose of converting markdown files into html blog posts to be hosted on my server. 

## Project Design
### Blog Posts Written As Markdown Files With YAML Metadata
The static site files are markdown files with YAML metadata. Markdown was chosen because it is easy to write by hand and the syntax will not interfere with the process of writing. 
### What Does the Static Site Generator Do? 
When run, the program goes through the files in `/content` and converts each markdown file into an html file and saves it to `/public`. If a change has been made to a file that already exists, the program will replace the file with an updated one.
### Github Actions 
This project incorporates a Github Worklow that runs on any push to the main branch that changes the `/content` directory. This workflow runs the static site generator to convert changes in the markdown files into the appropriate html files in `/public` and commits the result.

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



