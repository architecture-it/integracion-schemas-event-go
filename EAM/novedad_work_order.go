// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NovedadWorkOrder.avsc
 */
package EAMEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type NovedadWorkOrder struct {
	IdEquipoEAM *UnionNullString `json:"IdEquipoEAM"`

	CausaFalla *UnionNullString `json:"causaFalla"`

	TrabajoRealizado *UnionNullString `json:"trabajoRealizado"`

	EstadoFinal *UnionNullString `json:"estadoFinal"`

	IdOTEAM *UnionNullString `json:"IdOTEAM"`
}

const NovedadWorkOrderAvroCRC64Fingerprint = "\xedN\xca;\xdc\v<\x17"

func NewNovedadWorkOrder() NovedadWorkOrder {
	r := NovedadWorkOrder{}
	r.IdEquipoEAM = nil
	r.CausaFalla = nil
	r.TrabajoRealizado = nil
	r.EstadoFinal = nil
	r.IdOTEAM = nil
	return r
}

func DeserializeNovedadWorkOrder(r io.Reader) (NovedadWorkOrder, error) {
	t := NewNovedadWorkOrder()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNovedadWorkOrderFromSchema(r io.Reader, schema string) (NovedadWorkOrder, error) {
	t := NewNovedadWorkOrder()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNovedadWorkOrder(r NovedadWorkOrder, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.IdEquipoEAM, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CausaFalla, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TrabajoRealizado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstadoFinal, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.IdOTEAM, w)
	if err != nil {
		return err
	}
	return err
}

func (r NovedadWorkOrder) Serialize(w io.Writer) error {
	return writeNovedadWorkOrder(r, w)
}

func (r NovedadWorkOrder) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"IdEquipoEAM\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"causaFalla\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"trabajoRealizado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoFinal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"IdOTEAM\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.EAM.Events.Sharepoint.NovedadWorkOrder\",\"type\":\"record\"}"
}

func (r NovedadWorkOrder) SchemaName() string {
	return "Andreani.EAM.Events.Sharepoint.NovedadWorkOrder"
}

func (_ NovedadWorkOrder) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NovedadWorkOrder) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NovedadWorkOrder) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NovedadWorkOrder) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NovedadWorkOrder) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NovedadWorkOrder) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NovedadWorkOrder) SetString(v string)   { panic("Unsupported operation") }
func (_ NovedadWorkOrder) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NovedadWorkOrder) Get(i int) types.Field {
	switch i {
	case 0:
		r.IdEquipoEAM = NewUnionNullString()

		return r.IdEquipoEAM
	case 1:
		r.CausaFalla = NewUnionNullString()

		return r.CausaFalla
	case 2:
		r.TrabajoRealizado = NewUnionNullString()

		return r.TrabajoRealizado
	case 3:
		r.EstadoFinal = NewUnionNullString()

		return r.EstadoFinal
	case 4:
		r.IdOTEAM = NewUnionNullString()

		return r.IdOTEAM
	}
	panic("Unknown field index")
}

func (r *NovedadWorkOrder) SetDefault(i int) {
	switch i {
	case 0:
		r.IdEquipoEAM = nil
		return
	case 1:
		r.CausaFalla = nil
		return
	case 2:
		r.TrabajoRealizado = nil
		return
	case 3:
		r.EstadoFinal = nil
		return
	case 4:
		r.IdOTEAM = nil
		return
	}
	panic("Unknown field index")
}

func (r *NovedadWorkOrder) NullField(i int) {
	switch i {
	case 0:
		r.IdEquipoEAM = nil
		return
	case 1:
		r.CausaFalla = nil
		return
	case 2:
		r.TrabajoRealizado = nil
		return
	case 3:
		r.EstadoFinal = nil
		return
	case 4:
		r.IdOTEAM = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ NovedadWorkOrder) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NovedadWorkOrder) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NovedadWorkOrder) HintSize(int)                     { panic("Unsupported operation") }
func (_ NovedadWorkOrder) Finalize()                        {}

func (_ NovedadWorkOrder) AvroCRC64Fingerprint() []byte {
	return []byte(NovedadWorkOrderAvroCRC64Fingerprint)
}

func (r NovedadWorkOrder) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["IdEquipoEAM"], err = json.Marshal(r.IdEquipoEAM)
	if err != nil {
		return nil, err
	}
	output["causaFalla"], err = json.Marshal(r.CausaFalla)
	if err != nil {
		return nil, err
	}
	output["trabajoRealizado"], err = json.Marshal(r.TrabajoRealizado)
	if err != nil {
		return nil, err
	}
	output["estadoFinal"], err = json.Marshal(r.EstadoFinal)
	if err != nil {
		return nil, err
	}
	output["IdOTEAM"], err = json.Marshal(r.IdOTEAM)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NovedadWorkOrder) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["IdEquipoEAM"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdEquipoEAM); err != nil {
			return err
		}
	} else {
		r.IdEquipoEAM = NewUnionNullString()

		r.IdEquipoEAM = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["causaFalla"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CausaFalla); err != nil {
			return err
		}
	} else {
		r.CausaFalla = NewUnionNullString()

		r.CausaFalla = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["trabajoRealizado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TrabajoRealizado); err != nil {
			return err
		}
	} else {
		r.TrabajoRealizado = NewUnionNullString()

		r.TrabajoRealizado = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["estadoFinal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoFinal); err != nil {
			return err
		}
	} else {
		r.EstadoFinal = NewUnionNullString()

		r.EstadoFinal = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["IdOTEAM"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdOTEAM); err != nil {
			return err
		}
	} else {
		r.IdOTEAM = NewUnionNullString()

		r.IdOTEAM = nil
	}
	return nil
}
