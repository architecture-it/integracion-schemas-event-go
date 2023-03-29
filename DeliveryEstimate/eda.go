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

type EDA struct {
	EdaDias int32 `json:"edaDias"`

	EdaFecha int64 `json:"edaFecha"`

	NumeroEnvio string `json:"numeroEnvio"`

	CodigoDeContratoInterno string `json:"codigoDeContratoInterno"`

	CalculoEda CalculoEda `json:"CalculoEda"`

	Timestamp string `json:"Timestamp"`
}

const EDAAvroCRC64Fingerprint = "<\x96\xdff\x87R\x9cw"

func NewEDA() EDA {
	r := EDA{}
	r.CalculoEda = NewCalculoEda()

	return r
}

func DeserializeEDA(r io.Reader) (EDA, error) {
	t := NewEDA()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEDAFromSchema(r io.Reader, schema string) (EDA, error) {
	t := NewEDA()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEDA(r EDA, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.EdaDias, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.EdaFecha, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroEnvio, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoDeContratoInterno, w)
	if err != nil {
		return err
	}
	err = writeCalculoEda(r.CalculoEda, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Timestamp, w)
	if err != nil {
		return err
	}
	return err
}

func (r EDA) Serialize(w io.Writer) error {
	return writeEDA(r, w)
}

func (r EDA) Schema() string {
	return "{\"fields\":[{\"name\":\"edaDias\",\"type\":\"int\"},{\"name\":\"edaFecha\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"numeroEnvio\",\"type\":\"string\"},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"name\":\"CalculoEda\",\"type\":{\"fields\":[{\"name\":\"cpSucursalOrigen\",\"type\":\"string\"},{\"name\":\"cpSucursalDistribucion\",\"type\":\"string\"},{\"name\":\"cpDestino\",\"type\":\"string\"}],\"name\":\"CalculoEda\",\"type\":\"record\"}},{\"name\":\"Timestamp\",\"type\":\"string\"}],\"name\":\"Andreani.DeliveryEstimate.Events.Records.EDA\",\"type\":\"record\"}"
}

func (r EDA) SchemaName() string {
	return "Andreani.DeliveryEstimate.Events.Records.EDA"
}

func (_ EDA) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EDA) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EDA) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EDA) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EDA) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EDA) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EDA) SetString(v string)   { panic("Unsupported operation") }
func (_ EDA) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EDA) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.EdaDias}

		return w

	case 1:
		w := types.Long{Target: &r.EdaFecha}

		return w

	case 2:
		w := types.String{Target: &r.NumeroEnvio}

		return w

	case 3:
		w := types.String{Target: &r.CodigoDeContratoInterno}

		return w

	case 4:
		r.CalculoEda = NewCalculoEda()

		w := types.Record{Target: &r.CalculoEda}

		return w

	case 5:
		w := types.String{Target: &r.Timestamp}

		return w

	}
	panic("Unknown field index")
}

func (r *EDA) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EDA) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EDA) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EDA) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EDA) HintSize(int)                     { panic("Unsupported operation") }
func (_ EDA) Finalize()                        {}

func (_ EDA) AvroCRC64Fingerprint() []byte {
	return []byte(EDAAvroCRC64Fingerprint)
}

func (r EDA) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["edaDias"], err = json.Marshal(r.EdaDias)
	if err != nil {
		return nil, err
	}
	output["edaFecha"], err = json.Marshal(r.EdaFecha)
	if err != nil {
		return nil, err
	}
	output["numeroEnvio"], err = json.Marshal(r.NumeroEnvio)
	if err != nil {
		return nil, err
	}
	output["codigoDeContratoInterno"], err = json.Marshal(r.CodigoDeContratoInterno)
	if err != nil {
		return nil, err
	}
	output["CalculoEda"], err = json.Marshal(r.CalculoEda)
	if err != nil {
		return nil, err
	}
	output["Timestamp"], err = json.Marshal(r.Timestamp)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EDA) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["edaDias"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EdaDias); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for edaDias")
	}
	val = func() json.RawMessage {
		if v, ok := fields["edaFecha"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EdaFecha); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for edaFecha")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroEnvio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroEnvio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoDeContratoInterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeContratoInterno); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigoDeContratoInterno")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CalculoEda"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CalculoEda); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CalculoEda")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Timestamp"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Timestamp); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Timestamp")
	}
	return nil
}
