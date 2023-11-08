// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     SugerenciaContenedores.avsc
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

type SugerenciaContenedores struct {
	Identificacion Identificacion `json:"Identificacion"`

	OrdenWh string `json:"OrdenWh"`

	OrdenCliente string `json:"OrdenCliente"`

	Articulos []ArticuloContenedorPreparacion `json:"Articulos"`

	ContenedoresEmbalajeDeAlmacen []ContenedorEmbalaje `json:"ContenedoresEmbalajeDeAlmacen"`
}

const SugerenciaContenedoresAvroCRC64Fingerprint = "F*\xa1%U.\x1b\x94"

func NewSugerenciaContenedores() SugerenciaContenedores {
	r := SugerenciaContenedores{}
	r.Identificacion = NewIdentificacion()

	r.Articulos = make([]ArticuloContenedorPreparacion, 0)

	r.ContenedoresEmbalajeDeAlmacen = make([]ContenedorEmbalaje, 0)

	return r
}

func DeserializeSugerenciaContenedores(r io.Reader) (SugerenciaContenedores, error) {
	t := NewSugerenciaContenedores()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSugerenciaContenedoresFromSchema(r io.Reader, schema string) (SugerenciaContenedores, error) {
	t := NewSugerenciaContenedores()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSugerenciaContenedores(r SugerenciaContenedores, w io.Writer) error {
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
	err = writeArrayArticuloContenedorPreparacion(r.Articulos, w)
	if err != nil {
		return err
	}
	err = writeArrayContenedorEmbalaje(r.ContenedoresEmbalajeDeAlmacen, w)
	if err != nil {
		return err
	}
	return err
}

func (r SugerenciaContenedores) Serialize(w io.Writer) error {
	return writeSugerenciaContenedores(r, w)
}

func (r SugerenciaContenedores) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"},{\"default\":null,\"name\":\"PlantaOperacionId\",\"type\":[\"null\",\"int\"]}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.Empaquetado.Events.Common\",\"type\":\"record\"}},{\"name\":\"OrdenWh\",\"type\":\"string\"},{\"name\":\"OrdenCliente\",\"type\":\"string\"},{\"name\":\"Articulos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Sku\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"default\":null,\"name\":\"Ean\",\"type\":[\"null\",\"string\"]},{\"name\":\"NroLineaPedido\",\"type\":\"string\"},{\"name\":\"CantidadPedido\",\"type\":\"int\"},{\"name\":\"CantidadPickeada\",\"type\":\"int\"},{\"default\":null,\"name\":\"Lote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Serie\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Zona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoZona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DescripcionZona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Longitud\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Altura\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Ancho\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Peso\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"InstruccionesEmbalaje\",\"type\":[\"null\",\"string\"]}],\"name\":\"ArticuloContenedorPreparacion\",\"namespace\":\"Andreani.Empaquetado.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}},{\"name\":\"ContenedoresEmbalajeDeAlmacen\",\"type\":{\"items\":{\"fields\":[{\"name\":\"ContenedorId\",\"type\":\"string\"},{\"name\":\"ContenedorDescripcion\",\"type\":\"string\"},{\"default\":null,\"name\":\"Longuitud\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Altura\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Ancho\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Peso\",\"type\":[\"null\",\"float\"]},{\"name\":\"EsRetornable\",\"type\":\"boolean\"}],\"name\":\"ContenedorEmbalaje\",\"namespace\":\"Andreani.Empaquetado.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.Empaquetado.Events.Record.SugerenciaContenedores\",\"type\":\"record\"}"
}

func (r SugerenciaContenedores) SchemaName() string {
	return "Andreani.Empaquetado.Events.Record.SugerenciaContenedores"
}

func (_ SugerenciaContenedores) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ SugerenciaContenedores) SetInt(v int32)       { panic("Unsupported operation") }
func (_ SugerenciaContenedores) SetLong(v int64)      { panic("Unsupported operation") }
func (_ SugerenciaContenedores) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ SugerenciaContenedores) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ SugerenciaContenedores) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ SugerenciaContenedores) SetString(v string)   { panic("Unsupported operation") }
func (_ SugerenciaContenedores) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *SugerenciaContenedores) Get(i int) types.Field {
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
		r.Articulos = make([]ArticuloContenedorPreparacion, 0)

		w := ArrayArticuloContenedorPreparacionWrapper{Target: &r.Articulos}

		return w

	case 4:
		r.ContenedoresEmbalajeDeAlmacen = make([]ContenedorEmbalaje, 0)

		w := ArrayContenedorEmbalajeWrapper{Target: &r.ContenedoresEmbalajeDeAlmacen}

		return w

	}
	panic("Unknown field index")
}

func (r *SugerenciaContenedores) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *SugerenciaContenedores) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ SugerenciaContenedores) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ SugerenciaContenedores) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ SugerenciaContenedores) HintSize(int)                     { panic("Unsupported operation") }
func (_ SugerenciaContenedores) Finalize()                        {}

func (_ SugerenciaContenedores) AvroCRC64Fingerprint() []byte {
	return []byte(SugerenciaContenedoresAvroCRC64Fingerprint)
}

func (r SugerenciaContenedores) MarshalJSON() ([]byte, error) {
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
	output["Articulos"], err = json.Marshal(r.Articulos)
	if err != nil {
		return nil, err
	}
	output["ContenedoresEmbalajeDeAlmacen"], err = json.Marshal(r.ContenedoresEmbalajeDeAlmacen)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *SugerenciaContenedores) UnmarshalJSON(data []byte) error {
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
	val = func() json.RawMessage {
		if v, ok := fields["ContenedoresEmbalajeDeAlmacen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContenedoresEmbalajeDeAlmacen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ContenedoresEmbalajeDeAlmacen")
	}
	return nil
}
