// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     WosMaquinistaEstadoHojaDeRutaMv.avsc
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

type CambioEstadoHojaDeRutaMv struct {
	HojaDeRutaMvId int32 `json:"HojaDeRutaMvId"`

	TipoEstadoHojaDeRutaMvId int32 `json:"TipoEstadoHojaDeRutaMvId"`

	Observacion *UnionNullString `json:"Observacion"`
}

const CambioEstadoHojaDeRutaMvAvroCRC64Fingerprint = "\x05f\x19\xb6~\xaf'3"

func NewCambioEstadoHojaDeRutaMv() CambioEstadoHojaDeRutaMv {
	r := CambioEstadoHojaDeRutaMv{}
	r.Observacion = nil
	return r
}

func DeserializeCambioEstadoHojaDeRutaMv(r io.Reader) (CambioEstadoHojaDeRutaMv, error) {
	t := NewCambioEstadoHojaDeRutaMv()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCambioEstadoHojaDeRutaMvFromSchema(r io.Reader, schema string) (CambioEstadoHojaDeRutaMv, error) {
	t := NewCambioEstadoHojaDeRutaMv()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCambioEstadoHojaDeRutaMv(r CambioEstadoHojaDeRutaMv, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.HojaDeRutaMvId, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.TipoEstadoHojaDeRutaMvId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Observacion, w)
	if err != nil {
		return err
	}
	return err
}

func (r CambioEstadoHojaDeRutaMv) Serialize(w io.Writer) error {
	return writeCambioEstadoHojaDeRutaMv(r, w)
}

func (r CambioEstadoHojaDeRutaMv) Schema() string {
	return "{\"fields\":[{\"name\":\"HojaDeRutaMvId\",\"type\":\"int\"},{\"name\":\"TipoEstadoHojaDeRutaMvId\",\"type\":\"int\"},{\"default\":null,\"name\":\"Observacion\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.WosMaquinista.Events.CambioEstadoHojaDeRutaMvCommon.CambioEstadoHojaDeRutaMv\",\"type\":\"record\"}"
}

func (r CambioEstadoHojaDeRutaMv) SchemaName() string {
	return "Andreani.WosMaquinista.Events.CambioEstadoHojaDeRutaMvCommon.CambioEstadoHojaDeRutaMv"
}

func (_ CambioEstadoHojaDeRutaMv) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CambioEstadoHojaDeRutaMv) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CambioEstadoHojaDeRutaMv) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CambioEstadoHojaDeRutaMv) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CambioEstadoHojaDeRutaMv) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CambioEstadoHojaDeRutaMv) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CambioEstadoHojaDeRutaMv) SetString(v string)   { panic("Unsupported operation") }
func (_ CambioEstadoHojaDeRutaMv) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CambioEstadoHojaDeRutaMv) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.HojaDeRutaMvId}

		return w

	case 1:
		w := types.Int{Target: &r.TipoEstadoHojaDeRutaMvId}

		return w

	case 2:
		r.Observacion = NewUnionNullString()

		return r.Observacion
	}
	panic("Unknown field index")
}

func (r *CambioEstadoHojaDeRutaMv) SetDefault(i int) {
	switch i {
	case 2:
		r.Observacion = nil
		return
	}
	panic("Unknown field index")
}

func (r *CambioEstadoHojaDeRutaMv) NullField(i int) {
	switch i {
	case 2:
		r.Observacion = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ CambioEstadoHojaDeRutaMv) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CambioEstadoHojaDeRutaMv) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CambioEstadoHojaDeRutaMv) HintSize(int)                     { panic("Unsupported operation") }
func (_ CambioEstadoHojaDeRutaMv) Finalize()                        {}

func (_ CambioEstadoHojaDeRutaMv) AvroCRC64Fingerprint() []byte {
	return []byte(CambioEstadoHojaDeRutaMvAvroCRC64Fingerprint)
}

func (r CambioEstadoHojaDeRutaMv) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["HojaDeRutaMvId"], err = json.Marshal(r.HojaDeRutaMvId)
	if err != nil {
		return nil, err
	}
	output["TipoEstadoHojaDeRutaMvId"], err = json.Marshal(r.TipoEstadoHojaDeRutaMvId)
	if err != nil {
		return nil, err
	}
	output["Observacion"], err = json.Marshal(r.Observacion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CambioEstadoHojaDeRutaMv) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
		if v, ok := fields["TipoEstadoHojaDeRutaMvId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoEstadoHojaDeRutaMvId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoEstadoHojaDeRutaMvId")
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
	return nil
}
