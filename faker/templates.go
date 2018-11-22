package faker

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

// Template replaces all the found variables into the template with the actual
// results from the Faker.*() function.
//		fastFaker.Template("Hello {name}!") //Hello Jeromy Schmeler!
//
// All the "#" will be replaced by a digit and all the "?" by a ASCII letter.
//
// To use custom delimiters (instead of {}) see TemplateCustom()
// For more details and examples see TEMPLATES.md
func (f *Faker) Template(pattern string) string {

	//to allow simple patterns like phone numbers ##-###-###-###
	//and be backward compatibility with Generate()
	pattern = f.Numerify(pattern)
	pattern = f.Lexify(pattern)

	result, _ := f.TemplateCustom(pattern, "{", "}")
	return result
}

type keyPos struct {
	start, end   int
	variableFunc fakerer
}

// TemplateAllowedDelimiters the runes that are allowed as variable delimiters
// in Custom Templates. Must be ASCII (1 byte size) and not interfere with the regex expressions.
const TemplateAllowedDelimiters = "{}%#~<>-:@`"

//TODO group these 2 to a struct
var templateRegexCache = make(map[string]*regexp.Regexp)
var templateRegexMutex = sync.RWMutex{}

// TemplateCustom replaces all the found variables into the template with the actual
// results from the Faker.*() function.
//		fastFaker.Template("Hello {name}!") //Hello Jeromy Schmeler!
//
// It needs a start and end variable delimiters (delimStart, delimEnd).
// They can only contain runes from TemplateAllowedDelimiters.
// There is no constraint on the number of characters.
// Examples of valid delimiters: {{name}}, %%%name%%% or mixed: <<%name>>
//
// It uses Regex to find all the variables.
func (f *Faker) TemplateCustom(template, delimStart, delimEnd string) (string, error) {

	//edge case, the template is only a variable "name"
	if delimStart == "" || delimEnd == "" || len(template) < 3 {
		variableFunc, exists := templateVariables[template]
		if exists {
			return variableFunc(f), nil
		}
		return template, nil
	}

	for _, r := range delimStart + delimEnd {
		if strings.ContainsRune(TemplateAllowedDelimiters, r) == false {
			return template, fmt.Errorf("delimiters error, supported ones are: '%s'", TemplateAllowedDelimiters)
		}
	}

	//To allow more TemplateAllowedDelimiters we must add escape characters to each rune
	//for example for a delimiter "|||" we must transform it to "\|\|\|"

	//building a Regexp is pretty CPU and memory expensive, we reuse them
	//most likely all calls are using the same delimiters
	var pattern *regexp.Regexp
	cacheKey := delimStart + "///" + delimEnd

	templateRegexMutex.RLock()
	cachedRegex, exists := templateRegexCache[cacheKey]
	templateRegexMutex.RUnlock() //do not defer, we need it released ASAP

	if exists {
		pattern = cachedRegex
	} else {
		pattern = regexp.MustCompile(`(` + delimStart + `[a-zA-Z0-9_-]+` + delimEnd + `)`)

		templateRegexMutex.Lock()
		templateRegexCache[cacheKey] = pattern
		templateRegexMutex.Unlock()
	}

	templateAsByte := []byte(template)
	indexes := pattern.FindAllIndex(templateAsByte, -1)

	//filter templateVariables, we will find all the locations in the template
	//of all variables.
	var toReplace []keyPos
	for _, match := range indexes {
		start := match[0]
		end := match[1]

		variableWithDelim := string(templateAsByte[start:end])
		sizeDelimStart := len(delimStart)
		sizeDelimEnd := len(delimEnd)

		//remove the delimiters, eg: {name} => name
		variable := strings.ToLower(variableWithDelim[sizeDelimStart : len(variableWithDelim)-sizeDelimEnd])

		variableFunc, exists := templateVariables[variable]
		if exists == false {
			//the variable does not exists
			continue
		}
		toReplace = append(toReplace, keyPos{start, end, variableFunc})
	}

	if len(indexes) == 0 {
		return template, nil
	}

	//cannot use strings.Builder to keep 1.0 compatibility
	buff := bytes.Buffer{}
	buff.Grow(len(templateAsByte) + bytesPerWordEstimation*len(indexes)) //at least the input with some MagicNumber

	var lastEnd int
	//we go trough each byte and replace the variable with a value
	for _, posToReplace := range toReplace {
		buff.Write(templateAsByte[lastEnd:posToReplace.start])
		buff.WriteString(posToReplace.variableFunc(f))
		lastEnd = posToReplace.end
	}
	buff.Write(templateAsByte[lastEnd:])
	return buff.String(), nil
}
