// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoWhPropietarioEliminado.avsc
 */
package EventoPropietarioEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Cabecera struct {
	Propietario string `json:"Propietario"`

	Compañía *UnionNullString `json:"Compañía"`

	Descripción *UnionNullString `json:"Descripción"`

	IVA *UnionNullString `json:"IVA"`

	SeguimientoIntercambioPallet *UnionNullString `json:"SeguimientoIntercambioPallet"`

	DestinatarioCalle *UnionNullString `json:"DestinatarioCalle"`

	DestinatarioNumero *UnionNullString `json:"DestinatarioNumero"`

	DestinatarioPiso *UnionNullString `json:"DestinatarioPiso"`

	DestinatarioDepartamento *UnionNullString `json:"DestinatarioDepartamento"`

	Ciudad *UnionNullString `json:"Ciudad"`

	Provincia *UnionNullString `json:"Provincia"`

	CP *UnionNullString `json:"CP"`

	Pais *UnionNullString `json:"Pais"`

	ISOPais *UnionNullString `json:"ISOPais"`

	Contacto1 *UnionNullString `json:"Contacto1"`

	Telefono *UnionNullString `json:"Telefono"`

	Mail *UnionNullString `json:"Mail"`

	Fax *UnionNullString `json:"Fax"`

	DestinatarioCUIT *UnionNullString `json:"DestinatarioCUIT"`

	DestinatarioGLN *UnionNullString `json:"DestinatarioGLN"`

	DestinatarioAgente *UnionNullString `json:"DestinatarioAgente"`

	DestinatarioRamo *UnionNullString `json:"DestinatarioRamo"`

	Grilla *UnionNullString `json:"Grilla"`

	DiaTransmion *UnionNullString `json:"DiaTransmion"`

	EntregaLunes *UnionNullString `json:"EntregaLunes"`

	EntregaMartes *UnionNullString `json:"EntregaMartes"`

	EntregaMiercoles *UnionNullString `json:"EntregaMiercoles"`

	EntregaJueves *UnionNullString `json:"EntregaJueves"`

	EntregaViernes *UnionNullString `json:"EntregaViernes"`

	EntregaSabado *UnionNullString `json:"EntregaSabado"`

	EntregaDomingo *UnionNullString `json:"EntregaDomingo"`

	EntregaPostFeriado *UnionNullString `json:"EntregaPostFeriado"`

	Geolocalizacion *UnionNullString `json:"Geolocalizacion"`

	Latitud *UnionNullString `json:"Latitud"`

	Longitud *UnionNullString `json:"Longitud"`

	TipoIntegracion *UnionNullString `json:"TipoIntegracion"`

	Planta *UnionNullString `json:"Planta"`

	Nave *UnionNullString `json:"Nave"`

	ClienteSap *UnionNullString `json:"ClienteSap"`

	CodLevantamientoCuarentena *UnionNullString `json:"CodLevantamientoCuarentena"`

	WosTrazabilidad *UnionNullString `json:"WosTrazabilidad"`

	TipoEncolado *UnionNullString `json:"TipoEncolado"`

	ProcesosDondeInformaAnmat *UnionNullString `json:"ProcesosDondeInformaAnmat"`

	AcondiExterno *UnionNullString `json:"AcondiExterno"`

	InformaLevantamientoCuarentena *UnionNullString `json:"InformaLevantamientoCuarentena"`

	ValidaRemito *UnionNullString `json:"ValidaRemito"`

	CodAnmatAbastecimiento *UnionNullString `json:"CodAnmatAbastecimiento"`

	AbastecimientoSeriesCliente *UnionNullString `json:"AbastecimientoSeriesCliente"`

	TipoReceptorAnmat *UnionNullString `json:"TipoReceptorAnmat"`

	Tipo string `json:"Tipo"`

	AndreaniCOM *UnionNullString `json:"AndreaniCOM"`
}

const CabeceraAvroCRC64Fingerprint = "F'\x7f\xfd\xd7\xec\x8b\xc2"

func NewCabecera() Cabecera {
	r := Cabecera{}
	return r
}

func DeserializeCabecera(r io.Reader) (Cabecera, error) {
	t := NewCabecera()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCabeceraFromSchema(r io.Reader, schema string) (Cabecera, error) {
	t := NewCabecera()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCabecera(r Cabecera, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Compañía, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Descripción, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.IVA, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.SeguimientoIntercambioPallet, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioCalle, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioNumero, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioPiso, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioDepartamento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Ciudad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Provincia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CP, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Pais, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ISOPais, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Contacto1, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Telefono, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Mail, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Fax, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioCUIT, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioGLN, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioAgente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DestinatarioRamo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Grilla, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DiaTransmion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EntregaLunes, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EntregaMartes, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EntregaMiercoles, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EntregaJueves, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EntregaViernes, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EntregaSabado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EntregaDomingo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EntregaPostFeriado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Geolocalizacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Latitud, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Longitud, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoIntegracion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Planta, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Nave, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ClienteSap, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodLevantamientoCuarentena, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.WosTrazabilidad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoEncolado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ProcesosDondeInformaAnmat, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.AcondiExterno, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.InformaLevantamientoCuarentena, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ValidaRemito, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodAnmatAbastecimiento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.AbastecimientoSeriesCliente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoReceptorAnmat, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Tipo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.AndreaniCOM, w)
	if err != nil {
		return err
	}
	return err
}

func (r Cabecera) Serialize(w io.Writer) error {
	return writeCabecera(r, w)
}

func (r Cabecera) Schema() string {
	return "{\"fields\":[{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Compañía\",\"type\":[\"null\",\"string\"]},{\"name\":\"Descripción\",\"type\":[\"null\",\"string\"]},{\"name\":\"IVA\",\"type\":[\"null\",\"string\"]},{\"name\":\"SeguimientoIntercambioPallet\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioCalle\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioNumero\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioPiso\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioDepartamento\",\"type\":[\"null\",\"string\"]},{\"name\":\"Ciudad\",\"type\":[\"null\",\"string\"]},{\"name\":\"Provincia\",\"type\":[\"null\",\"string\"]},{\"name\":\"CP\",\"type\":[\"null\",\"string\"]},{\"name\":\"Pais\",\"type\":[\"null\",\"string\"]},{\"name\":\"ISOPais\",\"type\":[\"null\",\"string\"]},{\"name\":\"Contacto1\",\"type\":[\"null\",\"string\"]},{\"name\":\"Telefono\",\"type\":[\"null\",\"string\"]},{\"name\":\"Mail\",\"type\":[\"null\",\"string\"]},{\"name\":\"Fax\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioCUIT\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioGLN\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioAgente\",\"type\":[\"null\",\"string\"]},{\"name\":\"DestinatarioRamo\",\"type\":[\"null\",\"string\"]},{\"name\":\"Grilla\",\"type\":[\"null\",\"string\"]},{\"name\":\"DiaTransmion\",\"type\":[\"null\",\"string\"]},{\"name\":\"EntregaLunes\",\"type\":[\"null\",\"string\"]},{\"name\":\"EntregaMartes\",\"type\":[\"null\",\"string\"]},{\"name\":\"EntregaMiercoles\",\"type\":[\"null\",\"string\"]},{\"name\":\"EntregaJueves\",\"type\":[\"null\",\"string\"]},{\"name\":\"EntregaViernes\",\"type\":[\"null\",\"string\"]},{\"name\":\"EntregaSabado\",\"type\":[\"null\",\"string\"]},{\"name\":\"EntregaDomingo\",\"type\":[\"null\",\"string\"]},{\"name\":\"EntregaPostFeriado\",\"type\":[\"null\",\"string\"]},{\"name\":\"Geolocalizacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"Latitud\",\"type\":[\"null\",\"string\"]},{\"name\":\"Longitud\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoIntegracion\",\"type\":[\"null\",\"string\"]},{\"name\":\"Planta\",\"type\":[\"null\",\"string\"]},{\"name\":\"Nave\",\"type\":[\"null\",\"string\"]},{\"name\":\"ClienteSap\",\"type\":[\"null\",\"string\"]},{\"name\":\"CodLevantamientoCuarentena\",\"type\":[\"null\",\"string\"]},{\"name\":\"WosTrazabilidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoEncolado\",\"type\":[\"null\",\"string\"]},{\"name\":\"ProcesosDondeInformaAnmat\",\"type\":[\"null\",\"string\"]},{\"name\":\"AcondiExterno\",\"type\":[\"null\",\"string\"]},{\"name\":\"InformaLevantamientoCuarentena\",\"type\":[\"null\",\"string\"]},{\"name\":\"ValidaRemito\",\"type\":[\"null\",\"string\"]},{\"name\":\"CodAnmatAbastecimiento\",\"type\":[\"null\",\"string\"]},{\"name\":\"AbastecimientoSeriesCliente\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoReceptorAnmat\",\"type\":[\"null\",\"string\"]},{\"name\":\"Tipo\",\"type\":\"string\"},{\"name\":\"AndreaniCOM\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.EventoPropietario.Events.EliminadoCommon.Cabecera\",\"type\":\"record\"}"
}

func (r Cabecera) SchemaName() string {
	return "Andreani.EventoPropietario.Events.EliminadoCommon.Cabecera"
}

func (_ Cabecera) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Cabecera) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Cabecera) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Cabecera) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Cabecera) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Cabecera) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Cabecera) SetString(v string)   { panic("Unsupported operation") }
func (_ Cabecera) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Cabecera) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Propietario}

		return w

	case 1:
		r.Compañía = NewUnionNullString()

		return r.Compañía
	case 2:
		r.Descripción = NewUnionNullString()

		return r.Descripción
	case 3:
		r.IVA = NewUnionNullString()

		return r.IVA
	case 4:
		r.SeguimientoIntercambioPallet = NewUnionNullString()

		return r.SeguimientoIntercambioPallet
	case 5:
		r.DestinatarioCalle = NewUnionNullString()

		return r.DestinatarioCalle
	case 6:
		r.DestinatarioNumero = NewUnionNullString()

		return r.DestinatarioNumero
	case 7:
		r.DestinatarioPiso = NewUnionNullString()

		return r.DestinatarioPiso
	case 8:
		r.DestinatarioDepartamento = NewUnionNullString()

		return r.DestinatarioDepartamento
	case 9:
		r.Ciudad = NewUnionNullString()

		return r.Ciudad
	case 10:
		r.Provincia = NewUnionNullString()

		return r.Provincia
	case 11:
		r.CP = NewUnionNullString()

		return r.CP
	case 12:
		r.Pais = NewUnionNullString()

		return r.Pais
	case 13:
		r.ISOPais = NewUnionNullString()

		return r.ISOPais
	case 14:
		r.Contacto1 = NewUnionNullString()

		return r.Contacto1
	case 15:
		r.Telefono = NewUnionNullString()

		return r.Telefono
	case 16:
		r.Mail = NewUnionNullString()

		return r.Mail
	case 17:
		r.Fax = NewUnionNullString()

		return r.Fax
	case 18:
		r.DestinatarioCUIT = NewUnionNullString()

		return r.DestinatarioCUIT
	case 19:
		r.DestinatarioGLN = NewUnionNullString()

		return r.DestinatarioGLN
	case 20:
		r.DestinatarioAgente = NewUnionNullString()

		return r.DestinatarioAgente
	case 21:
		r.DestinatarioRamo = NewUnionNullString()

		return r.DestinatarioRamo
	case 22:
		r.Grilla = NewUnionNullString()

		return r.Grilla
	case 23:
		r.DiaTransmion = NewUnionNullString()

		return r.DiaTransmion
	case 24:
		r.EntregaLunes = NewUnionNullString()

		return r.EntregaLunes
	case 25:
		r.EntregaMartes = NewUnionNullString()

		return r.EntregaMartes
	case 26:
		r.EntregaMiercoles = NewUnionNullString()

		return r.EntregaMiercoles
	case 27:
		r.EntregaJueves = NewUnionNullString()

		return r.EntregaJueves
	case 28:
		r.EntregaViernes = NewUnionNullString()

		return r.EntregaViernes
	case 29:
		r.EntregaSabado = NewUnionNullString()

		return r.EntregaSabado
	case 30:
		r.EntregaDomingo = NewUnionNullString()

		return r.EntregaDomingo
	case 31:
		r.EntregaPostFeriado = NewUnionNullString()

		return r.EntregaPostFeriado
	case 32:
		r.Geolocalizacion = NewUnionNullString()

		return r.Geolocalizacion
	case 33:
		r.Latitud = NewUnionNullString()

		return r.Latitud
	case 34:
		r.Longitud = NewUnionNullString()

		return r.Longitud
	case 35:
		r.TipoIntegracion = NewUnionNullString()

		return r.TipoIntegracion
	case 36:
		r.Planta = NewUnionNullString()

		return r.Planta
	case 37:
		r.Nave = NewUnionNullString()

		return r.Nave
	case 38:
		r.ClienteSap = NewUnionNullString()

		return r.ClienteSap
	case 39:
		r.CodLevantamientoCuarentena = NewUnionNullString()

		return r.CodLevantamientoCuarentena
	case 40:
		r.WosTrazabilidad = NewUnionNullString()

		return r.WosTrazabilidad
	case 41:
		r.TipoEncolado = NewUnionNullString()

		return r.TipoEncolado
	case 42:
		r.ProcesosDondeInformaAnmat = NewUnionNullString()

		return r.ProcesosDondeInformaAnmat
	case 43:
		r.AcondiExterno = NewUnionNullString()

		return r.AcondiExterno
	case 44:
		r.InformaLevantamientoCuarentena = NewUnionNullString()

		return r.InformaLevantamientoCuarentena
	case 45:
		r.ValidaRemito = NewUnionNullString()

		return r.ValidaRemito
	case 46:
		r.CodAnmatAbastecimiento = NewUnionNullString()

		return r.CodAnmatAbastecimiento
	case 47:
		r.AbastecimientoSeriesCliente = NewUnionNullString()

		return r.AbastecimientoSeriesCliente
	case 48:
		r.TipoReceptorAnmat = NewUnionNullString()

		return r.TipoReceptorAnmat
	case 49:
		w := types.String{Target: &r.Tipo}

		return w

	case 50:
		r.AndreaniCOM = NewUnionNullString()

		return r.AndreaniCOM
	}
	panic("Unknown field index")
}

func (r *Cabecera) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Cabecera) NullField(i int) {
	switch i {
	case 1:
		r.Compañía = nil
		return
	case 2:
		r.Descripción = nil
		return
	case 3:
		r.IVA = nil
		return
	case 4:
		r.SeguimientoIntercambioPallet = nil
		return
	case 5:
		r.DestinatarioCalle = nil
		return
	case 6:
		r.DestinatarioNumero = nil
		return
	case 7:
		r.DestinatarioPiso = nil
		return
	case 8:
		r.DestinatarioDepartamento = nil
		return
	case 9:
		r.Ciudad = nil
		return
	case 10:
		r.Provincia = nil
		return
	case 11:
		r.CP = nil
		return
	case 12:
		r.Pais = nil
		return
	case 13:
		r.ISOPais = nil
		return
	case 14:
		r.Contacto1 = nil
		return
	case 15:
		r.Telefono = nil
		return
	case 16:
		r.Mail = nil
		return
	case 17:
		r.Fax = nil
		return
	case 18:
		r.DestinatarioCUIT = nil
		return
	case 19:
		r.DestinatarioGLN = nil
		return
	case 20:
		r.DestinatarioAgente = nil
		return
	case 21:
		r.DestinatarioRamo = nil
		return
	case 22:
		r.Grilla = nil
		return
	case 23:
		r.DiaTransmion = nil
		return
	case 24:
		r.EntregaLunes = nil
		return
	case 25:
		r.EntregaMartes = nil
		return
	case 26:
		r.EntregaMiercoles = nil
		return
	case 27:
		r.EntregaJueves = nil
		return
	case 28:
		r.EntregaViernes = nil
		return
	case 29:
		r.EntregaSabado = nil
		return
	case 30:
		r.EntregaDomingo = nil
		return
	case 31:
		r.EntregaPostFeriado = nil
		return
	case 32:
		r.Geolocalizacion = nil
		return
	case 33:
		r.Latitud = nil
		return
	case 34:
		r.Longitud = nil
		return
	case 35:
		r.TipoIntegracion = nil
		return
	case 36:
		r.Planta = nil
		return
	case 37:
		r.Nave = nil
		return
	case 38:
		r.ClienteSap = nil
		return
	case 39:
		r.CodLevantamientoCuarentena = nil
		return
	case 40:
		r.WosTrazabilidad = nil
		return
	case 41:
		r.TipoEncolado = nil
		return
	case 42:
		r.ProcesosDondeInformaAnmat = nil
		return
	case 43:
		r.AcondiExterno = nil
		return
	case 44:
		r.InformaLevantamientoCuarentena = nil
		return
	case 45:
		r.ValidaRemito = nil
		return
	case 46:
		r.CodAnmatAbastecimiento = nil
		return
	case 47:
		r.AbastecimientoSeriesCliente = nil
		return
	case 48:
		r.TipoReceptorAnmat = nil
		return
	case 50:
		r.AndreaniCOM = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Cabecera) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Cabecera) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Cabecera) HintSize(int)                     { panic("Unsupported operation") }
func (_ Cabecera) Finalize()                        {}

func (_ Cabecera) AvroCRC64Fingerprint() []byte {
	return []byte(CabeceraAvroCRC64Fingerprint)
}

func (r Cabecera) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["Compañía"], err = json.Marshal(r.Compañía)
	if err != nil {
		return nil, err
	}
	output["Descripción"], err = json.Marshal(r.Descripción)
	if err != nil {
		return nil, err
	}
	output["IVA"], err = json.Marshal(r.IVA)
	if err != nil {
		return nil, err
	}
	output["SeguimientoIntercambioPallet"], err = json.Marshal(r.SeguimientoIntercambioPallet)
	if err != nil {
		return nil, err
	}
	output["DestinatarioCalle"], err = json.Marshal(r.DestinatarioCalle)
	if err != nil {
		return nil, err
	}
	output["DestinatarioNumero"], err = json.Marshal(r.DestinatarioNumero)
	if err != nil {
		return nil, err
	}
	output["DestinatarioPiso"], err = json.Marshal(r.DestinatarioPiso)
	if err != nil {
		return nil, err
	}
	output["DestinatarioDepartamento"], err = json.Marshal(r.DestinatarioDepartamento)
	if err != nil {
		return nil, err
	}
	output["Ciudad"], err = json.Marshal(r.Ciudad)
	if err != nil {
		return nil, err
	}
	output["Provincia"], err = json.Marshal(r.Provincia)
	if err != nil {
		return nil, err
	}
	output["CP"], err = json.Marshal(r.CP)
	if err != nil {
		return nil, err
	}
	output["Pais"], err = json.Marshal(r.Pais)
	if err != nil {
		return nil, err
	}
	output["ISOPais"], err = json.Marshal(r.ISOPais)
	if err != nil {
		return nil, err
	}
	output["Contacto1"], err = json.Marshal(r.Contacto1)
	if err != nil {
		return nil, err
	}
	output["Telefono"], err = json.Marshal(r.Telefono)
	if err != nil {
		return nil, err
	}
	output["Mail"], err = json.Marshal(r.Mail)
	if err != nil {
		return nil, err
	}
	output["Fax"], err = json.Marshal(r.Fax)
	if err != nil {
		return nil, err
	}
	output["DestinatarioCUIT"], err = json.Marshal(r.DestinatarioCUIT)
	if err != nil {
		return nil, err
	}
	output["DestinatarioGLN"], err = json.Marshal(r.DestinatarioGLN)
	if err != nil {
		return nil, err
	}
	output["DestinatarioAgente"], err = json.Marshal(r.DestinatarioAgente)
	if err != nil {
		return nil, err
	}
	output["DestinatarioRamo"], err = json.Marshal(r.DestinatarioRamo)
	if err != nil {
		return nil, err
	}
	output["Grilla"], err = json.Marshal(r.Grilla)
	if err != nil {
		return nil, err
	}
	output["DiaTransmion"], err = json.Marshal(r.DiaTransmion)
	if err != nil {
		return nil, err
	}
	output["EntregaLunes"], err = json.Marshal(r.EntregaLunes)
	if err != nil {
		return nil, err
	}
	output["EntregaMartes"], err = json.Marshal(r.EntregaMartes)
	if err != nil {
		return nil, err
	}
	output["EntregaMiercoles"], err = json.Marshal(r.EntregaMiercoles)
	if err != nil {
		return nil, err
	}
	output["EntregaJueves"], err = json.Marshal(r.EntregaJueves)
	if err != nil {
		return nil, err
	}
	output["EntregaViernes"], err = json.Marshal(r.EntregaViernes)
	if err != nil {
		return nil, err
	}
	output["EntregaSabado"], err = json.Marshal(r.EntregaSabado)
	if err != nil {
		return nil, err
	}
	output["EntregaDomingo"], err = json.Marshal(r.EntregaDomingo)
	if err != nil {
		return nil, err
	}
	output["EntregaPostFeriado"], err = json.Marshal(r.EntregaPostFeriado)
	if err != nil {
		return nil, err
	}
	output["Geolocalizacion"], err = json.Marshal(r.Geolocalizacion)
	if err != nil {
		return nil, err
	}
	output["Latitud"], err = json.Marshal(r.Latitud)
	if err != nil {
		return nil, err
	}
	output["Longitud"], err = json.Marshal(r.Longitud)
	if err != nil {
		return nil, err
	}
	output["TipoIntegracion"], err = json.Marshal(r.TipoIntegracion)
	if err != nil {
		return nil, err
	}
	output["Planta"], err = json.Marshal(r.Planta)
	if err != nil {
		return nil, err
	}
	output["Nave"], err = json.Marshal(r.Nave)
	if err != nil {
		return nil, err
	}
	output["ClienteSap"], err = json.Marshal(r.ClienteSap)
	if err != nil {
		return nil, err
	}
	output["CodLevantamientoCuarentena"], err = json.Marshal(r.CodLevantamientoCuarentena)
	if err != nil {
		return nil, err
	}
	output["WosTrazabilidad"], err = json.Marshal(r.WosTrazabilidad)
	if err != nil {
		return nil, err
	}
	output["TipoEncolado"], err = json.Marshal(r.TipoEncolado)
	if err != nil {
		return nil, err
	}
	output["ProcesosDondeInformaAnmat"], err = json.Marshal(r.ProcesosDondeInformaAnmat)
	if err != nil {
		return nil, err
	}
	output["AcondiExterno"], err = json.Marshal(r.AcondiExterno)
	if err != nil {
		return nil, err
	}
	output["InformaLevantamientoCuarentena"], err = json.Marshal(r.InformaLevantamientoCuarentena)
	if err != nil {
		return nil, err
	}
	output["ValidaRemito"], err = json.Marshal(r.ValidaRemito)
	if err != nil {
		return nil, err
	}
	output["CodAnmatAbastecimiento"], err = json.Marshal(r.CodAnmatAbastecimiento)
	if err != nil {
		return nil, err
	}
	output["AbastecimientoSeriesCliente"], err = json.Marshal(r.AbastecimientoSeriesCliente)
	if err != nil {
		return nil, err
	}
	output["TipoReceptorAnmat"], err = json.Marshal(r.TipoReceptorAnmat)
	if err != nil {
		return nil, err
	}
	output["Tipo"], err = json.Marshal(r.Tipo)
	if err != nil {
		return nil, err
	}
	output["AndreaniCOM"], err = json.Marshal(r.AndreaniCOM)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Cabecera) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["Compañía"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Compañía); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Compañía")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Descripción"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Descripción); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Descripción")
	}
	val = func() json.RawMessage {
		if v, ok := fields["IVA"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IVA); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IVA")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SeguimientoIntercambioPallet"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SeguimientoIntercambioPallet); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SeguimientoIntercambioPallet")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioCalle"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioCalle); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioCalle")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioNumero"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioNumero); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioNumero")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioPiso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioPiso); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioPiso")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioDepartamento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioDepartamento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioDepartamento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Ciudad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ciudad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Ciudad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Provincia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Provincia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Provincia")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CP"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CP); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CP")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Pais"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Pais); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Pais")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ISOPais"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ISOPais); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ISOPais")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Contacto1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contacto1); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Contacto1")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Telefono"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Telefono); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Telefono")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Mail"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Mail); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Mail")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Fax"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Fax); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Fax")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioCUIT"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioCUIT); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioCUIT")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioGLN"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioGLN); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioGLN")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioAgente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioAgente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioAgente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DestinatarioRamo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DestinatarioRamo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DestinatarioRamo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Grilla"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Grilla); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Grilla")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DiaTransmion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DiaTransmion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DiaTransmion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EntregaLunes"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EntregaLunes); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EntregaLunes")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EntregaMartes"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EntregaMartes); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EntregaMartes")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EntregaMiercoles"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EntregaMiercoles); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EntregaMiercoles")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EntregaJueves"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EntregaJueves); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EntregaJueves")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EntregaViernes"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EntregaViernes); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EntregaViernes")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EntregaSabado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EntregaSabado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EntregaSabado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EntregaDomingo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EntregaDomingo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EntregaDomingo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EntregaPostFeriado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EntregaPostFeriado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EntregaPostFeriado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Geolocalizacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Geolocalizacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Geolocalizacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Latitud"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Latitud); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Latitud")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Longitud"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Longitud); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Longitud")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoIntegracion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoIntegracion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoIntegracion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Planta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Planta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Planta")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Nave"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Nave); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Nave")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ClienteSap"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ClienteSap); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ClienteSap")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodLevantamientoCuarentena"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodLevantamientoCuarentena); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodLevantamientoCuarentena")
	}
	val = func() json.RawMessage {
		if v, ok := fields["WosTrazabilidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.WosTrazabilidad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for WosTrazabilidad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoEncolado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoEncolado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoEncolado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ProcesosDondeInformaAnmat"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ProcesosDondeInformaAnmat); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ProcesosDondeInformaAnmat")
	}
	val = func() json.RawMessage {
		if v, ok := fields["AcondiExterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AcondiExterno); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for AcondiExterno")
	}
	val = func() json.RawMessage {
		if v, ok := fields["InformaLevantamientoCuarentena"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.InformaLevantamientoCuarentena); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for InformaLevantamientoCuarentena")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ValidaRemito"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ValidaRemito); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ValidaRemito")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodAnmatAbastecimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodAnmatAbastecimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodAnmatAbastecimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["AbastecimientoSeriesCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AbastecimientoSeriesCliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for AbastecimientoSeriesCliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoReceptorAnmat"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoReceptorAnmat); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoReceptorAnmat")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Tipo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tipo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Tipo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["AndreaniCOM"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AndreaniCOM); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for AndreaniCOM")
	}
	return nil
}