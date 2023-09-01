// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     WosTrazaAnmat.avsc
 */
package WosTrazabilidadEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type WosTrazaAnmat struct {
	Cabecera Cabecera `json:"Cabecera"`

	Lineas []Linea `json:"Lineas"`

	RespuestaAnmat *UnionNullRespuestaAnmat `json:"RespuestaAnmat"`
}

const WosTrazaAnmatAvroCRC64Fingerprint = "\xc9Y\x95\xa8\xcb&z\xd7"

func NewWosTrazaAnmat() WosTrazaAnmat {
	r := WosTrazaAnmat{}
	r.Cabecera = NewCabecera()

	r.Lineas = make([]Linea, 0)

	r.RespuestaAnmat = nil
	return r
}

func DeserializeWosTrazaAnmat(r io.Reader) (WosTrazaAnmat, error) {
	t := NewWosTrazaAnmat()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeWosTrazaAnmatFromSchema(r io.Reader, schema string) (WosTrazaAnmat, error) {
	t := NewWosTrazaAnmat()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeWosTrazaAnmat(r WosTrazaAnmat, w io.Writer) error {
	var err error
	err = writeCabecera(r.Cabecera, w)
	if err != nil {
		return err
	}
	err = writeArrayLinea(r.Lineas, w)
	if err != nil {
		return err
	}
	err = writeUnionNullRespuestaAnmat(r.RespuestaAnmat, w)
	if err != nil {
		return err
	}
	return err
}

func (r WosTrazaAnmat) Serialize(w io.Writer) error {
	return writeWosTrazaAnmat(r, w)
}

func (r WosTrazaAnmat) Schema() string {
	return "{\"fields\":[{\"name\":\"Cabecera\",\"type\":{\"fields\":[{\"name\":\"Serialkey\",\"type\":\"int\"},{\"name\":\"almacen\",\"type\":\"string\"},{\"name\":\"instancia\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"tipoDocumento\",\"type\":\"int\"},{\"name\":\"nroDocumento\",\"type\":\"string\"},{\"name\":\"nroDocumentoWMS\",\"type\":\"string\"},{\"name\":\"glnOrigen\",\"type\":\"string\"},{\"name\":\"glnDestino\",\"type\":\"string\"},{\"name\":\"estado\",\"type\":\"int\"},{\"name\":\"descripcionEstado\",\"type\":\"string\"},{\"name\":\"remito\",\"type\":\"string\"},{\"name\":\"idEventoAnmat\",\"type\":\"string\"}],\"name\":\"Cabecera\",\"namespace\":\"Andreani.WosTrazabilidad.Events.AnmatCommon\",\"type\":\"record\"}},{\"name\":\"Lineas\",\"type\":{\"items\":{\"fields\":[{\"name\":\"numeroDeLinea\",\"type\":\"int\"},{\"name\":\"numeroDeLineaWMS\",\"type\":\"string\"},{\"name\":\"numeroDeLineaCliente\",\"type\":\"string\"},{\"name\":\"sku\",\"type\":\"string\"},{\"name\":\"loteCajita\",\"type\":\"string\"},{\"name\":\"loteSecundario\",\"type\":\"string\"},{\"name\":\"fechaVencimiento\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"productoTrazable\",\"type\":\"boolean\"},{\"name\":\"cantidad\",\"type\":\"int\"},{\"name\":\"estado\",\"type\":\"int\"},{\"name\":\"descripcionEstado\",\"type\":\"string\"},{\"name\":\"gtin\",\"type\":\"string\"},{\"name\":\"Series\",\"type\":{\"items\":{\"fields\":[{\"name\":\"serie\",\"type\":\"string\"},{\"default\":null,\"name\":\"estado\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"descripcionEstado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoTransaccion\",\"type\":[\"null\",\"string\"]}],\"name\":\"Serie\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Linea\",\"namespace\":\"Andreani.WosTrazabilidad.Events.AnmatCommon\",\"type\":\"record\"},\"type\":\"array\"}},{\"default\":null,\"name\":\"RespuestaAnmat\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"codigoTransaccion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"resultado\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"errores\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]}],\"name\":\"RespuestaAnmat\",\"namespace\":\"Andreani.WosTrazabilidad.Events.AnmatCommon\",\"type\":\"record\"}]}],\"name\":\"Andreani.WosTrazabilidad.Events.Record.WosTrazaAnmat\",\"type\":\"record\"}"
}

func (r WosTrazaAnmat) SchemaName() string {
	return "Andreani.WosTrazabilidad.Events.Record.WosTrazaAnmat"
}

func (_ WosTrazaAnmat) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ WosTrazaAnmat) SetInt(v int32)       { panic("Unsupported operation") }
func (_ WosTrazaAnmat) SetLong(v int64)      { panic("Unsupported operation") }
func (_ WosTrazaAnmat) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ WosTrazaAnmat) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ WosTrazaAnmat) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ WosTrazaAnmat) SetString(v string)   { panic("Unsupported operation") }
func (_ WosTrazaAnmat) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *WosTrazaAnmat) Get(i int) types.Field {
	switch i {
	case 0:
		r.Cabecera = NewCabecera()

		w := types.Record{Target: &r.Cabecera}

		return w

	case 1:
		r.Lineas = make([]Linea, 0)

		w := ArrayLineaWrapper{Target: &r.Lineas}

		return w

	case 2:
		r.RespuestaAnmat = NewUnionNullRespuestaAnmat()

		return r.RespuestaAnmat
	}
	panic("Unknown field index")
}

func (r *WosTrazaAnmat) SetDefault(i int) {
	switch i {
	case 2:
		r.RespuestaAnmat = nil
		return
	}
	panic("Unknown field index")
}

func (r *WosTrazaAnmat) NullField(i int) {
	switch i {
	case 2:
		r.RespuestaAnmat = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ WosTrazaAnmat) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ WosTrazaAnmat) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ WosTrazaAnmat) HintSize(int)                     { panic("Unsupported operation") }
func (_ WosTrazaAnmat) Finalize()                        {}

func (_ WosTrazaAnmat) AvroCRC64Fingerprint() []byte {
	return []byte(WosTrazaAnmatAvroCRC64Fingerprint)
}

func (r WosTrazaAnmat) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Cabecera"], err = json.Marshal(r.Cabecera)
	if err != nil {
		return nil, err
	}
	output["Lineas"], err = json.Marshal(r.Lineas)
	if err != nil {
		return nil, err
	}
	output["RespuestaAnmat"], err = json.Marshal(r.RespuestaAnmat)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *WosTrazaAnmat) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
	val = func() json.RawMessage {
		if v, ok := fields["Lineas"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Lineas); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Lineas")
	}
	val = func() json.RawMessage {
		if v, ok := fields["RespuestaAnmat"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RespuestaAnmat); err != nil {
			return err
		}
	} else {
		r.RespuestaAnmat = NewUnionNullRespuestaAnmat()

		r.RespuestaAnmat = nil
	}
	return nil
}
