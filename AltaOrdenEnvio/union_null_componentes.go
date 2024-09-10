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

type UnionNullComponentesTypeEnum int

const (
	UnionNullComponentesTypeEnumComponentes UnionNullComponentesTypeEnum = 1
)

type UnionNullComponentes struct {
	Null        *types.NullVal
	Componentes Componentes
	UnionType   UnionNullComponentesTypeEnum
}

func writeUnionNullComponentes(r *UnionNullComponentes, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullComponentesTypeEnumComponentes:
		return writeComponentes(r.Componentes, w)
	}
	return fmt.Errorf("invalid value for *UnionNullComponentes")
}

func NewUnionNullComponentes() *UnionNullComponentes {
	return &UnionNullComponentes{}
}

func (r *UnionNullComponentes) Serialize(w io.Writer) error {
	return writeUnionNullComponentes(r, w)
}

func DeserializeUnionNullComponentes(r io.Reader) (*UnionNullComponentes, error) {
	t := NewUnionNullComponentes()
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

func DeserializeUnionNullComponentesFromSchema(r io.Reader, schema string) (*UnionNullComponentes, error) {
	t := NewUnionNullComponentes()
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

func (r *UnionNullComponentes) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"numeroAgrupador\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"numerosComponentesHijos\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"referencias\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]}],\"name\":\"Componentes\",\"type\":\"record\"}]"
}

func (_ *UnionNullComponentes) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullComponentes) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullComponentes) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullComponentes) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullComponentes) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullComponentes) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullComponentes) SetLong(v int64) {

	r.UnionType = (UnionNullComponentesTypeEnum)(v)
}

func (r *UnionNullComponentes) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Componentes = NewComponentes()
		return &types.Record{Target: (&r.Componentes)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullComponentes) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullComponentes) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullComponentes) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullComponentes) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullComponentes) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullComponentes) Finalize()                        {}

func (r *UnionNullComponentes) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullComponentesTypeEnumComponentes:
		return json.Marshal(map[string]interface{}{"Andreani.AltaOrdenEnvio.Events.Common.Componentes": r.Componentes})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullComponentes")
}

func (r *UnionNullComponentes) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.AltaOrdenEnvio.Events.Common.Componentes"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Componentes)
	}
	return fmt.Errorf("invalid value for *UnionNullComponentes")
}