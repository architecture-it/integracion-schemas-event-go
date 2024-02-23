// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CambioDeAlmacenConLpnSolicitadoV2.avsc
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

type CambioAlmacenConLpn struct {
	Propietario *UnionNullString `json:"propietario"`

	Articulo *UnionNullString `json:"articulo"`

	Lote *UnionNullString `json:"lote"`

	LpnId *UnionNullString `json:"lpnId"`

	Cantidad *UnionNullInt `json:"cantidad"`

	AlmacenOrigen *UnionNullString `json:"almacenOrigen"`

	AlmacenDestino *UnionNullString `json:"almacenDestino"`

	OrdenExterna *UnionNullString `json:"ordenExterna"`
}

const CambioAlmacenConLpnAvroCRC64Fingerprint = "\xe2D:P\x90\xfd\x10\xac"

func NewCambioAlmacenConLpn() CambioAlmacenConLpn {
	r := CambioAlmacenConLpn{}
	r.Propietario = nil
	r.Articulo = nil
	r.Lote = nil
	r.LpnId = nil
	r.Cantidad = nil
	r.AlmacenOrigen = nil
	r.AlmacenDestino = nil
	r.OrdenExterna = nil
	return r
}

func DeserializeCambioAlmacenConLpn(r io.Reader) (CambioAlmacenConLpn, error) {
	t := NewCambioAlmacenConLpn()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCambioAlmacenConLpnFromSchema(r io.Reader, schema string) (CambioAlmacenConLpn, error) {
	t := NewCambioAlmacenConLpn()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCambioAlmacenConLpn(r CambioAlmacenConLpn, w io.Writer) error {
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
	err = writeUnionNullString(r.LpnId, w)
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
	err = writeUnionNullString(r.OrdenExterna, w)
	if err != nil {
		return err
	}
	return err
}

func (r CambioAlmacenConLpn) Serialize(w io.Writer) error {
	return writeCambioAlmacenConLpn(r, w)
}

func (r CambioAlmacenConLpn) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"propietario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"articulo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"lote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"lpnId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cantidad\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"almacenOrigen\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"almacenDestino\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ordenExterna\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.Wapv2.Events.Record.CambioAlmacenConLpn\",\"type\":\"record\"}"
}

func (r CambioAlmacenConLpn) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.CambioAlmacenConLpn"
}

func (_ CambioAlmacenConLpn) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CambioAlmacenConLpn) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CambioAlmacenConLpn) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CambioAlmacenConLpn) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CambioAlmacenConLpn) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CambioAlmacenConLpn) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CambioAlmacenConLpn) SetString(v string)   { panic("Unsupported operation") }
func (_ CambioAlmacenConLpn) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CambioAlmacenConLpn) Get(i int) types.Field {
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
		r.LpnId = NewUnionNullString()

		return r.LpnId
	case 4:
		r.Cantidad = NewUnionNullInt()

		return r.Cantidad
	case 5:
		r.AlmacenOrigen = NewUnionNullString()

		return r.AlmacenOrigen
	case 6:
		r.AlmacenDestino = NewUnionNullString()

		return r.AlmacenDestino
	case 7:
		r.OrdenExterna = NewUnionNullString()

		return r.OrdenExterna
	}
	panic("Unknown field index")
}

func (r *CambioAlmacenConLpn) SetDefault(i int) {
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
		r.LpnId = nil
		return
	case 4:
		r.Cantidad = nil
		return
	case 5:
		r.AlmacenOrigen = nil
		return
	case 6:
		r.AlmacenDestino = nil
		return
	case 7:
		r.OrdenExterna = nil
		return
	}
	panic("Unknown field index")
}

func (r *CambioAlmacenConLpn) NullField(i int) {
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
		r.LpnId = nil
		return
	case 4:
		r.Cantidad = nil
		return
	case 5:
		r.AlmacenOrigen = nil
		return
	case 6:
		r.AlmacenDestino = nil
		return
	case 7:
		r.OrdenExterna = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ CambioAlmacenConLpn) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CambioAlmacenConLpn) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CambioAlmacenConLpn) HintSize(int)                     { panic("Unsupported operation") }
func (_ CambioAlmacenConLpn) Finalize()                        {}

func (_ CambioAlmacenConLpn) AvroCRC64Fingerprint() []byte {
	return []byte(CambioAlmacenConLpnAvroCRC64Fingerprint)
}

func (r CambioAlmacenConLpn) MarshalJSON() ([]byte, error) {
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
	output["lpnId"], err = json.Marshal(r.LpnId)
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
	output["ordenExterna"], err = json.Marshal(r.OrdenExterna)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CambioAlmacenConLpn) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["lpnId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LpnId); err != nil {
			return err
		}
	} else {
		r.LpnId = NewUnionNullString()

		r.LpnId = nil
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
	val = func() json.RawMessage {
		if v, ok := fields["ordenExterna"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenExterna); err != nil {
			return err
		}
	} else {
		r.OrdenExterna = NewUnionNullString()

		r.OrdenExterna = nil
	}
	return nil
}
