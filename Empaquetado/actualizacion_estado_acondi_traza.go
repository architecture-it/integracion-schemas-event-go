// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ActualizacionEstadoAcondiTraza.avsc
 */
package EmpaquetadoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ActualizacionEstadoAcondiTraza struct {
	Identificacion Identificacion `json:"Identificacion"`

	OrdenWh string `json:"OrdenWh"`

	OrdenCliente string `json:"OrdenCliente"`

	CodigoEstadoAcondi *UnionNullString `json:"CodigoEstadoAcondi"`

	DescripcionEstadoAcondi *UnionNullString `json:"DescripcionEstadoAcondi"`

	CodigoEstadoTraza *UnionNullString `json:"CodigoEstadoTraza"`

	DescripcionEstadoTraza *UnionNullString `json:"DescripcionEstadoTraza"`

	DescripcionErrorTraza *UnionNullString `json:"DescripcionErrorTraza"`

	DescripcionErrorAcondi *UnionNullString `json:"DescripcionErrorAcondi"`
}

const ActualizacionEstadoAcondiTrazaAvroCRC64Fingerprint = "XF\x1a\xea\x18\x1f\xebp"

func NewActualizacionEstadoAcondiTraza() ActualizacionEstadoAcondiTraza {
	r := ActualizacionEstadoAcondiTraza{}
	r.Identificacion = NewIdentificacion()

	r.CodigoEstadoAcondi = nil
	r.DescripcionEstadoAcondi = nil
	r.CodigoEstadoTraza = nil
	r.DescripcionEstadoTraza = nil
	r.DescripcionErrorTraza = nil
	r.DescripcionErrorAcondi = nil
	return r
}

func DeserializeActualizacionEstadoAcondiTraza(r io.Reader) (ActualizacionEstadoAcondiTraza, error) {
	t := NewActualizacionEstadoAcondiTraza()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeActualizacionEstadoAcondiTrazaFromSchema(r io.Reader, schema string) (ActualizacionEstadoAcondiTraza, error) {
	t := NewActualizacionEstadoAcondiTraza()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeActualizacionEstadoAcondiTraza(r ActualizacionEstadoAcondiTraza, w io.Writer) error {
	var err error
	err = writeIdentificacion(r.Identificacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.OrdenWh, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.OrdenCliente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoEstadoAcondi, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DescripcionEstadoAcondi, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoEstadoTraza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DescripcionEstadoTraza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DescripcionErrorTraza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DescripcionErrorAcondi, w)
	if err != nil {
		return err
	}
	return err
}

func (r ActualizacionEstadoAcondiTraza) Serialize(w io.Writer) error {
	return writeActualizacionEstadoAcondiTraza(r, w)
}

func (r ActualizacionEstadoAcondiTraza) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"},{\"default\":null,\"name\":\"PlantaOperacionId\",\"type\":[\"null\",\"int\"]}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.Empaquetado.Events.Common\",\"type\":\"record\"}},{\"name\":\"OrdenWh\",\"type\":\"string\"},{\"name\":\"OrdenCliente\",\"type\":\"string\"},{\"default\":null,\"name\":\"CodigoEstadoAcondi\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DescripcionEstadoAcondi\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoEstadoTraza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DescripcionEstadoTraza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DescripcionErrorTraza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DescripcionErrorAcondi\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.Empaquetado.Events.Record.ActualizacionEstadoAcondiTraza\",\"type\":\"record\"}"
}

func (r ActualizacionEstadoAcondiTraza) SchemaName() string {
	return "Andreani.Empaquetado.Events.Record.ActualizacionEstadoAcondiTraza"
}

func (_ ActualizacionEstadoAcondiTraza) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ActualizacionEstadoAcondiTraza) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ActualizacionEstadoAcondiTraza) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ActualizacionEstadoAcondiTraza) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ActualizacionEstadoAcondiTraza) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ActualizacionEstadoAcondiTraza) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ActualizacionEstadoAcondiTraza) SetString(v string)   { panic("Unsupported operation") }
func (_ ActualizacionEstadoAcondiTraza) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ActualizacionEstadoAcondiTraza) Get(i int) types.Field {
	switch i {
	case 0:
		r.Identificacion = NewIdentificacion()

		w := types.Record{Target: &r.Identificacion}

		return w

	case 1:
		w := types.String{Target: &r.OrdenWh}

		return w

	case 2:
		w := types.String{Target: &r.OrdenCliente}

		return w

	case 3:
		r.CodigoEstadoAcondi = NewUnionNullString()

		return r.CodigoEstadoAcondi
	case 4:
		r.DescripcionEstadoAcondi = NewUnionNullString()

		return r.DescripcionEstadoAcondi
	case 5:
		r.CodigoEstadoTraza = NewUnionNullString()

		return r.CodigoEstadoTraza
	case 6:
		r.DescripcionEstadoTraza = NewUnionNullString()

		return r.DescripcionEstadoTraza
	case 7:
		r.DescripcionErrorTraza = NewUnionNullString()

		return r.DescripcionErrorTraza
	case 8:
		r.DescripcionErrorAcondi = NewUnionNullString()

		return r.DescripcionErrorAcondi
	}
	panic("Unknown field index")
}

func (r *ActualizacionEstadoAcondiTraza) SetDefault(i int) {
	switch i {
	case 3:
		r.CodigoEstadoAcondi = nil
		return
	case 4:
		r.DescripcionEstadoAcondi = nil
		return
	case 5:
		r.CodigoEstadoTraza = nil
		return
	case 6:
		r.DescripcionEstadoTraza = nil
		return
	case 7:
		r.DescripcionErrorTraza = nil
		return
	case 8:
		r.DescripcionErrorAcondi = nil
		return
	}
	panic("Unknown field index")
}

func (r *ActualizacionEstadoAcondiTraza) NullField(i int) {
	switch i {
	case 3:
		r.CodigoEstadoAcondi = nil
		return
	case 4:
		r.DescripcionEstadoAcondi = nil
		return
	case 5:
		r.CodigoEstadoTraza = nil
		return
	case 6:
		r.DescripcionEstadoTraza = nil
		return
	case 7:
		r.DescripcionErrorTraza = nil
		return
	case 8:
		r.DescripcionErrorAcondi = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ActualizacionEstadoAcondiTraza) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ActualizacionEstadoAcondiTraza) AppendArray() types.Field { panic("Unsupported operation") }
func (_ ActualizacionEstadoAcondiTraza) HintSize(int)             { panic("Unsupported operation") }
func (_ ActualizacionEstadoAcondiTraza) Finalize()                {}

func (_ ActualizacionEstadoAcondiTraza) AvroCRC64Fingerprint() []byte {
	return []byte(ActualizacionEstadoAcondiTrazaAvroCRC64Fingerprint)
}

func (r ActualizacionEstadoAcondiTraza) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Identificacion"], err = json.Marshal(r.Identificacion)
	if err != nil {
		return nil, err
	}
	output["OrdenWh"], err = json.Marshal(r.OrdenWh)
	if err != nil {
		return nil, err
	}
	output["OrdenCliente"], err = json.Marshal(r.OrdenCliente)
	if err != nil {
		return nil, err
	}
	output["CodigoEstadoAcondi"], err = json.Marshal(r.CodigoEstadoAcondi)
	if err != nil {
		return nil, err
	}
	output["DescripcionEstadoAcondi"], err = json.Marshal(r.DescripcionEstadoAcondi)
	if err != nil {
		return nil, err
	}
	output["CodigoEstadoTraza"], err = json.Marshal(r.CodigoEstadoTraza)
	if err != nil {
		return nil, err
	}
	output["DescripcionEstadoTraza"], err = json.Marshal(r.DescripcionEstadoTraza)
	if err != nil {
		return nil, err
	}
	output["DescripcionErrorTraza"], err = json.Marshal(r.DescripcionErrorTraza)
	if err != nil {
		return nil, err
	}
	output["DescripcionErrorAcondi"], err = json.Marshal(r.DescripcionErrorAcondi)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ActualizacionEstadoAcondiTraza) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["OrdenWh"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenWh); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrdenWh")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OrdenCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenCliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrdenCliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoEstadoAcondi"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoEstadoAcondi); err != nil {
			return err
		}
	} else {
		r.CodigoEstadoAcondi = NewUnionNullString()

		r.CodigoEstadoAcondi = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DescripcionEstadoAcondi"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DescripcionEstadoAcondi); err != nil {
			return err
		}
	} else {
		r.DescripcionEstadoAcondi = NewUnionNullString()

		r.DescripcionEstadoAcondi = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoEstadoTraza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoEstadoTraza); err != nil {
			return err
		}
	} else {
		r.CodigoEstadoTraza = NewUnionNullString()

		r.CodigoEstadoTraza = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DescripcionEstadoTraza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DescripcionEstadoTraza); err != nil {
			return err
		}
	} else {
		r.DescripcionEstadoTraza = NewUnionNullString()

		r.DescripcionEstadoTraza = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DescripcionErrorTraza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DescripcionErrorTraza); err != nil {
			return err
		}
	} else {
		r.DescripcionErrorTraza = NewUnionNullString()

		r.DescripcionErrorTraza = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DescripcionErrorAcondi"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DescripcionErrorAcondi); err != nil {
			return err
		}
	} else {
		r.DescripcionErrorAcondi = NewUnionNullString()

		r.DescripcionErrorAcondi = nil
	}
	return nil
}
