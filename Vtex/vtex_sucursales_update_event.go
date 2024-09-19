// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     VtexSucursalesUpdateEvent.avsc
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

type VtexSucursalesUpdateEvent struct {
	User UserInfo `json:"User"`

	CarrierType CarrierType `json:"CarrierType"`
}

const VtexSucursalesUpdateEventAvroCRC64Fingerprint = "\x1cʻ\xf2=\xdf\xdcP"

func NewVtexSucursalesUpdateEvent() VtexSucursalesUpdateEvent {
	r := VtexSucursalesUpdateEvent{}
	r.User = NewUserInfo()

	return r
}

func DeserializeVtexSucursalesUpdateEvent(r io.Reader) (VtexSucursalesUpdateEvent, error) {
	t := NewVtexSucursalesUpdateEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeVtexSucursalesUpdateEventFromSchema(r io.Reader, schema string) (VtexSucursalesUpdateEvent, error) {
	t := NewVtexSucursalesUpdateEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeVtexSucursalesUpdateEvent(r VtexSucursalesUpdateEvent, w io.Writer) error {
	var err error
	err = writeUserInfo(r.User, w)
	if err != nil {
		return err
	}
	err = writeCarrierType(r.CarrierType, w)
	if err != nil {
		return err
	}
	return err
}

func (r VtexSucursalesUpdateEvent) Serialize(w io.Writer) error {
	return writeVtexSucursalesUpdateEvent(r, w)
}

func (r VtexSucursalesUpdateEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"User\",\"type\":{\"fields\":[{\"name\":\"Email\",\"type\":\"string\"},{\"name\":\"VtexURL\",\"type\":\"string\"},{\"name\":\"VtexAppKey\",\"type\":\"string\"},{\"name\":\"VtexAppToken\",\"type\":\"string\"},{\"name\":\"Usuario_login\",\"type\":\"long\"},{\"name\":\"Aol_id\",\"type\":\"long\"},{\"name\":\"Hostname\",\"type\":\"string\"},{\"name\":\"Contratos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"name\":\"TipoDeEnvio\",\"type\":\"string\"},{\"name\":\"ModoDeEntrega\",\"type\":\"string\"}],\"name\":\"Contrato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"UserInfo\",\"namespace\":\"Andreani.Vtex.Events.Record.VtexSucursalesOnboarding\",\"type\":\"record\"}},{\"name\":\"CarrierType\",\"type\":{\"name\":\"CarrierType\",\"namespace\":\"Andreani.Vtex.Events.Record.VtexSucursalesCreationNotification\",\"symbols\":[\"AndreaniSucursal\",\"AndreaniBigger\"],\"type\":\"enum\"}}],\"name\":\"Andreani.Vtex.Events.Record.VtexSucursalesUpdateEvent\",\"type\":\"record\"}"
}

func (r VtexSucursalesUpdateEvent) SchemaName() string {
	return "Andreani.Vtex.Events.Record.VtexSucursalesUpdateEvent"
}

func (_ VtexSucursalesUpdateEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ VtexSucursalesUpdateEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ VtexSucursalesUpdateEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ VtexSucursalesUpdateEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ VtexSucursalesUpdateEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ VtexSucursalesUpdateEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ VtexSucursalesUpdateEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ VtexSucursalesUpdateEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *VtexSucursalesUpdateEvent) Get(i int) types.Field {
	switch i {
	case 0:
		r.User = NewUserInfo()

		w := types.Record{Target: &r.User}

		return w

	case 1:
		w := CarrierTypeWrapper{Target: &r.CarrierType}

		return w

	}
	panic("Unknown field index")
}

func (r *VtexSucursalesUpdateEvent) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *VtexSucursalesUpdateEvent) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ VtexSucursalesUpdateEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ VtexSucursalesUpdateEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ VtexSucursalesUpdateEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ VtexSucursalesUpdateEvent) Finalize()                        {}

func (_ VtexSucursalesUpdateEvent) AvroCRC64Fingerprint() []byte {
	return []byte(VtexSucursalesUpdateEventAvroCRC64Fingerprint)
}

func (r VtexSucursalesUpdateEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["User"], err = json.Marshal(r.User)
	if err != nil {
		return nil, err
	}
	output["CarrierType"], err = json.Marshal(r.CarrierType)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *VtexSucursalesUpdateEvent) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["User"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.User); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for User")
	}
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
	return nil
}