// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ResponseGenerarPreenvios.avsc
 */
package PreEnvioBackendEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullDatosSucursalTypeEnum int

const (
	UnionNullDatosSucursalTypeEnumDatosSucursal UnionNullDatosSucursalTypeEnum = 1
)

type UnionNullDatosSucursal struct {
	Null          *types.NullVal
	DatosSucursal DatosSucursal
	UnionType     UnionNullDatosSucursalTypeEnum
}

func writeUnionNullDatosSucursal(r *UnionNullDatosSucursal, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullDatosSucursalTypeEnumDatosSucursal:
		return writeDatosSucursal(r.DatosSucursal, w)
	}
	return fmt.Errorf("invalid value for *UnionNullDatosSucursal")
}

func NewUnionNullDatosSucursal() *UnionNullDatosSucursal {
	return &UnionNullDatosSucursal{}
}

func (r *UnionNullDatosSucursal) Serialize(w io.Writer) error {
	return writeUnionNullDatosSucursal(r, w)
}

func DeserializeUnionNullDatosSucursal(r io.Reader) (*UnionNullDatosSucursal, error) {
	t := NewUnionNullDatosSucursal()
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

func DeserializeUnionNullDatosSucursalFromSchema(r io.Reader, schema string) (*UnionNullDatosSucursal, error) {
	t := NewUnionNullDatosSucursal()
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

func (r *UnionNullDatosSucursal) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Nomenclatura\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Descripcion\",\"type\":[\"null\",\"string\"]}],\"name\":\"DatosSucursal\",\"namespace\":\"Andreani.PreEnvioBackend.Events.Preenvio.Common\",\"type\":\"record\"}]"
}

func (_ *UnionNullDatosSucursal) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullDatosSucursal) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullDatosSucursal) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullDatosSucursal) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullDatosSucursal) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullDatosSucursal) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullDatosSucursal) SetLong(v int64) {

	r.UnionType = (UnionNullDatosSucursalTypeEnum)(v)
}

func (r *UnionNullDatosSucursal) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.DatosSucursal = NewDatosSucursal()
		return &types.Record{Target: (&r.DatosSucursal)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullDatosSucursal) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullDatosSucursal) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullDatosSucursal) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullDatosSucursal) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullDatosSucursal) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullDatosSucursal) Finalize()                        {}

func (r *UnionNullDatosSucursal) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullDatosSucursalTypeEnumDatosSucursal:
		return json.Marshal(map[string]interface{}{"Andreani.PreEnvioBackend.Events.Preenvio.Common.DatosSucursal": r.DatosSucursal})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullDatosSucursal")
}

func (r *UnionNullDatosSucursal) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.PreEnvioBackend.Events.Preenvio.Common.DatosSucursal"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.DatosSucursal)
	}
	return fmt.Errorf("invalid value for *UnionNullDatosSucursal")
}
