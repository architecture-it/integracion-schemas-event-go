// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     SelfTransportShipmentEvent.avsc
 */
package MEunoApiCorpoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type DireccionRemitente struct {
	IdSucursal *UnionNullInt `json:"idSucursal"`

	Calle string `json:"calle"`

	Numero string `json:"numero"`

	Piso string `json:"piso"`

	Unidad string `json:"unidad"`

	CodigoPostal string `json:"codigoPostal"`

	Localidad string `json:"localidad"`

	Provincia string `json:"provincia"`

	Observaciones string `json:"observaciones"`
}

const DireccionRemitenteAvroCRC64Fingerprint = "\xc47\x1f\x8e6\xdec\x04"

func NewDireccionRemitente() DireccionRemitente {
	r := DireccionRemitente{}
	r.IdSucursal = nil
	return r
}

func DeserializeDireccionRemitente(r io.Reader) (DireccionRemitente, error) {
	t := NewDireccionRemitente()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDireccionRemitenteFromSchema(r io.Reader, schema string) (DireccionRemitente, error) {
	t := NewDireccionRemitente()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDireccionRemitente(r DireccionRemitente, w io.Writer) error {
	var err error
	err = writeUnionNullInt(r.IdSucursal, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Calle, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Numero, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Piso, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Unidad, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoPostal, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Localidad, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Provincia, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Observaciones, w)
	if err != nil {
		return err
	}
	return err
}

func (r DireccionRemitente) Serialize(w io.Writer) error {
	return writeDireccionRemitente(r, w)
}

func (r DireccionRemitente) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"idSucursal\",\"type\":[\"null\",\"int\"]},{\"name\":\"calle\",\"type\":\"string\"},{\"name\":\"numero\",\"type\":\"string\"},{\"name\":\"piso\",\"type\":\"string\"},{\"name\":\"unidad\",\"type\":\"string\"},{\"name\":\"codigoPostal\",\"type\":\"string\"},{\"name\":\"localidad\",\"type\":\"string\"},{\"name\":\"provincia\",\"type\":\"string\"},{\"name\":\"observaciones\",\"type\":\"string\"}],\"name\":\"Andreani.MEunoApiCorpo.Events.Record.Structs.DireccionRemitente\",\"type\":\"record\"}"
}

func (r DireccionRemitente) SchemaName() string {
	return "Andreani.MEunoApiCorpo.Events.Record.Structs.DireccionRemitente"
}

func (_ DireccionRemitente) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DireccionRemitente) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DireccionRemitente) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DireccionRemitente) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DireccionRemitente) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DireccionRemitente) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DireccionRemitente) SetString(v string)   { panic("Unsupported operation") }
func (_ DireccionRemitente) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DireccionRemitente) Get(i int) types.Field {
	switch i {
	case 0:
		r.IdSucursal = NewUnionNullInt()

		return r.IdSucursal
	case 1:
		w := types.String{Target: &r.Calle}

		return w

	case 2:
		w := types.String{Target: &r.Numero}

		return w

	case 3:
		w := types.String{Target: &r.Piso}

		return w

	case 4:
		w := types.String{Target: &r.Unidad}

		return w

	case 5:
		w := types.String{Target: &r.CodigoPostal}

		return w

	case 6:
		w := types.String{Target: &r.Localidad}

		return w

	case 7:
		w := types.String{Target: &r.Provincia}

		return w

	case 8:
		w := types.String{Target: &r.Observaciones}

		return w

	}
	panic("Unknown field index")
}

func (r *DireccionRemitente) SetDefault(i int) {
	switch i {
	case 0:
		r.IdSucursal = nil
		return
	}
	panic("Unknown field index")
}

func (r *DireccionRemitente) NullField(i int) {
	switch i {
	case 0:
		r.IdSucursal = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ DireccionRemitente) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DireccionRemitente) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DireccionRemitente) HintSize(int)                     { panic("Unsupported operation") }
func (_ DireccionRemitente) Finalize()                        {}

func (_ DireccionRemitente) AvroCRC64Fingerprint() []byte {
	return []byte(DireccionRemitenteAvroCRC64Fingerprint)
}

func (r DireccionRemitente) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["idSucursal"], err = json.Marshal(r.IdSucursal)
	if err != nil {
		return nil, err
	}
	output["calle"], err = json.Marshal(r.Calle)
	if err != nil {
		return nil, err
	}
	output["numero"], err = json.Marshal(r.Numero)
	if err != nil {
		return nil, err
	}
	output["piso"], err = json.Marshal(r.Piso)
	if err != nil {
		return nil, err
	}
	output["unidad"], err = json.Marshal(r.Unidad)
	if err != nil {
		return nil, err
	}
	output["codigoPostal"], err = json.Marshal(r.CodigoPostal)
	if err != nil {
		return nil, err
	}
	output["localidad"], err = json.Marshal(r.Localidad)
	if err != nil {
		return nil, err
	}
	output["provincia"], err = json.Marshal(r.Provincia)
	if err != nil {
		return nil, err
	}
	output["observaciones"], err = json.Marshal(r.Observaciones)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DireccionRemitente) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["idSucursal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdSucursal); err != nil {
			return err
		}
	} else {
		r.IdSucursal = NewUnionNullInt()

		r.IdSucursal = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["calle"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Calle); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for calle")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numero"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Numero); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numero")
	}
	val = func() json.RawMessage {
		if v, ok := fields["piso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Piso); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for piso")
	}
	val = func() json.RawMessage {
		if v, ok := fields["unidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Unidad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for unidad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoPostal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoPostal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigoPostal")
	}
	val = func() json.RawMessage {
		if v, ok := fields["localidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Localidad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for localidad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["provincia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Provincia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for provincia")
	}
	val = func() json.RawMessage {
		if v, ok := fields["observaciones"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Observaciones); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for observaciones")
	}
	return nil
}
