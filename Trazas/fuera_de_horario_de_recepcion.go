// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     FueraDeHorarioDeRecepcion.avsc
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

type FueraDeHorarioDeRecepcion struct {
	Traza Traza `json:"traza"`
}

const FueraDeHorarioDeRecepcionAvroCRC64Fingerprint = "\v\x19\xf4!\xf7Ŋ\xfc"

func NewFueraDeHorarioDeRecepcion() FueraDeHorarioDeRecepcion {
	r := FueraDeHorarioDeRecepcion{}
	r.Traza = NewTraza()

	return r
}

func DeserializeFueraDeHorarioDeRecepcion(r io.Reader) (FueraDeHorarioDeRecepcion, error) {
	t := NewFueraDeHorarioDeRecepcion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeFueraDeHorarioDeRecepcionFromSchema(r io.Reader, schema string) (FueraDeHorarioDeRecepcion, error) {
	t := NewFueraDeHorarioDeRecepcion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeFueraDeHorarioDeRecepcion(r FueraDeHorarioDeRecepcion, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r FueraDeHorarioDeRecepcion) Serialize(w io.Writer) error {
	return writeFueraDeHorarioDeRecepcion(r, w)
}

func (r FueraDeHorarioDeRecepcion) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.FueraDeHorarioDeRecepcion\",\"type\":\"record\"}"
}

func (r FueraDeHorarioDeRecepcion) SchemaName() string {
	return "Integracion.Esquemas.Trazas.FueraDeHorarioDeRecepcion"
}

func (_ FueraDeHorarioDeRecepcion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ FueraDeHorarioDeRecepcion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ FueraDeHorarioDeRecepcion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ FueraDeHorarioDeRecepcion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ FueraDeHorarioDeRecepcion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ FueraDeHorarioDeRecepcion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ FueraDeHorarioDeRecepcion) SetString(v string)   { panic("Unsupported operation") }
func (_ FueraDeHorarioDeRecepcion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *FueraDeHorarioDeRecepcion) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *FueraDeHorarioDeRecepcion) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *FueraDeHorarioDeRecepcion) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ FueraDeHorarioDeRecepcion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ FueraDeHorarioDeRecepcion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ FueraDeHorarioDeRecepcion) HintSize(int)                     { panic("Unsupported operation") }
func (_ FueraDeHorarioDeRecepcion) Finalize()                        {}

func (_ FueraDeHorarioDeRecepcion) AvroCRC64Fingerprint() []byte {
	return []byte(FueraDeHorarioDeRecepcionAvroCRC64Fingerprint)
}

func (r FueraDeHorarioDeRecepcion) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *FueraDeHorarioDeRecepcion) UnmarshalJSON(data []byte) error {
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