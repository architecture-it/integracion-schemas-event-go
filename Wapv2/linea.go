// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListaDeDetalleDeLinea.avsc
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

type Linea struct {
	IdAndreani *UnionNullString `json:"idAndreani"`

	NumeroDeLinea string `json:"numeroDeLinea"`

	NumeroDeLineaWMS *UnionNullString `json:"numeroDeLineaWMS"`

	EstatusOTdeTRazaporLPN *UnionNullString `json:"estatusOTdeTRazaporLPN"`

	CantidadPedida int32 `json:"cantidadPedida"`

	ValorDeclarado string `json:"valorDeclarado"`

	CodigoActualizacion *UnionNullString `json:"codigoActualizacion"`

	UnidadMedida string `json:"unidadMedida"`

	Propietario string `json:"propietario"`

	ContratoWhAbastecimiento *UnionNullString `json:"contratoWhAbastecimiento"`

	TipoAcondicionamiento *UnionNullArrayString `json:"tipoAcondicionamiento"`

	TipoTraza *UnionNullString `json:"tipoTraza"`

	TipoControlCalidad *UnionNullArrayString `json:"tipoControlCalidad"`

	AccionControlCalidad *UnionNullArrayString `json:"accionControlCalidad"`

	UbicacionPredeterminada *UnionNullString `json:"ubicacionPredeterminada"`

	Almacen string `json:"almacen"`

	VidaUtilLote *UnionNullString `json:"vidaUtilLote"`

	AvisoContramuestra *UnionNullString `json:"avisoContramuestra"`

	OtrosDatos *UnionNullListaDePropiedades `json:"otrosDatos"`

	Articulo ArticuloAsn `json:"articulo"`

	EstatusOTdeAcondporLPN *UnionNullString `json:"estatusOTdeAcondporLPN"`

	CodigoDeTrazaANMAT *UnionNullString `json:"codigoDeTrazaANMAT"`

	ProductoTrazable *UnionNullString `json:"productoTrazable"`
}

const LineaAvroCRC64Fingerprint = "ؤ\x80OEF\xdau"

func NewLinea() Linea {
	r := Linea{}
	r.IdAndreani = nil
	r.NumeroDeLineaWMS = nil
	r.EstatusOTdeTRazaporLPN = nil
	r.CodigoActualizacion = nil
	r.TipoAcondicionamiento = nil
	r.TipoTraza = nil
	r.TipoControlCalidad = nil
	r.AccionControlCalidad = nil
	r.UbicacionPredeterminada = nil
	r.VidaUtilLote = nil
	r.AvisoContramuestra = nil
	r.OtrosDatos = nil
	r.Articulo = NewArticuloAsn()

	r.EstatusOTdeAcondporLPN = nil
	r.CodigoDeTrazaANMAT = nil
	r.ProductoTrazable = nil
	return r
}

func DeserializeLinea(r io.Reader) (Linea, error) {
	t := NewLinea()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeLineaFromSchema(r io.Reader, schema string) (Linea, error) {
	t := NewLinea()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeLinea(r Linea, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.IdAndreani, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroDeLinea, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeLineaWMS, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstatusOTdeTRazaporLPN, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.CantidadPedida, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ValorDeclarado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoActualizacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.UnidadMedida, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ContratoWhAbastecimiento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayString(r.TipoAcondicionamiento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoTraza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayString(r.TipoControlCalidad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayString(r.AccionControlCalidad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.UbicacionPredeterminada, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Almacen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.VidaUtilLote, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.AvisoContramuestra, w)
	if err != nil {
		return err
	}
	err = writeUnionNullListaDePropiedades(r.OtrosDatos, w)
	if err != nil {
		return err
	}
	err = writeArticuloAsn(r.Articulo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstatusOTdeAcondporLPN, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoDeTrazaANMAT, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ProductoTrazable, w)
	if err != nil {
		return err
	}
	return err
}

func (r Linea) Serialize(w io.Writer) error {
	return writeLinea(r, w)
}

func (r Linea) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"idAndreani\",\"type\":[\"null\",\"string\"]},{\"name\":\"numeroDeLinea\",\"type\":\"string\"},{\"default\":null,\"name\":\"numeroDeLineaWMS\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estatusOTdeTRazaporLPN\",\"type\":[\"null\",\"string\"]},{\"name\":\"cantidadPedida\",\"type\":\"int\"},{\"name\":\"valorDeclarado\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoActualizacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"unidadMedida\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"contratoWhAbastecimiento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoAcondicionamiento\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"tipoTraza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoControlCalidad\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"accionControlCalidad\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"ubicacionPredeterminada\",\"type\":[\"null\",\"string\"]},{\"name\":\"almacen\",\"type\":\"string\"},{\"default\":null,\"name\":\"vidaUtilLote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"avisoContramuestra\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"name\":\"articulo\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"ean\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"camposLibres\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.Metadato\"]},{\"name\":\"lote\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaDeVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"Andreani.Wapv2.Events.Record.ListaDePropiedades\"]}],\"name\":\"LoteAsn\",\"type\":\"record\"}},{\"default\":null,\"name\":\"ubicacionDestino\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"paqueteLote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"prodTrazable\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"consumirAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"entrarAntesDe\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"nomenclaturaServicio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"descripcionServicio\",\"type\":[\"null\",\"string\"]}],\"name\":\"ArticuloAsn\",\"type\":\"record\"}},{\"default\":null,\"name\":\"estatusOTdeAcondporLPN\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoDeTrazaANMAT\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"productoTrazable\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.Wapv2.Events.Record.Linea\",\"type\":\"record\"}"
}

func (r Linea) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.Linea"
}

func (_ Linea) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Linea) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Linea) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Linea) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Linea) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Linea) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Linea) SetString(v string)   { panic("Unsupported operation") }
func (_ Linea) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Linea) Get(i int) types.Field {
	switch i {
	case 0:
		r.IdAndreani = NewUnionNullString()

		return r.IdAndreani
	case 1:
		w := types.String{Target: &r.NumeroDeLinea}

		return w

	case 2:
		r.NumeroDeLineaWMS = NewUnionNullString()

		return r.NumeroDeLineaWMS
	case 3:
		r.EstatusOTdeTRazaporLPN = NewUnionNullString()

		return r.EstatusOTdeTRazaporLPN
	case 4:
		w := types.Int{Target: &r.CantidadPedida}

		return w

	case 5:
		w := types.String{Target: &r.ValorDeclarado}

		return w

	case 6:
		r.CodigoActualizacion = NewUnionNullString()

		return r.CodigoActualizacion
	case 7:
		w := types.String{Target: &r.UnidadMedida}

		return w

	case 8:
		w := types.String{Target: &r.Propietario}

		return w

	case 9:
		r.ContratoWhAbastecimiento = NewUnionNullString()

		return r.ContratoWhAbastecimiento
	case 10:
		r.TipoAcondicionamiento = NewUnionNullArrayString()

		return r.TipoAcondicionamiento
	case 11:
		r.TipoTraza = NewUnionNullString()

		return r.TipoTraza
	case 12:
		r.TipoControlCalidad = NewUnionNullArrayString()

		return r.TipoControlCalidad
	case 13:
		r.AccionControlCalidad = NewUnionNullArrayString()

		return r.AccionControlCalidad
	case 14:
		r.UbicacionPredeterminada = NewUnionNullString()

		return r.UbicacionPredeterminada
	case 15:
		w := types.String{Target: &r.Almacen}

		return w

	case 16:
		r.VidaUtilLote = NewUnionNullString()

		return r.VidaUtilLote
	case 17:
		r.AvisoContramuestra = NewUnionNullString()

		return r.AvisoContramuestra
	case 18:
		r.OtrosDatos = NewUnionNullListaDePropiedades()

		return r.OtrosDatos
	case 19:
		r.Articulo = NewArticuloAsn()

		w := types.Record{Target: &r.Articulo}

		return w

	case 20:
		r.EstatusOTdeAcondporLPN = NewUnionNullString()

		return r.EstatusOTdeAcondporLPN
	case 21:
		r.CodigoDeTrazaANMAT = NewUnionNullString()

		return r.CodigoDeTrazaANMAT
	case 22:
		r.ProductoTrazable = NewUnionNullString()

		return r.ProductoTrazable
	}
	panic("Unknown field index")
}

func (r *Linea) SetDefault(i int) {
	switch i {
	case 0:
		r.IdAndreani = nil
		return
	case 2:
		r.NumeroDeLineaWMS = nil
		return
	case 3:
		r.EstatusOTdeTRazaporLPN = nil
		return
	case 6:
		r.CodigoActualizacion = nil
		return
	case 10:
		r.TipoAcondicionamiento = nil
		return
	case 11:
		r.TipoTraza = nil
		return
	case 12:
		r.TipoControlCalidad = nil
		return
	case 13:
		r.AccionControlCalidad = nil
		return
	case 14:
		r.UbicacionPredeterminada = nil
		return
	case 16:
		r.VidaUtilLote = nil
		return
	case 17:
		r.AvisoContramuestra = nil
		return
	case 18:
		r.OtrosDatos = nil
		return
	case 20:
		r.EstatusOTdeAcondporLPN = nil
		return
	case 21:
		r.CodigoDeTrazaANMAT = nil
		return
	case 22:
		r.ProductoTrazable = nil
		return
	}
	panic("Unknown field index")
}

func (r *Linea) NullField(i int) {
	switch i {
	case 0:
		r.IdAndreani = nil
		return
	case 2:
		r.NumeroDeLineaWMS = nil
		return
	case 3:
		r.EstatusOTdeTRazaporLPN = nil
		return
	case 6:
		r.CodigoActualizacion = nil
		return
	case 9:
		r.ContratoWhAbastecimiento = nil
		return
	case 10:
		r.TipoAcondicionamiento = nil
		return
	case 11:
		r.TipoTraza = nil
		return
	case 12:
		r.TipoControlCalidad = nil
		return
	case 13:
		r.AccionControlCalidad = nil
		return
	case 14:
		r.UbicacionPredeterminada = nil
		return
	case 16:
		r.VidaUtilLote = nil
		return
	case 17:
		r.AvisoContramuestra = nil
		return
	case 18:
		r.OtrosDatos = nil
		return
	case 20:
		r.EstatusOTdeAcondporLPN = nil
		return
	case 21:
		r.CodigoDeTrazaANMAT = nil
		return
	case 22:
		r.ProductoTrazable = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Linea) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Linea) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Linea) HintSize(int)                     { panic("Unsupported operation") }
func (_ Linea) Finalize()                        {}

func (_ Linea) AvroCRC64Fingerprint() []byte {
	return []byte(LineaAvroCRC64Fingerprint)
}

func (r Linea) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["idAndreani"], err = json.Marshal(r.IdAndreani)
	if err != nil {
		return nil, err
	}
	output["numeroDeLinea"], err = json.Marshal(r.NumeroDeLinea)
	if err != nil {
		return nil, err
	}
	output["numeroDeLineaWMS"], err = json.Marshal(r.NumeroDeLineaWMS)
	if err != nil {
		return nil, err
	}
	output["estatusOTdeTRazaporLPN"], err = json.Marshal(r.EstatusOTdeTRazaporLPN)
	if err != nil {
		return nil, err
	}
	output["cantidadPedida"], err = json.Marshal(r.CantidadPedida)
	if err != nil {
		return nil, err
	}
	output["valorDeclarado"], err = json.Marshal(r.ValorDeclarado)
	if err != nil {
		return nil, err
	}
	output["codigoActualizacion"], err = json.Marshal(r.CodigoActualizacion)
	if err != nil {
		return nil, err
	}
	output["unidadMedida"], err = json.Marshal(r.UnidadMedida)
	if err != nil {
		return nil, err
	}
	output["propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["contratoWhAbastecimiento"], err = json.Marshal(r.ContratoWhAbastecimiento)
	if err != nil {
		return nil, err
	}
	output["tipoAcondicionamiento"], err = json.Marshal(r.TipoAcondicionamiento)
	if err != nil {
		return nil, err
	}
	output["tipoTraza"], err = json.Marshal(r.TipoTraza)
	if err != nil {
		return nil, err
	}
	output["tipoControlCalidad"], err = json.Marshal(r.TipoControlCalidad)
	if err != nil {
		return nil, err
	}
	output["accionControlCalidad"], err = json.Marshal(r.AccionControlCalidad)
	if err != nil {
		return nil, err
	}
	output["ubicacionPredeterminada"], err = json.Marshal(r.UbicacionPredeterminada)
	if err != nil {
		return nil, err
	}
	output["almacen"], err = json.Marshal(r.Almacen)
	if err != nil {
		return nil, err
	}
	output["vidaUtilLote"], err = json.Marshal(r.VidaUtilLote)
	if err != nil {
		return nil, err
	}
	output["avisoContramuestra"], err = json.Marshal(r.AvisoContramuestra)
	if err != nil {
		return nil, err
	}
	output["otrosDatos"], err = json.Marshal(r.OtrosDatos)
	if err != nil {
		return nil, err
	}
	output["articulo"], err = json.Marshal(r.Articulo)
	if err != nil {
		return nil, err
	}
	output["estatusOTdeAcondporLPN"], err = json.Marshal(r.EstatusOTdeAcondporLPN)
	if err != nil {
		return nil, err
	}
	output["codigoDeTrazaANMAT"], err = json.Marshal(r.CodigoDeTrazaANMAT)
	if err != nil {
		return nil, err
	}
	output["productoTrazable"], err = json.Marshal(r.ProductoTrazable)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Linea) UnmarshalJSON(data []byte) error {
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
		r.IdAndreani = NewUnionNullString()

		r.IdAndreani = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroDeLinea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeLinea); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroDeLinea")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroDeLineaWMS"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeLineaWMS); err != nil {
			return err
		}
	} else {
		r.NumeroDeLineaWMS = NewUnionNullString()

		r.NumeroDeLineaWMS = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["estatusOTdeTRazaporLPN"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstatusOTdeTRazaporLPN); err != nil {
			return err
		}
	} else {
		r.EstatusOTdeTRazaporLPN = NewUnionNullString()

		r.EstatusOTdeTRazaporLPN = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["cantidadPedida"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadPedida); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cantidadPedida")
	}
	val = func() json.RawMessage {
		if v, ok := fields["valorDeclarado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ValorDeclarado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for valorDeclarado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoActualizacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoActualizacion); err != nil {
			return err
		}
	} else {
		r.CodigoActualizacion = NewUnionNullString()

		r.CodigoActualizacion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["unidadMedida"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UnidadMedida); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for unidadMedida")
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
		if v, ok := fields["contratoWhAbastecimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContratoWhAbastecimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for contratoWhAbastecimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoAcondicionamiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoAcondicionamiento); err != nil {
			return err
		}
	} else {
		r.TipoAcondicionamiento = NewUnionNullArrayString()

		r.TipoAcondicionamiento = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoTraza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoTraza); err != nil {
			return err
		}
	} else {
		r.TipoTraza = NewUnionNullString()

		r.TipoTraza = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoControlCalidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoControlCalidad); err != nil {
			return err
		}
	} else {
		r.TipoControlCalidad = NewUnionNullArrayString()

		r.TipoControlCalidad = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["accionControlCalidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AccionControlCalidad); err != nil {
			return err
		}
	} else {
		r.AccionControlCalidad = NewUnionNullArrayString()

		r.AccionControlCalidad = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ubicacionPredeterminada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UbicacionPredeterminada); err != nil {
			return err
		}
	} else {
		r.UbicacionPredeterminada = NewUnionNullString()

		r.UbicacionPredeterminada = nil
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
		if v, ok := fields["vidaUtilLote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VidaUtilLote); err != nil {
			return err
		}
	} else {
		r.VidaUtilLote = NewUnionNullString()

		r.VidaUtilLote = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["avisoContramuestra"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AvisoContramuestra); err != nil {
			return err
		}
	} else {
		r.AvisoContramuestra = NewUnionNullString()

		r.AvisoContramuestra = nil
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
		r.OtrosDatos = NewUnionNullListaDePropiedades()

		r.OtrosDatos = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["articulo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Articulo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for articulo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["estatusOTdeAcondporLPN"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstatusOTdeAcondporLPN); err != nil {
			return err
		}
	} else {
		r.EstatusOTdeAcondporLPN = NewUnionNullString()

		r.EstatusOTdeAcondporLPN = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoDeTrazaANMAT"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeTrazaANMAT); err != nil {
			return err
		}
	} else {
		r.CodigoDeTrazaANMAT = NewUnionNullString()

		r.CodigoDeTrazaANMAT = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["productoTrazable"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ProductoTrazable); err != nil {
			return err
		}
	} else {
		r.ProductoTrazable = NewUnionNullString()

		r.ProductoTrazable = nil
	}
	return nil
}
