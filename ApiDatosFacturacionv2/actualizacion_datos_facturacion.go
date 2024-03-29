// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     DatosFacturacionSolicitadaV2.avsc
 */
package ApiDatosFacturacionv2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ActualizacionDatosFacturacion struct {
	OrdenExterna string `json:"OrdenExterna"`

	Remito string `json:"Remito"`

	Factura string `json:"Factura"`

	ValorDeclarado *UnionNullString `json:"ValorDeclarado"`

	ValorSeguro *UnionNullString `json:"ValorSeguro"`

	FechaFactura *UnionNullString `json:"FechaFactura"`
}

const ActualizacionDatosFacturacionAvroCRC64Fingerprint = "lK`FD\xc4\xd1;"

func NewActualizacionDatosFacturacion() ActualizacionDatosFacturacion {
	r := ActualizacionDatosFacturacion{}
	r.ValorDeclarado = nil
	r.ValorSeguro = nil
	r.FechaFactura = nil
	return r
}

func DeserializeActualizacionDatosFacturacion(r io.Reader) (ActualizacionDatosFacturacion, error) {
	t := NewActualizacionDatosFacturacion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeActualizacionDatosFacturacionFromSchema(r io.Reader, schema string) (ActualizacionDatosFacturacion, error) {
	t := NewActualizacionDatosFacturacion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeActualizacionDatosFacturacion(r ActualizacionDatosFacturacion, w io.Writer) error {
	var err error
	err = vm.WriteString(r.OrdenExterna, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Remito, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Factura, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ValorDeclarado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ValorSeguro, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaFactura, w)
	if err != nil {
		return err
	}
	return err
}

func (r ActualizacionDatosFacturacion) Serialize(w io.Writer) error {
	return writeActualizacionDatosFacturacion(r, w)
}

func (r ActualizacionDatosFacturacion) Schema() string {
	return "{\"fields\":[{\"name\":\"OrdenExterna\",\"type\":\"string\"},{\"name\":\"Remito\",\"type\":\"string\"},{\"name\":\"Factura\",\"type\":\"string\"},{\"default\":null,\"name\":\"ValorDeclarado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ValorSeguro\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaFactura\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.WapDatosFacturacionv2.Events.Record.ActualizacionDatosFacturacion\",\"type\":\"record\"}"
}

func (r ActualizacionDatosFacturacion) SchemaName() string {
	return "Andreani.WapDatosFacturacionv2.Events.Record.ActualizacionDatosFacturacion"
}

func (_ ActualizacionDatosFacturacion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ActualizacionDatosFacturacion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ActualizacionDatosFacturacion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ActualizacionDatosFacturacion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ActualizacionDatosFacturacion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ActualizacionDatosFacturacion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ActualizacionDatosFacturacion) SetString(v string)   { panic("Unsupported operation") }
func (_ ActualizacionDatosFacturacion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ActualizacionDatosFacturacion) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.OrdenExterna}

		return w

	case 1:
		w := types.String{Target: &r.Remito}

		return w

	case 2:
		w := types.String{Target: &r.Factura}

		return w

	case 3:
		r.ValorDeclarado = NewUnionNullString()

		return r.ValorDeclarado
	case 4:
		r.ValorSeguro = NewUnionNullString()

		return r.ValorSeguro
	case 5:
		r.FechaFactura = NewUnionNullString()

		return r.FechaFactura
	}
	panic("Unknown field index")
}

func (r *ActualizacionDatosFacturacion) SetDefault(i int) {
	switch i {
	case 3:
		r.ValorDeclarado = nil
		return
	case 4:
		r.ValorSeguro = nil
		return
	case 5:
		r.FechaFactura = nil
		return
	}
	panic("Unknown field index")
}

func (r *ActualizacionDatosFacturacion) NullField(i int) {
	switch i {
	case 3:
		r.ValorDeclarado = nil
		return
	case 4:
		r.ValorSeguro = nil
		return
	case 5:
		r.FechaFactura = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ActualizacionDatosFacturacion) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ActualizacionDatosFacturacion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ ActualizacionDatosFacturacion) HintSize(int)             { panic("Unsupported operation") }
func (_ ActualizacionDatosFacturacion) Finalize()                {}

func (_ ActualizacionDatosFacturacion) AvroCRC64Fingerprint() []byte {
	return []byte(ActualizacionDatosFacturacionAvroCRC64Fingerprint)
}

func (r ActualizacionDatosFacturacion) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["OrdenExterna"], err = json.Marshal(r.OrdenExterna)
	if err != nil {
		return nil, err
	}
	output["Remito"], err = json.Marshal(r.Remito)
	if err != nil {
		return nil, err
	}
	output["Factura"], err = json.Marshal(r.Factura)
	if err != nil {
		return nil, err
	}
	output["ValorDeclarado"], err = json.Marshal(r.ValorDeclarado)
	if err != nil {
		return nil, err
	}
	output["ValorSeguro"], err = json.Marshal(r.ValorSeguro)
	if err != nil {
		return nil, err
	}
	output["FechaFactura"], err = json.Marshal(r.FechaFactura)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ActualizacionDatosFacturacion) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["OrdenExterna"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenExterna); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrdenExterna")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Remito"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Remito); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Remito")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Factura"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Factura); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Factura")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ValorDeclarado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ValorDeclarado); err != nil {
			return err
		}
	} else {
		r.ValorDeclarado = NewUnionNullString()

		r.ValorDeclarado = nil
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
		r.ValorSeguro = NewUnionNullString()

		r.ValorSeguro = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaFactura"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaFactura); err != nil {
			return err
		}
	} else {
		r.FechaFactura = NewUnionNullString()

		r.FechaFactura = nil
	}
	return nil
}
