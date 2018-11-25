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

	//nolint template cannot fail with {}, see tests
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

//because regex creation is expensive we cache them
//we presume that it will be a Read heavy usage, as in
//the same delimiters `{}` will be used for many requests.
type regexCache struct {
	cache map[string]*regexp.Regexp
	sync.RWMutex
}

var templateRegexMutex = regexCache{
	cache: make(map[string]*regexp.Regexp),
}

var templateBuffersPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var templateKeyPosPool = sync.Pool{
	New: func() interface{} {
		return &keyPos{}
	},
}

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
		if !strings.ContainsRune(TemplateAllowedDelimiters, r) {
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
	cachedRegex, exists := templateRegexMutex.cache[cacheKey]
	templateRegexMutex.RUnlock() //do not defer, we need it released ASAP

	if exists {
		pattern = cachedRegex
	} else {
		pattern = regexp.MustCompile(`(` + delimStart + `[a-zA-Z0-9_-]+` + delimEnd + `)`)

		templateRegexMutex.Lock()
		templateRegexMutex.cache[cacheKey] = pattern
		templateRegexMutex.Unlock()
	}

	templateAsByte := []byte(template)
	indexes := pattern.FindAllIndex(templateAsByte, -1)

	//filter templateVariables, we will find all the locations in the template
	//of all variables.
	toReplace := make([]*keyPos, 0, len(indexes))
	for _, match := range indexes {
		pos := templateKeyPosPool.Get().(*keyPos)
		pos.start = match[0]
		pos.end = match[1]

		variableWithDelim := string(templateAsByte[pos.start:pos.end])
		sizeDelimStart := len(delimStart)
		sizeDelimEnd := len(delimEnd)

		//remove the delimiters, eg: {name} => name
		variable := strings.ToLower(variableWithDelim[sizeDelimStart : len(variableWithDelim)-sizeDelimEnd])

		pos.variableFunc, exists = templateVariables[variable]
		if !exists {
			//the variable does not exists
			templateKeyPosPool.Put(pos)
			continue
		}

		toReplace = append(toReplace, pos)
	}

	defer func() {
		//return all keyPos to the pool
		for _, pos := range toReplace {
			templateKeyPosPool.Put(pos)
		}
		toReplace = nil
	}()

	if len(indexes) == 0 {
		return template, nil
	}

	//cannot use strings.Builder to keep Go 1.0 compatibility
	buff := templateBuffersPool.Get().(*bytes.Buffer)
	buff.Reset()

	//we reuse the buffer instances, and their inner slices
	//shrink/grows without new allocations, if there is enough room
	//we end up reusing the same slices over and over
	//If the templates are very similar, the performance increases
	buff.Grow(len(templateAsByte) + bytesPerWordEstimation*len(indexes)) //at least the input with some MagicNumber

	var lastEnd int
	//we go trough each byte and replace the variable with a value
	for _, posToReplace := range toReplace {
		//nolint buffer never returns error
		buff.Write(templateAsByte[lastEnd:posToReplace.start])
		buff.WriteString(posToReplace.variableFunc(f))
		//nolint buffer never returns error
		lastEnd = posToReplace.end
	}
	//nolint buffer never returns error
	buff.Write(templateAsByte[lastEnd:])
	result := buff.String()

	templateBuffersPool.Put(buff)
	return result, nil
}
