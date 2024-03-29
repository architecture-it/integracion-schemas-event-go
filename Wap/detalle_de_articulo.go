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

type DetalleDeArticulo struct {
	Codigo string `json:"codigo"`

	Ean13 string `json:"ean13"`

	Propietario string `json:"propietario"`

	Lote LoteArticulo `json:"lote"`

	OtrosDatos []Metadato `json:"otrosDatos"`

	Descripcion string `json:"descripcion"`

	ClaseDeExpedicion string `json:"claseDeExpedicion"`

	ClaseDeArticulo string `json:"claseDeArticulo"`

	PaisDeOrigen string `json:"paisDeOrigen"`

	EsNumeroDeSerieDeEntradaUnico bool `json:"esNumeroDeSerieDeEntradaUnico"`

	RequiereCapturaDatosEntrada bool `json:"requiereCapturaDatosEntrada"`

	EsNumeroDeSerieSalidaUnico bool `json:"esNumeroDeSerieSalidaUnico"`

	RequiereCapturaDatosSalida bool `json:"requiereCapturaDatosSalida"`

	RequierecapturaTotalNumSeries bool `json:"requierecapturaTotalNumSeries"`

	Caracteristicas []Metadato `json:"caracteristicas"`

	Notas string `json:"notas"`

	InstruccionesDePreparacion string `json:"instruccionesDePreparacion"`

	VidaUtilEnDias int64 `json:"vidaUtilEnDias"`

	CodigoDeVidaUtil string `json:"codigoDeVidaUtil"`

	IndicadorDeVidaUtil string `json:"indicadorDeVidaUtil"`

	ConsumoEnDias int64 `json:"consumoEnDias"`

	VencimientoEnDias int64 `json:"vencimientoEnDias"`

	VidaUtilEntradaEnDias int64 `json:"vidaUtilEntradaEnDias"`

	AcondicionamientoSecundario string `json:"acondicionamientoSecundario"`

	ZonaRepo string `json:"zonaRepo"`

	Grupos []Metadato `json:"grupos"`

	Volumen float64 `json:"volumen"`

	PesoBruto float64 `json:"pesoBruto"`

	PesoTara float64 `json:"pesoTara"`

	PesoNeto float64 `json:"pesoNeto"`

	CamposLibres []Metadato `json:"camposLibres"`
}

const DetalleDeArticuloAvroCRC64Fingerprint = "\xfb+b\x1eu7P\xdd"

func NewDetalleDeArticulo() DetalleDeArticulo {
	r := DetalleDeArticulo{}
	r.Lote = NewLoteArticulo()

	r.OtrosDatos = make([]Metadato, 0)

	r.Caracteristicas = make([]Metadato, 0)

	r.Grupos = make([]Metadato, 0)

	r.CamposLibres = make([]Metadato, 0)

	return r
}

func DeserializeDetalleDeArticulo(r io.Reader) (DetalleDeArticulo, error) {
	t := NewDetalleDeArticulo()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDetalleDeArticuloFromSchema(r io.Reader, schema string) (DetalleDeArticulo, error) {
	t := NewDetalleDeArticulo()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDetalleDeArticulo(r DetalleDeArticulo, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Codigo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Ean13, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = writeLoteArticulo(r.Lote, w)
	if err != nil {
		return err
	}
	err = writeArrayMetadato(r.OtrosDatos, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Descripcion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ClaseDeExpedicion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ClaseDeArticulo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.PaisDeOrigen, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.EsNumeroDeSerieDeEntradaUnico, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.RequiereCapturaDatosEntrada, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.EsNumeroDeSerieSalidaUnico, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.RequiereCapturaDatosSalida, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.RequierecapturaTotalNumSeries, w)
	if err != nil {
		return err
	}
	err = writeArrayMetadato(r.Caracteristicas, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Notas, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.InstruccionesDePreparacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.VidaUtilEnDias, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoDeVidaUtil, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.IndicadorDeVidaUtil, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.ConsumoEnDias, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.VencimientoEnDias, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.VidaUtilEntradaEnDias, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.AcondicionamientoSecundario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ZonaRepo, w)
	if err != nil {
		return err
	}
	err = writeArrayMetadato(r.Grupos, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.Volumen, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.PesoBruto, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.PesoTara, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.PesoNeto, w)
	if err != nil {
		return err
	}
	err = writeArrayMetadato(r.CamposLibres, w)
	if err != nil {
		return err
	}
	return err
}

func (r DetalleDeArticulo) Serialize(w io.Writer) error {
	return writeDetalleDeArticulo(r, w)
}

func (r DetalleDeArticulo) Schema() string {
	return "{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"ean13\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"lote\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaDeVencimiento\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]}],\"name\":\"LoteArticulo\",\"type\":\"record\"}},{\"name\":\"otrosDatos\",\"type\":{\"items\":\"Andreani.Wap.Events.Record.Metadato\",\"type\":\"array\"}},{\"name\":\"descripcion\",\"type\":\"string\"},{\"name\":\"claseDeExpedicion\",\"type\":\"string\"},{\"name\":\"claseDeArticulo\",\"type\":\"string\"},{\"name\":\"paisDeOrigen\",\"type\":\"string\"},{\"name\":\"esNumeroDeSerieDeEntradaUnico\",\"type\":\"boolean\"},{\"name\":\"requiereCapturaDatosEntrada\",\"type\":\"boolean\"},{\"name\":\"esNumeroDeSerieSalidaUnico\",\"type\":\"boolean\"},{\"name\":\"requiereCapturaDatosSalida\",\"type\":\"boolean\"},{\"name\":\"requierecapturaTotalNumSeries\",\"type\":\"boolean\"},{\"name\":\"caracteristicas\",\"type\":{\"items\":\"Andreani.Wap.Events.Record.Metadato\",\"type\":\"array\"}},{\"name\":\"notas\",\"type\":\"string\"},{\"name\":\"instruccionesDePreparacion\",\"type\":\"string\"},{\"name\":\"vidaUtilEnDias\",\"type\":\"long\"},{\"name\":\"codigoDeVidaUtil\",\"type\":\"string\"},{\"name\":\"indicadorDeVidaUtil\",\"type\":\"string\"},{\"name\":\"consumoEnDias\",\"type\":\"long\"},{\"name\":\"vencimientoEnDias\",\"type\":\"long\"},{\"name\":\"vidaUtilEntradaEnDias\",\"type\":\"long\"},{\"name\":\"acondicionamientoSecundario\",\"type\":\"string\"},{\"name\":\"zonaRepo\",\"type\":\"string\"},{\"name\":\"grupos\",\"type\":{\"items\":\"Andreani.Wap.Events.Record.Metadato\",\"type\":\"array\"}},{\"name\":\"volumen\",\"type\":\"double\"},{\"name\":\"pesoBruto\",\"type\":\"double\"},{\"name\":\"pesoTara\",\"type\":\"double\"},{\"name\":\"pesoNeto\",\"type\":\"double\"},{\"name\":\"camposLibres\",\"type\":{\"items\":\"Andreani.Wap.Events.Record.Metadato\",\"type\":\"array\"}}],\"name\":\"Andreani.Wap.Events.Record.DetalleDeArticulo\",\"type\":\"record\"}"
}

func (r DetalleDeArticulo) SchemaName() string {
	return "Andreani.Wap.Events.Record.DetalleDeArticulo"
}

func (_ DetalleDeArticulo) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DetalleDeArticulo) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DetalleDeArticulo) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DetalleDeArticulo) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DetalleDeArticulo) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DetalleDeArticulo) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DetalleDeArticulo) SetString(v string)   { panic("Unsupported operation") }
func (_ DetalleDeArticulo) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DetalleDeArticulo) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Codigo}

		return w

	case 1:
		w := types.String{Target: &r.Ean13}

		return w

	case 2:
		w := types.String{Target: &r.Propietario}

		return w

	case 3:
		r.Lote = NewLoteArticulo()

		w := types.Record{Target: &r.Lote}

		return w

	case 4:
		r.OtrosDatos = make([]Metadato, 0)

		w := ArrayMetadatoWrapper{Target: &r.OtrosDatos}

		return w

	case 5:
		w := types.String{Target: &r.Descripcion}

		return w

	case 6:
		w := types.String{Target: &r.ClaseDeExpedicion}

		return w

	case 7:
		w := types.String{Target: &r.ClaseDeArticulo}

		return w

	case 8:
		w := types.String{Target: &r.PaisDeOrigen}

		return w

	case 9:
		w := types.Boolean{Target: &r.EsNumeroDeSerieDeEntradaUnico}

		return w

	case 10:
		w := types.Boolean{Target: &r.RequiereCapturaDatosEntrada}

		return w

	case 11:
		w := types.Boolean{Target: &r.EsNumeroDeSerieSalidaUnico}

		return w

	case 12:
		w := types.Boolean{Target: &r.RequiereCapturaDatosSalida}

		return w

	case 13:
		w := types.Boolean{Target: &r.RequierecapturaTotalNumSeries}

		return w

	case 14:
		r.Caracteristicas = make([]Metadato, 0)

		w := ArrayMetadatoWrapper{Target: &r.Caracteristicas}

		return w

	case 15:
		w := types.String{Target: &r.Notas}

		return w

	case 16:
		w := types.String{Target: &r.InstruccionesDePreparacion}

		return w

	case 17:
		w := types.Long{Target: &r.VidaUtilEnDias}

		return w

	case 18:
		w := types.String{Target: &r.CodigoDeVidaUtil}

		return w

	case 19:
		w := types.String{Target: &r.IndicadorDeVidaUtil}

		return w

	case 20:
		w := types.Long{Target: &r.ConsumoEnDias}

		return w

	case 21:
		w := types.Long{Target: &r.VencimientoEnDias}

		return w

	case 22:
		w := types.Long{Target: &r.VidaUtilEntradaEnDias}

		return w

	case 23:
		w := types.String{Target: &r.AcondicionamientoSecundario}

		return w

	case 24:
		w := types.String{Target: &r.ZonaRepo}

		return w

	case 25:
		r.Grupos = make([]Metadato, 0)

		w := ArrayMetadatoWrapper{Target: &r.Grupos}

		return w

	case 26:
		w := types.Double{Target: &r.Volumen}

		return w

	case 27:
		w := types.Double{Target: &r.PesoBruto}

		return w

	case 28:
		w := types.Double{Target: &r.PesoTara}

		return w

	case 29:
		w := types.Double{Target: &r.PesoNeto}

		return w

	case 30:
		r.CamposLibres = make([]Metadato, 0)

		w := ArrayMetadatoWrapper{Target: &r.CamposLibres}

		return w

	}
	panic("Unknown field index")
}

func (r *DetalleDeArticulo) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *DetalleDeArticulo) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ DetalleDeArticulo) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DetalleDeArticulo) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DetalleDeArticulo) HintSize(int)                     { panic("Unsupported operation") }
func (_ DetalleDeArticulo) Finalize()                        {}

func (_ DetalleDeArticulo) AvroCRC64Fingerprint() []byte {
	return []byte(DetalleDeArticuloAvroCRC64Fingerprint)
}

func (r DetalleDeArticulo) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["codigo"], err = json.Marshal(r.Codigo)
	if err != nil {
		return nil, err
	}
	output["ean13"], err = json.Marshal(r.Ean13)
	if err != nil {
		return nil, err
	}
	output["propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["lote"], err = json.Marshal(r.Lote)
	if err != nil {
		return nil, err
	}
	output["otrosDatos"], err = json.Marshal(r.OtrosDatos)
	if err != nil {
		return nil, err
	}
	output["descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	output["claseDeExpedicion"], err = json.Marshal(r.ClaseDeExpedicion)
	if err != nil {
		return nil, err
	}
	output["claseDeArticulo"], err = json.Marshal(r.ClaseDeArticulo)
	if err != nil {
		return nil, err
	}
	output["paisDeOrigen"], err = json.Marshal(r.PaisDeOrigen)
	if err != nil {
		return nil, err
	}
	output["esNumeroDeSerieDeEntradaUnico"], err = json.Marshal(r.EsNumeroDeSerieDeEntradaUnico)
	if err != nil {
		return nil, err
	}
	output["requiereCapturaDatosEntrada"], err = json.Marshal(r.RequiereCapturaDatosEntrada)
	if err != nil {
		return nil, err
	}
	output["esNumeroDeSerieSalidaUnico"], err = json.Marshal(r.EsNumeroDeSerieSalidaUnico)
	if err != nil {
		return nil, err
	}
	output["requiereCapturaDatosSalida"], err = json.Marshal(r.RequiereCapturaDatosSalida)
	if err != nil {
		return nil, err
	}
	output["requierecapturaTotalNumSeries"], err = json.Marshal(r.RequierecapturaTotalNumSeries)
	if err != nil {
		return nil, err
	}
	output["caracteristicas"], err = json.Marshal(r.Caracteristicas)
	if err != nil {
		return nil, err
	}
	output["notas"], err = json.Marshal(r.Notas)
	if err != nil {
		return nil, err
	}
	output["instruccionesDePreparacion"], err = json.Marshal(r.InstruccionesDePreparacion)
	if err != nil {
		return nil, err
	}
	output["vidaUtilEnDias"], err = json.Marshal(r.VidaUtilEnDias)
	if err != nil {
		return nil, err
	}
	output["codigoDeVidaUtil"], err = json.Marshal(r.CodigoDeVidaUtil)
	if err != nil {
		return nil, err
	}
	output["indicadorDeVidaUtil"], err = json.Marshal(r.IndicadorDeVidaUtil)
	if err != nil {
		return nil, err
	}
	output["consumoEnDias"], err = json.Marshal(r.ConsumoEnDias)
	if err != nil {
		return nil, err
	}
	output["vencimientoEnDias"], err = json.Marshal(r.VencimientoEnDias)
	if err != nil {
		return nil, err
	}
	output["vidaUtilEntradaEnDias"], err = json.Marshal(r.VidaUtilEntradaEnDias)
	if err != nil {
		return nil, err
	}
	output["acondicionamientoSecundario"], err = json.Marshal(r.AcondicionamientoSecundario)
	if err != nil {
		return nil, err
	}
	output["zonaRepo"], err = json.Marshal(r.ZonaRepo)
	if err != nil {
		return nil, err
	}
	output["grupos"], err = json.Marshal(r.Grupos)
	if err != nil {
		return nil, err
	}
	output["volumen"], err = json.Marshal(r.Volumen)
	if err != nil {
		return nil, err
	}
	output["pesoBruto"], err = json.Marshal(r.PesoBruto)
	if err != nil {
		return nil, err
	}
	output["pesoTara"], err = json.Marshal(r.PesoTara)
	if err != nil {
		return nil, err
	}
	output["pesoNeto"], err = json.Marshal(r.PesoNeto)
	if err != nil {
		return nil, err
	}
	output["camposLibres"], err = json.Marshal(r.CamposLibres)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DetalleDeArticulo) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["codigo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Codigo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ean13"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ean13); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ean13")
	}
	val = func() json.RawMessage {
		if v, ok := fields["propietario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Propietario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for propietario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["lote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Lote); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for lote")
	}
	val = func() json.RawMessage {
		if v, ok := fields["otrosDatos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OtrosDatos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for otrosDatos")
	}
	val = func() json.RawMessage {
		if v, ok := fields["descripcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Descripcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for descripcion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["claseDeExpedicion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ClaseDeExpedicion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for claseDeExpedicion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["claseDeArticulo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ClaseDeArticulo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for claseDeArticulo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["paisDeOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PaisDeOrigen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for paisDeOrigen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["esNumeroDeSerieDeEntradaUnico"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EsNumeroDeSerieDeEntradaUnico); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for esNumeroDeSerieDeEntradaUnico")
	}
	val = func() json.RawMessage {
		if v, ok := fields["requiereCapturaDatosEntrada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RequiereCapturaDatosEntrada); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for requiereCapturaDatosEntrada")
	}
	val = func() json.RawMessage {
		if v, ok := fields["esNumeroDeSerieSalidaUnico"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EsNumeroDeSerieSalidaUnico); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for esNumeroDeSerieSalidaUnico")
	}
	val = func() json.RawMessage {
		if v, ok := fields["requiereCapturaDatosSalida"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RequiereCapturaDatosSalida); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for requiereCapturaDatosSalida")
	}
	val = func() json.RawMessage {
		if v, ok := fields["requierecapturaTotalNumSeries"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RequierecapturaTotalNumSeries); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for requierecapturaTotalNumSeries")
	}
	val = func() json.RawMessage {
		if v, ok := fields["caracteristicas"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Caracteristicas); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for caracteristicas")
	}
	val = func() json.RawMessage {
		if v, ok := fields["notas"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Notas); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for notas")
	}
	val = func() json.RawMessage {
		if v, ok := fields["instruccionesDePreparacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.InstruccionesDePreparacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for instruccionesDePreparacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["vidaUtilEnDias"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VidaUtilEnDias); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for vidaUtilEnDias")
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoDeVidaUtil"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeVidaUtil); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigoDeVidaUtil")
	}
	val = func() json.RawMessage {
		if v, ok := fields["indicadorDeVidaUtil"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IndicadorDeVidaUtil); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for indicadorDeVidaUtil")
	}
	val = func() json.RawMessage {
		if v, ok := fields["consumoEnDias"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ConsumoEnDias); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for consumoEnDias")
	}
	val = func() json.RawMessage {
		if v, ok := fields["vencimientoEnDias"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VencimientoEnDias); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for vencimientoEnDias")
	}
	val = func() json.RawMessage {
		if v, ok := fields["vidaUtilEntradaEnDias"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VidaUtilEntradaEnDias); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for vidaUtilEntradaEnDias")
	}
	val = func() json.RawMessage {
		if v, ok := fields["acondicionamientoSecundario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AcondicionamientoSecundario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for acondicionamientoSecundario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["zonaRepo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ZonaRepo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for zonaRepo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["grupos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Grupos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for grupos")
	}
	val = func() json.RawMessage {
		if v, ok := fields["volumen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Volumen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for volumen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["pesoBruto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PesoBruto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for pesoBruto")
	}
	val = func() json.RawMessage {
		if v, ok := fields["pesoTara"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PesoTara); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for pesoTara")
	}
	val = func() json.RawMessage {
		if v, ok := fields["pesoNeto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PesoNeto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for pesoNeto")
	}
	val = func() json.RawMessage {
		if v, ok := fields["camposLibres"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CamposLibres); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for camposLibres")
	}
	return nil
}
