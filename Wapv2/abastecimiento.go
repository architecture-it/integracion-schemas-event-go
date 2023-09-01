// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CreacionDeAsnSolicitada.avsc
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

type Abastecimiento struct {
	IdAndreani string `json:"idAndreani"`

	ContratoWarehouse *UnionNullString `json:"contratoWarehouse"`

	MarketPlace *UnionNullString `json:"marketPlace"`

	NumeroOrdenExterna string `json:"numeroOrdenExterna"`

	FechaRecepcionEsperada *UnionNullLong `json:"fechaRecepcionEsperada"`

	FechaOrdenExterna *UnionNullLong `json:"fechaOrdenExterna"`

	PropietarioMercaderia string `json:"propietarioMercaderia"`

	TipoAsn string `json:"tipoAsn"`

	TipoRecepcion string `json:"tipoRecepcion"`

	OrdenCompraAsociada *UnionNullString `json:"ordenCompraAsociada"`

	CamposAdicionales *UnionNullListaDePropiedades `json:"camposAdicionales"`

	Origen Origen `json:"origen"`

	Transportista Transporte `json:"transportista"`

	LineaLista ListaDeDetalleDeLinea `json:"lineaLista"`
}

const AbastecimientoAvroCRC64Fingerprint = "\xcd\x04\xecX\xac\xe8-\xe7"

func NewAbastecimiento() Abastecimiento {
	r := Abastecimiento{}
	r.MarketPlace = nil
	r.FechaRecepcionEsperada = nil
	r.FechaOrdenExterna = nil
	r.OrdenCompraAsociada = nil
	r.CamposAdicionales = nil
	r.Origen = NewOrigen()

	r.Transportista = NewTransporte()

	r.LineaLista = NewListaDeDetalleDeLinea()

	return r
}

func DeserializeAbastecimiento(r io.Reader) (Abastecimiento, error) {
	t := NewAbastecimiento()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAbastecimientoFromSchema(r io.Reader, schema string) (Abastecimiento, error) {
	t := NewAbastecimiento()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAbastecimiento(r Abastecimiento, w io.Writer) error {
	var err error
	err = vm.WriteString(r.IdAndreani, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ContratoWarehouse, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.MarketPlace, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroOrdenExterna, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaRecepcionEsperada, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaOrdenExterna, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.PropietarioMercaderia, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TipoAsn, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TipoRecepcion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.OrdenCompraAsociada, w)
	if err != nil {
		return err
	}
	err = writeUnionNullListaDePropiedades(r.CamposAdicionales, w)
	if err != nil {
		return err
	}
	err = writeOrigen(r.Origen, w)
	if err != nil {
		return err
	}
	err = writeTransporte(r.Transportista, w)
	if err != nil {
		return err
	}
	err = writeListaDeDetalleDeLinea(r.LineaLista, w)
	if err != nil {
		return err
	}
	return err
}

func (r Abastecimiento) Serialize(w io.Writer) error {
	return writeAbastecimiento(r, w)
}

func (r Abastecimiento) Schema() string {
	return "{\"fields\":[{\"name\":\"idAndreani\",\"type\":\"string\"},{\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"marketPlace\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroOrdenExterna\",\"type\":\"string\"},{\"default\":null,\"name\":\"fechaRecepcionEsperada\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"fechaOrdenExterna\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"propietarioMercaderia\",\"type\":\"string\"},{\"name\":\"tipoAsn\",\"type\":\"string\"},{\"name\":\"tipoRecepcion\",\"type\":\"string\"},{\"default\":null,\"name\":\"ordenCompraAsociada\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"camposAdicionales\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"name\":\"origen\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"idOrigen\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contacto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosAdicionales\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.ListaDePropiedades\"]}],\"name\":\"Origen\",\"type\":\"record\"}},{\"name\":\"transportista\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"idTransportista\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosPersonales\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"numeroDeDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"idinternocliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"eMail\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"tipo\",\"type\":[\"null\",\"string\"]},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeTelefonos\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"contacto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoDeDocumento\",\"type\":[\"null\",{\"name\":\"TipoDeDocumento\",\"symbols\":[\"undefined\",\"DNI\",\"CUIT\",\"CUIL\"],\"type\":\"enum\"}]}],\"name\":\"DatosPersonales\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"referencia\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"transportista\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contenedor\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"muelleRecepcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroCita\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroGuia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"gln\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contacto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"correoElectronico\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoISOpais\",\"type\":[\"null\",\"string\"]}],\"name\":\"Referencia\",\"type\":\"record\"}]}],\"name\":\"Transporte\",\"type\":\"record\"}},{\"name\":\"lineaLista\",\"type\":{\"fields\":[{\"name\":\"detalleLinea\",\"type\":{\"items\":{\"fields\":[{\"default\":null,\"name\":\"idAndreani\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroDeLinea\",\"type\":\"string\"},{\"default\":null,\"name\":\"numeroDeLineaWMS\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estatusOTdeTRazaporLPN\",\"type\":[\"null\",\"string\"]},{\"name\":\"cantidadPedida\",\"type\":\"int\"},{\"name\":\"valorDeclarado\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoActualizacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"unidadMedida\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"contratoWhAbastecimiento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoAcondicionamiento\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"tipoTraza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoControlCalidad\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"accionControlCalidad\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"ubicacionPredeterminada\",\"type\":[\"null\",\"string\"]},{\"name\":\"almacen\",\"type\":\"string\"},{\"default\":null,\"name\":\"vidaUtilLote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"avisoContramuestra\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.ListaDePropiedades\"]},{\"name\":\"articulo\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"ean\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"camposLibres\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.Metadato\"]},{\"name\":\"lote\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaDeVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.ListaDePropiedades\"]}],\"name\":\"LoteAsn\",\"type\":\"record\"}},{\"default\":null,\"name\":\"ubicacionDestino\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"paqueteLote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"prodTrazable\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"consumirAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"entrarAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"nomenclaturaServicio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"descripcionServicio\",\"type\":[\"null\",\"string\"]}],\"name\":\"ArticuloAsn\",\"type\":\"record\"}},{\"default\":null,\"name\":\"estatusOTdeAcondporLPN\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoDeTrazaANMAT\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"productoTrazable\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"series\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]}],\"name\":\"Linea\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeDetalleDeLinea\",\"type\":\"record\"}}],\"name\":\"Andreani.Wapv2.Events.Record.Abastecimiento\",\"type\":\"record\"}"
}

func (r Abastecimiento) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.Abastecimiento"
}

func (_ Abastecimiento) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Abastecimiento) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Abastecimiento) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Abastecimiento) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Abastecimiento) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Abastecimiento) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Abastecimiento) SetString(v string)   { panic("Unsupported operation") }
func (_ Abastecimiento) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Abastecimiento) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.IdAndreani}

		return w

	case 1:
		r.ContratoWarehouse = NewUnionNullString()

		return r.ContratoWarehouse
	case 2:
		r.MarketPlace = NewUnionNullString()

		return r.MarketPlace
	case 3:
		w := types.String{Target: &r.NumeroOrdenExterna}

		return w

	case 4:
		r.FechaRecepcionEsperada = NewUnionNullLong()

		return r.FechaRecepcionEsperada
	case 5:
		r.FechaOrdenExterna = NewUnionNullLong()

		return r.FechaOrdenExterna
	case 6:
		w := types.String{Target: &r.PropietarioMercaderia}

		return w

	case 7:
		w := types.String{Target: &r.TipoAsn}

		return w

	case 8:
		w := types.String{Target: &r.TipoRecepcion}

		return w

	case 9:
		r.OrdenCompraAsociada = NewUnionNullString()

		return r.OrdenCompraAsociada
	case 10:
		r.CamposAdicionales = NewUnionNullListaDePropiedades()

		return r.CamposAdicionales
	case 11:
		r.Origen = NewOrigen()

		w := types.Record{Target: &r.Origen}

		return w

	case 12:
		r.Transportista = NewTransporte()

		w := types.Record{Target: &r.Transportista}

		return w

	case 13:
		r.LineaLista = NewListaDeDetalleDeLinea()

		w := types.Record{Target: &r.LineaLista}

		return w

	}
	panic("Unknown field index")
}

func (r *Abastecimiento) SetDefault(i int) {
	switch i {
	case 2:
		r.MarketPlace = nil
		return
	case 4:
		r.FechaRecepcionEsperada = nil
		return
	case 5:
		r.FechaOrdenExterna = nil
		return
	case 9:
		r.OrdenCompraAsociada = nil
		return
	case 10:
		r.CamposAdicionales = nil
		return
	}
	panic("Unknown field index")
}

func (r *Abastecimiento) NullField(i int) {
	switch i {
	case 1:
		r.ContratoWarehouse = nil
		return
	case 2:
		r.MarketPlace = nil
		return
	case 4:
		r.FechaRecepcionEsperada = nil
		return
	case 5:
		r.FechaOrdenExterna = nil
		return
	case 9:
		r.OrdenCompraAsociada = nil
		return
	case 10:
		r.CamposAdicionales = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Abastecimiento) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Abastecimiento) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Abastecimiento) HintSize(int)                     { panic("Unsupported operation") }
func (_ Abastecimiento) Finalize()                        {}

func (_ Abastecimiento) AvroCRC64Fingerprint() []byte {
	return []byte(AbastecimientoAvroCRC64Fingerprint)
}

func (r Abastecimiento) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["idAndreani"], err = json.Marshal(r.IdAndreani)
	if err != nil {
		return nil, err
	}
	output["contratoWarehouse"], err = json.Marshal(r.ContratoWarehouse)
	if err != nil {
		return nil, err
	}
	output["marketPlace"], err = json.Marshal(r.MarketPlace)
	if err != nil {
		return nil, err
	}
	output["numeroOrdenExterna"], err = json.Marshal(r.NumeroOrdenExterna)
	if err != nil {
		return nil, err
	}
	output["fechaRecepcionEsperada"], err = json.Marshal(r.FechaRecepcionEsperada)
	if err != nil {
		return nil, err
	}
	output["fechaOrdenExterna"], err = json.Marshal(r.FechaOrdenExterna)
	if err != nil {
		return nil, err
	}
	output["propietarioMercaderia"], err = json.Marshal(r.PropietarioMercaderia)
	if err != nil {
		return nil, err
	}
	output["tipoAsn"], err = json.Marshal(r.TipoAsn)
	if err != nil {
		return nil, err
	}
	output["tipoRecepcion"], err = json.Marshal(r.TipoRecepcion)
	if err != nil {
		return nil, err
	}
	output["ordenCompraAsociada"], err = json.Marshal(r.OrdenCompraAsociada)
	if err != nil {
		return nil, err
	}
	output["camposAdicionales"], err = json.Marshal(r.CamposAdicionales)
	if err != nil {
		return nil, err
	}
	output["origen"], err = json.Marshal(r.Origen)
	if err != nil {
		return nil, err
	}
	output["transportista"], err = json.Marshal(r.Transportista)
	if err != nil {
		return nil, err
	}
	output["lineaLista"], err = json.Marshal(r.LineaLista)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Abastecimiento) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["idAndreani"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdAndreani); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idAndreani")
	}
	val = func() json.RawMessage {
		if v, ok := fields["contratoWarehouse"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContratoWarehouse); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for contratoWarehouse")
	}
	val = func() json.RawMessage {
		if v, ok := fields["marketPlace"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MarketPlace); err != nil {
			return err
		}
	} else {
		r.MarketPlace = NewUnionNullString()

		r.MarketPlace = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroOrdenExterna"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroOrdenExterna); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroOrdenExterna")
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaRecepcionEsperada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaRecepcionEsperada); err != nil {
			return err
		}
	} else {
		r.FechaRecepcionEsperada = NewUnionNullLong()

		r.FechaRecepcionEsperada = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaOrdenExterna"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaOrdenExterna); err != nil {
			return err
		}
	} else {
		r.FechaOrdenExterna = NewUnionNullLong()

		r.FechaOrdenExterna = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["propietarioMercaderia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PropietarioMercaderia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for propietarioMercaderia")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoAsn"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoAsn); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tipoAsn")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoRecepcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoRecepcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tipoRecepcion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ordenCompraAsociada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenCompraAsociada); err != nil {
			return err
		}
	} else {
		r.OrdenCompraAsociada = NewUnionNullString()

		r.OrdenCompraAsociada = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["camposAdicionales"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CamposAdicionales); err != nil {
			return err
		}
	} else {
		r.CamposAdicionales = NewUnionNullListaDePropiedades()

		r.CamposAdicionales = nil
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
		if v, ok := fields["transportista"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Transportista); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for transportista")
	}
	val = func() json.RawMessage {
		if v, ok := fields["lineaLista"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LineaLista); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for lineaLista")
	}
	return nil
}
