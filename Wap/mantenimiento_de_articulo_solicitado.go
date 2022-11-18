// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     MantenimientoDeArticuloSolicitado.avsc
 */
package WapEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type MantenimientoDeArticuloSolicitado struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	DetalleDeArticulo DetalleDeArticulo `json:"detalleDeArticulo"`

	Topic string `json:"Topic"`
}

const MantenimientoDeArticuloSolicitadoAvroCRC64Fingerprint = "\xfe\xed\xc9\r\x8e\tL\xce"

func NewMantenimientoDeArticuloSolicitado() MantenimientoDeArticuloSolicitado {
	r := MantenimientoDeArticuloSolicitado{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.DetalleDeArticulo = NewDetalleDeArticulo()

	r.Topic = "Almacen/Solicitudes/MantenimientoDeArticuloSolicitado"
	return r
}

func DeserializeMantenimientoDeArticuloSolicitado(r io.Reader) (MantenimientoDeArticuloSolicitado, error) {
	t := NewMantenimientoDeArticuloSolicitado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMantenimientoDeArticuloSolicitadoFromSchema(r io.Reader, schema string) (MantenimientoDeArticuloSolicitado, error) {
	t := NewMantenimientoDeArticuloSolicitado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMantenimientoDeArticuloSolicitado(r MantenimientoDeArticuloSolicitado, w io.Writer) error {
	var err error
	err = writeSolicitudDeAccionAlmacen(r.SolicitudDeAccionAlmacen, w)
	if err != nil {
		return err
	}
	err = writeDetalleDeArticulo(r.DetalleDeArticulo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Topic, w)
	if err != nil {
		return err
	}
	return err
}

func (r MantenimientoDeArticuloSolicitado) Serialize(w io.Writer) error {
	return writeMantenimientoDeArticuloSolicitado(r, w)
}

func (r MantenimientoDeArticuloSolicitado) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"contrato\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"name\":\"planta\",\"type\":\"string\"},{\"default\":null,\"name\":\"planta_instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoDistribucion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"detalleDeArticulo\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"ean13\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"lote\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaDeVencimiento\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]}],\"name\":\"LoteArticulo\",\"type\":\"record\"}},{\"name\":\"otrosDatos\",\"type\":{\"items\":\"Andreani.Wap.Events.Record.Metadato\",\"type\":\"array\"}},{\"name\":\"descripcion\",\"type\":\"string\"},{\"name\":\"claseDeExpedicion\",\"type\":\"string\"},{\"name\":\"claseDeArticulo\",\"type\":\"string\"},{\"name\":\"paisDeOrigen\",\"type\":\"string\"},{\"name\":\"esNumeroDeSerieDeEntradaUnico\",\"type\":\"boolean\"},{\"name\":\"requiereCapturaDatosEntrada\",\"type\":\"boolean\"},{\"name\":\"esNumeroDeSerieSalidaUnico\",\"type\":\"boolean\"},{\"name\":\"requiereCapturaDatosSalida\",\"type\":\"boolean\"},{\"name\":\"requierecapturaTotalNumSeries\",\"type\":\"boolean\"},{\"name\":\"caracteristicas\",\"type\":{\"items\":\"Andreani.Wap.Events.Record.Metadato\",\"type\":\"array\"}},{\"name\":\"notas\",\"type\":\"string\"},{\"name\":\"instruccionesDePreparacion\",\"type\":\"string\"},{\"name\":\"vidaUtilEnDias\",\"type\":\"long\"},{\"name\":\"codigoDeVidaUtil\",\"type\":\"string\"},{\"name\":\"indicadorDeVidaUtil\",\"type\":\"string\"},{\"name\":\"consumoEnDias\",\"type\":\"long\"},{\"name\":\"vencimientoEnDias\",\"type\":\"long\"},{\"name\":\"vidaUtilEntradaEnDias\",\"type\":\"long\"},{\"name\":\"acondicionamientoSecundario\",\"type\":\"string\"},{\"name\":\"zonaRepo\",\"type\":\"string\"},{\"name\":\"grupos\",\"type\":{\"items\":\"Andreani.Wap.Events.Record.Metadato\",\"type\":\"array\"}},{\"name\":\"volumen\",\"type\":\"double\"},{\"name\":\"pesoBruto\",\"type\":\"double\"},{\"name\":\"pesoTara\",\"type\":\"double\"},{\"name\":\"pesoNeto\",\"type\":\"double\"},{\"name\":\"camposLibres\",\"type\":{\"items\":\"Andreani.Wap.Events.Record.Metadato\",\"type\":\"array\"}}],\"name\":\"DetalleDeArticulo\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/MantenimientoDeArticuloSolicitado\",\"name\":\"Topic\",\"type\":\"string\"}],\"name\":\"Andreani.Wap.Events.Record.MantenimientoDeArticuloSolicitado\",\"type\":\"record\"}"
}

func (r MantenimientoDeArticuloSolicitado) SchemaName() string {
	return "Andreani.Wap.Events.Record.MantenimientoDeArticuloSolicitado"
}

func (_ MantenimientoDeArticuloSolicitado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloSolicitado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloSolicitado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloSolicitado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloSolicitado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloSolicitado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloSolicitado) SetString(v string)   { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloSolicitado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MantenimientoDeArticuloSolicitado) Get(i int) types.Field {
	switch i {
	case 0:
		r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

		w := types.Record{Target: &r.SolicitudDeAccionAlmacen}

		return w

	case 1:
		r.DetalleDeArticulo = NewDetalleDeArticulo()

		w := types.Record{Target: &r.DetalleDeArticulo}

		return w

	case 2:
		w := types.String{Target: &r.Topic}

		return w

	}
	panic("Unknown field index")
}

func (r *MantenimientoDeArticuloSolicitado) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/MantenimientoDeArticuloSolicitado"
		return
	}
	panic("Unknown field index")
}

func (r *MantenimientoDeArticuloSolicitado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ MantenimientoDeArticuloSolicitado) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ MantenimientoDeArticuloSolicitado) AppendArray() types.Field { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloSolicitado) HintSize(int)             { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloSolicitado) Finalize()                {}

func (_ MantenimientoDeArticuloSolicitado) AvroCRC64Fingerprint() []byte {
	return []byte(MantenimientoDeArticuloSolicitadoAvroCRC64Fingerprint)
}

func (r MantenimientoDeArticuloSolicitado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["solicitudDeAccionAlmacen"], err = json.Marshal(r.SolicitudDeAccionAlmacen)
	if err != nil {
		return nil, err
	}
	output["detalleDeArticulo"], err = json.Marshal(r.DetalleDeArticulo)
	if err != nil {
		return nil, err
	}
	output["Topic"], err = json.Marshal(r.Topic)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *MantenimientoDeArticuloSolicitado) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["detalleDeArticulo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DetalleDeArticulo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for detalleDeArticulo")
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
		r.Topic = "Almacen/Solicitudes/MantenimientoDeArticuloSolicitado"
	}
	return nil
}
