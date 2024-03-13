// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TareaMobileConformadaEnGrupo.avsc
 */
package MobileOperacionesProducerEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type TareaMobileConformadaEnGrupo struct {
	Id *UnionNullInt `json:"id"`

	Notificada *UnionNullBool `json:"notificada"`

	Enviada *UnionNullBool `json:"enviada"`

	Verificacion *UnionNullString `json:"verificacion"`

	FechaCierre *UnionNullString `json:"fechaCierre"`

	Tarea *UnionNullInt `json:"tarea"`

	Tareas *UnionNullArrayInt `json:"tareas"`

	FechaGeneracion *UnionNullString `json:"fechaGeneracion"`

	FechaContacto *UnionNullString `json:"fechaContacto"`

	Contenido *UnionNullString `json:"contenido"`

	NumeroSeguimiento *UnionNullString `json:"numeroSeguimiento"`

	Latitud *UnionNullString `json:"latitud"`

	Longitud *UnionNullString `json:"longitud"`

	LatitudDomicilio *UnionNullString `json:"latitudDomicilio"`

	LongitudDomicilio *UnionNullString `json:"longitudDomicilio"`

	FechaEnvio *UnionNullString `json:"fechaEnvio"`

	Contacto *UnionNullBool `json:"contacto"`

	Imagenes *UnionNullString `json:"imagenes"`

	Observaciones *UnionNullArrayString `json:"observaciones"`

	Motivo *UnionNullString `json:"motivo"`

	Motivo_id *UnionNullString `json:"motivo_id"`

	Orientacion *UnionNullString `json:"orientacion"`
}

const TareaMobileConformadaEnGrupoAvroCRC64Fingerprint = "\x80}a\xee\xdc\xff\x99\xe7"

func NewTareaMobileConformadaEnGrupo() TareaMobileConformadaEnGrupo {
	r := TareaMobileConformadaEnGrupo{}
	r.Id = nil
	r.Notificada = nil
	r.Enviada = nil
	r.Verificacion = nil
	r.FechaCierre = nil
	r.Tarea = nil
	r.Tareas = nil
	r.FechaGeneracion = nil
	r.FechaContacto = nil
	r.Contenido = nil
	r.NumeroSeguimiento = nil
	r.Latitud = nil
	r.Longitud = nil
	r.LatitudDomicilio = nil
	r.LongitudDomicilio = nil
	r.FechaEnvio = nil
	r.Contacto = nil
	r.Imagenes = nil
	r.Observaciones = nil
	r.Motivo = nil
	r.Motivo_id = nil
	r.Orientacion = nil
	return r
}

func DeserializeTareaMobileConformadaEnGrupo(r io.Reader) (TareaMobileConformadaEnGrupo, error) {
	t := NewTareaMobileConformadaEnGrupo()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTareaMobileConformadaEnGrupoFromSchema(r io.Reader, schema string) (TareaMobileConformadaEnGrupo, error) {
	t := NewTareaMobileConformadaEnGrupo()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTareaMobileConformadaEnGrupo(r TareaMobileConformadaEnGrupo, w io.Writer) error {
	var err error
	err = writeUnionNullInt(r.Id, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.Notificada, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.Enviada, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Verificacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaCierre, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.Tarea, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayInt(r.Tareas, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaGeneracion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaContacto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Contenido, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroSeguimiento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Latitud, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Longitud, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LatitudDomicilio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LongitudDomicilio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaEnvio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.Contacto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Imagenes, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayString(r.Observaciones, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Motivo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Motivo_id, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Orientacion, w)
	if err != nil {
		return err
	}
	return err
}

func (r TareaMobileConformadaEnGrupo) Serialize(w io.Writer) error {
	return writeTareaMobileConformadaEnGrupo(r, w)
}

func (r TareaMobileConformadaEnGrupo) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"id\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"notificada\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"enviada\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"verificacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaCierre\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tarea\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"tareas\",\"type\":[\"null\",{\"items\":\"int\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"fechaGeneracion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaContacto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contenido\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroSeguimiento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"latitud\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"longitud\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"latitudDomicilio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"longitudDomicilio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contacto\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"imagenes\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"observaciones\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"motivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"motivo_id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"orientacion\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.MobileOperacionesProducer.Events.Record.TareaMobileConformadaEnGrupo\",\"type\":\"record\"}"
}

func (r TareaMobileConformadaEnGrupo) SchemaName() string {
	return "Andreani.MobileOperacionesProducer.Events.Record.TareaMobileConformadaEnGrupo"
}

func (_ TareaMobileConformadaEnGrupo) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ TareaMobileConformadaEnGrupo) SetInt(v int32)       { panic("Unsupported operation") }
func (_ TareaMobileConformadaEnGrupo) SetLong(v int64)      { panic("Unsupported operation") }
func (_ TareaMobileConformadaEnGrupo) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ TareaMobileConformadaEnGrupo) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ TareaMobileConformadaEnGrupo) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ TareaMobileConformadaEnGrupo) SetString(v string)   { panic("Unsupported operation") }
func (_ TareaMobileConformadaEnGrupo) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *TareaMobileConformadaEnGrupo) Get(i int) types.Field {
	switch i {
	case 0:
		r.Id = NewUnionNullInt()

		return r.Id
	case 1:
		r.Notificada = NewUnionNullBool()

		return r.Notificada
	case 2:
		r.Enviada = NewUnionNullBool()

		return r.Enviada
	case 3:
		r.Verificacion = NewUnionNullString()

		return r.Verificacion
	case 4:
		r.FechaCierre = NewUnionNullString()

		return r.FechaCierre
	case 5:
		r.Tarea = NewUnionNullInt()

		return r.Tarea
	case 6:
		r.Tareas = NewUnionNullArrayInt()

		return r.Tareas
	case 7:
		r.FechaGeneracion = NewUnionNullString()

		return r.FechaGeneracion
	case 8:
		r.FechaContacto = NewUnionNullString()

		return r.FechaContacto
	case 9:
		r.Contenido = NewUnionNullString()

		return r.Contenido
	case 10:
		r.NumeroSeguimiento = NewUnionNullString()

		return r.NumeroSeguimiento
	case 11:
		r.Latitud = NewUnionNullString()

		return r.Latitud
	case 12:
		r.Longitud = NewUnionNullString()

		return r.Longitud
	case 13:
		r.LatitudDomicilio = NewUnionNullString()

		return r.LatitudDomicilio
	case 14:
		r.LongitudDomicilio = NewUnionNullString()

		return r.LongitudDomicilio
	case 15:
		r.FechaEnvio = NewUnionNullString()

		return r.FechaEnvio
	case 16:
		r.Contacto = NewUnionNullBool()

		return r.Contacto
	case 17:
		r.Imagenes = NewUnionNullString()

		return r.Imagenes
	case 18:
		r.Observaciones = NewUnionNullArrayString()

		return r.Observaciones
	case 19:
		r.Motivo = NewUnionNullString()

		return r.Motivo
	case 20:
		r.Motivo_id = NewUnionNullString()

		return r.Motivo_id
	case 21:
		r.Orientacion = NewUnionNullString()

		return r.Orientacion
	}
	panic("Unknown field index")
}

func (r *TareaMobileConformadaEnGrupo) SetDefault(i int) {
	switch i {
	case 0:
		r.Id = nil
		return
	case 1:
		r.Notificada = nil
		return
	case 2:
		r.Enviada = nil
		return
	case 3:
		r.Verificacion = nil
		return
	case 4:
		r.FechaCierre = nil
		return
	case 5:
		r.Tarea = nil
		return
	case 6:
		r.Tareas = nil
		return
	case 7:
		r.FechaGeneracion = nil
		return
	case 8:
		r.FechaContacto = nil
		return
	case 9:
		r.Contenido = nil
		return
	case 10:
		r.NumeroSeguimiento = nil
		return
	case 11:
		r.Latitud = nil
		return
	case 12:
		r.Longitud = nil
		return
	case 13:
		r.LatitudDomicilio = nil
		return
	case 14:
		r.LongitudDomicilio = nil
		return
	case 15:
		r.FechaEnvio = nil
		return
	case 16:
		r.Contacto = nil
		return
	case 17:
		r.Imagenes = nil
		return
	case 18:
		r.Observaciones = nil
		return
	case 19:
		r.Motivo = nil
		return
	case 20:
		r.Motivo_id = nil
		return
	case 21:
		r.Orientacion = nil
		return
	}
	panic("Unknown field index")
}

func (r *TareaMobileConformadaEnGrupo) NullField(i int) {
	switch i {
	case 0:
		r.Id = nil
		return
	case 1:
		r.Notificada = nil
		return
	case 2:
		r.Enviada = nil
		return
	case 3:
		r.Verificacion = nil
		return
	case 4:
		r.FechaCierre = nil
		return
	case 5:
		r.Tarea = nil
		return
	case 6:
		r.Tareas = nil
		return
	case 7:
		r.FechaGeneracion = nil
		return
	case 8:
		r.FechaContacto = nil
		return
	case 9:
		r.Contenido = nil
		return
	case 10:
		r.NumeroSeguimiento = nil
		return
	case 11:
		r.Latitud = nil
		return
	case 12:
		r.Longitud = nil
		return
	case 13:
		r.LatitudDomicilio = nil
		return
	case 14:
		r.LongitudDomicilio = nil
		return
	case 15:
		r.FechaEnvio = nil
		return
	case 16:
		r.Contacto = nil
		return
	case 17:
		r.Imagenes = nil
		return
	case 18:
		r.Observaciones = nil
		return
	case 19:
		r.Motivo = nil
		return
	case 20:
		r.Motivo_id = nil
		return
	case 21:
		r.Orientacion = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ TareaMobileConformadaEnGrupo) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ TareaMobileConformadaEnGrupo) AppendArray() types.Field { panic("Unsupported operation") }
func (_ TareaMobileConformadaEnGrupo) HintSize(int)             { panic("Unsupported operation") }
func (_ TareaMobileConformadaEnGrupo) Finalize()                {}

func (_ TareaMobileConformadaEnGrupo) AvroCRC64Fingerprint() []byte {
	return []byte(TareaMobileConformadaEnGrupoAvroCRC64Fingerprint)
}

func (r TareaMobileConformadaEnGrupo) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["notificada"], err = json.Marshal(r.Notificada)
	if err != nil {
		return nil, err
	}
	output["enviada"], err = json.Marshal(r.Enviada)
	if err != nil {
		return nil, err
	}
	output["verificacion"], err = json.Marshal(r.Verificacion)
	if err != nil {
		return nil, err
	}
	output["fechaCierre"], err = json.Marshal(r.FechaCierre)
	if err != nil {
		return nil, err
	}
	output["tarea"], err = json.Marshal(r.Tarea)
	if err != nil {
		return nil, err
	}
	output["tareas"], err = json.Marshal(r.Tareas)
	if err != nil {
		return nil, err
	}
	output["fechaGeneracion"], err = json.Marshal(r.FechaGeneracion)
	if err != nil {
		return nil, err
	}
	output["fechaContacto"], err = json.Marshal(r.FechaContacto)
	if err != nil {
		return nil, err
	}
	output["contenido"], err = json.Marshal(r.Contenido)
	if err != nil {
		return nil, err
	}
	output["numeroSeguimiento"], err = json.Marshal(r.NumeroSeguimiento)
	if err != nil {
		return nil, err
	}
	output["latitud"], err = json.Marshal(r.Latitud)
	if err != nil {
		return nil, err
	}
	output["longitud"], err = json.Marshal(r.Longitud)
	if err != nil {
		return nil, err
	}
	output["latitudDomicilio"], err = json.Marshal(r.LatitudDomicilio)
	if err != nil {
		return nil, err
	}
	output["longitudDomicilio"], err = json.Marshal(r.LongitudDomicilio)
	if err != nil {
		return nil, err
	}
	output["fechaEnvio"], err = json.Marshal(r.FechaEnvio)
	if err != nil {
		return nil, err
	}
	output["contacto"], err = json.Marshal(r.Contacto)
	if err != nil {
		return nil, err
	}
	output["imagenes"], err = json.Marshal(r.Imagenes)
	if err != nil {
		return nil, err
	}
	output["observaciones"], err = json.Marshal(r.Observaciones)
	if err != nil {
		return nil, err
	}
	output["motivo"], err = json.Marshal(r.Motivo)
	if err != nil {
		return nil, err
	}
	output["motivo_id"], err = json.Marshal(r.Motivo_id)
	if err != nil {
		return nil, err
	}
	output["orientacion"], err = json.Marshal(r.Orientacion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *TareaMobileConformadaEnGrupo) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		r.Id = NewUnionNullInt()

		r.Id = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["notificada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Notificada); err != nil {
			return err
		}
	} else {
		r.Notificada = NewUnionNullBool()

		r.Notificada = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["enviada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Enviada); err != nil {
			return err
		}
	} else {
		r.Enviada = NewUnionNullBool()

		r.Enviada = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["verificacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Verificacion); err != nil {
			return err
		}
	} else {
		r.Verificacion = NewUnionNullString()

		r.Verificacion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaCierre"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaCierre); err != nil {
			return err
		}
	} else {
		r.FechaCierre = NewUnionNullString()

		r.FechaCierre = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["tarea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tarea); err != nil {
			return err
		}
	} else {
		r.Tarea = NewUnionNullInt()

		r.Tarea = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["tareas"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tareas); err != nil {
			return err
		}
	} else {
		r.Tareas = NewUnionNullArrayInt()

		r.Tareas = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaGeneracion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaGeneracion); err != nil {
			return err
		}
	} else {
		r.FechaGeneracion = NewUnionNullString()

		r.FechaGeneracion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaContacto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaContacto); err != nil {
			return err
		}
	} else {
		r.FechaContacto = NewUnionNullString()

		r.FechaContacto = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["contenido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contenido); err != nil {
			return err
		}
	} else {
		r.Contenido = NewUnionNullString()

		r.Contenido = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroSeguimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroSeguimiento); err != nil {
			return err
		}
	} else {
		r.NumeroSeguimiento = NewUnionNullString()

		r.NumeroSeguimiento = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["latitud"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Latitud); err != nil {
			return err
		}
	} else {
		r.Latitud = NewUnionNullString()

		r.Latitud = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["longitud"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Longitud); err != nil {
			return err
		}
	} else {
		r.Longitud = NewUnionNullString()

		r.Longitud = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["latitudDomicilio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LatitudDomicilio); err != nil {
			return err
		}
	} else {
		r.LatitudDomicilio = NewUnionNullString()

		r.LatitudDomicilio = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["longitudDomicilio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LongitudDomicilio); err != nil {
			return err
		}
	} else {
		r.LongitudDomicilio = NewUnionNullString()

		r.LongitudDomicilio = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaEnvio); err != nil {
			return err
		}
	} else {
		r.FechaEnvio = NewUnionNullString()

		r.FechaEnvio = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["contacto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contacto); err != nil {
			return err
		}
	} else {
		r.Contacto = NewUnionNullBool()

		r.Contacto = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["imagenes"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Imagenes); err != nil {
			return err
		}
	} else {
		r.Imagenes = NewUnionNullString()

		r.Imagenes = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["observaciones"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Observaciones); err != nil {
			return err
		}
	} else {
		r.Observaciones = NewUnionNullArrayString()

		r.Observaciones = nil
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
		if v, ok := fields["motivo_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Motivo_id); err != nil {
			return err
		}
	} else {
		r.Motivo_id = NewUnionNullString()

		r.Motivo_id = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["orientacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Orientacion); err != nil {
			return err
		}
	} else {
		r.Orientacion = NewUnionNullString()

		r.Orientacion = nil
	}
	return nil
}