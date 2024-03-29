// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TyDExpedidoFallida.avsc
 */
package WosExpedicionTyDEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type TyDExpedidoFallida struct {
	Identificacion Identificacion `json:"Identificacion"`

	Cabecera Cabecera `json:"Cabecera"`
}

const TyDExpedidoFallidaAvroCRC64Fingerprint = "\x05\f\xfc\xa62\x8e*\xca"

func NewTyDExpedidoFallida() TyDExpedidoFallida {
	r := TyDExpedidoFallida{}
	r.Identificacion = NewIdentificacion()

	r.Cabecera = NewCabecera()

	return r
}

func DeserializeTyDExpedidoFallida(r io.Reader) (TyDExpedidoFallida, error) {
	t := NewTyDExpedidoFallida()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTyDExpedidoFallidaFromSchema(r io.Reader, schema string) (TyDExpedidoFallida, error) {
	t := NewTyDExpedidoFallida()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTyDExpedidoFallida(r TyDExpedidoFallida, w io.Writer) error {
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

func (r TyDExpedidoFallida) Serialize(w io.Writer) error {
	return writeTyDExpedidoFallida(r, w)
}

func (r TyDExpedidoFallida) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.WosExpedicionTyD.Events.TyDExpedidoFallidaCommon\",\"type\":\"record\"}},{\"name\":\"Cabecera\",\"type\":{\"fields\":[{\"name\":\"OrdenWh\",\"type\":\"string\"},{\"name\":\"OrdenCliente\",\"type\":\"string\"},{\"name\":\"NumeroEnvio\",\"type\":\"string\"},{\"name\":\"ContratoTMS\",\"type\":\"string\"},{\"name\":\"MensajeError\",\"type\":\"string\"},{\"name\":\"PlantaDeOperacionId\",\"type\":\"int\"},{\"name\":\"Puerta\",\"type\":\"string\"}],\"name\":\"Cabecera\",\"namespace\":\"Andreani.WosExpedicionTyD.Events.TyDExpedidoFallidaCommon\",\"type\":\"record\"}}],\"name\":\"Andreani.WosExpedicionTyD.Events.Record.TyDExpedidoFallida\",\"type\":\"record\"}"
}

func (r TyDExpedidoFallida) SchemaName() string {
	return "Andreani.WosExpedicionTyD.Events.Record.TyDExpedidoFallida"
}

func (_ TyDExpedidoFallida) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ TyDExpedidoFallida) SetInt(v int32)       { panic("Unsupported operation") }
func (_ TyDExpedidoFallida) SetLong(v int64)      { panic("Unsupported operation") }
func (_ TyDExpedidoFallida) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ TyDExpedidoFallida) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ TyDExpedidoFallida) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ TyDExpedidoFallida) SetString(v string)   { panic("Unsupported operation") }
func (_ TyDExpedidoFallida) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *TyDExpedidoFallida) Get(i int) types.Field {
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

func (r *TyDExpedidoFallida) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *TyDExpedidoFallida) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ TyDExpedidoFallida) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ TyDExpedidoFallida) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ TyDExpedidoFallida) HintSize(int)                     { panic("Unsupported operation") }
func (_ TyDExpedidoFallida) Finalize()                        {}

func (_ TyDExpedidoFallida) AvroCRC64Fingerprint() []byte {
	return []byte(TyDExpedidoFallidaAvroCRC64Fingerprint)
}

func (r TyDExpedidoFallida) MarshalJSON() ([]byte, error) {
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

func (r *TyDExpedidoFallida) UnmarshalJSON(data []byte) error {
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
