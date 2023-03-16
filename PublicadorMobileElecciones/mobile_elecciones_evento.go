// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     MobileEleccionesEvento.avsc
 */
package PublicadorMobileEleccionesEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type MobileEleccionesEvento struct {
	OperacionId int32 `json:"operacionId"`

	TareaId int32 `json:"tareaId"`

	TipoDeServicioNombre *UnionNullString `json:"tipoDeServicioNombre"`

	Contenido *UnionNullArrayPropiedadAdicionalElecciones `json:"contenido"`

	Imagenes *UnionNullArrayPropiedadAdicionalElecciones `json:"imagenes"`

	FechaGeneracion *UnionNullString `json:"fechaGeneracion"`

	FechaEnvio *UnionNullString `json:"fechaEnvio"`

	MotivoCodigo *UnionNullString `json:"motivoCodigo"`
}

const MobileEleccionesEventoAvroCRC64Fingerprint = "\xca\xfc\xb8\x80\x95\xf0\xf16"

func NewMobileEleccionesEvento() MobileEleccionesEvento {
	r := MobileEleccionesEvento{}
	r.TipoDeServicioNombre = nil
	r.Contenido = nil
	r.Imagenes = nil
	r.FechaGeneracion = nil
	r.FechaEnvio = nil
	r.MotivoCodigo = nil
	return r
}

func DeserializeMobileEleccionesEvento(r io.Reader) (MobileEleccionesEvento, error) {
	t := NewMobileEleccionesEvento()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMobileEleccionesEventoFromSchema(r io.Reader, schema string) (MobileEleccionesEvento, error) {
	t := NewMobileEleccionesEvento()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMobileEleccionesEvento(r MobileEleccionesEvento, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.OperacionId, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.TareaId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoDeServicioNombre, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayPropiedadAdicionalElecciones(r.Contenido, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayPropiedadAdicionalElecciones(r.Imagenes, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaGeneracion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaEnvio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.MotivoCodigo, w)
	if err != nil {
		return err
	}
	return err
}

func (r MobileEleccionesEvento) Serialize(w io.Writer) error {
	return writeMobileEleccionesEvento(r, w)
}

func (r MobileEleccionesEvento) Schema() string {
	return "{\"fields\":[{\"name\":\"operacionId\",\"type\":\"int\"},{\"name\":\"tareaId\",\"type\":\"int\"},{\"default\":null,\"name\":\"tipoDeServicioNombre\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contenido\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"key\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"value\",\"type\":[\"null\",\"string\"]}],\"name\":\"PropiedadAdicionalElecciones\",\"namespace\":\"Andreani.PublicadorMobileElecciones.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"imagenes\",\"type\":[\"null\",{\"items\":\"Andreani.PublicadorMobileElecciones.Events.Common.PropiedadAdicionalElecciones\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"fechaGeneracion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"motivoCodigo\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.PublicadorMobileElecciones.Events.Record.MobileEleccionesEvento\",\"type\":\"record\"}"
}

func (r MobileEleccionesEvento) SchemaName() string {
	return "Andreani.PublicadorMobileElecciones.Events.Record.MobileEleccionesEvento"
}

func (_ MobileEleccionesEvento) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ MobileEleccionesEvento) SetInt(v int32)       { panic("Unsupported operation") }
func (_ MobileEleccionesEvento) SetLong(v int64)      { panic("Unsupported operation") }
func (_ MobileEleccionesEvento) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ MobileEleccionesEvento) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ MobileEleccionesEvento) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ MobileEleccionesEvento) SetString(v string)   { panic("Unsupported operation") }
func (_ MobileEleccionesEvento) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MobileEleccionesEvento) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.OperacionId}

		return w

	case 1:
		w := types.Int{Target: &r.TareaId}

		return w

	case 2:
		r.TipoDeServicioNombre = NewUnionNullString()

		return r.TipoDeServicioNombre
	case 3:
		r.Contenido = NewUnionNullArrayPropiedadAdicionalElecciones()

		return r.Contenido
	case 4:
		r.Imagenes = NewUnionNullArrayPropiedadAdicionalElecciones()

		return r.Imagenes
	case 5:
		r.FechaGeneracion = NewUnionNullString()

		return r.FechaGeneracion
	case 6:
		r.FechaEnvio = NewUnionNullString()

		return r.FechaEnvio
	case 7:
		r.MotivoCodigo = NewUnionNullString()

		return r.MotivoCodigo
	}
	panic("Unknown field index")
}

func (r *MobileEleccionesEvento) SetDefault(i int) {
	switch i {
	case 2:
		r.TipoDeServicioNombre = nil
		return
	case 3:
		r.Contenido = nil
		return
	case 4:
		r.Imagenes = nil
		return
	case 5:
		r.FechaGeneracion = nil
		return
	case 6:
		r.FechaEnvio = nil
		return
	case 7:
		r.MotivoCodigo = nil
		return
	}
	panic("Unknown field index")
}

func (r *MobileEleccionesEvento) NullField(i int) {
	switch i {
	case 2:
		r.TipoDeServicioNombre = nil
		return
	case 3:
		r.Contenido = nil
		return
	case 4:
		r.Imagenes = nil
		return
	case 5:
		r.FechaGeneracion = nil
		return
	case 6:
		r.FechaEnvio = nil
		return
	case 7:
		r.MotivoCodigo = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ MobileEleccionesEvento) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ MobileEleccionesEvento) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ MobileEleccionesEvento) HintSize(int)                     { panic("Unsupported operation") }
func (_ MobileEleccionesEvento) Finalize()                        {}

func (_ MobileEleccionesEvento) AvroCRC64Fingerprint() []byte {
	return []byte(MobileEleccionesEventoAvroCRC64Fingerprint)
}

func (r MobileEleccionesEvento) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["operacionId"], err = json.Marshal(r.OperacionId)
	if err != nil {
		return nil, err
	}
	output["tareaId"], err = json.Marshal(r.TareaId)
	if err != nil {
		return nil, err
	}
	output["tipoDeServicioNombre"], err = json.Marshal(r.TipoDeServicioNombre)
	if err != nil {
		return nil, err
	}
	output["contenido"], err = json.Marshal(r.Contenido)
	if err != nil {
		return nil, err
	}
	output["imagenes"], err = json.Marshal(r.Imagenes)
	if err != nil {
		return nil, err
	}
	output["fechaGeneracion"], err = json.Marshal(r.FechaGeneracion)
	if err != nil {
		return nil, err
	}
	output["fechaEnvio"], err = json.Marshal(r.FechaEnvio)
	if err != nil {
		return nil, err
	}
	output["motivoCodigo"], err = json.Marshal(r.MotivoCodigo)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *MobileEleccionesEvento) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["operacionId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OperacionId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for operacionId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tareaId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TareaId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tareaId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoDeServicioNombre"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeServicioNombre); err != nil {
			return err
		}
	} else {
		r.TipoDeServicioNombre = NewUnionNullString()

		r.TipoDeServicioNombre = nil
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
		r.Contenido = NewUnionNullArrayPropiedadAdicionalElecciones()

		r.Contenido = nil
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
		r.Imagenes = NewUnionNullArrayPropiedadAdicionalElecciones()

		r.Imagenes = nil
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
		if v, ok := fields["motivoCodigo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MotivoCodigo); err != nil {
			return err
		}
	} else {
		r.MotivoCodigo = NewUnionNullString()

		r.MotivoCodigo = nil
	}
	return nil
}
