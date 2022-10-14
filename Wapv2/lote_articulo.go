// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoDeAlmacenSCE.avsc
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

type LoteArticulo struct {
	LoteDeFabricante *UnionNullString `json:"loteDeFabricante"`

	LoteSecundario *UnionNullString `json:"loteSecundario"`

	FechaDeVencimiento *UnionNullLong `json:"fechaDeVencimiento"`

	OtrosDatos *UnionNullString `json:"otrosDatos"`
}

const LoteArticuloAvroCRC64Fingerprint = "\x84\x03z'\x06\x00\x97\xec"

func NewLoteArticulo() LoteArticulo {
	r := LoteArticulo{}
	r.LoteDeFabricante = nil
	r.LoteSecundario = nil
	r.OtrosDatos = nil
	return r
}

func DeserializeLoteArticulo(r io.Reader) (LoteArticulo, error) {
	t := NewLoteArticulo()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeLoteArticuloFromSchema(r io.Reader, schema string) (LoteArticulo, error) {
	t := NewLoteArticulo()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeLoteArticulo(r LoteArticulo, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.LoteDeFabricante, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LoteSecundario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaDeVencimiento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.OtrosDatos, w)
	if err != nil {
		return err
	}
	return err
}

func (r LoteArticulo) Serialize(w io.Writer) error {
	return writeLoteArticulo(r, w)
}

func (r LoteArticulo) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaDeVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.Wapv2.Events.Record.LoteArticulo\",\"type\":\"record\"}"
}

func (r LoteArticulo) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.LoteArticulo"
}

func (_ LoteArticulo) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ LoteArticulo) SetInt(v int32)       { panic("Unsupported operation") }
func (_ LoteArticulo) SetLong(v int64)      { panic("Unsupported operation") }
func (_ LoteArticulo) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ LoteArticulo) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ LoteArticulo) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ LoteArticulo) SetString(v string)   { panic("Unsupported operation") }
func (_ LoteArticulo) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *LoteArticulo) Get(i int) types.Field {
	switch i {
	case 0:
		r.LoteDeFabricante = NewUnionNullString()

		return r.LoteDeFabricante
	case 1:
		r.LoteSecundario = NewUnionNullString()

		return r.LoteSecundario
	case 2:
		r.FechaDeVencimiento = NewUnionNullLong()

		return r.FechaDeVencimiento
	case 3:
		r.OtrosDatos = NewUnionNullString()

		return r.OtrosDatos
	}
	panic("Unknown field index")
}

func (r *LoteArticulo) SetDefault(i int) {
	switch i {
	case 0:
		r.LoteDeFabricante = nil
		return
	case 1:
		r.LoteSecundario = nil
		return
	case 3:
		r.OtrosDatos = nil
		return
	}
	panic("Unknown field index")
}

func (r *LoteArticulo) NullField(i int) {
	switch i {
	case 0:
		r.LoteDeFabricante = nil
		return
	case 1:
		r.LoteSecundario = nil
		return
	case 2:
		r.FechaDeVencimiento = nil
		return
	case 3:
		r.OtrosDatos = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ LoteArticulo) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ LoteArticulo) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ LoteArticulo) HintSize(int)                     { panic("Unsupported operation") }
func (_ LoteArticulo) Finalize()                        {}

func (_ LoteArticulo) AvroCRC64Fingerprint() []byte {
	return []byte(LoteArticuloAvroCRC64Fingerprint)
}

func (r LoteArticulo) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["loteDeFabricante"], err = json.Marshal(r.LoteDeFabricante)
	if err != nil {
		return nil, err
	}
	output["loteSecundario"], err = json.Marshal(r.LoteSecundario)
	if err != nil {
		return nil, err
	}
	output["fechaDeVencimiento"], err = json.Marshal(r.FechaDeVencimiento)
	if err != nil {
		return nil, err
	}
	output["otrosDatos"], err = json.Marshal(r.OtrosDatos)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *LoteArticulo) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["loteDeFabricante"]; ok {
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
		if v, ok := fields["loteSecundario"]; ok {
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
		if v, ok := fields["fechaDeVencimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaDeVencimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fechaDeVencimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["otrosDatos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OtrosDatos); err != nil {
			return err
		}
	} else {
		r.OtrosDatos = NewUnionNullString()

		r.OtrosDatos = nil
	}
	return nil
}
