package faker

import (
	"reflect"
)

// Struct fills in exported elements of a struct with random data
// based on the value of `fake` tag of exported elements.
// Use `fake:"skip"` to explicitly skip an element.
// All built-in types are supported, with templating support
// for string types.
// Attention: Because it uses reflection it has a performance penalty!
func (f *Faker) Struct(v interface{}) {
	f.r(reflect.TypeOf(v), reflect.ValueOf(v), "")
}

func (f *Faker) r(t reflect.Type, v reflect.Value, template string) {
	switch t.Kind() {
	case reflect.Ptr:
		f.rPointer(t, v, template)
	case reflect.Struct:
		f.rStruct(t, v)
	case reflect.String:
		f.rString(template, v)
	case reflect.Uint8:
		v.SetUint(uint64(f.Uint8()))
	case reflect.Uint16:
		v.SetUint(uint64(f.Uint16()))
	case reflect.Uint32:
		v.SetUint(uint64(f.Uint32()))
	case reflect.Uint64:
		//capped at [0, math.MaxInt64)
		v.SetUint(f.Uint64())
	case reflect.Int:
		v.SetInt(f.Int64())
	case reflect.Int8:
		v.SetInt(int64(f.Int8()))
	case reflect.Int16:
		v.SetInt(int64(f.Int16()))
	case reflect.Int32:
		v.SetInt(int64(f.Int32()))
	case reflect.Int64:
		v.SetInt(f.Int64())
	case reflect.Float64:
		v.SetFloat(f.Float64())
	case reflect.Float32:
		v.SetFloat(float64(f.Float32()))
	case reflect.Bool:
		v.SetBool(f.Bool())
	}
}

func (f *Faker) rString(template string, v reflect.Value) {

	if template == "" {
		template = "{word} {word} {word} {word}"
	}
	r := f.Template(template)
	v.SetString(r)
}

func (f *Faker) rStruct(t reflect.Type, v reflect.Value) {
	n := t.NumField()
	for i := 0; i < n; i++ {
		elementT := t.Field(i)
		elementV := v.Field(i)
		fake := true
		t, ok := elementT.Tag.Lookup("fake")
		if ok && t == "skip" {
			fake = false
		}
		if fake && elementV.CanSet() {
			f.r(elementT.Type, elementV, t)
		}
	}
}

func (f *Faker) rPointer(t reflect.Type, v reflect.Value, template string) {
	elemT := t.Elem()
	if v.IsNil() {
		nv := reflect.New(elemT)
		f.r(elemT, nv.Elem(), template)
		v.Set(nv)
	} else {
		f.r(elemT, v.Elem(), template)
	}
}
