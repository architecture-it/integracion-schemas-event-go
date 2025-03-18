// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnviosEvent.avsc
 */
package CartaDocumentoBackendEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullArrayPaqueteTypeEnum int

const (
	UnionNullArrayPaqueteTypeEnumArrayPaquete UnionNullArrayPaqueteTypeEnum = 1
)

type UnionNullArrayPaquete struct {
	Null         *types.NullVal
	ArrayPaquete []Paquete
	UnionType    UnionNullArrayPaqueteTypeEnum
}

func writeUnionNullArrayPaquete(r *UnionNullArrayPaquete, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayPaqueteTypeEnumArrayPaquete:
		return writeArrayPaquete(r.ArrayPaquete, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayPaquete")
}

func NewUnionNullArrayPaquete() *UnionNullArrayPaquete {
	return &UnionNullArrayPaquete{}
}

func (r *UnionNullArrayPaquete) Serialize(w io.Writer) error {
	return writeUnionNullArrayPaquete(r, w)
}

func DeserializeUnionNullArrayPaquete(r io.Reader) (*UnionNullArrayPaquete, error) {
	t := NewUnionNullArrayPaquete()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionNullArrayPaqueteFromSchema(r io.Reader, schema string) (*UnionNullArrayPaquete, error) {
	t := NewUnionNullArrayPaquete()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r *UnionNullArrayPaquete) Schema() string {
	return "[\"null\",{\"items\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"TipoId\",\"type\":\"string\"},{\"name\":\"Tipo\",\"type\":\"string\"},{\"name\":\"Alto\",\"type\":\"string\"},{\"name\":\"Ancho\",\"type\":\"string\"},{\"name\":\"Largo\",\"type\":\"string\"},{\"name\":\"Peso\",\"type\":\"string\"},{\"name\":\"ValorDeclarado\",\"type\":\"int\"},{\"default\":null,\"name\":\"NumeroDeBulto\",\"type\":[\"null\",\"string\"]}],\"name\":\"Paquete\",\"type\":\"record\"},\"type\":\"array\"}]"
}

func (_ *UnionNullArrayPaquete) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullArrayPaquete) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullArrayPaquete) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullArrayPaquete) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullArrayPaquete) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullArrayPaquete) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullArrayPaquete) SetLong(v int64) {

	r.UnionType = (UnionNullArrayPaqueteTypeEnum)(v)
}

func (r *UnionNullArrayPaquete) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayPaquete = make([]Paquete, 0)
		return &ArrayPaqueteWrapper{Target: (&r.ArrayPaquete)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayPaquete) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullArrayPaquete) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullArrayPaquete) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullArrayPaquete) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullArrayPaquete) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullArrayPaquete) Finalize()                        {}

func (r *UnionNullArrayPaquete) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayPaqueteTypeEnumArrayPaquete:
		return json.Marshal(map[string]interface{}{"array": r.ArrayPaquete})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayPaquete")
}

func (r *UnionNullArrayPaquete) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayPaquete)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayPaquete")
}
