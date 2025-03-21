// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     IncidenciaCreada.avsc
 */
package IncidenciasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type IncidenciaCreada struct {
	IdOrigen *UnionNullInt `json:"IdOrigen"`

	IdPrioridad *UnionNullInt `json:"IdPrioridad"`

	Denunciante *UnionNullString `json:"Denunciante"`

	Observaciones *UnionNullString `json:"Observaciones"`

	IdOperacion *UnionNullInt `json:"IdOperacion"`

	Propietario *UnionNullString `json:"Propietario"`

	IdMotivoPorPropietario *UnionNullInt `json:"IdMotivoPorPropietario"`

	IdFuenteEvidencia *UnionNullInt `json:"IdFuenteEvidencia"`

	Archivos *UnionNullArrayString `json:"Archivos"`

	Fotos *UnionNullArrayString `json:"Fotos"`

	Parametros *UnionNullArrayParametro `json:"Parametros"`

	Instancia *UnionNullString `json:"Instancia"`

	Almacen *UnionNullString `json:"Almacen"`

	LPN *UnionNullString `json:"LPN"`

	IdProceso *UnionNullInt `json:"IdProceso"`

	IdentificadorExterno *UnionNullString `json:"IdentificadorExterno"`

	IdUnidadOperativa *UnionNullString `json:"IdUnidadOperativa"`
}

const IncidenciaCreadaAvroCRC64Fingerprint = "\xa2\x9a\xb2jޟ1\x12"

func NewIncidenciaCreada() IncidenciaCreada {
	r := IncidenciaCreada{}
	r.IdOrigen = nil
	r.IdPrioridad = nil
	r.Denunciante = nil
	r.Observaciones = nil
	r.IdOperacion = nil
	r.Propietario = nil
	r.IdMotivoPorPropietario = nil
	r.IdFuenteEvidencia = nil
	r.Archivos = nil
	r.Fotos = nil
	r.Parametros = nil
	r.Instancia = nil
	r.Almacen = nil
	r.LPN = nil
	r.IdProceso = nil
	r.IdentificadorExterno = nil
	r.IdUnidadOperativa = nil
	return r
}

func DeserializeIncidenciaCreada(r io.Reader) (IncidenciaCreada, error) {
	t := NewIncidenciaCreada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeIncidenciaCreadaFromSchema(r io.Reader, schema string) (IncidenciaCreada, error) {
	t := NewIncidenciaCreada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeIncidenciaCreada(r IncidenciaCreada, w io.Writer) error {
	var err error
	err = writeUnionNullInt(r.IdOrigen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.IdPrioridad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Denunciante, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Observaciones, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.IdOperacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.IdMotivoPorPropietario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.IdFuenteEvidencia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayString(r.Archivos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayString(r.Fotos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayParametro(r.Parametros, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Instancia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Almacen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LPN, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.IdProceso, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.IdentificadorExterno, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.IdUnidadOperativa, w)
	if err != nil {
		return err
	}
	return err
}

func (r IncidenciaCreada) Serialize(w io.Writer) error {
	return writeIncidenciaCreada(r, w)
}

func (r IncidenciaCreada) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"IdOrigen\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"IdPrioridad\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"Denunciante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Observaciones\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"IdOperacion\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"Propietario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"IdMotivoPorPropietario\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"IdFuenteEvidencia\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"Archivos\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"Fotos\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"Parametros\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"IdParametro\",\"type\":\"int\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Valor\",\"type\":\"string\"}],\"name\":\"Parametro\",\"namespace\":\"Andreani.Incidencias.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"Instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Almacen\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LPN\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"IdProceso\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"IdentificadorExterno\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"IdUnidadOperativa\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.Incidencias.Events.Record.IncidenciaCreada\",\"type\":\"record\"}"
}

func (r IncidenciaCreada) SchemaName() string {
	return "Andreani.Incidencias.Events.Record.IncidenciaCreada"
}

func (_ IncidenciaCreada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ IncidenciaCreada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ IncidenciaCreada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ IncidenciaCreada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ IncidenciaCreada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ IncidenciaCreada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ IncidenciaCreada) SetString(v string)   { panic("Unsupported operation") }
func (_ IncidenciaCreada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *IncidenciaCreada) Get(i int) types.Field {
	switch i {
	case 0:
		r.IdOrigen = NewUnionNullInt()

		return r.IdOrigen
	case 1:
		r.IdPrioridad = NewUnionNullInt()

		return r.IdPrioridad
	case 2:
		r.Denunciante = NewUnionNullString()

		return r.Denunciante
	case 3:
		r.Observaciones = NewUnionNullString()

		return r.Observaciones
	case 4:
		r.IdOperacion = NewUnionNullInt()

		return r.IdOperacion
	case 5:
		r.Propietario = NewUnionNullString()

		return r.Propietario
	case 6:
		r.IdMotivoPorPropietario = NewUnionNullInt()

		return r.IdMotivoPorPropietario
	case 7:
		r.IdFuenteEvidencia = NewUnionNullInt()

		return r.IdFuenteEvidencia
	case 8:
		r.Archivos = NewUnionNullArrayString()

		return r.Archivos
	case 9:
		r.Fotos = NewUnionNullArrayString()

		return r.Fotos
	case 10:
		r.Parametros = NewUnionNullArrayParametro()

		return r.Parametros
	case 11:
		r.Instancia = NewUnionNullString()

		return r.Instancia
	case 12:
		r.Almacen = NewUnionNullString()

		return r.Almacen
	case 13:
		r.LPN = NewUnionNullString()

		return r.LPN
	case 14:
		r.IdProceso = NewUnionNullInt()

		return r.IdProceso
	case 15:
		r.IdentificadorExterno = NewUnionNullString()

		return r.IdentificadorExterno
	case 16:
		r.IdUnidadOperativa = NewUnionNullString()

		return r.IdUnidadOperativa
	}
	panic("Unknown field index")
}

func (r *IncidenciaCreada) SetDefault(i int) {
	switch i {
	case 0:
		r.IdOrigen = nil
		return
	case 1:
		r.IdPrioridad = nil
		return
	case 2:
		r.Denunciante = nil
		return
	case 3:
		r.Observaciones = nil
		return
	case 4:
		r.IdOperacion = nil
		return
	case 5:
		r.Propietario = nil
		return
	case 6:
		r.IdMotivoPorPropietario = nil
		return
	case 7:
		r.IdFuenteEvidencia = nil
		return
	case 8:
		r.Archivos = nil
		return
	case 9:
		r.Fotos = nil
		return
	case 10:
		r.Parametros = nil
		return
	case 11:
		r.Instancia = nil
		return
	case 12:
		r.Almacen = nil
		return
	case 13:
		r.LPN = nil
		return
	case 14:
		r.IdProceso = nil
		return
	case 15:
		r.IdentificadorExterno = nil
		return
	case 16:
		r.IdUnidadOperativa = nil
		return
	}
	panic("Unknown field index")
}

func (r *IncidenciaCreada) NullField(i int) {
	switch i {
	case 0:
		r.IdOrigen = nil
		return
	case 1:
		r.IdPrioridad = nil
		return
	case 2:
		r.Denunciante = nil
		return
	case 3:
		r.Observaciones = nil
		return
	case 4:
		r.IdOperacion = nil
		return
	case 5:
		r.Propietario = nil
		return
	case 6:
		r.IdMotivoPorPropietario = nil
		return
	case 7:
		r.IdFuenteEvidencia = nil
		return
	case 8:
		r.Archivos = nil
		return
	case 9:
		r.Fotos = nil
		return
	case 10:
		r.Parametros = nil
		return
	case 11:
		r.Instancia = nil
		return
	case 12:
		r.Almacen = nil
		return
	case 13:
		r.LPN = nil
		return
	case 14:
		r.IdProceso = nil
		return
	case 15:
		r.IdentificadorExterno = nil
		return
	case 16:
		r.IdUnidadOperativa = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ IncidenciaCreada) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ IncidenciaCreada) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ IncidenciaCreada) HintSize(int)                     { panic("Unsupported operation") }
func (_ IncidenciaCreada) Finalize()                        {}

func (_ IncidenciaCreada) AvroCRC64Fingerprint() []byte {
	return []byte(IncidenciaCreadaAvroCRC64Fingerprint)
}

func (r IncidenciaCreada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["IdOrigen"], err = json.Marshal(r.IdOrigen)
	if err != nil {
		return nil, err
	}
	output["IdPrioridad"], err = json.Marshal(r.IdPrioridad)
	if err != nil {
		return nil, err
	}
	output["Denunciante"], err = json.Marshal(r.Denunciante)
	if err != nil {
		return nil, err
	}
	output["Observaciones"], err = json.Marshal(r.Observaciones)
	if err != nil {
		return nil, err
	}
	output["IdOperacion"], err = json.Marshal(r.IdOperacion)
	if err != nil {
		return nil, err
	}
	output["Propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["IdMotivoPorPropietario"], err = json.Marshal(r.IdMotivoPorPropietario)
	if err != nil {
		return nil, err
	}
	output["IdFuenteEvidencia"], err = json.Marshal(r.IdFuenteEvidencia)
	if err != nil {
		return nil, err
	}
	output["Archivos"], err = json.Marshal(r.Archivos)
	if err != nil {
		return nil, err
	}
	output["Fotos"], err = json.Marshal(r.Fotos)
	if err != nil {
		return nil, err
	}
	output["Parametros"], err = json.Marshal(r.Parametros)
	if err != nil {
		return nil, err
	}
	output["Instancia"], err = json.Marshal(r.Instancia)
	if err != nil {
		return nil, err
	}
	output["Almacen"], err = json.Marshal(r.Almacen)
	if err != nil {
		return nil, err
	}
	output["LPN"], err = json.Marshal(r.LPN)
	if err != nil {
		return nil, err
	}
	output["IdProceso"], err = json.Marshal(r.IdProceso)
	if err != nil {
		return nil, err
	}
	output["IdentificadorExterno"], err = json.Marshal(r.IdentificadorExterno)
	if err != nil {
		return nil, err
	}
	output["IdUnidadOperativa"], err = json.Marshal(r.IdUnidadOperativa)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *IncidenciaCreada) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["IdOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdOrigen); err != nil {
			return err
		}
	} else {
		r.IdOrigen = NewUnionNullInt()

		r.IdOrigen = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["IdPrioridad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdPrioridad); err != nil {
			return err
		}
	} else {
		r.IdPrioridad = NewUnionNullInt()

		r.IdPrioridad = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Denunciante"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Denunciante); err != nil {
			return err
		}
	} else {
		r.Denunciante = NewUnionNullString()

		r.Denunciante = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Observaciones"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Observaciones); err != nil {
			return err
		}
	} else {
		r.Observaciones = NewUnionNullString()

		r.Observaciones = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["IdOperacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdOperacion); err != nil {
			return err
		}
	} else {
		r.IdOperacion = NewUnionNullInt()

		r.IdOperacion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Propietario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Propietario); err != nil {
			return err
		}
	} else {
		r.Propietario = NewUnionNullString()

		r.Propietario = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["IdMotivoPorPropietario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdMotivoPorPropietario); err != nil {
			return err
		}
	} else {
		r.IdMotivoPorPropietario = NewUnionNullInt()

		r.IdMotivoPorPropietario = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["IdFuenteEvidencia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdFuenteEvidencia); err != nil {
			return err
		}
	} else {
		r.IdFuenteEvidencia = NewUnionNullInt()

		r.IdFuenteEvidencia = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Archivos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Archivos); err != nil {
			return err
		}
	} else {
		r.Archivos = NewUnionNullArrayString()

		r.Archivos = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Fotos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Fotos); err != nil {
			return err
		}
	} else {
		r.Fotos = NewUnionNullArrayString()

		r.Fotos = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Parametros"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Parametros); err != nil {
			return err
		}
	} else {
		r.Parametros = NewUnionNullArrayParametro()

		r.Parametros = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Instancia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Instancia); err != nil {
			return err
		}
	} else {
		r.Instancia = NewUnionNullString()

		r.Instancia = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Almacen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Almacen); err != nil {
			return err
		}
	} else {
		r.Almacen = NewUnionNullString()

		r.Almacen = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["LPN"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LPN); err != nil {
			return err
		}
	} else {
		r.LPN = NewUnionNullString()

		r.LPN = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["IdProceso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdProceso); err != nil {
			return err
		}
	} else {
		r.IdProceso = NewUnionNullInt()

		r.IdProceso = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["IdentificadorExterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdentificadorExterno); err != nil {
			return err
		}
	} else {
		r.IdentificadorExterno = NewUnionNullString()

		r.IdentificadorExterno = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["IdUnidadOperativa"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdUnidadOperativa); err != nil {
			return err
		}
	} else {
		r.IdUnidadOperativa = NewUnionNullString()

		r.IdUnidadOperativa = nil
	}
	return nil
}
