package slob

import "reflect"

type input interface {
	Read(...interface{})
	Render(*render)
}

type StructInput struct {
	structSet []*Struct
}

type Struct struct {
	Name   string
	Fields []*Field
}

type Field struct {
	Name      string
	Type      string
	TransName string
}

var skipFieldsList = map[string]int{
	"CreateAt": 1,
	"UpdateAt": 1,
	"DeleteAt": 1,
}

func SkipFields(name string) bool {
	_, ok := skipFieldsList[name]
	return ok
}

func TransName(name string) (transStr string) {
	newName := make([]rune, 0)
	for idx, chr := range name {
		if 'A' <= chr && 'Z' >= chr {
			if idx > 0 {
				newName = append(newName, '_')
			}
			chr -= ('A' - 'a')
		}
		newName = append(newName, chr)
	}
	return string(newName)
}

func (i *StructInput) Read(model ...interface{}) {
	i.structSet = make([]*Struct, 0)
	for _, obj := range model {
		st := new(Struct)
		typ := reflect.TypeOf(obj).Elem()
		st.Name = typ.Name()
		tmp := typ.NumField()
		fields := make([]*Field, 0)
		for i := 0; i < tmp; i++ {
			field := typ.Field(i)
			fld := new(Field)
			if SkipFields(field.Name) {
				continue
			}
			fld.Name = field.Name
			fld.Type = field.Type.Name()
			fld.TransName = TransName(fld.Name)
			fields = append(fields, fld)
		}
		st.Fields = fields
		i.structSet = append(i.structSet, st)
	}

}

func (i *StructInput) Render(r *render) {
	for _, obj := range i.structSet {
		fileName := TransName(obj.Name)
		r.SetFileName(fileName).Set("struct", obj).execute()
	}
}

var slob_input input = new(StructInput)

func SetInput(inp input) {
	slob_input = inp
}
