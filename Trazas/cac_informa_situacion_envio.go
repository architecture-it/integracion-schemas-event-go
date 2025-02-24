// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CacInformaSituacionEnvio.avsc
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

type CacInformaSituacionEnvio struct {
	Traza Traza `json:"traza"`
}

const CacInformaSituacionEnvioAvroCRC64Fingerprint = "\xa48'<\x80ӄ\xe5"

func NewCacInformaSituacionEnvio() CacInformaSituacionEnvio {
	r := CacInformaSituacionEnvio{}
	r.Traza = NewTraza()

	return r
}

func DeserializeCacInformaSituacionEnvio(r io.Reader) (CacInformaSituacionEnvio, error) {
	t := NewCacInformaSituacionEnvio()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCacInformaSituacionEnvioFromSchema(r io.Reader, schema string) (CacInformaSituacionEnvio, error) {
	t := NewCacInformaSituacionEnvio()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCacInformaSituacionEnvio(r CacInformaSituacionEnvio, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r CacInformaSituacionEnvio) Serialize(w io.Writer) error {
	return writeCacInformaSituacionEnvio(r, w)
}

func (r CacInformaSituacionEnvio) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"LogisticShipmentType\",\"type\":[\"null\",\"string\"]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.CacInformaSituacionEnvio\",\"type\":\"record\"}"
}

func (r CacInformaSituacionEnvio) SchemaName() string {
	return "Integracion.Esquemas.Trazas.CacInformaSituacionEnvio"
}

func (_ CacInformaSituacionEnvio) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CacInformaSituacionEnvio) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CacInformaSituacionEnvio) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CacInformaSituacionEnvio) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CacInformaSituacionEnvio) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CacInformaSituacionEnvio) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CacInformaSituacionEnvio) SetString(v string)   { panic("Unsupported operation") }
func (_ CacInformaSituacionEnvio) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CacInformaSituacionEnvio) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *CacInformaSituacionEnvio) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *CacInformaSituacionEnvio) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ CacInformaSituacionEnvio) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CacInformaSituacionEnvio) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CacInformaSituacionEnvio) HintSize(int)                     { panic("Unsupported operation") }
func (_ CacInformaSituacionEnvio) Finalize()                        {}

func (_ CacInformaSituacionEnvio) AvroCRC64Fingerprint() []byte {
	return []byte(CacInformaSituacionEnvioAvroCRC64Fingerprint)
}

func (r CacInformaSituacionEnvio) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CacInformaSituacionEnvio) UnmarshalJSON(data []byte) error {
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
