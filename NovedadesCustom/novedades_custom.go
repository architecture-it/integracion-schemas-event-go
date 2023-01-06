// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NovedadesCustom.avsc
 */
package NovedadesCustomEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type NovedadesCustom struct {
	Traza *UnionNullTraza `json:"traza"`
}

const NovedadesCustomAvroCRC64Fingerprint = "\x00\x99\xba\xa0\xde+\"\xbe"

func NewNovedadesCustom() NovedadesCustom {
	r := NovedadesCustom{}
	r.Traza = nil
	return r
}

func DeserializeNovedadesCustom(r io.Reader) (NovedadesCustom, error) {
	t := NewNovedadesCustom()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNovedadesCustomFromSchema(r io.Reader, schema string) (NovedadesCustom, error) {
	t := NewNovedadesCustom()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNovedadesCustom(r NovedadesCustom, w io.Writer) error {
	var err error
	err = writeUnionNullTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r NovedadesCustom) Serialize(w io.Writer) error {
	return writeNovedadesCustom(r, w)
}

func (r NovedadesCustom) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"traza\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Evento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeContratoInterno\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"CicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Motivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SubMotivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DatosAdicionales\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"Key\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Value\",\"type\":[\"null\",\"string\"]}],\"name\":\"Metadata\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"Traza\",\"namespace\":\"Andreani.NovedadesCustom.Events.Common\",\"type\":\"record\"}]}],\"name\":\"Andreani.NovedadesCustom.Events.Record.NovedadesCustom\",\"type\":\"record\"}"
}

func (r NovedadesCustom) SchemaName() string {
	return "Andreani.NovedadesCustom.Events.Record.NovedadesCustom"
}

func (_ NovedadesCustom) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NovedadesCustom) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NovedadesCustom) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NovedadesCustom) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NovedadesCustom) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NovedadesCustom) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NovedadesCustom) SetString(v string)   { panic("Unsupported operation") }
func (_ NovedadesCustom) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NovedadesCustom) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewUnionNullTraza()

		return r.Traza
	}
	panic("Unknown field index")
}

func (r *NovedadesCustom) SetDefault(i int) {
	switch i {
	case 0:
		r.Traza = nil
		return
	}
	panic("Unknown field index")
}

func (r *NovedadesCustom) NullField(i int) {
	switch i {
	case 0:
		r.Traza = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ NovedadesCustom) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NovedadesCustom) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NovedadesCustom) HintSize(int)                     { panic("Unsupported operation") }
func (_ NovedadesCustom) Finalize()                        {}

func (_ NovedadesCustom) AvroCRC64Fingerprint() []byte {
	return []byte(NovedadesCustomAvroCRC64Fingerprint)
}

func (r NovedadesCustom) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NovedadesCustom) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["traza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Traza); err != nil {
			return err
		}
	} else {
		r.Traza = NewUnionNullTraza()

		r.Traza = nil
	}
	return nil
}
