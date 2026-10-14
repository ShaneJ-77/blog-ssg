//generate html pages

package generator

import (
	"html/template"
	"path/filepath"
	"os"
	"bytes"

	"blog-ssg/internal/parser"
)

type Page struct {
	parser.Metadata
	Content template.HTML
}

func MakePosts(dir string) ([]string, error){
	var posts []string
	

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && info.Name() == "drafts" {
        	return filepath.SkipDir
   		}
		
		if !info.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			metadata, content, err := parser.ParseFrontMatter(content)
			htmlContent, err := parser.ConvertMarkdown(string(content))
			if err != nil {
				return err
			}
			page := Page{*metadata, template.HTML(htmlContent)}
			//title := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			postFileName, err := createPage(page)
			if err != nil {
				return err
		}

		posts = append(posts, postFileName)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func createPage(page Page) (string, error) {
	postFileDir := page.Metadata.Endpoint + "/index.html"
	path := filepath.Join("public",postFileDir)

	var tmpl *template.Template
	tmpl, err := tmpl.ParseGlob("templates/*html")
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, page.Template, page); err != nil {
		return "", err
	}

	if err := os.MkdirAll(filepath.Join("public", page.Metadata.Endpoint), 0755); err != nil {
		return "", err
	}

	if err := os.WriteFile(path, buf.Bytes(), 0644); err != nil {
		return "", err
	}
	
	return postFileDir, nil
}

