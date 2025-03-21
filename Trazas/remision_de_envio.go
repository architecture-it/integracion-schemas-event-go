// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     RemisionDeEnvio.avsc
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

type RemisionDeEnvio struct {
	Traza Traza `json:"traza"`

	NuevaSucursalDestino *UnionNullDatosSucursal `json:"nuevaSucursalDestino"`

	EsControlado *UnionNullBool `json:"esControlado"`
}

const RemisionDeEnvioAvroCRC64Fingerprint = "\xccQ\xdc\xfb\xf4g\x90\xe0"

func NewRemisionDeEnvio() RemisionDeEnvio {
	r := RemisionDeEnvio{}
	r.Traza = NewTraza()

	r.NuevaSucursalDestino = nil
	r.EsControlado = nil
	return r
}

func DeserializeRemisionDeEnvio(r io.Reader) (RemisionDeEnvio, error) {
	t := NewRemisionDeEnvio()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRemisionDeEnvioFromSchema(r io.Reader, schema string) (RemisionDeEnvio, error) {
	t := NewRemisionDeEnvio()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRemisionDeEnvio(r RemisionDeEnvio, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDatosSucursal(r.NuevaSucursalDestino, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.EsControlado, w)
	if err != nil {
		return err
	}
	return err
}

func (r RemisionDeEnvio) Serialize(w io.Writer) error {
	return writeRemisionDeEnvio(r, w)
}

func (r RemisionDeEnvio) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"default\":null,\"name\":\"nuevaSucursalDestino\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosSucursal\"]},{\"default\":null,\"name\":\"esControlado\",\"type\":[\"null\",\"boolean\"]}],\"name\":\"Integracion.Esquemas.Trazas.RemisionDeEnvio\",\"type\":\"record\"}"
}

func (r RemisionDeEnvio) SchemaName() string {
	return "Integracion.Esquemas.Trazas.RemisionDeEnvio"
}

func (_ RemisionDeEnvio) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ RemisionDeEnvio) SetInt(v int32)       { panic("Unsupported operation") }
func (_ RemisionDeEnvio) SetLong(v int64)      { panic("Unsupported operation") }
func (_ RemisionDeEnvio) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ RemisionDeEnvio) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ RemisionDeEnvio) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ RemisionDeEnvio) SetString(v string)   { panic("Unsupported operation") }
func (_ RemisionDeEnvio) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RemisionDeEnvio) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		r.NuevaSucursalDestino = NewUnionNullDatosSucursal()

		return r.NuevaSucursalDestino
	case 2:
		r.EsControlado = NewUnionNullBool()

		return r.EsControlado
	}
	panic("Unknown field index")
}

func (r *RemisionDeEnvio) SetDefault(i int) {
	switch i {
	case 1:
		r.NuevaSucursalDestino = nil
		return
	case 2:
		r.EsControlado = nil
		return
	}
	panic("Unknown field index")
}

func (r *RemisionDeEnvio) NullField(i int) {
	switch i {
	case 1:
		r.NuevaSucursalDestino = nil
		return
	case 2:
		r.EsControlado = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ RemisionDeEnvio) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ RemisionDeEnvio) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ RemisionDeEnvio) HintSize(int)                     { panic("Unsupported operation") }
func (_ RemisionDeEnvio) Finalize()                        {}

func (_ RemisionDeEnvio) AvroCRC64Fingerprint() []byte {
	return []byte(RemisionDeEnvioAvroCRC64Fingerprint)
}

func (r RemisionDeEnvio) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["nuevaSucursalDestino"], err = json.Marshal(r.NuevaSucursalDestino)
	if err != nil {
		return nil, err
	}
	output["esControlado"], err = json.Marshal(r.EsControlado)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *RemisionDeEnvio) UnmarshalJSON(data []byte) error {
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
	val = func() json.RawMessage {
		if v, ok := fields["nuevaSucursalDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NuevaSucursalDestino); err != nil {
			return err
		}
	} else {
		r.NuevaSucursalDestino = NewUnionNullDatosSucursal()

		r.NuevaSucursalDestino = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["esControlado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EsControlado); err != nil {
			return err
		}
	} else {
		r.EsControlado = NewUnionNullBool()

		r.EsControlado = nil
	}
	return nil
}
