// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhRecepcionCreaccion.avsc
 */
package EventoWhRecepcionEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Detalle struct {
	Propietario string `json:"Propietario"`

	SKU string `json:"SKU"`

	UbicacionDestino *UnionNullString `json:"UbicacionDestino"`

	LPNDestino *UnionNullString `json:"LPNDestino"`

	PaqueteLote *UnionNullString `json:"PaqueteLote"`

	LoteCajitaFabricante *UnionNullString `json:"LoteCajitaFabricante"`

	LoteSecundario *UnionNullString `json:"LoteSecundario"`

	FechaFabricacion *UnionNullLong `json:"FechaFabricacion"`

	FechaVencimiento *UnionNullLong `json:"FechaVencimiento"`

	ProductoTrazable *UnionNullString `json:"ProductoTrazable"`

	AlmacenConsumo *UnionNullString `json:"AlmacenConsumo"`

	EstadoLote *UnionNullString `json:"EstadoLote"`

	BloqueoUbicacion *UnionNullString `json:"BloqueoUbicacion"`

	VidaUtilLote *UnionNullString `json:"VidaUtilLote"`

	EntregaAntesDe *UnionNullLong `json:"EntregaAntesDe"`

	ConsumoAntesDe *UnionNullLong `json:"ConsumoAntesDe"`

	Contramuestras *UnionNullString `json:"Contramuestras"`

	EstadoOTAcondi *UnionNullString `json:"EstadoOTAcondi"`

	EstadoOTTraza *UnionNullString `json:"EstadoOTTraza"`

	TipoAcondi *UnionNullString `json:"TipoAcondi"`

	TipoTraza *UnionNullString `json:"TipoTraza"`

	ContratoServicioIngreso *UnionNullString `json:"ContratoServicioIngreso"`

	NomenclaturaContratoServicioIngreso *UnionNullString `json:"NomenclaturaContratoServicioIngreso"`

	DescripcionServicioIngreso *UnionNullString `json:"DescripcionServicioIngreso"`

	TipoLineaMatriz *UnionNullString `json:"TipoLineaMatriz"`

	CodConCalidad *UnionNullString `json:"CodConCalidad"`

	AccionConCalidad *UnionNullString `json:"AccionConCalidad"`

	ResultadoConCalidad *UnionNullString `json:"ResultadoConCalidad"`

	LineaExterna *UnionNullString `json:"LineaExterna"`

	CantEsperada float32 `json:"CantEsperada"`

	CantRecibida float32 `json:"CantRecibida"`

	ValorDeclaradoLinea *UnionNullString `json:"ValorDeclaradoLinea"`

	UnidadMedida *UnionNullString `json:"UnidadMedida"`

	LineaRecepcionWH *UnionNullString `json:"LineaRecepcionWH"`
}

const DetalleAvroCRC64Fingerprint = "e؍\xd8\x19\xa9\xd4\xfb"

func NewDetalle() Detalle {
	r := Detalle{}
	return r
}

func DeserializeDetalle(r io.Reader) (Detalle, error) {
	t := NewDetalle()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDetalleFromSchema(r io.Reader, schema string) (Detalle, error) {
	t := NewDetalle()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDetalle(r Detalle, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.SKU, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.UbicacionDestino, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LPNDestino, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.PaqueteLote, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LoteCajitaFabricante, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LoteSecundario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaFabricacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaVencimiento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ProductoTrazable, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.AlmacenConsumo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstadoLote, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.BloqueoUbicacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.VidaUtilLote, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.EntregaAntesDe, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.ConsumoAntesDe, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Contramuestras, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstadoOTAcondi, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstadoOTTraza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoAcondi, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoTraza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ContratoServicioIngreso, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NomenclaturaContratoServicioIngreso, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DescripcionServicioIngreso, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoLineaMatriz, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodConCalidad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.AccionConCalidad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ResultadoConCalidad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LineaExterna, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.CantEsperada, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.CantRecibida, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ValorDeclaradoLinea, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.UnidadMedida, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LineaRecepcionWH, w)
	if err != nil {
		return err
	}
	return err
}

func (r Detalle) Serialize(w io.Writer) error {
	return writeDetalle(r, w)
}

func (r Detalle) Schema() string {
	return "{\"fields\":[{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"SKU\",\"type\":\"string\"},{\"name\":\"UbicacionDestino\",\"type\":[\"null\",\"string\"]},{\"name\":\"LPNDestino\",\"type\":[\"null\",\"string\"]},{\"name\":\"PaqueteLote\",\"type\":[\"null\",\"string\"]},{\"name\":\"LoteCajitaFabricante\",\"type\":[\"null\",\"string\"]},{\"name\":\"LoteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ProductoTrazable\",\"type\":[\"null\",\"string\"]},{\"name\":\"AlmacenConsumo\",\"type\":[\"null\",\"string\"]},{\"name\":\"EstadoLote\",\"type\":[\"null\",\"string\"]},{\"name\":\"BloqueoUbicacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"VidaUtilLote\",\"type\":[\"null\",\"string\"]},{\"name\":\"EntregaAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ConsumoAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"Contramuestras\",\"type\":[\"null\",\"string\"]},{\"name\":\"EstadoOTAcondi\",\"type\":[\"null\",\"string\"]},{\"name\":\"EstadoOTTraza\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoAcondi\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoTraza\",\"type\":[\"null\",\"string\"]},{\"name\":\"ContratoServicioIngreso\",\"type\":[\"null\",\"string\"]},{\"name\":\"NomenclaturaContratoServicioIngreso\",\"type\":[\"null\",\"string\"]},{\"name\":\"DescripcionServicioIngreso\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoLineaMatriz\",\"type\":[\"null\",\"string\"]},{\"name\":\"CodConCalidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"AccionConCalidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"ResultadoConCalidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"LineaExterna\",\"type\":[\"null\",\"string\"]},{\"name\":\"CantEsperada\",\"type\":\"float\"},{\"name\":\"CantRecibida\",\"type\":\"float\"},{\"name\":\"ValorDeclaradoLinea\",\"type\":[\"null\",\"string\"]},{\"name\":\"UnidadMedida\",\"type\":[\"null\",\"string\"]},{\"name\":\"LineaRecepcionWH\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.EventoWhRecepcion.Events.RecepcionCreacionCommon.Detalle\",\"type\":\"record\"}"
}

func (r Detalle) SchemaName() string {
	return "Andreani.EventoWhRecepcion.Events.RecepcionCreacionCommon.Detalle"
}

func (_ Detalle) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Detalle) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Detalle) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Detalle) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Detalle) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Detalle) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Detalle) SetString(v string)   { panic("Unsupported operation") }
func (_ Detalle) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Detalle) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Propietario}

		return w

	case 1:
		w := types.String{Target: &r.SKU}

		return w

	case 2:
		r.UbicacionDestino = NewUnionNullString()

		return r.UbicacionDestino
	case 3:
		r.LPNDestino = NewUnionNullString()

		return r.LPNDestino
	case 4:
		r.PaqueteLote = NewUnionNullString()

		return r.PaqueteLote
	case 5:
		r.LoteCajitaFabricante = NewUnionNullString()

		return r.LoteCajitaFabricante
	case 6:
		r.LoteSecundario = NewUnionNullString()

		return r.LoteSecundario
	case 7:
		r.FechaFabricacion = NewUnionNullLong()

		return r.FechaFabricacion
	case 8:
		r.FechaVencimiento = NewUnionNullLong()

		return r.FechaVencimiento
	case 9:
		r.ProductoTrazable = NewUnionNullString()

		return r.ProductoTrazable
	case 10:
		r.AlmacenConsumo = NewUnionNullString()

		return r.AlmacenConsumo
	case 11:
		r.EstadoLote = NewUnionNullString()

		return r.EstadoLote
	case 12:
		r.BloqueoUbicacion = NewUnionNullString()

		return r.BloqueoUbicacion
	case 13:
		r.VidaUtilLote = NewUnionNullString()

		return r.VidaUtilLote
	case 14:
		r.EntregaAntesDe = NewUnionNullLong()

		return r.EntregaAntesDe
	case 15:
		r.ConsumoAntesDe = NewUnionNullLong()

		return r.ConsumoAntesDe
	case 16:
		r.Contramuestras = NewUnionNullString()

		return r.Contramuestras
	case 17:
		r.EstadoOTAcondi = NewUnionNullString()

		return r.EstadoOTAcondi
	case 18:
		r.EstadoOTTraza = NewUnionNullString()

		return r.EstadoOTTraza
	case 19:
		r.TipoAcondi = NewUnionNullString()

		return r.TipoAcondi
	case 20:
		r.TipoTraza = NewUnionNullString()

		return r.TipoTraza
	case 21:
		r.ContratoServicioIngreso = NewUnionNullString()

		return r.ContratoServicioIngreso
	case 22:
		r.NomenclaturaContratoServicioIngreso = NewUnionNullString()

		return r.NomenclaturaContratoServicioIngreso
	case 23:
		r.DescripcionServicioIngreso = NewUnionNullString()

		return r.DescripcionServicioIngreso
	case 24:
		r.TipoLineaMatriz = NewUnionNullString()

		return r.TipoLineaMatriz
	case 25:
		r.CodConCalidad = NewUnionNullString()

		return r.CodConCalidad
	case 26:
		r.AccionConCalidad = NewUnionNullString()

		return r.AccionConCalidad
	case 27:
		r.ResultadoConCalidad = NewUnionNullString()

		return r.ResultadoConCalidad
	case 28:
		r.LineaExterna = NewUnionNullString()

		return r.LineaExterna
	case 29:
		w := types.Float{Target: &r.CantEsperada}

		return w

	case 30:
		w := types.Float{Target: &r.CantRecibida}

		return w

	case 31:
		r.ValorDeclaradoLinea = NewUnionNullString()

		return r.ValorDeclaradoLinea
	case 32:
		r.UnidadMedida = NewUnionNullString()

		return r.UnidadMedida
	case 33:
		r.LineaRecepcionWH = NewUnionNullString()

		return r.LineaRecepcionWH
	}
	panic("Unknown field index")
}

func (r *Detalle) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Detalle) NullField(i int) {
	switch i {
	case 2:
		r.UbicacionDestino = nil
		return
	case 3:
		r.LPNDestino = nil
		return
	case 4:
		r.PaqueteLote = nil
		return
	case 5:
		r.LoteCajitaFabricante = nil
		return
	case 6:
		r.LoteSecundario = nil
		return
	case 7:
		r.FechaFabricacion = nil
		return
	case 8:
		r.FechaVencimiento = nil
		return
	case 9:
		r.ProductoTrazable = nil
		return
	case 10:
		r.AlmacenConsumo = nil
		return
	case 11:
		r.EstadoLote = nil
		return
	case 12:
		r.BloqueoUbicacion = nil
		return
	case 13:
		r.VidaUtilLote = nil
		return
	case 14:
		r.EntregaAntesDe = nil
		return
	case 15:
		r.ConsumoAntesDe = nil
		return
	case 16:
		r.Contramuestras = nil
		return
	case 17:
		r.EstadoOTAcondi = nil
		return
	case 18:
		r.EstadoOTTraza = nil
		return
	case 19:
		r.TipoAcondi = nil
		return
	case 20:
		r.TipoTraza = nil
		return
	case 21:
		r.ContratoServicioIngreso = nil
		return
	case 22:
		r.NomenclaturaContratoServicioIngreso = nil
		return
	case 23:
		r.DescripcionServicioIngreso = nil
		return
	case 24:
		r.TipoLineaMatriz = nil
		return
	case 25:
		r.CodConCalidad = nil
		return
	case 26:
		r.AccionConCalidad = nil
		return
	case 27:
		r.ResultadoConCalidad = nil
		return
	case 28:
		r.LineaExterna = nil
		return
	case 31:
		r.ValorDeclaradoLinea = nil
		return
	case 32:
		r.UnidadMedida = nil
		return
	case 33:
		r.LineaRecepcionWH = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Detalle) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Detalle) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Detalle) HintSize(int)                     { panic("Unsupported operation") }
func (_ Detalle) Finalize()                        {}

func (_ Detalle) AvroCRC64Fingerprint() []byte {
	return []byte(DetalleAvroCRC64Fingerprint)
}

func (r Detalle) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["SKU"], err = json.Marshal(r.SKU)
	if err != nil {
		return nil, err
	}
	output["UbicacionDestino"], err = json.Marshal(r.UbicacionDestino)
	if err != nil {
		return nil, err
	}
	output["LPNDestino"], err = json.Marshal(r.LPNDestino)
	if err != nil {
		return nil, err
	}
	output["PaqueteLote"], err = json.Marshal(r.PaqueteLote)
	if err != nil {
		return nil, err
	}
	output["LoteCajitaFabricante"], err = json.Marshal(r.LoteCajitaFabricante)
	if err != nil {
		return nil, err
	}
	output["LoteSecundario"], err = json.Marshal(r.LoteSecundario)
	if err != nil {
		return nil, err
	}
	output["FechaFabricacion"], err = json.Marshal(r.FechaFabricacion)
	if err != nil {
		return nil, err
	}
	output["FechaVencimiento"], err = json.Marshal(r.FechaVencimiento)
	if err != nil {
		return nil, err
	}
	output["ProductoTrazable"], err = json.Marshal(r.ProductoTrazable)
	if err != nil {
		return nil, err
	}
	output["AlmacenConsumo"], err = json.Marshal(r.AlmacenConsumo)
	if err != nil {
		return nil, err
	}
	output["EstadoLote"], err = json.Marshal(r.EstadoLote)
	if err != nil {
		return nil, err
	}
	output["BloqueoUbicacion"], err = json.Marshal(r.BloqueoUbicacion)
	if err != nil {
		return nil, err
	}
	output["VidaUtilLote"], err = json.Marshal(r.VidaUtilLote)
	if err != nil {
		return nil, err
	}
	output["EntregaAntesDe"], err = json.Marshal(r.EntregaAntesDe)
	if err != nil {
		return nil, err
	}
	output["ConsumoAntesDe"], err = json.Marshal(r.ConsumoAntesDe)
	if err != nil {
		return nil, err
	}
	output["Contramuestras"], err = json.Marshal(r.Contramuestras)
	if err != nil {
		return nil, err
	}
	output["EstadoOTAcondi"], err = json.Marshal(r.EstadoOTAcondi)
	if err != nil {
		return nil, err
	}
	output["EstadoOTTraza"], err = json.Marshal(r.EstadoOTTraza)
	if err != nil {
		return nil, err
	}
	output["TipoAcondi"], err = json.Marshal(r.TipoAcondi)
	if err != nil {
		return nil, err
	}
	output["TipoTraza"], err = json.Marshal(r.TipoTraza)
	if err != nil {
		return nil, err
	}
	output["ContratoServicioIngreso"], err = json.Marshal(r.ContratoServicioIngreso)
	if err != nil {
		return nil, err
	}
	output["NomenclaturaContratoServicioIngreso"], err = json.Marshal(r.NomenclaturaContratoServicioIngreso)
	if err != nil {
		return nil, err
	}
	output["DescripcionServicioIngreso"], err = json.Marshal(r.DescripcionServicioIngreso)
	if err != nil {
		return nil, err
	}
	output["TipoLineaMatriz"], err = json.Marshal(r.TipoLineaMatriz)
	if err != nil {
		return nil, err
	}
	output["CodConCalidad"], err = json.Marshal(r.CodConCalidad)
	if err != nil {
		return nil, err
	}
	output["AccionConCalidad"], err = json.Marshal(r.AccionConCalidad)
	if err != nil {
		return nil, err
	}
	output["ResultadoConCalidad"], err = json.Marshal(r.ResultadoConCalidad)
	if err != nil {
		return nil, err
	}
	output["LineaExterna"], err = json.Marshal(r.LineaExterna)
	if err != nil {
		return nil, err
	}
	output["CantEsperada"], err = json.Marshal(r.CantEsperada)
	if err != nil {
		return nil, err
	}
	output["CantRecibida"], err = json.Marshal(r.CantRecibida)
	if err != nil {
		return nil, err
	}
	output["ValorDeclaradoLinea"], err = json.Marshal(r.ValorDeclaradoLinea)
	if err != nil {
		return nil, err
	}
	output["UnidadMedida"], err = json.Marshal(r.UnidadMedida)
	if err != nil {
		return nil, err
	}
	output["LineaRecepcionWH"], err = json.Marshal(r.LineaRecepcionWH)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Detalle) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Propietario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Propietario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Propietario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SKU"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SKU); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SKU")
	}
	val = func() json.RawMessage {
		if v, ok := fields["UbicacionDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UbicacionDestino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for UbicacionDestino")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LPNDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LPNDestino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LPNDestino")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PaqueteLote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PaqueteLote); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PaqueteLote")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LoteCajitaFabricante"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteCajitaFabricante); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LoteCajitaFabricante")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LoteSecundario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteSecundario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LoteSecundario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaFabricacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaFabricacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaFabricacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaVencimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaVencimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaVencimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ProductoTrazable"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ProductoTrazable); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ProductoTrazable")
	}
	val = func() json.RawMessage {
		if v, ok := fields["AlmacenConsumo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AlmacenConsumo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for AlmacenConsumo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EstadoLote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoLote); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EstadoLote")
	}
	val = func() json.RawMessage {
		if v, ok := fields["BloqueoUbicacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BloqueoUbicacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for BloqueoUbicacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["VidaUtilLote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VidaUtilLote); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for VidaUtilLote")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EntregaAntesDe"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EntregaAntesDe); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EntregaAntesDe")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ConsumoAntesDe"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ConsumoAntesDe); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ConsumoAntesDe")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Contramuestras"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contramuestras); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Contramuestras")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EstadoOTAcondi"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoOTAcondi); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EstadoOTAcondi")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EstadoOTTraza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoOTTraza); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EstadoOTTraza")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoAcondi"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoAcondi); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoAcondi")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoTraza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoTraza); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoTraza")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ContratoServicioIngreso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContratoServicioIngreso); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ContratoServicioIngreso")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NomenclaturaContratoServicioIngreso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NomenclaturaContratoServicioIngreso); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NomenclaturaContratoServicioIngreso")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DescripcionServicioIngreso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DescripcionServicioIngreso); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DescripcionServicioIngreso")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoLineaMatriz"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoLineaMatriz); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoLineaMatriz")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodConCalidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodConCalidad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodConCalidad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["AccionConCalidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AccionConCalidad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for AccionConCalidad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ResultadoConCalidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ResultadoConCalidad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ResultadoConCalidad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LineaExterna"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LineaExterna); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LineaExterna")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CantEsperada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantEsperada); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CantEsperada")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CantRecibida"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantRecibida); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CantRecibida")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ValorDeclaradoLinea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ValorDeclaradoLinea); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ValorDeclaradoLinea")
	}
	val = func() json.RawMessage {
		if v, ok := fields["UnidadMedida"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UnidadMedida); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for UnidadMedida")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LineaRecepcionWH"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LineaRecepcionWH); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LineaRecepcionWH")
	}
	return nil
}
