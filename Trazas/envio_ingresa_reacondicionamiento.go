// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnvioIngresaReacondicionamiento.avsc
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

type EnvioIngresaReacondicionamiento struct {
	Traza Traza `json:"traza"`
}

const EnvioIngresaReacondicionamientoAvroCRC64Fingerprint = "\x8c\xa5\x1d\xc6D\xb4\x88$"

func NewEnvioIngresaReacondicionamiento() EnvioIngresaReacondicionamiento {
	r := EnvioIngresaReacondicionamiento{}
	r.Traza = NewTraza()

	return r
}

func DeserializeEnvioIngresaReacondicionamiento(r io.Reader) (EnvioIngresaReacondicionamiento, error) {
	t := NewEnvioIngresaReacondicionamiento()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnvioIngresaReacondicionamientoFromSchema(r io.Reader, schema string) (EnvioIngresaReacondicionamiento, error) {
	t := NewEnvioIngresaReacondicionamiento()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnvioIngresaReacondicionamiento(r EnvioIngresaReacondicionamiento, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r EnvioIngresaReacondicionamiento) Serialize(w io.Writer) error {
	return writeEnvioIngresaReacondicionamiento(r, w)
}

func (r EnvioIngresaReacondicionamiento) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.EnvioIngresaReacondicionamiento\",\"type\":\"record\"}"
}

func (r EnvioIngresaReacondicionamiento) SchemaName() string {
	return "Integracion.Esquemas.Trazas.EnvioIngresaReacondicionamiento"
}

func (_ EnvioIngresaReacondicionamiento) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EnvioIngresaReacondicionamiento) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EnvioIngresaReacondicionamiento) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EnvioIngresaReacondicionamiento) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EnvioIngresaReacondicionamiento) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EnvioIngresaReacondicionamiento) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EnvioIngresaReacondicionamiento) SetString(v string)   { panic("Unsupported operation") }
func (_ EnvioIngresaReacondicionamiento) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EnvioIngresaReacondicionamiento) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *EnvioIngresaReacondicionamiento) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EnvioIngresaReacondicionamiento) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EnvioIngresaReacondicionamiento) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ EnvioIngresaReacondicionamiento) AppendArray() types.Field { panic("Unsupported operation") }
func (_ EnvioIngresaReacondicionamiento) HintSize(int)             { panic("Unsupported operation") }
func (_ EnvioIngresaReacondicionamiento) Finalize()                {}

func (_ EnvioIngresaReacondicionamiento) AvroCRC64Fingerprint() []byte {
	return []byte(EnvioIngresaReacondicionamientoAvroCRC64Fingerprint)
}

func (r EnvioIngresaReacondicionamiento) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EnvioIngresaReacondicionamiento) UnmarshalJSON(data []byte) error {
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