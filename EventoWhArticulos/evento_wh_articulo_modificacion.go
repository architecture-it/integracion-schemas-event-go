// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhArticuloModificacion.avsc
 */
package EventoWhArticulosEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EventoWhArticuloModificacion struct {
	Identificacion Identificacion `json:"Identificacion"`

	Cabecera Cabecera `json:"Cabecera"`
}

const EventoWhArticuloModificacionAvroCRC64Fingerprint = "ϳ\x80\x9f\xad\xad3\xb7"

func NewEventoWhArticuloModificacion() EventoWhArticuloModificacion {
	r := EventoWhArticuloModificacion{}
	r.Identificacion = NewIdentificacion()

	r.Cabecera = NewCabecera()

	return r
}

func DeserializeEventoWhArticuloModificacion(r io.Reader) (EventoWhArticuloModificacion, error) {
	t := NewEventoWhArticuloModificacion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoWhArticuloModificacionFromSchema(r io.Reader, schema string) (EventoWhArticuloModificacion, error) {
	t := NewEventoWhArticuloModificacion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoWhArticuloModificacion(r EventoWhArticuloModificacion, w io.Writer) error {
	var err error
	err = writeIdentificacion(r.Identificacion, w)
	if err != nil {
		return err
	}
	err = writeCabecera(r.Cabecera, w)
	if err != nil {
		return err
	}
	return err
}

func (r EventoWhArticuloModificacion) Serialize(w io.Writer) error {
	return writeEventoWhArticuloModificacion(r, w)
}

func (r EventoWhArticuloModificacion) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.EventoWhArticulos.Events.EventoArticuloModificacion\",\"type\":\"record\"}},{\"name\":\"Cabecera\",\"type\":{\"fields\":[{\"name\":\"TipoRotacionABC\",\"type\":[\"null\",\"string\"]},{\"name\":\"TieneAcondi\",\"type\":[\"null\",\"string\"]},{\"name\":\"FinTemporada\",\"type\":[\"null\",\"string\"]},{\"name\":\"InicioTemporada\",\"type\":[\"null\",\"string\"]},{\"name\":\"Coleccion\",\"type\":[\"null\",\"string\"]},{\"name\":\"Color\",\"type\":[\"null\",\"string\"]},{\"name\":\"Atributos\",\"type\":[\"null\",\"string\"]},{\"name\":\"PrecioAsociadoAlSKU\",\"type\":[\"null\",\"float\"]},{\"name\":\"PaisOrigen\",\"type\":[\"null\",\"string\"]},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"name\":\"DigitosMinimosSeries\",\"type\":[\"null\",\"string\"]},{\"name\":\"DigitosMaximosSeries\",\"type\":[\"null\",\"string\"]},{\"name\":\"LimiteImpresion\",\"type\":[\"null\",\"string\"]},{\"name\":\"ContadorGeneracionSeries\",\"type\":[\"null\",\"string\"]},{\"name\":\"Rubro\",\"type\":[\"null\",\"string\"]},{\"name\":\"Pavu\",\"type\":[\"null\",\"string\"]},{\"name\":\"Psicotropico\",\"type\":[\"null\",\"string\"]},{\"name\":\"Temperatura\",\"type\":[\"null\",\"string\"]},{\"name\":\"Serializado\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoValidacionLote\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoDatamatrix\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoAgrupadora\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoEtiqueta\",\"type\":[\"null\",\"string\"]},{\"name\":\"PickeaTodos\",\"type\":[\"null\",\"string\"]},{\"name\":\"SerieDirigida\",\"type\":[\"null\",\"string\"]},{\"name\":\"GeneraSerie\",\"type\":[\"null\",\"string\"]},{\"name\":\"ControlaSeries\",\"type\":[\"null\",\"string\"]},{\"name\":\"InformaLevantamientoCuarentena\",\"type\":[\"null\",\"string\"]},{\"name\":\"EventoLevantamientoCuarentena\",\"type\":[\"null\",\"string\"]},{\"name\":\"GeneraAgrupadora\",\"type\":[\"null\",\"string\"]},{\"name\":\"AltaPorAPI\",\"type\":[\"null\",\"string\"]},{\"name\":\"PropietarioEcommerce\",\"type\":[\"null\",\"string\"]},{\"name\":\"CodigoCliente\",\"type\":[\"null\",\"string\"]},{\"name\":\"SKUOriginalCliente\",\"type\":[\"null\",\"string\"]},{\"name\":\"ReservadoFFC\",\"type\":[\"null\",\"string\"]},{\"name\":\"Usuario\",\"type\":[\"null\",\"string\"]},{\"name\":\"ControlSeriesRecepcion\",\"type\":[\"null\",\"string\"]},{\"name\":\"Notas\",\"type\":[\"null\",\"string\"]},{\"name\":\"InstruccionesPreparacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"ControlSeriesExpedicion\",\"type\":[\"null\",\"string\"]},{\"name\":\"CodigoPaquete\",\"type\":\"string\"},{\"name\":\"PercioLinea\",\"type\":[\"null\",\"float\"]},{\"name\":\"UbicacionControlCalidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoRotacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"Temporada\",\"type\":[\"null\",\"string\"]},{\"name\":\"VidaUtilEnDias\",\"type\":[\"null\",\"int\"]},{\"name\":\"TipoValidacionVidaUtil\",\"type\":[\"null\",\"string\"]},{\"name\":\"IndicadorVidaUtil\",\"type\":[\"null\",\"string\"]},{\"name\":\"VidaUtilEntrada\",\"type\":[\"null\",\"int\"]},{\"name\":\"CodigoArticulo\",\"type\":\"string\"},{\"name\":\"Talle\",\"type\":[\"null\",\"string\"]},{\"name\":\"VolumenCentimetros\",\"type\":\"float\"},{\"name\":\"PesoBrutoKg\",\"type\":\"float\"},{\"name\":\"PesoNetoKg\",\"type\":\"float\"},{\"name\":\"CategoriaStock\",\"type\":[\"null\",\"string\"]},{\"name\":\"Estilo\",\"type\":[\"null\",\"string\"]},{\"name\":\"Tema\",\"type\":[\"null\",\"string\"]},{\"name\":\"ConsumoAntesDeXDias\",\"type\":[\"null\",\"int\"]},{\"name\":\"ConsumoVencimiento\",\"type\":[\"null\",\"int\"]},{\"name\":\"ValidaLoteWOS\",\"type\":[\"null\",\"string\"]},{\"name\":\"ValidaSerieWOS\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ArticulosAlternativos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"ArticuloAlternativo\",\"type\":\"string\"},{\"name\":\"Tipo\",\"type\":\"string\"}],\"name\":\"ArticulosAlternativos\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"Cabecera\",\"namespace\":\"Andreani.EventoWhArticulos.Events.EventoArticuloModificacion\",\"type\":\"record\"}}],\"name\":\"Andreani.EventoWhArticulos.Events.Record.EventoWhArticuloModificacion\",\"type\":\"record\"}"
}

func (r EventoWhArticuloModificacion) SchemaName() string {
	return "Andreani.EventoWhArticulos.Events.Record.EventoWhArticuloModificacion"
}

func (_ EventoWhArticuloModificacion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoWhArticuloModificacion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoWhArticuloModificacion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoWhArticuloModificacion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoWhArticuloModificacion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoWhArticuloModificacion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoWhArticuloModificacion) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoWhArticuloModificacion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoWhArticuloModificacion) Get(i int) types.Field {
	switch i {
	case 0:
		r.Identificacion = NewIdentificacion()

		w := types.Record{Target: &r.Identificacion}

		return w

	case 1:
		r.Cabecera = NewCabecera()

		w := types.Record{Target: &r.Cabecera}

		return w

	}
	panic("Unknown field index")
}

func (r *EventoWhArticuloModificacion) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoWhArticuloModificacion) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EventoWhArticuloModificacion) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ EventoWhArticuloModificacion) AppendArray() types.Field { panic("Unsupported operation") }
func (_ EventoWhArticuloModificacion) HintSize(int)             { panic("Unsupported operation") }
func (_ EventoWhArticuloModificacion) Finalize()                {}

func (_ EventoWhArticuloModificacion) AvroCRC64Fingerprint() []byte {
	return []byte(EventoWhArticuloModificacionAvroCRC64Fingerprint)
}

func (r EventoWhArticuloModificacion) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(output)
}

func (r *EventoWhArticuloModificacion) UnmarshalJSON(data []byte) error {
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
	return nil
}
