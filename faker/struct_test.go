package faker

import (
	"fmt"
	"testing"
)

type Basic struct {
	s string
	S string
}

type Nested struct {
	A   string
	B   *Basic
	bar *Basic
}

type BuiltIn struct {
	Uint8   *uint8
	Uint16  *uint16
	Uint32  *uint32
	Uint64  *uint64
	Int     *int
	Int8    *int8
	Int16   *int16
	Int32   *int32
	Int64   *int64
	Float32 *float32
	Float64 *float64
	Bool    *bool
}

type Template struct {
	Number *string `fake:"#"`
	Name   *string `fake:"{person.first}"`
}

func TestStructBasic(t *testing.T) {
	var basic Basic
	Global.Struct(&basic)
	if basic.s != "" {
		t.Error("unexported field is not populated")
	}
	if basic.S == "" {
		t.Error("exported field is populated")
	}
}

func TestStructNested(t *testing.T) {
	var nested Nested
	Global.Struct(&nested)
	if nested.A == "" {
		t.Error("exported string field is populated")
	}
	if nested.B == nil {
		t.Error("exported struct field is populated")
	}
	if nested.B.S == "" {
		t.Error("nested struct string field is not populated")
	}
	if nested.bar != nil {
		t.Error("nested struct bar should be be populated due to unexported variable")
	}
}

func TestStructBuiltInTypes(t *testing.T) {
	var builtIn BuiltIn
	Global.Struct(&builtIn)
	if builtIn.Uint8 == nil {
		t.Error("builtIn Uint8 was not set")
	}
	if builtIn.Uint16 == nil {
		t.Error("builtIn Uint16 was not set")
	}
	if builtIn.Uint32 == nil {
		t.Error("builtIn Uint32 was not set")
	}
	if builtIn.Uint64 == nil {
		t.Error("builtIn Uint64 was not set")
	}
	if builtIn.Int == nil {
		t.Error("builtIn int was not set")
	}
	if builtIn.Int8 == nil {
		t.Error("builtIn int8 was not set")
	}
	if builtIn.Int16 == nil {
		t.Error("builtIn int16 was not set")
	}
	if builtIn.Int32 == nil {
		t.Error("builtIn int32 was not set")
	}
	if builtIn.Int64 == nil {
		t.Error("builtIn int64 was not set")
	}
	if builtIn.Float32 == nil {
		t.Error("builtIn float32 was not set")
	}
	if builtIn.Float64 == nil {
		t.Error("builtIn float64 was not set")
	}
	if builtIn.Bool == nil {
		t.Error("builtIn bool was not set")
	}
}

func TestStructWithTemplate(t *testing.T) {
	var template Template
	Global.Struct(&template)
	if *template.Number == "" {
		t.Error("template number should set to number value")
	}
	if *template.Name == "" {
		t.Error("template number should set to number value")
	}
}

func ExampleFaker_Struct() {
	Global.Seed(42)
	type Foo struct {
		//for all the possible template variables see TEMPLATE_VARIABLES.md
		Browser string `fake:"{browser}"`
		Drink   string `fake:"{beername}"`
		Phone   string `fake:"##-###-###"`
		Blob    string `fake:"?????"`
		Int     int
		Dice    uint8
		Pointer *int
		Skip    *string `fake:"skip"`
	}
	var f Foo
	Global.Struct(&f)
	fmt.Printf("%s\n", f.Browser)
	fmt.Printf("%s\n", f.Drink)
	fmt.Printf("%s\n", f.Phone)
	fmt.Printf("%s\n", f.Blob)
	fmt.Printf("%d\n", f.Int)
	fmt.Printf("%d\n", f.Dice)
	fmt.Printf("%d\n", *f.Pointer)
	fmt.Printf("%v\n", f.Skip)
	// Output: firefox
	//Samuel Smith’s Oatmeal Stout
	//80-357-683
	//ptneu
	//-4923402592883905764
	//4
	//-2620303913143990366
	//<nil>
}
