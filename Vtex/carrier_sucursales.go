// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     VtexSucursalesCreationNotificationEvent.avsc
 */
package VtexEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type CarrierSucursales struct {
	CarrierType CarrierType `json:"CarrierType"`

	SucursalesIds []int32 `json:"SucursalesIds"`
}

const CarrierSucursalesAvroCRC64Fingerprint = "\x89N\xf2\xe4\x88-\xee\xbc"

func NewCarrierSucursales() CarrierSucursales {
	r := CarrierSucursales{}
	r.SucursalesIds = make([]int32, 0)

	return r
}

func DeserializeCarrierSucursales(r io.Reader) (CarrierSucursales, error) {
	t := NewCarrierSucursales()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCarrierSucursalesFromSchema(r io.Reader, schema string) (CarrierSucursales, error) {
	t := NewCarrierSucursales()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCarrierSucursales(r CarrierSucursales, w io.Writer) error {
	var err error
	err = writeCarrierType(r.CarrierType, w)
	if err != nil {
		return err
	}
	err = writeArrayInt(r.SucursalesIds, w)
	if err != nil {
		return err
	}
	return err
}

func (r CarrierSucursales) Serialize(w io.Writer) error {
	return writeCarrierSucursales(r, w)
}

func (r CarrierSucursales) Schema() string {
	return "{\"fields\":[{\"name\":\"CarrierType\",\"type\":{\"name\":\"CarrierType\",\"symbols\":[\"AndreaniSucursal\",\"AndreaniBigger\"],\"type\":\"enum\"}},{\"name\":\"SucursalesIds\",\"type\":{\"items\":\"int\",\"type\":\"array\"}}],\"name\":\"Andreani.Vtex.Events.Record.VtexSucursalesCreationNotification.CarrierSucursales\",\"type\":\"record\"}"
}

func (r CarrierSucursales) SchemaName() string {
	return "Andreani.Vtex.Events.Record.VtexSucursalesCreationNotification.CarrierSucursales"
}

func (_ CarrierSucursales) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CarrierSucursales) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CarrierSucursales) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CarrierSucursales) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CarrierSucursales) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CarrierSucursales) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CarrierSucursales) SetString(v string)   { panic("Unsupported operation") }
func (_ CarrierSucursales) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CarrierSucursales) Get(i int) types.Field {
	switch i {
	case 0:
		w := CarrierTypeWrapper{Target: &r.CarrierType}

		return w

	case 1:
		r.SucursalesIds = make([]int32, 0)

		w := ArrayIntWrapper{Target: &r.SucursalesIds}

		return w

	}
	panic("Unknown field index")
}

func (r *CarrierSucursales) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *CarrierSucursales) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ CarrierSucursales) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CarrierSucursales) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CarrierSucursales) HintSize(int)                     { panic("Unsupported operation") }
func (_ CarrierSucursales) Finalize()                        {}

func (_ CarrierSucursales) AvroCRC64Fingerprint() []byte {
	return []byte(CarrierSucursalesAvroCRC64Fingerprint)
}

func (r CarrierSucursales) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["CarrierType"], err = json.Marshal(r.CarrierType)
	if err != nil {
		return nil, err
	}
	output["SucursalesIds"], err = json.Marshal(r.SucursalesIds)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CarrierSucursales) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["CarrierType"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CarrierType); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CarrierType")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SucursalesIds"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalesIds); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SucursalesIds")
	}
	return nil
}
