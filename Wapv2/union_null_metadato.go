// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListaDeDetalleDeLinea.avsc
 */
package Wapv2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullMetadatoTypeEnum int

const (
	UnionNullMetadatoTypeEnumMetadato UnionNullMetadatoTypeEnum = 1
)

type UnionNullMetadato struct {
	Null      *types.NullVal
	Metadato  Metadato
	UnionType UnionNullMetadatoTypeEnum
}

func writeUnionNullMetadato(r *UnionNullMetadato, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullMetadatoTypeEnumMetadato:
		return writeMetadato(r.Metadato, w)
	}
	return fmt.Errorf("invalid value for *UnionNullMetadato")
}

func NewUnionNullMetadato() *UnionNullMetadato {
	return &UnionNullMetadato{}
}

func (r *UnionNullMetadato) Serialize(w io.Writer) error {
	return writeUnionNullMetadato(r, w)
}

func DeserializeUnionNullMetadato(r io.Reader) (*UnionNullMetadato, error) {
	t := NewUnionNullMetadato()
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

func DeserializeUnionNullMetadatoFromSchema(r io.Reader, schema string) (*UnionNullMetadato, error) {
	t := NewUnionNullMetadato()
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

func (r *UnionNullMetadato) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"}]"
}

func (_ *UnionNullMetadato) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullMetadato) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullMetadato) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullMetadato) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullMetadato) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullMetadato) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullMetadato) SetLong(v int64) {

	r.UnionType = (UnionNullMetadatoTypeEnum)(v)
}

func (r *UnionNullMetadato) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Metadato = NewMetadato()
		return &types.Record{Target: (&r.Metadato)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullMetadato) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullMetadato) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullMetadato) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullMetadato) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullMetadato) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullMetadato) Finalize()                        {}

func (r *UnionNullMetadato) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullMetadatoTypeEnumMetadato:
		return json.Marshal(map[string]interface{}{"Andreani.Wapv2.Events.Record.Metadato": r.Metadato})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullMetadato")
}

func (r *UnionNullMetadato) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.Wapv2.Events.Record.Metadato"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Metadato)
	}
	return fmt.Errorf("invalid value for *UnionNullMetadato")
}
