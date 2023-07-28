// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ArticulosDeBulto.avsc
 */
package AuditoriaEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ArticulosDeBulto struct {
	Identificacion Identificacion `json:"Identificacion"`

	OrdenWh string `json:"OrdenWh"`

	OrdenCliente string `json:"OrdenCliente"`

	CodigoEmbalaje string `json:"CodigoEmbalaje"`

	Articulos []ArticuloAuditar `json:"Articulos"`
}

const ArticulosDeBultoAvroCRC64Fingerprint = "6zDǴ\xa7\xa7D"

func NewArticulosDeBulto() ArticulosDeBulto {
	r := ArticulosDeBulto{}
	r.Identificacion = NewIdentificacion()

	r.Articulos = make([]ArticuloAuditar, 0)

	return r
}

func DeserializeArticulosDeBulto(r io.Reader) (ArticulosDeBulto, error) {
	t := NewArticulosDeBulto()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeArticulosDeBultoFromSchema(r io.Reader, schema string) (ArticulosDeBulto, error) {
	t := NewArticulosDeBulto()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeArticulosDeBulto(r ArticulosDeBulto, w io.Writer) error {
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
	err = vm.WriteString(r.CodigoEmbalaje, w)
	if err != nil {
		return err
	}
	err = writeArrayArticuloAuditar(r.Articulos, w)
	if err != nil {
		return err
	}
	return err
}

func (r ArticulosDeBulto) Serialize(w io.Writer) error {
	return writeArticulosDeBulto(r, w)
}

func (r ArticulosDeBulto) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"},{\"default\":null,\"name\":\"PlantaOperacionId\",\"type\":[\"null\",\"int\"]}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.Auditoria.Events.Common\",\"type\":\"record\"}},{\"name\":\"OrdenWh\",\"type\":\"string\"},{\"name\":\"OrdenCliente\",\"type\":\"string\"},{\"name\":\"CodigoEmbalaje\",\"type\":\"string\"},{\"name\":\"Articulos\",\"type\":{\"items\":{\"fields\":[{\"default\":null,\"name\":\"Ean\",\"type\":[\"null\",\"string\"]},{\"name\":\"CodigoCliente\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"name\":\"CantidadControlada\",\"type\":\"int\"},{\"name\":\"CantidadPickeada\",\"type\":\"int\"},{\"name\":\"NroLineaPedido\",\"type\":\"string\"}],\"name\":\"ArticuloAuditar\",\"namespace\":\"Andreani.Auditoria.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.Auditoria.Events.Record.ArticulosDeBulto\",\"type\":\"record\"}"
}

func (r ArticulosDeBulto) SchemaName() string {
	return "Andreani.Auditoria.Events.Record.ArticulosDeBulto"
}

func (_ ArticulosDeBulto) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ArticulosDeBulto) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ArticulosDeBulto) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ArticulosDeBulto) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ArticulosDeBulto) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ArticulosDeBulto) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ArticulosDeBulto) SetString(v string)   { panic("Unsupported operation") }
func (_ ArticulosDeBulto) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ArticulosDeBulto) Get(i int) types.Field {
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
		w := types.String{Target: &r.CodigoEmbalaje}

		return w

	case 4:
		r.Articulos = make([]ArticuloAuditar, 0)

		w := ArrayArticuloAuditarWrapper{Target: &r.Articulos}

		return w

	}
	panic("Unknown field index")
}

func (r *ArticulosDeBulto) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ArticulosDeBulto) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ArticulosDeBulto) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArticulosDeBulto) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ArticulosDeBulto) HintSize(int)                     { panic("Unsupported operation") }
func (_ ArticulosDeBulto) Finalize()                        {}

func (_ ArticulosDeBulto) AvroCRC64Fingerprint() []byte {
	return []byte(ArticulosDeBultoAvroCRC64Fingerprint)
}

func (r ArticulosDeBulto) MarshalJSON() ([]byte, error) {
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
	output["CodigoEmbalaje"], err = json.Marshal(r.CodigoEmbalaje)
	if err != nil {
		return nil, err
	}
	output["Articulos"], err = json.Marshal(r.Articulos)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ArticulosDeBulto) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["CodigoEmbalaje"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoEmbalaje); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoEmbalaje")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Articulos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Articulos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Articulos")
	}
	return nil
}
