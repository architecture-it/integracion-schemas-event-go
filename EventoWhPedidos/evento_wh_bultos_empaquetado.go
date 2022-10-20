// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhBultosEmpaquetado.avsc
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

type EventoWhBultosEmpaquetado struct {
	Identificacion Identificacion `json:"Identificacion"`

	Cabecera Cabecera `json:"Cabecera"`

	Detalle []Detalle `json:"Detalle"`
}

const EventoWhBultosEmpaquetadoAvroCRC64Fingerprint = "hXb\xc6!VoI"

func NewEventoWhBultosEmpaquetado() EventoWhBultosEmpaquetado {
	r := EventoWhBultosEmpaquetado{}
	r.Identificacion = NewIdentificacion()

	r.Cabecera = NewCabecera()

	r.Detalle = make([]Detalle, 0)

	return r
}

func DeserializeEventoWhBultosEmpaquetado(r io.Reader) (EventoWhBultosEmpaquetado, error) {
	t := NewEventoWhBultosEmpaquetado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoWhBultosEmpaquetadoFromSchema(r io.Reader, schema string) (EventoWhBultosEmpaquetado, error) {
	t := NewEventoWhBultosEmpaquetado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoWhBultosEmpaquetado(r EventoWhBultosEmpaquetado, w io.Writer) error {
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
	return err
}

func (r EventoWhBultosEmpaquetado) Serialize(w io.Writer) error {
	return writeEventoWhBultosEmpaquetado(r, w)
}

func (r EventoWhBultosEmpaquetado) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.EventoWhPedidos.Events.EmpaquetadoCommon\",\"type\":\"record\"}},{\"name\":\"Cabecera\",\"type\":{\"fields\":[{\"name\":\"OrdenWH\",\"type\":\"string\"},{\"name\":\"OrdenCliente\",\"type\":\"string\"},{\"name\":\"Remito\",\"type\":[\"null\",\"string\"]},{\"name\":\"CodigoDestinatario\",\"type\":\"string\"},{\"name\":\"Acondicionamiento\",\"type\":[\"null\",\"string\"]},{\"name\":\"AlmacenCliente\",\"type\":[\"null\",\"string\"]},{\"name\":\"PrioridadApiPlani\",\"type\":[\"null\",\"string\"]},{\"name\":\"CuentaTMS\",\"type\":[\"null\",\"string\"]},{\"name\":\"ContratoTMS\",\"type\":[\"null\",\"string\"]},{\"name\":\"FacturaLegal\",\"type\":[\"null\",\"string\"]},{\"name\":\"FacturaInterna\",\"type\":[\"null\",\"string\"]},{\"name\":\"TieneGestionCobranza\",\"type\":[\"null\",\"string\"]},{\"name\":\"ValorSeguro\",\"type\":[\"null\",\"string\"]},{\"name\":\"AdmiteCambioLoteDirigido\",\"type\":[\"null\",\"string\"]},{\"name\":\"AdmitePickingParcial\",\"type\":[\"null\",\"string\"]},{\"name\":\"ImprimeDocumentacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"NumeroCale\",\"type\":[\"null\",\"string\"]},{\"name\":\"Cot\",\"type\":[\"null\",\"string\"]},{\"name\":\"ComprobanteIngresosBrutos\",\"type\":[\"null\",\"string\"]},{\"name\":\"ContratoServicioAlmacenes\",\"type\":[\"null\",\"string\"]},{\"name\":\"Marketplace\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoPedidoMatriz\",\"type\":[\"null\",\"string\"]},{\"name\":\"Contacto\",\"type\":[\"null\",\"string\"]},{\"name\":\"Destinario\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioCalle\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinarioNumero\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioPiso\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioDepartamento\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioGLNDNI\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioCiudad\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioProvincia\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioCodigoPostal\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioTelefono\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioEmail\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenCalle\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenNumero\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenPiso\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenDepartamento\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenEmail\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrigenTelefono\",\"type\":[\"null\",\"string\"]},{\"name\":\"RemitenteCalle\",\"type\":[\"null\",\"string\"]},{\"name\":\"RemitenteNumero\",\"type\":[\"null\",\"string\"]},{\"name\":\"RemitentePiso\",\"type\":[\"null\",\"string\"]},{\"name\":\"RemitenteDepartamento\",\"type\":[\"null\",\"string\"]},{\"name\":\"RemitenteEmail\",\"type\":[\"null\",\"string\"]},{\"name\":\"RemitenteTelefono\",\"type\":[\"null\",\"string\"]},{\"name\":\"RemitenteCompra\",\"type\":[\"null\",\"string\"]},{\"name\":\"RemitenteDni\",\"type\":[\"null\",\"string\"]},{\"name\":\"RemitenteTipo\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrdenCompra\",\"type\":[\"null\",\"string\"]},{\"name\":\"ClientePadre\",\"type\":\"string\"},{\"name\":\"CodigoEstado\",\"type\":\"string\"},{\"name\":\"CodigoDescripcion\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoPedidoCodigo\",\"type\":\"string\"},{\"name\":\"TipoPedidoDescripcion\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoGrilla\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaEnvioFacturar\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaFacturacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"PrecioValorFC\",\"type\":[\"null\",\"string\"]},{\"name\":\"NumeroEnvio\",\"type\":[\"null\",\"string\"]},{\"name\":\"NotasPedido\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaCita\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaEntrega\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaExpedicionSolicitada\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaExpedicionProgramada\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaExpedicionPrometida\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaEntregaPlanificada\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaEntregaProgramada\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"GrupoOrdenes\",\"type\":\"string\"},{\"name\":\"FranjaHorario\",\"type\":[\"null\",\"string\"]},{\"name\":\"EstadoOTAcondi\",\"type\":[\"null\",\"string\"]},{\"name\":\"EstadoOTTraz\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaEvento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"CantidadTotal\",\"type\":\"float\"},{\"name\":\"CantidadExpedidaTotal\",\"type\":\"float\"},{\"name\":\"CantidadPickeadaTotal\",\"type\":\"float\"}],\"name\":\"Cabecera\",\"namespace\":\"Andreani.EventoWhPedidos.Events.EmpaquetadoCommon\",\"type\":\"record\"}},{\"name\":\"Detalle\",\"type\":{\"items\":{\"fields\":[{\"name\":\"SKU\",\"type\":\"string\"},{\"name\":\"LineaPedidoWH\",\"type\":\"string\"},{\"name\":\"LineaExterna\",\"type\":\"string\"},{\"name\":\"PaqueteLote\",\"type\":\"string\"},{\"name\":\"LoteCajitaFabricante\",\"type\":\"string\"},{\"name\":\"LoteSecundario\",\"type\":\"string\"},{\"name\":\"FechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ProductoTrazable\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"EstadoLote\",\"type\":\"string\"},{\"name\":\"BloqueoUbicacion\",\"type\":\"string\"},{\"name\":\"VitaUtilLote\",\"type\":\"string\"},{\"name\":\"EntregaAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ConsumoAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"EsfuerzoLineaApiPlani\",\"type\":[\"null\",\"string\"]},{\"name\":\"VidaUtil\",\"type\":[\"null\",\"string\"]},{\"name\":\"Cantidad\",\"type\":\"float\"},{\"name\":\"CantidadPickeada\",\"type\":\"float\"},{\"name\":\"CantidadExpedida\",\"type\":\"float\"},{\"name\":\"EstadoLineaCodigo\",\"type\":\"string\"},{\"name\":\"EstadoLineaDescripcion\",\"type\":\"string\"},{\"name\":\"CantidadBultos\",\"type\":\"string\"},{\"name\":\"Series\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Serie\",\"type\":\"string\"}],\"name\":\"Series\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Detalle\",\"namespace\":\"Andreani.EventoWhPedidos.Events.EmpaquetadoCommon\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.EventoWhPedidos.Events.Record.EventoWhBultosEmpaquetado\",\"type\":\"record\"}"
}

func (r EventoWhBultosEmpaquetado) SchemaName() string {
	return "Andreani.EventoWhPedidos.Events.Record.EventoWhBultosEmpaquetado"
}

func (_ EventoWhBultosEmpaquetado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoWhBultosEmpaquetado) Get(i int) types.Field {
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

	}
	panic("Unknown field index")
}

func (r *EventoWhBultosEmpaquetado) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoWhBultosEmpaquetado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EventoWhBultosEmpaquetado) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) HintSize(int)                     { panic("Unsupported operation") }
func (_ EventoWhBultosEmpaquetado) Finalize()                        {}

func (_ EventoWhBultosEmpaquetado) AvroCRC64Fingerprint() []byte {
	return []byte(EventoWhBultosEmpaquetadoAvroCRC64Fingerprint)
}

func (r EventoWhBultosEmpaquetado) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(output)
}

func (r *EventoWhBultosEmpaquetado) UnmarshalJSON(data []byte) error {
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
	return nil
}
