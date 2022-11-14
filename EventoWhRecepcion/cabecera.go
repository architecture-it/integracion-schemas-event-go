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

type Cabecera struct {
	Propietario string `json:"Propietario"`

	RecepcionWH string `json:"RecepcionWH"`

	Remito *UnionNullString `json:"Remito"`

	OrdenCompra *UnionNullString `json:"OrdenCompra"`

	TipoRecepcion *UnionNullString `json:"TipoRecepcion"`

	ReferenciaTransportista *UnionNullString `json:"ReferenciaTransportista"`

	ReferenciaContenedor *UnionNullString `json:"ReferenciaContenedor"`

	Muelle *UnionNullString `json:"Muelle"`

	NumeroCita *UnionNullString `json:"NumeroCita"`

	ReferenciaProveedor *UnionNullString `json:"ReferenciaProveedor"`

	NumeroGuia *UnionNullString `json:"NumeroGuia"`

	Calle *UnionNullString `json:"Calle"`

	Numero *UnionNullString `json:"Numero"`

	Piso *UnionNullString `json:"Piso"`

	Dpto *UnionNullString `json:"Dpto"`

	GLN *UnionNullString `json:"GLN"`

	CiudadExpedidor *UnionNullString `json:"CiudadExpedidor"`

	Contacto *UnionNullString `json:"Contacto"`

	Email *UnionNullString `json:"Email"`

	CodISOPais *UnionNullString `json:"CodISOPais"`

	Telefono *UnionNullString `json:"Telefono"`

	EstadoExpedidor *UnionNullString `json:"EstadoExpedidor"`

	CodPostalExpedidor *UnionNullString `json:"CodPostalExpedidor"`

	FechaLlegada *UnionNullLong `json:"FechaLlegada"`

	FechaRecepcion *UnionNullLong `json:"FechaRecepcion"`

	CantEsperadaTotal float32 `json:"CantEsperadaTotal"`

	CantRecibidaTotal float32 `json:"CantRecibidaTotal"`
}

const CabeceraAvroCRC64Fingerprint = "\xfbظ\xc1\x9bU\b\xc6"

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
	err = vm.WriteString(r.RecepcionWH, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Remito, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.OrdenCompra, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoRecepcion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ReferenciaTransportista, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ReferenciaContenedor, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Muelle, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroCita, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ReferenciaProveedor, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroGuia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Calle, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Numero, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Piso, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Dpto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.GLN, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CiudadExpedidor, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Contacto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Email, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodISOPais, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Telefono, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstadoExpedidor, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodPostalExpedidor, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaLlegada, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaRecepcion, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.CantEsperadaTotal, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.CantRecibidaTotal, w)
	if err != nil {
		return err
	}
	return err
}

func (r Cabecera) Serialize(w io.Writer) error {
	return writeCabecera(r, w)
}

func (r Cabecera) Schema() string {
	return "{\"fields\":[{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"RecepcionWH\",\"type\":\"string\"},{\"name\":\"Remito\",\"type\":[\"null\",\"string\"]},{\"name\":\"OrdenCompra\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoRecepcion\",\"type\":[\"null\",\"string\"]},{\"name\":\"ReferenciaTransportista\",\"type\":[\"null\",\"string\"]},{\"name\":\"ReferenciaContenedor\",\"type\":[\"null\",\"string\"]},{\"name\":\"Muelle\",\"type\":[\"null\",\"string\"]},{\"name\":\"NumeroCita\",\"type\":[\"null\",\"string\"]},{\"name\":\"ReferenciaProveedor\",\"type\":[\"null\",\"string\"]},{\"name\":\"NumeroGuia\",\"type\":[\"null\",\"string\"]},{\"name\":\"Calle\",\"type\":[\"null\",\"string\"]},{\"name\":\"Numero\",\"type\":[\"null\",\"string\"]},{\"name\":\"Piso\",\"type\":[\"null\",\"string\"]},{\"name\":\"Dpto\",\"type\":[\"null\",\"string\"]},{\"name\":\"GLN\",\"type\":[\"null\",\"string\"]},{\"name\":\"CiudadExpedidor\",\"type\":[\"null\",\"string\"]},{\"name\":\"Contacto\",\"type\":[\"null\",\"string\"]},{\"name\":\"Email\",\"type\":[\"null\",\"string\"]},{\"name\":\"CodISOPais\",\"type\":[\"null\",\"string\"]},{\"name\":\"Telefono\",\"type\":[\"null\",\"string\"]},{\"name\":\"EstadoExpedidor\",\"type\":[\"null\",\"string\"]},{\"name\":\"CodPostalExpedidor\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaLlegada\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"FechaRecepcion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"CantEsperadaTotal\",\"type\":\"float\"},{\"name\":\"CantRecibidaTotal\",\"type\":\"float\"}],\"name\":\"Andreani.EventoWhRecepcion.Events.RecepcionCreacionCommon.Cabecera\",\"type\":\"record\"}"
}

func (r Cabecera) SchemaName() string {
	return "Andreani.EventoWhRecepcion.Events.RecepcionCreacionCommon.Cabecera"
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
		w := types.String{Target: &r.RecepcionWH}

		return w

	case 2:
		r.Remito = NewUnionNullString()

		return r.Remito
	case 3:
		r.OrdenCompra = NewUnionNullString()

		return r.OrdenCompra
	case 4:
		r.TipoRecepcion = NewUnionNullString()

		return r.TipoRecepcion
	case 5:
		r.ReferenciaTransportista = NewUnionNullString()

		return r.ReferenciaTransportista
	case 6:
		r.ReferenciaContenedor = NewUnionNullString()

		return r.ReferenciaContenedor
	case 7:
		r.Muelle = NewUnionNullString()

		return r.Muelle
	case 8:
		r.NumeroCita = NewUnionNullString()

		return r.NumeroCita
	case 9:
		r.ReferenciaProveedor = NewUnionNullString()

		return r.ReferenciaProveedor
	case 10:
		r.NumeroGuia = NewUnionNullString()

		return r.NumeroGuia
	case 11:
		r.Calle = NewUnionNullString()

		return r.Calle
	case 12:
		r.Numero = NewUnionNullString()

		return r.Numero
	case 13:
		r.Piso = NewUnionNullString()

		return r.Piso
	case 14:
		r.Dpto = NewUnionNullString()

		return r.Dpto
	case 15:
		r.GLN = NewUnionNullString()

		return r.GLN
	case 16:
		r.CiudadExpedidor = NewUnionNullString()

		return r.CiudadExpedidor
	case 17:
		r.Contacto = NewUnionNullString()

		return r.Contacto
	case 18:
		r.Email = NewUnionNullString()

		return r.Email
	case 19:
		r.CodISOPais = NewUnionNullString()

		return r.CodISOPais
	case 20:
		r.Telefono = NewUnionNullString()

		return r.Telefono
	case 21:
		r.EstadoExpedidor = NewUnionNullString()

		return r.EstadoExpedidor
	case 22:
		r.CodPostalExpedidor = NewUnionNullString()

		return r.CodPostalExpedidor
	case 23:
		r.FechaLlegada = NewUnionNullLong()

		return r.FechaLlegada
	case 24:
		r.FechaRecepcion = NewUnionNullLong()

		return r.FechaRecepcion
	case 25:
		w := types.Float{Target: &r.CantEsperadaTotal}

		return w

	case 26:
		w := types.Float{Target: &r.CantRecibidaTotal}

		return w

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
	case 2:
		r.Remito = nil
		return
	case 3:
		r.OrdenCompra = nil
		return
	case 4:
		r.TipoRecepcion = nil
		return
	case 5:
		r.ReferenciaTransportista = nil
		return
	case 6:
		r.ReferenciaContenedor = nil
		return
	case 7:
		r.Muelle = nil
		return
	case 8:
		r.NumeroCita = nil
		return
	case 9:
		r.ReferenciaProveedor = nil
		return
	case 10:
		r.NumeroGuia = nil
		return
	case 11:
		r.Calle = nil
		return
	case 12:
		r.Numero = nil
		return
	case 13:
		r.Piso = nil
		return
	case 14:
		r.Dpto = nil
		return
	case 15:
		r.GLN = nil
		return
	case 16:
		r.CiudadExpedidor = nil
		return
	case 17:
		r.Contacto = nil
		return
	case 18:
		r.Email = nil
		return
	case 19:
		r.CodISOPais = nil
		return
	case 20:
		r.Telefono = nil
		return
	case 21:
		r.EstadoExpedidor = nil
		return
	case 22:
		r.CodPostalExpedidor = nil
		return
	case 23:
		r.FechaLlegada = nil
		return
	case 24:
		r.FechaRecepcion = nil
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
	output["RecepcionWH"], err = json.Marshal(r.RecepcionWH)
	if err != nil {
		return nil, err
	}
	output["Remito"], err = json.Marshal(r.Remito)
	if err != nil {
		return nil, err
	}
	output["OrdenCompra"], err = json.Marshal(r.OrdenCompra)
	if err != nil {
		return nil, err
	}
	output["TipoRecepcion"], err = json.Marshal(r.TipoRecepcion)
	if err != nil {
		return nil, err
	}
	output["ReferenciaTransportista"], err = json.Marshal(r.ReferenciaTransportista)
	if err != nil {
		return nil, err
	}
	output["ReferenciaContenedor"], err = json.Marshal(r.ReferenciaContenedor)
	if err != nil {
		return nil, err
	}
	output["Muelle"], err = json.Marshal(r.Muelle)
	if err != nil {
		return nil, err
	}
	output["NumeroCita"], err = json.Marshal(r.NumeroCita)
	if err != nil {
		return nil, err
	}
	output["ReferenciaProveedor"], err = json.Marshal(r.ReferenciaProveedor)
	if err != nil {
		return nil, err
	}
	output["NumeroGuia"], err = json.Marshal(r.NumeroGuia)
	if err != nil {
		return nil, err
	}
	output["Calle"], err = json.Marshal(r.Calle)
	if err != nil {
		return nil, err
	}
	output["Numero"], err = json.Marshal(r.Numero)
	if err != nil {
		return nil, err
	}
	output["Piso"], err = json.Marshal(r.Piso)
	if err != nil {
		return nil, err
	}
	output["Dpto"], err = json.Marshal(r.Dpto)
	if err != nil {
		return nil, err
	}
	output["GLN"], err = json.Marshal(r.GLN)
	if err != nil {
		return nil, err
	}
	output["CiudadExpedidor"], err = json.Marshal(r.CiudadExpedidor)
	if err != nil {
		return nil, err
	}
	output["Contacto"], err = json.Marshal(r.Contacto)
	if err != nil {
		return nil, err
	}
	output["Email"], err = json.Marshal(r.Email)
	if err != nil {
		return nil, err
	}
	output["CodISOPais"], err = json.Marshal(r.CodISOPais)
	if err != nil {
		return nil, err
	}
	output["Telefono"], err = json.Marshal(r.Telefono)
	if err != nil {
		return nil, err
	}
	output["EstadoExpedidor"], err = json.Marshal(r.EstadoExpedidor)
	if err != nil {
		return nil, err
	}
	output["CodPostalExpedidor"], err = json.Marshal(r.CodPostalExpedidor)
	if err != nil {
		return nil, err
	}
	output["FechaLlegada"], err = json.Marshal(r.FechaLlegada)
	if err != nil {
		return nil, err
	}
	output["FechaRecepcion"], err = json.Marshal(r.FechaRecepcion)
	if err != nil {
		return nil, err
	}
	output["CantEsperadaTotal"], err = json.Marshal(r.CantEsperadaTotal)
	if err != nil {
		return nil, err
	}
	output["CantRecibidaTotal"], err = json.Marshal(r.CantRecibidaTotal)
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
		if v, ok := fields["RecepcionWH"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RecepcionWH); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for RecepcionWH")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Remito"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Remito); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Remito")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OrdenCompra"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenCompra); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrdenCompra")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoRecepcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoRecepcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoRecepcion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ReferenciaTransportista"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ReferenciaTransportista); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ReferenciaTransportista")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ReferenciaContenedor"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ReferenciaContenedor); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ReferenciaContenedor")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Muelle"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Muelle); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Muelle")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroCita"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroCita); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NumeroCita")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ReferenciaProveedor"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ReferenciaProveedor); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ReferenciaProveedor")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroGuia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroGuia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NumeroGuia")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Calle"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Calle); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Calle")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Numero"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Numero); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Numero")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Piso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Piso); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Piso")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Dpto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Dpto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Dpto")
	}
	val = func() json.RawMessage {
		if v, ok := fields["GLN"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.GLN); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for GLN")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CiudadExpedidor"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CiudadExpedidor); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CiudadExpedidor")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Contacto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Contacto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Contacto")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Email"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Email); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Email")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodISOPais"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodISOPais); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodISOPais")
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
		if v, ok := fields["EstadoExpedidor"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoExpedidor); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EstadoExpedidor")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodPostalExpedidor"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodPostalExpedidor); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodPostalExpedidor")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaLlegada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaLlegada); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaLlegada")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaRecepcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaRecepcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaRecepcion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CantEsperadaTotal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantEsperadaTotal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CantEsperadaTotal")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CantRecibidaTotal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantRecibidaTotal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CantRecibidaTotal")
	}
	return nil
}
