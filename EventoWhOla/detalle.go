// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhOlaLanzadaWosPicking.avsc
 */
package EventoWhOlaEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Detalle struct {
	TaskDetailKey string `json:"TaskDetailKey"`

	TaskType string `json:"TaskType"`

	StorerKey string `json:"StorerKey"`

	Sku string `json:"Sku"`
}

const DetalleAvroCRC64Fingerprint = "\x01\x8b\xc0\x8aס_0"

func NewDetalle() Detalle {
	r := Detalle{}
	return r
}

func DeserializeDetalle(r io.Reader) (Detalle, error) {
	t := NewDetalle()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDetalleFromSchema(r io.Reader, schema string) (Detalle, error) {
	t := NewDetalle()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDetalle(r Detalle, w io.Writer) error {
	var err error
	err = vm.WriteString(r.TaskDetailKey, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TaskType, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.StorerKey, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Sku, w)
	if err != nil {
		return err
	}
	return err
}

func (r Detalle) Serialize(w io.Writer) error {
	return writeDetalle(r, w)
}

func (r Detalle) Schema() string {
	return "{\"fields\":[{\"name\":\"TaskDetailKey\",\"type\":\"string\"},{\"name\":\"TaskType\",\"type\":\"string\"},{\"name\":\"StorerKey\",\"type\":\"string\"},{\"name\":\"Sku\",\"type\":\"string\"}],\"name\":\"Andreani.EventoWhOla.Events.LanzadaWosPickingCommon.Detalle\",\"type\":\"record\"}"
}

func (r Detalle) SchemaName() string {
	return "Andreani.EventoWhOla.Events.LanzadaWosPickingCommon.Detalle"
}

func (_ Detalle) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Detalle) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Detalle) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Detalle) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Detalle) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Detalle) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Detalle) SetString(v string)   { panic("Unsupported operation") }
func (_ Detalle) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Detalle) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.TaskDetailKey}

		return w

	case 1:
		w := types.String{Target: &r.TaskType}

		return w

	case 2:
		w := types.String{Target: &r.StorerKey}

		return w

	case 3:
		w := types.String{Target: &r.Sku}

		return w

	}
	panic("Unknown field index")
}

func (r *Detalle) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Detalle) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Detalle) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Detalle) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Detalle) HintSize(int)                     { panic("Unsupported operation") }
func (_ Detalle) Finalize()                        {}

func (_ Detalle) AvroCRC64Fingerprint() []byte {
	return []byte(DetalleAvroCRC64Fingerprint)
}

func (r Detalle) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["TaskDetailKey"], err = json.Marshal(r.TaskDetailKey)
	if err != nil {
		return nil, err
	}
	output["TaskType"], err = json.Marshal(r.TaskType)
	if err != nil {
		return nil, err
	}
	output["StorerKey"], err = json.Marshal(r.StorerKey)
	if err != nil {
		return nil, err
	}
	output["Sku"], err = json.Marshal(r.Sku)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Detalle) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["TaskDetailKey"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TaskDetailKey); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TaskDetailKey")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TaskType"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TaskType); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TaskType")
	}
	val = func() json.RawMessage {
		if v, ok := fields["StorerKey"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StorerKey); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for StorerKey")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Sku"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sku); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Sku")
	}
	return nil
}
