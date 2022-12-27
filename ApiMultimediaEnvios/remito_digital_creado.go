// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     RemitoDigitalCreado.avsc
 */
package ApiMultimediaEnviosEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type RemitoDigitalCreado struct {
	Timestamp string `json:"Timestamp"`

	Remitente string `json:"Remitente"`

	IDTransaccion string `json:"IDTransaccion"`

	Contrato string `json:"Contrato"`

	Almacen string `json:"Almacen"`

	Planta string `json:"Planta"`

	NumeroDeEnvio string `json:"NumeroDeEnvio"`

	ImagenURL string `json:"ImagenURL"`

	Remito Anexo `json:"Remito"`
}

const RemitoDigitalCreadoAvroCRC64Fingerprint = "!\xd1fȗ[\xec\x90"

func NewRemitoDigitalCreado() RemitoDigitalCreado {
	r := RemitoDigitalCreado{}
	r.Remito = NewAnexo()

	return r
}

func DeserializeRemitoDigitalCreado(r io.Reader) (RemitoDigitalCreado, error) {
	t := NewRemitoDigitalCreado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRemitoDigitalCreadoFromSchema(r io.Reader, schema string) (RemitoDigitalCreado, error) {
	t := NewRemitoDigitalCreado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRemitoDigitalCreado(r RemitoDigitalCreado, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Timestamp, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Remitente, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.IDTransaccion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Contrato, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Almacen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Planta, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroDeEnvio, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ImagenURL, w)
	if err != nil {
		return err
	}
	err = writeAnexo(r.Remito, w)
	if err != nil {
		return err
	}
	return err
}

func (r RemitoDigitalCreado) Serialize(w io.Writer) error {
	return writeRemitoDigitalCreado(r, w)
}

func (r RemitoDigitalCreado) Schema() string {
	return "{\"fields\":[{\"name\":\"Timestamp\",\"type\":\"string\"},{\"name\":\"Remitente\",\"type\":\"string\"},{\"name\":\"IDTransaccion\",\"type\":\"string\"},{\"name\":\"Contrato\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Planta\",\"type\":\"string\"},{\"name\":\"NumeroDeEnvio\",\"type\":\"string\"},{\"name\":\"ImagenURL\",\"type\":\"string\"},{\"name\":\"Remito\",\"type\":{\"fields\":[{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"NumeroDePedido\",\"type\":\"string\"},{\"name\":\"NumeroFacturaLegal\",\"type\":\"string\"},{\"name\":\"NumeroRemito\",\"type\":\"string\"},{\"name\":\"RemitoURL\",\"type\":\"string\"}],\"name\":\"Anexo\",\"namespace\":\"Andreani.ApiMultimediaEnvios.Events.Common\",\"type\":\"record\"}}],\"name\":\"Andreani.ApiMultimediaEnvios.Events.Record.RemitoDigitalCreado\",\"type\":\"record\"}"
}

func (r RemitoDigitalCreado) SchemaName() string {
	return "Andreani.ApiMultimediaEnvios.Events.Record.RemitoDigitalCreado"
}

func (_ RemitoDigitalCreado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ RemitoDigitalCreado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ RemitoDigitalCreado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ RemitoDigitalCreado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ RemitoDigitalCreado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ RemitoDigitalCreado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ RemitoDigitalCreado) SetString(v string)   { panic("Unsupported operation") }
func (_ RemitoDigitalCreado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RemitoDigitalCreado) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Timestamp}

		return w

	case 1:
		w := types.String{Target: &r.Remitente}

		return w

	case 2:
		w := types.String{Target: &r.IDTransaccion}

		return w

	case 3:
		w := types.String{Target: &r.Contrato}

		return w

	case 4:
		w := types.String{Target: &r.Almacen}

		return w

	case 5:
		w := types.String{Target: &r.Planta}

		return w

	case 6:
		w := types.String{Target: &r.NumeroDeEnvio}

		return w

	case 7:
		w := types.String{Target: &r.ImagenURL}

		return w

	case 8:
		r.Remito = NewAnexo()

		w := types.Record{Target: &r.Remito}

		return w

	}
	panic("Unknown field index")
}

func (r *RemitoDigitalCreado) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *RemitoDigitalCreado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ RemitoDigitalCreado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ RemitoDigitalCreado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ RemitoDigitalCreado) HintSize(int)                     { panic("Unsupported operation") }
func (_ RemitoDigitalCreado) Finalize()                        {}

func (_ RemitoDigitalCreado) AvroCRC64Fingerprint() []byte {
	return []byte(RemitoDigitalCreadoAvroCRC64Fingerprint)
}

func (r RemitoDigitalCreado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Timestamp"], err = json.Marshal(r.Timestamp)
	if err != nil {
		return nil, err
	}
	output["Remitente"], err = json.Marshal(r.Remitente)
	if err != nil {
		return nil, err
	}
	output["IDTransaccion"], err = json.Marshal(r.IDTransaccion)
	if err != nil {
		return nil, err
	}
	output["Contrato"], err = json.Marshal(r.Contrato)
	if err != nil {
		return nil, err
	}
	output["Almacen"], err = json.Marshal(r.Almacen)
	if err != nil {
		return nil, err
	}
	output["Planta"], err = json.Marshal(r.Planta)
	if err != nil {
		return nil, err
	}
	output["NumeroDeEnvio"], err = json.Marshal(r.NumeroDeEnvio)
	if err != nil {
		return nil, err
	}
	output["ImagenURL"], err = json.Marshal(r.ImagenURL)
	if err != nil {
		return nil, err
	}
	output["Remito"], err = json.Marshal(r.Remito)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *RemitoDigitalCreado) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
		if v, ok := fields["Remitente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Remitente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Remitente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["IDTransaccion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IDTransaccion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IDTransaccion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Contrato"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contrato); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Contrato")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Almacen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Almacen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Almacen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Planta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Planta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Planta")
	}
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
		if v, ok := fields["ImagenURL"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ImagenURL); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ImagenURL")
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
	return nil
}