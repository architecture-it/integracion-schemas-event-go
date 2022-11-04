// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     OrdenDeEnvioSolicitada.avsc
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

type OrdenDeEnvioSolicitada struct {
	Remitente *UnionNullString `json:"Remitente"`

	Destinatario *UnionNullString `json:"Destinatario"`

	NumeroDeOrden *UnionNullString `json:"NumeroDeOrden"`

	Vencimiento int32 `json:"Vencimiento"`

	CodigoDeEnvio *UnionNullString `json:"CodigoDeEnvio"`

	Cuando *UnionNullString `json:"Cuando"`

	CodigoDeContratoInterno *UnionNullString `json:"CodigoDeContratoInterno"`

	Solicitante *UnionNullString `json:"Solicitante"`

	Comentario *UnionNullString `json:"Comentario"`

	Topic *UnionNullString `json:"Topic"`

	DatosDeLaOrden *UnionNullDetalleDeOrdenDeEnvio `json:"datosDeLaOrden"`
}

const OrdenDeEnvioSolicitadaAvroCRC64Fingerprint = "\xe29\xa4\xdfhЮ\x0e"

func NewOrdenDeEnvioSolicitada() OrdenDeEnvioSolicitada {
	r := OrdenDeEnvioSolicitada{}
	r.Remitente = nil
	r.Destinatario = nil
	r.NumeroDeOrden = nil
	r.CodigoDeEnvio = nil
	r.Cuando = nil
	r.CodigoDeContratoInterno = nil
	r.Solicitante = nil
	r.Comentario = nil
	r.Topic = nil
	r.DatosDeLaOrden = nil
	return r
}

func DeserializeOrdenDeEnvioSolicitada(r io.Reader) (OrdenDeEnvioSolicitada, error) {
	t := NewOrdenDeEnvioSolicitada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeOrdenDeEnvioSolicitadaFromSchema(r io.Reader, schema string) (OrdenDeEnvioSolicitada, error) {
	t := NewOrdenDeEnvioSolicitada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeOrdenDeEnvioSolicitada(r OrdenDeEnvioSolicitada, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Remitente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Destinatario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeOrden, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Vencimiento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoDeEnvio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Cuando, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoDeContratoInterno, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Solicitante, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Comentario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Topic, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDetalleDeOrdenDeEnvio(r.DatosDeLaOrden, w)
	if err != nil {
		return err
	}
	return err
}

func (r OrdenDeEnvioSolicitada) Serialize(w io.Writer) error {
	return writeOrdenDeEnvioSolicitada(r, w)
}

func (r OrdenDeEnvioSolicitada) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"Remitente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"name\":\"Vencimiento\",\"type\":{\"logicalType\":\"date\",\"type\":\"int\"}},{\"default\":null,\"name\":\"CodigoDeEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Cuando\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoDeContratoInterno\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Solicitante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Topic\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosDeLaOrden\",\"type\":[\"null\",{\"fields\":[{\"name\":\"numeroDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"remitente\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idInternoDelCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"eMail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"tipo\",\"type\":\"string\"},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeTelefonos\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"agrupador\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipoDeDocumento\",\"type\":{\"name\":\"TipoDeDocumento\",\"symbols\":[\"undefined\",\"DNI\",\"CUIT\",\"CUIL\"],\"type\":\"enum\"}}],\"name\":\"DatosPersonales\",\"type\":\"record\"}]},{\"name\":\"destinatario\",\"type\":\"Andreani.AltaOrdenEnvio.Events.Common.DatosPersonales\"},{\"default\":null,\"name\":\"destinatarioAlternativo\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.DatosPersonales\"]},{\"name\":\"destino\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"postal\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"abreviaturaProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoDeDireccion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoCiudad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefono\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO -3166-2 (https://es.wikipedia.org/wiki/ISO_3166-2:AR)\",\"name\":\"codigoISOProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO 3166-1 alpha-2 (https://es.wikipedia.org/wiki/ISO_3166-1#Tabla_de_c.C3.B3digos_alfa-2_asignados_o_reservados)\",\"name\":\"codigoISOPais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"casillaDeCorreo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"region\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]}],\"name\":\"Direccion\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"datosSucursal\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"ID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nomenclatura\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"direccion\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Region\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDePropiedades\"]},{\"default\":null,\"name\":\"casillaDeCorreo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"region\",\"type\":[\"null\",\"string\"]}],\"name\":\"DireccionPostal\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"datosAdicionales\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDePropiedades\"]},{\"default\":null,\"name\":\"listaDeTelefonos\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDeTelefonos\"]}],\"name\":\"Sucursal\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"GeoCoordenada\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"latitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"longitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"elevacion\",\"type\":[\"null\",\"double\"]}],\"name\":\"GeoCoordenada\",\"type\":\"record\"}]}],\"name\":\"LugarDeRetiroOEntrega\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"origen\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.LugarDeRetiroOEntrega\"]},{\"name\":\"contrato\",\"type\":\"string\"},{\"default\":null,\"name\":\"tipoDeServicio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeRemito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remitosComplementarios\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"centroDeCostosDelCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"categoriaDeProducto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"categoriaDeFacturacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"valorACobrar\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"fechaPactadaDeEntrega\",\"type\":[\"null\",{\"fields\":[{\"name\":\"horaDesde\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"name\":\"horaHasta\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]}],\"name\":\"FechaPactadaDeEntrega\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"productoAEntregar\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"productoARetirar\",\"type\":[\"null\",\"string\"]},{\"name\":\"pagoDestino\",\"type\":{\"name\":\"TipoPagoDestino\",\"symbols\":[\"undefined\",\"P\",\"D\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"sucursalDeDistribucion\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"sucursalCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"listaPaquetes\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaPaquetes\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"pesoEnKg\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"altoEnCm\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"anchoEnCm\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"largoEnCm\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"referenciasDelCliente\",\"type\":[\"null\",{\"type\":\"map\",\"values\":\"string\"}]},{\"default\":null,\"name\":\"volumenEnCm3\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"valorDeclaradoSinImpuesto\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"valorDeclaradoConImpuesto\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"numeroDeBulto\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"valorDeclarado\",\"type\":[\"null\",\"double\"]}],\"name\":\"DetalleDePaquete\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaPaquetes\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"codigoVerificadorDeEntrega\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cantidadDeBultos\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"agrupadorDeBulto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idPedido\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pagoEnMostrador\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.TipoPagoDestino\"]}],\"name\":\"DetalleDeOrdenDeEnvio\",\"namespace\":\"Andreani.AltaOrdenEnvio.Events.Common\",\"type\":\"record\"}]}],\"name\":\"Andreani.AltaOrdenEnvio.Events.Record.OrdenDeEnvioSolicitada\",\"type\":\"record\"}"
}

func (r OrdenDeEnvioSolicitada) SchemaName() string {
	return "Andreani.AltaOrdenEnvio.Events.Record.OrdenDeEnvioSolicitada"
}

func (_ OrdenDeEnvioSolicitada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ OrdenDeEnvioSolicitada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ OrdenDeEnvioSolicitada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ OrdenDeEnvioSolicitada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ OrdenDeEnvioSolicitada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ OrdenDeEnvioSolicitada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ OrdenDeEnvioSolicitada) SetString(v string)   { panic("Unsupported operation") }
func (_ OrdenDeEnvioSolicitada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *OrdenDeEnvioSolicitada) Get(i int) types.Field {
	switch i {
	case 0:
		r.Remitente = NewUnionNullString()

		return r.Remitente
	case 1:
		r.Destinatario = NewUnionNullString()

		return r.Destinatario
	case 2:
		r.NumeroDeOrden = NewUnionNullString()

		return r.NumeroDeOrden
	case 3:
		w := types.Int{Target: &r.Vencimiento}

		return w

	case 4:
		r.CodigoDeEnvio = NewUnionNullString()

		return r.CodigoDeEnvio
	case 5:
		r.Cuando = NewUnionNullString()

		return r.Cuando
	case 6:
		r.CodigoDeContratoInterno = NewUnionNullString()

		return r.CodigoDeContratoInterno
	case 7:
		r.Solicitante = NewUnionNullString()

		return r.Solicitante
	case 8:
		r.Comentario = NewUnionNullString()

		return r.Comentario
	case 9:
		r.Topic = NewUnionNullString()

		return r.Topic
	case 10:
		r.DatosDeLaOrden = NewUnionNullDetalleDeOrdenDeEnvio()

		return r.DatosDeLaOrden
	}
	panic("Unknown field index")
}

func (r *OrdenDeEnvioSolicitada) SetDefault(i int) {
	switch i {
	case 0:
		r.Remitente = nil
		return
	case 1:
		r.Destinatario = nil
		return
	case 2:
		r.NumeroDeOrden = nil
		return
	case 4:
		r.CodigoDeEnvio = nil
		return
	case 5:
		r.Cuando = nil
		return
	case 6:
		r.CodigoDeContratoInterno = nil
		return
	case 7:
		r.Solicitante = nil
		return
	case 8:
		r.Comentario = nil
		return
	case 9:
		r.Topic = nil
		return
	case 10:
		r.DatosDeLaOrden = nil
		return
	}
	panic("Unknown field index")
}

func (r *OrdenDeEnvioSolicitada) NullField(i int) {
	switch i {
	case 0:
		r.Remitente = nil
		return
	case 1:
		r.Destinatario = nil
		return
	case 2:
		r.NumeroDeOrden = nil
		return
	case 4:
		r.CodigoDeEnvio = nil
		return
	case 5:
		r.Cuando = nil
		return
	case 6:
		r.CodigoDeContratoInterno = nil
		return
	case 7:
		r.Solicitante = nil
		return
	case 8:
		r.Comentario = nil
		return
	case 9:
		r.Topic = nil
		return
	case 10:
		r.DatosDeLaOrden = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ OrdenDeEnvioSolicitada) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ OrdenDeEnvioSolicitada) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ OrdenDeEnvioSolicitada) HintSize(int)                     { panic("Unsupported operation") }
func (_ OrdenDeEnvioSolicitada) Finalize()                        {}

func (_ OrdenDeEnvioSolicitada) AvroCRC64Fingerprint() []byte {
	return []byte(OrdenDeEnvioSolicitadaAvroCRC64Fingerprint)
}

func (r OrdenDeEnvioSolicitada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Remitente"], err = json.Marshal(r.Remitente)
	if err != nil {
		return nil, err
	}
	output["Destinatario"], err = json.Marshal(r.Destinatario)
	if err != nil {
		return nil, err
	}
	output["NumeroDeOrden"], err = json.Marshal(r.NumeroDeOrden)
	if err != nil {
		return nil, err
	}
	output["Vencimiento"], err = json.Marshal(r.Vencimiento)
	if err != nil {
		return nil, err
	}
	output["CodigoDeEnvio"], err = json.Marshal(r.CodigoDeEnvio)
	if err != nil {
		return nil, err
	}
	output["Cuando"], err = json.Marshal(r.Cuando)
	if err != nil {
		return nil, err
	}
	output["CodigoDeContratoInterno"], err = json.Marshal(r.CodigoDeContratoInterno)
	if err != nil {
		return nil, err
	}
	output["Solicitante"], err = json.Marshal(r.Solicitante)
	if err != nil {
		return nil, err
	}
	output["Comentario"], err = json.Marshal(r.Comentario)
	if err != nil {
		return nil, err
	}
	output["Topic"], err = json.Marshal(r.Topic)
	if err != nil {
		return nil, err
	}
	output["datosDeLaOrden"], err = json.Marshal(r.DatosDeLaOrden)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *OrdenDeEnvioSolicitada) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
		r.Remitente = NewUnionNullString()

		r.Remitente = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Destinatario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Destinatario); err != nil {
			return err
		}
	} else {
		r.Destinatario = NewUnionNullString()

		r.Destinatario = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeOrden"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeOrden); err != nil {
			return err
		}
	} else {
		r.NumeroDeOrden = NewUnionNullString()

		r.NumeroDeOrden = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Vencimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Vencimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Vencimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoDeEnvio"]; ok {
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
		if v, ok := fields["Cuando"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cuando); err != nil {
			return err
		}
	} else {
		r.Cuando = NewUnionNullString()

		r.Cuando = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoDeContratoInterno"]; ok {
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
		if v, ok := fields["Solicitante"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Solicitante); err != nil {
			return err
		}
	} else {
		r.Solicitante = NewUnionNullString()

		r.Solicitante = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Comentario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Comentario); err != nil {
			return err
		}
	} else {
		r.Comentario = NewUnionNullString()

		r.Comentario = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Topic"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Topic); err != nil {
			return err
		}
	} else {
		r.Topic = NewUnionNullString()

		r.Topic = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["datosDeLaOrden"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DatosDeLaOrden); err != nil {
			return err
		}
	} else {
		r.DatosDeLaOrden = NewUnionNullDetalleDeOrdenDeEnvio()

		r.DatosDeLaOrden = nil
	}
	return nil
}
