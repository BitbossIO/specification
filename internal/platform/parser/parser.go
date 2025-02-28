package parser

import (
	"fmt"
	"html"
	"math"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func FetchFiles(srcPath, packageName, version string) []string {

	dir := filepath.FromSlash(srcPath + "/" + packageName + "/" + version)
	filenames := []string{}

	fn := func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, ".yaml") {
			return nil
		}

		// Do not go in to sub directories
		if filepath.Dir(path) != dir {
			return nil
		}

		filenames = append(filenames, path)

		return nil
	}

	if err := filepath.Walk(dir, fn); err != nil {
		panic(err)
	}

	return filenames
}

func reformat(s string, prefix string) string {
	parts := strings.Split(s, " ")

	lines := []string{}

	line := prefix

	for _, p := range parts {
		if len(p) == 0 {
			continue
		}

		if len(line)+len(p) > 74 {
			// line length exceeded. Add the line to our lines
			lines = append(lines, line)

			// start a new line
			line = prefix
		}

		// append the word to the line
		line = fmt.Sprintf("%v %v", line, p)
	}

	// make sure to append any remaining non-empty line
	if line != prefix {
		lines = append(lines, line)
	}

	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func SnakeCase(str string) string {
	result := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	result = matchAllCap.ReplaceAllString(result, "${1}_${2}")
	return strings.ToLower(result)
}

func KebabCase(str string) string {
	result := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	result = matchAllCap.ReplaceAllString(result, "${1}-${2}")
	return strings.ToLower(result)
}

func TemplateToFile(distPath string, data interface{}, inFile, outFile string) {
	f, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}

	tmplFuncs := template.FuncMap{
		"minus": func(a, b int) int {
			return a - b
		},
		"padding": func(str string, size int) string {
			return strings.Repeat(" ", int(math.Max(float64(size-len(str)), 0)))
		},
		"comment": func(str, chr string) string {
			return reformat(html.UnescapeString(str), chr)
		},
		"snakecase": func(str string) string {
			return SnakeCase(str)
		},
	}

	tmpl := template.Must(template.New(path.Base(inFile)).Funcs(tmplFuncs).ParseFiles(inFile))
	if err := tmpl.Execute(f, data); err != nil {
		panic(err)
	}
}
