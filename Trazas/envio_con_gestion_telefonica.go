// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnvioConGestionTelefonica.avsc
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

type EnvioConGestionTelefonica struct {
	Traza Traza `json:"traza"`
}

const EnvioConGestionTelefonicaAvroCRC64Fingerprint = "54oE=\x96<\xca"

func NewEnvioConGestionTelefonica() EnvioConGestionTelefonica {
	r := EnvioConGestionTelefonica{}
	r.Traza = NewTraza()

	return r
}

func DeserializeEnvioConGestionTelefonica(r io.Reader) (EnvioConGestionTelefonica, error) {
	t := NewEnvioConGestionTelefonica()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnvioConGestionTelefonicaFromSchema(r io.Reader, schema string) (EnvioConGestionTelefonica, error) {
	t := NewEnvioConGestionTelefonica()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnvioConGestionTelefonica(r EnvioConGestionTelefonica, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r EnvioConGestionTelefonica) Serialize(w io.Writer) error {
	return writeEnvioConGestionTelefonica(r, w)
}

func (r EnvioConGestionTelefonica) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"LogisticShipmentType\",\"type\":[\"null\",\"string\"]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.EnvioConGestionTelefonica\",\"type\":\"record\"}"
}

func (r EnvioConGestionTelefonica) SchemaName() string {
	return "Integracion.Esquemas.Trazas.EnvioConGestionTelefonica"
}

func (_ EnvioConGestionTelefonica) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EnvioConGestionTelefonica) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EnvioConGestionTelefonica) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EnvioConGestionTelefonica) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EnvioConGestionTelefonica) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EnvioConGestionTelefonica) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EnvioConGestionTelefonica) SetString(v string)   { panic("Unsupported operation") }
func (_ EnvioConGestionTelefonica) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EnvioConGestionTelefonica) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *EnvioConGestionTelefonica) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EnvioConGestionTelefonica) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EnvioConGestionTelefonica) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EnvioConGestionTelefonica) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EnvioConGestionTelefonica) HintSize(int)                     { panic("Unsupported operation") }
func (_ EnvioConGestionTelefonica) Finalize()                        {}

func (_ EnvioConGestionTelefonica) AvroCRC64Fingerprint() []byte {
	return []byte(EnvioConGestionTelefonicaAvroCRC64Fingerprint)
}

func (r EnvioConGestionTelefonica) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EnvioConGestionTelefonica) UnmarshalJSON(data []byte) error {
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
