// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     DevolucionReembolsoAlCliente.avsc
 */
package TrazasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type DevolucionReembolsoAlCliente struct {
	Traza Traza `json:"traza"`

	NumeroDeEnvio string `json:"numeroDeEnvio"`

	Factura string `json:"factura"`

	Fecha string `json:"fecha"`

	Motivo string `json:"motivo"`
}

const DevolucionReembolsoAlClienteAvroCRC64Fingerprint = "\x92\xa4\x03<\xe4l/a"

func NewDevolucionReembolsoAlCliente() DevolucionReembolsoAlCliente {
	r := DevolucionReembolsoAlCliente{}
	r.Traza = NewTraza()

	return r
}

func DeserializeDevolucionReembolsoAlCliente(r io.Reader) (DevolucionReembolsoAlCliente, error) {
	t := NewDevolucionReembolsoAlCliente()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDevolucionReembolsoAlClienteFromSchema(r io.Reader, schema string) (DevolucionReembolsoAlCliente, error) {
	t := NewDevolucionReembolsoAlCliente()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDevolucionReembolsoAlCliente(r DevolucionReembolsoAlCliente, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroDeEnvio, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Factura, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Fecha, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Motivo, w)
	if err != nil {
		return err
	}
	return err
}

func (r DevolucionReembolsoAlCliente) Serialize(w io.Writer) error {
	return writeDevolucionReembolsoAlCliente(r, w)
}

func (r DevolucionReembolsoAlCliente) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"LogisticShipmentType\",\"type\":[\"null\",\"string\"]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"name\":\"numeroDeEnvio\",\"type\":\"string\"},{\"name\":\"factura\",\"type\":\"string\"},{\"name\":\"fecha\",\"type\":\"string\"},{\"name\":\"motivo\",\"type\":\"string\"}],\"name\":\"Integracion.Esquemas.Trazas.DevolucionReembolsoAlCliente\",\"type\":\"record\"}"
}

func (r DevolucionReembolsoAlCliente) SchemaName() string {
	return "Integracion.Esquemas.Trazas.DevolucionReembolsoAlCliente"
}

func (_ DevolucionReembolsoAlCliente) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DevolucionReembolsoAlCliente) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DevolucionReembolsoAlCliente) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DevolucionReembolsoAlCliente) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DevolucionReembolsoAlCliente) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DevolucionReembolsoAlCliente) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DevolucionReembolsoAlCliente) SetString(v string)   { panic("Unsupported operation") }
func (_ DevolucionReembolsoAlCliente) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DevolucionReembolsoAlCliente) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		w := types.String{Target: &r.NumeroDeEnvio}

		return w

	case 2:
		w := types.String{Target: &r.Factura}

		return w

	case 3:
		w := types.String{Target: &r.Fecha}

		return w

	case 4:
		w := types.String{Target: &r.Motivo}

		return w

	}
	panic("Unknown field index")
}

func (r *DevolucionReembolsoAlCliente) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *DevolucionReembolsoAlCliente) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ DevolucionReembolsoAlCliente) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ DevolucionReembolsoAlCliente) AppendArray() types.Field { panic("Unsupported operation") }
func (_ DevolucionReembolsoAlCliente) HintSize(int)             { panic("Unsupported operation") }
func (_ DevolucionReembolsoAlCliente) Finalize()                {}

func (_ DevolucionReembolsoAlCliente) AvroCRC64Fingerprint() []byte {
	return []byte(DevolucionReembolsoAlClienteAvroCRC64Fingerprint)
}

func (r DevolucionReembolsoAlCliente) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["numeroDeEnvio"], err = json.Marshal(r.NumeroDeEnvio)
	if err != nil {
		return nil, err
	}
	output["factura"], err = json.Marshal(r.Factura)
	if err != nil {
		return nil, err
	}
	output["fecha"], err = json.Marshal(r.Fecha)
	if err != nil {
		return nil, err
	}
	output["motivo"], err = json.Marshal(r.Motivo)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DevolucionReembolsoAlCliente) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["traza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Traza); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for traza")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeEnvio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroDeEnvio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["factura"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Factura); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for factura")
	}
	val = func() json.RawMessage {
		if v, ok := fields["fecha"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Fecha); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fecha")
	}
	val = func() json.RawMessage {
		if v, ok := fields["motivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Motivo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for motivo")
	}
	return nil
}
