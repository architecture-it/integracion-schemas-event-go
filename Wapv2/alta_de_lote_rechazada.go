// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AltaDeLoteRechazada.avsc
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

type AltaDeLoteRechazada struct {
	SolicitudDeAccionAlmacen SolicitudDeAccionAlmacen `json:"solicitudDeAccionAlmacen"`

	LoteSCE LoteSCE `json:"loteSCE"`

	Topic string `json:"Topic"`

	Razon string `json:"razon"`
}

const AltaDeLoteRechazadaAvroCRC64Fingerprint = "c\a\xa8\x06\x88\xd6\x17|"

func NewAltaDeLoteRechazada() AltaDeLoteRechazada {
	r := AltaDeLoteRechazada{}
	r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

	r.LoteSCE = NewLoteSCE()

	r.Topic = "Almacen/Solicitudes/AltaDeLoteRechazada"
	return r
}

func DeserializeAltaDeLoteRechazada(r io.Reader) (AltaDeLoteRechazada, error) {
	t := NewAltaDeLoteRechazada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAltaDeLoteRechazadaFromSchema(r io.Reader, schema string) (AltaDeLoteRechazada, error) {
	t := NewAltaDeLoteRechazada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAltaDeLoteRechazada(r AltaDeLoteRechazada, w io.Writer) error {
	var err error
	err = writeSolicitudDeAccionAlmacen(r.SolicitudDeAccionAlmacen, w)
	if err != nil {
		return err
	}
	err = writeLoteSCE(r.LoteSCE, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Topic, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Razon, w)
	if err != nil {
		return err
	}
	return err
}

func (r AltaDeLoteRechazada) Serialize(w io.Writer) error {
	return writeAltaDeLoteRechazada(r, w)
}

func (r AltaDeLoteRechazada) Schema() string {
	return "{\"fields\":[{\"name\":\"solicitudDeAccionAlmacen\",\"type\":{\"fields\":[{\"name\":\"eventoDeNegocio\",\"type\":{\"fields\":[{\"name\":\"timestamp\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"remitente\",\"type\":\"string\"},{\"default\":null,\"name\":\"destinatario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"vencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"EventoDeNegocio\",\"type\":\"record\"}},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"default\":null,\"name\":\"instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoDistribucion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]}],\"name\":\"SolicitudDeAccionAlmacen\",\"type\":\"record\"}},{\"name\":\"loteSCE\",\"type\":{\"fields\":[{\"name\":\"articulo\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"estadolote\",\"type\":\"string\"},{\"name\":\"loteCaja\",\"type\":\"string\"},{\"name\":\"loteSecundario\",\"type\":\"string\"},{\"name\":\"trazable\",\"type\":\"int\"},{\"name\":\"zonaConsumo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"fechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"paquete\",\"type\":[\"null\",\"string\"]},{\"name\":\"vidaUtil\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"etregaantesde\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"consumoantesde\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"estadoDeUbicacionLote\",\"type\":[\"null\",\"string\"]}],\"name\":\"LoteSCE\",\"type\":\"record\"}},{\"default\":\"Almacen/Solicitudes/AltaDeLoteRechazada\",\"name\":\"Topic\",\"type\":\"string\"},{\"name\":\"razon\",\"type\":\"string\"}],\"name\":\"Andreani.Wapv2.Events.Record.AltaDeLoteRechazada\",\"type\":\"record\"}"
}

func (r AltaDeLoteRechazada) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.AltaDeLoteRechazada"
}

func (_ AltaDeLoteRechazada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ AltaDeLoteRechazada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ AltaDeLoteRechazada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ AltaDeLoteRechazada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ AltaDeLoteRechazada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ AltaDeLoteRechazada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ AltaDeLoteRechazada) SetString(v string)   { panic("Unsupported operation") }
func (_ AltaDeLoteRechazada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AltaDeLoteRechazada) Get(i int) types.Field {
	switch i {
	case 0:
		r.SolicitudDeAccionAlmacen = NewSolicitudDeAccionAlmacen()

		w := types.Record{Target: &r.SolicitudDeAccionAlmacen}

		return w

	case 1:
		r.LoteSCE = NewLoteSCE()

		w := types.Record{Target: &r.LoteSCE}

		return w

	case 2:
		w := types.String{Target: &r.Topic}

		return w

	case 3:
		w := types.String{Target: &r.Razon}

		return w

	}
	panic("Unknown field index")
}

func (r *AltaDeLoteRechazada) SetDefault(i int) {
	switch i {
	case 2:
		r.Topic = "Almacen/Solicitudes/AltaDeLoteRechazada"
		return
	}
	panic("Unknown field index")
}

func (r *AltaDeLoteRechazada) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ AltaDeLoteRechazada) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ AltaDeLoteRechazada) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ AltaDeLoteRechazada) HintSize(int)                     { panic("Unsupported operation") }
func (_ AltaDeLoteRechazada) Finalize()                        {}

func (_ AltaDeLoteRechazada) AvroCRC64Fingerprint() []byte {
	return []byte(AltaDeLoteRechazadaAvroCRC64Fingerprint)
}

func (r AltaDeLoteRechazada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["solicitudDeAccionAlmacen"], err = json.Marshal(r.SolicitudDeAccionAlmacen)
	if err != nil {
		return nil, err
	}
	output["loteSCE"], err = json.Marshal(r.LoteSCE)
	if err != nil {
		return nil, err
	}
	output["Topic"], err = json.Marshal(r.Topic)
	if err != nil {
		return nil, err
	}
	output["razon"], err = json.Marshal(r.Razon)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *AltaDeLoteRechazada) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["loteSCE"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteSCE); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for loteSCE")
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
		r.Topic = "Almacen/Solicitudes/AltaDeLoteRechazada"
	}
	val = func() json.RawMessage {
		if v, ok := fields["razon"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Razon); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for razon")
	}
	return nil
}
