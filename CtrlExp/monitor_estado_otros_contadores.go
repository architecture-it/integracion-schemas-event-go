// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     MonitorEstadoOtrosContadores.avsc
 */
package CtrlExpEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type MonitorEstadoOtrosContadores struct {
	Contador string `json:"Contador"`

	Valor int32 `json:"Valor"`
}

const MonitorEstadoOtrosContadoresAvroCRC64Fingerprint = "\xd8\xd9Nb\xe6\x17\xaco"

func NewMonitorEstadoOtrosContadores() MonitorEstadoOtrosContadores {
	r := MonitorEstadoOtrosContadores{}
	return r
}

func DeserializeMonitorEstadoOtrosContadores(r io.Reader) (MonitorEstadoOtrosContadores, error) {
	t := NewMonitorEstadoOtrosContadores()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMonitorEstadoOtrosContadoresFromSchema(r io.Reader, schema string) (MonitorEstadoOtrosContadores, error) {
	t := NewMonitorEstadoOtrosContadores()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMonitorEstadoOtrosContadores(r MonitorEstadoOtrosContadores, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Contador, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Valor, w)
	if err != nil {
		return err
	}
	return err
}

func (r MonitorEstadoOtrosContadores) Serialize(w io.Writer) error {
	return writeMonitorEstadoOtrosContadores(r, w)
}

func (r MonitorEstadoOtrosContadores) Schema() string {
	return "{\"fields\":[{\"name\":\"Contador\",\"type\":\"string\"},{\"name\":\"Valor\",\"type\":\"int\"}],\"name\":\"Andreani.CtrlExp.Events.Record.MonitorEstadoOtrosContadores\",\"type\":\"record\"}"
}

func (r MonitorEstadoOtrosContadores) SchemaName() string {
	return "Andreani.CtrlExp.Events.Record.MonitorEstadoOtrosContadores"
}

func (_ MonitorEstadoOtrosContadores) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ MonitorEstadoOtrosContadores) SetInt(v int32)       { panic("Unsupported operation") }
func (_ MonitorEstadoOtrosContadores) SetLong(v int64)      { panic("Unsupported operation") }
func (_ MonitorEstadoOtrosContadores) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ MonitorEstadoOtrosContadores) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ MonitorEstadoOtrosContadores) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ MonitorEstadoOtrosContadores) SetString(v string)   { panic("Unsupported operation") }
func (_ MonitorEstadoOtrosContadores) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MonitorEstadoOtrosContadores) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Contador}

		return w

	case 1:
		w := types.Int{Target: &r.Valor}

		return w

	}
	panic("Unknown field index")
}

func (r *MonitorEstadoOtrosContadores) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *MonitorEstadoOtrosContadores) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ MonitorEstadoOtrosContadores) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ MonitorEstadoOtrosContadores) AppendArray() types.Field { panic("Unsupported operation") }
func (_ MonitorEstadoOtrosContadores) HintSize(int)             { panic("Unsupported operation") }
func (_ MonitorEstadoOtrosContadores) Finalize()                {}

func (_ MonitorEstadoOtrosContadores) AvroCRC64Fingerprint() []byte {
	return []byte(MonitorEstadoOtrosContadoresAvroCRC64Fingerprint)
}

func (r MonitorEstadoOtrosContadores) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Contador"], err = json.Marshal(r.Contador)
	if err != nil {
		return nil, err
	}
	output["Valor"], err = json.Marshal(r.Valor)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *MonitorEstadoOtrosContadores) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Contador"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contador); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Contador")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Valor"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Valor); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Valor")
	}
	return nil
}
