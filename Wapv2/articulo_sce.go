// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ArticuloSCE.avsc
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

type ArticuloSCE struct {
	Codigo string `json:"codigo"`

	Ean string `json:"ean"`

	TipoEan string `json:"tipoEan"`

	Propietario string `json:"propietario"`

	Descripcion string `json:"descripcion"`

	DatosLogisticos DatosLogisticos `json:"datosLogisticos"`

	OtrosDatos *UnionNullMetadato `json:"otrosDatos"`

	Notas *UnionNullString `json:"notas"`

	ClaseDeArticulo *UnionNullString `json:"claseDeArticulo"`

	PaisDeOrigen *UnionNullString `json:"paisDeOrigen"`

	EsNumeroDeSerieDeEntradaUnico *UnionNullBool `json:"esNumeroDeSerieDeEntradaUnico"`

	EsNumeroDeSerieSalidaUnico *UnionNullBool `json:"esNumeroDeSerieSalidaUnico"`

	InstruccionesDePreparacion *UnionNullString `json:"instruccionesDePreparacion"`

	VidaUtilSalidaEnDias *UnionNullInt `json:"vidaUtilSalidaEnDias"`

	CodigoDeVidaUtil *UnionNullString `json:"codigoDeVidaUtil"`

	IndicadorDeVidaUtil *UnionNullString `json:"indicadorDeVidaUtil"`

	ConsumoEnDias *UnionNullInt `json:"consumoEnDias"`

	VencimientoEnDias *UnionNullInt `json:"vencimientoEnDias"`

	VidaUtilEntradaEnDias *UnionNullInt `json:"vidaUtilEntradaEnDias"`

	Serializado *UnionNullString `json:"serializado"`

	Grupos *UnionNullMetadato `json:"grupos"`

	CamposLibres *UnionNullMetadato `json:"camposLibres"`

	Coleccion *UnionNullString `json:"coleccion"`

	Tema *UnionNullString `json:"tema"`

	Temporada *UnionNullString `json:"temporada"`

	Estilo *UnionNullString `json:"estilo"`

	Color *UnionNullString `json:"color"`

	IniciodeTemporada *UnionNullString `json:"iniciodeTemporada"`

	FindeTemporada *UnionNullString `json:"findeTemporada"`

	Talle *UnionNullString `json:"talle"`

	Rubro *UnionNullString `json:"rubro"`

	Pavu *UnionNullString `json:"pavu"`

	Psicotropico *UnionNullString `json:"psicotropico"`

	Temperatura *UnionNullString `json:"temperatura"`
}

const ArticuloSCEAvroCRC64Fingerprint = "b\xd8\xddf1\xa0l\x1b"

func NewArticuloSCE() ArticuloSCE {
	r := ArticuloSCE{}
	r.DatosLogisticos = NewDatosLogisticos()

	r.OtrosDatos = nil
	r.Grupos = nil
	r.CamposLibres = nil
	return r
}

func DeserializeArticuloSCE(r io.Reader) (ArticuloSCE, error) {
	t := NewArticuloSCE()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeArticuloSCEFromSchema(r io.Reader, schema string) (ArticuloSCE, error) {
	t := NewArticuloSCE()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeArticuloSCE(r ArticuloSCE, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Codigo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Ean, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TipoEan, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Descripcion, w)
	if err != nil {
		return err
	}
	err = writeDatosLogisticos(r.DatosLogisticos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullMetadato(r.OtrosDatos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Notas, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ClaseDeArticulo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.PaisDeOrigen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.EsNumeroDeSerieDeEntradaUnico, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.EsNumeroDeSerieSalidaUnico, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.InstruccionesDePreparacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.VidaUtilSalidaEnDias, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoDeVidaUtil, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.IndicadorDeVidaUtil, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.ConsumoEnDias, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.VencimientoEnDias, w)
	if err != nil {
		return err
	}
	err = writeUnionNullInt(r.VidaUtilEntradaEnDias, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Serializado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullMetadato(r.Grupos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullMetadato(r.CamposLibres, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Coleccion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Tema, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Temporada, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Estilo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Color, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.IniciodeTemporada, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FindeTemporada, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Talle, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Rubro, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Pavu, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Psicotropico, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Temperatura, w)
	if err != nil {
		return err
	}
	return err
}

func (r ArticuloSCE) Serialize(w io.Writer) error {
	return writeArticuloSCE(r, w)
}

func (r ArticuloSCE) Schema() string {
	return "{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"ean\",\"type\":\"string\"},{\"name\":\"tipoEan\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"descripcion\",\"type\":\"string\"},{\"name\":\"datosLogisticos\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"volumen\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"pesoBruto\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"pesoTara\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"pesoNeto\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"cantidadporPaquete\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"cantidadporCaja\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"cantidadporPallet\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"alturaUnidad\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"largoUnidad\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"anchoUnidad\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"pesoPack\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"alturaPack\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"largoPack\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"anchoPack\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"pesoCaja\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"altoCaja\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"largoCaja\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"anchoCaja\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"pesoPallet\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"altoPallet\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"largoPallet\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"anchoPallet\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"nivelesporPallet\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"cajasporNivel\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"altura\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"longitud\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"anchura\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"volumencubico\",\"type\":[\"null\",\"float\"]}],\"name\":\"DatosLogisticos\",\"type\":\"record\"}},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"}]},{\"name\":\"notas\",\"type\":[\"null\",\"string\"]},{\"name\":\"claseDeArticulo\",\"type\":[\"null\",\"string\"]},{\"name\":\"paisDeOrigen\",\"type\":[\"null\",\"string\"]},{\"name\":\"esNumeroDeSerieDeEntradaUnico\",\"type\":[\"null\",\"boolean\"]},{\"name\":\"esNumeroDeSerieSalidaUnico\",\"type\":[\"null\",\"boolean\"]},{\"name\":\"instruccionesDePreparacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"vidaUtilSalidaEnDias\",\"type\":[\"null\",\"int\"]},{\"name\":\"codigoDeVidaUtil\",\"type\":[\"null\",\"string\"]},{\"name\":\"indicadorDeVidaUtil\",\"type\":[\"null\",\"string\"]},{\"name\":\"consumoEnDias\",\"type\":[\"null\",\"int\"]},{\"name\":\"vencimientoEnDias\",\"type\":[\"null\",\"int\"]},{\"name\":\"vidaUtilEntradaEnDias\",\"type\":[\"null\",\"int\"]},{\"name\":\"serializado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"grupos\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.Metadato\"]},{\"default\":null,\"name\":\"camposLibres\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.Metadato\"]},{\"name\":\"coleccion\",\"type\":[\"null\",\"string\"]},{\"name\":\"tema\",\"type\":[\"null\",\"string\"]},{\"name\":\"temporada\",\"type\":[\"null\",\"string\"]},{\"name\":\"estilo\",\"type\":[\"null\",\"string\"]},{\"name\":\"color\",\"type\":[\"null\",\"string\"]},{\"name\":\"iniciodeTemporada\",\"type\":[\"null\",\"string\"]},{\"name\":\"findeTemporada\",\"type\":[\"null\",\"string\"]},{\"name\":\"talle\",\"type\":[\"null\",\"string\"]},{\"name\":\"rubro\",\"type\":[\"null\",\"string\"]},{\"name\":\"pavu\",\"type\":[\"null\",\"string\"]},{\"name\":\"psicotropico\",\"type\":[\"null\",\"string\"]},{\"name\":\"temperatura\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.Wapv2.Events.Record.ArticuloSCE\",\"type\":\"record\"}"
}

func (r ArticuloSCE) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.ArticuloSCE"
}

func (_ ArticuloSCE) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ArticuloSCE) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ArticuloSCE) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ArticuloSCE) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ArticuloSCE) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ArticuloSCE) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ArticuloSCE) SetString(v string)   { panic("Unsupported operation") }
func (_ ArticuloSCE) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ArticuloSCE) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Codigo}

		return w

	case 1:
		w := types.String{Target: &r.Ean}

		return w

	case 2:
		w := types.String{Target: &r.TipoEan}

		return w

	case 3:
		w := types.String{Target: &r.Propietario}

		return w

	case 4:
		w := types.String{Target: &r.Descripcion}

		return w

	case 5:
		r.DatosLogisticos = NewDatosLogisticos()

		w := types.Record{Target: &r.DatosLogisticos}

		return w

	case 6:
		r.OtrosDatos = NewUnionNullMetadato()

		return r.OtrosDatos
	case 7:
		r.Notas = NewUnionNullString()

		return r.Notas
	case 8:
		r.ClaseDeArticulo = NewUnionNullString()

		return r.ClaseDeArticulo
	case 9:
		r.PaisDeOrigen = NewUnionNullString()

		return r.PaisDeOrigen
	case 10:
		r.EsNumeroDeSerieDeEntradaUnico = NewUnionNullBool()

		return r.EsNumeroDeSerieDeEntradaUnico
	case 11:
		r.EsNumeroDeSerieSalidaUnico = NewUnionNullBool()

		return r.EsNumeroDeSerieSalidaUnico
	case 12:
		r.InstruccionesDePreparacion = NewUnionNullString()

		return r.InstruccionesDePreparacion
	case 13:
		r.VidaUtilSalidaEnDias = NewUnionNullInt()

		return r.VidaUtilSalidaEnDias
	case 14:
		r.CodigoDeVidaUtil = NewUnionNullString()

		return r.CodigoDeVidaUtil
	case 15:
		r.IndicadorDeVidaUtil = NewUnionNullString()

		return r.IndicadorDeVidaUtil
	case 16:
		r.ConsumoEnDias = NewUnionNullInt()

		return r.ConsumoEnDias
	case 17:
		r.VencimientoEnDias = NewUnionNullInt()

		return r.VencimientoEnDias
	case 18:
		r.VidaUtilEntradaEnDias = NewUnionNullInt()

		return r.VidaUtilEntradaEnDias
	case 19:
		r.Serializado = NewUnionNullString()

		return r.Serializado
	case 20:
		r.Grupos = NewUnionNullMetadato()

		return r.Grupos
	case 21:
		r.CamposLibres = NewUnionNullMetadato()

		return r.CamposLibres
	case 22:
		r.Coleccion = NewUnionNullString()

		return r.Coleccion
	case 23:
		r.Tema = NewUnionNullString()

		return r.Tema
	case 24:
		r.Temporada = NewUnionNullString()

		return r.Temporada
	case 25:
		r.Estilo = NewUnionNullString()

		return r.Estilo
	case 26:
		r.Color = NewUnionNullString()

		return r.Color
	case 27:
		r.IniciodeTemporada = NewUnionNullString()

		return r.IniciodeTemporada
	case 28:
		r.FindeTemporada = NewUnionNullString()

		return r.FindeTemporada
	case 29:
		r.Talle = NewUnionNullString()

		return r.Talle
	case 30:
		r.Rubro = NewUnionNullString()

		return r.Rubro
	case 31:
		r.Pavu = NewUnionNullString()

		return r.Pavu
	case 32:
		r.Psicotropico = NewUnionNullString()

		return r.Psicotropico
	case 33:
		r.Temperatura = NewUnionNullString()

		return r.Temperatura
	}
	panic("Unknown field index")
}

func (r *ArticuloSCE) SetDefault(i int) {
	switch i {
	case 6:
		r.OtrosDatos = nil
		return
	case 20:
		r.Grupos = nil
		return
	case 21:
		r.CamposLibres = nil
		return
	}
	panic("Unknown field index")
}

func (r *ArticuloSCE) NullField(i int) {
	switch i {
	case 6:
		r.OtrosDatos = nil
		return
	case 7:
		r.Notas = nil
		return
	case 8:
		r.ClaseDeArticulo = nil
		return
	case 9:
		r.PaisDeOrigen = nil
		return
	case 10:
		r.EsNumeroDeSerieDeEntradaUnico = nil
		return
	case 11:
		r.EsNumeroDeSerieSalidaUnico = nil
		return
	case 12:
		r.InstruccionesDePreparacion = nil
		return
	case 13:
		r.VidaUtilSalidaEnDias = nil
		return
	case 14:
		r.CodigoDeVidaUtil = nil
		return
	case 15:
		r.IndicadorDeVidaUtil = nil
		return
	case 16:
		r.ConsumoEnDias = nil
		return
	case 17:
		r.VencimientoEnDias = nil
		return
	case 18:
		r.VidaUtilEntradaEnDias = nil
		return
	case 19:
		r.Serializado = nil
		return
	case 20:
		r.Grupos = nil
		return
	case 21:
		r.CamposLibres = nil
		return
	case 22:
		r.Coleccion = nil
		return
	case 23:
		r.Tema = nil
		return
	case 24:
		r.Temporada = nil
		return
	case 25:
		r.Estilo = nil
		return
	case 26:
		r.Color = nil
		return
	case 27:
		r.IniciodeTemporada = nil
		return
	case 28:
		r.FindeTemporada = nil
		return
	case 29:
		r.Talle = nil
		return
	case 30:
		r.Rubro = nil
		return
	case 31:
		r.Pavu = nil
		return
	case 32:
		r.Psicotropico = nil
		return
	case 33:
		r.Temperatura = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ArticuloSCE) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArticuloSCE) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ArticuloSCE) HintSize(int)                     { panic("Unsupported operation") }
func (_ ArticuloSCE) Finalize()                        {}

func (_ ArticuloSCE) AvroCRC64Fingerprint() []byte {
	return []byte(ArticuloSCEAvroCRC64Fingerprint)
}

func (r ArticuloSCE) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["codigo"], err = json.Marshal(r.Codigo)
	if err != nil {
		return nil, err
	}
	output["ean"], err = json.Marshal(r.Ean)
	if err != nil {
		return nil, err
	}
	output["tipoEan"], err = json.Marshal(r.TipoEan)
	if err != nil {
		return nil, err
	}
	output["propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	output["datosLogisticos"], err = json.Marshal(r.DatosLogisticos)
	if err != nil {
		return nil, err
	}
	output["otrosDatos"], err = json.Marshal(r.OtrosDatos)
	if err != nil {
		return nil, err
	}
	output["notas"], err = json.Marshal(r.Notas)
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
	output["esNumeroDeSerieSalidaUnico"], err = json.Marshal(r.EsNumeroDeSerieSalidaUnico)
	if err != nil {
		return nil, err
	}
	output["instruccionesDePreparacion"], err = json.Marshal(r.InstruccionesDePreparacion)
	if err != nil {
		return nil, err
	}
	output["vidaUtilSalidaEnDias"], err = json.Marshal(r.VidaUtilSalidaEnDias)
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
	output["serializado"], err = json.Marshal(r.Serializado)
	if err != nil {
		return nil, err
	}
	output["grupos"], err = json.Marshal(r.Grupos)
	if err != nil {
		return nil, err
	}
	output["camposLibres"], err = json.Marshal(r.CamposLibres)
	if err != nil {
		return nil, err
	}
	output["coleccion"], err = json.Marshal(r.Coleccion)
	if err != nil {
		return nil, err
	}
	output["tema"], err = json.Marshal(r.Tema)
	if err != nil {
		return nil, err
	}
	output["temporada"], err = json.Marshal(r.Temporada)
	if err != nil {
		return nil, err
	}
	output["estilo"], err = json.Marshal(r.Estilo)
	if err != nil {
		return nil, err
	}
	output["color"], err = json.Marshal(r.Color)
	if err != nil {
		return nil, err
	}
	output["iniciodeTemporada"], err = json.Marshal(r.IniciodeTemporada)
	if err != nil {
		return nil, err
	}
	output["findeTemporada"], err = json.Marshal(r.FindeTemporada)
	if err != nil {
		return nil, err
	}
	output["talle"], err = json.Marshal(r.Talle)
	if err != nil {
		return nil, err
	}
	output["rubro"], err = json.Marshal(r.Rubro)
	if err != nil {
		return nil, err
	}
	output["pavu"], err = json.Marshal(r.Pavu)
	if err != nil {
		return nil, err
	}
	output["psicotropico"], err = json.Marshal(r.Psicotropico)
	if err != nil {
		return nil, err
	}
	output["temperatura"], err = json.Marshal(r.Temperatura)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ArticuloSCE) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["ean"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ean); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ean")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoEan"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoEan); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tipoEan")
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
		if v, ok := fields["datosLogisticos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DatosLogisticos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for datosLogisticos")
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
		r.OtrosDatos = NewUnionNullMetadato()

		r.OtrosDatos = nil
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
		if v, ok := fields["vidaUtilSalidaEnDias"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VidaUtilSalidaEnDias); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for vidaUtilSalidaEnDias")
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
		if v, ok := fields["serializado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Serializado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for serializado")
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
		r.Grupos = NewUnionNullMetadato()

		r.Grupos = nil
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
		r.CamposLibres = NewUnionNullMetadato()

		r.CamposLibres = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["coleccion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Coleccion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for coleccion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tema"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tema); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tema")
	}
	val = func() json.RawMessage {
		if v, ok := fields["temporada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Temporada); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for temporada")
	}
	val = func() json.RawMessage {
		if v, ok := fields["estilo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estilo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for estilo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["color"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Color); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for color")
	}
	val = func() json.RawMessage {
		if v, ok := fields["iniciodeTemporada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IniciodeTemporada); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for iniciodeTemporada")
	}
	val = func() json.RawMessage {
		if v, ok := fields["findeTemporada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FindeTemporada); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for findeTemporada")
	}
	val = func() json.RawMessage {
		if v, ok := fields["talle"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Talle); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for talle")
	}
	val = func() json.RawMessage {
		if v, ok := fields["rubro"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Rubro); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for rubro")
	}
	val = func() json.RawMessage {
		if v, ok := fields["pavu"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Pavu); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for pavu")
	}
	val = func() json.RawMessage {
		if v, ok := fields["psicotropico"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Psicotropico); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for psicotropico")
	}
	val = func() json.RawMessage {
		if v, ok := fields["temperatura"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Temperatura); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for temperatura")
	}
	return nil
}
