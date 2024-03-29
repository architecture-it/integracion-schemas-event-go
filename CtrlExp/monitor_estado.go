// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     MonitorEstado.avsc
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

type MonitorEstado struct {
	Nombre string `json:"Nombre"`

	Tipo string `json:"Tipo"`

	Estado bool `json:"Estado"`

	Ts int64 `json:"Ts"`

	OtrosContadores *UnionNullArrayMonitorEstadoOtrosContadores `json:"OtrosContadores"`

	OtrosSemaforos *UnionNullArrayMonitorEstadoOtrosSemaforos `json:"OtrosSemaforos"`

	Otros *UnionNullArrayMonitorEstadoOtros `json:"Otros"`

	SchemaDb *UnionNullString `json:"SchemaDb"`
}

const MonitorEstadoAvroCRC64Fingerprint = "\t/\xa9^7\xb9\xcf\x1b"

func NewMonitorEstado() MonitorEstado {
	r := MonitorEstado{}
	r.SchemaDb = nil
	return r
}

func DeserializeMonitorEstado(r io.Reader) (MonitorEstado, error) {
	t := NewMonitorEstado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMonitorEstadoFromSchema(r io.Reader, schema string) (MonitorEstado, error) {
	t := NewMonitorEstado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMonitorEstado(r MonitorEstado, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Nombre, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Tipo, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Estado, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Ts, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayMonitorEstadoOtrosContadores(r.OtrosContadores, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayMonitorEstadoOtrosSemaforos(r.OtrosSemaforos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayMonitorEstadoOtros(r.Otros, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.SchemaDb, w)
	if err != nil {
		return err
	}
	return err
}

func (r MonitorEstado) Serialize(w io.Writer) error {
	return writeMonitorEstado(r, w)
}

func (r MonitorEstado) Schema() string {
	return "{\"fields\":[{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Tipo\",\"type\":\"string\"},{\"name\":\"Estado\",\"type\":\"boolean\"},{\"name\":\"Ts\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"OtrosContadores\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"Contador\",\"type\":\"string\"},{\"name\":\"Valor\",\"type\":\"int\"}],\"name\":\"MonitorEstadoOtrosContadores\",\"type\":\"record\"},\"type\":\"array\"}]},{\"name\":\"OtrosSemaforos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"Semaforo\",\"type\":\"string\"},{\"name\":\"Valor\",\"type\":\"boolean\"}],\"name\":\"MonitorEstadoOtrosSemaforos\",\"type\":\"record\"},\"type\":\"array\"}]},{\"name\":\"Otros\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"Referencia\",\"type\":\"string\"}],\"name\":\"MonitorEstadoOtros\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"SchemaDb\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.CtrlExp.Events.Record.MonitorEstado\",\"type\":\"record\"}"
}

func (r MonitorEstado) SchemaName() string {
	return "Andreani.CtrlExp.Events.Record.MonitorEstado"
}

func (_ MonitorEstado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ MonitorEstado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ MonitorEstado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ MonitorEstado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ MonitorEstado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ MonitorEstado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ MonitorEstado) SetString(v string)   { panic("Unsupported operation") }
func (_ MonitorEstado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MonitorEstado) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Nombre}

		return w

	case 1:
		w := types.String{Target: &r.Tipo}

		return w

	case 2:
		w := types.Boolean{Target: &r.Estado}

		return w

	case 3:
		w := types.Long{Target: &r.Ts}

		return w

	case 4:
		r.OtrosContadores = NewUnionNullArrayMonitorEstadoOtrosContadores()

		return r.OtrosContadores
	case 5:
		r.OtrosSemaforos = NewUnionNullArrayMonitorEstadoOtrosSemaforos()

		return r.OtrosSemaforos
	case 6:
		r.Otros = NewUnionNullArrayMonitorEstadoOtros()

		return r.Otros
	case 7:
		r.SchemaDb = NewUnionNullString()

		return r.SchemaDb
	}
	panic("Unknown field index")
}

func (r *MonitorEstado) SetDefault(i int) {
	switch i {
	case 7:
		r.SchemaDb = nil
		return
	}
	panic("Unknown field index")
}

func (r *MonitorEstado) NullField(i int) {
	switch i {
	case 4:
		r.OtrosContadores = nil
		return
	case 5:
		r.OtrosSemaforos = nil
		return
	case 6:
		r.Otros = nil
		return
	case 7:
		r.SchemaDb = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ MonitorEstado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ MonitorEstado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ MonitorEstado) HintSize(int)                     { panic("Unsupported operation") }
func (_ MonitorEstado) Finalize()                        {}

func (_ MonitorEstado) AvroCRC64Fingerprint() []byte {
	return []byte(MonitorEstadoAvroCRC64Fingerprint)
}

func (r MonitorEstado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	output["Tipo"], err = json.Marshal(r.Tipo)
	if err != nil {
		return nil, err
	}
	output["Estado"], err = json.Marshal(r.Estado)
	if err != nil {
		return nil, err
	}
	output["Ts"], err = json.Marshal(r.Ts)
	if err != nil {
		return nil, err
	}
	output["OtrosContadores"], err = json.Marshal(r.OtrosContadores)
	if err != nil {
		return nil, err
	}
	output["OtrosSemaforos"], err = json.Marshal(r.OtrosSemaforos)
	if err != nil {
		return nil, err
	}
	output["Otros"], err = json.Marshal(r.Otros)
	if err != nil {
		return nil, err
	}
	output["SchemaDb"], err = json.Marshal(r.SchemaDb)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *MonitorEstado) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Nombre"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Nombre); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Nombre")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Tipo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tipo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Tipo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Estado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Estado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Ts"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ts); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Ts")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OtrosContadores"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OtrosContadores); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OtrosContadores")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OtrosSemaforos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OtrosSemaforos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OtrosSemaforos")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Otros"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Otros); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Otros")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SchemaDb"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SchemaDb); err != nil {
			return err
		}
	} else {
		r.SchemaDb = NewUnionNullString()

		r.SchemaDb = nil
	}
	return nil
}
