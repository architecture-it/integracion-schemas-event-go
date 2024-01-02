// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NovedadAnalizada.avsc
 */
package NovedadesCustomEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type NovedadAnalizada struct {
	NovedadUnificada NovedadUnificada `json:"novedadUnificada"`

	IdCliente string `json:"idCliente"`

	NombreDeCliente string `json:"nombreDeCliente"`

	IdModelo int32 `json:"idModelo"`

	GeneradorDestino *UnionNullString `json:"generadorDestino"`

	DespachadorDestino *UnionNullString `json:"despachadorDestino"`
}

const NovedadAnalizadaAvroCRC64Fingerprint = "\xb0\xcfV#\xed\xd9E\xc0"

func NewNovedadAnalizada() NovedadAnalizada {
	r := NovedadAnalizada{}
	r.NovedadUnificada = NewNovedadUnificada()

	r.GeneradorDestino = nil
	r.DespachadorDestino = nil
	return r
}

func DeserializeNovedadAnalizada(r io.Reader) (NovedadAnalizada, error) {
	t := NewNovedadAnalizada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNovedadAnalizadaFromSchema(r io.Reader, schema string) (NovedadAnalizada, error) {
	t := NewNovedadAnalizada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNovedadAnalizada(r NovedadAnalizada, w io.Writer) error {
	var err error
	err = writeNovedadUnificada(r.NovedadUnificada, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.IdCliente, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NombreDeCliente, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.IdModelo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.GeneradorDestino, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DespachadorDestino, w)
	if err != nil {
		return err
	}
	return err
}

func (r NovedadAnalizada) Serialize(w io.Writer) error {
	return writeNovedadAnalizada(r, w)
}

func (r NovedadAnalizada) Schema() string {
	return "{\"fields\":[{\"name\":\"novedadUnificada\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"Traza\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Evento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeContratoInterno\",\"type\":[\"null\",\"string\"]},{\"name\":\"Cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"CicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Motivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SubMotivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Remitente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"EstadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SucursalAsociada\",\"type\":[\"null\",\"string\"]}],\"name\":\"Traza\",\"namespace\":\"Andreani.NovedadesCustom.Events.Common\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"DatosAdicionales\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"default\":null,\"name\":\"Key\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Value\",\"type\":[\"null\",\"string\"]}],\"name\":\"Metadata\",\"namespace\":\"Andreani.NovedadesCustom.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"NovedadUnificada\",\"type\":\"record\"}},{\"name\":\"idCliente\",\"type\":\"string\"},{\"name\":\"nombreDeCliente\",\"type\":\"string\"},{\"name\":\"idModelo\",\"type\":\"int\"},{\"default\":null,\"name\":\"generadorDestino\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"despachadorDestino\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.NovedadesCustom.Events.Record.NovedadAnalizada\",\"type\":\"record\"}"
}

func (r NovedadAnalizada) SchemaName() string {
	return "Andreani.NovedadesCustom.Events.Record.NovedadAnalizada"
}

func (_ NovedadAnalizada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NovedadAnalizada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NovedadAnalizada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NovedadAnalizada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NovedadAnalizada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NovedadAnalizada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NovedadAnalizada) SetString(v string)   { panic("Unsupported operation") }
func (_ NovedadAnalizada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NovedadAnalizada) Get(i int) types.Field {
	switch i {
	case 0:
		r.NovedadUnificada = NewNovedadUnificada()

		w := types.Record{Target: &r.NovedadUnificada}

		return w

	case 1:
		w := types.String{Target: &r.IdCliente}

		return w

	case 2:
		w := types.String{Target: &r.NombreDeCliente}

		return w

	case 3:
		w := types.Int{Target: &r.IdModelo}

		return w

	case 4:
		r.GeneradorDestino = NewUnionNullString()

		return r.GeneradorDestino
	case 5:
		r.DespachadorDestino = NewUnionNullString()

		return r.DespachadorDestino
	}
	panic("Unknown field index")
}

func (r *NovedadAnalizada) SetDefault(i int) {
	switch i {
	case 4:
		r.GeneradorDestino = nil
		return
	case 5:
		r.DespachadorDestino = nil
		return
	}
	panic("Unknown field index")
}

func (r *NovedadAnalizada) NullField(i int) {
	switch i {
	case 4:
		r.GeneradorDestino = nil
		return
	case 5:
		r.DespachadorDestino = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ NovedadAnalizada) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NovedadAnalizada) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NovedadAnalizada) HintSize(int)                     { panic("Unsupported operation") }
func (_ NovedadAnalizada) Finalize()                        {}

func (_ NovedadAnalizada) AvroCRC64Fingerprint() []byte {
	return []byte(NovedadAnalizadaAvroCRC64Fingerprint)
}

func (r NovedadAnalizada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["novedadUnificada"], err = json.Marshal(r.NovedadUnificada)
	if err != nil {
		return nil, err
	}
	output["idCliente"], err = json.Marshal(r.IdCliente)
	if err != nil {
		return nil, err
	}
	output["nombreDeCliente"], err = json.Marshal(r.NombreDeCliente)
	if err != nil {
		return nil, err
	}
	output["idModelo"], err = json.Marshal(r.IdModelo)
	if err != nil {
		return nil, err
	}
	output["generadorDestino"], err = json.Marshal(r.GeneradorDestino)
	if err != nil {
		return nil, err
	}
	output["despachadorDestino"], err = json.Marshal(r.DespachadorDestino)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NovedadAnalizada) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["novedadUnificada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NovedadUnificada); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for novedadUnificada")
	}
	val = func() json.RawMessage {
		if v, ok := fields["idCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdCliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idCliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["nombreDeCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NombreDeCliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for nombreDeCliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["idModelo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdModelo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idModelo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["generadorDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.GeneradorDestino); err != nil {
			return err
		}
	} else {
		r.GeneradorDestino = NewUnionNullString()

		r.GeneradorDestino = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["despachadorDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DespachadorDestino); err != nil {
			return err
		}
	} else {
		r.DespachadorDestino = NewUnionNullString()

		r.DespachadorDestino = nil
	}
	return nil
}
