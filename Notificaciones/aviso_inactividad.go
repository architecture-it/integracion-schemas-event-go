// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AvisoInactividad.avsc
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

type AvisoInactividad struct {
	Sistema string `json:"sistema"`

	Cuando int64 `json:"cuando"`

	CodigoDeContratoInterno string `json:"codigoDeContratoInterno"`

	CodigoDeEnvio string `json:"codigoDeEnvio"`

	DiasInactivo int32 `json:"diasInactivo"`

	DestinatarioNombre string `json:"destinatarioNombre"`

	DestinatarioEmail string `json:"destinatarioEmail"`

	DestinatarioTelefono string `json:"destinatarioTelefono"`

	IdMovimiento int32 `json:"idMovimiento"`

	Movimiento string `json:"movimiento"`

	IdSubMovimiento int32 `json:"idSubMovimiento"`

	SubMovimiento string `json:"subMovimiento"`

	Motivo string `json:"motivo"`

	IdTipoContacto int32 `json:"idTipoContacto"`

	TipoContacto string `json:"tipoContacto"`
}

const AvisoInactividadAvroCRC64Fingerprint = "\xaf\xd5m\x90:\xf3\x1a\xeb"

func NewAvisoInactividad() AvisoInactividad {
	r := AvisoInactividad{}
	return r
}

func DeserializeAvisoInactividad(r io.Reader) (AvisoInactividad, error) {
	t := NewAvisoInactividad()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAvisoInactividadFromSchema(r io.Reader, schema string) (AvisoInactividad, error) {
	t := NewAvisoInactividad()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAvisoInactividad(r AvisoInactividad, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Sistema, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Cuando, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoDeContratoInterno, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoDeEnvio, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.DiasInactivo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.DestinatarioNombre, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.DestinatarioEmail, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.DestinatarioTelefono, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.IdMovimiento, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Movimiento, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.IdSubMovimiento, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.SubMovimiento, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Motivo, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.IdTipoContacto, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TipoContacto, w)
	if err != nil {
		return err
	}
	return err
}

func (r AvisoInactividad) Serialize(w io.Writer) error {
	return writeAvisoInactividad(r, w)
}

func (r AvisoInactividad) Schema() string {
	return "{\"fields\":[{\"name\":\"sistema\",\"type\":\"string\"},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"name\":\"diasInactivo\",\"type\":\"int\"},{\"name\":\"destinatarioNombre\",\"type\":\"string\"},{\"name\":\"destinatarioEmail\",\"type\":\"string\"},{\"name\":\"destinatarioTelefono\",\"type\":\"string\"},{\"name\":\"idMovimiento\",\"type\":\"int\"},{\"name\":\"movimiento\",\"type\":\"string\"},{\"name\":\"idSubMovimiento\",\"type\":\"int\"},{\"name\":\"subMovimiento\",\"type\":\"string\"},{\"name\":\"motivo\",\"type\":\"string\"},{\"name\":\"idTipoContacto\",\"type\":\"int\"},{\"name\":\"tipoContacto\",\"type\":\"string\"}],\"name\":\"Andreani.Notificaciones.Events.Records.AvisoInactividad\",\"type\":\"record\"}"
}

func (r AvisoInactividad) SchemaName() string {
	return "Andreani.Notificaciones.Events.Records.AvisoInactividad"
}

func (_ AvisoInactividad) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ AvisoInactividad) SetInt(v int32)       { panic("Unsupported operation") }
func (_ AvisoInactividad) SetLong(v int64)      { panic("Unsupported operation") }
func (_ AvisoInactividad) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ AvisoInactividad) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ AvisoInactividad) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ AvisoInactividad) SetString(v string)   { panic("Unsupported operation") }
func (_ AvisoInactividad) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AvisoInactividad) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Sistema}

		return w

	case 1:
		w := types.Long{Target: &r.Cuando}

		return w

	case 2:
		w := types.String{Target: &r.CodigoDeContratoInterno}

		return w

	case 3:
		w := types.String{Target: &r.CodigoDeEnvio}

		return w

	case 4:
		w := types.Int{Target: &r.DiasInactivo}

		return w

	case 5:
		w := types.String{Target: &r.DestinatarioNombre}

		return w

	case 6:
		w := types.String{Target: &r.DestinatarioEmail}

		return w

	case 7:
		w := types.String{Target: &r.DestinatarioTelefono}

		return w

	case 8:
		w := types.Int{Target: &r.IdMovimiento}

		return w

	case 9:
		w := types.String{Target: &r.Movimiento}

		return w

	case 10:
		w := types.Int{Target: &r.IdSubMovimiento}

		return w

	case 11:
		w := types.String{Target: &r.SubMovimiento}

		return w

	case 12:
		w := types.String{Target: &r.Motivo}

		return w

	case 13:
		w := types.Int{Target: &r.IdTipoContacto}

		return w

	case 14:
		w := types.String{Target: &r.TipoContacto}

		return w

	}
	panic("Unknown field index")
}

func (r *AvisoInactividad) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *AvisoInactividad) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ AvisoInactividad) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ AvisoInactividad) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ AvisoInactividad) HintSize(int)                     { panic("Unsupported operation") }
func (_ AvisoInactividad) Finalize()                        {}

func (_ AvisoInactividad) AvroCRC64Fingerprint() []byte {
	return []byte(AvisoInactividadAvroCRC64Fingerprint)
}

func (r AvisoInactividad) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["sistema"], err = json.Marshal(r.Sistema)
	if err != nil {
		return nil, err
	}
	output["cuando"], err = json.Marshal(r.Cuando)
	if err != nil {
		return nil, err
	}
	output["codigoDeContratoInterno"], err = json.Marshal(r.CodigoDeContratoInterno)
	if err != nil {
		return nil, err
	}
	output["codigoDeEnvio"], err = json.Marshal(r.CodigoDeEnvio)
	if err != nil {
		return nil, err
	}
	output["diasInactivo"], err = json.Marshal(r.DiasInactivo)
	if err != nil {
		return nil, err
	}
	output["destinatarioNombre"], err = json.Marshal(r.DestinatarioNombre)
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
	output["idMovimiento"], err = json.Marshal(r.IdMovimiento)
	if err != nil {
		return nil, err
	}
	output["movimiento"], err = json.Marshal(r.Movimiento)
	if err != nil {
		return nil, err
	}
	output["idSubMovimiento"], err = json.Marshal(r.IdSubMovimiento)
	if err != nil {
		return nil, err
	}
	output["subMovimiento"], err = json.Marshal(r.SubMovimiento)
	if err != nil {
		return nil, err
	}
	output["motivo"], err = json.Marshal(r.Motivo)
	if err != nil {
		return nil, err
	}
	output["idTipoContacto"], err = json.Marshal(r.IdTipoContacto)
	if err != nil {
		return nil, err
	}
	output["tipoContacto"], err = json.Marshal(r.TipoContacto)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *AvisoInactividad) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["sistema"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sistema); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for sistema")
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
		return fmt.Errorf("no value specified for cuando")
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
		return fmt.Errorf("no value specified for codigoDeContratoInterno")
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
		return fmt.Errorf("no value specified for codigoDeEnvio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["diasInactivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DiasInactivo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for diasInactivo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["destinatarioNombre"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioNombre); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for destinatarioNombre")
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
		return fmt.Errorf("no value specified for destinatarioEmail")
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
		return fmt.Errorf("no value specified for destinatarioTelefono")
	}
	val = func() json.RawMessage {
		if v, ok := fields["idMovimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdMovimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idMovimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["movimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Movimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for movimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["idSubMovimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdSubMovimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idSubMovimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["subMovimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SubMovimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for subMovimiento")
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
		if v, ok := fields["idTipoContacto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdTipoContacto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idTipoContacto")
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
		return fmt.Errorf("no value specified for tipoContacto")
	}
	return nil
}
