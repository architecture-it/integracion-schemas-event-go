// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AppOperacionHDRControlada.avsc
 */
package MobileOperacionesProducerEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type AppOperacionHDRControlada struct {
	NumeroHdr string `json:"numeroHdr"`

	NumerosSeguimiento []string `json:"numerosSeguimiento"`
}

const AppOperacionHDRControladaAvroCRC64Fingerprint = "\xb3Nq\xb0\xf5\xb7\xf2\xbb"

func NewAppOperacionHDRControlada() AppOperacionHDRControlada {
	r := AppOperacionHDRControlada{}
	r.NumerosSeguimiento = make([]string, 0)

	return r
}

func DeserializeAppOperacionHDRControlada(r io.Reader) (AppOperacionHDRControlada, error) {
	t := NewAppOperacionHDRControlada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAppOperacionHDRControladaFromSchema(r io.Reader, schema string) (AppOperacionHDRControlada, error) {
	t := NewAppOperacionHDRControlada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAppOperacionHDRControlada(r AppOperacionHDRControlada, w io.Writer) error {
	var err error
	err = vm.WriteString(r.NumeroHdr, w)
	if err != nil {
		return err
	}
	err = writeArrayString(r.NumerosSeguimiento, w)
	if err != nil {
		return err
	}
	return err
}

func (r AppOperacionHDRControlada) Serialize(w io.Writer) error {
	return writeAppOperacionHDRControlada(r, w)
}

func (r AppOperacionHDRControlada) Schema() string {
	return "{\"fields\":[{\"name\":\"numeroHdr\",\"type\":\"string\"},{\"name\":\"numerosSeguimiento\",\"type\":{\"items\":\"string\",\"type\":\"array\"}}],\"name\":\"Andreani.MobileOperacionesProducer.Events.Record.AppOperacionHDRControlada\",\"type\":\"record\"}"
}

func (r AppOperacionHDRControlada) SchemaName() string {
	return "Andreani.MobileOperacionesProducer.Events.Record.AppOperacionHDRControlada"
}

func (_ AppOperacionHDRControlada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ AppOperacionHDRControlada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ AppOperacionHDRControlada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ AppOperacionHDRControlada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ AppOperacionHDRControlada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ AppOperacionHDRControlada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ AppOperacionHDRControlada) SetString(v string)   { panic("Unsupported operation") }
func (_ AppOperacionHDRControlada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AppOperacionHDRControlada) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.NumeroHdr}

		return w

	case 1:
		r.NumerosSeguimiento = make([]string, 0)

		w := ArrayStringWrapper{Target: &r.NumerosSeguimiento}

		return w

	}
	panic("Unknown field index")
}

func (r *AppOperacionHDRControlada) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *AppOperacionHDRControlada) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ AppOperacionHDRControlada) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ AppOperacionHDRControlada) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ AppOperacionHDRControlada) HintSize(int)                     { panic("Unsupported operation") }
func (_ AppOperacionHDRControlada) Finalize()                        {}

func (_ AppOperacionHDRControlada) AvroCRC64Fingerprint() []byte {
	return []byte(AppOperacionHDRControladaAvroCRC64Fingerprint)
}

func (r AppOperacionHDRControlada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["numeroHdr"], err = json.Marshal(r.NumeroHdr)
	if err != nil {
		return nil, err
	}
	output["numerosSeguimiento"], err = json.Marshal(r.NumerosSeguimiento)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *AppOperacionHDRControlada) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["numeroHdr"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroHdr); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroHdr")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numerosSeguimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumerosSeguimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numerosSeguimiento")
	}
	return nil
}