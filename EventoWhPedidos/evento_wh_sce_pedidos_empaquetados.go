// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhSCEPedidosEmpaquetados.avsc
 */
package EventoWhPedidosEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EventoWhSCEPedidosEmpaquetados struct {
	Identificacion Identificacion `json:"Identificacion"`

	Cabecera Cabecera `json:"Cabecera"`

	Detalle []Detalle `json:"Detalle"`

	DetalleBultos []DetalleBultos `json:"DetalleBultos"`
}

const EventoWhSCEPedidosEmpaquetadosAvroCRC64Fingerprint = "\xf4\xd6\tUo\x90\xfcg"

func NewEventoWhSCEPedidosEmpaquetados() EventoWhSCEPedidosEmpaquetados {
	r := EventoWhSCEPedidosEmpaquetados{}
	r.Identificacion = NewIdentificacion()

	r.Cabecera = NewCabecera()

	r.Detalle = make([]Detalle, 0)

	r.DetalleBultos = make([]DetalleBultos, 0)

	return r
}

func DeserializeEventoWhSCEPedidosEmpaquetados(r io.Reader) (EventoWhSCEPedidosEmpaquetados, error) {
	t := NewEventoWhSCEPedidosEmpaquetados()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoWhSCEPedidosEmpaquetadosFromSchema(r io.Reader, schema string) (EventoWhSCEPedidosEmpaquetados, error) {
	t := NewEventoWhSCEPedidosEmpaquetados()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoWhSCEPedidosEmpaquetados(r EventoWhSCEPedidosEmpaquetados, w io.Writer) error {
	var err error
	err = writeIdentificacion(r.Identificacion, w)
	if err != nil {
		return err
	}
	err = writeCabecera(r.Cabecera, w)
	if err != nil {
		return err
	}
	err = writeArrayDetalle(r.Detalle, w)
	if err != nil {
		return err
	}
	err = writeArrayDetalleBultos(r.DetalleBultos, w)
	if err != nil {
		return err
	}
	return err
}

func (r EventoWhSCEPedidosEmpaquetados) Serialize(w io.Writer) error {
	return writeEventoWhSCEPedidosEmpaquetados(r, w)
}

func (r EventoWhSCEPedidosEmpaquetados) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.EventoWhPedidos.Events.SCEPedidosEmpaquetadosCommon\",\"type\":\"record\"}},{\"name\":\"Cabecera\",\"type\":{\"fields\":[{\"name\":\"OrdenWH\",\"type\":\"string\"},{\"name\":\"OrdenCliente\",\"type\":\"string\"},{\"default\":null,\"name\":\"Remito\",\"type\":[\"null\",\"string\"]},{\"name\":\"CodigoDestinatario\",\"type\":\"string\"},{\"default\":null,\"name\":\"TipoDeIntegracion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Acondicionamiento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"AlmacenCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PrioridadApiPlani\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CuentaTMS\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ContratoTMS\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FacturaLegal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FacturaInterna\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TieneGestionCobranza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ValorSeguro\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"AdmiteCambioLoteDirigido\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NoAdmitePickingParcial\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ImprimeDocumentacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroCale\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Cot\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ComprobanteIngresosBrutos\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ContratoServicioAlmacenes\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Marketplace\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoPedidoMatriz\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Contacto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Destinario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DestinatarioCalle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DestinarioNumero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DestinatarioPiso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DestinatarioDepartamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DestinatarioGLNDNI\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DestinatarioCiudad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DestinatarioProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DestinatarioCodigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DestinatarioTelefono\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DestinatarioEmail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"OrdenCompra\",\"type\":[\"null\",\"string\"]},{\"name\":\"ClientePadre\",\"type\":\"string\"},{\"name\":\"CodigoEstado\",\"type\":\"string\"},{\"default\":null,\"name\":\"CodigoDescripcion\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoPedidoCodigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"TipoPedidoDescripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoGrilla\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaEnvioFacturar\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaFacturacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"PrecioValorFC\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NotasPedido\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaCita\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaEntrega\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaExpedicionSolicitada\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaExpedicionProgramada\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaExpedicionPrometida\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaEntregaPlanificada\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaEntregaProgramada\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"GrupoOrdenes\",\"type\":\"string\"},{\"default\":null,\"name\":\"FranjaHorario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"EstadoOTAcondi\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"EstadoOTTraz\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaEvento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"CreacionPedido\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaPedido\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"GlnOrigen\",\"type\":[\"null\",\"string\"]},{\"name\":\"CantidadTotal\",\"type\":\"float\"},{\"name\":\"CantidadExpedidaTotal\",\"type\":\"float\"},{\"name\":\"CantidadPickeadaTotal\",\"type\":\"float\"},{\"default\":null,\"name\":\"DestinatarioDomicilioAdicional\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Mail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoIsoPais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"EstadoControlCalidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DescripcionTipoPedido\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LinkImpresionRemito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LinkImpresionEtiquetaPedido\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoANMAT\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ErrorOtTraza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ErrorOtAcondi\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ErrorValidaciones\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PropietarioEcommerce\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NovedadesAcondi\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroVale\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoEmbalaje\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroTaco\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"UrlQR\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroOrdenOriginalEcommerce\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NotasEmbalaje\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Propietario\",\"type\":[\"null\",\"string\"]}],\"name\":\"Cabecera\",\"namespace\":\"Andreani.EventoWhPedidos.Events.SCEPedidosEmpaquetadosCommon\",\"type\":\"record\"}},{\"name\":\"Detalle\",\"type\":{\"items\":{\"fields\":[{\"name\":\"SKU\",\"type\":\"string\"},{\"name\":\"LineaPedidoWH\",\"type\":\"string\"},{\"name\":\"LineaExterna\",\"type\":\"string\"},{\"name\":\"PaqueteLote\",\"type\":\"string\"},{\"name\":\"LoteCajitaFabricante\",\"type\":\"string\"},{\"name\":\"LoteSecundario\",\"type\":\"string\"},{\"default\":null,\"name\":\"FechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ProductoTrazable\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"EstadoLote\",\"type\":\"string\"},{\"name\":\"BloqueoUbicacion\",\"type\":\"string\"},{\"name\":\"VidaUtilLote\",\"type\":\"string\"},{\"default\":null,\"name\":\"EntregaAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"ConsumoAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"EsfuerzoLineaApiPlani\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"VidaUtil\",\"type\":[\"null\",\"string\"]},{\"name\":\"Cantidad\",\"type\":\"float\"},{\"name\":\"CantidadPickeada\",\"type\":\"float\"},{\"name\":\"CantidadExpedida\",\"type\":\"float\"},{\"name\":\"EstadoLineaCodigo\",\"type\":\"string\"},{\"name\":\"EstadoLineaDescripcion\",\"type\":\"string\"},{\"name\":\"CantidadBultos\",\"type\":\"string\"},{\"name\":\"Series\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Serie\",\"type\":\"string\"}],\"name\":\"Series\",\"type\":\"record\"},\"type\":\"array\"}},{\"default\":null,\"name\":\"Contrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Nomenclatura\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoPedido\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoPedidoDescripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DetalleAcondicionamiento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"EstadoOTAcondicionamiento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoFrio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigosAcondi\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NovedadesAcondi\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Propietario\",\"type\":[\"null\",\"string\"]}],\"name\":\"Detalle\",\"namespace\":\"Andreani.EventoWhPedidos.Events.SCEPedidosEmpaquetadosCommon\",\"type\":\"record\"},\"type\":\"array\"}},{\"name\":\"DetalleBultos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"OrdenWH\",\"type\":\"string\"},{\"name\":\"OrdenCliente\",\"type\":\"string\"},{\"default\":null,\"name\":\"Remito\",\"type\":[\"null\",\"string\"]},{\"name\":\"LineaPedidoWH\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"SKU\",\"type\":\"string\"},{\"name\":\"CantidadEmbalada\",\"type\":\"float\"},{\"name\":\"EtiquetaEmbalaje\",\"type\":\"string\"},{\"default\":null,\"name\":\"TipoDeContenedor\",\"type\":[\"null\",\"string\"]},{\"name\":\"PaqueteLote\",\"type\":\"string\"},{\"name\":\"LoteCajitaFabricante\",\"type\":\"string\"},{\"name\":\"LoteSecundario\",\"type\":\"string\"},{\"default\":null,\"name\":\"FechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ProductoTrazable\",\"type\":\"string\"},{\"name\":\"AlmacenConsumo\",\"type\":\"string\"},{\"name\":\"EstadoLote\",\"type\":\"string\"},{\"name\":\"BloqueoUbicacion\",\"type\":\"string\"},{\"name\":\"VidaUtilLote\",\"type\":\"string\"},{\"default\":null,\"name\":\"EntregaAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"ConsumoAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"DetalleBultos\",\"namespace\":\"Andreani.EventoWhPedidos.Events.SCEPedidosEmpaquetadosCommon\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.EventoWhPedidos.Events.Record.EventoWhSCEPedidosEmpaquetados\",\"type\":\"record\"}"
}

func (r EventoWhSCEPedidosEmpaquetados) SchemaName() string {
	return "Andreani.EventoWhPedidos.Events.Record.EventoWhSCEPedidosEmpaquetados"
}

func (_ EventoWhSCEPedidosEmpaquetados) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoWhSCEPedidosEmpaquetados) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoWhSCEPedidosEmpaquetados) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoWhSCEPedidosEmpaquetados) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoWhSCEPedidosEmpaquetados) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoWhSCEPedidosEmpaquetados) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoWhSCEPedidosEmpaquetados) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoWhSCEPedidosEmpaquetados) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoWhSCEPedidosEmpaquetados) Get(i int) types.Field {
	switch i {
	case 0:
		r.Identificacion = NewIdentificacion()

		w := types.Record{Target: &r.Identificacion}

		return w

	case 1:
		r.Cabecera = NewCabecera()

		w := types.Record{Target: &r.Cabecera}

		return w

	case 2:
		r.Detalle = make([]Detalle, 0)

		w := ArrayDetalleWrapper{Target: &r.Detalle}

		return w

	case 3:
		r.DetalleBultos = make([]DetalleBultos, 0)

		w := ArrayDetalleBultosWrapper{Target: &r.DetalleBultos}

		return w

	}
	panic("Unknown field index")
}

func (r *EventoWhSCEPedidosEmpaquetados) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoWhSCEPedidosEmpaquetados) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EventoWhSCEPedidosEmpaquetados) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ EventoWhSCEPedidosEmpaquetados) AppendArray() types.Field { panic("Unsupported operation") }
func (_ EventoWhSCEPedidosEmpaquetados) HintSize(int)             { panic("Unsupported operation") }
func (_ EventoWhSCEPedidosEmpaquetados) Finalize()                {}

func (_ EventoWhSCEPedidosEmpaquetados) AvroCRC64Fingerprint() []byte {
	return []byte(EventoWhSCEPedidosEmpaquetadosAvroCRC64Fingerprint)
}

func (r EventoWhSCEPedidosEmpaquetados) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Identificacion"], err = json.Marshal(r.Identificacion)
	if err != nil {
		return nil, err
	}
	output["Cabecera"], err = json.Marshal(r.Cabecera)
	if err != nil {
		return nil, err
	}
	output["Detalle"], err = json.Marshal(r.Detalle)
	if err != nil {
		return nil, err
	}
	output["DetalleBultos"], err = json.Marshal(r.DetalleBultos)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EventoWhSCEPedidosEmpaquetados) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Identificacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Identificacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Identificacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Cabecera"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cabecera); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Cabecera")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Detalle"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Detalle); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Detalle")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DetalleBultos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DetalleBultos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DetalleBultos")
	}
	return nil
}
