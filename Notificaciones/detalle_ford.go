// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     DetalleFord.avsc
 */
package NotificacionesEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type DetalleFord struct {
	CodigoDeEnvio *UnionNullString `json:"codigoDeEnvio"`

	Remito *UnionNullString `json:"remito"`

	NotaDespacho *UnionNullString `json:"notaDespacho"`

	Bultos *UnionNullLong `json:"bultos"`

	FechaAdmision *UnionNullLong `json:"fechaAdmision"`

	FechaInsercion *UnionNullLong `json:"fechaInsercion"`
}

const DetalleFordAvroCRC64Fingerprint = "ĚW\xdb(\xd3\xfd "

func NewDetalleFord() DetalleFord {
	r := DetalleFord{}
	r.CodigoDeEnvio = nil
	r.Remito = nil
	r.NotaDespacho = nil
	r.Bultos = nil
	r.FechaAdmision = nil
	r.FechaInsercion = nil
	return r
}

func DeserializeDetalleFord(r io.Reader) (DetalleFord, error) {
	t := NewDetalleFord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDetalleFordFromSchema(r io.Reader, schema string) (DetalleFord, error) {
	t := NewDetalleFord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDetalleFord(r DetalleFord, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.CodigoDeEnvio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Remito, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NotaDespacho, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.Bultos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaAdmision, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaInsercion, w)
	if err != nil {
		return err
	}
	return err
}

func (r DetalleFord) Serialize(w io.Writer) error {
	return writeDetalleFord(r, w)
}

func (r DetalleFord) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"codigoDeEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"notaDespacho\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"bultos\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"fechaAdmision\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"fechaInsercion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"Andreani.Notificaciones.Events.Records.DetalleFord\",\"type\":\"record\"}"
}

func (r DetalleFord) SchemaName() string {
	return "Andreani.Notificaciones.Events.Records.DetalleFord"
}

func (_ DetalleFord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DetalleFord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DetalleFord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DetalleFord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DetalleFord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DetalleFord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DetalleFord) SetString(v string)   { panic("Unsupported operation") }
func (_ DetalleFord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DetalleFord) Get(i int) types.Field {
	switch i {
	case 0:
		r.CodigoDeEnvio = NewUnionNullString()

		return r.CodigoDeEnvio
	case 1:
		r.Remito = NewUnionNullString()

		return r.Remito
	case 2:
		r.NotaDespacho = NewUnionNullString()

		return r.NotaDespacho
	case 3:
		r.Bultos = NewUnionNullLong()

		return r.Bultos
	case 4:
		r.FechaAdmision = NewUnionNullLong()

		return r.FechaAdmision
	case 5:
		r.FechaInsercion = NewUnionNullLong()

		return r.FechaInsercion
	}
	panic("Unknown field index")
}

func (r *DetalleFord) SetDefault(i int) {
	switch i {
	case 0:
		r.CodigoDeEnvio = nil
		return
	case 1:
		r.Remito = nil
		return
	case 2:
		r.NotaDespacho = nil
		return
	case 3:
		r.Bultos = nil
		return
	case 4:
		r.FechaAdmision = nil
		return
	case 5:
		r.FechaInsercion = nil
		return
	}
	panic("Unknown field index")
}

func (r *DetalleFord) NullField(i int) {
	switch i {
	case 0:
		r.CodigoDeEnvio = nil
		return
	case 1:
		r.Remito = nil
		return
	case 2:
		r.NotaDespacho = nil
		return
	case 3:
		r.Bultos = nil
		return
	case 4:
		r.FechaAdmision = nil
		return
	case 5:
		r.FechaInsercion = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ DetalleFord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DetalleFord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DetalleFord) HintSize(int)                     { panic("Unsupported operation") }
func (_ DetalleFord) Finalize()                        {}

func (_ DetalleFord) AvroCRC64Fingerprint() []byte {
	return []byte(DetalleFordAvroCRC64Fingerprint)
}

func (r DetalleFord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["codigoDeEnvio"], err = json.Marshal(r.CodigoDeEnvio)
	if err != nil {
		return nil, err
	}
	output["remito"], err = json.Marshal(r.Remito)
	if err != nil {
		return nil, err
	}
	output["notaDespacho"], err = json.Marshal(r.NotaDespacho)
	if err != nil {
		return nil, err
	}
	output["bultos"], err = json.Marshal(r.Bultos)
	if err != nil {
		return nil, err
	}
	output["fechaAdmision"], err = json.Marshal(r.FechaAdmision)
	if err != nil {
		return nil, err
	}
	output["fechaInsercion"], err = json.Marshal(r.FechaInsercion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DetalleFord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["codigoDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeEnvio); err != nil {
			return err
		}
	} else {
		r.CodigoDeEnvio = NewUnionNullString()

		r.CodigoDeEnvio = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["remito"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Remito); err != nil {
			return err
		}
	} else {
		r.Remito = NewUnionNullString()

		r.Remito = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["notaDespacho"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NotaDespacho); err != nil {
			return err
		}
	} else {
		r.NotaDespacho = NewUnionNullString()

		r.NotaDespacho = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["bultos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Bultos); err != nil {
			return err
		}
	} else {
		r.Bultos = NewUnionNullLong()

		r.Bultos = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaAdmision"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaAdmision); err != nil {
			return err
		}
	} else {
		r.FechaAdmision = NewUnionNullLong()

		r.FechaAdmision = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaInsercion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaInsercion); err != nil {
			return err
		}
	} else {
		r.FechaInsercion = NewUnionNullLong()

		r.FechaInsercion = nil
	}
	return nil
}
