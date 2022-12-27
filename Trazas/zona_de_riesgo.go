// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ZonaDeRiesgo.avsc
 */
package TrazasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ZonaDeRiesgo struct {
	Traza Traza `json:"traza"`
}

const ZonaDeRiesgoAvroCRC64Fingerprint = "^\xcc\xe9\xe2\xdf+\xc3\xcb"

func NewZonaDeRiesgo() ZonaDeRiesgo {
	r := ZonaDeRiesgo{}
	r.Traza = NewTraza()

	return r
}

func DeserializeZonaDeRiesgo(r io.Reader) (ZonaDeRiesgo, error) {
	t := NewZonaDeRiesgo()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeZonaDeRiesgoFromSchema(r io.Reader, schema string) (ZonaDeRiesgo, error) {
	t := NewZonaDeRiesgo()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeZonaDeRiesgo(r ZonaDeRiesgo, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r ZonaDeRiesgo) Serialize(w io.Writer) error {
	return writeZonaDeRiesgo(r, w)
}

func (r ZonaDeRiesgo) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.ZonaDeRiesgo\",\"type\":\"record\"}"
}

func (r ZonaDeRiesgo) SchemaName() string {
	return "Integracion.Esquemas.Trazas.ZonaDeRiesgo"
}

func (_ ZonaDeRiesgo) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ZonaDeRiesgo) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ZonaDeRiesgo) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ZonaDeRiesgo) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ZonaDeRiesgo) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ZonaDeRiesgo) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ZonaDeRiesgo) SetString(v string)   { panic("Unsupported operation") }
func (_ ZonaDeRiesgo) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ZonaDeRiesgo) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *ZonaDeRiesgo) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ZonaDeRiesgo) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ZonaDeRiesgo) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ZonaDeRiesgo) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ZonaDeRiesgo) HintSize(int)                     { panic("Unsupported operation") }
func (_ ZonaDeRiesgo) Finalize()                        {}

func (_ ZonaDeRiesgo) AvroCRC64Fingerprint() []byte {
	return []byte(ZonaDeRiesgoAvroCRC64Fingerprint)
}

func (r ZonaDeRiesgo) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ZonaDeRiesgo) UnmarshalJSON(data []byte) error {
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
		return fmt.Errorf("no value specified for traza")
	}
	return nil
}