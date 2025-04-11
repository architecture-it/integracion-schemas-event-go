// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     UOHdrClosureDistribution.avsc
 */
package UOPublisherHdrEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type NovedadEnvio struct {
	NumeroDeEnvio string `json:"NumeroDeEnvio"`

	Motivo string `json:"Motivo"`

	SubMotivo string `json:"SubMotivo"`

	Observaciones string `json:"Observaciones"`
}

const NovedadEnvioAvroCRC64Fingerprint = "\xfd$q\x8eΎvp"

func NewNovedadEnvio() NovedadEnvio {
	r := NovedadEnvio{}
	return r
}

func DeserializeNovedadEnvio(r io.Reader) (NovedadEnvio, error) {
	t := NewNovedadEnvio()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNovedadEnvioFromSchema(r io.Reader, schema string) (NovedadEnvio, error) {
	t := NewNovedadEnvio()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNovedadEnvio(r NovedadEnvio, w io.Writer) error {
	var err error
	err = vm.WriteString(r.NumeroDeEnvio, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Motivo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.SubMotivo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Observaciones, w)
	if err != nil {
		return err
	}
	return err
}

func (r NovedadEnvio) Serialize(w io.Writer) error {
	return writeNovedadEnvio(r, w)
}

func (r NovedadEnvio) Schema() string {
	return "{\"fields\":[{\"name\":\"NumeroDeEnvio\",\"type\":\"string\"},{\"name\":\"Motivo\",\"type\":\"string\"},{\"name\":\"SubMotivo\",\"type\":\"string\"},{\"name\":\"Observaciones\",\"type\":\"string\"}],\"name\":\"Andreani.UOPublisherHdr.Events.Common.NovedadEnvio\",\"type\":\"record\"}"
}

func (r NovedadEnvio) SchemaName() string {
	return "Andreani.UOPublisherHdr.Events.Common.NovedadEnvio"
}

func (_ NovedadEnvio) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NovedadEnvio) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NovedadEnvio) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NovedadEnvio) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NovedadEnvio) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NovedadEnvio) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NovedadEnvio) SetString(v string)   { panic("Unsupported operation") }
func (_ NovedadEnvio) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NovedadEnvio) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.NumeroDeEnvio}

		return w

	case 1:
		w := types.String{Target: &r.Motivo}

		return w

	case 2:
		w := types.String{Target: &r.SubMotivo}

		return w

	case 3:
		w := types.String{Target: &r.Observaciones}

		return w

	}
	panic("Unknown field index")
}

func (r *NovedadEnvio) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NovedadEnvio) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ NovedadEnvio) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NovedadEnvio) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NovedadEnvio) HintSize(int)                     { panic("Unsupported operation") }
func (_ NovedadEnvio) Finalize()                        {}

func (_ NovedadEnvio) AvroCRC64Fingerprint() []byte {
	return []byte(NovedadEnvioAvroCRC64Fingerprint)
}

func (r NovedadEnvio) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["NumeroDeEnvio"], err = json.Marshal(r.NumeroDeEnvio)
	if err != nil {
		return nil, err
	}
	output["Motivo"], err = json.Marshal(r.Motivo)
	if err != nil {
		return nil, err
	}
	output["SubMotivo"], err = json.Marshal(r.SubMotivo)
	if err != nil {
		return nil, err
	}
	output["Observaciones"], err = json.Marshal(r.Observaciones)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NovedadEnvio) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeEnvio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NumeroDeEnvio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Motivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Motivo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Motivo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SubMotivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SubMotivo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SubMotivo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Observaciones"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Observaciones); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Observaciones")
	}
	return nil
}
