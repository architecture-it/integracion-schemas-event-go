// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Contrato.avsc
 */
package ConstanciasElectronicasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Contrato struct {
	CodigoDeContratoInterno string `json:"codigoDeContratoInterno"`

	CodigoDeClienteInterno string `json:"codigoDeClienteInterno"`

	Status EntityStatus `json:"status"`
}

const ContratoAvroCRC64Fingerprint = "\x9e\xc3}\r\xd8t\xd3\xd9"

func NewContrato() Contrato {
	r := Contrato{}
	return r
}

func DeserializeContrato(r io.Reader) (Contrato, error) {
	t := NewContrato()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeContratoFromSchema(r io.Reader, schema string) (Contrato, error) {
	t := NewContrato()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeContrato(r Contrato, w io.Writer) error {
	var err error
	err = vm.WriteString(r.CodigoDeContratoInterno, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoDeClienteInterno, w)
	if err != nil {
		return err
	}
	err = writeEntityStatus(r.Status, w)
	if err != nil {
		return err
	}
	return err
}

func (r Contrato) Serialize(w io.Writer) error {
	return writeContrato(r, w)
}

func (r Contrato) Schema() string {
	return "{\"fields\":[{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"name\":\"codigoDeClienteInterno\",\"type\":\"string\"},{\"name\":\"status\",\"type\":{\"name\":\"EntityStatus\",\"symbols\":[\"Inactive\",\"Active\",\"Expired\",\"Deleted\"],\"type\":\"enum\"}}],\"name\":\"Andreani.ConstanciasElectronicas.Common.Contrato\",\"type\":\"record\"}"
}

func (r Contrato) SchemaName() string {
	return "Andreani.ConstanciasElectronicas.Common.Contrato"
}

func (_ Contrato) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Contrato) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Contrato) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Contrato) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Contrato) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Contrato) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Contrato) SetString(v string)   { panic("Unsupported operation") }
func (_ Contrato) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Contrato) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.CodigoDeContratoInterno}

		return w

	case 1:
		w := types.String{Target: &r.CodigoDeClienteInterno}

		return w

	case 2:
		w := EntityStatusWrapper{Target: &r.Status}

		return w

	}
	panic("Unknown field index")
}

func (r *Contrato) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Contrato) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Contrato) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Contrato) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Contrato) HintSize(int)                     { panic("Unsupported operation") }
func (_ Contrato) Finalize()                        {}

func (_ Contrato) AvroCRC64Fingerprint() []byte {
	return []byte(ContratoAvroCRC64Fingerprint)
}

func (r Contrato) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["codigoDeContratoInterno"], err = json.Marshal(r.CodigoDeContratoInterno)
	if err != nil {
		return nil, err
	}
	output["codigoDeClienteInterno"], err = json.Marshal(r.CodigoDeClienteInterno)
	if err != nil {
		return nil, err
	}
	output["status"], err = json.Marshal(r.Status)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Contrato) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
		if v, ok := fields["codigoDeClienteInterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeClienteInterno); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigoDeClienteInterno")
	}
	val = func() json.RawMessage {
		if v, ok := fields["status"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Status); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for status")
	}
	return nil
}
