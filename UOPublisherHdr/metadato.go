// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     UOHdrCreada.avsc
 */
package UOPublisherHdrEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Metadato struct {
	Meta string `json:"Meta"`

	Contenido string `json:"contenido"`
}

const MetadatoAvroCRC64Fingerprint = "o\x03\xaa\xc8H\xb31\x15"

func NewMetadato() Metadato {
	r := Metadato{}
	return r
}

func DeserializeMetadato(r io.Reader) (Metadato, error) {
	t := NewMetadato()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMetadatoFromSchema(r io.Reader, schema string) (Metadato, error) {
	t := NewMetadato()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMetadato(r Metadato, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Meta, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Contenido, w)
	if err != nil {
		return err
	}
	return err
}

func (r Metadato) Serialize(w io.Writer) error {
	return writeMetadato(r, w)
}

func (r Metadato) Schema() string {
	return "{\"fields\":[{\"name\":\"Meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Andreani.UOPublisherHdr.Events.Common.Metadato\",\"type\":\"record\"}"
}

func (r Metadato) SchemaName() string {
	return "Andreani.UOPublisherHdr.Events.Common.Metadato"
}

func (_ Metadato) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Metadato) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Metadato) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Metadato) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Metadato) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Metadato) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Metadato) SetString(v string)   { panic("Unsupported operation") }
func (_ Metadato) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Metadato) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Meta}

		return w

	case 1:
		w := types.String{Target: &r.Contenido}

		return w

	}
	panic("Unknown field index")
}

func (r *Metadato) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Metadato) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Metadato) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Metadato) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Metadato) HintSize(int)                     { panic("Unsupported operation") }
func (_ Metadato) Finalize()                        {}

func (_ Metadato) AvroCRC64Fingerprint() []byte {
	return []byte(MetadatoAvroCRC64Fingerprint)
}

func (r Metadato) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Meta"], err = json.Marshal(r.Meta)
	if err != nil {
		return nil, err
	}
	output["contenido"], err = json.Marshal(r.Contenido)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Metadato) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Meta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Meta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Meta")
	}
	val = func() json.RawMessage {
		if v, ok := fields["contenido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contenido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for contenido")
	}
	return nil
}
