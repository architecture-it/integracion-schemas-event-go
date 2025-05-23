// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CambioDeStock.avsc
 */
package WarehouseStockEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Ajuste struct {
	StockTotal *UnionNullFloat `json:"StockTotal"`

	StockDisponible *UnionNullFloat `json:"StockDisponible"`

	StockEnTransito *UnionNullFloat `json:"StockEnTransito"`

	StockAnteriorAjuste *UnionNullFloat `json:"StockAnteriorAjuste"`
}

const AjusteAvroCRC64Fingerprint = "@\xb9i#H\x95\x8f["

func NewAjuste() Ajuste {
	r := Ajuste{}
	r.StockTotal = nil
	r.StockDisponible = nil
	r.StockEnTransito = nil
	r.StockAnteriorAjuste = nil
	return r
}

func DeserializeAjuste(r io.Reader) (Ajuste, error) {
	t := NewAjuste()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAjusteFromSchema(r io.Reader, schema string) (Ajuste, error) {
	t := NewAjuste()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAjuste(r Ajuste, w io.Writer) error {
	var err error
	err = writeUnionNullFloat(r.StockTotal, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFloat(r.StockDisponible, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFloat(r.StockEnTransito, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFloat(r.StockAnteriorAjuste, w)
	if err != nil {
		return err
	}
	return err
}

func (r Ajuste) Serialize(w io.Writer) error {
	return writeAjuste(r, w)
}

func (r Ajuste) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"StockTotal\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"StockDisponible\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"StockEnTransito\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"StockAnteriorAjuste\",\"type\":[\"null\",\"float\"]}],\"name\":\"Andreani.WarehouseStock.Events.StockCommon.Ajuste\",\"type\":\"record\"}"
}

func (r Ajuste) SchemaName() string {
	return "Andreani.WarehouseStock.Events.StockCommon.Ajuste"
}

func (_ Ajuste) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Ajuste) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Ajuste) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Ajuste) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Ajuste) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Ajuste) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Ajuste) SetString(v string)   { panic("Unsupported operation") }
func (_ Ajuste) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Ajuste) Get(i int) types.Field {
	switch i {
	case 0:
		r.StockTotal = NewUnionNullFloat()

		return r.StockTotal
	case 1:
		r.StockDisponible = NewUnionNullFloat()

		return r.StockDisponible
	case 2:
		r.StockEnTransito = NewUnionNullFloat()

		return r.StockEnTransito
	case 3:
		r.StockAnteriorAjuste = NewUnionNullFloat()

		return r.StockAnteriorAjuste
	}
	panic("Unknown field index")
}

func (r *Ajuste) SetDefault(i int) {
	switch i {
	case 0:
		r.StockTotal = nil
		return
	case 1:
		r.StockDisponible = nil
		return
	case 2:
		r.StockEnTransito = nil
		return
	case 3:
		r.StockAnteriorAjuste = nil
		return
	}
	panic("Unknown field index")
}

func (r *Ajuste) NullField(i int) {
	switch i {
	case 0:
		r.StockTotal = nil
		return
	case 1:
		r.StockDisponible = nil
		return
	case 2:
		r.StockEnTransito = nil
		return
	case 3:
		r.StockAnteriorAjuste = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Ajuste) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Ajuste) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Ajuste) HintSize(int)                     { panic("Unsupported operation") }
func (_ Ajuste) Finalize()                        {}

func (_ Ajuste) AvroCRC64Fingerprint() []byte {
	return []byte(AjusteAvroCRC64Fingerprint)
}

func (r Ajuste) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["StockTotal"], err = json.Marshal(r.StockTotal)
	if err != nil {
		return nil, err
	}
	output["StockDisponible"], err = json.Marshal(r.StockDisponible)
	if err != nil {
		return nil, err
	}
	output["StockEnTransito"], err = json.Marshal(r.StockEnTransito)
	if err != nil {
		return nil, err
	}
	output["StockAnteriorAjuste"], err = json.Marshal(r.StockAnteriorAjuste)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Ajuste) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["StockTotal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StockTotal); err != nil {
			return err
		}
	} else {
		r.StockTotal = NewUnionNullFloat()

		r.StockTotal = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["StockDisponible"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StockDisponible); err != nil {
			return err
		}
	} else {
		r.StockDisponible = NewUnionNullFloat()

		r.StockDisponible = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["StockEnTransito"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StockEnTransito); err != nil {
			return err
		}
	} else {
		r.StockEnTransito = NewUnionNullFloat()

		r.StockEnTransito = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["StockAnteriorAjuste"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StockAnteriorAjuste); err != nil {
			return err
		}
	} else {
		r.StockAnteriorAjuste = NewUnionNullFloat()

		r.StockAnteriorAjuste = nil
	}
	return nil
}
