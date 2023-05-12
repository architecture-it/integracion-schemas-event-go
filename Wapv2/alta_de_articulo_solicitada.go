// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AltaDeArticuloSolicitada.avsc
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

type AltaDeArticuloSolicitada struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	ArticuloSCE ArticuloSCE `json:"articuloSCE"`

	Topic string `json:"Topic"`
}

const AltaDeArticuloSolicitadaAvroCRC64Fingerprint = "\x8b\xb1+d\v\x00\x84\x94"

func NewAltaDeArticuloSolicitada() AltaDeArticuloSolicitada {
	r := AltaDeArticuloSolicitada{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.ArticuloSCE = NewArticuloSCE()

	r.Topic = "Almacen/Solicitudes/AltaDeArticuloSolicitada"
	return r
}

func DeserializeAltaDeArticuloSolicitada(r io.Reader) (AltaDeArticuloSolicitada, error) {
	t := NewAltaDeArticuloSolicitada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAltaDeArticuloSolicitadaFromSchema(r io.Reader, schema string) (AltaDeArticuloSolicitada, error) {
	t := NewAltaDeArticuloSolicitada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAltaDeArticuloSolicitada(r AltaDeArticuloSolicitada, w io.Writer) error {
	var err error
	err = writeSolicitudDeAccionAlmacen(r.SolicitudDeAccionAlmacen, w)
	if err != nil {
		return err
	}
	err = writeArticuloSCE(r.ArticuloSCE, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Topic, w)
	if err != nil {
		return err
	}
	return err
}

func (r AltaDeArticuloSolicitada) Serialize(w io.Writer) error {
	return writeAltaDeArticuloSolicitada(r, w)
}

func (r AltaDeArticuloSolicitada) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"default\":null,\"name\":\"instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoDistribucion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"articuloSCE\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"ean\",\"type\":\"string\"},{\"name\":\"tipoEan\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"descripcion\",\"type\":\"string\"},{\"name\":\"datosLogisticos\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"volumen\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"pesoBruto\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"pesoTara\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"pesoNeto\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"cantidadporPaquete\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"cantidadporCaja\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"cantidadporPallet\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"alturaUnidad\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"largoUnidad\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"anchoUnidad\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"pesoPack\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"alturaPack\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"largoPack\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"anchoPack\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"pesoCaja\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"altoCaja\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"largoCaja\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"anchoCaja\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"pesoPallet\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"altoPallet\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"largoPallet\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"anchoPallet\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"nivelesporPallet\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"cajasporNivel\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"altura\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"longitud\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"anchura\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"volumencubico\",\"type\":[\"null\",\"float\"]}],\"name\":\"DatosLogisticos\",\"type\":\"record\"}},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"}]},{\"name\":\"notas\",\"type\":[\"null\",\"string\"]},{\"name\":\"claseDeArticulo\",\"type\":[\"null\",\"string\"]},{\"name\":\"paisDeOrigen\",\"type\":[\"null\",\"string\"]},{\"name\":\"esNumeroDeSerieDeEntradaUnico\",\"type\":[\"null\",\"boolean\"]},{\"name\":\"esNumeroDeSerieSalidaUnico\",\"type\":[\"null\",\"boolean\"]},{\"name\":\"instruccionesDePreparacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"vidaUtilSalidaEnDias\",\"type\":[\"null\",\"int\"]},{\"name\":\"codigoDeVidaUtil\",\"type\":[\"null\",\"string\"]},{\"name\":\"indicadorDeVidaUtil\",\"type\":[\"null\",\"string\"]},{\"name\":\"consumoEnDias\",\"type\":[\"null\",\"int\"]},{\"name\":\"vencimientoEnDias\",\"type\":[\"null\",\"int\"]},{\"name\":\"vidaUtilEntradaEnDias\",\"type\":[\"null\",\"int\"]},{\"name\":\"serializado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"grupos\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.Metadato\"]},{\"default\":null,\"name\":\"camposLibres\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.Metadato\"]},{\"name\":\"coleccion\",\"type\":[\"null\",\"string\"]},{\"name\":\"tema\",\"type\":[\"null\",\"string\"]},{\"name\":\"temporada\",\"type\":[\"null\",\"string\"]},{\"name\":\"estilo\",\"type\":[\"null\",\"string\"]},{\"name\":\"color\",\"type\":[\"null\",\"string\"]},{\"name\":\"iniciodeTemporada\",\"type\":[\"null\",\"string\"]},{\"name\":\"findeTemporada\",\"type\":[\"null\",\"string\"]},{\"name\":\"talle\",\"type\":[\"null\",\"string\"]},{\"name\":\"rubro\",\"type\":[\"null\",\"string\"]},{\"name\":\"pavu\",\"type\":[\"null\",\"string\"]},{\"name\":\"psicotropico\",\"type\":[\"null\",\"string\"]},{\"name\":\"temperatura\",\"type\":[\"null\",\"string\"]}],\"name\":\"ArticuloSCE\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/AltaDeArticuloSolicitada\",\"name\":\"Topic\",\"type\":\"string\"}],\"name\":\"Andreani.Wapv2.Events.Record.AltaDeArticuloSolicitada\",\"type\":\"record\"}"
}

func (r AltaDeArticuloSolicitada) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.AltaDeArticuloSolicitada"
}

func (_ AltaDeArticuloSolicitada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ AltaDeArticuloSolicitada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ AltaDeArticuloSolicitada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ AltaDeArticuloSolicitada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ AltaDeArticuloSolicitada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ AltaDeArticuloSolicitada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ AltaDeArticuloSolicitada) SetString(v string)   { panic("Unsupported operation") }
func (_ AltaDeArticuloSolicitada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AltaDeArticuloSolicitada) Get(i int) types.Field {
	switch i {
	case 0:
		r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

		w := types.Record{Target: &r.SolicitudDeAccionAlmacen}

		return w

	case 1:
		r.ArticuloSCE = NewArticuloSCE()

		w := types.Record{Target: &r.ArticuloSCE}

		return w

	case 2:
		w := types.String{Target: &r.Topic}

		return w

	}
	panic("Unknown field index")
}

func (r *AltaDeArticuloSolicitada) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/AltaDeArticuloSolicitada"
		return
	}
	panic("Unknown field index")
}

func (r *AltaDeArticuloSolicitada) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ AltaDeArticuloSolicitada) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ AltaDeArticuloSolicitada) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ AltaDeArticuloSolicitada) HintSize(int)                     { panic("Unsupported operation") }
func (_ AltaDeArticuloSolicitada) Finalize()                        {}

func (_ AltaDeArticuloSolicitada) AvroCRC64Fingerprint() []byte {
	return []byte(AltaDeArticuloSolicitadaAvroCRC64Fingerprint)
}

func (r AltaDeArticuloSolicitada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["solicitudDeAccionAlmacen"], err = json.Marshal(r.SolicitudDeAccionAlmacen)
	if err != nil {
		return nil, err
	}
	output["articuloSCE"], err = json.Marshal(r.ArticuloSCE)
	if err != nil {
		return nil, err
	}
	output["Topic"], err = json.Marshal(r.Topic)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *AltaDeArticuloSolicitada) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["articuloSCE"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ArticuloSCE); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for articuloSCE")
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
		r.Topic = "Almacen/Solicitudes/AltaDeArticuloSolicitada"
	}
	return nil
}
