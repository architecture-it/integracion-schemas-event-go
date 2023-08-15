// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CambioDeAlmacenSinLpnSolicitadoV2.avsc
 */
package Wapv2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type CambioAlmacenSinLpn struct {
	Propietario *UnionNullString `json:"propietario"`

	Articulo *UnionNullString `json:"articulo"`

	Lote *UnionNullString `json:"lote"`

	Cantidad *UnionNullInt `json:"cantidad"`

	AlmacenOrigen *UnionNullString `json:"almacenOrigen"`

	AlmacenDestino *UnionNullString `json:"almacenDestino"`
}

const CambioAlmacenSinLpnAvroCRC64Fingerprint = "!\\L\f\xc4P\xd9F"

func NewCambioAlmacenSinLpn() CambioAlmacenSinLpn {
	r := CambioAlmacenSinLpn{}
	r.Propietario = nil
	r.Articulo = nil
	r.Lote = nil
	r.Cantidad = nil
	r.AlmacenOrigen = nil
	r.AlmacenDestino = nil
	return r
}

func DeserializeCambioAlmacenSinLpn(r io.Reader) (CambioAlmacenSinLpn, error) {
	t := NewCambioAlmacenSinLpn()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCambioAlmacenSinLpnFromSchema(r io.Reader, schema string) (CambioAlmacenSinLpn, error) {
	t := NewCambioAlmacenSinLpn()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCambioAlmacenSinLpn(r CambioAlmacenSinLpn, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Articulo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Lote, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.Cantidad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.AlmacenOrigen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.AlmacenDestino, w)
	if err != nil {
		return err
	}
	return err
}

func (r CambioAlmacenSinLpn) Serialize(w io.Writer) error {
	return writeCambioAlmacenSinLpn(r, w)
}

func (r CambioAlmacenSinLpn) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"propietario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"articulo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"lote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cantidad\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"almacenOrigen\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"almacenDestino\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.Wapv2.Events.Record.CambioAlmacenSinLpn\",\"type\":\"record\"}"
}

func (r CambioAlmacenSinLpn) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.CambioAlmacenSinLpn"
}

func (_ CambioAlmacenSinLpn) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CambioAlmacenSinLpn) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CambioAlmacenSinLpn) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CambioAlmacenSinLpn) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CambioAlmacenSinLpn) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CambioAlmacenSinLpn) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CambioAlmacenSinLpn) SetString(v string)   { panic("Unsupported operation") }
func (_ CambioAlmacenSinLpn) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CambioAlmacenSinLpn) Get(i int) types.Field {
	switch i {
	case 0:
		r.Propietario = NewUnionNullString()

		return r.Propietario
	case 1:
		r.Articulo = NewUnionNullString()

		return r.Articulo
	case 2:
		r.Lote = NewUnionNullString()

		return r.Lote
	case 3:
		r.Cantidad = NewUnionNullInt()

		return r.Cantidad
	case 4:
		r.AlmacenOrigen = NewUnionNullString()

		return r.AlmacenOrigen
	case 5:
		r.AlmacenDestino = NewUnionNullString()

		return r.AlmacenDestino
	}
	panic("Unknown field index")
}

func (r *CambioAlmacenSinLpn) SetDefault(i int) {
	switch i {
	case 0:
		r.Propietario = nil
		return
	case 1:
		r.Articulo = nil
		return
	case 2:
		r.Lote = nil
		return
	case 3:
		r.Cantidad = nil
		return
	case 4:
		r.AlmacenOrigen = nil
		return
	case 5:
		r.AlmacenDestino = nil
		return
	}
	panic("Unknown field index")
}

func (r *CambioAlmacenSinLpn) NullField(i int) {
	switch i {
	case 0:
		r.Propietario = nil
		return
	case 1:
		r.Articulo = nil
		return
	case 2:
		r.Lote = nil
		return
	case 3:
		r.Cantidad = nil
		return
	case 4:
		r.AlmacenOrigen = nil
		return
	case 5:
		r.AlmacenDestino = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ CambioAlmacenSinLpn) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CambioAlmacenSinLpn) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CambioAlmacenSinLpn) HintSize(int)                     { panic("Unsupported operation") }
func (_ CambioAlmacenSinLpn) Finalize()                        {}

func (_ CambioAlmacenSinLpn) AvroCRC64Fingerprint() []byte {
	return []byte(CambioAlmacenSinLpnAvroCRC64Fingerprint)
}

func (r CambioAlmacenSinLpn) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["articulo"], err = json.Marshal(r.Articulo)
	if err != nil {
		return nil, err
	}
	output["lote"], err = json.Marshal(r.Lote)
	if err != nil {
		return nil, err
	}
	output["cantidad"], err = json.Marshal(r.Cantidad)
	if err != nil {
		return nil, err
	}
	output["almacenOrigen"], err = json.Marshal(r.AlmacenOrigen)
	if err != nil {
		return nil, err
	}
	output["almacenDestino"], err = json.Marshal(r.AlmacenDestino)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CambioAlmacenSinLpn) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["propietario"]; ok {
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
		if v, ok := fields["articulo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Articulo); err != nil {
			return err
		}
	} else {
		r.Articulo = NewUnionNullString()

		r.Articulo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["lote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Lote); err != nil {
			return err
		}
	} else {
		r.Lote = NewUnionNullString()

		r.Lote = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["cantidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cantidad); err != nil {
			return err
		}
	} else {
		r.Cantidad = NewUnionNullInt()

		r.Cantidad = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["almacenOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AlmacenOrigen); err != nil {
			return err
		}
	} else {
		r.AlmacenOrigen = NewUnionNullString()

		r.AlmacenOrigen = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["almacenDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AlmacenDestino); err != nil {
			return err
		}
	} else {
		r.AlmacenDestino = NewUnionNullString()

		r.AlmacenDestino = nil
	}
	return nil
}