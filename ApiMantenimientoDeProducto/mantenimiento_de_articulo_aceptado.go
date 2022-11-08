// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     MantenimientoDeArticuloAceptado.avsc
 */
package ApiMantenimientoDeProductoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type MantenimientoDeArticuloAceptado struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	DetalleDeArticulo DetalleDeArticulo `json:"detalleDeArticulo"`

	Topic string `json:"Topic"`
}

const MantenimientoDeArticuloAceptadoAvroCRC64Fingerprint = "\x1e\x89b\xf9Z\xfd\xdd1"

func NewMantenimientoDeArticuloAceptado() MantenimientoDeArticuloAceptado {
	r := MantenimientoDeArticuloAceptado{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.DetalleDeArticulo = NewDetalleDeArticulo()

	r.Topic = "Almacen/Solicitudes/MantenimientoDeArticuloAceptado"
	return r
}

func DeserializeMantenimientoDeArticuloAceptado(r io.Reader) (MantenimientoDeArticuloAceptado, error) {
	t := NewMantenimientoDeArticuloAceptado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeMantenimientoDeArticuloAceptadoFromSchema(r io.Reader, schema string) (MantenimientoDeArticuloAceptado, error) {
	t := NewMantenimientoDeArticuloAceptado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeMantenimientoDeArticuloAceptado(r MantenimientoDeArticuloAceptado, w io.Writer) error {
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

func (r MantenimientoDeArticuloAceptado) Serialize(w io.Writer) error {
	return writeMantenimientoDeArticuloAceptado(r, w)
}

func (r MantenimientoDeArticuloAceptado) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"contrato\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"name\":\"planta\",\"type\":\"string\"}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"detalleDeArticulo\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"ean13\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"lote\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"loteDeFabricante\",\"type\":\"string\"},{\"name\":\"loteSecundario\",\"type\":\"string\"},{\"name\":\"fechaDeVencimiento\",\"type\":\"string\"},{\"name\":\"otrosDatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Lote\",\"type\":\"record\"}},{\"name\":\"otrosDatos\",\"type\":{\"items\":\"Andreani.ApiMantenimientoDeProducto.Events.Record.Metadato\",\"type\":\"array\"}},{\"name\":\"descripcion\",\"type\":\"string\"},{\"name\":\"claseDeExpedicion\",\"type\":\"string\"},{\"name\":\"claseDeArticulo\",\"type\":\"string\"},{\"name\":\"paisDeOrigen\",\"type\":\"string\"},{\"name\":\"esNumeroDeSerieDeEntradaUnico\",\"type\":\"boolean\"},{\"name\":\"requiereCapturaDatosEntrada\",\"type\":\"boolean\"},{\"name\":\"esNumeroDeSerieSalidaUnico\",\"type\":\"boolean\"},{\"name\":\"requiereCapturaDatosSalida\",\"type\":\"boolean\"},{\"name\":\"requierecapturaTotalNumSeries\",\"type\":\"boolean\"},{\"name\":\"caracteristicas\",\"type\":{\"items\":\"Andreani.ApiMantenimientoDeProducto.Events.Record.Metadato\",\"type\":\"array\"}},{\"name\":\"notas\",\"type\":\"string\"},{\"name\":\"instruccionesDePreparacion\",\"type\":\"string\"},{\"name\":\"vidaUtilEnDias\",\"type\":\"long\"},{\"name\":\"codigoDeVidaUtil\",\"type\":\"string\"},{\"name\":\"indicadorDeVidaUtil\",\"type\":\"string\"},{\"name\":\"consumoEnDias\",\"type\":\"long\"},{\"name\":\"vencimientoEnDias\",\"type\":\"long\"},{\"name\":\"vidaUtilEntradaEnDias\",\"type\":\"long\"},{\"name\":\"acondicionamientoSecundario\",\"type\":\"string\"},{\"name\":\"zonaRepo\",\"type\":\"string\"},{\"name\":\"grupos\",\"type\":{\"items\":\"Andreani.ApiMantenimientoDeProducto.Events.Record.Metadato\",\"type\":\"array\"}},{\"name\":\"volumen\",\"type\":\"double\"},{\"name\":\"pesoBruto\",\"type\":\"double\"},{\"name\":\"pesoTara\",\"type\":\"double\"},{\"name\":\"pesoNeto\",\"type\":\"double\"},{\"name\":\"camposLibres\",\"type\":{\"items\":\"Andreani.ApiMantenimientoDeProducto.Events.Record.Metadato\",\"type\":\"array\"}}],\"name\":\"DetalleDeArticulo\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/MantenimientoDeArticuloAceptado\",\"name\":\"Topic\",\"type\":\"string\"}],\"name\":\"Andreani.ApiMantenimientoDeProducto.Events.Record.MantenimientoDeArticuloAceptado\",\"type\":\"record\"}"
}

func (r MantenimientoDeArticuloAceptado) SchemaName() string {
	return "Andreani.ApiMantenimientoDeProducto.Events.Record.MantenimientoDeArticuloAceptado"
}

func (_ MantenimientoDeArticuloAceptado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloAceptado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloAceptado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloAceptado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloAceptado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloAceptado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloAceptado) SetString(v string)   { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloAceptado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *MantenimientoDeArticuloAceptado) Get(i int) types.Field {
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

func (r *MantenimientoDeArticuloAceptado) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/MantenimientoDeArticuloAceptado"
		return
	}
	panic("Unknown field index")
}

func (r *MantenimientoDeArticuloAceptado) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ MantenimientoDeArticuloAceptado) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ MantenimientoDeArticuloAceptado) AppendArray() types.Field { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloAceptado) HintSize(int)             { panic("Unsupported operation") }
func (_ MantenimientoDeArticuloAceptado) Finalize()                {}

func (_ MantenimientoDeArticuloAceptado) AvroCRC64Fingerprint() []byte {
	return []byte(MantenimientoDeArticuloAceptadoAvroCRC64Fingerprint)
}

func (r MantenimientoDeArticuloAceptado) MarshalJSON() ([]byte, error) {
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

func (r *MantenimientoDeArticuloAceptado) UnmarshalJSON(data []byte) error {
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
		r.Topic = "Almacen/Solicitudes/MantenimientoDeArticuloAceptado"
	}
	return nil
}
