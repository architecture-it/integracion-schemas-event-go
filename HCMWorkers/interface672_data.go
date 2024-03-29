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

type Interface672Data struct {
	Empresa *UnionNullString `json:"Empresa"`

	Legajo *UnionNullLong `json:"Legajo"`

	Apellido string `json:"Apellido"`

	Nombre string `json:"Nombre"`

	FechaNacimiento string `json:"FechaNacimiento"`

	PaisDeNacimiento *UnionNullString `json:"PaisDeNacimiento"`

	Nacionalidad *UnionNullString `json:"Nacionalidad"`

	FechaIngresoAlPais *UnionNullString `json:"FechaIngresoAlPais"`

	EstadoCivil string `json:"EstadoCivil"`

	Sexo *UnionNullString `json:"Sexo"`

	FechaAlta string `json:"FechaAlta"`

	CodigoDePais int32 `json:"CodigoDePais"`

	TipoDocumento1 string `json:"TipoDocumento1"`

	NroDocumento1 int64 `json:"NroDocumento1"`

	TipoDocumento2 *UnionNullString `json:"TipoDocumento2"`

	NroDocumento2 *UnionNullLong `json:"NroDocumento2"`

	Email *UnionNullString `json:"Email"`

	Estudia *UnionNullString `json:"Estudia"`

	NivelDeEstudio *UnionNullString `json:"NivelDeEstudio"`

	Contrato *UnionNullString `json:"Contrato"`

	FechaDeVtoContrato *UnionNullString `json:"FechaDeVtoContrato"`

	Estado *UnionNullString `json:"Estado"`

	CausaDeBaja *UnionNullString `json:"CausaDeBaja"`

	FechaDeBaja *UnionNullString `json:"FechaDeBaja"`

	ModeloOrganizacional *UnionNullString `json:"ModeloOrganizacional"`

	PeriodoPrueba *UnionNullString `json:"PeriodoPrueba"`

	FechaVencimientoDePrueba *UnionNullString `json:"FechaVencimientoDePrueba"`

	TipoDeEstructura1 *UnionNullString `json:"TipoDeEstructura1"`

	Estructura1 *UnionNullString `json:"Estructura1"`

	TipoDeEstructura2 *UnionNullString `json:"TipoDeEstructura2"`

	Estructura2 *UnionNullString `json:"Estructura2"`

	TipoDeEstructura3 *UnionNullString `json:"TipoDeEstructura3"`

	Estructura3 *UnionNullString `json:"Estructura3"`

	TipoDeEstructura4 *UnionNullString `json:"TipoDeEstructura4"`

	Estructura4 *UnionNullString `json:"Estructura4"`

	TipoDeEstructura5 *UnionNullString `json:"TipoDeEstructura5"`

	Estructura5 *UnionNullString `json:"Estructura5"`

	TipoDeIL1 *UnionNullString `json:"TipoDeIL1"`

	NumeroDeIL1 *UnionNullString `json:"NumeroDeIL1"`

	NumeroDeExpediente1 *UnionNullString `json:"NumeroDeExpediente1"`

	TipoDeIL2 *UnionNullString `json:"TipoDeIL2"`

	NumeroDeIL2 *UnionNullString `json:"NumeroDeIL2"`

	NumeroDeExpediente2 *UnionNullString `json:"NumeroDeExpediente2"`

	DocumentoPrincipal int32 `json:"DocumentoPrincipal"`
}

const Interface672DataAvroCRC64Fingerprint = "9\xe4K\xd1\xe4c\xe2\xbf"

func NewInterface672Data() Interface672Data {
	r := Interface672Data{}
	r.Empresa = nil
	r.Legajo = nil
	r.PaisDeNacimiento = nil
	r.Nacionalidad = nil
	r.FechaIngresoAlPais = nil
	r.Sexo = nil
	r.TipoDocumento2 = nil
	r.NroDocumento2 = nil
	r.Email = nil
	r.Estudia = nil
	r.NivelDeEstudio = nil
	r.Contrato = nil
	r.FechaDeVtoContrato = nil
	r.Estado = nil
	r.CausaDeBaja = nil
	r.FechaDeBaja = nil
	r.ModeloOrganizacional = nil
	r.PeriodoPrueba = nil
	r.FechaVencimientoDePrueba = nil
	r.TipoDeEstructura1 = nil
	r.Estructura1 = nil
	r.TipoDeEstructura2 = nil
	r.Estructura2 = nil
	r.TipoDeEstructura3 = nil
	r.Estructura3 = nil
	r.TipoDeEstructura4 = nil
	r.Estructura4 = nil
	r.TipoDeEstructura5 = nil
	r.Estructura5 = nil
	r.TipoDeIL1 = nil
	r.NumeroDeIL1 = nil
	r.NumeroDeExpediente1 = nil
	r.TipoDeIL2 = nil
	r.NumeroDeIL2 = nil
	r.NumeroDeExpediente2 = nil
	return r
}

func DeserializeInterface672Data(r io.Reader) (Interface672Data, error) {
	t := NewInterface672Data()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeInterface672DataFromSchema(r io.Reader, schema string) (Interface672Data, error) {
	t := NewInterface672Data()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeInterface672Data(r Interface672Data, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Empresa, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.Legajo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Apellido, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Nombre, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaNacimiento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.PaisDeNacimiento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Nacionalidad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaIngresoAlPais, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EstadoCivil, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Sexo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaAlta, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.CodigoDePais, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TipoDocumento1, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.NroDocumento1, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoDocumento2, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.NroDocumento2, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Email, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Estudia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NivelDeEstudio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Contrato, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaDeVtoContrato, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Estado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CausaDeBaja, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaDeBaja, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ModeloOrganizacional, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.PeriodoPrueba, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaVencimientoDePrueba, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoDeEstructura1, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Estructura1, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoDeEstructura2, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Estructura2, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoDeEstructura3, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Estructura3, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoDeEstructura4, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Estructura4, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoDeEstructura5, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Estructura5, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoDeIL1, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeIL1, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeExpediente1, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoDeIL2, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeIL2, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeExpediente2, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.DocumentoPrincipal, w)
	if err != nil {
		return err
	}
	return err
}

func (r Interface672Data) Serialize(w io.Writer) error {
	return writeInterface672Data(r, w)
}

func (r Interface672Data) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Empresa\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Legajo\",\"type\":[\"null\",\"long\"]},{\"name\":\"Apellido\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"FechaNacimiento\",\"type\":\"string\"},{\"default\":null,\"name\":\"PaisDeNacimiento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Nacionalidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaIngresoAlPais\",\"type\":[\"null\",\"string\"]},{\"name\":\"EstadoCivil\",\"type\":\"string\"},{\"default\":null,\"name\":\"Sexo\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaAlta\",\"type\":\"string\"},{\"name\":\"CodigoDePais\",\"type\":\"int\"},{\"name\":\"TipoDocumento1\",\"type\":\"string\"},{\"name\":\"NroDocumento1\",\"type\":\"long\"},{\"default\":null,\"name\":\"TipoDocumento2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NroDocumento2\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"Email\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Estudia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NivelDeEstudio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Contrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaDeVtoContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Estado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CausaDeBaja\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaDeBaja\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ModeloOrganizacional\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PeriodoPrueba\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaVencimientoDePrueba\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeEstructura1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Estructura1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeEstructura2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Estructura2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeEstructura3\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Estructura3\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeEstructura4\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Estructura4\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeEstructura5\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Estructura5\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeIL1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeIL1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeExpediente1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeIL2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeIL2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeExpediente2\",\"type\":[\"null\",\"string\"]},{\"name\":\"DocumentoPrincipal\",\"type\":\"int\"}],\"name\":\"Andreani.HCMWorkers.Events.Record.Interface672Data\",\"type\":\"record\"}"
}

func (r Interface672Data) SchemaName() string {
	return "Andreani.HCMWorkers.Events.Record.Interface672Data"
}

func (_ Interface672Data) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Interface672Data) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Interface672Data) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Interface672Data) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Interface672Data) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Interface672Data) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Interface672Data) SetString(v string)   { panic("Unsupported operation") }
func (_ Interface672Data) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Interface672Data) Get(i int) types.Field {
	switch i {
	case 0:
		r.Empresa = NewUnionNullString()

		return r.Empresa
	case 1:
		r.Legajo = NewUnionNullLong()

		return r.Legajo
	case 2:
		w := types.String{Target: &r.Apellido}

		return w

	case 3:
		w := types.String{Target: &r.Nombre}

		return w

	case 4:
		w := types.String{Target: &r.FechaNacimiento}

		return w

	case 5:
		r.PaisDeNacimiento = NewUnionNullString()

		return r.PaisDeNacimiento
	case 6:
		r.Nacionalidad = NewUnionNullString()

		return r.Nacionalidad
	case 7:
		r.FechaIngresoAlPais = NewUnionNullString()

		return r.FechaIngresoAlPais
	case 8:
		w := types.String{Target: &r.EstadoCivil}

		return w

	case 9:
		r.Sexo = NewUnionNullString()

		return r.Sexo
	case 10:
		w := types.String{Target: &r.FechaAlta}

		return w

	case 11:
		w := types.Int{Target: &r.CodigoDePais}

		return w

	case 12:
		w := types.String{Target: &r.TipoDocumento1}

		return w

	case 13:
		w := types.Long{Target: &r.NroDocumento1}

		return w

	case 14:
		r.TipoDocumento2 = NewUnionNullString()

		return r.TipoDocumento2
	case 15:
		r.NroDocumento2 = NewUnionNullLong()

		return r.NroDocumento2
	case 16:
		r.Email = NewUnionNullString()

		return r.Email
	case 17:
		r.Estudia = NewUnionNullString()

		return r.Estudia
	case 18:
		r.NivelDeEstudio = NewUnionNullString()

		return r.NivelDeEstudio
	case 19:
		r.Contrato = NewUnionNullString()

		return r.Contrato
	case 20:
		r.FechaDeVtoContrato = NewUnionNullString()

		return r.FechaDeVtoContrato
	case 21:
		r.Estado = NewUnionNullString()

		return r.Estado
	case 22:
		r.CausaDeBaja = NewUnionNullString()

		return r.CausaDeBaja
	case 23:
		r.FechaDeBaja = NewUnionNullString()

		return r.FechaDeBaja
	case 24:
		r.ModeloOrganizacional = NewUnionNullString()

		return r.ModeloOrganizacional
	case 25:
		r.PeriodoPrueba = NewUnionNullString()

		return r.PeriodoPrueba
	case 26:
		r.FechaVencimientoDePrueba = NewUnionNullString()

		return r.FechaVencimientoDePrueba
	case 27:
		r.TipoDeEstructura1 = NewUnionNullString()

		return r.TipoDeEstructura1
	case 28:
		r.Estructura1 = NewUnionNullString()

		return r.Estructura1
	case 29:
		r.TipoDeEstructura2 = NewUnionNullString()

		return r.TipoDeEstructura2
	case 30:
		r.Estructura2 = NewUnionNullString()

		return r.Estructura2
	case 31:
		r.TipoDeEstructura3 = NewUnionNullString()

		return r.TipoDeEstructura3
	case 32:
		r.Estructura3 = NewUnionNullString()

		return r.Estructura3
	case 33:
		r.TipoDeEstructura4 = NewUnionNullString()

		return r.TipoDeEstructura4
	case 34:
		r.Estructura4 = NewUnionNullString()

		return r.Estructura4
	case 35:
		r.TipoDeEstructura5 = NewUnionNullString()

		return r.TipoDeEstructura5
	case 36:
		r.Estructura5 = NewUnionNullString()

		return r.Estructura5
	case 37:
		r.TipoDeIL1 = NewUnionNullString()

		return r.TipoDeIL1
	case 38:
		r.NumeroDeIL1 = NewUnionNullString()

		return r.NumeroDeIL1
	case 39:
		r.NumeroDeExpediente1 = NewUnionNullString()

		return r.NumeroDeExpediente1
	case 40:
		r.TipoDeIL2 = NewUnionNullString()

		return r.TipoDeIL2
	case 41:
		r.NumeroDeIL2 = NewUnionNullString()

		return r.NumeroDeIL2
	case 42:
		r.NumeroDeExpediente2 = NewUnionNullString()

		return r.NumeroDeExpediente2
	case 43:
		w := types.Int{Target: &r.DocumentoPrincipal}

		return w

	}
	panic("Unknown field index")
}

func (r *Interface672Data) SetDefault(i int) {
	switch i {
	case 0:
		r.Empresa = nil
		return
	case 1:
		r.Legajo = nil
		return
	case 5:
		r.PaisDeNacimiento = nil
		return
	case 6:
		r.Nacionalidad = nil
		return
	case 7:
		r.FechaIngresoAlPais = nil
		return
	case 9:
		r.Sexo = nil
		return
	case 14:
		r.TipoDocumento2 = nil
		return
	case 15:
		r.NroDocumento2 = nil
		return
	case 16:
		r.Email = nil
		return
	case 17:
		r.Estudia = nil
		return
	case 18:
		r.NivelDeEstudio = nil
		return
	case 19:
		r.Contrato = nil
		return
	case 20:
		r.FechaDeVtoContrato = nil
		return
	case 21:
		r.Estado = nil
		return
	case 22:
		r.CausaDeBaja = nil
		return
	case 23:
		r.FechaDeBaja = nil
		return
	case 24:
		r.ModeloOrganizacional = nil
		return
	case 25:
		r.PeriodoPrueba = nil
		return
	case 26:
		r.FechaVencimientoDePrueba = nil
		return
	case 27:
		r.TipoDeEstructura1 = nil
		return
	case 28:
		r.Estructura1 = nil
		return
	case 29:
		r.TipoDeEstructura2 = nil
		return
	case 30:
		r.Estructura2 = nil
		return
	case 31:
		r.TipoDeEstructura3 = nil
		return
	case 32:
		r.Estructura3 = nil
		return
	case 33:
		r.TipoDeEstructura4 = nil
		return
	case 34:
		r.Estructura4 = nil
		return
	case 35:
		r.TipoDeEstructura5 = nil
		return
	case 36:
		r.Estructura5 = nil
		return
	case 37:
		r.TipoDeIL1 = nil
		return
	case 38:
		r.NumeroDeIL1 = nil
		return
	case 39:
		r.NumeroDeExpediente1 = nil
		return
	case 40:
		r.TipoDeIL2 = nil
		return
	case 41:
		r.NumeroDeIL2 = nil
		return
	case 42:
		r.NumeroDeExpediente2 = nil
		return
	}
	panic("Unknown field index")
}

func (r *Interface672Data) NullField(i int) {
	switch i {
	case 0:
		r.Empresa = nil
		return
	case 1:
		r.Legajo = nil
		return
	case 5:
		r.PaisDeNacimiento = nil
		return
	case 6:
		r.Nacionalidad = nil
		return
	case 7:
		r.FechaIngresoAlPais = nil
		return
	case 9:
		r.Sexo = nil
		return
	case 14:
		r.TipoDocumento2 = nil
		return
	case 15:
		r.NroDocumento2 = nil
		return
	case 16:
		r.Email = nil
		return
	case 17:
		r.Estudia = nil
		return
	case 18:
		r.NivelDeEstudio = nil
		return
	case 19:
		r.Contrato = nil
		return
	case 20:
		r.FechaDeVtoContrato = nil
		return
	case 21:
		r.Estado = nil
		return
	case 22:
		r.CausaDeBaja = nil
		return
	case 23:
		r.FechaDeBaja = nil
		return
	case 24:
		r.ModeloOrganizacional = nil
		return
	case 25:
		r.PeriodoPrueba = nil
		return
	case 26:
		r.FechaVencimientoDePrueba = nil
		return
	case 27:
		r.TipoDeEstructura1 = nil
		return
	case 28:
		r.Estructura1 = nil
		return
	case 29:
		r.TipoDeEstructura2 = nil
		return
	case 30:
		r.Estructura2 = nil
		return
	case 31:
		r.TipoDeEstructura3 = nil
		return
	case 32:
		r.Estructura3 = nil
		return
	case 33:
		r.TipoDeEstructura4 = nil
		return
	case 34:
		r.Estructura4 = nil
		return
	case 35:
		r.TipoDeEstructura5 = nil
		return
	case 36:
		r.Estructura5 = nil
		return
	case 37:
		r.TipoDeIL1 = nil
		return
	case 38:
		r.NumeroDeIL1 = nil
		return
	case 39:
		r.NumeroDeExpediente1 = nil
		return
	case 40:
		r.TipoDeIL2 = nil
		return
	case 41:
		r.NumeroDeIL2 = nil
		return
	case 42:
		r.NumeroDeExpediente2 = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Interface672Data) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Interface672Data) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Interface672Data) HintSize(int)                     { panic("Unsupported operation") }
func (_ Interface672Data) Finalize()                        {}

func (_ Interface672Data) AvroCRC64Fingerprint() []byte {
	return []byte(Interface672DataAvroCRC64Fingerprint)
}

func (r Interface672Data) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Empresa"], err = json.Marshal(r.Empresa)
	if err != nil {
		return nil, err
	}
	output["Legajo"], err = json.Marshal(r.Legajo)
	if err != nil {
		return nil, err
	}
	output["Apellido"], err = json.Marshal(r.Apellido)
	if err != nil {
		return nil, err
	}
	output["Nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	output["FechaNacimiento"], err = json.Marshal(r.FechaNacimiento)
	if err != nil {
		return nil, err
	}
	output["PaisDeNacimiento"], err = json.Marshal(r.PaisDeNacimiento)
	if err != nil {
		return nil, err
	}
	output["Nacionalidad"], err = json.Marshal(r.Nacionalidad)
	if err != nil {
		return nil, err
	}
	output["FechaIngresoAlPais"], err = json.Marshal(r.FechaIngresoAlPais)
	if err != nil {
		return nil, err
	}
	output["EstadoCivil"], err = json.Marshal(r.EstadoCivil)
	if err != nil {
		return nil, err
	}
	output["Sexo"], err = json.Marshal(r.Sexo)
	if err != nil {
		return nil, err
	}
	output["FechaAlta"], err = json.Marshal(r.FechaAlta)
	if err != nil {
		return nil, err
	}
	output["CodigoDePais"], err = json.Marshal(r.CodigoDePais)
	if err != nil {
		return nil, err
	}
	output["TipoDocumento1"], err = json.Marshal(r.TipoDocumento1)
	if err != nil {
		return nil, err
	}
	output["NroDocumento1"], err = json.Marshal(r.NroDocumento1)
	if err != nil {
		return nil, err
	}
	output["TipoDocumento2"], err = json.Marshal(r.TipoDocumento2)
	if err != nil {
		return nil, err
	}
	output["NroDocumento2"], err = json.Marshal(r.NroDocumento2)
	if err != nil {
		return nil, err
	}
	output["Email"], err = json.Marshal(r.Email)
	if err != nil {
		return nil, err
	}
	output["Estudia"], err = json.Marshal(r.Estudia)
	if err != nil {
		return nil, err
	}
	output["NivelDeEstudio"], err = json.Marshal(r.NivelDeEstudio)
	if err != nil {
		return nil, err
	}
	output["Contrato"], err = json.Marshal(r.Contrato)
	if err != nil {
		return nil, err
	}
	output["FechaDeVtoContrato"], err = json.Marshal(r.FechaDeVtoContrato)
	if err != nil {
		return nil, err
	}
	output["Estado"], err = json.Marshal(r.Estado)
	if err != nil {
		return nil, err
	}
	output["CausaDeBaja"], err = json.Marshal(r.CausaDeBaja)
	if err != nil {
		return nil, err
	}
	output["FechaDeBaja"], err = json.Marshal(r.FechaDeBaja)
	if err != nil {
		return nil, err
	}
	output["ModeloOrganizacional"], err = json.Marshal(r.ModeloOrganizacional)
	if err != nil {
		return nil, err
	}
	output["PeriodoPrueba"], err = json.Marshal(r.PeriodoPrueba)
	if err != nil {
		return nil, err
	}
	output["FechaVencimientoDePrueba"], err = json.Marshal(r.FechaVencimientoDePrueba)
	if err != nil {
		return nil, err
	}
	output["TipoDeEstructura1"], err = json.Marshal(r.TipoDeEstructura1)
	if err != nil {
		return nil, err
	}
	output["Estructura1"], err = json.Marshal(r.Estructura1)
	if err != nil {
		return nil, err
	}
	output["TipoDeEstructura2"], err = json.Marshal(r.TipoDeEstructura2)
	if err != nil {
		return nil, err
	}
	output["Estructura2"], err = json.Marshal(r.Estructura2)
	if err != nil {
		return nil, err
	}
	output["TipoDeEstructura3"], err = json.Marshal(r.TipoDeEstructura3)
	if err != nil {
		return nil, err
	}
	output["Estructura3"], err = json.Marshal(r.Estructura3)
	if err != nil {
		return nil, err
	}
	output["TipoDeEstructura4"], err = json.Marshal(r.TipoDeEstructura4)
	if err != nil {
		return nil, err
	}
	output["Estructura4"], err = json.Marshal(r.Estructura4)
	if err != nil {
		return nil, err
	}
	output["TipoDeEstructura5"], err = json.Marshal(r.TipoDeEstructura5)
	if err != nil {
		return nil, err
	}
	output["Estructura5"], err = json.Marshal(r.Estructura5)
	if err != nil {
		return nil, err
	}
	output["TipoDeIL1"], err = json.Marshal(r.TipoDeIL1)
	if err != nil {
		return nil, err
	}
	output["NumeroDeIL1"], err = json.Marshal(r.NumeroDeIL1)
	if err != nil {
		return nil, err
	}
	output["NumeroDeExpediente1"], err = json.Marshal(r.NumeroDeExpediente1)
	if err != nil {
		return nil, err
	}
	output["TipoDeIL2"], err = json.Marshal(r.TipoDeIL2)
	if err != nil {
		return nil, err
	}
	output["NumeroDeIL2"], err = json.Marshal(r.NumeroDeIL2)
	if err != nil {
		return nil, err
	}
	output["NumeroDeExpediente2"], err = json.Marshal(r.NumeroDeExpediente2)
	if err != nil {
		return nil, err
	}
	output["DocumentoPrincipal"], err = json.Marshal(r.DocumentoPrincipal)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Interface672Data) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Empresa"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Empresa); err != nil {
			return err
		}
	} else {
		r.Empresa = NewUnionNullString()

		r.Empresa = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Legajo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Legajo); err != nil {
			return err
		}
	} else {
		r.Legajo = NewUnionNullLong()

		r.Legajo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Apellido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Apellido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Apellido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Nombre"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Nombre); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Nombre")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaNacimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaNacimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaNacimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PaisDeNacimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PaisDeNacimiento); err != nil {
			return err
		}
	} else {
		r.PaisDeNacimiento = NewUnionNullString()

		r.PaisDeNacimiento = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Nacionalidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Nacionalidad); err != nil {
			return err
		}
	} else {
		r.Nacionalidad = NewUnionNullString()

		r.Nacionalidad = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaIngresoAlPais"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaIngresoAlPais); err != nil {
			return err
		}
	} else {
		r.FechaIngresoAlPais = NewUnionNullString()

		r.FechaIngresoAlPais = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["EstadoCivil"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoCivil); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EstadoCivil")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Sexo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sexo); err != nil {
			return err
		}
	} else {
		r.Sexo = NewUnionNullString()

		r.Sexo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaAlta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaAlta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaAlta")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoDePais"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDePais); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoDePais")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoDocumento1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDocumento1); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoDocumento1")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NroDocumento1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NroDocumento1); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NroDocumento1")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoDocumento2"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDocumento2); err != nil {
			return err
		}
	} else {
		r.TipoDocumento2 = NewUnionNullString()

		r.TipoDocumento2 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["NroDocumento2"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NroDocumento2); err != nil {
			return err
		}
	} else {
		r.NroDocumento2 = NewUnionNullLong()

		r.NroDocumento2 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Email"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Email); err != nil {
			return err
		}
	} else {
		r.Email = NewUnionNullString()

		r.Email = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Estudia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estudia); err != nil {
			return err
		}
	} else {
		r.Estudia = NewUnionNullString()

		r.Estudia = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["NivelDeEstudio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NivelDeEstudio); err != nil {
			return err
		}
	} else {
		r.NivelDeEstudio = NewUnionNullString()

		r.NivelDeEstudio = nil
	}
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
		if v, ok := fields["FechaDeVtoContrato"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaDeVtoContrato); err != nil {
			return err
		}
	} else {
		r.FechaDeVtoContrato = NewUnionNullString()

		r.FechaDeVtoContrato = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Estado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estado); err != nil {
			return err
		}
	} else {
		r.Estado = NewUnionNullString()

		r.Estado = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["CausaDeBaja"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CausaDeBaja); err != nil {
			return err
		}
	} else {
		r.CausaDeBaja = NewUnionNullString()

		r.CausaDeBaja = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaDeBaja"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaDeBaja); err != nil {
			return err
		}
	} else {
		r.FechaDeBaja = NewUnionNullString()

		r.FechaDeBaja = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ModeloOrganizacional"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ModeloOrganizacional); err != nil {
			return err
		}
	} else {
		r.ModeloOrganizacional = NewUnionNullString()

		r.ModeloOrganizacional = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["PeriodoPrueba"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PeriodoPrueba); err != nil {
			return err
		}
	} else {
		r.PeriodoPrueba = NewUnionNullString()

		r.PeriodoPrueba = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaVencimientoDePrueba"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaVencimientoDePrueba); err != nil {
			return err
		}
	} else {
		r.FechaVencimientoDePrueba = NewUnionNullString()

		r.FechaVencimientoDePrueba = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoDeEstructura1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeEstructura1); err != nil {
			return err
		}
	} else {
		r.TipoDeEstructura1 = NewUnionNullString()

		r.TipoDeEstructura1 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Estructura1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estructura1); err != nil {
			return err
		}
	} else {
		r.Estructura1 = NewUnionNullString()

		r.Estructura1 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoDeEstructura2"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeEstructura2); err != nil {
			return err
		}
	} else {
		r.TipoDeEstructura2 = NewUnionNullString()

		r.TipoDeEstructura2 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Estructura2"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estructura2); err != nil {
			return err
		}
	} else {
		r.Estructura2 = NewUnionNullString()

		r.Estructura2 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoDeEstructura3"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeEstructura3); err != nil {
			return err
		}
	} else {
		r.TipoDeEstructura3 = NewUnionNullString()

		r.TipoDeEstructura3 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Estructura3"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estructura3); err != nil {
			return err
		}
	} else {
		r.Estructura3 = NewUnionNullString()

		r.Estructura3 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoDeEstructura4"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeEstructura4); err != nil {
			return err
		}
	} else {
		r.TipoDeEstructura4 = NewUnionNullString()

		r.TipoDeEstructura4 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Estructura4"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estructura4); err != nil {
			return err
		}
	} else {
		r.Estructura4 = NewUnionNullString()

		r.Estructura4 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoDeEstructura5"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeEstructura5); err != nil {
			return err
		}
	} else {
		r.TipoDeEstructura5 = NewUnionNullString()

		r.TipoDeEstructura5 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Estructura5"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estructura5); err != nil {
			return err
		}
	} else {
		r.Estructura5 = NewUnionNullString()

		r.Estructura5 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoDeIL1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeIL1); err != nil {
			return err
		}
	} else {
		r.TipoDeIL1 = NewUnionNullString()

		r.TipoDeIL1 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeIL1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeIL1); err != nil {
			return err
		}
	} else {
		r.NumeroDeIL1 = NewUnionNullString()

		r.NumeroDeIL1 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeExpediente1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeExpediente1); err != nil {
			return err
		}
	} else {
		r.NumeroDeExpediente1 = NewUnionNullString()

		r.NumeroDeExpediente1 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoDeIL2"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeIL2); err != nil {
			return err
		}
	} else {
		r.TipoDeIL2 = NewUnionNullString()

		r.TipoDeIL2 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeIL2"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeIL2); err != nil {
			return err
		}
	} else {
		r.NumeroDeIL2 = NewUnionNullString()

		r.NumeroDeIL2 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeExpediente2"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeExpediente2); err != nil {
			return err
		}
	} else {
		r.NumeroDeExpediente2 = NewUnionNullString()

		r.NumeroDeExpediente2 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DocumentoPrincipal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DocumentoPrincipal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DocumentoPrincipal")
	}
	return nil
}
