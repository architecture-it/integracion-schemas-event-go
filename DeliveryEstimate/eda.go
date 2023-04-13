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
	DiasHabiles int32 `json:"diasHabiles"`

	Eda int64 `json:"eda"`

	NumeroEnvio string `json:"numeroEnvio"`

	Contrato string `json:"contrato"`

	CalculoEda CalculoEda `json:"CalculoEda"`

	Timestamp string `json:"Timestamp"`

	FechaAlta int64 `json:"FechaAlta"`
}

const EDAAvroCRC64Fingerprint = "\xcd\xf9 \x05)\x9d\xde\xcd"

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
	err = vm.WriteInt(r.DiasHabiles, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Eda, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroEnvio, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Contrato, w)
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
	err = vm.WriteLong(r.FechaAlta, w)
	if err != nil {
		return err
	}
	return err
}

func (r EDA) Serialize(w io.Writer) error {
	return writeEDA(r, w)
}

func (r EDA) Schema() string {
	return "{\"fields\":[{\"name\":\"diasHabiles\",\"type\":\"int\"},{\"name\":\"eda\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"numeroEnvio\",\"type\":\"string\"},{\"name\":\"contrato\",\"type\":\"string\"},{\"name\":\"CalculoEda\",\"type\":{\"fields\":[{\"name\":\"cpSucursalOrigen\",\"type\":\"string\"},{\"name\":\"cpSucursalDistribucion\",\"type\":\"string\"},{\"name\":\"cpDestino\",\"type\":\"string\"}],\"name\":\"CalculoEda\",\"type\":\"record\"}},{\"name\":\"Timestamp\",\"type\":\"string\"},{\"name\":\"FechaAlta\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}}],\"name\":\"Andreani.DeliveryEstimate.Events.Records.EDA\",\"type\":\"record\"}"
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
		w := types.Int{Target: &r.DiasHabiles}

		return w

	case 1:
		w := types.Long{Target: &r.Eda}

		return w

	case 2:
		w := types.String{Target: &r.NumeroEnvio}

		return w

	case 3:
		w := types.String{Target: &r.Contrato}

		return w

	case 4:
		r.CalculoEda = NewCalculoEda()

		w := types.Record{Target: &r.CalculoEda}

		return w

	case 5:
		w := types.String{Target: &r.Timestamp}

		return w

	case 6:
		w := types.Long{Target: &r.FechaAlta}

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
	output["diasHabiles"], err = json.Marshal(r.DiasHabiles)
	if err != nil {
		return nil, err
	}
	output["eda"], err = json.Marshal(r.Eda)
	if err != nil {
		return nil, err
	}
	output["numeroEnvio"], err = json.Marshal(r.NumeroEnvio)
	if err != nil {
		return nil, err
	}
	output["contrato"], err = json.Marshal(r.Contrato)
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
	output["FechaAlta"], err = json.Marshal(r.FechaAlta)
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
		if v, ok := fields["diasHabiles"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DiasHabiles); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for diasHabiles")
	}
	val = func() json.RawMessage {
		if v, ok := fields["eda"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Eda); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for eda")
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
		if v, ok := fields["contrato"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contrato); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for contrato")
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
	val = func() json.RawMessage {
		if v, ok := fields["FechaAlta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaAlta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaAlta")
	}
	return nil
}
