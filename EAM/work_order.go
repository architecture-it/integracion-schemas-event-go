// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     WorkOrder.avsc
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

type WorkOrder struct {
	Id string `json:"id"`

	Id_equipo string `json:"id_equipo"`

	Planta string `json:"planta"`

	Descripcion string `json:"descripcion"`

	User_report string `json:"user_report"`

	Clase string `json:"clase"`

	TipoOT string `json:"tipoOT"`

	SubTipoOT string `json:"subTipoOT"`
}

const WorkOrderAvroCRC64Fingerprint = "\x16\xbcW\xfb\xd6.\xe4\x18"

func NewWorkOrder() WorkOrder {
	r := WorkOrder{}
	return r
}

func DeserializeWorkOrder(r io.Reader) (WorkOrder, error) {
	t := NewWorkOrder()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeWorkOrderFromSchema(r io.Reader, schema string) (WorkOrder, error) {
	t := NewWorkOrder()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeWorkOrder(r WorkOrder, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Id_equipo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Planta, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Descripcion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.User_report, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Clase, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TipoOT, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.SubTipoOT, w)
	if err != nil {
		return err
	}
	return err
}

func (r WorkOrder) Serialize(w io.Writer) error {
	return writeWorkOrder(r, w)
}

func (r WorkOrder) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"id_equipo\",\"type\":\"string\"},{\"name\":\"planta\",\"type\":\"string\"},{\"name\":\"descripcion\",\"type\":\"string\"},{\"name\":\"user_report\",\"type\":\"string\"},{\"name\":\"clase\",\"type\":\"string\"},{\"name\":\"tipoOT\",\"type\":\"string\"},{\"name\":\"subTipoOT\",\"type\":\"string\"}],\"name\":\"Andreani.EAM.Events.Sharepoint.WorkOrder\",\"type\":\"record\"}"
}

func (r WorkOrder) SchemaName() string {
	return "Andreani.EAM.Events.Sharepoint.WorkOrder"
}

func (_ WorkOrder) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ WorkOrder) SetInt(v int32)       { panic("Unsupported operation") }
func (_ WorkOrder) SetLong(v int64)      { panic("Unsupported operation") }
func (_ WorkOrder) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ WorkOrder) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ WorkOrder) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ WorkOrder) SetString(v string)   { panic("Unsupported operation") }
func (_ WorkOrder) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *WorkOrder) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Id_equipo}

		return w

	case 2:
		w := types.String{Target: &r.Planta}

		return w

	case 3:
		w := types.String{Target: &r.Descripcion}

		return w

	case 4:
		w := types.String{Target: &r.User_report}

		return w

	case 5:
		w := types.String{Target: &r.Clase}

		return w

	case 6:
		w := types.String{Target: &r.TipoOT}

		return w

	case 7:
		w := types.String{Target: &r.SubTipoOT}

		return w

	}
	panic("Unknown field index")
}

func (r *WorkOrder) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *WorkOrder) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ WorkOrder) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ WorkOrder) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ WorkOrder) HintSize(int)                     { panic("Unsupported operation") }
func (_ WorkOrder) Finalize()                        {}

func (_ WorkOrder) AvroCRC64Fingerprint() []byte {
	return []byte(WorkOrderAvroCRC64Fingerprint)
}

func (r WorkOrder) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["id_equipo"], err = json.Marshal(r.Id_equipo)
	if err != nil {
		return nil, err
	}
	output["planta"], err = json.Marshal(r.Planta)
	if err != nil {
		return nil, err
	}
	output["descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	output["user_report"], err = json.Marshal(r.User_report)
	if err != nil {
		return nil, err
	}
	output["clase"], err = json.Marshal(r.Clase)
	if err != nil {
		return nil, err
	}
	output["tipoOT"], err = json.Marshal(r.TipoOT)
	if err != nil {
		return nil, err
	}
	output["subTipoOT"], err = json.Marshal(r.SubTipoOT)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *WorkOrder) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["id_equipo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id_equipo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id_equipo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["planta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Planta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for planta")
	}
	val = func() json.RawMessage {
		if v, ok := fields["descripcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Descripcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for descripcion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["user_report"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.User_report); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for user_report")
	}
	val = func() json.RawMessage {
		if v, ok := fields["clase"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Clase); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for clase")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoOT"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoOT); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tipoOT")
	}
	val = func() json.RawMessage {
		if v, ok := fields["subTipoOT"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SubTipoOT); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for subTipoOT")
	}
	return nil
}
