// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Phase.avsc
 */
package appnameEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Phase struct {
	Token string `json:"Token"`

	Tenant string `json:"Tenant"`

	Data *UnionNullPhaseData `json:"Data"`
}

const PhaseAvroCRC64Fingerprint = "2\r\xd8QN\xf3\xf2\xf4"

func NewPhase() Phase {
	r := Phase{}
	r.Data = nil
	return r
}

func DeserializePhase(r io.Reader) (Phase, error) {
	t := NewPhase()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePhaseFromSchema(r io.Reader, schema string) (Phase, error) {
	t := NewPhase()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePhase(r Phase, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Token, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Tenant, w)
	if err != nil {
		return err
	}
	err = writeUnionNullPhaseData(r.Data, w)
	if err != nil {
		return err
	}
	return err
}

func (r Phase) Serialize(w io.Writer) error {
	return writePhase(r, w)
}

func (r Phase) Schema() string {
	return "{\"fields\":[{\"name\":\"Token\",\"type\":\"string\"},{\"name\":\"Tenant\",\"type\":\"string\"},{\"default\":null,\"name\":\"Data\",\"type\":[\"null\",{\"fields\":[{\"name\":\"EndDate\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"StartDate\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"AdvanceNotice\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DecouplingCause\",\"type\":[\"null\",\"string\"]},{\"name\":\"Compensation\",\"type\":\"boolean\"},{\"name\":\"Holidays\",\"type\":\"boolean\"},{\"name\":\"Real\",\"type\":\"boolean\"},{\"name\":\"RecognizedStartDate\",\"type\":\"boolean\"},{\"name\":\"Salary\",\"type\":\"boolean\"},{\"name\":\"Status\",\"type\":\"boolean\"},{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"EmployeeId\",\"type\":\"int\"}],\"name\":\"PhaseData\",\"namespace\":\"Andreani.RHpro.Events.Common\",\"type\":\"record\"}]}],\"name\":\"Andreani.RHpro.Events.Record.Phase\",\"type\":\"record\"}"
}

func (r Phase) SchemaName() string {
	return "Andreani.RHpro.Events.Record.Phase"
}

func (_ Phase) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Phase) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Phase) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Phase) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Phase) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Phase) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Phase) SetString(v string)   { panic("Unsupported operation") }
func (_ Phase) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Phase) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Token}

		return w

	case 1:
		w := types.String{Target: &r.Tenant}

		return w

	case 2:
		r.Data = NewUnionNullPhaseData()

		return r.Data
	}
	panic("Unknown field index")
}

func (r *Phase) SetDefault(i int) {
	switch i {
	case 2:
		r.Data = nil
		return
	}
	panic("Unknown field index")
}

func (r *Phase) NullField(i int) {
	switch i {
	case 2:
		r.Data = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Phase) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Phase) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Phase) HintSize(int)                     { panic("Unsupported operation") }
func (_ Phase) Finalize()                        {}

func (_ Phase) AvroCRC64Fingerprint() []byte {
	return []byte(PhaseAvroCRC64Fingerprint)
}

func (r Phase) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Token"], err = json.Marshal(r.Token)
	if err != nil {
		return nil, err
	}
	output["Tenant"], err = json.Marshal(r.Tenant)
	if err != nil {
		return nil, err
	}
	output["Data"], err = json.Marshal(r.Data)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Phase) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Token"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Token); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Token")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Tenant"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tenant); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Tenant")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Data"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Data); err != nil {
			return err
		}
	} else {
		r.Data = NewUnionNullPhaseData()

		r.Data = nil
	}
	return nil
}
