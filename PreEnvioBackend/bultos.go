// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     GenerarPreenvios.avsc
 */
package PreEnvioBackendEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Bultos struct {
	Kilos *UnionNullFloat `json:"Kilos"`

	ValorDeclaradoConImpuestos *UnionNullFloat `json:"ValorDeclaradoConImpuestos"`

	VolumenCm *UnionNullFloat `json:"VolumenCm"`

	Referencias *UnionNullArrayString `json:"Referencias"`
}

const BultosAvroCRC64Fingerprint = "\xf5\x14\xfe\x0e\xa5)\x88\x1b"

func NewBultos() Bultos {
	r := Bultos{}
	r.Kilos = nil
	r.ValorDeclaradoConImpuestos = nil
	r.VolumenCm = nil
	r.Referencias = nil
	return r
}

func DeserializeBultos(r io.Reader) (Bultos, error) {
	t := NewBultos()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeBultosFromSchema(r io.Reader, schema string) (Bultos, error) {
	t := NewBultos()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeBultos(r Bultos, w io.Writer) error {
	var err error
	err = writeUnionNullFloat(r.Kilos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFloat(r.ValorDeclaradoConImpuestos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFloat(r.VolumenCm, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayString(r.Referencias, w)
	if err != nil {
		return err
	}
	return err
}

func (r Bultos) Serialize(w io.Writer) error {
	return writeBultos(r, w)
}

func (r Bultos) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Kilos\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"ValorDeclaradoConImpuestos\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"VolumenCm\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Referencias\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]}],\"name\":\"Andreani.PreEnvioBackend.Events.Preenvio.Common.Bultos\",\"type\":\"record\"}"
}

func (r Bultos) SchemaName() string {
	return "Andreani.PreEnvioBackend.Events.Preenvio.Common.Bultos"
}

func (_ Bultos) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Bultos) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Bultos) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Bultos) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Bultos) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Bultos) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Bultos) SetString(v string)   { panic("Unsupported operation") }
func (_ Bultos) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Bultos) Get(i int) types.Field {
	switch i {
	case 0:
		r.Kilos = NewUnionNullFloat()

		return r.Kilos
	case 1:
		r.ValorDeclaradoConImpuestos = NewUnionNullFloat()

		return r.ValorDeclaradoConImpuestos
	case 2:
		r.VolumenCm = NewUnionNullFloat()

		return r.VolumenCm
	case 3:
		r.Referencias = NewUnionNullArrayString()

		return r.Referencias
	}
	panic("Unknown field index")
}

func (r *Bultos) SetDefault(i int) {
	switch i {
	case 0:
		r.Kilos = nil
		return
	case 1:
		r.ValorDeclaradoConImpuestos = nil
		return
	case 2:
		r.VolumenCm = nil
		return
	case 3:
		r.Referencias = nil
		return
	}
	panic("Unknown field index")
}

func (r *Bultos) NullField(i int) {
	switch i {
	case 0:
		r.Kilos = nil
		return
	case 1:
		r.ValorDeclaradoConImpuestos = nil
		return
	case 2:
		r.VolumenCm = nil
		return
	case 3:
		r.Referencias = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Bultos) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Bultos) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Bultos) HintSize(int)                     { panic("Unsupported operation") }
func (_ Bultos) Finalize()                        {}

func (_ Bultos) AvroCRC64Fingerprint() []byte {
	return []byte(BultosAvroCRC64Fingerprint)
}

func (r Bultos) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Kilos"], err = json.Marshal(r.Kilos)
	if err != nil {
		return nil, err
	}
	output["ValorDeclaradoConImpuestos"], err = json.Marshal(r.ValorDeclaradoConImpuestos)
	if err != nil {
		return nil, err
	}
	output["VolumenCm"], err = json.Marshal(r.VolumenCm)
	if err != nil {
		return nil, err
	}
	output["Referencias"], err = json.Marshal(r.Referencias)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Bultos) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Kilos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Kilos); err != nil {
			return err
		}
	} else {
		r.Kilos = NewUnionNullFloat()

		r.Kilos = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ValorDeclaradoConImpuestos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ValorDeclaradoConImpuestos); err != nil {
			return err
		}
	} else {
		r.ValorDeclaradoConImpuestos = NewUnionNullFloat()

		r.ValorDeclaradoConImpuestos = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["VolumenCm"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VolumenCm); err != nil {
			return err
		}
	} else {
		r.VolumenCm = NewUnionNullFloat()

		r.VolumenCm = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Referencias"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Referencias); err != nil {
			return err
		}
	} else {
		r.Referencias = NewUnionNullArrayString()

		r.Referencias = nil
	}
	return nil
}