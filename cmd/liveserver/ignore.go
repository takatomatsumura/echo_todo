// based on github.com/sabhiram/go-gitignore but with some fix

package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type IgnorePattern struct {
	Pattern  *regexp.Regexp
	Negative bool
}

type Ignore struct {
	Patterns []*IgnorePattern
}

// compile subset of gitignore pattern format to regular expression.
func compilePattern(pattern string) (*regexp.Regexp, bool) {
	// [Rule 2] comment
	if pattern[0] == '#' {
		return nil, false
	}

	// [Rule 3] remove trailing spaces
	// TODO: treat escaped spaces
	pattern = strings.TrimRight(pattern, " ")

	// [Rule 1] blank line
	if pattern == "" {
		return nil, false
	}

	// [Rule 4] handle leading "!"
	negative := false
	if pattern[0] == '!' {
		negative = true
		pattern = pattern[1:]
	}

	// [Rule 2, 4] handle escaped leading '#' or '!'
	if strings.HasPrefix(pattern, `\#`) || strings.HasPrefix(pattern, `\!`) {
		pattern = pattern[1:]
	}

	// [Rule 6] a separator at middle of the pattern
	if regexp.MustCompile(`[^\/+]/.`).MatchString(pattern) && pattern[0] != '/' {
		pattern = "/" + pattern
	}

	// escape "."
	pattern = strings.ReplaceAll(pattern, ".", `\.`)

	magicStar := "#$~"

	// handle "/**/"
	if strings.HasPrefix(pattern, "/**/") {
		pattern = pattern[1:]
	}
	pattern = strings.ReplaceAll(pattern, "/**/", `(?:/|/.+/)`)
	pattern = strings.ReplaceAll(pattern, "**/", `(?:|.`+magicStar+`/)`)
	pattern = strings.ReplaceAll(pattern, "/**", `(?:|/.`+magicStar+`)`)

	// handle "*"
	pattern = strings.ReplaceAll(pattern, `\*`, `\`+magicStar)
	pattern = strings.ReplaceAll(pattern, "*", `[^/]*`)

	// handle "?"
	pattern = strings.ReplaceAll(pattern, "?", `[^/]`)

	pattern = strings.ReplaceAll(pattern, magicStar, "*")

	if pattern[len(pattern)-1] == '/' {
		pattern = pattern + ".*$"
	} else {
		pattern = pattern + "/?.*$"
	}
	if pattern[0] == '/' {
		pattern = "^/?" + pattern[1:]
	} else {
		pattern = "^(?:.*/)?" + pattern
	}

	return regexp.MustCompile(pattern), negative
}

func ParseIgnoreFile(path string) (*Ignore, error) {
	ignore := &Ignore{}

	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		ignore.AddPattern(line)
	}

	return ignore, nil
}

func (ignore *Ignore) AddPattern(pattern string) {
	ip, negative := compilePattern(pattern)
	if ip != nil {
		ignore.Patterns = append(ignore.Patterns, &IgnorePattern{
			Pattern:  ip,
			Negative: negative,
		})
	}
}

func (ignore *Ignore) AddPatterns(patterns ...string) {
	for _, pattern := range patterns {
		ignore.AddPattern(pattern)
	}
}

func (ignore *Ignore) Match(path string) bool {
	path = strings.ReplaceAll(path, string(os.PathSeparator), "/")

	match := false
	for _, p := range ignore.Patterns {
		if p.Pattern.MatchString(path) {
			if !p.Negative {
				match = true
			} else if match {
				match = false
			}
		}
	}
	return match
}
