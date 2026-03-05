package main

import (
	"os"
	"path/filepath"
	"testing"
)

var testDomains = map[string]bool{
	"engineering":  true,
	"architecture": true,
	"product":      true,
}

func writeTempFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	return path
}

func TestValidateFile_Valid(t *testing.T) {
	dir := t.TempDir()
	path := writeTempFile(t, dir, "my-skill.md", `---
id: my-skill
type: skill
domain: engineering
---

# My Skill
`)
	errs := validateFile(path, "skill", testDomains)
	if len(errs) != 0 {
		t.Errorf("expected no errors, got: %v", errs)
	}
}

func TestValidateFile_NoFrontmatter(t *testing.T) {
	dir := t.TempDir()
	path := writeTempFile(t, dir, "my-skill.md", `# My Skill

no frontmatter here
`)
	errs := validateFile(path, "skill", testDomains)
	if len(errs) != 1 || errs[0].Message != "frontmatter (---) not found" {
		t.Errorf("expected frontmatter error, got: %v", errs)
	}
}

func TestValidateFile_MissingID(t *testing.T) {
	dir := t.TempDir()
	path := writeTempFile(t, dir, "my-skill.md", `---
type: skill
domain: engineering
---
`)
	errs := validateFile(path, "skill", testDomains)
	assertContainsMessage(t, errs, "required field 'id' is missing or empty")
}

func TestValidateFile_MissingType(t *testing.T) {
	dir := t.TempDir()
	path := writeTempFile(t, dir, "my-skill.md", `---
id: my-skill
domain: engineering
---
`)
	errs := validateFile(path, "skill", testDomains)
	assertContainsMessage(t, errs, "required field 'type' is missing or empty")
}

func TestValidateFile_MissingDomain(t *testing.T) {
	dir := t.TempDir()
	path := writeTempFile(t, dir, "my-skill.md", `---
id: my-skill
type: skill
---
`)
	errs := validateFile(path, "skill", testDomains)
	assertContainsMessage(t, errs, "required field 'domain' is missing or empty")
}

func TestValidateFile_InvalidType(t *testing.T) {
	dir := t.TempDir()
	path := writeTempFile(t, dir, "my-skill.md", `---
id: my-skill
type: unknown
domain: engineering
---
`)
	errs := validateFile(path, "skill", testDomains)
	assertContainsMessagePrefix(t, errs, "invalid type value 'unknown'")
}

func TestValidateFile_TypeMismatchDirectory(t *testing.T) {
	dir := t.TempDir()
	// type: persona が skills/ ディレクトリに置かれているケース
	path := writeTempFile(t, dir, "my-skill.md", `---
id: my-skill
type: persona
domain: engineering
---
`)
	errs := validateFile(path, "skill", testDomains)
	assertContainsMessagePrefix(t, errs, "type 'persona' is not allowed in this directory")
}

func TestValidateFile_InvalidDomain(t *testing.T) {
	dir := t.TempDir()
	path := writeTempFile(t, dir, "my-skill.md", `---
id: my-skill
type: skill
domain: invalid-domain
---
`)
	errs := validateFile(path, "skill", testDomains)
	assertContainsMessagePrefix(t, errs, "invalid domain value 'invalid-domain'")
}

func TestValidateFile_IDFilenameMismatch(t *testing.T) {
	dir := t.TempDir()
	// ファイル名は my-skill.md だが id は other-id
	path := writeTempFile(t, dir, "my-skill.md", `---
id: other-id
type: skill
domain: engineering
---
`)
	errs := validateFile(path, "skill", testDomains)
	assertContainsMessagePrefix(t, errs, "id 'other-id' does not match filename 'my-skill'")
}

func TestValidateFile_MultipleErrors(t *testing.T) {
	dir := t.TempDir()
	// id・type・domain すべて欠如
	path := writeTempFile(t, dir, "my-skill.md", `---
sources: []
---
`)
	errs := validateFile(path, "skill", testDomains)
	if len(errs) < 3 {
		t.Errorf("expected at least 3 errors, got %d: %v", len(errs), errs)
	}
}

func TestLoadConfig(t *testing.T) {
	dir := t.TempDir()
	configContent := `domains:
  - engineering
  - architecture
  - product
`
	if err := os.WriteFile(filepath.Join(dir, "config.yaml"), []byte(configContent), 0644); err != nil {
		t.Fatalf("failed to write config.yaml: %v", err)
	}

	domains, err := loadConfig(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"engineering", "architecture", "product"}
	for _, d := range expected {
		if !domains[d] {
			t.Errorf("expected domain '%s' to be present", d)
		}
	}
	if len(domains) != len(expected) {
		t.Errorf("expected %d domains, got %d", len(expected), len(domains))
	}
}

func TestLoadConfig_Missing(t *testing.T) {
	_, err := loadConfig("/nonexistent/path")
	if err == nil {
		t.Error("expected error for missing config.yaml, got nil")
	}
}

// helpers

func assertContainsMessage(t *testing.T, errs []ValidationError, msg string) {
	t.Helper()
	for _, e := range errs {
		if e.Message == msg {
			return
		}
	}
	t.Errorf("expected error message %q not found in: %v", msg, errs)
}

func assertContainsMessagePrefix(t *testing.T, errs []ValidationError, prefix string) {
	t.Helper()
	for _, e := range errs {
		if len(e.Message) >= len(prefix) && e.Message[:len(prefix)] == prefix {
			return
		}
	}
	t.Errorf("expected error message with prefix %q not found in: %v", prefix, errs)
}
