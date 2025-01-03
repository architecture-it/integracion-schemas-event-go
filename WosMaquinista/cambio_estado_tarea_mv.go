// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     WosMaquinistaEstadoTareaMv.avsc
 */
package WosMaquinistaEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type CambioEstadoTareaMv struct {
	TareaMvId int32 `json:"TareaMvId"`

	HojaDeRutaMvId int32 `json:"HojaDeRutaMvId"`

	TipoEstadoTareaMvId int32 `json:"TipoEstadoTareaMvId"`

	TareaMvIdSce string `json:"TareaMvIdSce"`

	PlantaOperacionId int32 `json:"PlantaOperacionId"`

	Cantidad int32 `json:"Cantidad"`

	TareaInterna *UnionNullBool `json:"TareaInterna"`

	Contenedor *UnionNullString `json:"Contenedor"`

	Observacion *UnionNullString `json:"Observacion"`

	Usuario *UnionNullString `json:"Usuario"`
}

const CambioEstadoTareaMvAvroCRC64Fingerprint = "\xc0\x03@Ԝ{\x06\xb0"

func NewCambioEstadoTareaMv() CambioEstadoTareaMv {
	r := CambioEstadoTareaMv{}
	r.TareaInterna = nil
	r.Contenedor = nil
	r.Observacion = nil
	r.Usuario = nil
	return r
}

func DeserializeCambioEstadoTareaMv(r io.Reader) (CambioEstadoTareaMv, error) {
	t := NewCambioEstadoTareaMv()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCambioEstadoTareaMvFromSchema(r io.Reader, schema string) (CambioEstadoTareaMv, error) {
	t := NewCambioEstadoTareaMv()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCambioEstadoTareaMv(r CambioEstadoTareaMv, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.TareaMvId, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.HojaDeRutaMvId, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.TipoEstadoTareaMvId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TareaMvIdSce, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.PlantaOperacionId, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Cantidad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.TareaInterna, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Contenedor, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Observacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Usuario, w)
	if err != nil {
		return err
	}
	return err
}

func (r CambioEstadoTareaMv) Serialize(w io.Writer) error {
	return writeCambioEstadoTareaMv(r, w)
}

func (r CambioEstadoTareaMv) Schema() string {
	return "{\"fields\":[{\"name\":\"TareaMvId\",\"type\":\"int\"},{\"name\":\"HojaDeRutaMvId\",\"type\":\"int\"},{\"name\":\"TipoEstadoTareaMvId\",\"type\":\"int\"},{\"name\":\"TareaMvIdSce\",\"type\":\"string\"},{\"name\":\"PlantaOperacionId\",\"type\":\"int\"},{\"name\":\"Cantidad\",\"type\":\"int\"},{\"default\":null,\"name\":\"TareaInterna\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Contenedor\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Observacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Usuario\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.WosMaquinista.Events.CambioEstadoTareaMvCommon.CambioEstadoTareaMv\",\"type\":\"record\"}"
}

func (r CambioEstadoTareaMv) SchemaName() string {
	return "Andreani.WosMaquinista.Events.CambioEstadoTareaMvCommon.CambioEstadoTareaMv"
}

func (_ CambioEstadoTareaMv) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CambioEstadoTareaMv) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CambioEstadoTareaMv) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CambioEstadoTareaMv) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CambioEstadoTareaMv) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CambioEstadoTareaMv) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CambioEstadoTareaMv) SetString(v string)   { panic("Unsupported operation") }
func (_ CambioEstadoTareaMv) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CambioEstadoTareaMv) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.TareaMvId}

		return w

	case 1:
		w := types.Int{Target: &r.HojaDeRutaMvId}

		return w

	case 2:
		w := types.Int{Target: &r.TipoEstadoTareaMvId}

		return w

	case 3:
		w := types.String{Target: &r.TareaMvIdSce}

		return w

	case 4:
		w := types.Int{Target: &r.PlantaOperacionId}

		return w

	case 5:
		w := types.Int{Target: &r.Cantidad}

		return w

	case 6:
		r.TareaInterna = NewUnionNullBool()

		return r.TareaInterna
	case 7:
		r.Contenedor = NewUnionNullString()

		return r.Contenedor
	case 8:
		r.Observacion = NewUnionNullString()

		return r.Observacion
	case 9:
		r.Usuario = NewUnionNullString()

		return r.Usuario
	}
	panic("Unknown field index")
}

func (r *CambioEstadoTareaMv) SetDefault(i int) {
	switch i {
	case 6:
		r.TareaInterna = nil
		return
	case 7:
		r.Contenedor = nil
		return
	case 8:
		r.Observacion = nil
		return
	case 9:
		r.Usuario = nil
		return
	}
	panic("Unknown field index")
}

func (r *CambioEstadoTareaMv) NullField(i int) {
	switch i {
	case 6:
		r.TareaInterna = nil
		return
	case 7:
		r.Contenedor = nil
		return
	case 8:
		r.Observacion = nil
		return
	case 9:
		r.Usuario = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ CambioEstadoTareaMv) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CambioEstadoTareaMv) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CambioEstadoTareaMv) HintSize(int)                     { panic("Unsupported operation") }
func (_ CambioEstadoTareaMv) Finalize()                        {}

func (_ CambioEstadoTareaMv) AvroCRC64Fingerprint() []byte {
	return []byte(CambioEstadoTareaMvAvroCRC64Fingerprint)
}

func (r CambioEstadoTareaMv) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["TareaMvId"], err = json.Marshal(r.TareaMvId)
	if err != nil {
		return nil, err
	}
	output["HojaDeRutaMvId"], err = json.Marshal(r.HojaDeRutaMvId)
	if err != nil {
		return nil, err
	}
	output["TipoEstadoTareaMvId"], err = json.Marshal(r.TipoEstadoTareaMvId)
	if err != nil {
		return nil, err
	}
	output["TareaMvIdSce"], err = json.Marshal(r.TareaMvIdSce)
	if err != nil {
		return nil, err
	}
	output["PlantaOperacionId"], err = json.Marshal(r.PlantaOperacionId)
	if err != nil {
		return nil, err
	}
	output["Cantidad"], err = json.Marshal(r.Cantidad)
	if err != nil {
		return nil, err
	}
	output["TareaInterna"], err = json.Marshal(r.TareaInterna)
	if err != nil {
		return nil, err
	}
	output["Contenedor"], err = json.Marshal(r.Contenedor)
	if err != nil {
		return nil, err
	}
	output["Observacion"], err = json.Marshal(r.Observacion)
	if err != nil {
		return nil, err
	}
	output["Usuario"], err = json.Marshal(r.Usuario)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CambioEstadoTareaMv) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["TareaMvId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TareaMvId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TareaMvId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["HojaDeRutaMvId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.HojaDeRutaMvId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for HojaDeRutaMvId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoEstadoTareaMvId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoEstadoTareaMvId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoEstadoTareaMvId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TareaMvIdSce"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TareaMvIdSce); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TareaMvIdSce")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PlantaOperacionId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PlantaOperacionId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PlantaOperacionId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Cantidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cantidad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Cantidad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TareaInterna"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TareaInterna); err != nil {
			return err
		}
	} else {
		r.TareaInterna = NewUnionNullBool()

		r.TareaInterna = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Contenedor"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contenedor); err != nil {
			return err
		}
	} else {
		r.Contenedor = NewUnionNullString()

		r.Contenedor = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Observacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Observacion); err != nil {
			return err
		}
	} else {
		r.Observacion = NewUnionNullString()

		r.Observacion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Usuario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Usuario); err != nil {
			return err
		}
	} else {
		r.Usuario = NewUnionNullString()

		r.Usuario = nil
	}
	return nil
}
