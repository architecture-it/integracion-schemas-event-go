// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ObservadoFueraDelRangoHorario.avsc
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

type ObservadoFueraDelRangoHorario struct {
	Traza Traza `json:"traza"`

	SucursalDeRecepcion *UnionNullDatosSucursal `json:"sucursalDeRecepcion"`
}

const ObservadoFueraDelRangoHorarioAvroCRC64Fingerprint = "\xf1q\xf7A\xf5`\xda\xf6"

func NewObservadoFueraDelRangoHorario() ObservadoFueraDelRangoHorario {
	r := ObservadoFueraDelRangoHorario{}
	r.Traza = NewTraza()

	r.SucursalDeRecepcion = nil
	return r
}

func DeserializeObservadoFueraDelRangoHorario(r io.Reader) (ObservadoFueraDelRangoHorario, error) {
	t := NewObservadoFueraDelRangoHorario()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeObservadoFueraDelRangoHorarioFromSchema(r io.Reader, schema string) (ObservadoFueraDelRangoHorario, error) {
	t := NewObservadoFueraDelRangoHorario()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeObservadoFueraDelRangoHorario(r ObservadoFueraDelRangoHorario, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDatosSucursal(r.SucursalDeRecepcion, w)
	if err != nil {
		return err
	}
	return err
}

func (r ObservadoFueraDelRangoHorario) Serialize(w io.Writer) error {
	return writeObservadoFueraDelRangoHorario(r, w)
}

func (r ObservadoFueraDelRangoHorario) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"default\":null,\"name\":\"sucursalDeRecepcion\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosSucursal\"]}],\"name\":\"Integracion.Esquemas.Trazas.ObservadoFueraDelRangoHorario\",\"type\":\"record\"}"
}

func (r ObservadoFueraDelRangoHorario) SchemaName() string {
	return "Integracion.Esquemas.Trazas.ObservadoFueraDelRangoHorario"
}

func (_ ObservadoFueraDelRangoHorario) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ObservadoFueraDelRangoHorario) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ObservadoFueraDelRangoHorario) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ObservadoFueraDelRangoHorario) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ObservadoFueraDelRangoHorario) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ObservadoFueraDelRangoHorario) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ObservadoFueraDelRangoHorario) SetString(v string)   { panic("Unsupported operation") }
func (_ ObservadoFueraDelRangoHorario) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ObservadoFueraDelRangoHorario) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		r.SucursalDeRecepcion = NewUnionNullDatosSucursal()

		return r.SucursalDeRecepcion
	}
	panic("Unknown field index")
}

func (r *ObservadoFueraDelRangoHorario) SetDefault(i int) {
	switch i {
	case 1:
		r.SucursalDeRecepcion = nil
		return
	}
	panic("Unknown field index")
}

func (r *ObservadoFueraDelRangoHorario) NullField(i int) {
	switch i {
	case 1:
		r.SucursalDeRecepcion = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ObservadoFueraDelRangoHorario) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ObservadoFueraDelRangoHorario) AppendArray() types.Field { panic("Unsupported operation") }
func (_ ObservadoFueraDelRangoHorario) HintSize(int)             { panic("Unsupported operation") }
func (_ ObservadoFueraDelRangoHorario) Finalize()                {}

func (_ ObservadoFueraDelRangoHorario) AvroCRC64Fingerprint() []byte {
	return []byte(ObservadoFueraDelRangoHorarioAvroCRC64Fingerprint)
}

func (r ObservadoFueraDelRangoHorario) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["sucursalDeRecepcion"], err = json.Marshal(r.SucursalDeRecepcion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ObservadoFueraDelRangoHorario) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["sucursalDeRecepcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalDeRecepcion); err != nil {
			return err
		}
	} else {
		r.SucursalDeRecepcion = NewUnionNullDatosSucursal()

		r.SucursalDeRecepcion = nil
	}
	return nil
}