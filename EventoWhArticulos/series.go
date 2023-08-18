// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Series.avsc
 */
package EventoWhArticulosEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Series struct {
	Serie string `json:"Serie"`

	Data2 string `json:"Data2"`

	Data3 string `json:"Data3"`
}

const SeriesAvroCRC64Fingerprint = "\x00\x89\xd8\xc1\x04]^\x14"

func NewSeries() Series {
	r := Series{}
	return r
}

func DeserializeSeries(r io.Reader) (Series, error) {
	t := NewSeries()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSeriesFromSchema(r io.Reader, schema string) (Series, error) {
	t := NewSeries()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSeries(r Series, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Serie, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Data2, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Data3, w)
	if err != nil {
		return err
	}
	return err
}

func (r Series) Serialize(w io.Writer) error {
	return writeSeries(r, w)
}

func (r Series) Schema() string {
	return "{\"fields\":[{\"name\":\"Serie\",\"type\":\"string\"},{\"name\":\"Data2\",\"type\":\"string\"},{\"name\":\"Data3\",\"type\":\"string\"}],\"name\":\"Andreani.EventoWhArticulos.Events.ArticuloAjusteCommon.Series\",\"type\":\"record\"}"
}

func (r Series) SchemaName() string {
	return "Andreani.EventoWhArticulos.Events.ArticuloAjusteCommon.Series"
}

func (_ Series) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Series) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Series) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Series) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Series) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Series) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Series) SetString(v string)   { panic("Unsupported operation") }
func (_ Series) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Series) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Serie}

		return w

	case 1:
		w := types.String{Target: &r.Data2}

		return w

	case 2:
		w := types.String{Target: &r.Data3}

		return w

	}
	panic("Unknown field index")
}

func (r *Series) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Series) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Series) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Series) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Series) HintSize(int)                     { panic("Unsupported operation") }
func (_ Series) Finalize()                        {}

func (_ Series) AvroCRC64Fingerprint() []byte {
	return []byte(SeriesAvroCRC64Fingerprint)
}

func (r Series) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Serie"], err = json.Marshal(r.Serie)
	if err != nil {
		return nil, err
	}
	output["Data2"], err = json.Marshal(r.Data2)
	if err != nil {
		return nil, err
	}
	output["Data3"], err = json.Marshal(r.Data3)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Series) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Serie"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Serie); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Serie")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Data2"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Data2); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Data2")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Data3"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Data3); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Data3")
	}
	return nil
}
