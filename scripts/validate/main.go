package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// ディレクトリ名 -> 許可される type 値のマッピング
var dirTypeMap = map[string]string{
	"agents":    "persona",
	"skills":    "skill",
	"reviews":   "review",
	"artifacts": "artifact",
}

var validTypes = map[string]bool{
	"persona":  true,
	"skill":    true,
	"review":   true,
	"artifact": true,
}

var frontmatterRe = regexp.MustCompile(`(?s)^---\n(.+?)\n---`)

type Frontmatter struct {
	ID     string `yaml:"id"`
	Type   string `yaml:"type"`
	Domain string `yaml:"domain"`
}

type ValidationError struct {
	FilePath string
	Message  string
}

func (e ValidationError) String() string {
	return fmt.Sprintf("[ERROR] %s: %s", e.FilePath, e.Message)
}

func main() {
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	var errors []ValidationError

	for dirName, expectedType := range dirTypeMap {
		dirPath := filepath.Join(root, dirName)

		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			continue
		}

		err := filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() || !strings.HasSuffix(path, ".md") {
				return nil
			}

			errs := validateFile(path, expectedType)
			errors = append(errors, errs...)
			return nil
		})

		if err != nil {
			fmt.Fprintf(os.Stderr, "walk error: %v\n", err)
			os.Exit(1)
		}
	}

	if len(errors) > 0 {
		for _, e := range errors {
			fmt.Println(e)
		}
		fmt.Printf("\n%d error(s) found.\n", len(errors))
		os.Exit(1)
	}

	fmt.Println("All files are valid.")
}

func validateFile(path, expectedType string) []ValidationError {
	var errors []ValidationError

	content, err := os.ReadFile(path)
	if err != nil {
		return append(errors, ValidationError{FilePath: path, Message: fmt.Sprintf("failed to read file: %v", err)})
	}

	matches := frontmatterRe.FindSubmatch(content)
	if matches == nil {
		return append(errors, ValidationError{FilePath: path, Message: "frontmatter (---) not found"})
	}

	var fm Frontmatter
	if err := yaml.Unmarshal(matches[1], &fm); err != nil {
		return append(errors, ValidationError{FilePath: path, Message: fmt.Sprintf("failed to parse frontmatter YAML: %v", err)})
	}

	// 必須フィールドチェック
	if fm.ID == "" {
		errors = append(errors, ValidationError{FilePath: path, Message: "required field 'id' is missing or empty"})
	}
	if fm.Type == "" {
		errors = append(errors, ValidationError{FilePath: path, Message: "required field 'type' is missing or empty"})
	}
	if fm.Domain == "" {
		errors = append(errors, ValidationError{FilePath: path, Message: "required field 'domain' is missing or empty"})
	}

	// type 値チェック
	if fm.Type != "" && !validTypes[fm.Type] {
		errors = append(errors, ValidationError{FilePath: path, Message: fmt.Sprintf("invalid type value '%s': must be one of persona, skill, review, artifact", fm.Type)})
	}

	// ディレクトリ↔type 対応チェック
	if fm.Type != "" && validTypes[fm.Type] && fm.Type != expectedType {
		errors = append(errors, ValidationError{FilePath: path, Message: fmt.Sprintf("type '%s' is not allowed in this directory (expected '%s')", fm.Type, expectedType)})
	}

	// id とファイル名の一致チェック
	if fm.ID != "" {
		basename := strings.TrimSuffix(filepath.Base(path), ".md")
		if fm.ID != basename {
			errors = append(errors, ValidationError{FilePath: path, Message: fmt.Sprintf("id '%s' does not match filename '%s'", fm.ID, basename)})
		}
	}

	return errors
}
