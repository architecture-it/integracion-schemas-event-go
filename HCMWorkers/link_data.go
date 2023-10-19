// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     WorkerData.avsc
 */
package HCMWorkersEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type LinkData struct {
	Rel string `json:"rel"`

	Href string `json:"href"`

	Name string `json:"name"`

	Kind string `json:"kind"`

	Properties *UnionNullPropertiesData `json:"properties"`
}

const LinkDataAvroCRC64Fingerprint = "\x14\xf7\xa1Kt\xca\xcd\xfb"

func NewLinkData() LinkData {
	r := LinkData{}
	r.Properties = nil
	return r
}

func DeserializeLinkData(r io.Reader) (LinkData, error) {
	t := NewLinkData()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeLinkDataFromSchema(r io.Reader, schema string) (LinkData, error) {
	t := NewLinkData()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeLinkData(r LinkData, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Rel, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Href, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Kind, w)
	if err != nil {
		return err
	}
	err = writeUnionNullPropertiesData(r.Properties, w)
	if err != nil {
		return err
	}
	return err
}

func (r LinkData) Serialize(w io.Writer) error {
	return writeLinkData(r, w)
}

func (r LinkData) Schema() string {
	return "{\"fields\":[{\"name\":\"rel\",\"type\":\"string\"},{\"name\":\"href\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"kind\",\"type\":\"string\"},{\"default\":null,\"name\":\"properties\",\"type\":[\"null\",{\"fields\":[{\"name\":\"changeIndicator\",\"type\":\"string\"}],\"name\":\"PropertiesData\",\"type\":\"record\"}]}],\"name\":\"Andreani.HCMWorkers.Events.Record.LinkData\",\"type\":\"record\"}"
}

func (r LinkData) SchemaName() string {
	return "Andreani.HCMWorkers.Events.Record.LinkData"
}

func (_ LinkData) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ LinkData) SetInt(v int32)       { panic("Unsupported operation") }
func (_ LinkData) SetLong(v int64)      { panic("Unsupported operation") }
func (_ LinkData) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ LinkData) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ LinkData) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ LinkData) SetString(v string)   { panic("Unsupported operation") }
func (_ LinkData) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *LinkData) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Rel}

		return w

	case 1:
		w := types.String{Target: &r.Href}

		return w

	case 2:
		w := types.String{Target: &r.Name}

		return w

	case 3:
		w := types.String{Target: &r.Kind}

		return w

	case 4:
		r.Properties = NewUnionNullPropertiesData()

		return r.Properties
	}
	panic("Unknown field index")
}

func (r *LinkData) SetDefault(i int) {
	switch i {
	case 4:
		r.Properties = nil
		return
	}
	panic("Unknown field index")
}

func (r *LinkData) NullField(i int) {
	switch i {
	case 4:
		r.Properties = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ LinkData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ LinkData) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ LinkData) HintSize(int)                     { panic("Unsupported operation") }
func (_ LinkData) Finalize()                        {}

func (_ LinkData) AvroCRC64Fingerprint() []byte {
	return []byte(LinkDataAvroCRC64Fingerprint)
}

func (r LinkData) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["rel"], err = json.Marshal(r.Rel)
	if err != nil {
		return nil, err
	}
	output["href"], err = json.Marshal(r.Href)
	if err != nil {
		return nil, err
	}
	output["name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	output["kind"], err = json.Marshal(r.Kind)
	if err != nil {
		return nil, err
	}
	output["properties"], err = json.Marshal(r.Properties)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *LinkData) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["rel"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Rel); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for rel")
	}
	val = func() json.RawMessage {
		if v, ok := fields["href"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Href); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for href")
	}
	val = func() json.RawMessage {
		if v, ok := fields["name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["kind"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Kind); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for kind")
	}
	val = func() json.RawMessage {
		if v, ok := fields["properties"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Properties); err != nil {
			return err
		}
	} else {
		r.Properties = NewUnionNullPropertiesData()

		r.Properties = nil
	}
	return nil
}
