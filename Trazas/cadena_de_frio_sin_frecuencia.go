// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CadenaDeFrioSinFrecuencia.avsc
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

type CadenaDeFrioSinFrecuencia struct {
	Traza Traza `json:"traza"`
}

const CadenaDeFrioSinFrecuenciaAvroCRC64Fingerprint = "\xe6\xe5|\xaam0\xc0\x03"

func NewCadenaDeFrioSinFrecuencia() CadenaDeFrioSinFrecuencia {
	r := CadenaDeFrioSinFrecuencia{}
	r.Traza = NewTraza()

	return r
}

func DeserializeCadenaDeFrioSinFrecuencia(r io.Reader) (CadenaDeFrioSinFrecuencia, error) {
	t := NewCadenaDeFrioSinFrecuencia()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCadenaDeFrioSinFrecuenciaFromSchema(r io.Reader, schema string) (CadenaDeFrioSinFrecuencia, error) {
	t := NewCadenaDeFrioSinFrecuencia()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCadenaDeFrioSinFrecuencia(r CadenaDeFrioSinFrecuencia, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r CadenaDeFrioSinFrecuencia) Serialize(w io.Writer) error {
	return writeCadenaDeFrioSinFrecuencia(r, w)
}

func (r CadenaDeFrioSinFrecuencia) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"LogisticShipmentType\",\"type\":[\"null\",\"string\"]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.CadenaDeFrioSinFrecuencia\",\"type\":\"record\"}"
}

func (r CadenaDeFrioSinFrecuencia) SchemaName() string {
	return "Integracion.Esquemas.Trazas.CadenaDeFrioSinFrecuencia"
}

func (_ CadenaDeFrioSinFrecuencia) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CadenaDeFrioSinFrecuencia) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CadenaDeFrioSinFrecuencia) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CadenaDeFrioSinFrecuencia) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CadenaDeFrioSinFrecuencia) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CadenaDeFrioSinFrecuencia) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CadenaDeFrioSinFrecuencia) SetString(v string)   { panic("Unsupported operation") }
func (_ CadenaDeFrioSinFrecuencia) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CadenaDeFrioSinFrecuencia) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *CadenaDeFrioSinFrecuencia) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *CadenaDeFrioSinFrecuencia) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ CadenaDeFrioSinFrecuencia) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CadenaDeFrioSinFrecuencia) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CadenaDeFrioSinFrecuencia) HintSize(int)                     { panic("Unsupported operation") }
func (_ CadenaDeFrioSinFrecuencia) Finalize()                        {}

func (_ CadenaDeFrioSinFrecuencia) AvroCRC64Fingerprint() []byte {
	return []byte(CadenaDeFrioSinFrecuenciaAvroCRC64Fingerprint)
}

func (r CadenaDeFrioSinFrecuencia) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CadenaDeFrioSinFrecuencia) UnmarshalJSON(data []byte) error {
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
