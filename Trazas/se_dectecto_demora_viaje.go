// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     SeDectectoDemoraViaje.avsc
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

type SeDectectoDemoraViaje struct {
	Traza Traza `json:"traza"`
}

const SeDectectoDemoraViajeAvroCRC64Fingerprint = "\xeb\xeat\xc3\xcb}\xb7H"

func NewSeDectectoDemoraViaje() SeDectectoDemoraViaje {
	r := SeDectectoDemoraViaje{}
	r.Traza = NewTraza()

	return r
}

func DeserializeSeDectectoDemoraViaje(r io.Reader) (SeDectectoDemoraViaje, error) {
	t := NewSeDectectoDemoraViaje()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSeDectectoDemoraViajeFromSchema(r io.Reader, schema string) (SeDectectoDemoraViaje, error) {
	t := NewSeDectectoDemoraViaje()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSeDectectoDemoraViaje(r SeDectectoDemoraViaje, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r SeDectectoDemoraViaje) Serialize(w io.Writer) error {
	return writeSeDectectoDemoraViaje(r, w)
}

func (r SeDectectoDemoraViaje) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.SeDectectoDemoraViaje\",\"type\":\"record\"}"
}

func (r SeDectectoDemoraViaje) SchemaName() string {
	return "Integracion.Esquemas.Trazas.SeDectectoDemoraViaje"
}

func (_ SeDectectoDemoraViaje) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ SeDectectoDemoraViaje) SetInt(v int32)       { panic("Unsupported operation") }
func (_ SeDectectoDemoraViaje) SetLong(v int64)      { panic("Unsupported operation") }
func (_ SeDectectoDemoraViaje) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ SeDectectoDemoraViaje) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ SeDectectoDemoraViaje) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ SeDectectoDemoraViaje) SetString(v string)   { panic("Unsupported operation") }
func (_ SeDectectoDemoraViaje) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *SeDectectoDemoraViaje) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *SeDectectoDemoraViaje) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *SeDectectoDemoraViaje) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ SeDectectoDemoraViaje) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ SeDectectoDemoraViaje) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ SeDectectoDemoraViaje) HintSize(int)                     { panic("Unsupported operation") }
func (_ SeDectectoDemoraViaje) Finalize()                        {}

func (_ SeDectectoDemoraViaje) AvroCRC64Fingerprint() []byte {
	return []byte(SeDectectoDemoraViajeAvroCRC64Fingerprint)
}

func (r SeDectectoDemoraViaje) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *SeDectectoDemoraViaje) UnmarshalJSON(data []byte) error {
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
