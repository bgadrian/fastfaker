package faker

import (
	"fmt"
	"strings"
	"testing"
)

func TestFaker_TemplateCustom(t *testing.T) {
	for _, test := range []struct {
		template, stardDelim, endDelim, should string
	}{
		{"%name%", "%", "%", "Jeromy Schmeler"},
		{"%%name%%", "%", "%", "%Jeromy Schmeler%"},
		{"%%name%%", "%%", "%%", "Jeromy Schmeler"},
		{"#name#", "#", "#", "Jeromy Schmeler"},
		{"##name##", "##", "##", "Jeromy Schmeler"},
		{"{{name}}", "{{", "}}", "Jeromy Schmeler"},
		{"{{{name}}}", "{{{", "}}}", "Jeromy Schmeler"},
		{"~name~", "~", "~", "Jeromy Schmeler"},
		{"~name}", "~", "}", "Jeromy Schmeler"},
		{"{{name}", "{{", "}", "Jeromy Schmeler"},
		{"<name>", "<", ">", "Jeromy Schmeler"},
	} {
		fastFaker := NewFastFaker()
		fastFaker.Seed(42)
		got, err := fastFaker.TemplateCustom(test.template, test.stardDelim, test.endDelim)
		if err != nil {
			t.Error(err)
		}
		if got == test.should {
			continue
		}

		t.Errorf("fot template '%s', got '%s' should '%s'",
			test.template, got, test.should)
	}
}
func TestFaker_TemplateCustomEdgeCases(t *testing.T) {
	fastFaker := NewFastFaker()

	res, err := fastFaker.TemplateCustom("{name}", "", "}")
	if err != nil {
		t.Error(err)
	}
	if res != "{name}" {
		t.Error("should have left the template unchanged with an empty delimiter")
	}

	res, err = fastFaker.TemplateCustom("|name|", "|", "|")
	if res != "|name|" {
		t.Error("should have left the template unchanged with an invalid delimiter")
	}
	if err == nil {
		t.Error("should have returned error with an invalid delimiter")
	}

	res, err = fastFaker.TemplateCustom("name", "", "")
	if err != nil {
		t.Error(err)
	}
	if res == "name" {
		t.Error("should work with a single variable")
	}
}

func TestFaker_TemplateCustomDelimiters(t *testing.T) {
	should := "😀o😀Jeromy Schmeler😀😀"
	for _, count := range []int{1, 2, 3, 4, 5, 10} {
		for _, delimRune := range TemplateAllowedDelimiters {
			delim := strings.Repeat(string(delimRune), count)
			template := fmt.Sprintf("😀o😀%sname%s😀😀", delim, delim)

			fastFaker := NewFastFaker()
			fastFaker.Seed(42)
			got, err := fastFaker.TemplateCustom(template, delim, delim)
			if err != nil {
				t.Error(err)
			}
			if got == should {
				continue
			}

			t.Errorf("fot template '%s', got '%s' should '%s'",
				template, got, should)
		}
	}
}

func TestFaker_Template(t *testing.T) {
	for _, test := range []struct {
		template, should string
	}{
		{"", ""},
		{"{", "{"},
		{"{{", "{{"},
		{"{}", "{}"},
		{"{notfound}", "{notfound}"},
		{"Alfa", "Alfa"},
		{"Al{fa", "Al{fa"},
		{"{name}", "Jeromy Schmeler"},
		{"{name}}", "Jeromy Schmeler}"},
		{"{{name}}", "{Jeromy Schmeler}"},
		{"{{name}", "{Jeromy Schmeler"},
		{"{name}!", "Jeromy Schmeler!"},
		{"X{name}X", "XJeromy SchmelerX"},
		{"X{name}X{name}X", "XJeromy SchmelerXKim SteuberX"},
		{"{name", "{name"},
		{"name}", "name}"},
		{"name{", "name{"},
	} {
		fastFaker := NewFastFaker()
		fastFaker.Seed(42)
		got := fastFaker.Template(test.template)
		if got == test.should {
			continue
		}

		t.Errorf("fot template '%s', got '%s' should '%s'",
			test.template, got, test.should)
	}
}

func TestFaker_TemplateJson(t *testing.T) {
	for _, test := range []struct {
		template, should string
	}{
		{`{name:"{name}"}`, `{name:"Jeromy Schmeler"}`},
		{`["{name}", "{name}"]`, `["Jeromy Schmeler", "Kim Steuber"]`},

		{`{name:"{name}", age: {digit}}`, `{name:"Jeromy Schmeler", age: 8}`},
	} {
		fastFaker := NewFastFaker()
		fastFaker.Seed(42)
		got := fastFaker.Template(test.template)

		if got == test.should {
			continue
		}

		t.Errorf("fot template '%s', got '%s' should '%s'",
			test.template, got, test.should)
	}
}

func ExampleFaker_Template() {
	template := "Hello {name}!"

	fastFaker := NewFastFaker() // not concurrent safe, see NewSafeFaker()
	fastFaker.Seed(42)          //for each seed value will generate a different result

	fmt.Printf("%s\n", fastFaker.Template(template))
	fmt.Printf("%s\n", fastFaker.Template(template))

	// Output: Hello Jeromy Schmeler!
	//Hello Kim Steuber!
}

func ExampleFaker_TemplateJSON() {
	template := `{name:"{name}", age: {digit}, confirmed: {booltext}}`

	fastFaker := NewFastFaker() // not concurrent safe, see NewSafeFaker()
	fastFaker.Seed(42)          //for each seed value will generate a different result

	fmt.Printf("%s\n", fastFaker.Template(template))
	// Output:{name:"Jeromy Schmeler", age: 8, confirmed: false}
}

func ExampleFaker_TemplateHTML() {
	template := `<ul class="person">
	<li>Name: {name}</li>
	<li>Age: ##</li>
	<li>Number: {phone}</li>
	<li>Address: {street}, {city} {country}</li>
</ul>`

	fastFaker := NewFastFaker() // not concurrent safe, see NewSafeFaker()
	fastFaker.Seed(42)          //for each seed value will generate a different result

	fmt.Printf("%s\n", fastFaker.Template(template))
	// Output: <ul class="person">
	//	<li>Name: Kim Steuber</li>
	//	<li>Age: 57</li>
	//	<li>Number: 3576839758</li>
	//	<li>Address: 21542 North Clubview, Schimmelborough Mozambique</li>
	//</ul>
}
func ExampleFaker_TemplateYAML() {
	template := `
invoice: ######
date: {year}-{month}-{day}
bill-to: {username}
	given: {name}
	city: {city}
	state: {state}
	postal: {zip}
product:
	-	sku: ??###?
		quantity: {uint8}
		description: {word}
		price: {uint8}.##
tax: {uint8}.##
total: {uint16}.##`

	fastFaker := NewFastFaker() // not concurrent safe, see NewSafeFaker()
	fastFaker.Seed(42)          //for each seed value will generate a different result

	fmt.Printf("%s\n", fastFaker.Template(template))
	// Output: invoice: 578035
	//date: 1910-September-2
	//bill-to: Little4541
	//	given: Lilly McClure
	//	city: Pacochaport
	//	state: Arizona
	//	postal: 59847
	//product:
	//	-	sku: vu768n
	//		quantity: 27
	//		description: hic
	//		price: 235.39
	//tax: 96.75
	//total: 49128.82
}

func TestFaker_TemplateVariables(t *testing.T) {
	fastFaker := NewFastFaker() // not concurrent safe, see NewSafeFaker()

	for variable := range templateVariables {
		template := fmt.Sprintf("{%s}", variable)
		if fastFaker.Template(template) == template {
			t.Errorf("failed for variable %s", variable)
		}
	}
}

func BenchmarkFaker_Template(b *testing.B) {
	template := "😀o😀{name} {word} {uint8}"
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.Template(template)
	}
}

func BenchmarkFaker_TemplateCustom(b *testing.B) {
	template := "😀o😀{name} {word} {uint8}"
	fastFaker := NewFastFaker()
	for i := 0; i < b.N; i++ {
		fastFaker.TemplateCustom(template, "{", "}")
	}
}
