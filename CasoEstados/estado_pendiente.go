// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EstadoPendiente.avsc
 */
package CasoEstadosEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EstadoPendiente struct {
	Ticketnumber string `json:"Ticketnumber"`

	Subjectid string `json:"Subjectid"`

	And_numerodeenvio *UnionNullString `json:"And_numerodeenvio"`

	Title string `json:"Title"`

	Description *UnionNullString `json:"Description"`

	Createdon string `json:"Createdon"`

	Cac_numerodeenvioincorrecto *UnionNullString `json:"Cac_numerodeenvioincorrecto"`

	Customername string `json:"Customername"`

	Cac_areainterna *UnionNullString `json:"Cac_areainterna"`

	StatusCodeName string `json:"StatusCodeName"`

	IncidentId *UnionNullString `json:"IncidentId"`

	Origen *UnionNullString `json:"Origen"`
}

const EstadoPendienteAvroCRC64Fingerprint = "81\x047\xaf\x1cnK"

func NewEstadoPendiente() EstadoPendiente {
	r := EstadoPendiente{}
	return r
}

func DeserializeEstadoPendiente(r io.Reader) (EstadoPendiente, error) {
	t := NewEstadoPendiente()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEstadoPendienteFromSchema(r io.Reader, schema string) (EstadoPendiente, error) {
	t := NewEstadoPendiente()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEstadoPendiente(r EstadoPendiente, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Ticketnumber, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Subjectid, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.And_numerodeenvio, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Title, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Description, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Createdon, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Cac_numerodeenvioincorrecto, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Customername, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Cac_areainterna, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.StatusCodeName, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.IncidentId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Origen, w)
	if err != nil {
		return err
	}
	return err
}

func (r EstadoPendiente) Serialize(w io.Writer) error {
	return writeEstadoPendiente(r, w)
}

func (r EstadoPendiente) Schema() string {
	return "{\"fields\":[{\"name\":\"Ticketnumber\",\"type\":\"string\"},{\"name\":\"Subjectid\",\"type\":\"string\"},{\"name\":\"And_numerodeenvio\",\"type\":[\"null\",\"string\"]},{\"name\":\"Title\",\"type\":\"string\"},{\"name\":\"Description\",\"type\":[\"null\",\"string\"]},{\"name\":\"Createdon\",\"type\":\"string\"},{\"name\":\"Cac_numerodeenvioincorrecto\",\"type\":[\"null\",\"string\"]},{\"name\":\"Customername\",\"type\":\"string\"},{\"name\":\"Cac_areainterna\",\"type\":[\"null\",\"string\"]},{\"name\":\"StatusCodeName\",\"type\":\"string\"},{\"name\":\"IncidentId\",\"type\":[\"null\",\"string\"]},{\"name\":\"Origen\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.CasoEstados.Events.Record.EstadoPendiente\",\"type\":\"record\"}"
}

func (r EstadoPendiente) SchemaName() string {
	return "Andreani.CasoEstados.Events.Record.EstadoPendiente"
}

func (_ EstadoPendiente) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EstadoPendiente) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EstadoPendiente) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EstadoPendiente) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EstadoPendiente) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EstadoPendiente) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EstadoPendiente) SetString(v string)   { panic("Unsupported operation") }
func (_ EstadoPendiente) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EstadoPendiente) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Ticketnumber}

		return w

	case 1:
		w := types.String{Target: &r.Subjectid}

		return w

	case 2:
		r.And_numerodeenvio = NewUnionNullString()

		return r.And_numerodeenvio
	case 3:
		w := types.String{Target: &r.Title}

		return w

	case 4:
		r.Description = NewUnionNullString()

		return r.Description
	case 5:
		w := types.String{Target: &r.Createdon}

		return w

	case 6:
		r.Cac_numerodeenvioincorrecto = NewUnionNullString()

		return r.Cac_numerodeenvioincorrecto
	case 7:
		w := types.String{Target: &r.Customername}

		return w

	case 8:
		r.Cac_areainterna = NewUnionNullString()

		return r.Cac_areainterna
	case 9:
		w := types.String{Target: &r.StatusCodeName}

		return w

	case 10:
		r.IncidentId = NewUnionNullString()

		return r.IncidentId
	case 11:
		r.Origen = NewUnionNullString()

		return r.Origen
	}
	panic("Unknown field index")
}

func (r *EstadoPendiente) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EstadoPendiente) NullField(i int) {
	switch i {
	case 2:
		r.And_numerodeenvio = nil
		return
	case 4:
		r.Description = nil
		return
	case 6:
		r.Cac_numerodeenvioincorrecto = nil
		return
	case 8:
		r.Cac_areainterna = nil
		return
	case 10:
		r.IncidentId = nil
		return
	case 11:
		r.Origen = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ EstadoPendiente) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EstadoPendiente) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EstadoPendiente) HintSize(int)                     { panic("Unsupported operation") }
func (_ EstadoPendiente) Finalize()                        {}

func (_ EstadoPendiente) AvroCRC64Fingerprint() []byte {
	return []byte(EstadoPendienteAvroCRC64Fingerprint)
}

func (r EstadoPendiente) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Ticketnumber"], err = json.Marshal(r.Ticketnumber)
	if err != nil {
		return nil, err
	}
	output["Subjectid"], err = json.Marshal(r.Subjectid)
	if err != nil {
		return nil, err
	}
	output["And_numerodeenvio"], err = json.Marshal(r.And_numerodeenvio)
	if err != nil {
		return nil, err
	}
	output["Title"], err = json.Marshal(r.Title)
	if err != nil {
		return nil, err
	}
	output["Description"], err = json.Marshal(r.Description)
	if err != nil {
		return nil, err
	}
	output["Createdon"], err = json.Marshal(r.Createdon)
	if err != nil {
		return nil, err
	}
	output["Cac_numerodeenvioincorrecto"], err = json.Marshal(r.Cac_numerodeenvioincorrecto)
	if err != nil {
		return nil, err
	}
	output["Customername"], err = json.Marshal(r.Customername)
	if err != nil {
		return nil, err
	}
	output["Cac_areainterna"], err = json.Marshal(r.Cac_areainterna)
	if err != nil {
		return nil, err
	}
	output["StatusCodeName"], err = json.Marshal(r.StatusCodeName)
	if err != nil {
		return nil, err
	}
	output["IncidentId"], err = json.Marshal(r.IncidentId)
	if err != nil {
		return nil, err
	}
	output["Origen"], err = json.Marshal(r.Origen)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EstadoPendiente) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Ticketnumber"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ticketnumber); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Ticketnumber")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Subjectid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Subjectid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Subjectid")
	}
	val = func() json.RawMessage {
		if v, ok := fields["And_numerodeenvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.And_numerodeenvio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for And_numerodeenvio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Title"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Title); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Title")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Description"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Description); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Description")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Createdon"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Createdon); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Createdon")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Cac_numerodeenvioincorrecto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cac_numerodeenvioincorrecto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Cac_numerodeenvioincorrecto")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Customername"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Customername); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Customername")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Cac_areainterna"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cac_areainterna); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Cac_areainterna")
	}
	val = func() json.RawMessage {
		if v, ok := fields["StatusCodeName"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StatusCodeName); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for StatusCodeName")
	}
	val = func() json.RawMessage {
		if v, ok := fields["IncidentId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IncidentId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IncidentId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Origen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Origen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Origen")
	}
	return nil
}
