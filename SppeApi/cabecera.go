// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     OrdenDeEnvioRechazada.avsc
 */
package SppeApiEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Cabecera struct {
	OrdenWh string `json:"OrdenWh"`

	OrdenCliente string `json:"OrdenCliente"`

	ContratoTMS string `json:"ContratoTMS"`

	ValorSeguro *UnionNullString `json:"ValorSeguro"`

	MensajeError string `json:"MensajeError"`
}

const CabeceraAvroCRC64Fingerprint = "\xb5\x11\x8c\xf1\xe3\xees\x8c"

func NewCabecera() Cabecera {
	r := Cabecera{}
	return r
}

func DeserializeCabecera(r io.Reader) (Cabecera, error) {
	t := NewCabecera()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCabeceraFromSchema(r io.Reader, schema string) (Cabecera, error) {
	t := NewCabecera()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCabecera(r Cabecera, w io.Writer) error {
	var err error
	err = vm.WriteString(r.OrdenWh, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.OrdenCliente, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ContratoTMS, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ValorSeguro, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.MensajeError, w)
	if err != nil {
		return err
	}
	return err
}

func (r Cabecera) Serialize(w io.Writer) error {
	return writeCabecera(r, w)
}

func (r Cabecera) Schema() string {
	return "{\"fields\":[{\"name\":\"OrdenWh\",\"type\":\"string\"},{\"name\":\"OrdenCliente\",\"type\":\"string\"},{\"name\":\"ContratoTMS\",\"type\":\"string\"},{\"name\":\"ValorSeguro\",\"type\":[\"null\",\"string\"]},{\"name\":\"MensajeError\",\"type\":\"string\"}],\"name\":\"Andreani.SppeApi.Events.OrdenDeEnvioRechazadaCommon.Cabecera\",\"type\":\"record\"}"
}

func (r Cabecera) SchemaName() string {
	return "Andreani.SppeApi.Events.OrdenDeEnvioRechazadaCommon.Cabecera"
}

func (_ Cabecera) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Cabecera) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Cabecera) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Cabecera) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Cabecera) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Cabecera) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Cabecera) SetString(v string)   { panic("Unsupported operation") }
func (_ Cabecera) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Cabecera) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.OrdenWh}

		return w

	case 1:
		w := types.String{Target: &r.OrdenCliente}

		return w

	case 2:
		w := types.String{Target: &r.ContratoTMS}

		return w

	case 3:
		r.ValorSeguro = NewUnionNullString()

		return r.ValorSeguro
	case 4:
		w := types.String{Target: &r.MensajeError}

		return w

	}
	panic("Unknown field index")
}

func (r *Cabecera) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Cabecera) NullField(i int) {
	switch i {
	case 3:
		r.ValorSeguro = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Cabecera) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Cabecera) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Cabecera) HintSize(int)                     { panic("Unsupported operation") }
func (_ Cabecera) Finalize()                        {}

func (_ Cabecera) AvroCRC64Fingerprint() []byte {
	return []byte(CabeceraAvroCRC64Fingerprint)
}

func (r Cabecera) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["OrdenWh"], err = json.Marshal(r.OrdenWh)
	if err != nil {
		return nil, err
	}
	output["OrdenCliente"], err = json.Marshal(r.OrdenCliente)
	if err != nil {
		return nil, err
	}
	output["ContratoTMS"], err = json.Marshal(r.ContratoTMS)
	if err != nil {
		return nil, err
	}
	output["ValorSeguro"], err = json.Marshal(r.ValorSeguro)
	if err != nil {
		return nil, err
	}
	output["MensajeError"], err = json.Marshal(r.MensajeError)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Cabecera) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["OrdenWh"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenWh); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrdenWh")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OrdenCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenCliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrdenCliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ContratoTMS"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContratoTMS); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ContratoTMS")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ValorSeguro"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ValorSeguro); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ValorSeguro")
	}
	val = func() json.RawMessage {
		if v, ok := fields["MensajeError"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MensajeError); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for MensajeError")
	}
	return nil
}
