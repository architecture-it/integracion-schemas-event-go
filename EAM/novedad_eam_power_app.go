// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NovedadEAMPowerApp.avsc
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

type NovedadEAMPowerApp struct {
	Tipo string `json:"Tipo"`

	IdPowerApp string `json:"IdPowerApp"`

	NuevoEstado string `json:"NuevoEstado"`

	OT *UnionNullWorkOrder `json:"OT"`

	Equipo *UnionNullAsset `json:"Equipo"`
}

const NovedadEAMPowerAppAvroCRC64Fingerprint = "\xeb]\x85\xcbS?*1"

func NewNovedadEAMPowerApp() NovedadEAMPowerApp {
	r := NovedadEAMPowerApp{}
	r.OT = nil
	r.Equipo = nil
	return r
}

func DeserializeNovedadEAMPowerApp(r io.Reader) (NovedadEAMPowerApp, error) {
	t := NewNovedadEAMPowerApp()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNovedadEAMPowerAppFromSchema(r io.Reader, schema string) (NovedadEAMPowerApp, error) {
	t := NewNovedadEAMPowerApp()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNovedadEAMPowerApp(r NovedadEAMPowerApp, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Tipo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.IdPowerApp, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NuevoEstado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullWorkOrder(r.OT, w)
	if err != nil {
		return err
	}
	err = writeUnionNullAsset(r.Equipo, w)
	if err != nil {
		return err
	}
	return err
}

func (r NovedadEAMPowerApp) Serialize(w io.Writer) error {
	return writeNovedadEAMPowerApp(r, w)
}

func (r NovedadEAMPowerApp) Schema() string {
	return "{\"fields\":[{\"name\":\"Tipo\",\"type\":\"string\"},{\"name\":\"IdPowerApp\",\"type\":\"string\"},{\"name\":\"NuevoEstado\",\"type\":\"string\"},{\"default\":null,\"name\":\"OT\",\"type\":[\"null\",{\"fields\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"id_equipo\",\"type\":\"string\"},{\"name\":\"planta\",\"type\":\"string\"},{\"name\":\"descripcion\",\"type\":\"string\"},{\"name\":\"user_report\",\"type\":\"string\"}],\"name\":\"WorkOrder\",\"namespace\":\"Andreani.EAM.Events.Sharepoint\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Equipo\",\"type\":[\"null\",{\"fields\":[{\"name\":\"id\",\"type\":\"int\"},{\"name\":\"tipo_objeto\",\"type\":\"string\"},{\"name\":\"descripcion\",\"type\":\"string\"},{\"name\":\"clase\",\"type\":\"string\"},{\"name\":\"codigo_costo\",\"type\":\"string\"},{\"name\":\"estado\",\"type\":\"string\"},{\"name\":\"fecha_alta\",\"type\":\"string\"},{\"name\":\"organizacion\",\"type\":\"string\"},{\"name\":\"fabricante\",\"type\":\"string\"},{\"name\":\"modelo\",\"type\":\"string\"},{\"name\":\"nro_serie\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"fueraDeServicio\",\"type\":\"boolean\"},{\"name\":\"propulsion\",\"type\":\"string\"},{\"name\":\"cod_eam\",\"type\":[\"null\",\"string\"]}],\"name\":\"Asset\",\"namespace\":\"Andreani.EAM.Events.Sharepoint\",\"type\":\"record\"}]}],\"name\":\"Andreani.EAM.Events.Record.NovedadEAMPowerApp\",\"type\":\"record\"}"
}

func (r NovedadEAMPowerApp) SchemaName() string {
	return "Andreani.EAM.Events.Record.NovedadEAMPowerApp"
}

func (_ NovedadEAMPowerApp) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NovedadEAMPowerApp) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NovedadEAMPowerApp) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NovedadEAMPowerApp) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NovedadEAMPowerApp) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NovedadEAMPowerApp) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NovedadEAMPowerApp) SetString(v string)   { panic("Unsupported operation") }
func (_ NovedadEAMPowerApp) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NovedadEAMPowerApp) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Tipo}

		return w

	case 1:
		w := types.String{Target: &r.IdPowerApp}

		return w

	case 2:
		w := types.String{Target: &r.NuevoEstado}

		return w

	case 3:
		r.OT = NewUnionNullWorkOrder()

		return r.OT
	case 4:
		r.Equipo = NewUnionNullAsset()

		return r.Equipo
	}
	panic("Unknown field index")
}

func (r *NovedadEAMPowerApp) SetDefault(i int) {
	switch i {
	case 3:
		r.OT = nil
		return
	case 4:
		r.Equipo = nil
		return
	}
	panic("Unknown field index")
}

func (r *NovedadEAMPowerApp) NullField(i int) {
	switch i {
	case 3:
		r.OT = nil
		return
	case 4:
		r.Equipo = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ NovedadEAMPowerApp) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NovedadEAMPowerApp) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NovedadEAMPowerApp) HintSize(int)                     { panic("Unsupported operation") }
func (_ NovedadEAMPowerApp) Finalize()                        {}

func (_ NovedadEAMPowerApp) AvroCRC64Fingerprint() []byte {
	return []byte(NovedadEAMPowerAppAvroCRC64Fingerprint)
}

func (r NovedadEAMPowerApp) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Tipo"], err = json.Marshal(r.Tipo)
	if err != nil {
		return nil, err
	}
	output["IdPowerApp"], err = json.Marshal(r.IdPowerApp)
	if err != nil {
		return nil, err
	}
	output["NuevoEstado"], err = json.Marshal(r.NuevoEstado)
	if err != nil {
		return nil, err
	}
	output["OT"], err = json.Marshal(r.OT)
	if err != nil {
		return nil, err
	}
	output["Equipo"], err = json.Marshal(r.Equipo)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NovedadEAMPowerApp) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Tipo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tipo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Tipo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["IdPowerApp"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdPowerApp); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IdPowerApp")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NuevoEstado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NuevoEstado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NuevoEstado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OT"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OT); err != nil {
			return err
		}
	} else {
		r.OT = NewUnionNullWorkOrder()

		r.OT = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Equipo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Equipo); err != nil {
			return err
		}
	} else {
		r.Equipo = NewUnionNullAsset()

		r.Equipo = nil
	}
	return nil
}
