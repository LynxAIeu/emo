package core

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"unicode"
)

type Ref struct {
	Name    string `json:"name"`
	Emoji   string `json:"emoji"`
	IsError bool   `json:"isError"`
}

var doNotAppendError = []string{"error", "warn", "info", "debug", "panic", "ok", "found"}

func containsInfoWarnError(name string) bool {
	name = strings.ToLower(name)
	for token := range slices.Values(doNotAppendError) {
		if strings.Contains(name, token) {
			return true
		}
	}
	return false
}

func Sanitize(ref []Ref) []Ref {
	names := make(map[string]struct{}, len(ref))
	for item := range slices.Values(ref) {
		_, found := names[item.Name]
		if found {
			log.Panic("Duplicated name=", item.Name, " in ")
		}
		names[item.Name] = struct{}{}
	}

	// add error functions
	for item := range slices.Values(ref) {
		// skip if function name already contains any of: error, warn, info, debug...
		if containsInfoWarnError(item.Name) {
			continue
		}

		errorFunction := item.Name + "Error"

		// skip if error functions already present
		_, found := names[errorFunction]
		if found {
			continue
		}

		ref = append(ref, Ref{errorFunction, item.Emoji, true})
	}

	slices.SortFunc(ref, func(a, b Ref) int { return strings.Compare(a.Name, b.Name) })

	return ref
}

func GetRef() []Ref {
	fn, err := filepath.Abs("./codegen/ref.json")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("[codegen] Open referential: ", fn)

	b, err := os.ReadFile(fn)
	if err != nil {
		log.Panic(err)
	}

	ref := []Ref{}
	err = json.Unmarshal(b, &ref)
	if err != nil {
		log.Panic(err)
	}

	return ref
}

func Write(fn, code string) {
	fn, err := filepath.Abs(fn)
	if err != nil {
		log.Panic(err)
	}

	file, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o755)
	if err != nil {
		log.Panic(err)
	}

	n, err := file.Write([]byte(code))
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("[codegen] File: %s (%d bytes)"+"\n", fn, n)
}

func Uncapitalized(s string) string {
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

func SnakeCase(s string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
