// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListInterface672.avsc
 */
package HCMWorkersEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ListInterface672 struct {
	Interfaces []Interface672Data `json:"interfaces"`
}

const ListInterface672AvroCRC64Fingerprint = "\xb2j\x87\xddm\x99\x9ah"

func NewListInterface672() ListInterface672 {
	r := ListInterface672{}
	r.Interfaces = make([]Interface672Data, 0)

	return r
}

func DeserializeListInterface672(r io.Reader) (ListInterface672, error) {
	t := NewListInterface672()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeListInterface672FromSchema(r io.Reader, schema string) (ListInterface672, error) {
	t := NewListInterface672()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeListInterface672(r ListInterface672, w io.Writer) error {
	var err error
	err = writeArrayInterface672Data(r.Interfaces, w)
	if err != nil {
		return err
	}
	return err
}

func (r ListInterface672) Serialize(w io.Writer) error {
	return writeListInterface672(r, w)
}

func (r ListInterface672) Schema() string {
	return "{\"fields\":[{\"name\":\"interfaces\",\"type\":{\"items\":{\"fields\":[{\"default\":null,\"name\":\"Empresa\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Legajo\",\"type\":[\"null\",\"long\"]},{\"name\":\"Apellido\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"FechaNacimiento\",\"type\":\"string\"},{\"default\":null,\"name\":\"PaisDeNacimiento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Nacionalidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaIngresoAlPais\",\"type\":[\"null\",\"string\"]},{\"name\":\"EstadoCivil\",\"type\":\"string\"},{\"default\":null,\"name\":\"Sexo\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaAlta\",\"type\":\"string\"},{\"name\":\"CodigoDePais\",\"type\":\"int\"},{\"name\":\"TipoDocumento1\",\"type\":\"string\"},{\"name\":\"NroDocumento1\",\"type\":\"long\"},{\"default\":null,\"name\":\"TipoDocumento2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NroDocumento2\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"Email\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Estudia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NivelDeEstudio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Contrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaDeVtoContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Estado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CausaDeBaja\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaDeBaja\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ModeloOrganizacional\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PeriodoPrueba\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaVencimientoDePrueba\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeEstructura1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Estructura1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeEstructura2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Estructura2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeEstructura3\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Estructura3\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeEstructura4\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Estructura4\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeEstructura5\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Estructura5\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeIL1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeIL1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeExpediente1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeIL2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeIL2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeExpediente2\",\"type\":[\"null\",\"string\"]},{\"name\":\"DocumentoPrincipal\",\"type\":\"int\"}],\"name\":\"Interface672Data\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.HCMWorkers.Events.Record.ListInterface672\",\"type\":\"record\"}"
}

func (r ListInterface672) SchemaName() string {
	return "Andreani.HCMWorkers.Events.Record.ListInterface672"
}

func (_ ListInterface672) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ListInterface672) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ListInterface672) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ListInterface672) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ListInterface672) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ListInterface672) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ListInterface672) SetString(v string)   { panic("Unsupported operation") }
func (_ ListInterface672) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ListInterface672) Get(i int) types.Field {
	switch i {
	case 0:
		r.Interfaces = make([]Interface672Data, 0)

		w := ArrayInterface672DataWrapper{Target: &r.Interfaces}

		return w

	}
	panic("Unknown field index")
}

func (r *ListInterface672) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ListInterface672) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ListInterface672) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ListInterface672) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ListInterface672) HintSize(int)                     { panic("Unsupported operation") }
func (_ ListInterface672) Finalize()                        {}

func (_ ListInterface672) AvroCRC64Fingerprint() []byte {
	return []byte(ListInterface672AvroCRC64Fingerprint)
}

func (r ListInterface672) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["interfaces"], err = json.Marshal(r.Interfaces)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ListInterface672) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["interfaces"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Interfaces); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for interfaces")
	}
	return nil
}
