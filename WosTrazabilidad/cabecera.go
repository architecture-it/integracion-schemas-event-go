// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     WosTrazaAnmat.avsc
 */
package WosTrazabilidadEvents

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
	Serialkey int32 `json:"Serialkey"`

	Almacen string `json:"almacen"`

	Instancia string `json:"instancia"`

	Propietario string `json:"propietario"`

	TipoDocumento int32 `json:"tipoDocumento"`

	NroDocumento string `json:"nroDocumento"`

	NroDocumentoWMS string `json:"nroDocumentoWMS"`

	GlnOrigen string `json:"glnOrigen"`

	GlnDestino string `json:"glnDestino"`

	Estado int32 `json:"estado"`

	DescripcionEstado string `json:"descripcionEstado"`

	Remito string `json:"remito"`

	IdEventoAnmat string `json:"idEventoAnmat"`

	Destinatario *UnionNullDestinatario `json:"destinatario"`
}

const CabeceraAvroCRC64Fingerprint = "\xb5\xe3\x9bR\xf7>fO"

func NewCabecera() Cabecera {
	r := Cabecera{}
	r.Destinatario = nil
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
	err = vm.WriteInt(r.Serialkey, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Almacen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Instancia, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.TipoDocumento, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NroDocumento, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NroDocumentoWMS, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.GlnOrigen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.GlnDestino, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Estado, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.DescripcionEstado, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Remito, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.IdEventoAnmat, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDestinatario(r.Destinatario, w)
	if err != nil {
		return err
	}
	return err
}

func (r Cabecera) Serialize(w io.Writer) error {
	return writeCabecera(r, w)
}

func (r Cabecera) Schema() string {
	return "{\"fields\":[{\"name\":\"Serialkey\",\"type\":\"int\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"name\":\"instancia\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"tipoDocumento\",\"type\":\"int\"},{\"name\":\"nroDocumento\",\"type\":\"string\"},{\"name\":\"nroDocumentoWMS\",\"type\":\"string\"},{\"name\":\"glnOrigen\",\"type\":\"string\"},{\"name\":\"glnDestino\",\"type\":\"string\"},{\"name\":\"estado\",\"type\":\"int\"},{\"name\":\"descripcionEstado\",\"type\":\"string\"},{\"name\":\"remito\",\"type\":\"string\"},{\"name\":\"idEventoAnmat\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Apellido\",\"type\":\"string\"},{\"name\":\"TipoDocumento\",\"type\":\"int\"}],\"name\":\"Destinatario\",\"type\":\"record\"}]}],\"name\":\"Andreani.WosTrazabilidad.Events.AnmatCommon.Cabecera\",\"type\":\"record\"}"
}

func (r Cabecera) SchemaName() string {
	return "Andreani.WosTrazabilidad.Events.AnmatCommon.Cabecera"
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
		w := types.Int{Target: &r.Serialkey}

		return w

	case 1:
		w := types.String{Target: &r.Almacen}

		return w

	case 2:
		w := types.String{Target: &r.Instancia}

		return w

	case 3:
		w := types.String{Target: &r.Propietario}

		return w

	case 4:
		w := types.Int{Target: &r.TipoDocumento}

		return w

	case 5:
		w := types.String{Target: &r.NroDocumento}

		return w

	case 6:
		w := types.String{Target: &r.NroDocumentoWMS}

		return w

	case 7:
		w := types.String{Target: &r.GlnOrigen}

		return w

	case 8:
		w := types.String{Target: &r.GlnDestino}

		return w

	case 9:
		w := types.Int{Target: &r.Estado}

		return w

	case 10:
		w := types.String{Target: &r.DescripcionEstado}

		return w

	case 11:
		w := types.String{Target: &r.Remito}

		return w

	case 12:
		w := types.String{Target: &r.IdEventoAnmat}

		return w

	case 13:
		r.Destinatario = NewUnionNullDestinatario()

		return r.Destinatario
	}
	panic("Unknown field index")
}

func (r *Cabecera) SetDefault(i int) {
	switch i {
	case 13:
		r.Destinatario = nil
		return
	}
	panic("Unknown field index")
}

func (r *Cabecera) NullField(i int) {
	switch i {
	case 13:
		r.Destinatario = nil
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
	output["Serialkey"], err = json.Marshal(r.Serialkey)
	if err != nil {
		return nil, err
	}
	output["almacen"], err = json.Marshal(r.Almacen)
	if err != nil {
		return nil, err
	}
	output["instancia"], err = json.Marshal(r.Instancia)
	if err != nil {
		return nil, err
	}
	output["propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["tipoDocumento"], err = json.Marshal(r.TipoDocumento)
	if err != nil {
		return nil, err
	}
	output["nroDocumento"], err = json.Marshal(r.NroDocumento)
	if err != nil {
		return nil, err
	}
	output["nroDocumentoWMS"], err = json.Marshal(r.NroDocumentoWMS)
	if err != nil {
		return nil, err
	}
	output["glnOrigen"], err = json.Marshal(r.GlnOrigen)
	if err != nil {
		return nil, err
	}
	output["glnDestino"], err = json.Marshal(r.GlnDestino)
	if err != nil {
		return nil, err
	}
	output["estado"], err = json.Marshal(r.Estado)
	if err != nil {
		return nil, err
	}
	output["descripcionEstado"], err = json.Marshal(r.DescripcionEstado)
	if err != nil {
		return nil, err
	}
	output["remito"], err = json.Marshal(r.Remito)
	if err != nil {
		return nil, err
	}
	output["idEventoAnmat"], err = json.Marshal(r.IdEventoAnmat)
	if err != nil {
		return nil, err
	}
	output["destinatario"], err = json.Marshal(r.Destinatario)
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
		if v, ok := fields["Serialkey"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Serialkey); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Serialkey")
	}
	val = func() json.RawMessage {
		if v, ok := fields["almacen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Almacen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for almacen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["instancia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Instancia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for instancia")
	}
	val = func() json.RawMessage {
		if v, ok := fields["propietario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Propietario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for propietario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoDocumento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDocumento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tipoDocumento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["nroDocumento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NroDocumento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for nroDocumento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["nroDocumentoWMS"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NroDocumentoWMS); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for nroDocumentoWMS")
	}
	val = func() json.RawMessage {
		if v, ok := fields["glnOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.GlnOrigen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for glnOrigen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["glnDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.GlnDestino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for glnDestino")
	}
	val = func() json.RawMessage {
		if v, ok := fields["estado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for estado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["descripcionEstado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DescripcionEstado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for descripcionEstado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["remito"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Remito); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for remito")
	}
	val = func() json.RawMessage {
		if v, ok := fields["idEventoAnmat"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdEventoAnmat); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idEventoAnmat")
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
		r.Destinatario = NewUnionNullDestinatario()

		r.Destinatario = nil
	}
	return nil
}
