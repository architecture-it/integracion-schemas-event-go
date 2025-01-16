// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ExpedicionHojaDeRutaCabecera.avsc
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

type ExpedicionHojaDeRutaCabecera struct {
	Traza Traza `json:"traza"`

	SucursalQueExpide *UnionNullDatosSucursal `json:"sucursalQueExpide"`

	MedioDeExpedicion *UnionNullString `json:"medioDeExpedicion"`
}

const ExpedicionHojaDeRutaCabeceraAvroCRC64Fingerprint = "7\x922!K\n\xa4\x99"

func NewExpedicionHojaDeRutaCabecera() ExpedicionHojaDeRutaCabecera {
	r := ExpedicionHojaDeRutaCabecera{}
	r.Traza = NewTraza()

	r.SucursalQueExpide = nil
	r.MedioDeExpedicion = nil
	return r
}

func DeserializeExpedicionHojaDeRutaCabecera(r io.Reader) (ExpedicionHojaDeRutaCabecera, error) {
	t := NewExpedicionHojaDeRutaCabecera()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeExpedicionHojaDeRutaCabeceraFromSchema(r io.Reader, schema string) (ExpedicionHojaDeRutaCabecera, error) {
	t := NewExpedicionHojaDeRutaCabecera()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeExpedicionHojaDeRutaCabecera(r ExpedicionHojaDeRutaCabecera, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDatosSucursal(r.SucursalQueExpide, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.MedioDeExpedicion, w)
	if err != nil {
		return err
	}
	return err
}

func (r ExpedicionHojaDeRutaCabecera) Serialize(w io.Writer) error {
	return writeExpedicionHojaDeRutaCabecera(r, w)
}

func (r ExpedicionHojaDeRutaCabecera) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"default\":null,\"name\":\"sucursalQueExpide\",\"type\":[\"null\",\"Integracion.Esquemas.Referencias.DatosSucursal\"]},{\"default\":null,\"name\":\"medioDeExpedicion\",\"type\":[\"null\",\"string\"]}],\"name\":\"Integracion.Esquemas.Trazas.ExpedicionHojaDeRutaCabecera\",\"type\":\"record\"}"
}

func (r ExpedicionHojaDeRutaCabecera) SchemaName() string {
	return "Integracion.Esquemas.Trazas.ExpedicionHojaDeRutaCabecera"
}

func (_ ExpedicionHojaDeRutaCabecera) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaCabecera) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaCabecera) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaCabecera) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaCabecera) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaCabecera) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaCabecera) SetString(v string)   { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaCabecera) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ExpedicionHojaDeRutaCabecera) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		r.SucursalQueExpide = NewUnionNullDatosSucursal()

		return r.SucursalQueExpide
	case 2:
		r.MedioDeExpedicion = NewUnionNullString()

		return r.MedioDeExpedicion
	}
	panic("Unknown field index")
}

func (r *ExpedicionHojaDeRutaCabecera) SetDefault(i int) {
	switch i {
	case 1:
		r.SucursalQueExpide = nil
		return
	case 2:
		r.MedioDeExpedicion = nil
		return
	}
	panic("Unknown field index")
}

func (r *ExpedicionHojaDeRutaCabecera) NullField(i int) {
	switch i {
	case 1:
		r.SucursalQueExpide = nil
		return
	case 2:
		r.MedioDeExpedicion = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ExpedicionHojaDeRutaCabecera) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ExpedicionHojaDeRutaCabecera) AppendArray() types.Field { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaCabecera) HintSize(int)             { panic("Unsupported operation") }
func (_ ExpedicionHojaDeRutaCabecera) Finalize()                {}

func (_ ExpedicionHojaDeRutaCabecera) AvroCRC64Fingerprint() []byte {
	return []byte(ExpedicionHojaDeRutaCabeceraAvroCRC64Fingerprint)
}

func (r ExpedicionHojaDeRutaCabecera) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["sucursalQueExpide"], err = json.Marshal(r.SucursalQueExpide)
	if err != nil {
		return nil, err
	}
	output["medioDeExpedicion"], err = json.Marshal(r.MedioDeExpedicion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ExpedicionHojaDeRutaCabecera) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["sucursalQueExpide"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalQueExpide); err != nil {
			return err
		}
	} else {
		r.SucursalQueExpide = NewUnionNullDatosSucursal()

		r.SucursalQueExpide = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["medioDeExpedicion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MedioDeExpedicion); err != nil {
			return err
		}
	} else {
		r.MedioDeExpedicion = NewUnionNullString()

		r.MedioDeExpedicion = nil
	}
	return nil
}
