// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnvioRecanalizado.avsc
 */
package SppSchemaEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EnvioRecanalizado struct {
	Traza EventoDeTraza `json:"traza"`

	NuevaSucursalDeDistribucion *UnionNullDatosSucursal `json:"nuevaSucursalDeDistribucion"`
}

const EnvioRecanalizadoAvroCRC64Fingerprint = "m\xc0J;\xf8\xf0\\\x99"

func NewEnvioRecanalizado() EnvioRecanalizado {
	r := EnvioRecanalizado{}
	r.Traza = NewEventoDeTraza()

	return r
}

func DeserializeEnvioRecanalizado(r io.Reader) (EnvioRecanalizado, error) {
	t := NewEnvioRecanalizado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnvioRecanalizadoFromSchema(r io.Reader, schema string) (EnvioRecanalizado, error) {
	t := NewEnvioRecanalizado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnvioRecanalizado(r EnvioRecanalizado, w io.Writer) error {
	var err error
	err = writeEventoDeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDatosSucursal(r.NuevaSucursalDeDistribucion, w)
	if err != nil {
		return err
	}
	return err
}

func (r EnvioRecanalizado) Serialize(w io.Writer) error {
	return writeEnvioRecanalizado(r, w)
}

func (r EnvioRecanalizado) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":\"string\"},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":[\"null\",\"string\"]},{\"name\":\"nombre\",\"type\":[\"null\",\"string\"]}],\"name\":\"DatosSucursal\",\"type\":\"record\"}]}],\"name\":\"EventoDeTraza\",\"type\":\"record\"}},{\"name\":\"nuevaSucursalDeDistribucion\",\"type\":[\"null\",\"Andreani.SppSchema.Events.DatosSucursal\"]}],\"name\":\"Andreani.SppSchema.Events.EnvioRecanalizado\",\"type\":\"record\"}"
}

func (r EnvioRecanalizado) SchemaName() string {
	return "Andreani.SppSchema.Events.EnvioRecanalizado"
}

func (_ EnvioRecanalizado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EnvioRecanalizado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EnvioRecanalizado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EnvioRecanalizado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EnvioRecanalizado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EnvioRecanalizado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EnvioRecanalizado) SetString(v string)   { panic("Unsupported operation") }
func (_ EnvioRecanalizado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EnvioRecanalizado) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewEventoDeTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		r.NuevaSucursalDeDistribucion = NewUnionNullDatosSucursal()

		return r.NuevaSucursalDeDistribucion
	}
	panic("Unknown field index")
}

func (r *EnvioRecanalizado) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EnvioRecanalizado) NullField(i int) {
	switch i {
	case 1:
		r.NuevaSucursalDeDistribucion = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ EnvioRecanalizado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EnvioRecanalizado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EnvioRecanalizado) HintSize(int)                     { panic("Unsupported operation") }
func (_ EnvioRecanalizado) Finalize()                        {}

func (_ EnvioRecanalizado) AvroCRC64Fingerprint() []byte {
	return []byte(EnvioRecanalizadoAvroCRC64Fingerprint)
}

func (r EnvioRecanalizado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["nuevaSucursalDeDistribucion"], err = json.Marshal(r.NuevaSucursalDeDistribucion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EnvioRecanalizado) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["nuevaSucursalDeDistribucion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NuevaSucursalDeDistribucion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for nuevaSucursalDeDistribucion")
	}
	return nil
}
