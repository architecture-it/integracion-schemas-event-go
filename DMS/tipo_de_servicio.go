// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TipoDeServicio.avsc
 */
package DMSEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type TipoDeServicio struct {
	TipoDeServicioId int32 `json:"TipoDeServicioId"`

	Nombre string `json:"Nombre"`
}

const TipoDeServicioAvroCRC64Fingerprint = "\xe7\x19\xc4\xcc(\xeb\xba\x1b"

func NewTipoDeServicio() TipoDeServicio {
	r := TipoDeServicio{}
	return r
}

func DeserializeTipoDeServicio(r io.Reader) (TipoDeServicio, error) {
	t := NewTipoDeServicio()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTipoDeServicioFromSchema(r io.Reader, schema string) (TipoDeServicio, error) {
	t := NewTipoDeServicio()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTipoDeServicio(r TipoDeServicio, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.TipoDeServicioId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Nombre, w)
	if err != nil {
		return err
	}
	return err
}

func (r TipoDeServicio) Serialize(w io.Writer) error {
	return writeTipoDeServicio(r, w)
}

func (r TipoDeServicio) Schema() string {
	return "{\"fields\":[{\"name\":\"TipoDeServicioId\",\"type\":\"int\"},{\"name\":\"Nombre\",\"type\":\"string\"}],\"name\":\"Andreani.DMS.Events.Hdr.Common.TipoDeServicio\",\"type\":\"record\"}"
}

func (r TipoDeServicio) SchemaName() string {
	return "Andreani.DMS.Events.Hdr.Common.TipoDeServicio"
}

func (_ TipoDeServicio) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ TipoDeServicio) SetInt(v int32)       { panic("Unsupported operation") }
func (_ TipoDeServicio) SetLong(v int64)      { panic("Unsupported operation") }
func (_ TipoDeServicio) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ TipoDeServicio) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ TipoDeServicio) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ TipoDeServicio) SetString(v string)   { panic("Unsupported operation") }
func (_ TipoDeServicio) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *TipoDeServicio) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.TipoDeServicioId}

		return w

	case 1:
		w := types.String{Target: &r.Nombre}

		return w

	}
	panic("Unknown field index")
}

func (r *TipoDeServicio) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *TipoDeServicio) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ TipoDeServicio) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ TipoDeServicio) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ TipoDeServicio) HintSize(int)                     { panic("Unsupported operation") }
func (_ TipoDeServicio) Finalize()                        {}

func (_ TipoDeServicio) AvroCRC64Fingerprint() []byte {
	return []byte(TipoDeServicioAvroCRC64Fingerprint)
}

func (r TipoDeServicio) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["TipoDeServicioId"], err = json.Marshal(r.TipoDeServicioId)
	if err != nil {
		return nil, err
	}
	output["Nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *TipoDeServicio) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["TipoDeServicioId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeServicioId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoDeServicioId")
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
	return nil
}
