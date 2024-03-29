// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     RespuestaZonaDeRiesgo.avsc
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

type Respuesta struct {
	Fecha string `json:"fecha"`

	RespondeA string `json:"respondeA"`

	Comentario string `json:"comentario"`

	IdNovedad string `json:"idNovedad"`

	NumeroDeOrden string `json:"numeroDeOrden"`

	CierraPregunta *UnionNullBool `json:"cierraPregunta"`

	EsParaCliente *UnionNullBool `json:"esParaCliente"`
}

const RespuestaAvroCRC64Fingerprint = "ϧ3\x1ed\x00\xee\xa0"

func NewRespuesta() Respuesta {
	r := Respuesta{}
	return r
}

func DeserializeRespuesta(r io.Reader) (Respuesta, error) {
	t := NewRespuesta()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRespuestaFromSchema(r io.Reader, schema string) (Respuesta, error) {
	t := NewRespuesta()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRespuesta(r Respuesta, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Fecha, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.RespondeA, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Comentario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.IdNovedad, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroDeOrden, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.CierraPregunta, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.EsParaCliente, w)
	if err != nil {
		return err
	}
	return err
}

func (r Respuesta) Serialize(w io.Writer) error {
	return writeRespuesta(r, w)
}

func (r Respuesta) Schema() string {
	return "{\"fields\":[{\"name\":\"fecha\",\"type\":\"string\"},{\"name\":\"respondeA\",\"type\":\"string\"},{\"name\":\"comentario\",\"type\":\"string\"},{\"name\":\"idNovedad\",\"type\":\"string\"},{\"name\":\"numeroDeOrden\",\"type\":\"string\"},{\"name\":\"cierraPregunta\",\"type\":[\"null\",\"boolean\"]},{\"name\":\"esParaCliente\",\"type\":[\"null\",\"boolean\"]}],\"name\":\"Integracion.Esquemas.Respuestas.Respuesta\",\"type\":\"record\"}"
}

func (r Respuesta) SchemaName() string {
	return "Integracion.Esquemas.Respuestas.Respuesta"
}

func (_ Respuesta) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Respuesta) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Respuesta) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Respuesta) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Respuesta) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Respuesta) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Respuesta) SetString(v string)   { panic("Unsupported operation") }
func (_ Respuesta) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Respuesta) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Fecha}

		return w

	case 1:
		w := types.String{Target: &r.RespondeA}

		return w

	case 2:
		w := types.String{Target: &r.Comentario}

		return w

	case 3:
		w := types.String{Target: &r.IdNovedad}

		return w

	case 4:
		w := types.String{Target: &r.NumeroDeOrden}

		return w

	case 5:
		r.CierraPregunta = NewUnionNullBool()

		return r.CierraPregunta
	case 6:
		r.EsParaCliente = NewUnionNullBool()

		return r.EsParaCliente
	}
	panic("Unknown field index")
}

func (r *Respuesta) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Respuesta) NullField(i int) {
	switch i {
	case 5:
		r.CierraPregunta = nil
		return
	case 6:
		r.EsParaCliente = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Respuesta) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Respuesta) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Respuesta) HintSize(int)                     { panic("Unsupported operation") }
func (_ Respuesta) Finalize()                        {}

func (_ Respuesta) AvroCRC64Fingerprint() []byte {
	return []byte(RespuestaAvroCRC64Fingerprint)
}

func (r Respuesta) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["fecha"], err = json.Marshal(r.Fecha)
	if err != nil {
		return nil, err
	}
	output["respondeA"], err = json.Marshal(r.RespondeA)
	if err != nil {
		return nil, err
	}
	output["comentario"], err = json.Marshal(r.Comentario)
	if err != nil {
		return nil, err
	}
	output["idNovedad"], err = json.Marshal(r.IdNovedad)
	if err != nil {
		return nil, err
	}
	output["numeroDeOrden"], err = json.Marshal(r.NumeroDeOrden)
	if err != nil {
		return nil, err
	}
	output["cierraPregunta"], err = json.Marshal(r.CierraPregunta)
	if err != nil {
		return nil, err
	}
	output["esParaCliente"], err = json.Marshal(r.EsParaCliente)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Respuesta) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["fecha"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Fecha); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fecha")
	}
	val = func() json.RawMessage {
		if v, ok := fields["respondeA"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RespondeA); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for respondeA")
	}
	val = func() json.RawMessage {
		if v, ok := fields["comentario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Comentario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for comentario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["idNovedad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdNovedad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idNovedad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroDeOrden"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeOrden); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroDeOrden")
	}
	val = func() json.RawMessage {
		if v, ok := fields["cierraPregunta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CierraPregunta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cierraPregunta")
	}
	val = func() json.RawMessage {
		if v, ok := fields["esParaCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EsParaCliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for esParaCliente")
	}
	return nil
}
