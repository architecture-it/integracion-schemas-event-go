// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NovedadMasiva.avsc
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
	IdModelo int64 `json:"idModelo"`

	DestinatarioEmail *UnionNullString `json:"destinatarioEmail"`

	DestinatarioTelefono *UnionNullString `json:"destinatarioTelefono"`

	Canal *UnionNullString `json:"canal"`

	CodigoDeEnvio *UnionNullString `json:"codigoDeEnvio"`

	SucursalActual *UnionNullString `json:"sucursalActual"`

	Segmento *UnionNullString `json:"segmento"`

	Cliente *UnionNullString `json:"cliente"`

	CodigoDeContratoInterno *UnionNullString `json:"codigoDeContratoInterno"`

	TipoContacto *UnionNullString `json:"tipoContacto"`

	CodigoPostal *UnionNullString `json:"codigoPostal"`

	Provincia *UnionNullString `json:"provincia"`

	Cuando *UnionNullLong `json:"cuando"`

	AdmisionDrop *UnionNullBool `json:"AdmisionDrop"`
}

const NovedadEventualAvroCRC64Fingerprint = ":\xd1\xdd1k\xdd\xd4N"

func NewNovedadEventual() NovedadEventual {
	r := NovedadEventual{}
	r.DestinatarioEmail = nil
	r.DestinatarioTelefono = nil
	r.Canal = nil
	r.CodigoDeEnvio = nil
	r.SucursalActual = nil
	r.Segmento = nil
	r.Cliente = nil
	r.CodigoDeContratoInterno = nil
	r.TipoContacto = nil
	r.CodigoPostal = nil
	r.Provincia = nil
	r.Cuando = nil
	r.AdmisionDrop = nil
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
	err = vm.WriteLong(r.IdModelo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioEmail, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioTelefono, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Canal, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoDeEnvio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.SucursalActual, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Segmento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Cliente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoDeContratoInterno, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoContacto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoPostal, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Provincia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.Cuando, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.AdmisionDrop, w)
	if err != nil {
		return err
	}
	return err
}

func (r NovedadEventual) Serialize(w io.Writer) error {
	return writeNovedadEventual(r, w)
}

func (r NovedadEventual) Schema() string {
	return "{\"fields\":[{\"name\":\"idModelo\",\"type\":\"long\"},{\"default\":null,\"name\":\"destinatarioEmail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"destinatarioTelefono\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"canal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoDeEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalActual\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"segmento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoDeContratoInterno\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoContacto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"provincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cuando\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"AdmisionDrop\",\"type\":[\"null\",\"boolean\"]}],\"name\":\"Andreani.Notificaciones.Events.Records.NovedadEventual\",\"type\":\"record\"}"
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
		w := types.Long{Target: &r.IdModelo}

		return w

	case 1:
		r.DestinatarioEmail = NewUnionNullString()

		return r.DestinatarioEmail
	case 2:
		r.DestinatarioTelefono = NewUnionNullString()

		return r.DestinatarioTelefono
	case 3:
		r.Canal = NewUnionNullString()

		return r.Canal
	case 4:
		r.CodigoDeEnvio = NewUnionNullString()

		return r.CodigoDeEnvio
	case 5:
		r.SucursalActual = NewUnionNullString()

		return r.SucursalActual
	case 6:
		r.Segmento = NewUnionNullString()

		return r.Segmento
	case 7:
		r.Cliente = NewUnionNullString()

		return r.Cliente
	case 8:
		r.CodigoDeContratoInterno = NewUnionNullString()

		return r.CodigoDeContratoInterno
	case 9:
		r.TipoContacto = NewUnionNullString()

		return r.TipoContacto
	case 10:
		r.CodigoPostal = NewUnionNullString()

		return r.CodigoPostal
	case 11:
		r.Provincia = NewUnionNullString()

		return r.Provincia
	case 12:
		r.Cuando = NewUnionNullLong()

		return r.Cuando
	case 13:
		r.AdmisionDrop = NewUnionNullBool()

		return r.AdmisionDrop
	}
	panic("Unknown field index")
}

func (r *NovedadEventual) SetDefault(i int) {
	switch i {
	case 1:
		r.DestinatarioEmail = nil
		return
	case 2:
		r.DestinatarioTelefono = nil
		return
	case 3:
		r.Canal = nil
		return
	case 4:
		r.CodigoDeEnvio = nil
		return
	case 5:
		r.SucursalActual = nil
		return
	case 6:
		r.Segmento = nil
		return
	case 7:
		r.Cliente = nil
		return
	case 8:
		r.CodigoDeContratoInterno = nil
		return
	case 9:
		r.TipoContacto = nil
		return
	case 10:
		r.CodigoPostal = nil
		return
	case 11:
		r.Provincia = nil
		return
	case 12:
		r.Cuando = nil
		return
	case 13:
		r.AdmisionDrop = nil
		return
	}
	panic("Unknown field index")
}

func (r *NovedadEventual) NullField(i int) {
	switch i {
	case 1:
		r.DestinatarioEmail = nil
		return
	case 2:
		r.DestinatarioTelefono = nil
		return
	case 3:
		r.Canal = nil
		return
	case 4:
		r.CodigoDeEnvio = nil
		return
	case 5:
		r.SucursalActual = nil
		return
	case 6:
		r.Segmento = nil
		return
	case 7:
		r.Cliente = nil
		return
	case 8:
		r.CodigoDeContratoInterno = nil
		return
	case 9:
		r.TipoContacto = nil
		return
	case 10:
		r.CodigoPostal = nil
		return
	case 11:
		r.Provincia = nil
		return
	case 12:
		r.Cuando = nil
		return
	case 13:
		r.AdmisionDrop = nil
		return
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
	output["idModelo"], err = json.Marshal(r.IdModelo)
	if err != nil {
		return nil, err
	}
	output["destinatarioEmail"], err = json.Marshal(r.DestinatarioEmail)
	if err != nil {
		return nil, err
	}
	output["destinatarioTelefono"], err = json.Marshal(r.DestinatarioTelefono)
	if err != nil {
		return nil, err
	}
	output["canal"], err = json.Marshal(r.Canal)
	if err != nil {
		return nil, err
	}
	output["codigoDeEnvio"], err = json.Marshal(r.CodigoDeEnvio)
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
	output["codigoDeContratoInterno"], err = json.Marshal(r.CodigoDeContratoInterno)
	if err != nil {
		return nil, err
	}
	output["tipoContacto"], err = json.Marshal(r.TipoContacto)
	if err != nil {
		return nil, err
	}
	output["codigoPostal"], err = json.Marshal(r.CodigoPostal)
	if err != nil {
		return nil, err
	}
	output["provincia"], err = json.Marshal(r.Provincia)
	if err != nil {
		return nil, err
	}
	output["cuando"], err = json.Marshal(r.Cuando)
	if err != nil {
		return nil, err
	}
	output["AdmisionDrop"], err = json.Marshal(r.AdmisionDrop)
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
		if v, ok := fields["idModelo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdModelo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idModelo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["destinatarioEmail"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioEmail); err != nil {
			return err
		}
	} else {
		r.DestinatarioEmail = NewUnionNullString()

		r.DestinatarioEmail = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["destinatarioTelefono"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioTelefono); err != nil {
			return err
		}
	} else {
		r.DestinatarioTelefono = NewUnionNullString()

		r.DestinatarioTelefono = nil
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
		r.Canal = NewUnionNullString()

		r.Canal = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeEnvio); err != nil {
			return err
		}
	} else {
		r.CodigoDeEnvio = NewUnionNullString()

		r.CodigoDeEnvio = nil
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
		r.SucursalActual = NewUnionNullString()

		r.SucursalActual = nil
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
		r.Segmento = NewUnionNullString()

		r.Segmento = nil
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
		r.Cliente = NewUnionNullString()

		r.Cliente = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoDeContratoInterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeContratoInterno); err != nil {
			return err
		}
	} else {
		r.CodigoDeContratoInterno = NewUnionNullString()

		r.CodigoDeContratoInterno = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoContacto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoContacto); err != nil {
			return err
		}
	} else {
		r.TipoContacto = NewUnionNullString()

		r.TipoContacto = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoPostal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoPostal); err != nil {
			return err
		}
	} else {
		r.CodigoPostal = NewUnionNullString()

		r.CodigoPostal = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["provincia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Provincia); err != nil {
			return err
		}
	} else {
		r.Provincia = NewUnionNullString()

		r.Provincia = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["cuando"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cuando); err != nil {
			return err
		}
	} else {
		r.Cuando = NewUnionNullLong()

		r.Cuando = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["AdmisionDrop"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AdmisionDrop); err != nil {
			return err
		}
	} else {
		r.AdmisionDrop = NewUnionNullBool()

		r.AdmisionDrop = nil
	}
	return nil
}
