// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     SharepointWorkOrder.avsc
 */
package EAMEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type SharepointWorkOrder struct {
	WorkOrder WorkOrder `json:"WorkOrder"`
}

const SharepointWorkOrderAvroCRC64Fingerprint = "\xe3\xfd\x8d2+\xf7c\t"

func NewSharepointWorkOrder() SharepointWorkOrder {
	r := SharepointWorkOrder{}
	r.WorkOrder = NewWorkOrder()

	return r
}

func DeserializeSharepointWorkOrder(r io.Reader) (SharepointWorkOrder, error) {
	t := NewSharepointWorkOrder()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSharepointWorkOrderFromSchema(r io.Reader, schema string) (SharepointWorkOrder, error) {
	t := NewSharepointWorkOrder()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSharepointWorkOrder(r SharepointWorkOrder, w io.Writer) error {
	var err error
	err = writeWorkOrder(r.WorkOrder, w)
	if err != nil {
		return err
	}
	return err
}

func (r SharepointWorkOrder) Serialize(w io.Writer) error {
	return writeSharepointWorkOrder(r, w)
}

func (r SharepointWorkOrder) Schema() string {
	return "{\"fields\":[{\"name\":\"WorkOrder\",\"type\":{\"fields\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"id_equipo\",\"type\":\"string\"},{\"name\":\"planta\",\"type\":\"string\"},{\"name\":\"descripcion\",\"type\":\"string\"},{\"name\":\"user_report\",\"type\":\"string\"},{\"name\":\"clase\",\"type\":\"string\"},{\"name\":\"tipoOT\",\"type\":\"string\"},{\"name\":\"subTipoOT\",\"type\":\"string\"},{\"name\":\"desvio\",\"type\":\"string\"},{\"name\":\"user_report_email\",\"type\":\"string\"},{\"default\":null,\"name\":\"propietario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"date_report\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ceco\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"equipo_id_catalog\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"equipo_modelo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"equipo_marca\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"equipo_inventario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"equipo_categoria\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"equipo_descripcion\",\"type\":[\"null\",\"string\"]}],\"name\":\"WorkOrder\",\"namespace\":\"Andreani.EAM.Events.Sharepoint\",\"type\":\"record\"}}],\"name\":\"Andreani.EAM.Events.Record.SharepointWorkOrder\",\"type\":\"record\"}"
}

func (r SharepointWorkOrder) SchemaName() string {
	return "Andreani.EAM.Events.Record.SharepointWorkOrder"
}

func (_ SharepointWorkOrder) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ SharepointWorkOrder) SetInt(v int32)       { panic("Unsupported operation") }
func (_ SharepointWorkOrder) SetLong(v int64)      { panic("Unsupported operation") }
func (_ SharepointWorkOrder) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ SharepointWorkOrder) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ SharepointWorkOrder) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ SharepointWorkOrder) SetString(v string)   { panic("Unsupported operation") }
func (_ SharepointWorkOrder) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *SharepointWorkOrder) Get(i int) types.Field {
	switch i {
	case 0:
		r.WorkOrder = NewWorkOrder()

		w := types.Record{Target: &r.WorkOrder}

		return w

	}
	panic("Unknown field index")
}

func (r *SharepointWorkOrder) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *SharepointWorkOrder) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ SharepointWorkOrder) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ SharepointWorkOrder) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ SharepointWorkOrder) HintSize(int)                     { panic("Unsupported operation") }
func (_ SharepointWorkOrder) Finalize()                        {}

func (_ SharepointWorkOrder) AvroCRC64Fingerprint() []byte {
	return []byte(SharepointWorkOrderAvroCRC64Fingerprint)
}

func (r SharepointWorkOrder) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["WorkOrder"], err = json.Marshal(r.WorkOrder)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *SharepointWorkOrder) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["WorkOrder"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.WorkOrder); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for WorkOrder")
	}
	return nil
}
