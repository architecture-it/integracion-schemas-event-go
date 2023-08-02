// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnviosEvent.avsc
 */
package PersonaBackendEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EnviosEvent struct {
	PedidoId string `json:"PedidoId"`

	UsuarioId string `json:"UsuarioId"`

	UsuarioLoginId string `json:"UsuarioLoginId"`

	Estado string `json:"Estado"`

	NumeroDeOperacion *UnionNullString `json:"NumeroDeOperacion"`

	Envios *UnionNullArrayEnvio `json:"Envios"`
}

const EnviosEventAvroCRC64Fingerprint = "R\x00E\xa4\xaf\xc6K@"

func NewEnviosEvent() EnviosEvent {
	r := EnviosEvent{}
	r.NumeroDeOperacion = nil
	r.Envios = nil
	return r
}

func DeserializeEnviosEvent(r io.Reader) (EnviosEvent, error) {
	t := NewEnviosEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnviosEventFromSchema(r io.Reader, schema string) (EnviosEvent, error) {
	t := NewEnviosEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnviosEvent(r EnviosEvent, w io.Writer) error {
	var err error
	err = vm.WriteString(r.PedidoId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.UsuarioId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.UsuarioLoginId, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Estado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeOperacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayEnvio(r.Envios, w)
	if err != nil {
		return err
	}
	return err
}

func (r EnviosEvent) Serialize(w io.Writer) error {
	return writeEnviosEvent(r, w)
}

func (r EnviosEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"PedidoId\",\"type\":\"string\"},{\"name\":\"UsuarioId\",\"type\":\"string\"},{\"name\":\"UsuarioLoginId\",\"type\":\"string\"},{\"name\":\"Estado\",\"type\":\"string\"},{\"default\":null,\"name\":\"NumeroDeOperacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Envios\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"TipoDeEnvio\",\"type\":\"string\"},{\"name\":\"ModoDeEntrega\",\"type\":\"string\"},{\"name\":\"NumeroInterno\",\"type\":\"string\"},{\"name\":\"NumeroDeSeguimiento\",\"type\":\"string\"},{\"name\":\"NumeroDeContrato\",\"type\":\"string\"},{\"name\":\"Estado\",\"type\":\"string\"},{\"name\":\"RemitoFactura\",\"type\":\"string\"},{\"name\":\"FechaCreacion\",\"type\":\"long\"},{\"default\":null,\"name\":\"Origen\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Tipo\",\"type\":\"string\"},{\"name\":\"Calle\",\"type\":\"string\"},{\"default\":null,\"name\":\"Numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Unidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"},{\"name\":\"Provincia\",\"type\":\"string\"},{\"default\":null,\"name\":\"SucursalId\",\"type\":[\"null\",\"string\"]}],\"name\":\"Origen\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Destino\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Tipo\",\"type\":\"string\"},{\"name\":\"Calle\",\"type\":\"string\"},{\"default\":null,\"name\":\"Numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Unidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"},{\"name\":\"Provincia\",\"type\":\"string\"},{\"default\":null,\"name\":\"ObservacionesAdicionales\",\"type\":[\"null\",\"string\"]}],\"name\":\"Destino\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Remitente\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Nombre\",\"type\":\"string\"},{\"default\":null,\"name\":\"Apellido\",\"type\":[\"null\",\"string\"]},{\"name\":\"Email\",\"type\":\"string\"},{\"name\":\"Telefono\",\"type\":\"string\"},{\"name\":\"Dni\",\"type\":\"string\"},{\"name\":\"Calle\",\"type\":\"string\"},{\"default\":null,\"name\":\"Numero\",\"type\":[\"null\",\"string\"]},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"},{\"default\":null,\"name\":\"Piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Unidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"Pais\",\"type\":\"string\"},{\"name\":\"Region\",\"type\":\"string\"}],\"name\":\"Remitente\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Destinatario\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Apellido\",\"type\":\"string\"},{\"name\":\"Email\",\"type\":\"string\"},{\"name\":\"Telefono\",\"type\":\"string\"},{\"name\":\"Dni\",\"type\":\"string\"}],\"name\":\"Destinatario\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Paquetes\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"TipoId\",\"type\":\"string\"},{\"name\":\"Tipo\",\"type\":\"string\"},{\"name\":\"Alto\",\"type\":\"string\"},{\"name\":\"Ancho\",\"type\":\"string\"},{\"name\":\"Largo\",\"type\":\"string\"},{\"name\":\"Peso\",\"type\":\"string\"},{\"name\":\"ValorDeclarado\",\"type\":\"int\"},{\"default\":null,\"name\":\"NumeroDeBulto\",\"type\":[\"null\",\"string\"]}],\"name\":\"Paquete\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"DatosConstancia\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Url\",\"type\":\"string\"},{\"name\":\"NumeroPermisionaria\",\"type\":\"string\"},{\"name\":\"SucursalDistribucion\",\"type\":\"string\"},{\"name\":\"SucursalDistribucionNomenclatura\",\"type\":\"string\"},{\"name\":\"SucursalDistribucionDescripcion\",\"type\":\"string\"},{\"name\":\"SucursalDistribucionId\",\"type\":\"string\"},{\"name\":\"SucursalRendicion\",\"type\":\"string\"},{\"name\":\"SucursalRendicionNomenclatura\",\"type\":\"string\"},{\"name\":\"SucursalRendicionDescripcion\",\"type\":\"string\"},{\"name\":\"SucursalRendicionId\",\"type\":\"string\"},{\"name\":\"CodigoSucursalCabecera\",\"type\":\"string\"},{\"name\":\"SucursalAbastecedoraDescripcion\",\"type\":\"string\"},{\"name\":\"SucursalAbastecedoraId\",\"type\":\"string\"},{\"name\":\"CodigoZonaReparto\",\"type\":\"string\"}],\"name\":\"DatosConstancia\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Cliente\",\"type\":[\"null\",{\"fields\":[{\"name\":\"CodigoAndreani\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Logo\",\"type\":\"string\"}],\"name\":\"Cliente\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Cotizacion\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"IdMercadoPago\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Iva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PesoAforado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SeguroDistribucionSinIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SeguroDistribucionConIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DistribucionSinIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DistribucionConIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TarifaSinIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TarifaConIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Descuento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Cupon\",\"type\":[\"null\",\"string\"]}],\"name\":\"Cotizacion\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Canal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Referencia\",\"type\":[\"null\",\"string\"]}],\"name\":\"Envio\",\"namespace\":\"Andreani.PersonaBackend.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"Andreani.PersonaBackend.Events.Record.EnviosEvent\",\"type\":\"record\"}"
}

func (r EnviosEvent) SchemaName() string {
	return "Andreani.PersonaBackend.Events.Record.EnviosEvent"
}

func (_ EnviosEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EnviosEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EnviosEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EnviosEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EnviosEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EnviosEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EnviosEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ EnviosEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EnviosEvent) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.PedidoId}

		return w

	case 1:
		w := types.String{Target: &r.UsuarioId}

		return w

	case 2:
		w := types.String{Target: &r.UsuarioLoginId}

		return w

	case 3:
		w := types.String{Target: &r.Estado}

		return w

	case 4:
		r.NumeroDeOperacion = NewUnionNullString()

		return r.NumeroDeOperacion
	case 5:
		r.Envios = NewUnionNullArrayEnvio()

		return r.Envios
	}
	panic("Unknown field index")
}

func (r *EnviosEvent) SetDefault(i int) {
	switch i {
	case 4:
		r.NumeroDeOperacion = nil
		return
	case 5:
		r.Envios = nil
		return
	}
	panic("Unknown field index")
}

func (r *EnviosEvent) NullField(i int) {
	switch i {
	case 4:
		r.NumeroDeOperacion = nil
		return
	case 5:
		r.Envios = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ EnviosEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EnviosEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EnviosEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ EnviosEvent) Finalize()                        {}

func (_ EnviosEvent) AvroCRC64Fingerprint() []byte {
	return []byte(EnviosEventAvroCRC64Fingerprint)
}

func (r EnviosEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["PedidoId"], err = json.Marshal(r.PedidoId)
	if err != nil {
		return nil, err
	}
	output["UsuarioId"], err = json.Marshal(r.UsuarioId)
	if err != nil {
		return nil, err
	}
	output["UsuarioLoginId"], err = json.Marshal(r.UsuarioLoginId)
	if err != nil {
		return nil, err
	}
	output["Estado"], err = json.Marshal(r.Estado)
	if err != nil {
		return nil, err
	}
	output["NumeroDeOperacion"], err = json.Marshal(r.NumeroDeOperacion)
	if err != nil {
		return nil, err
	}
	output["Envios"], err = json.Marshal(r.Envios)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EnviosEvent) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["PedidoId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PedidoId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PedidoId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["UsuarioId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UsuarioId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for UsuarioId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["UsuarioLoginId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UsuarioLoginId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for UsuarioLoginId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Estado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Estado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeOperacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeOperacion); err != nil {
			return err
		}
	} else {
		r.NumeroDeOperacion = NewUnionNullString()

		r.NumeroDeOperacion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Envios"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Envios); err != nil {
			return err
		}
	} else {
		r.Envios = NewUnionNullArrayEnvio()

		r.Envios = nil
	}
	return nil
}
