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

type WosMaquinistaEstadoTareaMv struct {
	CambioEstadoTareaMv CambioEstadoTareaMv `json:"CambioEstadoTareaMv"`
}

const WosMaquinistaEstadoTareaMvAvroCRC64Fingerprint = "\xd9\xf9[\xe1\x16ɏ]"

func NewWosMaquinistaEstadoTareaMv() WosMaquinistaEstadoTareaMv {
	r := WosMaquinistaEstadoTareaMv{}
	r.CambioEstadoTareaMv = NewCambioEstadoTareaMv()

	return r
}

func DeserializeWosMaquinistaEstadoTareaMv(r io.Reader) (WosMaquinistaEstadoTareaMv, error) {
	t := NewWosMaquinistaEstadoTareaMv()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeWosMaquinistaEstadoTareaMvFromSchema(r io.Reader, schema string) (WosMaquinistaEstadoTareaMv, error) {
	t := NewWosMaquinistaEstadoTareaMv()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeWosMaquinistaEstadoTareaMv(r WosMaquinistaEstadoTareaMv, w io.Writer) error {
	var err error
	err = writeCambioEstadoTareaMv(r.CambioEstadoTareaMv, w)
	if err != nil {
		return err
	}
	return err
}

func (r WosMaquinistaEstadoTareaMv) Serialize(w io.Writer) error {
	return writeWosMaquinistaEstadoTareaMv(r, w)
}

func (r WosMaquinistaEstadoTareaMv) Schema() string {
	return "{\"fields\":[{\"name\":\"CambioEstadoTareaMv\",\"type\":{\"fields\":[{\"name\":\"TareaMvId\",\"type\":\"int\"},{\"name\":\"TareaMvIdExterno\",\"type\":\"string\"},{\"name\":\"PlantaOperacionId\",\"type\":\"int\"},{\"name\":\"EstadoId\",\"type\":\"int\"},{\"name\":\"Fecha\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"UsuarioId\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"Usuario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TiempoReal\",\"type\":[\"null\",\"int\"]}],\"name\":\"CambioEstadoTareaMv\",\"namespace\":\"Andreani.WosMaquinista.Events.CambioEstadoTareaMvCommon\",\"type\":\"record\"}}],\"name\":\"Andreani.WosMaquinista.Events.Record.WosMaquinistaEstadoTareaMv\",\"type\":\"record\"}"
}

func (r WosMaquinistaEstadoTareaMv) SchemaName() string {
	return "Andreani.WosMaquinista.Events.Record.WosMaquinistaEstadoTareaMv"
}

func (_ WosMaquinistaEstadoTareaMv) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ WosMaquinistaEstadoTareaMv) SetInt(v int32)       { panic("Unsupported operation") }
func (_ WosMaquinistaEstadoTareaMv) SetLong(v int64)      { panic("Unsupported operation") }
func (_ WosMaquinistaEstadoTareaMv) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ WosMaquinistaEstadoTareaMv) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ WosMaquinistaEstadoTareaMv) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ WosMaquinistaEstadoTareaMv) SetString(v string)   { panic("Unsupported operation") }
func (_ WosMaquinistaEstadoTareaMv) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *WosMaquinistaEstadoTareaMv) Get(i int) types.Field {
	switch i {
	case 0:
		r.CambioEstadoTareaMv = NewCambioEstadoTareaMv()

		w := types.Record{Target: &r.CambioEstadoTareaMv}

		return w

	}
	panic("Unknown field index")
}

func (r *WosMaquinistaEstadoTareaMv) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *WosMaquinistaEstadoTareaMv) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ WosMaquinistaEstadoTareaMv) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ WosMaquinistaEstadoTareaMv) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ WosMaquinistaEstadoTareaMv) HintSize(int)                     { panic("Unsupported operation") }
func (_ WosMaquinistaEstadoTareaMv) Finalize()                        {}

func (_ WosMaquinistaEstadoTareaMv) AvroCRC64Fingerprint() []byte {
	return []byte(WosMaquinistaEstadoTareaMvAvroCRC64Fingerprint)
}

func (r WosMaquinistaEstadoTareaMv) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["CambioEstadoTareaMv"], err = json.Marshal(r.CambioEstadoTareaMv)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *WosMaquinistaEstadoTareaMv) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["CambioEstadoTareaMv"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CambioEstadoTareaMv); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CambioEstadoTareaMv")
	}
	return nil
}
