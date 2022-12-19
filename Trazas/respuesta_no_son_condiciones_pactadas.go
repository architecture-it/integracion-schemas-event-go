// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     RespuestaNoSonCondicionesPactadas.avsc
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

type RespuestaNoSonCondicionesPactadas struct {
	Respuesta Respuesta `json:"respuesta"`

	Traza Traza `json:"traza"`
}

const RespuestaNoSonCondicionesPactadasAvroCRC64Fingerprint = "\xef\x8f\xde@\xc1:\x01b"

func NewRespuestaNoSonCondicionesPactadas() RespuestaNoSonCondicionesPactadas {
	r := RespuestaNoSonCondicionesPactadas{}
	r.Respuesta = NewRespuesta()

	r.Traza = NewTraza()

	return r
}

func DeserializeRespuestaNoSonCondicionesPactadas(r io.Reader) (RespuestaNoSonCondicionesPactadas, error) {
	t := NewRespuestaNoSonCondicionesPactadas()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRespuestaNoSonCondicionesPactadasFromSchema(r io.Reader, schema string) (RespuestaNoSonCondicionesPactadas, error) {
	t := NewRespuestaNoSonCondicionesPactadas()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRespuestaNoSonCondicionesPactadas(r RespuestaNoSonCondicionesPactadas, w io.Writer) error {
	var err error
	err = writeRespuesta(r.Respuesta, w)
	if err != nil {
		return err
	}
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	return err
}

func (r RespuestaNoSonCondicionesPactadas) Serialize(w io.Writer) error {
	return writeRespuestaNoSonCondicionesPactadas(r, w)
}

func (r RespuestaNoSonCondicionesPactadas) Schema() string {
	return "{\"fields\":[{\"name\":\"respuesta\",\"type\":{\"fields\":[{\"name\":\"fecha\",\"type\":\"string\"},{\"name\":\"respondeA\",\"type\":\"string\"},{\"name\":\"comentario\",\"type\":\"string\"},{\"name\":\"idNovedad\",\"type\":\"string\"},{\"name\":\"numeroDeOrden\",\"type\":\"string\"},{\"name\":\"cierraPregunta\",\"type\":[\"null\",\"boolean\"]},{\"name\":\"esParaCliente\",\"type\":[\"null\",\"boolean\"]}],\"name\":\"Respuesta\",\"namespace\":\"Integracion.Esquemas.Respuestas\",\"type\":\"record\"}},{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.RespuestaNoSonCondicionesPactadas\",\"type\":\"record\"}"
}

func (r RespuestaNoSonCondicionesPactadas) SchemaName() string {
	return "Integracion.Esquemas.Trazas.RespuestaNoSonCondicionesPactadas"
}

func (_ RespuestaNoSonCondicionesPactadas) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ RespuestaNoSonCondicionesPactadas) SetInt(v int32)       { panic("Unsupported operation") }
func (_ RespuestaNoSonCondicionesPactadas) SetLong(v int64)      { panic("Unsupported operation") }
func (_ RespuestaNoSonCondicionesPactadas) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ RespuestaNoSonCondicionesPactadas) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ RespuestaNoSonCondicionesPactadas) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ RespuestaNoSonCondicionesPactadas) SetString(v string)   { panic("Unsupported operation") }
func (_ RespuestaNoSonCondicionesPactadas) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RespuestaNoSonCondicionesPactadas) Get(i int) types.Field {
	switch i {
	case 0:
		r.Respuesta = NewRespuesta()

		w := types.Record{Target: &r.Respuesta}

		return w

	case 1:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	}
	panic("Unknown field index")
}

func (r *RespuestaNoSonCondicionesPactadas) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *RespuestaNoSonCondicionesPactadas) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ RespuestaNoSonCondicionesPactadas) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ RespuestaNoSonCondicionesPactadas) AppendArray() types.Field { panic("Unsupported operation") }
func (_ RespuestaNoSonCondicionesPactadas) HintSize(int)             { panic("Unsupported operation") }
func (_ RespuestaNoSonCondicionesPactadas) Finalize()                {}

func (_ RespuestaNoSonCondicionesPactadas) AvroCRC64Fingerprint() []byte {
	return []byte(RespuestaNoSonCondicionesPactadasAvroCRC64Fingerprint)
}

func (r RespuestaNoSonCondicionesPactadas) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["respuesta"], err = json.Marshal(r.Respuesta)
	if err != nil {
		return nil, err
	}
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *RespuestaNoSonCondicionesPactadas) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["respuesta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Respuesta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for respuesta")
	}
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
