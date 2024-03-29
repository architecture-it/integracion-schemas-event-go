// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Tarifa.avsc
 */
package CotizarBackendEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Tarifa struct {
	SeguroDistribucion string `json:"SeguroDistribucion"`

	Distribucion string `json:"Distribucion"`

	Total string `json:"Total"`
}

const TarifaAvroCRC64Fingerprint = "(\xe9<:\x88\x06\x16P"

func NewTarifa() Tarifa {
	r := Tarifa{}
	return r
}

func DeserializeTarifa(r io.Reader) (Tarifa, error) {
	t := NewTarifa()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTarifaFromSchema(r io.Reader, schema string) (Tarifa, error) {
	t := NewTarifa()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTarifa(r Tarifa, w io.Writer) error {
	var err error
	err = vm.WriteString(r.SeguroDistribucion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Distribucion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Total, w)
	if err != nil {
		return err
	}
	return err
}

func (r Tarifa) Serialize(w io.Writer) error {
	return writeTarifa(r, w)
}

func (r Tarifa) Schema() string {
	return "{\"fields\":[{\"name\":\"SeguroDistribucion\",\"type\":\"string\"},{\"name\":\"Distribucion\",\"type\":\"string\"},{\"name\":\"Total\",\"type\":\"string\"}],\"name\":\"Andreani.CotizarBackend.Events.Cotizar.Common.Tarifa\",\"type\":\"record\"}"
}

func (r Tarifa) SchemaName() string {
	return "Andreani.CotizarBackend.Events.Cotizar.Common.Tarifa"
}

func (_ Tarifa) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Tarifa) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Tarifa) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Tarifa) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Tarifa) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Tarifa) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Tarifa) SetString(v string)   { panic("Unsupported operation") }
func (_ Tarifa) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Tarifa) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.SeguroDistribucion}

		return w

	case 1:
		w := types.String{Target: &r.Distribucion}

		return w

	case 2:
		w := types.String{Target: &r.Total}

		return w

	}
	panic("Unknown field index")
}

func (r *Tarifa) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Tarifa) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Tarifa) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Tarifa) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Tarifa) HintSize(int)                     { panic("Unsupported operation") }
func (_ Tarifa) Finalize()                        {}

func (_ Tarifa) AvroCRC64Fingerprint() []byte {
	return []byte(TarifaAvroCRC64Fingerprint)
}

func (r Tarifa) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["SeguroDistribucion"], err = json.Marshal(r.SeguroDistribucion)
	if err != nil {
		return nil, err
	}
	output["Distribucion"], err = json.Marshal(r.Distribucion)
	if err != nil {
		return nil, err
	}
	output["Total"], err = json.Marshal(r.Total)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Tarifa) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["SeguroDistribucion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SeguroDistribucion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SeguroDistribucion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Distribucion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Distribucion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Distribucion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Total"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Total); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Total")
	}
	return nil
}
