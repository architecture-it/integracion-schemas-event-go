// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     RemitenteEtiqueta.avsc
 */
package CorporativoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type RemitenteEtiqueta struct {
	Nombre string `json:"Nombre"`

	Apellido string `json:"Apellido"`

	Telefono string `json:"Telefono"`

	DNI string `json:"DNI"`

	Email string `json:"Email"`

	Direccion DireccionRemitenteEtiqueta `json:"Direccion"`
}

const RemitenteEtiquetaAvroCRC64Fingerprint = "\x8f\x8f,\x1e\fZ\x1a\xdf"

func NewRemitenteEtiqueta() RemitenteEtiqueta {
	r := RemitenteEtiqueta{}
	r.Direccion = NewDireccionRemitenteEtiqueta()

	return r
}

func DeserializeRemitenteEtiqueta(r io.Reader) (RemitenteEtiqueta, error) {
	t := NewRemitenteEtiqueta()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRemitenteEtiquetaFromSchema(r io.Reader, schema string) (RemitenteEtiqueta, error) {
	t := NewRemitenteEtiqueta()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRemitenteEtiqueta(r RemitenteEtiqueta, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Nombre, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Apellido, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Telefono, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.DNI, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Email, w)
	if err != nil {
		return err
	}
	err = writeDireccionRemitenteEtiqueta(r.Direccion, w)
	if err != nil {
		return err
	}
	return err
}

func (r RemitenteEtiqueta) Serialize(w io.Writer) error {
	return writeRemitenteEtiqueta(r, w)
}

func (r RemitenteEtiqueta) Schema() string {
	return "{\"fields\":[{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Apellido\",\"type\":\"string\"},{\"name\":\"Telefono\",\"type\":\"string\"},{\"name\":\"DNI\",\"type\":\"string\"},{\"name\":\"Email\",\"type\":\"string\"},{\"name\":\"Direccion\",\"type\":{\"fields\":[{\"name\":\"Calle\",\"type\":\"string\"},{\"name\":\"Numero\",\"type\":\"string\"},{\"default\":null,\"name\":\"Piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Unidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"},{\"name\":\"Provincia\",\"type\":\"string\"},{\"default\":null,\"name\":\"DatosAdicionales\",\"type\":[\"null\",\"string\"]},{\"name\":\"Latitud\",\"type\":\"string\"},{\"name\":\"Longitud\",\"type\":\"string\"}],\"name\":\"DireccionRemitenteEtiqueta\",\"type\":\"record\"}}],\"name\":\"Andreani.Corporativo.Events.Record.RemitenteEtiqueta\",\"type\":\"record\"}"
}

func (r RemitenteEtiqueta) SchemaName() string {
	return "Andreani.Corporativo.Events.Record.RemitenteEtiqueta"
}

func (_ RemitenteEtiqueta) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ RemitenteEtiqueta) SetInt(v int32)       { panic("Unsupported operation") }
func (_ RemitenteEtiqueta) SetLong(v int64)      { panic("Unsupported operation") }
func (_ RemitenteEtiqueta) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ RemitenteEtiqueta) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ RemitenteEtiqueta) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ RemitenteEtiqueta) SetString(v string)   { panic("Unsupported operation") }
func (_ RemitenteEtiqueta) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RemitenteEtiqueta) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Nombre}

		return w

	case 1:
		w := types.String{Target: &r.Apellido}

		return w

	case 2:
		w := types.String{Target: &r.Telefono}

		return w

	case 3:
		w := types.String{Target: &r.DNI}

		return w

	case 4:
		w := types.String{Target: &r.Email}

		return w

	case 5:
		r.Direccion = NewDireccionRemitenteEtiqueta()

		w := types.Record{Target: &r.Direccion}

		return w

	}
	panic("Unknown field index")
}

func (r *RemitenteEtiqueta) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *RemitenteEtiqueta) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ RemitenteEtiqueta) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ RemitenteEtiqueta) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ RemitenteEtiqueta) HintSize(int)                     { panic("Unsupported operation") }
func (_ RemitenteEtiqueta) Finalize()                        {}

func (_ RemitenteEtiqueta) AvroCRC64Fingerprint() []byte {
	return []byte(RemitenteEtiquetaAvroCRC64Fingerprint)
}

func (r RemitenteEtiqueta) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	output["Apellido"], err = json.Marshal(r.Apellido)
	if err != nil {
		return nil, err
	}
	output["Telefono"], err = json.Marshal(r.Telefono)
	if err != nil {
		return nil, err
	}
	output["DNI"], err = json.Marshal(r.DNI)
	if err != nil {
		return nil, err
	}
	output["Email"], err = json.Marshal(r.Email)
	if err != nil {
		return nil, err
	}
	output["Direccion"], err = json.Marshal(r.Direccion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *RemitenteEtiqueta) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Nombre"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Nombre); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Nombre")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Apellido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Apellido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Apellido")
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
		if v, ok := fields["DNI"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DNI); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DNI")
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
		if v, ok := fields["Direccion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Direccion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Direccion")
	}
	return nil
}
