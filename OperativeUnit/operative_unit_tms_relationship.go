// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     OperativeUnitTMSRelationship.avsc
 */
package OperativeUnitEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type OperativeUnitTMSRelationship struct {
	IdIntegra int64 `json:"idIntegra"`

	NumberIntegra string `json:"numberIntegra"`

	IdAlertran string `json:"idAlertran"`
}

const OperativeUnitTMSRelationshipAvroCRC64Fingerprint = "\x95\xbe$(\xb3\xc5\v\xe5"

func NewOperativeUnitTMSRelationship() OperativeUnitTMSRelationship {
	r := OperativeUnitTMSRelationship{}
	return r
}

func DeserializeOperativeUnitTMSRelationship(r io.Reader) (OperativeUnitTMSRelationship, error) {
	t := NewOperativeUnitTMSRelationship()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeOperativeUnitTMSRelationshipFromSchema(r io.Reader, schema string) (OperativeUnitTMSRelationship, error) {
	t := NewOperativeUnitTMSRelationship()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeOperativeUnitTMSRelationship(r OperativeUnitTMSRelationship, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.IdIntegra, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumberIntegra, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.IdAlertran, w)
	if err != nil {
		return err
	}
	return err
}

func (r OperativeUnitTMSRelationship) Serialize(w io.Writer) error {
	return writeOperativeUnitTMSRelationship(r, w)
}

func (r OperativeUnitTMSRelationship) Schema() string {
	return "{\"fields\":[{\"name\":\"idIntegra\",\"type\":\"long\"},{\"name\":\"numberIntegra\",\"type\":\"string\"},{\"name\":\"idAlertran\",\"type\":\"string\"}],\"name\":\"Andreani.OperativeUnit.Events.Record.OperativeUnitTMSRelationship\",\"type\":\"record\"}"
}

func (r OperativeUnitTMSRelationship) SchemaName() string {
	return "Andreani.OperativeUnit.Events.Record.OperativeUnitTMSRelationship"
}

func (_ OperativeUnitTMSRelationship) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ OperativeUnitTMSRelationship) SetInt(v int32)       { panic("Unsupported operation") }
func (_ OperativeUnitTMSRelationship) SetLong(v int64)      { panic("Unsupported operation") }
func (_ OperativeUnitTMSRelationship) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ OperativeUnitTMSRelationship) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ OperativeUnitTMSRelationship) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ OperativeUnitTMSRelationship) SetString(v string)   { panic("Unsupported operation") }
func (_ OperativeUnitTMSRelationship) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *OperativeUnitTMSRelationship) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.IdIntegra}

		return w

	case 1:
		w := types.String{Target: &r.NumberIntegra}

		return w

	case 2:
		w := types.String{Target: &r.IdAlertran}

		return w

	}
	panic("Unknown field index")
}

func (r *OperativeUnitTMSRelationship) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *OperativeUnitTMSRelationship) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ OperativeUnitTMSRelationship) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ OperativeUnitTMSRelationship) AppendArray() types.Field { panic("Unsupported operation") }
func (_ OperativeUnitTMSRelationship) HintSize(int)             { panic("Unsupported operation") }
func (_ OperativeUnitTMSRelationship) Finalize()                {}

func (_ OperativeUnitTMSRelationship) AvroCRC64Fingerprint() []byte {
	return []byte(OperativeUnitTMSRelationshipAvroCRC64Fingerprint)
}

func (r OperativeUnitTMSRelationship) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["idIntegra"], err = json.Marshal(r.IdIntegra)
	if err != nil {
		return nil, err
	}
	output["numberIntegra"], err = json.Marshal(r.NumberIntegra)
	if err != nil {
		return nil, err
	}
	output["idAlertran"], err = json.Marshal(r.IdAlertran)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *OperativeUnitTMSRelationship) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["idIntegra"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdIntegra); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idIntegra")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numberIntegra"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumberIntegra); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numberIntegra")
	}
	val = func() json.RawMessage {
		if v, ok := fields["idAlertran"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdAlertran); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idAlertran")
	}
	return nil
}
