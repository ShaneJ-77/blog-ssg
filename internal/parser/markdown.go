//markdown processing

package parser

import (
	"bytes"

	"github.com/yuin/goldmark"
)

func ConvertMarkdown(content string) (string, error) {
	var buf bytes.Buffer
	md := goldmark.New()
	if err := md.Convert([]byte(content), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}