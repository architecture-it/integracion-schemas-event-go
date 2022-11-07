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

type TransactionEvent struct {
	NumerosAndreani *UnionNullArrayString `json:"numerosAndreani"`

	EstadoDeLaOrden *UnionNullString `json:"estadoDeLaOrden"`

	AgrupadorDeBulto *UnionNullString `json:"agrupadorDeBulto"`

	Contrato *UnionNullContrato `json:"contrato"`

	Request *UnionNullRequest `json:"request"`

	Response *UnionNullResponse `json:"response"`
}

const TransactionEventAvroCRC64Fingerprint = "\xc1F\xe8\x89+\xd3\xe0\xc2"

func NewTransactionEvent() TransactionEvent {
	r := TransactionEvent{}
	r.NumerosAndreani = nil
	r.EstadoDeLaOrden = nil
	r.AgrupadorDeBulto = nil
	r.Contrato = nil
	return r
}

func DeserializeTransactionEvent(r io.Reader) (TransactionEvent, error) {
	t := NewTransactionEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTransactionEventFromSchema(r io.Reader, schema string) (TransactionEvent, error) {
	t := NewTransactionEvent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTransactionEvent(r TransactionEvent, w io.Writer) error {
	var err error
	err = writeUnionNullArrayString(r.NumerosAndreani, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstadoDeLaOrden, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.AgrupadorDeBulto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullContrato(r.Contrato, w)
	if err != nil {
		return err
	}
	err = writeUnionNullRequest(r.Request, w)
	if err != nil {
		return err
	}
	err = writeUnionNullResponse(r.Response, w)
	if err != nil {
		return err
	}
	return err
}

func (r TransactionEvent) Serialize(w io.Writer) error {
	return writeTransactionEvent(r, w)
}

func (r TransactionEvent) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"numerosAndreani\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"estadoDeLaOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"agrupadorDeBulto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contrato\",\"type\":[\"null\",{\"fields\":[{\"name\":\"contratoTms\",\"type\":[\"null\",\"string\"]},{\"name\":\"clienteTmsCodigo\",\"type\":[\"null\",\"string\"]},{\"name\":\"codigo\",\"type\":[\"null\",\"string\"]},{\"name\":\"Descripcion\",\"type\":[\"null\",\"string\"]},{\"name\":\"SolicitanteCodigo\",\"type\":[\"null\",\"string\"]},{\"name\":\"SolicitanteDescripcion\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioCodigo\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioDescripcion\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioNombreCorto\",\"type\":[\"null\",\"string\"]},{\"name\":\"Segmento\",\"type\":[\"null\",\"string\"]},{\"name\":\"PlazoEntrega\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoEntrega\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoServicio\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoServicioCodigo\",\"type\":[\"null\",\"string\"]},{\"name\":\"InsertTS\",\"type\":[\"null\",\"string\"]},{\"name\":\"Cuit\",\"type\":[\"null\",\"string\"]},{\"name\":\"Canal\",\"type\":[\"null\",\"string\"]},{\"name\":\"SucursalRendicion\",\"type\":[\"null\",\"string\"]}],\"name\":\"Contrato\",\"namespace\":\"Andreani.AltaOrdenEnvio.Events.Common\",\"type\":\"record\"}]},{\"name\":\"request\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Bultos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"Kilos\",\"type\":\"double\"},{\"name\":\"LargoCm\",\"type\":\"double\"},{\"name\":\"AltoCm\",\"type\":\"double\"},{\"name\":\"AnchoCm\",\"type\":\"double\"},{\"name\":\"VolumenCm\",\"type\":\"double\"},{\"name\":\"ValorDeclaradoSinImpuestos\",\"type\":\"double\"},{\"name\":\"ValorDeclaradoConImpuestos\",\"type\":\"double\"},{\"name\":\"ValorDeclarado\",\"type\":\"double\"},{\"default\":null,\"name\":\"Descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Referencias\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"NumeroDeEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PagoEnMostrador\",\"type\":[\"null\",{\"fields\":[{\"name\":\"MontoACobrar\",\"type\":\"float\"},{\"default\":null,\"name\":\"documentoTipo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"documentoNumero\",\"type\":[\"null\",\"string\"]}],\"name\":\"PagoEnMostrador\",\"type\":\"record\"}]}],\"name\":\"BultoRequest\",\"type\":\"record\"},\"type\":\"array\"}]},{\"name\":\"Contrato\",\"type\":\"string\"},{\"default\":null,\"name\":\"TipoServicio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SucursalClienteID\",\"type\":[\"null\",\"int\"]},{\"name\":\"origen\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"postal\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"abreviaturaProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoDeDireccion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoCiudad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefono\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO -3166-2 (https://es.wikipedia.org/wiki/ISO_3166-2:AR)\",\"name\":\"codigoISOProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO 3166-1 alpha-2 (https://es.wikipedia.org/wiki/ISO_3166-1#Tabla_de_c.C3.B3digos_alfa-2_asignados_o_reservados)\",\"name\":\"codigoISOPais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":\"Andreani.AltaOrdenEnvio.Events.Common.Metadato\",\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"casillaDeCorreo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"region\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]}],\"name\":\"Direccion\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"datosSucursal\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"ID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nomenclatura\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"direccion\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Region\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDePropiedades\"]},{\"default\":null,\"name\":\"casillaDeCorreo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"region\",\"type\":[\"null\",\"string\"]}],\"name\":\"DireccionPostal\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"datosAdicionales\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDePropiedades\"]},{\"default\":null,\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"tipo\",\"type\":\"string\"},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeTelefonos\",\"type\":\"record\"}]}],\"name\":\"Sucursal\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"GeoCoordenada\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"latitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"longitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"elevacion\",\"type\":[\"null\",\"double\"]}],\"name\":\"GeoCoordenada\",\"type\":\"record\"}]}],\"name\":\"LugarDeRetiroOEntrega\",\"type\":\"record\"}]},{\"name\":\"destino\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.LugarDeRetiroOEntrega\"]},{\"name\":\"idPedido\",\"type\":[\"null\",\"string\"]},{\"name\":\"remitente\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"email\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"documentoTipo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"documentoNumero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDeTelefonos\"]}],\"name\":\"Persona\",\"type\":\"record\"}]},{\"name\":\"destinatario\",\"type\":[\"null\",{\"items\":\"Andreani.AltaOrdenEnvio.Events.Common.Persona\",\"type\":\"array\"}]},{\"name\":\"remito\",\"type\":[\"null\",{\"fields\":[{\"name\":\"complementarios\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"name\":\"numeroDeRemito\",\"type\":[\"null\",\"string\"]}],\"name\":\"Remito\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"centroDeCosto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"productoAEntregar\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"productoARetirar\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoProducto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"categoriaDeFacturacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pagoDestino\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"valorACobrar\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"codigoVerificadorDeEntrega\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaYHorario\",\"type\":[\"null\",{\"fields\":[{\"name\":\"fecha\",\"type\":[\"null\",\"string\"]},{\"name\":\"horaDesde\",\"type\":[\"null\",\"string\"]},{\"name\":\"horaHasta\",\"type\":[\"null\",\"string\"]}],\"name\":\"FechaYHorario\",\"type\":\"record\"}]}],\"name\":\"Request\",\"namespace\":\"Andreani.AltaOrdenEnvio.Events.Common\",\"type\":\"record\"}]},{\"name\":\"response\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"estado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"motivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalDeDistribucion\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"ID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nomenclatura\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]}],\"name\":\"SucursalResponse\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"sucursalDeImposicion\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.SucursalResponse\"]},{\"default\":null,\"name\":\"sucursalDeRendicion\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.SucursalResponse\"]},{\"default\":null,\"name\":\"sucursalAbastecedora\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.SucursalResponse\"]},{\"default\":null,\"name\":\"fechaCreacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"zonaDeReparto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDePermisionaria\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"descripcionServicio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"etiquetaRemito\",\"type\":[\"null\",\"string\"]},{\"name\":\"Bultos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"NumeroDeBulto\",\"type\":\"string\"},{\"name\":\"NumeroDeEnvio\",\"type\":\"string\"},{\"name\":\"Totalizador\",\"type\":\"string\"},{\"default\":null,\"name\":\"Linking\",\"type\":[\"null\",{\"items\":\"Andreani.AltaOrdenEnvio.Events.Common.Metadato\",\"type\":\"array\"}]}],\"name\":\"BultoResponse\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"FechaEstimadaDeEntrega\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"huellaDeCarbono\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"gastoEnergetico\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"agrupadorDeBultos\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"etiquetasPorAgrupador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"etiquetasDocumentoDeCambio\",\"type\":[\"null\",\"string\"]}],\"name\":\"Response\",\"namespace\":\"Andreani.AltaOrdenEnvio.Events.Common\",\"type\":\"record\"}]}],\"name\":\"Andreani.AltaOrdenEnvio.Events.Record.TransactionEvent\",\"type\":\"record\"}"
}

func (r TransactionEvent) SchemaName() string {
	return "Andreani.AltaOrdenEnvio.Events.Record.TransactionEvent"
}

func (_ TransactionEvent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ TransactionEvent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ TransactionEvent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ TransactionEvent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ TransactionEvent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ TransactionEvent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ TransactionEvent) SetString(v string)   { panic("Unsupported operation") }
func (_ TransactionEvent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *TransactionEvent) Get(i int) types.Field {
	switch i {
	case 0:
		r.NumerosAndreani = NewUnionNullArrayString()

		return r.NumerosAndreani
	case 1:
		r.EstadoDeLaOrden = NewUnionNullString()

		return r.EstadoDeLaOrden
	case 2:
		r.AgrupadorDeBulto = NewUnionNullString()

		return r.AgrupadorDeBulto
	case 3:
		r.Contrato = NewUnionNullContrato()

		return r.Contrato
	case 4:
		r.Request = NewUnionNullRequest()

		return r.Request
	case 5:
		r.Response = NewUnionNullResponse()

		return r.Response
	}
	panic("Unknown field index")
}

func (r *TransactionEvent) SetDefault(i int) {
	switch i {
	case 0:
		r.NumerosAndreani = nil
		return
	case 1:
		r.EstadoDeLaOrden = nil
		return
	case 2:
		r.AgrupadorDeBulto = nil
		return
	case 3:
		r.Contrato = nil
		return
	}
	panic("Unknown field index")
}

func (r *TransactionEvent) NullField(i int) {
	switch i {
	case 0:
		r.NumerosAndreani = nil
		return
	case 1:
		r.EstadoDeLaOrden = nil
		return
	case 2:
		r.AgrupadorDeBulto = nil
		return
	case 3:
		r.Contrato = nil
		return
	case 4:
		r.Request = nil
		return
	case 5:
		r.Response = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ TransactionEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ TransactionEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ TransactionEvent) HintSize(int)                     { panic("Unsupported operation") }
func (_ TransactionEvent) Finalize()                        {}

func (_ TransactionEvent) AvroCRC64Fingerprint() []byte {
	return []byte(TransactionEventAvroCRC64Fingerprint)
}

func (r TransactionEvent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["numerosAndreani"], err = json.Marshal(r.NumerosAndreani)
	if err != nil {
		return nil, err
	}
	output["estadoDeLaOrden"], err = json.Marshal(r.EstadoDeLaOrden)
	if err != nil {
		return nil, err
	}
	output["agrupadorDeBulto"], err = json.Marshal(r.AgrupadorDeBulto)
	if err != nil {
		return nil, err
	}
	output["contrato"], err = json.Marshal(r.Contrato)
	if err != nil {
		return nil, err
	}
	output["request"], err = json.Marshal(r.Request)
	if err != nil {
		return nil, err
	}
	output["response"], err = json.Marshal(r.Response)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *TransactionEvent) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["numerosAndreani"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumerosAndreani); err != nil {
			return err
		}
	} else {
		r.NumerosAndreani = NewUnionNullArrayString()

		r.NumerosAndreani = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["estadoDeLaOrden"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoDeLaOrden); err != nil {
			return err
		}
	} else {
		r.EstadoDeLaOrden = NewUnionNullString()

		r.EstadoDeLaOrden = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["agrupadorDeBulto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AgrupadorDeBulto); err != nil {
			return err
		}
	} else {
		r.AgrupadorDeBulto = NewUnionNullString()

		r.AgrupadorDeBulto = nil
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
		r.Contrato = NewUnionNullContrato()

		r.Contrato = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["request"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Request); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for request")
	}
	val = func() json.RawMessage {
		if v, ok := fields["response"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Response); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for response")
	}
	return nil
}
