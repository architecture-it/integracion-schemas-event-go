// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TransactionEvent.avsc
 */
package AltaOrdenEnvioEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Request struct {
	Bultos *UnionNullArrayBultoRequest `json:"Bultos"`

	Contrato string `json:"Contrato"`

	TipoServicio *UnionNullString `json:"TipoServicio"`

	SucursalClienteID *UnionNullInt `json:"SucursalClienteID"`

	Origen *UnionNullLugarDeRetiroOEntrega `json:"origen"`

	Destino *UnionNullLugarDeRetiroOEntrega `json:"destino"`

	IdPedido *UnionNullString `json:"idPedido"`

	Remitente *UnionNullPersona `json:"remitente"`

	Destinatario *UnionNullArrayPersona `json:"destinatario"`

	Remito *UnionNullRemito `json:"remito"`

	CentroDeCosto *UnionNullString `json:"centroDeCosto"`

	ProductoAEntregar *UnionNullString `json:"productoAEntregar"`

	ProductoARetirar *UnionNullString `json:"productoARetirar"`

	TipoProducto *UnionNullString `json:"tipoProducto"`

	CategoriaDeFacturacion *UnionNullString `json:"categoriaDeFacturacion"`

	PagoDestino *UnionNullInt `json:"pagoDestino"`

	ValorACobrar *UnionNullFloat `json:"valorACobrar"`

	CodigoVerificadorDeEntrega *UnionNullString `json:"codigoVerificadorDeEntrega"`

	FechaYHorario *UnionNullFechaYHorario `json:"fechaYHorario"`

	PagoEnMostrador *UnionNullPagoEnMostrador `json:"PagoEnMostrador"`
}

const RequestAvroCRC64Fingerprint = "\x98\x90\xa8;\x9e\xac\xf2\a"

func NewRequest() Request {
	r := Request{}
	r.TipoServicio = nil
	r.SucursalClienteID = nil
	r.CentroDeCosto = nil
	r.ProductoAEntregar = nil
	r.ProductoARetirar = nil
	r.TipoProducto = nil
	r.CategoriaDeFacturacion = nil
	r.PagoDestino = nil
	r.ValorACobrar = nil
	r.CodigoVerificadorDeEntrega = nil
	r.FechaYHorario = nil
	r.PagoEnMostrador = nil
	return r
}

func DeserializeRequest(r io.Reader) (Request, error) {
	t := NewRequest()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRequestFromSchema(r io.Reader, schema string) (Request, error) {
	t := NewRequest()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRequest(r Request, w io.Writer) error {
	var err error
	err = writeUnionNullArrayBultoRequest(r.Bultos, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Contrato, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoServicio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.SucursalClienteID, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLugarDeRetiroOEntrega(r.Origen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLugarDeRetiroOEntrega(r.Destino, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.IdPedido, w)
	if err != nil {
		return err
	}
	err = writeUnionNullPersona(r.Remitente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayPersona(r.Destinatario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullRemito(r.Remito, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CentroDeCosto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ProductoAEntregar, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ProductoARetirar, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoProducto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CategoriaDeFacturacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.PagoDestino, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFloat(r.ValorACobrar, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoVerificadorDeEntrega, w)
	if err != nil {
		return err
	}
	err = writeUnionNullFechaYHorario(r.FechaYHorario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullPagoEnMostrador(r.PagoEnMostrador, w)
	if err != nil {
		return err
	}
	return err
}

func (r Request) Serialize(w io.Writer) error {
	return writeRequest(r, w)
}

func (r Request) Schema() string {
	return "{\"fields\":[{\"name\":\"Bultos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"Kilos\",\"type\":\"double\"},{\"name\":\"LargoCm\",\"type\":\"double\"},{\"name\":\"AltoCm\",\"type\":\"double\"},{\"name\":\"AnchoCm\",\"type\":\"double\"},{\"name\":\"VolumenCm\",\"type\":\"double\"},{\"name\":\"ValorDeclaradoSinImpuestos\",\"type\":\"double\"},{\"name\":\"ValorDeclaradoConImpuestos\",\"type\":\"double\"},{\"name\":\"ValorDeclarado\",\"type\":\"double\"},{\"default\":null,\"name\":\"Descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Referencias\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"NumeroDeEnvio\",\"type\":[\"null\",\"string\"]}],\"name\":\"BultoRequest\",\"type\":\"record\"},\"type\":\"array\"}]},{\"name\":\"Contrato\",\"type\":\"string\"},{\"default\":null,\"name\":\"TipoServicio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SucursalClienteID\",\"type\":[\"null\",\"int\"]},{\"name\":\"origen\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"postal\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"abreviaturaProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoDeDireccion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoCiudad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefono\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO -3166-2 (https://es.wikipedia.org/wiki/ISO_3166-2:AR)\",\"name\":\"codigoISOProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO 3166-1 alpha-2 (https://es.wikipedia.org/wiki/ISO_3166-1#Tabla_de_c.C3.B3digos_alfa-2_asignados_o_reservados)\",\"name\":\"codigoISOPais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":\"Andreani.AltaOrdenEnvio.Events.Common.Metadato\",\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"casillaDeCorreo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"region\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]}],\"name\":\"Direccion\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"datosSucursal\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"ID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nomenclatura\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"direccion\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDePropiedades\"]},{\"default\":null,\"name\":\"casillaDeCorreo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"region\",\"type\":[\"null\",\"string\"]}],\"name\":\"DireccionPostal\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"datosAdicionales\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDePropiedades\"]},{\"default\":null,\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"tipo\",\"type\":\"string\"},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeTelefonos\",\"type\":\"record\"}]}],\"name\":\"Sucursal\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"GeoCoordenada\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"latitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"longitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"elevacion\",\"type\":[\"null\",\"double\"]}],\"name\":\"GeoCoordenada\",\"type\":\"record\"}]}],\"name\":\"LugarDeRetiroOEntrega\",\"type\":\"record\"}]},{\"name\":\"destino\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.LugarDeRetiroOEntrega\"]},{\"name\":\"idPedido\",\"type\":[\"null\",\"string\"]},{\"name\":\"remitente\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"email\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"documentoTipo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"documentoNumero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDeTelefonos\"]}],\"name\":\"Persona\",\"type\":\"record\"}]},{\"name\":\"destinatario\",\"type\":[\"null\",{\"items\":\"Andreani.AltaOrdenEnvio.Events.Common.Persona\",\"type\":\"array\"}]},{\"name\":\"remito\",\"type\":[\"null\",{\"fields\":[{\"name\":\"complementarios\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"name\":\"numeroDeRemito\",\"type\":[\"null\",\"string\"]}],\"name\":\"Remito\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"centroDeCosto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"productoAEntregar\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"productoARetirar\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoProducto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"categoriaDeFacturacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pagoDestino\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"valorACobrar\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"codigoVerificadorDeEntrega\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaYHorario\",\"type\":[\"null\",{\"fields\":[{\"name\":\"fecha\",\"type\":[\"null\",\"string\"]},{\"name\":\"horaDesde\",\"type\":[\"null\",\"string\"]},{\"name\":\"horaHasta\",\"type\":[\"null\",\"string\"]}],\"name\":\"FechaYHorario\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"PagoEnMostrador\",\"type\":[\"null\",{\"fields\":[{\"name\":\"MontoACobrar\",\"type\":\"float\"},{\"default\":null,\"name\":\"documentoTipo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"documentoNumero\",\"type\":[\"null\",\"string\"]}],\"name\":\"PagoEnMostrador\",\"type\":\"record\"}]}],\"name\":\"Andreani.AltaOrdenEnvio.Events.Common.Request\",\"type\":\"record\"}"
}

func (r Request) SchemaName() string {
	return "Andreani.AltaOrdenEnvio.Events.Common.Request"
}

func (_ Request) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Request) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Request) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Request) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Request) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Request) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Request) SetString(v string)   { panic("Unsupported operation") }
func (_ Request) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Request) Get(i int) types.Field {
	switch i {
	case 0:
		r.Bultos = NewUnionNullArrayBultoRequest()

		return r.Bultos
	case 1:
		w := types.String{Target: &r.Contrato}

		return w

	case 2:
		r.TipoServicio = NewUnionNullString()

		return r.TipoServicio
	case 3:
		r.SucursalClienteID = NewUnionNullInt()

		return r.SucursalClienteID
	case 4:
		r.Origen = NewUnionNullLugarDeRetiroOEntrega()

		return r.Origen
	case 5:
		r.Destino = NewUnionNullLugarDeRetiroOEntrega()

		return r.Destino
	case 6:
		r.IdPedido = NewUnionNullString()

		return r.IdPedido
	case 7:
		r.Remitente = NewUnionNullPersona()

		return r.Remitente
	case 8:
		r.Destinatario = NewUnionNullArrayPersona()

		return r.Destinatario
	case 9:
		r.Remito = NewUnionNullRemito()

		return r.Remito
	case 10:
		r.CentroDeCosto = NewUnionNullString()

		return r.CentroDeCosto
	case 11:
		r.ProductoAEntregar = NewUnionNullString()

		return r.ProductoAEntregar
	case 12:
		r.ProductoARetirar = NewUnionNullString()

		return r.ProductoARetirar
	case 13:
		r.TipoProducto = NewUnionNullString()

		return r.TipoProducto
	case 14:
		r.CategoriaDeFacturacion = NewUnionNullString()

		return r.CategoriaDeFacturacion
	case 15:
		r.PagoDestino = NewUnionNullInt()

		return r.PagoDestino
	case 16:
		r.ValorACobrar = NewUnionNullFloat()

		return r.ValorACobrar
	case 17:
		r.CodigoVerificadorDeEntrega = NewUnionNullString()

		return r.CodigoVerificadorDeEntrega
	case 18:
		r.FechaYHorario = NewUnionNullFechaYHorario()

		return r.FechaYHorario
	case 19:
		r.PagoEnMostrador = NewUnionNullPagoEnMostrador()

		return r.PagoEnMostrador
	}
	panic("Unknown field index")
}

func (r *Request) SetDefault(i int) {
	switch i {
	case 2:
		r.TipoServicio = nil
		return
	case 3:
		r.SucursalClienteID = nil
		return
	case 10:
		r.CentroDeCosto = nil
		return
	case 11:
		r.ProductoAEntregar = nil
		return
	case 12:
		r.ProductoARetirar = nil
		return
	case 13:
		r.TipoProducto = nil
		return
	case 14:
		r.CategoriaDeFacturacion = nil
		return
	case 15:
		r.PagoDestino = nil
		return
	case 16:
		r.ValorACobrar = nil
		return
	case 17:
		r.CodigoVerificadorDeEntrega = nil
		return
	case 18:
		r.FechaYHorario = nil
		return
	case 19:
		r.PagoEnMostrador = nil
		return
	}
	panic("Unknown field index")
}

func (r *Request) NullField(i int) {
	switch i {
	case 0:
		r.Bultos = nil
		return
	case 2:
		r.TipoServicio = nil
		return
	case 3:
		r.SucursalClienteID = nil
		return
	case 4:
		r.Origen = nil
		return
	case 5:
		r.Destino = nil
		return
	case 6:
		r.IdPedido = nil
		return
	case 7:
		r.Remitente = nil
		return
	case 8:
		r.Destinatario = nil
		return
	case 9:
		r.Remito = nil
		return
	case 10:
		r.CentroDeCosto = nil
		return
	case 11:
		r.ProductoAEntregar = nil
		return
	case 12:
		r.ProductoARetirar = nil
		return
	case 13:
		r.TipoProducto = nil
		return
	case 14:
		r.CategoriaDeFacturacion = nil
		return
	case 15:
		r.PagoDestino = nil
		return
	case 16:
		r.ValorACobrar = nil
		return
	case 17:
		r.CodigoVerificadorDeEntrega = nil
		return
	case 18:
		r.FechaYHorario = nil
		return
	case 19:
		r.PagoEnMostrador = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Request) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Request) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Request) HintSize(int)                     { panic("Unsupported operation") }
func (_ Request) Finalize()                        {}

func (_ Request) AvroCRC64Fingerprint() []byte {
	return []byte(RequestAvroCRC64Fingerprint)
}

func (r Request) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Bultos"], err = json.Marshal(r.Bultos)
	if err != nil {
		return nil, err
	}
	output["Contrato"], err = json.Marshal(r.Contrato)
	if err != nil {
		return nil, err
	}
	output["TipoServicio"], err = json.Marshal(r.TipoServicio)
	if err != nil {
		return nil, err
	}
	output["SucursalClienteID"], err = json.Marshal(r.SucursalClienteID)
	if err != nil {
		return nil, err
	}
	output["origen"], err = json.Marshal(r.Origen)
	if err != nil {
		return nil, err
	}
	output["destino"], err = json.Marshal(r.Destino)
	if err != nil {
		return nil, err
	}
	output["idPedido"], err = json.Marshal(r.IdPedido)
	if err != nil {
		return nil, err
	}
	output["remitente"], err = json.Marshal(r.Remitente)
	if err != nil {
		return nil, err
	}
	output["destinatario"], err = json.Marshal(r.Destinatario)
	if err != nil {
		return nil, err
	}
	output["remito"], err = json.Marshal(r.Remito)
	if err != nil {
		return nil, err
	}
	output["centroDeCosto"], err = json.Marshal(r.CentroDeCosto)
	if err != nil {
		return nil, err
	}
	output["productoAEntregar"], err = json.Marshal(r.ProductoAEntregar)
	if err != nil {
		return nil, err
	}
	output["productoARetirar"], err = json.Marshal(r.ProductoARetirar)
	if err != nil {
		return nil, err
	}
	output["tipoProducto"], err = json.Marshal(r.TipoProducto)
	if err != nil {
		return nil, err
	}
	output["categoriaDeFacturacion"], err = json.Marshal(r.CategoriaDeFacturacion)
	if err != nil {
		return nil, err
	}
	output["pagoDestino"], err = json.Marshal(r.PagoDestino)
	if err != nil {
		return nil, err
	}
	output["valorACobrar"], err = json.Marshal(r.ValorACobrar)
	if err != nil {
		return nil, err
	}
	output["codigoVerificadorDeEntrega"], err = json.Marshal(r.CodigoVerificadorDeEntrega)
	if err != nil {
		return nil, err
	}
	output["fechaYHorario"], err = json.Marshal(r.FechaYHorario)
	if err != nil {
		return nil, err
	}
	output["PagoEnMostrador"], err = json.Marshal(r.PagoEnMostrador)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Request) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Bultos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Bultos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Bultos")
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
		if v, ok := fields["TipoServicio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoServicio); err != nil {
			return err
		}
	} else {
		r.TipoServicio = NewUnionNullString()

		r.TipoServicio = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["SucursalClienteID"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalClienteID); err != nil {
			return err
		}
	} else {
		r.SucursalClienteID = NewUnionNullInt()

		r.SucursalClienteID = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["origen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Origen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for origen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["destino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Destino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for destino")
	}
	val = func() json.RawMessage {
		if v, ok := fields["idPedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdPedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idPedido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["remitente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Remitente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for remitente")
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
		if v, ok := fields["centroDeCosto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CentroDeCosto); err != nil {
			return err
		}
	} else {
		r.CentroDeCosto = NewUnionNullString()

		r.CentroDeCosto = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["productoAEntregar"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ProductoAEntregar); err != nil {
			return err
		}
	} else {
		r.ProductoAEntregar = NewUnionNullString()

		r.ProductoAEntregar = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["productoARetirar"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ProductoARetirar); err != nil {
			return err
		}
	} else {
		r.ProductoARetirar = NewUnionNullString()

		r.ProductoARetirar = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoProducto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoProducto); err != nil {
			return err
		}
	} else {
		r.TipoProducto = NewUnionNullString()

		r.TipoProducto = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["categoriaDeFacturacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CategoriaDeFacturacion); err != nil {
			return err
		}
	} else {
		r.CategoriaDeFacturacion = NewUnionNullString()

		r.CategoriaDeFacturacion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["pagoDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PagoDestino); err != nil {
			return err
		}
	} else {
		r.PagoDestino = NewUnionNullInt()

		r.PagoDestino = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["valorACobrar"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ValorACobrar); err != nil {
			return err
		}
	} else {
		r.ValorACobrar = NewUnionNullFloat()

		r.ValorACobrar = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoVerificadorDeEntrega"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoVerificadorDeEntrega); err != nil {
			return err
		}
	} else {
		r.CodigoVerificadorDeEntrega = NewUnionNullString()

		r.CodigoVerificadorDeEntrega = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaYHorario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaYHorario); err != nil {
			return err
		}
	} else {
		r.FechaYHorario = NewUnionNullFechaYHorario()

		r.FechaYHorario = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["PagoEnMostrador"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PagoEnMostrador); err != nil {
			return err
		}
	} else {
		r.PagoEnMostrador = NewUnionNullPagoEnMostrador()

		r.PagoEnMostrador = nil
	}
	return nil
}
