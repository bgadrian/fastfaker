# FastFaker Templates

Template is the most powerful FastFaker feature. It allows custom patterns/templates (JSON, YAML, HTML, ...) of text to be filled with over 110 random types of [data (variables)](./TEMPLATE_VARIABLES.md).

It can be used directly (faker.Template* methods) or via the faker.Struct fill method and `fake:` tags. 

```go
//instead of:
fmt.Sprintf("I'm %s, call me at %s!", fastFaker.Name(), fastFaker.Numerify("###-###-####"))
//you can do:
fastFaker.Template("I'm {name}, call me at ###-###-####!") // I'm John, call me at 152-335-8761!
```

### Faker.Template()

The `Template()` function replaces:
* all the `#` to digits (calls `Faker.Numerify()`)
* all the `?` to ASCII lower case letters (calls `Faker.Lexify()`)
* all the [variables](./TEMPLATE_VARIABLES.md) with the apropiate function calls

```go
fastFaker.Template("I'm {name}, call me at ###-###-####!") // I'm John, call me at 152-335-8761!
```

You can use any kind of text, template or encoding like HTML or JSON:
```go
package main
import "fmt"
import "github.com/bgadrian/fastfaker/faker"

func main() {
	fastFaker := faker.NewSafeFaker()

	templateJSON := `{name:"{name}", age: {digit}}`
	fmt.Printf("%s\n", fastFaker.Template(templateJSON))
	// Output:{name:"Jeromy Schmeler", age: 8}

	templateHTML := `
            <ul class="person">
                <li>Name: {name}</li>
                <li>Age: ##</li>
                <li>Number: {phone}</li>
                <li>Address: {street}, {city} {country}</li>
            </ul>`
	
	
	fmt.Printf("%s\n", fastFaker.Template(templateHTML))
	// Output: <ul class="person">
	//	<li>Name: Kim Steuber</li>
	//	<li>Age: 57</li>
	//	<li>Number: 3576839758</li>
	//	<li>Address: 21542 North Clubview, Schimmelborough Mozambique</li>
	//</ul>
}
```

### Faker.TemplateCustom()

This function does NOT replace the `#` and `?`, but rather allows a more advanced usage of the Template feature. Instead of the default delimiters `{}`, you can use others: `{}%#~<>-:@`:

```go
    fastFaker.TemplateCustom("%buzzword%","%","%") //productivity
    fastFaker.TemplateCustom("<beername>","<",">") //Trappistes Rochefort 10
    fastFaker.TemplateCustom("~domainname~","~","~") //dynamicdistributed.info
    fastFaker.TemplateCustom("{{hexcolor%%","{{","%%") //#v08d8x
```

### Struct

The templates can be read from the strings Tag property of any struct and populated using the `Faker.Struct` method. It calls the `Faker.Template()` method for all `fake:""` tags.

```go
    type Foo struct {
            Browser string `fake:"{browser}"`
            Drink   string `fake:"{beername}"`
            Phone   string `fake:"##-###-###"`
            Blob    string `fake:"??????"`
    }
    bar := &Foo{}
    fastFaker.Struct(bar)
    
    fmt.Println(bar.Browser) //firefox 
    fmt.Println(bar.Browser) //Samuel Smith’s Oatmeal Stout
    fmt.Println(bar.Browser) //-3651589698752897048
    fmt.Println(bar.Blob)    //skrdcq
```

### Variables
Variables are aliases for the `Faker` public methods.

The templates recognize over 110 variables, you can see the [full list with examples here](./TEMPLATE_VARIABLES.md).