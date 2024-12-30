// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ActualizacionEnvioSolicitada.avsc
 */
package AltaOrdenEnvioEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ActualizacionEnvioSolicitada struct {
	Contrato *UnionNullString `json:"Contrato"`

	CodigoDeEnvio *UnionNullString `json:"CodigoDeEnvio"`

	BultoRequest *UnionNullBultoRequest `json:"BultoRequest"`
}

const ActualizacionEnvioSolicitadaAvroCRC64Fingerprint = "B9{\xe2 \x83W\xd0"

func NewActualizacionEnvioSolicitada() ActualizacionEnvioSolicitada {
	r := ActualizacionEnvioSolicitada{}
	r.Contrato = nil
	r.CodigoDeEnvio = nil
	r.BultoRequest = nil
	return r
}

func DeserializeActualizacionEnvioSolicitada(r io.Reader) (ActualizacionEnvioSolicitada, error) {
	t := NewActualizacionEnvioSolicitada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeActualizacionEnvioSolicitadaFromSchema(r io.Reader, schema string) (ActualizacionEnvioSolicitada, error) {
	t := NewActualizacionEnvioSolicitada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeActualizacionEnvioSolicitada(r ActualizacionEnvioSolicitada, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Contrato, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoDeEnvio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBultoRequest(r.BultoRequest, w)
	if err != nil {
		return err
	}
	return err
}

func (r ActualizacionEnvioSolicitada) Serialize(w io.Writer) error {
	return writeActualizacionEnvioSolicitada(r, w)
}

func (r ActualizacionEnvioSolicitada) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Contrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoDeEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"BultoRequest\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Kilos\",\"type\":\"double\"},{\"name\":\"LargoCm\",\"type\":\"double\"},{\"name\":\"AltoCm\",\"type\":\"double\"},{\"name\":\"AnchoCm\",\"type\":\"double\"},{\"name\":\"VolumenCm\",\"type\":\"double\"},{\"name\":\"ValorDeclaradoSinImpuestos\",\"type\":\"double\"},{\"name\":\"ValorDeclaradoConImpuestos\",\"type\":\"double\"},{\"name\":\"ValorDeclarado\",\"type\":\"double\"},{\"default\":null,\"name\":\"Descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"EAN\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentes\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"numeroAgrupador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesHijos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"numeroHijo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"referencias\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]}],\"name\":\"ComponenteHijo\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"Componentes\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Referencias\",\"type\":[\"null\",{\"items\":\"Andreani.AltaOrdenEnvio.Events.Common.Metadato\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"NumeroDeEnvio\",\"type\":[\"null\",\"string\"]}],\"name\":\"BultoRequest\",\"namespace\":\"Andreani.AltaOrdenEnvio.Events.Common\",\"type\":\"record\"}]}],\"name\":\"Andreani.AltaOrdenEnvio.Events.Record.ActualizacionEnvioSolicitada\",\"type\":\"record\"}"
}

func (r ActualizacionEnvioSolicitada) SchemaName() string {
	return "Andreani.AltaOrdenEnvio.Events.Record.ActualizacionEnvioSolicitada"
}

func (_ ActualizacionEnvioSolicitada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ActualizacionEnvioSolicitada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ActualizacionEnvioSolicitada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ActualizacionEnvioSolicitada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ActualizacionEnvioSolicitada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ActualizacionEnvioSolicitada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ActualizacionEnvioSolicitada) SetString(v string)   { panic("Unsupported operation") }
func (_ ActualizacionEnvioSolicitada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ActualizacionEnvioSolicitada) Get(i int) types.Field {
	switch i {
	case 0:
		r.Contrato = NewUnionNullString()

		return r.Contrato
	case 1:
		r.CodigoDeEnvio = NewUnionNullString()

		return r.CodigoDeEnvio
	case 2:
		r.BultoRequest = NewUnionNullBultoRequest()

		return r.BultoRequest
	}
	panic("Unknown field index")
}

func (r *ActualizacionEnvioSolicitada) SetDefault(i int) {
	switch i {
	case 0:
		r.Contrato = nil
		return
	case 1:
		r.CodigoDeEnvio = nil
		return
	case 2:
		r.BultoRequest = nil
		return
	}
	panic("Unknown field index")
}

func (r *ActualizacionEnvioSolicitada) NullField(i int) {
	switch i {
	case 0:
		r.Contrato = nil
		return
	case 1:
		r.CodigoDeEnvio = nil
		return
	case 2:
		r.BultoRequest = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ActualizacionEnvioSolicitada) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ActualizacionEnvioSolicitada) AppendArray() types.Field { panic("Unsupported operation") }
func (_ ActualizacionEnvioSolicitada) HintSize(int)             { panic("Unsupported operation") }
func (_ ActualizacionEnvioSolicitada) Finalize()                {}

func (_ ActualizacionEnvioSolicitada) AvroCRC64Fingerprint() []byte {
	return []byte(ActualizacionEnvioSolicitadaAvroCRC64Fingerprint)
}

func (r ActualizacionEnvioSolicitada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Contrato"], err = json.Marshal(r.Contrato)
	if err != nil {
		return nil, err
	}
	output["CodigoDeEnvio"], err = json.Marshal(r.CodigoDeEnvio)
	if err != nil {
		return nil, err
	}
	output["BultoRequest"], err = json.Marshal(r.BultoRequest)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ActualizacionEnvioSolicitada) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Contrato"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contrato); err != nil {
			return err
		}
	} else {
		r.Contrato = NewUnionNullString()

		r.Contrato = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeEnvio); err != nil {
			return err
		}
	} else {
		r.CodigoDeEnvio = NewUnionNullString()

		r.CodigoDeEnvio = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["BultoRequest"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BultoRequest); err != nil {
			return err
		}
	} else {
		r.BultoRequest = NewUnionNullBultoRequest()

		r.BultoRequest = nil
	}
	return nil
}
