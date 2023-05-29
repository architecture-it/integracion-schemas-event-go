// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoDeAlmacenSCE.avsc
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

type UnionNullListaDeRemitosTypeEnum int

const (
	UnionNullListaDeRemitosTypeEnumListaDeRemitos UnionNullListaDeRemitosTypeEnum = 1
)

type UnionNullListaDeRemitos struct {
	Null           *types.NullVal
	ListaDeRemitos ListaDeRemitos
	UnionType      UnionNullListaDeRemitosTypeEnum
}

func writeUnionNullListaDeRemitos(r *UnionNullListaDeRemitos, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullListaDeRemitosTypeEnumListaDeRemitos:
		return writeListaDeRemitos(r.ListaDeRemitos, w)
	}
	return fmt.Errorf("invalid value for *UnionNullListaDeRemitos")
}

func NewUnionNullListaDeRemitos() *UnionNullListaDeRemitos {
	return &UnionNullListaDeRemitos{}
}

func (r *UnionNullListaDeRemitos) Serialize(w io.Writer) error {
	return writeUnionNullListaDeRemitos(r, w)
}

func DeserializeUnionNullListaDeRemitos(r io.Reader) (*UnionNullListaDeRemitos, error) {
	t := NewUnionNullListaDeRemitos()
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

func DeserializeUnionNullListaDeRemitosFromSchema(r io.Reader, schema string) (*UnionNullListaDeRemitos, error) {
	t := NewUnionNullListaDeRemitos()
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

func (r *UnionNullListaDeRemitos) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"documentos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"numeroEntidad\",\"type\":\"string\"},{\"name\":\"documento\",\"type\":\"string\"}],\"name\":\"Remito\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeRemitos\",\"type\":\"record\"}]"
}

func (_ *UnionNullListaDeRemitos) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullListaDeRemitos) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullListaDeRemitos) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullListaDeRemitos) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullListaDeRemitos) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullListaDeRemitos) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullListaDeRemitos) SetLong(v int64) {

	r.UnionType = (UnionNullListaDeRemitosTypeEnum)(v)
}

func (r *UnionNullListaDeRemitos) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.ListaDeRemitos = NewListaDeRemitos()
		return &types.Record{Target: (&r.ListaDeRemitos)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullListaDeRemitos) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullListaDeRemitos) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullListaDeRemitos) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullListaDeRemitos) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullListaDeRemitos) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullListaDeRemitos) Finalize()                        {}

func (r *UnionNullListaDeRemitos) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullListaDeRemitosTypeEnumListaDeRemitos:
		return json.Marshal(map[string]interface{}{"Andreani.Wapv2.Events.Record.ListaDeRemitos": r.ListaDeRemitos})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullListaDeRemitos")
}

func (r *UnionNullListaDeRemitos) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.Wapv2.Events.Record.ListaDeRemitos"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.ListaDeRemitos)
	}
	return fmt.Errorf("invalid value for *UnionNullListaDeRemitos")
}