// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AppDistribucionRetiroDeValoresRetiradoV2.avsc
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

type AppDistribucionRetiroDeValoresRetiradoV2 struct {
	Pagos []PagoV2 `json:"pagos"`

	FechaGeneracion string `json:"fechaGeneracion"`

	Envios []Envio `json:"envios"`

	Usuario Usuario `json:"usuario"`

	EsRectificacion *UnionNullBool `json:"esRectificacion"`
}

const AppDistribucionRetiroDeValoresRetiradoV2AvroCRC64Fingerprint = "\u07b3\x8ch\U000589fb"

func NewAppDistribucionRetiroDeValoresRetiradoV2() AppDistribucionRetiroDeValoresRetiradoV2 {
	r := AppDistribucionRetiroDeValoresRetiradoV2{}
	r.Pagos = make([]PagoV2, 0)

	r.Envios = make([]Envio, 0)

	r.Usuario = NewUsuario()

	return r
}

func DeserializeAppDistribucionRetiroDeValoresRetiradoV2(r io.Reader) (AppDistribucionRetiroDeValoresRetiradoV2, error) {
	t := NewAppDistribucionRetiroDeValoresRetiradoV2()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAppDistribucionRetiroDeValoresRetiradoV2FromSchema(r io.Reader, schema string) (AppDistribucionRetiroDeValoresRetiradoV2, error) {
	t := NewAppDistribucionRetiroDeValoresRetiradoV2()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAppDistribucionRetiroDeValoresRetiradoV2(r AppDistribucionRetiroDeValoresRetiradoV2, w io.Writer) error {
	var err error
	err = writeArrayPagoV2(r.Pagos, w)
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
	err = writeUnionNullBool(r.EsRectificacion, w)
	if err != nil {
		return err
	}
	return err
}

func (r AppDistribucionRetiroDeValoresRetiradoV2) Serialize(w io.Writer) error {
	return writeAppDistribucionRetiroDeValoresRetiradoV2(r, w)
}

func (r AppDistribucionRetiroDeValoresRetiradoV2) Schema() string {
	return "{\"fields\":[{\"name\":\"pagos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"tipoPago\",\"type\":\"string\"},{\"name\":\"montoCobrado\",\"type\":\"string\"},{\"name\":\"comprobante\",\"type\":\"string\"},{\"name\":\"detalles\",\"type\":[\"null\",{\"fields\":[{\"name\":\"bancoEmisor\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaDeEmision\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaDeVencimiento\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroDeCheque\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroDeRetencion\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroNotaCredito\",\"type\":[\"null\",\"string\"]},{\"name\":\"notas\",\"type\":[\"null\",\"string\"]}],\"name\":\"DetallePagoV2\",\"type\":\"record\"}]}],\"name\":\"PagoV2\",\"type\":\"record\"},\"type\":\"array\"}},{\"name\":\"fechaGeneracion\",\"type\":\"string\"},{\"name\":\"envios\",\"type\":{\"items\":{\"fields\":[{\"name\":\"numeroSeguimiento\",\"type\":\"string\"},{\"name\":\"unidadOperativa\",\"type\":\"int\"},{\"name\":\"tareaId\",\"type\":\"int\"},{\"name\":\"montoSolicitado\",\"type\":\"string\"}],\"name\":\"Envio\",\"type\":\"record\"},\"type\":\"array\"}},{\"name\":\"usuario\",\"type\":{\"fields\":[{\"name\":\"username\",\"type\":\"string\"},{\"name\":\"dni\",\"type\":\"string\"},{\"name\":\"nombre\",\"type\":\"string\"},{\"name\":\"apellido\",\"type\":\"string\"}],\"name\":\"Usuario\",\"type\":\"record\"}},{\"name\":\"esRectificacion\",\"type\":[\"null\",\"boolean\"]}],\"name\":\"Andreani.MobileOperacionesProducer.Events.Record.AppDistribucionRetiroDeValoresRetiradoV2\",\"type\":\"record\"}"
}

func (r AppDistribucionRetiroDeValoresRetiradoV2) SchemaName() string {
	return "Andreani.MobileOperacionesProducer.Events.Record.AppDistribucionRetiroDeValoresRetiradoV2"
}

func (_ AppDistribucionRetiroDeValoresRetiradoV2) SetBoolean(v bool)  { panic("Unsupported operation") }
func (_ AppDistribucionRetiroDeValoresRetiradoV2) SetInt(v int32)     { panic("Unsupported operation") }
func (_ AppDistribucionRetiroDeValoresRetiradoV2) SetLong(v int64)    { panic("Unsupported operation") }
func (_ AppDistribucionRetiroDeValoresRetiradoV2) SetFloat(v float32) { panic("Unsupported operation") }
func (_ AppDistribucionRetiroDeValoresRetiradoV2) SetDouble(v float64) {
	panic("Unsupported operation")
}
func (_ AppDistribucionRetiroDeValoresRetiradoV2) SetBytes(v []byte)  { panic("Unsupported operation") }
func (_ AppDistribucionRetiroDeValoresRetiradoV2) SetString(v string) { panic("Unsupported operation") }
func (_ AppDistribucionRetiroDeValoresRetiradoV2) SetUnionElem(v int64) {
	panic("Unsupported operation")
}

func (r *AppDistribucionRetiroDeValoresRetiradoV2) Get(i int) types.Field {
	switch i {
	case 0:
		r.Pagos = make([]PagoV2, 0)

		w := ArrayPagoV2Wrapper{Target: &r.Pagos}

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

	case 4:
		r.EsRectificacion = NewUnionNullBool()

		return r.EsRectificacion
	}
	panic("Unknown field index")
}

func (r *AppDistribucionRetiroDeValoresRetiradoV2) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *AppDistribucionRetiroDeValoresRetiradoV2) NullField(i int) {
	switch i {
	case 4:
		r.EsRectificacion = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ AppDistribucionRetiroDeValoresRetiradoV2) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ AppDistribucionRetiroDeValoresRetiradoV2) AppendArray() types.Field {
	panic("Unsupported operation")
}
func (_ AppDistribucionRetiroDeValoresRetiradoV2) HintSize(int) { panic("Unsupported operation") }
func (_ AppDistribucionRetiroDeValoresRetiradoV2) Finalize()    {}

func (_ AppDistribucionRetiroDeValoresRetiradoV2) AvroCRC64Fingerprint() []byte {
	return []byte(AppDistribucionRetiroDeValoresRetiradoV2AvroCRC64Fingerprint)
}

func (r AppDistribucionRetiroDeValoresRetiradoV2) MarshalJSON() ([]byte, error) {
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
	output["esRectificacion"], err = json.Marshal(r.EsRectificacion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *AppDistribucionRetiroDeValoresRetiradoV2) UnmarshalJSON(data []byte) error {
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
	val = func() json.RawMessage {
		if v, ok := fields["esRectificacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EsRectificacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for esRectificacion")
	}
	return nil
}
