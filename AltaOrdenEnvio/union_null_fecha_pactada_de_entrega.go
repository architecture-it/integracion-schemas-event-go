// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     OrdenDeEnvioSolicitada.avsc
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

type UnionNullFechaPactadaDeEntregaTypeEnum int

const (
	UnionNullFechaPactadaDeEntregaTypeEnumFechaPactadaDeEntrega UnionNullFechaPactadaDeEntregaTypeEnum = 1
)

type UnionNullFechaPactadaDeEntrega struct {
	Null                  *types.NullVal
	FechaPactadaDeEntrega FechaPactadaDeEntrega
	UnionType             UnionNullFechaPactadaDeEntregaTypeEnum
}

func writeUnionNullFechaPactadaDeEntrega(r *UnionNullFechaPactadaDeEntrega, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullFechaPactadaDeEntregaTypeEnumFechaPactadaDeEntrega:
		return writeFechaPactadaDeEntrega(r.FechaPactadaDeEntrega, w)
	}
	return fmt.Errorf("invalid value for *UnionNullFechaPactadaDeEntrega")
}

func NewUnionNullFechaPactadaDeEntrega() *UnionNullFechaPactadaDeEntrega {
	return &UnionNullFechaPactadaDeEntrega{}
}

func (r *UnionNullFechaPactadaDeEntrega) Serialize(w io.Writer) error {
	return writeUnionNullFechaPactadaDeEntrega(r, w)
}

func DeserializeUnionNullFechaPactadaDeEntrega(r io.Reader) (*UnionNullFechaPactadaDeEntrega, error) {
	t := NewUnionNullFechaPactadaDeEntrega()
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

func DeserializeUnionNullFechaPactadaDeEntregaFromSchema(r io.Reader, schema string) (*UnionNullFechaPactadaDeEntrega, error) {
	t := NewUnionNullFechaPactadaDeEntrega()
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

func (r *UnionNullFechaPactadaDeEntrega) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"fecha\",\"type\":[\"null\",\"string\"]},{\"name\":\"horaDesde\",\"type\":[\"null\",\"string\"]},{\"name\":\"horaHasta\",\"type\":[\"null\",\"string\"]}],\"name\":\"FechaPactadaDeEntrega\",\"type\":\"record\"}]"
}

func (_ *UnionNullFechaPactadaDeEntrega) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullFechaPactadaDeEntrega) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullFechaPactadaDeEntrega) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullFechaPactadaDeEntrega) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullFechaPactadaDeEntrega) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullFechaPactadaDeEntrega) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullFechaPactadaDeEntrega) SetLong(v int64) {

	r.UnionType = (UnionNullFechaPactadaDeEntregaTypeEnum)(v)
}

func (r *UnionNullFechaPactadaDeEntrega) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.FechaPactadaDeEntrega = NewFechaPactadaDeEntrega()
		return &types.Record{Target: (&r.FechaPactadaDeEntrega)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullFechaPactadaDeEntrega) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullFechaPactadaDeEntrega) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullFechaPactadaDeEntrega) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullFechaPactadaDeEntrega) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullFechaPactadaDeEntrega) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullFechaPactadaDeEntrega) Finalize()                {}

func (r *UnionNullFechaPactadaDeEntrega) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullFechaPactadaDeEntregaTypeEnumFechaPactadaDeEntrega:
		return json.Marshal(map[string]interface{}{"Andreani.AltaOrdenEnvio.Events.Common.FechaPactadaDeEntrega": r.FechaPactadaDeEntrega})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullFechaPactadaDeEntrega")
}

func (r *UnionNullFechaPactadaDeEntrega) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.AltaOrdenEnvio.Events.Common.FechaPactadaDeEntrega"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.FechaPactadaDeEntrega)
	}
	return fmt.Errorf("invalid value for *UnionNullFechaPactadaDeEntrega")
}
