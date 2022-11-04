// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TransactionEvent.avsc
 */
package AltaOrdenEnvioEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullArrayMetadatoTypeEnum int

const (
	UnionNullArrayMetadatoTypeEnumArrayMetadato UnionNullArrayMetadatoTypeEnum = 1
)

type UnionNullArrayMetadato struct {
	Null          *types.NullVal
	ArrayMetadato []Metadato
	UnionType     UnionNullArrayMetadatoTypeEnum
}

func writeUnionNullArrayMetadato(r *UnionNullArrayMetadato, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullArrayMetadatoTypeEnumArrayMetadato:
		return writeArrayMetadato(r.ArrayMetadato, w)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayMetadato")
}

func NewUnionNullArrayMetadato() *UnionNullArrayMetadato {
	return &UnionNullArrayMetadato{}
}

func (r *UnionNullArrayMetadato) Serialize(w io.Writer) error {
	return writeUnionNullArrayMetadato(r, w)
}

func DeserializeUnionNullArrayMetadato(r io.Reader) (*UnionNullArrayMetadato, error) {
	t := NewUnionNullArrayMetadato()
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

func DeserializeUnionNullArrayMetadatoFromSchema(r io.Reader, schema string) (*UnionNullArrayMetadato, error) {
	t := NewUnionNullArrayMetadato()
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

func (r *UnionNullArrayMetadato) Schema() string {
	return "[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]"
}

func (_ *UnionNullArrayMetadato) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullArrayMetadato) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullArrayMetadato) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullArrayMetadato) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullArrayMetadato) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullArrayMetadato) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullArrayMetadato) SetLong(v int64) {

	r.UnionType = (UnionNullArrayMetadatoTypeEnum)(v)
}

func (r *UnionNullArrayMetadato) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ArrayMetadato = make([]Metadato, 0)
		return &ArrayMetadatoWrapper{Target: (&r.ArrayMetadato)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullArrayMetadato) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullArrayMetadato) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullArrayMetadato) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullArrayMetadato) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullArrayMetadato) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullArrayMetadato) Finalize()                        {}

func (r *UnionNullArrayMetadato) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullArrayMetadatoTypeEnumArrayMetadato:
		return json.Marshal(map[string]interface{}{"array": r.ArrayMetadato})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullArrayMetadato")
}

func (r *UnionNullArrayMetadato) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["array"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ArrayMetadato)
	}
	return fmt.Errorf("invalid value for *UnionNullArrayMetadato")
}
