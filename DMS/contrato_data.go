// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     HdrCreada.avsc
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

type ContratoData struct {
	ContratoId string `json:"ContratoId"`

	Motivos *UnionNullArrayString `json:"Motivos"`

	CondicionDeEntregaId *UnionNullString `json:"CondicionDeEntregaId"`
}

const ContratoDataAvroCRC64Fingerprint = "\xc9U\x1c_\xfe\xf0\x98\x1f"

func NewContratoData() ContratoData {
	r := ContratoData{}
	r.Motivos = nil
	r.CondicionDeEntregaId = nil
	return r
}

func DeserializeContratoData(r io.Reader) (ContratoData, error) {
	t := NewContratoData()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeContratoDataFromSchema(r io.Reader, schema string) (ContratoData, error) {
	t := NewContratoData()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeContratoData(r ContratoData, w io.Writer) error {
	var err error
	err = vm.WriteString(r.ContratoId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayString(r.Motivos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CondicionDeEntregaId, w)
	if err != nil {
		return err
	}
	return err
}

func (r ContratoData) Serialize(w io.Writer) error {
	return writeContratoData(r, w)
}

func (r ContratoData) Schema() string {
	return "{\"fields\":[{\"name\":\"ContratoId\",\"type\":\"string\"},{\"default\":null,\"name\":\"Motivos\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"CondicionDeEntregaId\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.DMS.Events.Hdr.Common.ContratoData\",\"type\":\"record\"}"
}

func (r ContratoData) SchemaName() string {
	return "Andreani.DMS.Events.Hdr.Common.ContratoData"
}

func (_ ContratoData) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ContratoData) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ContratoData) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ContratoData) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ContratoData) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ContratoData) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ContratoData) SetString(v string)   { panic("Unsupported operation") }
func (_ ContratoData) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ContratoData) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.ContratoId}

		return w

	case 1:
		r.Motivos = NewUnionNullArrayString()

		return r.Motivos
	case 2:
		r.CondicionDeEntregaId = NewUnionNullString()

		return r.CondicionDeEntregaId
	}
	panic("Unknown field index")
}

func (r *ContratoData) SetDefault(i int) {
	switch i {
	case 1:
		r.Motivos = nil
		return
	case 2:
		r.CondicionDeEntregaId = nil
		return
	}
	panic("Unknown field index")
}

func (r *ContratoData) NullField(i int) {
	switch i {
	case 1:
		r.Motivos = nil
		return
	case 2:
		r.CondicionDeEntregaId = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ContratoData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ContratoData) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ContratoData) HintSize(int)                     { panic("Unsupported operation") }
func (_ ContratoData) Finalize()                        {}

func (_ ContratoData) AvroCRC64Fingerprint() []byte {
	return []byte(ContratoDataAvroCRC64Fingerprint)
}

func (r ContratoData) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["ContratoId"], err = json.Marshal(r.ContratoId)
	if err != nil {
		return nil, err
	}
	output["Motivos"], err = json.Marshal(r.Motivos)
	if err != nil {
		return nil, err
	}
	output["CondicionDeEntregaId"], err = json.Marshal(r.CondicionDeEntregaId)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ContratoData) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["ContratoId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContratoId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ContratoId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Motivos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Motivos); err != nil {
			return err
		}
	} else {
		r.Motivos = NewUnionNullArrayString()

		r.Motivos = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["CondicionDeEntregaId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CondicionDeEntregaId); err != nil {
			return err
		}
	} else {
		r.CondicionDeEntregaId = NewUnionNullString()

		r.CondicionDeEntregaId = nil
	}
	return nil
}