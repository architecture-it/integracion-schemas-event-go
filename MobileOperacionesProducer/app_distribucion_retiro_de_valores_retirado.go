// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AppDistribucionRetiroDeValoresRetirado.avsc
 */
package MobileOperacionesProducerEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type AppDistribucionRetiroDeValoresRetirado struct {
	Pagos []Pago `json:"pagos"`

	FechaGeneracion string `json:"fechaGeneracion"`

	Envios []Envio `json:"envios"`

	Usuario Usuario `json:"usuario"`
}

const AppDistribucionRetiroDeValoresRetiradoAvroCRC64Fingerprint = "3SXU\xa6s\xa1\x8a"

func NewAppDistribucionRetiroDeValoresRetirado() AppDistribucionRetiroDeValoresRetirado {
	r := AppDistribucionRetiroDeValoresRetirado{}
	r.Pagos = make([]Pago, 0)

	r.Envios = make([]Envio, 0)

	r.Usuario = NewUsuario()

	return r
}

func DeserializeAppDistribucionRetiroDeValoresRetirado(r io.Reader) (AppDistribucionRetiroDeValoresRetirado, error) {
	t := NewAppDistribucionRetiroDeValoresRetirado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAppDistribucionRetiroDeValoresRetiradoFromSchema(r io.Reader, schema string) (AppDistribucionRetiroDeValoresRetirado, error) {
	t := NewAppDistribucionRetiroDeValoresRetirado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAppDistribucionRetiroDeValoresRetirado(r AppDistribucionRetiroDeValoresRetirado, w io.Writer) error {
	var err error
	err = writeArrayPago(r.Pagos, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaGeneracion, w)
	if err != nil {
		return err
	}
	err = writeArrayEnvio(r.Envios, w)
	if err != nil {
		return err
	}
	err = writeUsuario(r.Usuario, w)
	if err != nil {
		return err
	}
	return err
}

func (r AppDistribucionRetiroDeValoresRetirado) Serialize(w io.Writer) error {
	return writeAppDistribucionRetiroDeValoresRetirado(r, w)
}

func (r AppDistribucionRetiroDeValoresRetirado) Schema() string {
	return "{\"fields\":[{\"name\":\"pagos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"tipoPago\",\"type\":\"string\"},{\"name\":\"montoCobrado\",\"type\":\"string\"},{\"name\":\"comprobante\",\"type\":\"string\"},{\"name\":\"detalles\",\"type\":[\"null\",{\"fields\":[{\"name\":\"bancoEmisor\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaDePago\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroDeCheque\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroDeRetencion\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroNotaCredito\",\"type\":[\"null\",\"string\"]},{\"name\":\"notas\",\"type\":[\"null\",\"string\"]}],\"name\":\"DetallePago\",\"type\":\"record\"}]}],\"name\":\"Pago\",\"type\":\"record\"},\"type\":\"array\"}},{\"name\":\"fechaGeneracion\",\"type\":\"string\"},{\"name\":\"envios\",\"type\":{\"items\":{\"fields\":[{\"name\":\"numeroSeguimiento\",\"type\":\"string\"},{\"name\":\"unidadOperativa\",\"type\":\"int\"},{\"name\":\"tareaId\",\"type\":\"int\"},{\"name\":\"montoSolicitado\",\"type\":\"string\"}],\"name\":\"Envio\",\"type\":\"record\"},\"type\":\"array\"}},{\"name\":\"usuario\",\"type\":{\"fields\":[{\"name\":\"username\",\"type\":\"string\"},{\"name\":\"dni\",\"type\":\"string\"},{\"name\":\"nombre\",\"type\":\"string\"},{\"name\":\"apellido\",\"type\":\"string\"}],\"name\":\"Usuario\",\"type\":\"record\"}}],\"name\":\"Andreani.MobileOperacionesProducer.Events.Record.AppDistribucionRetiroDeValoresRetirado\",\"type\":\"record\"}"
}

func (r AppDistribucionRetiroDeValoresRetirado) SchemaName() string {
	return "Andreani.MobileOperacionesProducer.Events.Record.AppDistribucionRetiroDeValoresRetirado"
}

func (_ AppDistribucionRetiroDeValoresRetirado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ AppDistribucionRetiroDeValoresRetirado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ AppDistribucionRetiroDeValoresRetirado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ AppDistribucionRetiroDeValoresRetirado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ AppDistribucionRetiroDeValoresRetirado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ AppDistribucionRetiroDeValoresRetirado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ AppDistribucionRetiroDeValoresRetirado) SetString(v string)   { panic("Unsupported operation") }
func (_ AppDistribucionRetiroDeValoresRetirado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AppDistribucionRetiroDeValoresRetirado) Get(i int) types.Field {
	switch i {
	case 0:
		r.Pagos = make([]Pago, 0)

		w := ArrayPagoWrapper{Target: &r.Pagos}

		return w

	case 1:
		w := types.String{Target: &r.FechaGeneracion}

		return w

	case 2:
		r.Envios = make([]Envio, 0)

		w := ArrayEnvioWrapper{Target: &r.Envios}

		return w

	case 3:
		r.Usuario = NewUsuario()

		w := types.Record{Target: &r.Usuario}

		return w

	}
	panic("Unknown field index")
}

func (r *AppDistribucionRetiroDeValoresRetirado) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *AppDistribucionRetiroDeValoresRetirado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ AppDistribucionRetiroDeValoresRetirado) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ AppDistribucionRetiroDeValoresRetirado) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ AppDistribucionRetiroDeValoresRetirado) HintSize(int) { panic("Unsupported operation") }
func (_ AppDistribucionRetiroDeValoresRetirado) Finalize()    {}

func (_ AppDistribucionRetiroDeValoresRetirado) AvroCRC64Fingerprint() []byte {
	return []byte(AppDistribucionRetiroDeValoresRetiradoAvroCRC64Fingerprint)
}

func (r AppDistribucionRetiroDeValoresRetirado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["pagos"], err = json.Marshal(r.Pagos)
	if err != nil {
		return nil, err
	}
	output["fechaGeneracion"], err = json.Marshal(r.FechaGeneracion)
	if err != nil {
		return nil, err
	}
	output["envios"], err = json.Marshal(r.Envios)
	if err != nil {
		return nil, err
	}
	output["usuario"], err = json.Marshal(r.Usuario)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *AppDistribucionRetiroDeValoresRetirado) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["pagos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Pagos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for pagos")
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaGeneracion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaGeneracion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fechaGeneracion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["envios"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Envios); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for envios")
	}
	val = func() json.RawMessage {
		if v, ok := fields["usuario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Usuario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for usuario")
	}
	return nil
}
