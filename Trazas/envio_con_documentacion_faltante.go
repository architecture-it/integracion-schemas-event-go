// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnvioConDocumentacionFaltante.avsc
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

type EnvioConDocumentacionFaltante struct {
	Traza Traza `json:"traza"`
}

const EnvioConDocumentacionFaltanteAvroCRC64Fingerprint = "\xb0\x8e\x19\xf3\xa8\xb0\xfd\xb4"

func NewEnvioConDocumentacionFaltante() EnvioConDocumentacionFaltante {
	r := EnvioConDocumentacionFaltante{}
	r.Traza = NewTraza()

	return r
}

func DeserializeEnvioConDocumentacionFaltante(r io.Reader) (EnvioConDocumentacionFaltante, error) {
	t := NewEnvioConDocumentacionFaltante()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnvioConDocumentacionFaltanteFromSchema(r io.Reader, schema string) (EnvioConDocumentacionFaltante, error) {
	t := NewEnvioConDocumentacionFaltante()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnvioConDocumentacionFaltante(r EnvioConDocumentacionFaltante, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r EnvioConDocumentacionFaltante) Serialize(w io.Writer) error {
	return writeEnvioConDocumentacionFaltante(r, w)
}

func (r EnvioConDocumentacionFaltante) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.EnvioConDocumentacionFaltante\",\"type\":\"record\"}"
}

func (r EnvioConDocumentacionFaltante) SchemaName() string {
	return "Integracion.Esquemas.Trazas.EnvioConDocumentacionFaltante"
}

func (_ EnvioConDocumentacionFaltante) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EnvioConDocumentacionFaltante) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EnvioConDocumentacionFaltante) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EnvioConDocumentacionFaltante) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EnvioConDocumentacionFaltante) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EnvioConDocumentacionFaltante) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EnvioConDocumentacionFaltante) SetString(v string)   { panic("Unsupported operation") }
func (_ EnvioConDocumentacionFaltante) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EnvioConDocumentacionFaltante) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *EnvioConDocumentacionFaltante) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EnvioConDocumentacionFaltante) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EnvioConDocumentacionFaltante) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ EnvioConDocumentacionFaltante) AppendArray() types.Field { panic("Unsupported operation") }
func (_ EnvioConDocumentacionFaltante) HintSize(int)             { panic("Unsupported operation") }
func (_ EnvioConDocumentacionFaltante) Finalize()                {}

func (_ EnvioConDocumentacionFaltante) AvroCRC64Fingerprint() []byte {
	return []byte(EnvioConDocumentacionFaltanteAvroCRC64Fingerprint)
}

func (r EnvioConDocumentacionFaltante) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EnvioConDocumentacionFaltante) UnmarshalJSON(data []byte) error {
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
