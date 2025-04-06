// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EmailEtiquetaEvent.avsc
 */
package CorporativoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EmailEtiquetaEvent struct {
	Envios []EnvioEtiqueta `json:"Envios"`

	FechaSolicitud *UnionNullString `json:"FechaSolicitud"`

	Empresa *UnionNullEmpresaEtiqueta `json:"Empresa"`
}

const EmailEtiquetaEventAvroCRC64Fingerprint = "_p@\x10\xae\x85x\x04"

func NewEmailEtiquetaEvent() EmailEtiquetaEvent {
	r := EmailEtiquetaEvent{}
	r.Envios = make([]EnvioEtiqueta, 0)

	r.FechaSolicitud = nil
	r.Empresa = nil
	return r
}

func DeserializeEmailEtiquetaEvent(r io.Reader) (EmailEtiquetaEvent, error) {
	t := NewEmailEtiquetaEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEmailEtiquetaEventFromSchema(r io.Reader, schema string) (EmailEtiquetaEvent, error) {
	t := NewEmailEtiquetaEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEmailEtiquetaEvent(r EmailEtiquetaEvent, w io.Writer) error {
	var err error
	err = writeArrayEnvioEtiqueta(r.Envios, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaSolicitud, w)
	if err != nil {
		return err
	}
	err = writeUnionNullEmpresaEtiqueta(r.Empresa, w)
	if err != nil {
		return err
	}
	return err
}

func (r EmailEtiquetaEvent) Serialize(w io.Writer) error {
	return writeEmailEtiquetaEvent(r, w)
}

func (r EmailEtiquetaEvent) Schema() string {
	return "{\"fields\":[{\"name\":\"Envios\",\"type\":{\"items\":{\"fields\":[{\"name\":\"EnvioId\",\"type\":\"string\"},{\"name\":\"UsuarioLoginId\",\"type\":\"string\"},{\"name\":\"NumeroSeguimiento\",\"type\":\"string\"},{\"name\":\"NumeroInterno\",\"type\":\"string\"},{\"name\":\"NumeroContrato\",\"type\":\"string\"},{\"name\":\"Tipo\",\"type\":\"string\"},{\"default\":null,\"name\":\"RemitoFactura\",\"type\":[\"null\",\"string\"]},{\"name\":\"Canal\",\"type\":\"string\"},{\"name\":\"PesoTotal\",\"type\":\"string\"},{\"name\":\"Template\",\"type\":\"string\"},{\"name\":\"TipoDeEnvio\",\"type\":{\"fields\":[{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"RemitoFactura\",\"type\":\"string\"}],\"name\":\"TipoDeEnvioEtiqueta\",\"type\":\"record\"}},{\"name\":\"ModoDeEntrega\",\"type\":{\"fields\":[{\"name\":\"Nombre\",\"type\":\"string\"}],\"name\":\"ModoDeEntregaEtiqueta\",\"type\":\"record\"}},{\"name\":\"Pedido\",\"type\":{\"fields\":[{\"name\":\"Usuario\",\"type\":{\"fields\":[{\"name\":\"Cliente\",\"type\":{\"fields\":[{\"name\":\"CodigoAndreani\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Logo\",\"type\":\"string\"}],\"name\":\"ClienteEtiqueta\",\"type\":\"record\"}}],\"name\":\"UsuarioEtiqueta\",\"type\":\"record\"}}],\"name\":\"PedidoEtiqueta\",\"type\":\"record\"}},{\"name\":\"Destino\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"DireccionDestino\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Calle\",\"type\":\"string\"},{\"name\":\"Numero\",\"type\":\"string\"},{\"default\":null,\"name\":\"Piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Unidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"},{\"name\":\"Provincia\",\"type\":\"string\"},{\"default\":null,\"name\":\"DatosAdicionales\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Latitud\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Longitud\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ObservacionesAdicionales\",\"type\":[\"null\",\"string\"]}],\"name\":\"DireccionDestinoEtiqueta\",\"type\":\"record\"}]}],\"name\":\"DestinoEtiqueta\",\"type\":\"record\"}},{\"name\":\"Destinatario\",\"type\":{\"fields\":[{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Apellido\",\"type\":\"string\"},{\"name\":\"Telefono\",\"type\":\"string\"},{\"name\":\"DNI\",\"type\":\"string\"},{\"name\":\"Email\",\"type\":\"string\"},{\"name\":\"DestinatarioEmail\",\"type\":\"string\"}],\"name\":\"DestinatarioEtiqueta\",\"type\":\"record\"}},{\"name\":\"Origen\",\"type\":{\"fields\":[{\"name\":\"Tipo\",\"type\":\"string\"},{\"name\":\"Calle\",\"type\":\"string\"},{\"name\":\"Numero\",\"type\":\"string\"},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"Provincia\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"},{\"name\":\"SucursalId\",\"type\":\"string\"}],\"name\":\"OrigenEtiqueta\",\"type\":\"record\"}},{\"name\":\"Remitente\",\"type\":{\"fields\":[{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Apellido\",\"type\":\"string\"},{\"default\":null,\"name\":\"Telefono\",\"type\":[\"null\",\"string\"]},{\"name\":\"DNI\",\"type\":\"string\"},{\"name\":\"Email\",\"type\":\"string\"},{\"name\":\"Direccion\",\"type\":{\"fields\":[{\"name\":\"Calle\",\"type\":\"string\"},{\"name\":\"Numero\",\"type\":\"string\"},{\"default\":null,\"name\":\"Piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Unidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"},{\"name\":\"Provincia\",\"type\":\"string\"},{\"default\":null,\"name\":\"DatosAdicionales\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Latitud\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Longitud\",\"type\":[\"null\",\"string\"]}],\"name\":\"DireccionRemitenteEtiqueta\",\"type\":\"record\"}}],\"name\":\"RemitenteEtiqueta\",\"type\":\"record\"}},{\"name\":\"DatosConstancia\",\"type\":{\"fields\":[{\"name\":\"Url\",\"type\":\"string\"},{\"name\":\"NumeroPermisionaria\",\"type\":\"string\"},{\"name\":\"SucursalDistribucion\",\"type\":\"string\"},{\"name\":\"SucursalDistribucionNomenclatura\",\"type\":\"string\"},{\"name\":\"SucursalDistribucionDescripcion\",\"type\":\"string\"},{\"name\":\"SucursalDistribucionId\",\"type\":\"string\"},{\"default\":null,\"name\":\"SucursalRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SucursalRendicionNomenclatura\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SucursalRendicionDescripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SucursalRendicionId\",\"type\":[\"null\",\"string\"]},{\"name\":\"CodigoSucursalCabecera\",\"type\":\"string\"},{\"default\":null,\"name\":\"SucursalAbastecedoraDescripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SucursalAbastecedoraId\",\"type\":[\"null\",\"string\"]},{\"name\":\"CodigoZonaReparto\",\"type\":\"string\"}],\"name\":\"DatosConstanciaEtiqueta\",\"type\":\"record\"}},{\"name\":\"Paquetes\",\"type\":{\"items\":{\"fields\":[{\"default\":null,\"name\":\"Item\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Nombre\",\"type\":\"string\"}],\"name\":\"ItemEtiqueta\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Alto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Ancho\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Largo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Peso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ValorDeclarado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeBulto\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"TotalDeBultos\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"NumeroDePaquete\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Volumen\",\"type\":[\"null\",\"string\"]}],\"name\":\"PaqueteEtiqueta\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"EnvioEtiqueta\",\"type\":\"record\"},\"type\":\"array\"}},{\"default\":null,\"name\":\"FechaSolicitud\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Empresa\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Nombre\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"RazonSocial\",\"type\":[\"null\",\"string\"]}],\"name\":\"EmpresaEtiqueta\",\"type\":\"record\"}]}],\"name\":\"Andreani.Corporativo.Events.Record.EmailEtiquetaEvent\",\"type\":\"record\"}"
}

func (r EmailEtiquetaEvent) SchemaName() string {
	return "Andreani.Corporativo.Events.Record.EmailEtiquetaEvent"
}

func (_ EmailEtiquetaEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EmailEtiquetaEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EmailEtiquetaEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EmailEtiquetaEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EmailEtiquetaEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EmailEtiquetaEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EmailEtiquetaEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ EmailEtiquetaEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EmailEtiquetaEvent) Get(i int) types.Field {
	switch i {
	case 0:
		r.Envios = make([]EnvioEtiqueta, 0)

		w := ArrayEnvioEtiquetaWrapper{Target: &r.Envios}

		return w

	case 1:
		r.FechaSolicitud = NewUnionNullString()

		return r.FechaSolicitud
	case 2:
		r.Empresa = NewUnionNullEmpresaEtiqueta()

		return r.Empresa
	}
	panic("Unknown field index")
}

func (r *EmailEtiquetaEvent) SetDefault(i int) {
	switch i {
	case 1:
		r.FechaSolicitud = nil
		return
	case 2:
		r.Empresa = nil
		return
	}
	panic("Unknown field index")
}

func (r *EmailEtiquetaEvent) NullField(i int) {
	switch i {
	case 1:
		r.FechaSolicitud = nil
		return
	case 2:
		r.Empresa = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ EmailEtiquetaEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EmailEtiquetaEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EmailEtiquetaEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ EmailEtiquetaEvent) Finalize()                        {}

func (_ EmailEtiquetaEvent) AvroCRC64Fingerprint() []byte {
	return []byte(EmailEtiquetaEventAvroCRC64Fingerprint)
}

func (r EmailEtiquetaEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Envios"], err = json.Marshal(r.Envios)
	if err != nil {
		return nil, err
	}
	output["FechaSolicitud"], err = json.Marshal(r.FechaSolicitud)
	if err != nil {
		return nil, err
	}
	output["Empresa"], err = json.Marshal(r.Empresa)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EmailEtiquetaEvent) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
		return fmt.Errorf("no value specified for Envios")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaSolicitud"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaSolicitud); err != nil {
			return err
		}
	} else {
		r.FechaSolicitud = NewUnionNullString()

		r.FechaSolicitud = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Empresa"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Empresa); err != nil {
			return err
		}
	} else {
		r.Empresa = NewUnionNullEmpresaEtiqueta()

		r.Empresa = nil
	}
	return nil
}
