// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NovedadEventual.avsc
 */
package NotificacionesEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type NovedadEventual struct {
	TipoPendiente string `json:"tipoPendiente"`

	Motivo string `json:"motivo"`

	Email string `json:"email"`

	Telefono string `json:"telefono"`

	Canal string `json:"canal"`

	Envio string `json:"envio"`

	SucursalActual string `json:"sucursalActual"`

	Segmento string `json:"segmento"`

	Cliente string `json:"cliente"`

	Contrato string `json:"contrato"`

	TipoServicio string `json:"tipoServicio"`

	Destinatario string `json:"destinatario"`

	Domicilio string `json:"domicilio"`
}

const NovedadEventualAvroCRC64Fingerprint = "vy<\x0e\xae\x98\xcd\xf3"

func NewNovedadEventual() NovedadEventual {
	r := NovedadEventual{}
	return r
}

func DeserializeNovedadEventual(r io.Reader) (NovedadEventual, error) {
	t := NewNovedadEventual()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNovedadEventualFromSchema(r io.Reader, schema string) (NovedadEventual, error) {
	t := NewNovedadEventual()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNovedadEventual(r NovedadEventual, w io.Writer) error {
	var err error
	err = vm.WriteString(r.TipoPendiente, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Motivo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Email, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Telefono, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Canal, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Envio, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.SucursalActual, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Segmento, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Cliente, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Contrato, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TipoServicio, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Destinatario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Domicilio, w)
	if err != nil {
		return err
	}
	return err
}

func (r NovedadEventual) Serialize(w io.Writer) error {
	return writeNovedadEventual(r, w)
}

func (r NovedadEventual) Schema() string {
	return "{\"fields\":[{\"name\":\"tipoPendiente\",\"type\":\"string\"},{\"name\":\"motivo\",\"type\":\"string\"},{\"name\":\"email\",\"type\":\"string\"},{\"name\":\"telefono\",\"type\":\"string\"},{\"name\":\"canal\",\"type\":\"string\"},{\"name\":\"envio\",\"type\":\"string\"},{\"name\":\"sucursalActual\",\"type\":\"string\"},{\"name\":\"segmento\",\"type\":\"string\"},{\"name\":\"cliente\",\"type\":\"string\"},{\"name\":\"contrato\",\"type\":\"string\"},{\"name\":\"tipoServicio\",\"type\":\"string\"},{\"name\":\"destinatario\",\"type\":\"string\"},{\"name\":\"domicilio\",\"type\":\"string\"}],\"name\":\"Andreani.Notificaciones.Events.Records.NovedadEventual\",\"type\":\"record\"}"
}

func (r NovedadEventual) SchemaName() string {
	return "Andreani.Notificaciones.Events.Records.NovedadEventual"
}

func (_ NovedadEventual) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NovedadEventual) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NovedadEventual) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NovedadEventual) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NovedadEventual) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NovedadEventual) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NovedadEventual) SetString(v string)   { panic("Unsupported operation") }
func (_ NovedadEventual) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NovedadEventual) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.TipoPendiente}

		return w

	case 1:
		w := types.String{Target: &r.Motivo}

		return w

	case 2:
		w := types.String{Target: &r.Email}

		return w

	case 3:
		w := types.String{Target: &r.Telefono}

		return w

	case 4:
		w := types.String{Target: &r.Canal}

		return w

	case 5:
		w := types.String{Target: &r.Envio}

		return w

	case 6:
		w := types.String{Target: &r.SucursalActual}

		return w

	case 7:
		w := types.String{Target: &r.Segmento}

		return w

	case 8:
		w := types.String{Target: &r.Cliente}

		return w

	case 9:
		w := types.String{Target: &r.Contrato}

		return w

	case 10:
		w := types.String{Target: &r.TipoServicio}

		return w

	case 11:
		w := types.String{Target: &r.Destinatario}

		return w

	case 12:
		w := types.String{Target: &r.Domicilio}

		return w

	}
	panic("Unknown field index")
}

func (r *NovedadEventual) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NovedadEventual) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ NovedadEventual) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NovedadEventual) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NovedadEventual) HintSize(int)                     { panic("Unsupported operation") }
func (_ NovedadEventual) Finalize()                        {}

func (_ NovedadEventual) AvroCRC64Fingerprint() []byte {
	return []byte(NovedadEventualAvroCRC64Fingerprint)
}

func (r NovedadEventual) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["tipoPendiente"], err = json.Marshal(r.TipoPendiente)
	if err != nil {
		return nil, err
	}
	output["motivo"], err = json.Marshal(r.Motivo)
	if err != nil {
		return nil, err
	}
	output["email"], err = json.Marshal(r.Email)
	if err != nil {
		return nil, err
	}
	output["telefono"], err = json.Marshal(r.Telefono)
	if err != nil {
		return nil, err
	}
	output["canal"], err = json.Marshal(r.Canal)
	if err != nil {
		return nil, err
	}
	output["envio"], err = json.Marshal(r.Envio)
	if err != nil {
		return nil, err
	}
	output["sucursalActual"], err = json.Marshal(r.SucursalActual)
	if err != nil {
		return nil, err
	}
	output["segmento"], err = json.Marshal(r.Segmento)
	if err != nil {
		return nil, err
	}
	output["cliente"], err = json.Marshal(r.Cliente)
	if err != nil {
		return nil, err
	}
	output["contrato"], err = json.Marshal(r.Contrato)
	if err != nil {
		return nil, err
	}
	output["tipoServicio"], err = json.Marshal(r.TipoServicio)
	if err != nil {
		return nil, err
	}
	output["destinatario"], err = json.Marshal(r.Destinatario)
	if err != nil {
		return nil, err
	}
	output["domicilio"], err = json.Marshal(r.Domicilio)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NovedadEventual) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["tipoPendiente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoPendiente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tipoPendiente")
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
	val = func() json.RawMessage {
		if v, ok := fields["email"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Email); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for email")
	}
	val = func() json.RawMessage {
		if v, ok := fields["telefono"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Telefono); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for telefono")
	}
	val = func() json.RawMessage {
		if v, ok := fields["canal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Canal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for canal")
	}
	val = func() json.RawMessage {
		if v, ok := fields["envio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Envio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for envio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["sucursalActual"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalActual); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for sucursalActual")
	}
	val = func() json.RawMessage {
		if v, ok := fields["segmento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Segmento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for segmento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["cliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["contrato"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contrato); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for contrato")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoServicio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoServicio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tipoServicio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["destinatario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Destinatario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for destinatario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["domicilio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Domicilio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for domicilio")
	}
	return nil
}
