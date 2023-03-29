// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EDA.avsc
 */
package DeliveryEstimateEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type CalculoEda struct {
	CpSucursalOrigen string `json:"cpSucursalOrigen"`

	CpSucursalDistribucion string `json:"cpSucursalDistribucion"`

	CpDestino string `json:"cpDestino"`
}

const CalculoEdaAvroCRC64Fingerprint = "U\xd270\xbb\x1e|y"

func NewCalculoEda() CalculoEda {
	r := CalculoEda{}
	return r
}

func DeserializeCalculoEda(r io.Reader) (CalculoEda, error) {
	t := NewCalculoEda()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCalculoEdaFromSchema(r io.Reader, schema string) (CalculoEda, error) {
	t := NewCalculoEda()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCalculoEda(r CalculoEda, w io.Writer) error {
	var err error
	err = vm.WriteString(r.CpSucursalOrigen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CpSucursalDistribucion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CpDestino, w)
	if err != nil {
		return err
	}
	return err
}

func (r CalculoEda) Serialize(w io.Writer) error {
	return writeCalculoEda(r, w)
}

func (r CalculoEda) Schema() string {
	return "{\"fields\":[{\"name\":\"cpSucursalOrigen\",\"type\":\"string\"},{\"name\":\"cpSucursalDistribucion\",\"type\":\"string\"},{\"name\":\"cpDestino\",\"type\":\"string\"}],\"name\":\"Andreani.DeliveryEstimate.Events.Records.CalculoEda\",\"type\":\"record\"}"
}

func (r CalculoEda) SchemaName() string {
	return "Andreani.DeliveryEstimate.Events.Records.CalculoEda"
}

func (_ CalculoEda) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CalculoEda) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CalculoEda) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CalculoEda) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CalculoEda) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CalculoEda) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CalculoEda) SetString(v string)   { panic("Unsupported operation") }
func (_ CalculoEda) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CalculoEda) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.CpSucursalOrigen}

		return w

	case 1:
		w := types.String{Target: &r.CpSucursalDistribucion}

		return w

	case 2:
		w := types.String{Target: &r.CpDestino}

		return w

	}
	panic("Unknown field index")
}

func (r *CalculoEda) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *CalculoEda) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ CalculoEda) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CalculoEda) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CalculoEda) HintSize(int)                     { panic("Unsupported operation") }
func (_ CalculoEda) Finalize()                        {}

func (_ CalculoEda) AvroCRC64Fingerprint() []byte {
	return []byte(CalculoEdaAvroCRC64Fingerprint)
}

func (r CalculoEda) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["cpSucursalOrigen"], err = json.Marshal(r.CpSucursalOrigen)
	if err != nil {
		return nil, err
	}
	output["cpSucursalDistribucion"], err = json.Marshal(r.CpSucursalDistribucion)
	if err != nil {
		return nil, err
	}
	output["cpDestino"], err = json.Marshal(r.CpDestino)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CalculoEda) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["cpSucursalOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CpSucursalOrigen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cpSucursalOrigen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["cpSucursalDistribucion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CpSucursalDistribucion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cpSucursalDistribucion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["cpDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CpDestino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cpDestino")
	}
	return nil
}
