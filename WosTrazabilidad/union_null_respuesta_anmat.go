// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     WosTrazaAnmat.avsc
 */
package WosTrazabilidadEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullRespuestaAnmatTypeEnum int

const (
	UnionNullRespuestaAnmatTypeEnumRespuestaAnmat UnionNullRespuestaAnmatTypeEnum = 1
)

type UnionNullRespuestaAnmat struct {
	Null           *types.NullVal
	RespuestaAnmat RespuestaAnmat
	UnionType      UnionNullRespuestaAnmatTypeEnum
}

func writeUnionNullRespuestaAnmat(r *UnionNullRespuestaAnmat, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullRespuestaAnmatTypeEnumRespuestaAnmat:
		return writeRespuestaAnmat(r.RespuestaAnmat, w)
	}
	return fmt.Errorf("invalid value for *UnionNullRespuestaAnmat")
}

func NewUnionNullRespuestaAnmat() *UnionNullRespuestaAnmat {
	return &UnionNullRespuestaAnmat{}
}

func (r *UnionNullRespuestaAnmat) Serialize(w io.Writer) error {
	return writeUnionNullRespuestaAnmat(r, w)
}

func DeserializeUnionNullRespuestaAnmat(r io.Reader) (*UnionNullRespuestaAnmat, error) {
	t := NewUnionNullRespuestaAnmat()
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

func DeserializeUnionNullRespuestaAnmatFromSchema(r io.Reader, schema string) (*UnionNullRespuestaAnmat, error) {
	t := NewUnionNullRespuestaAnmat()
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

func (r *UnionNullRespuestaAnmat) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"codigoTransaccion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"resultado\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"errores\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]}],\"name\":\"RespuestaAnmat\",\"namespace\":\"Andreani.WosTrazabilidad.Events.AnmatCommon\",\"type\":\"record\"}]"
}

func (_ *UnionNullRespuestaAnmat) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullRespuestaAnmat) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullRespuestaAnmat) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullRespuestaAnmat) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullRespuestaAnmat) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullRespuestaAnmat) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullRespuestaAnmat) SetLong(v int64) {

	r.UnionType = (UnionNullRespuestaAnmatTypeEnum)(v)
}

func (r *UnionNullRespuestaAnmat) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.RespuestaAnmat = NewRespuestaAnmat()
		return &types.Record{Target: (&r.RespuestaAnmat)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullRespuestaAnmat) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullRespuestaAnmat) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullRespuestaAnmat) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullRespuestaAnmat) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullRespuestaAnmat) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullRespuestaAnmat) Finalize()                        {}

func (r *UnionNullRespuestaAnmat) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullRespuestaAnmatTypeEnumRespuestaAnmat:
		return json.Marshal(map[string]interface{}{"Andreani.WosTrazabilidad.Events.AnmatCommon.RespuestaAnmat": r.RespuestaAnmat})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullRespuestaAnmat")
}

func (r *UnionNullRespuestaAnmat) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.WosTrazabilidad.Events.AnmatCommon.RespuestaAnmat"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.RespuestaAnmat)
	}
	return fmt.Errorf("invalid value for *UnionNullRespuestaAnmat")
}
