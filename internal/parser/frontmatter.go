//YAML frontmatter parsing

package parser

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

type Metadata struct {
	Title       string   `yaml:"title"`
	Endpoint    string   `yaml:"endpoint"`
	Date        string   `yaml:"date"`
	Tags        []string `yaml:"tags"`
	Description string   `yaml:"description"`
	Template    string   `yaml:"template"`
}

func ParseFrontMatter(content []byte) (*Metadata, []byte, error) {
	parts := bytes.Split(content, []byte("---"))
	if len(parts) < 3 {
		return nil, content, nil
	}

	var metadata Metadata
	err := yaml.Unmarshal(parts[1], &metadata)
	if err != nil {
		return nil, content, err
	}

	return &metadata, bytes.Join(parts[2:], []byte("---")), nil
}
