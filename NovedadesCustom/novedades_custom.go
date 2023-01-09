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
	Traza *UnionNullTraza `json:"Traza"`

	DatosAdicionales *UnionNullArrayMetadata `json:"DatosAdicionales"`
}

const NovedadesCustomAvroCRC64Fingerprint = "\xc9\x05\x85\xf2gMm\xbc"

func NewNovedadesCustom() NovedadesCustom {
	r := NovedadesCustom{}
	r.Traza = nil
	r.DatosAdicionales = nil
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
	err = writeUnionNullArrayMetadata(r.DatosAdicionales, w)
	if err != nil {
		return err
	}
	return err
}

func (r NovedadesCustom) Serialize(w io.Writer) error {
	return writeNovedadesCustom(r, w)
}

func (r NovedadesCustom) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Traza\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Evento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeContratoInterno\",\"type\":[\"null\",\"string\"]},{\"name\":\"Cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"CicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Motivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SubMotivo\",\"type\":[\"null\",\"string\"]}],\"name\":\"Traza\",\"namespace\":\"Andreani.NovedadesCustom.Events.Common\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"DatosAdicionales\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"Key\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Value\",\"type\":[\"null\",\"string\"]}],\"name\":\"Metadata\",\"namespace\":\"Andreani.NovedadesCustom.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"Andreani.NovedadesCustom.Events.Record.NovedadesCustom\",\"type\":\"record\"}"
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
	case 1:
		r.DatosAdicionales = NewUnionNullArrayMetadata()

		return r.DatosAdicionales
	}
	panic("Unknown field index")
}

func (r *NovedadesCustom) SetDefault(i int) {
	switch i {
	case 0:
		r.Traza = nil
		return
	case 1:
		r.DatosAdicionales = nil
		return
	}
	panic("Unknown field index")
}

func (r *NovedadesCustom) NullField(i int) {
	switch i {
	case 0:
		r.Traza = nil
		return
	case 1:
		r.DatosAdicionales = nil
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
	output["Traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["DatosAdicionales"], err = json.Marshal(r.DatosAdicionales)
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
		if v, ok := fields["Traza"]; ok {
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
	val = func() json.RawMessage {
		if v, ok := fields["DatosAdicionales"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DatosAdicionales); err != nil {
			return err
		}
	} else {
		r.DatosAdicionales = NewUnionNullArrayMetadata()

		r.DatosAdicionales = nil
	}
	return nil
}
