// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     SinFrecuenciaDeViaje.avsc
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

type SinFrecuenciaDeViaje struct {
	Traza Traza `json:"traza"`
}

const SinFrecuenciaDeViajeAvroCRC64Fingerprint = "5\x06\xd8\xfe\x9a\u0087\x11"

func NewSinFrecuenciaDeViaje() SinFrecuenciaDeViaje {
	r := SinFrecuenciaDeViaje{}
	r.Traza = NewTraza()

	return r
}

func DeserializeSinFrecuenciaDeViaje(r io.Reader) (SinFrecuenciaDeViaje, error) {
	t := NewSinFrecuenciaDeViaje()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSinFrecuenciaDeViajeFromSchema(r io.Reader, schema string) (SinFrecuenciaDeViaje, error) {
	t := NewSinFrecuenciaDeViaje()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSinFrecuenciaDeViaje(r SinFrecuenciaDeViaje, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r SinFrecuenciaDeViaje) Serialize(w io.Writer) error {
	return writeSinFrecuenciaDeViaje(r, w)
}

func (r SinFrecuenciaDeViaje) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.SinFrecuenciaDeViaje\",\"type\":\"record\"}"
}

func (r SinFrecuenciaDeViaje) SchemaName() string {
	return "Integracion.Esquemas.Trazas.SinFrecuenciaDeViaje"
}

func (_ SinFrecuenciaDeViaje) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ SinFrecuenciaDeViaje) SetInt(v int32)       { panic("Unsupported operation") }
func (_ SinFrecuenciaDeViaje) SetLong(v int64)      { panic("Unsupported operation") }
func (_ SinFrecuenciaDeViaje) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ SinFrecuenciaDeViaje) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ SinFrecuenciaDeViaje) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ SinFrecuenciaDeViaje) SetString(v string)   { panic("Unsupported operation") }
func (_ SinFrecuenciaDeViaje) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *SinFrecuenciaDeViaje) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *SinFrecuenciaDeViaje) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *SinFrecuenciaDeViaje) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ SinFrecuenciaDeViaje) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ SinFrecuenciaDeViaje) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ SinFrecuenciaDeViaje) HintSize(int)                     { panic("Unsupported operation") }
func (_ SinFrecuenciaDeViaje) Finalize()                        {}

func (_ SinFrecuenciaDeViaje) AvroCRC64Fingerprint() []byte {
	return []byte(SinFrecuenciaDeViajeAvroCRC64Fingerprint)
}

func (r SinFrecuenciaDeViaje) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *SinFrecuenciaDeViaje) UnmarshalJSON(data []byte) error {
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
