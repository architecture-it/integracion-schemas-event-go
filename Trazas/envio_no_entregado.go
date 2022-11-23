// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnvioNoEntregado.avsc
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

type EnvioNoEntregado struct {
	Traza Traza `json:"traza"`

	Motivo *UnionNullString `json:"motivo"`

	Submotivo *UnionNullString `json:"submotivo"`

	ReemplazadoPorNumeroDeEnvio *UnionNullString `json:"reemplazadoPorNumeroDeEnvio"`
}

const EnvioNoEntregadoAvroCRC64Fingerprint = "\x03W\bǓc\xc9\xe7"

func NewEnvioNoEntregado() EnvioNoEntregado {
	r := EnvioNoEntregado{}
	r.Traza = NewTraza()

	r.Motivo = nil
	r.Submotivo = nil
	r.ReemplazadoPorNumeroDeEnvio = nil
	return r
}

func DeserializeEnvioNoEntregado(r io.Reader) (EnvioNoEntregado, error) {
	t := NewEnvioNoEntregado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnvioNoEntregadoFromSchema(r io.Reader, schema string) (EnvioNoEntregado, error) {
	t := NewEnvioNoEntregado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnvioNoEntregado(r EnvioNoEntregado, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Motivo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Submotivo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ReemplazadoPorNumeroDeEnvio, w)
	if err != nil {
		return err
	}
	return err
}

func (r EnvioNoEntregado) Serialize(w io.Writer) error {
	return writeEnvioNoEntregado(r, w)
}

func (r EnvioNoEntregado) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"default\":null,\"name\":\"motivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"submotivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"reemplazadoPorNumeroDeEnvio\",\"type\":[\"null\",\"string\"]}],\"name\":\"Integracion.Esquemas.Trazas.EnvioNoEntregado\",\"type\":\"record\"}"
}

func (r EnvioNoEntregado) SchemaName() string {
	return "Integracion.Esquemas.Trazas.EnvioNoEntregado"
}

func (_ EnvioNoEntregado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EnvioNoEntregado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EnvioNoEntregado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EnvioNoEntregado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EnvioNoEntregado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EnvioNoEntregado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EnvioNoEntregado) SetString(v string)   { panic("Unsupported operation") }
func (_ EnvioNoEntregado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EnvioNoEntregado) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		r.Motivo = NewUnionNullString()

		return r.Motivo
	case 2:
		r.Submotivo = NewUnionNullString()

		return r.Submotivo
	case 3:
		r.ReemplazadoPorNumeroDeEnvio = NewUnionNullString()

		return r.ReemplazadoPorNumeroDeEnvio
	}
	panic("Unknown field index")
}

func (r *EnvioNoEntregado) SetDefault(i int) {
	switch i {
	case 1:
		r.Motivo = nil
		return
	case 2:
		r.Submotivo = nil
		return
	case 3:
		r.ReemplazadoPorNumeroDeEnvio = nil
		return
	}
	panic("Unknown field index")
}

func (r *EnvioNoEntregado) NullField(i int) {
	switch i {
	case 1:
		r.Motivo = nil
		return
	case 2:
		r.Submotivo = nil
		return
	case 3:
		r.ReemplazadoPorNumeroDeEnvio = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ EnvioNoEntregado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EnvioNoEntregado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EnvioNoEntregado) HintSize(int)                     { panic("Unsupported operation") }
func (_ EnvioNoEntregado) Finalize()                        {}

func (_ EnvioNoEntregado) AvroCRC64Fingerprint() []byte {
	return []byte(EnvioNoEntregadoAvroCRC64Fingerprint)
}

func (r EnvioNoEntregado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["motivo"], err = json.Marshal(r.Motivo)
	if err != nil {
		return nil, err
	}
	output["submotivo"], err = json.Marshal(r.Submotivo)
	if err != nil {
		return nil, err
	}
	output["reemplazadoPorNumeroDeEnvio"], err = json.Marshal(r.ReemplazadoPorNumeroDeEnvio)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EnvioNoEntregado) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["motivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Motivo); err != nil {
			return err
		}
	} else {
		r.Motivo = NewUnionNullString()

		r.Motivo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["submotivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Submotivo); err != nil {
			return err
		}
	} else {
		r.Submotivo = NewUnionNullString()

		r.Submotivo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["reemplazadoPorNumeroDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ReemplazadoPorNumeroDeEnvio); err != nil {
			return err
		}
	} else {
		r.ReemplazadoPorNumeroDeEnvio = NewUnionNullString()

		r.ReemplazadoPorNumeroDeEnvio = nil
	}
	return nil
}
