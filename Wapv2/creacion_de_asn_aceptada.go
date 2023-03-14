// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CreacionDeAsnAceptada.avsc
 */
package Wapv2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type CreacionDeAsnAceptada struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	Abastecimiento Abastecimiento `json:"abastecimiento"`

	Topic string `json:"Topic"`
}

const CreacionDeAsnAceptadaAvroCRC64Fingerprint = "\x8fU\xa2O\xa1\x97̇"

func NewCreacionDeAsnAceptada() CreacionDeAsnAceptada {
	r := CreacionDeAsnAceptada{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.Abastecimiento = NewAbastecimiento()

	r.Topic = "Almacen/Solicitudes/CreacionDeAsnAceptada"
	return r
}

func DeserializeCreacionDeAsnAceptada(r io.Reader) (CreacionDeAsnAceptada, error) {
	t := NewCreacionDeAsnAceptada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCreacionDeAsnAceptadaFromSchema(r io.Reader, schema string) (CreacionDeAsnAceptada, error) {
	t := NewCreacionDeAsnAceptada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCreacionDeAsnAceptada(r CreacionDeAsnAceptada, w io.Writer) error {
	var err error
	err = writeSolicitudDeAccionAlmacen(r.SolicitudDeAccionAlmacen, w)
	if err != nil {
		return err
	}
	err = writeAbastecimiento(r.Abastecimiento, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Topic, w)
	if err != nil {
		return err
	}
	return err
}

func (r CreacionDeAsnAceptada) Serialize(w io.Writer) error {
	return writeCreacionDeAsnAceptada(r, w)
}

func (r CreacionDeAsnAceptada) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"default\":null,\"name\":\"instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoDistribucion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"abastecimiento\",\"type\":{\"fields\":[{\"name\":\"idAndreani\",\"type\":\"string\"},{\"name\":\"contratoWarehouse\",\"type\":\"string\"},{\"name\":\"numeroOrdenExterna\",\"type\":\"string\"},{\"default\":null,\"name\":\"fechaRecepcionEsperada\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"fechaOrdenExterna\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"codigoDeTrazaANMAT\",\"type\":[\"null\",\"string\"]},{\"name\":\"propietarioMercaderia\",\"type\":\"string\"},{\"name\":\"tipoAsn\",\"type\":\"string\"},{\"name\":\"tipoRecepcion\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDeLote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ordenCompraAsociada\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"camposAdicionales\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"name\":\"origen\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"idOrigen\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contacto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosAdicionales\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.ListaDePropiedades\"]}],\"name\":\"Origen\",\"type\":\"record\"}},{\"name\":\"transportista\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"idTransportista\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosPersonales\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idinternocliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"eMail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"tipo\",\"type\":[\"null\",\"string\"]},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeTelefonos\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"contacto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoDeDocumento\",\"type\":[\"null\",{\"name\":\"TipoDeDocumento\",\"symbols\":[\"undefined\",\"DNI\",\"CUIT\",\"CUIL\"],\"type\":\"enum\"}]}],\"name\":\"DatosPersonales\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"referencia\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"transportista\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contenedor\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"muelleRecepcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroCita\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroGuia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"gln\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contacto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"correoElectronico\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoISOpais\",\"type\":[\"null\",\"string\"]}],\"name\":\"Referencia\",\"type\":\"record\"}]}],\"name\":\"Transporte\",\"type\":\"record\"}},{\"name\":\"lineaLista\",\"type\":{\"fields\":[{\"name\":\"detalleLinea\",\"type\":{\"items\":{\"fields\":[{\"default\":null,\"name\":\"idAndreani\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroDeLinea\",\"type\":\"string\"},{\"default\":null,\"name\":\"numeroDeLineaWMS\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estatusOTdeTRazaporLPN\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroOrdenExterna\",\"type\":\"string\"},{\"name\":\"cantidadPedida\",\"type\":\"int\"},{\"name\":\"valorDeclarado\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoActualizacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"unidadMedida\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"contratoWhAbastecimiento\",\"type\":\"string\"},{\"default\":null,\"name\":\"tipoAcondicionamiento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoTraza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoControlCalidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"accionControlCalidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ubicacionPredeterminada\",\"type\":[\"null\",\"string\"]},{\"name\":\"almacen\",\"type\":\"string\"},{\"default\":null,\"name\":\"vidaUtilLote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"avisoContramuestra\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.ListaDePropiedades\"]},{\"name\":\"articulo\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"ean\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"camposLibres\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.Metadato\"]},{\"name\":\"lote\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaDeVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.ListaDePropiedades\"]}],\"name\":\"LoteAsn\",\"type\":\"record\"}},{\"default\":null,\"name\":\"ubicacionDestino\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"paqueteLote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"prodTrazable\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"consumirAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"entrarAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"nomenclaturaServicio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"descripcionServicio\",\"type\":[\"null\",\"string\"]}],\"name\":\"ArticuloAsn\",\"type\":\"record\"}}],\"name\":\"Linea\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeDetalleDeLinea\",\"type\":\"record\"}},{\"default\":null,\"name\":\"tiposDeAcondicionamientoSecundario\",\"type\":[\"null\",{\"fields\":[{\"name\":\"tiposDeAcondicionamientoSecundario\",\"type\":{\"items\":{\"fields\":[{\"name\":\"tipoDeAcondi\",\"type\":\"int\"}],\"name\":\"TiposDeAcondicionamientoSecundario\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeTiposDeAcondicionamientoSecundario\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"tiposDeControlDeCalidad\",\"type\":[\"null\",{\"fields\":[{\"name\":\"tiposDeControlDeCalidad\",\"type\":{\"items\":{\"fields\":[{\"name\":\"tipoDeControl\",\"type\":\"int\"}],\"name\":\"TiposDeControlDeCalidad\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeTiposDeControlDeCalidad\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"tiposDeAcciones\",\"type\":[\"null\",{\"fields\":[{\"name\":\"tiposDeAcciones\",\"type\":{\"items\":{\"fields\":[{\"name\":\"tipoDeAccion\",\"type\":\"int\"}],\"name\":\"TiposDeAcciones\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeTiposDeAcciones\",\"type\":\"record\"}]}],\"name\":\"Abastecimiento\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/CreacionDeAsnAceptada\",\"name\":\"Topic\",\"type\":\"string\"}],\"name\":\"Andreani.Wapv2.Events.Record.CreacionDeAsnAceptada\",\"type\":\"record\"}"
}

func (r CreacionDeAsnAceptada) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.CreacionDeAsnAceptada"
}

func (_ CreacionDeAsnAceptada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CreacionDeAsnAceptada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CreacionDeAsnAceptada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CreacionDeAsnAceptada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CreacionDeAsnAceptada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CreacionDeAsnAceptada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CreacionDeAsnAceptada) SetString(v string)   { panic("Unsupported operation") }
func (_ CreacionDeAsnAceptada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CreacionDeAsnAceptada) Get(i int) types.Field {
	switch i {
	case 0:
		r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

		w := types.Record{Target: &r.SolicitudDeAccionAlmacen}

		return w

	case 1:
		r.Abastecimiento = NewAbastecimiento()

		w := types.Record{Target: &r.Abastecimiento}

		return w

	case 2:
		w := types.String{Target: &r.Topic}

		return w

	}
	panic("Unknown field index")
}

func (r *CreacionDeAsnAceptada) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/CreacionDeAsnAceptada"
		return
	}
	panic("Unknown field index")
}

func (r *CreacionDeAsnAceptada) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ CreacionDeAsnAceptada) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CreacionDeAsnAceptada) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CreacionDeAsnAceptada) HintSize(int)                     { panic("Unsupported operation") }
func (_ CreacionDeAsnAceptada) Finalize()                        {}

func (_ CreacionDeAsnAceptada) AvroCRC64Fingerprint() []byte {
	return []byte(CreacionDeAsnAceptadaAvroCRC64Fingerprint)
}

func (r CreacionDeAsnAceptada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["solicitudDeAccionAlmacen"], err = json.Marshal(r.SolicitudDeAccionAlmacen)
	if err != nil {
		return nil, err
	}
	output["abastecimiento"], err = json.Marshal(r.Abastecimiento)
	if err != nil {
		return nil, err
	}
	output["Topic"], err = json.Marshal(r.Topic)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CreacionDeAsnAceptada) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["solicitudDeAccionAlmacen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SolicitudDeAccionAlmacen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for solicitudDeAccionAlmacen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["abastecimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Abastecimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for abastecimiento")
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
		r.Topic = "Almacen/Solicitudes/CreacionDeAsnAceptada"
	}
	return nil
}
