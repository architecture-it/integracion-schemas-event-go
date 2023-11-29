// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListInterface2018.avsc
 */
package HCMInterface2018Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ListInterface2018 struct {
	Interfaces []Interface2018Data `json:"interfaces"`
}

const ListInterface2018AvroCRC64Fingerprint = "\xd9J\xec\xda\xe1\xfc]C"

func NewListInterface2018() ListInterface2018 {
	r := ListInterface2018{}
	r.Interfaces = make([]Interface2018Data, 0)

	return r
}

func DeserializeListInterface2018(r io.Reader) (ListInterface2018, error) {
	t := NewListInterface2018()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeListInterface2018FromSchema(r io.Reader, schema string) (ListInterface2018, error) {
	t := NewListInterface2018()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeListInterface2018(r ListInterface2018, w io.Writer) error {
	var err error
	err = writeArrayInterface2018Data(r.Interfaces, w)
	if err != nil {
		return err
	}
	return err
}

func (r ListInterface2018) Serialize(w io.Writer) error {
	return writeListInterface2018(r, w)
}

func (r ListInterface2018) Schema() string {
	return "{\"fields\":[{\"name\":\"interfaces\",\"type\":{\"items\":{\"fields\":[{\"name\":\"LegajoDelEmpleado\",\"type\":\"long\"},{\"name\":\"Apellido\",\"type\":\"string\"},{\"name\":\"Nombres\",\"type\":\"string\"},{\"name\":\"FechaNacimiento\",\"type\":\"string\"},{\"name\":\"PaisDeNacimiento\",\"type\":\"string\"},{\"name\":\"Nacionalidad\",\"type\":\"string\"},{\"name\":\"EstadoCivil\",\"type\":\"string\"},{\"name\":\"Sexo\",\"type\":\"string\"},{\"name\":\"Parentesco\",\"type\":\"string\"},{\"name\":\"Incapacitado\",\"type\":\"string\"},{\"name\":\"FechaDeDiscapacidad\",\"type\":\"string\"},{\"name\":\"TipoDocumento1\",\"type\":\"string\"},{\"name\":\"NroDocumento1\",\"type\":\"long\"},{\"name\":\"ObraSocial\",\"type\":\"string\"},{\"name\":\"PlanOS\",\"type\":\"string\"},{\"name\":\"AvisarAnteEmergencia\",\"type\":\"string\"},{\"name\":\"PagaSalarioFamiliar\",\"type\":\"string\"},{\"name\":\"Ganancias\",\"type\":\"string\"},{\"name\":\"FechaInicioVinculo\",\"type\":\"string\"},{\"name\":\"FechaVto1\",\"type\":\"string\"},{\"name\":\"ItemDDJJ\",\"type\":\"string\"},{\"name\":\"FechaDesdeDDJJ\",\"type\":\"string\"},{\"name\":\"FechaHastaDDJJ\",\"type\":\"string\"},{\"name\":\"TipoDocumento2\",\"type\":\"string\"},{\"name\":\"Documento2\",\"type\":\"string\"},{\"name\":\"Guarderia\",\"type\":\"string\"},{\"name\":\"FechaVto2\",\"type\":\"string\"},{\"name\":\"Adopcion\",\"type\":\"string\"},{\"name\":\"Estudia\",\"type\":\"string\"},{\"name\":\"PeriodoEscolar\",\"type\":\"string\"},{\"name\":\"NivelDeEstudio\",\"type\":\"string\"},{\"name\":\"GradoAnio\",\"type\":\"string\"},{\"name\":\"Anio\",\"type\":\"string\"},{\"name\":\"FechaDeInicio\",\"type\":\"string\"},{\"name\":\"ConstInicial\",\"type\":\"string\"},{\"name\":\"ConstFinal\",\"type\":\"string\"},{\"name\":\"FechaDocumentacion\",\"type\":\"string\"},{\"name\":\"Acta\",\"type\":\"string\"},{\"name\":\"Tomo\",\"type\":\"string\"},{\"name\":\"Folio\",\"type\":\"string\"},{\"name\":\"Tribunal\",\"type\":\"string\"},{\"name\":\"Juzgado\",\"type\":\"string\"},{\"name\":\"Secretaria\",\"type\":\"string\"},{\"name\":\"Comuna\",\"type\":\"string\"}],\"name\":\"Interface2018Data\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.HCMInterface2018.Events.Record.ListInterface2018\",\"type\":\"record\"}"
}

func (r ListInterface2018) SchemaName() string {
	return "Andreani.HCMInterface2018.Events.Record.ListInterface2018"
}

func (_ ListInterface2018) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ListInterface2018) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ListInterface2018) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ListInterface2018) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ListInterface2018) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ListInterface2018) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ListInterface2018) SetString(v string)   { panic("Unsupported operation") }
func (_ ListInterface2018) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ListInterface2018) Get(i int) types.Field {
	switch i {
	case 0:
		r.Interfaces = make([]Interface2018Data, 0)

		w := ArrayInterface2018DataWrapper{Target: &r.Interfaces}

		return w

	}
	panic("Unknown field index")
}

func (r *ListInterface2018) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ListInterface2018) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ListInterface2018) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ListInterface2018) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ListInterface2018) HintSize(int)                     { panic("Unsupported operation") }
func (_ ListInterface2018) Finalize()                        {}

func (_ ListInterface2018) AvroCRC64Fingerprint() []byte {
	return []byte(ListInterface2018AvroCRC64Fingerprint)
}

func (r ListInterface2018) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["interfaces"], err = json.Marshal(r.Interfaces)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ListInterface2018) UnmarshalJSON(data []byte) error {
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
