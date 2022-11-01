// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     MantenimientoDeArticuloSolicitado.avsc
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

type MantenimientoDeArticuloSolicitado struct {
	Contrato string `json:"contrato"`

	Almacen string `json:"almacen"`

	AlmacenSAP *UnionNullString `json:"almacenSAP"`

	Planta string `json:"planta"`

	DetalleDeArticulo DetalleDeArticulo `json:"detalleDeArticulo"`
}

const MantenimientoDeArticuloSolicitadoAvroCRC64Fingerprint = "\xbdڑ\xf0\x18CL\xd2"

func NewMantenimientoDeArticuloSolicitado() MantenimientoDeArticuloSolicitado {
	r := MantenimientoDeArticuloSolicitado{}
	r.AlmacenSAP = nil
	r.DetalleDeArticulo = NewDetalleDeArticulo()

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
	err = vm.WriteString(r.Contrato, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Almacen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.AlmacenSAP, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Planta, w)
	if err != nil {
		return err
	}
	err = writeDetalleDeArticulo(r.DetalleDeArticulo, w)
	if err != nil {
		return err
	}
	return err
}

func (r MantenimientoDeArticuloSolicitado) Serialize(w io.Writer) error {
	return writeMantenimientoDeArticuloSolicitado(r, w)
}

func (r MantenimientoDeArticuloSolicitado) Schema() string {
	return "{\"fields\":[{\"name\":\"contrato\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"default\":null,\"name\":\"almacenSAP\",\"type\":[\"null\",\"string\"]},{\"name\":\"planta\",\"type\":\"string\"},{\"name\":\"detalleDeArticulo\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"ean13\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"lote\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"loteDeFabricante\",\"type\":\"string\"},{\"name\":\"loteSecundario\",\"type\":\"string\"},{\"name\":\"fechaDeVencimiento\",\"type\":\"string\"},{\"name\":\"otrosDatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Lote\",\"type\":\"record\"}},{\"name\":\"otrosDatos\",\"type\":{\"items\":\"Andreani.ApiMantenimientoDeProducto.Events.Record.Metadato\",\"type\":\"array\"}},{\"name\":\"descripcion\",\"type\":\"string\"},{\"name\":\"claseDeExpedicion\",\"type\":\"string\"},{\"name\":\"claseDeArticulo\",\"type\":\"string\"},{\"name\":\"paisDeOrigen\",\"type\":\"string\"},{\"name\":\"esNumeroDeSerieDeEntradaUnico\",\"type\":\"boolean\"},{\"name\":\"requiereCapturaDatosEntrada\",\"type\":\"boolean\"},{\"name\":\"esNumeroDeSerieSalidaUnico\",\"type\":\"boolean\"},{\"name\":\"requiereCapturaDatosSalida\",\"type\":\"boolean\"},{\"name\":\"requierecapturaTotalNumSeries\",\"type\":\"boolean\"},{\"name\":\"caracteristicas\",\"type\":{\"items\":\"Andreani.ApiMantenimientoDeProducto.Events.Record.Metadato\",\"type\":\"array\"}},{\"name\":\"notas\",\"type\":\"string\"},{\"name\":\"instruccionesDePreparacion\",\"type\":\"string\"},{\"name\":\"vidaUtilEnDias\",\"type\":\"long\"},{\"name\":\"codigoDeVidaUtil\",\"type\":\"string\"},{\"name\":\"indicadorDeVidaUtil\",\"type\":\"string\"},{\"name\":\"consumoEnDias\",\"type\":\"long\"},{\"name\":\"vencimientoEnDias\",\"type\":\"long\"},{\"name\":\"vidaUtilEntradaEnDias\",\"type\":\"long\"},{\"name\":\"acondicionamientoSecundario\",\"type\":\"string\"},{\"name\":\"zonaRepo\",\"type\":\"string\"},{\"name\":\"grupos\",\"type\":{\"items\":\"Andreani.ApiMantenimientoDeProducto.Events.Record.Metadato\",\"type\":\"array\"}},{\"name\":\"volumen\",\"type\":\"double\"},{\"name\":\"pesoBruto\",\"type\":\"double\"},{\"name\":\"pesoNeto\",\"type\":\"double\"},{\"name\":\"camposLibres\",\"type\":{\"items\":\"Andreani.ApiMantenimientoDeProducto.Events.Record.Metadato\",\"type\":\"array\"}}],\"name\":\"DetalleDeArticulo\",\"type\":\"record\"}}],\"name\":\"Andreani.ApiMantenimientoDeProducto.Events.Record.MantenimientoDeArticuloSolicitado\",\"type\":\"record\"}"
}

func (r MantenimientoDeArticuloSolicitado) SchemaName() string {
	return "Andreani.ApiMantenimientoDeProducto.Events.Record.MantenimientoDeArticuloSolicitado"
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
		w := types.String{Target: &r.Contrato}

		return w

	case 1:
		w := types.String{Target: &r.Almacen}

		return w

	case 2:
		r.AlmacenSAP = NewUnionNullString()

		return r.AlmacenSAP
	case 3:
		w := types.String{Target: &r.Planta}

		return w

	case 4:
		r.DetalleDeArticulo = NewDetalleDeArticulo()

		w := types.Record{Target: &r.DetalleDeArticulo}

		return w

	}
	panic("Unknown field index")
}

func (r *MantenimientoDeArticuloSolicitado) SetDefault(i int) {
	switch i {
	case 2:
		r.AlmacenSAP = nil
		return
	}
	panic("Unknown field index")
}

func (r *MantenimientoDeArticuloSolicitado) NullField(i int) {
	switch i {
	case 2:
		r.AlmacenSAP = nil
		return
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
	output["contrato"], err = json.Marshal(r.Contrato)
	if err != nil {
		return nil, err
	}
	output["almacen"], err = json.Marshal(r.Almacen)
	if err != nil {
		return nil, err
	}
	output["almacenSAP"], err = json.Marshal(r.AlmacenSAP)
	if err != nil {
		return nil, err
	}
	output["planta"], err = json.Marshal(r.Planta)
	if err != nil {
		return nil, err
	}
	output["detalleDeArticulo"], err = json.Marshal(r.DetalleDeArticulo)
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
		return fmt.Errorf("no value specified for contrato")
	}
	val = func() json.RawMessage {
		if v, ok := fields["almacen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Almacen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for almacen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["almacenSAP"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AlmacenSAP); err != nil {
			return err
		}
	} else {
		r.AlmacenSAP = NewUnionNullString()

		r.AlmacenSAP = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["planta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Planta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for planta")
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
	return nil
}
