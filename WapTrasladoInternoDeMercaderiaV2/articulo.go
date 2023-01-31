// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     WapTrasladoInternoDeMercaderiaSolicitadaV2.avsc
 */
package WapTrasladoInternoDeMercaderiaV2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Articulo struct {
	Cantidad float32 `json:"Cantidad"`

	CodigoDestino string `json:"CodigoDestino"`

	CodigoOrigen string `json:"CodigoOrigen"`

	LoteDeFabricante *UnionNullString `json:"LoteDeFabricante"`

	LoteSecundario *UnionNullString `json:"LoteSecundario"`

	LPN *UnionNullString `json:"LPN"`
}

const ArticuloAvroCRC64Fingerprint = "\ue2a2_<\x1bj\xfb"

func NewArticulo() Articulo {
	r := Articulo{}
	r.LoteDeFabricante = nil
	r.LoteSecundario = nil
	r.LPN = nil
	return r
}

func DeserializeArticulo(r io.Reader) (Articulo, error) {
	t := NewArticulo()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeArticuloFromSchema(r io.Reader, schema string) (Articulo, error) {
	t := NewArticulo()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeArticulo(r Articulo, w io.Writer) error {
	var err error
	err = vm.WriteFloat(r.Cantidad, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoDestino, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoOrigen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LoteDeFabricante, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LoteSecundario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LPN, w)
	if err != nil {
		return err
	}
	return err
}

func (r Articulo) Serialize(w io.Writer) error {
	return writeArticulo(r, w)
}

func (r Articulo) Schema() string {
	return "{\"fields\":[{\"name\":\"Cantidad\",\"type\":\"float\"},{\"name\":\"CodigoDestino\",\"type\":\"string\"},{\"name\":\"CodigoOrigen\",\"type\":\"string\"},{\"default\":null,\"name\":\"LoteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LoteSecundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LPN\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.WapTrasladoInternoDeMercaderiaV2.Events.Record.Articulo\",\"type\":\"record\"}"
}

func (r Articulo) SchemaName() string {
	return "Andreani.WapTrasladoInternoDeMercaderiaV2.Events.Record.Articulo"
}

func (_ Articulo) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Articulo) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Articulo) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Articulo) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Articulo) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Articulo) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Articulo) SetString(v string)   { panic("Unsupported operation") }
func (_ Articulo) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Articulo) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Float{Target: &r.Cantidad}

		return w

	case 1:
		w := types.String{Target: &r.CodigoDestino}

		return w

	case 2:
		w := types.String{Target: &r.CodigoOrigen}

		return w

	case 3:
		r.LoteDeFabricante = NewUnionNullString()

		return r.LoteDeFabricante
	case 4:
		r.LoteSecundario = NewUnionNullString()

		return r.LoteSecundario
	case 5:
		r.LPN = NewUnionNullString()

		return r.LPN
	}
	panic("Unknown field index")
}

func (r *Articulo) SetDefault(i int) {
	switch i {
	case 3:
		r.LoteDeFabricante = nil
		return
	case 4:
		r.LoteSecundario = nil
		return
	case 5:
		r.LPN = nil
		return
	}
	panic("Unknown field index")
}

func (r *Articulo) NullField(i int) {
	switch i {
	case 3:
		r.LoteDeFabricante = nil
		return
	case 4:
		r.LoteSecundario = nil
		return
	case 5:
		r.LPN = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Articulo) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Articulo) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Articulo) HintSize(int)                     { panic("Unsupported operation") }
func (_ Articulo) Finalize()                        {}

func (_ Articulo) AvroCRC64Fingerprint() []byte {
	return []byte(ArticuloAvroCRC64Fingerprint)
}

func (r Articulo) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Cantidad"], err = json.Marshal(r.Cantidad)
	if err != nil {
		return nil, err
	}
	output["CodigoDestino"], err = json.Marshal(r.CodigoDestino)
	if err != nil {
		return nil, err
	}
	output["CodigoOrigen"], err = json.Marshal(r.CodigoOrigen)
	if err != nil {
		return nil, err
	}
	output["LoteDeFabricante"], err = json.Marshal(r.LoteDeFabricante)
	if err != nil {
		return nil, err
	}
	output["LoteSecundario"], err = json.Marshal(r.LoteSecundario)
	if err != nil {
		return nil, err
	}
	output["LPN"], err = json.Marshal(r.LPN)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Articulo) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
		if v, ok := fields["CodigoDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDestino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoDestino")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoOrigen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoOrigen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LoteDeFabricante"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteDeFabricante); err != nil {
			return err
		}
	} else {
		r.LoteDeFabricante = NewUnionNullString()

		r.LoteDeFabricante = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["LoteSecundario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteSecundario); err != nil {
			return err
		}
	} else {
		r.LoteSecundario = NewUnionNullString()

		r.LoteSecundario = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["LPN"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LPN); err != nil {
			return err
		}
	} else {
		r.LPN = NewUnionNullString()

		r.LPN = nil
	}
	return nil
}
