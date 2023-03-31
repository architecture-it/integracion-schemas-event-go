// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Tarea.avsc
 */
package EventoWhTareasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Tarea struct {
	IDTarea string `json:"IDTarea"`

	OrdenWH *UnionNullString `json:"OrdenWH"`

	TipoDeTarea string `json:"TipoDeTarea"`

	Propietario string `json:"Propietario"`

	SKU string `json:"SKU"`

	LoteInternoWH string `json:"LoteInternoWH"`

	CantidadTarea string `json:"CantidadTarea"`

	UbicacionInicial string `json:"UbicacionInicial"`

	IDInicial *UnionNullString `json:"IDInicial"`

	UbicacionFinal string `json:"UbicacionFinal"`

	IDFinal string `json:"IDFinal"`

	Estado string `json:"Estado"`

	Prioridad string `json:"Prioridad"`

	Usuario string `json:"Usuario"`

	FechaInicioTarea string `json:"FechaInicioTarea"`

	FechaFinTarea string `json:"FechaFinTarea"`

	LineaDeOrden *UnionNullString `json:"LineaDeOrden"`

	NumeroDeOleadaWH *UnionNullString `json:"NumeroDeOleadaWH"`

	FechaDeCreacion string `json:"FechaDeCreacion"`

	UsuarioCreacion string `json:"UsuarioCreacion"`

	FechaEdicion string `json:"FechaEdicion"`

	UsuarioEdicion string `json:"UsuarioEdicion"`

	PesoTara string `json:"PesoTara"`

	PesoNeto string `json:"PesoNeto"`

	PesoBruto string `json:"PesoBruto"`

	NumeroDeAsignacionWH *UnionNullString `json:"NumeroDeAsignacionWH"`

	TipoDeOrigen *UnionNullString `json:"TipoDeOrigen"`

	LoteSecundario string `json:"LoteSecundario"`

	FechaFabricacion *UnionNullLong `json:"FechaFabricacion"`

	FechaVencimiento *UnionNullLong `json:"FechaVencimiento"`

	ProductoTrazable string `json:"ProductoTrazable"`

	AlmacenConsumo string `json:"AlmacenConsumo"`

	EstadoLote string `json:"EstadoLote"`

	BloqueoUbicacion string `json:"BloqueoUbicacion"`

	VidaUtilLote string `json:"VidaUtilLote"`

	EntregaAntesDe string `json:"EntregaAntesDe"`

	ConsumoAntesDe string `json:"ConsumoAntesDe"`
}

const TareaAvroCRC64Fingerprint = "\x15\xba>\xf9\xa2\xda\xff\xb4"

func NewTarea() Tarea {
	r := Tarea{}
	r.OrdenWH = nil
	r.IDInicial = nil
	r.LineaDeOrden = nil
	r.NumeroDeOleadaWH = nil
	r.NumeroDeAsignacionWH = nil
	r.TipoDeOrigen = nil
	r.FechaFabricacion = nil
	r.FechaVencimiento = nil
	return r
}

func DeserializeTarea(r io.Reader) (Tarea, error) {
	t := NewTarea()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTareaFromSchema(r io.Reader, schema string) (Tarea, error) {
	t := NewTarea()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTarea(r Tarea, w io.Writer) error {
	var err error
	err = vm.WriteString(r.IDTarea, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.OrdenWH, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TipoDeTarea, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.SKU, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.LoteInternoWH, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CantidadTarea, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.UbicacionInicial, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.IDInicial, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.UbicacionFinal, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.IDFinal, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Estado, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Prioridad, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Usuario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaInicioTarea, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaFinTarea, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LineaDeOrden, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeOleadaWH, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaDeCreacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.UsuarioCreacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaEdicion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.UsuarioEdicion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.PesoTara, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.PesoNeto, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.PesoBruto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeAsignacionWH, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoDeOrigen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.LoteSecundario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaFabricacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaVencimiento, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ProductoTrazable, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.AlmacenConsumo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EstadoLote, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.BloqueoUbicacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.VidaUtilLote, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EntregaAntesDe, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ConsumoAntesDe, w)
	if err != nil {
		return err
	}
	return err
}

func (r Tarea) Serialize(w io.Writer) error {
	return writeTarea(r, w)
}

func (r Tarea) Schema() string {
	return "{\"fields\":[{\"name\":\"IDTarea\",\"type\":\"string\"},{\"default\":null,\"name\":\"OrdenWH\",\"type\":[\"null\",\"string\"]},{\"name\":\"TipoDeTarea\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"SKU\",\"type\":\"string\"},{\"name\":\"LoteInternoWH\",\"type\":\"string\"},{\"name\":\"CantidadTarea\",\"type\":\"string\"},{\"name\":\"UbicacionInicial\",\"type\":\"string\"},{\"default\":null,\"name\":\"IDInicial\",\"type\":[\"null\",\"string\"]},{\"name\":\"UbicacionFinal\",\"type\":\"string\"},{\"name\":\"IDFinal\",\"type\":\"string\"},{\"name\":\"Estado\",\"type\":\"string\"},{\"name\":\"Prioridad\",\"type\":\"string\"},{\"name\":\"Usuario\",\"type\":\"string\"},{\"name\":\"FechaInicioTarea\",\"type\":\"string\"},{\"name\":\"FechaFinTarea\",\"type\":\"string\"},{\"default\":null,\"name\":\"LineaDeOrden\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeOleadaWH\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaDeCreacion\",\"type\":\"string\"},{\"name\":\"UsuarioCreacion\",\"type\":\"string\"},{\"name\":\"FechaEdicion\",\"type\":\"string\"},{\"name\":\"UsuarioEdicion\",\"type\":\"string\"},{\"name\":\"PesoTara\",\"type\":\"string\"},{\"name\":\"PesoNeto\",\"type\":\"string\"},{\"name\":\"PesoBruto\",\"type\":\"string\"},{\"default\":null,\"name\":\"NumeroDeAsignacionWH\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeOrigen\",\"type\":[\"null\",\"string\"]},{\"name\":\"LoteSecundario\",\"type\":\"string\"},{\"default\":null,\"name\":\"FechaFabricacion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"FechaVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"name\":\"ProductoTrazable\",\"type\":\"string\"},{\"name\":\"AlmacenConsumo\",\"type\":\"string\"},{\"name\":\"EstadoLote\",\"type\":\"string\"},{\"name\":\"BloqueoUbicacion\",\"type\":\"string\"},{\"name\":\"VidaUtilLote\",\"type\":\"string\"},{\"name\":\"EntregaAntesDe\",\"type\":\"string\"},{\"name\":\"ConsumoAntesDe\",\"type\":\"string\"}],\"name\":\"Andreani.EventoWhTareas.Events.TareasMobilePickingCommon.Tarea\",\"type\":\"record\"}"
}

func (r Tarea) SchemaName() string {
	return "Andreani.EventoWhTareas.Events.TareasMobilePickingCommon.Tarea"
}

func (_ Tarea) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Tarea) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Tarea) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Tarea) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Tarea) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Tarea) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Tarea) SetString(v string)   { panic("Unsupported operation") }
func (_ Tarea) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Tarea) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.IDTarea}

		return w

	case 1:
		r.OrdenWH = NewUnionNullString()

		return r.OrdenWH
	case 2:
		w := types.String{Target: &r.TipoDeTarea}

		return w

	case 3:
		w := types.String{Target: &r.Propietario}

		return w

	case 4:
		w := types.String{Target: &r.SKU}

		return w

	case 5:
		w := types.String{Target: &r.LoteInternoWH}

		return w

	case 6:
		w := types.String{Target: &r.CantidadTarea}

		return w

	case 7:
		w := types.String{Target: &r.UbicacionInicial}

		return w

	case 8:
		r.IDInicial = NewUnionNullString()

		return r.IDInicial
	case 9:
		w := types.String{Target: &r.UbicacionFinal}

		return w

	case 10:
		w := types.String{Target: &r.IDFinal}

		return w

	case 11:
		w := types.String{Target: &r.Estado}

		return w

	case 12:
		w := types.String{Target: &r.Prioridad}

		return w

	case 13:
		w := types.String{Target: &r.Usuario}

		return w

	case 14:
		w := types.String{Target: &r.FechaInicioTarea}

		return w

	case 15:
		w := types.String{Target: &r.FechaFinTarea}

		return w

	case 16:
		r.LineaDeOrden = NewUnionNullString()

		return r.LineaDeOrden
	case 17:
		r.NumeroDeOleadaWH = NewUnionNullString()

		return r.NumeroDeOleadaWH
	case 18:
		w := types.String{Target: &r.FechaDeCreacion}

		return w

	case 19:
		w := types.String{Target: &r.UsuarioCreacion}

		return w

	case 20:
		w := types.String{Target: &r.FechaEdicion}

		return w

	case 21:
		w := types.String{Target: &r.UsuarioEdicion}

		return w

	case 22:
		w := types.String{Target: &r.PesoTara}

		return w

	case 23:
		w := types.String{Target: &r.PesoNeto}

		return w

	case 24:
		w := types.String{Target: &r.PesoBruto}

		return w

	case 25:
		r.NumeroDeAsignacionWH = NewUnionNullString()

		return r.NumeroDeAsignacionWH
	case 26:
		r.TipoDeOrigen = NewUnionNullString()

		return r.TipoDeOrigen
	case 27:
		w := types.String{Target: &r.LoteSecundario}

		return w

	case 28:
		r.FechaFabricacion = NewUnionNullLong()

		return r.FechaFabricacion
	case 29:
		r.FechaVencimiento = NewUnionNullLong()

		return r.FechaVencimiento
	case 30:
		w := types.String{Target: &r.ProductoTrazable}

		return w

	case 31:
		w := types.String{Target: &r.AlmacenConsumo}

		return w

	case 32:
		w := types.String{Target: &r.EstadoLote}

		return w

	case 33:
		w := types.String{Target: &r.BloqueoUbicacion}

		return w

	case 34:
		w := types.String{Target: &r.VidaUtilLote}

		return w

	case 35:
		w := types.String{Target: &r.EntregaAntesDe}

		return w

	case 36:
		w := types.String{Target: &r.ConsumoAntesDe}

		return w

	}
	panic("Unknown field index")
}

func (r *Tarea) SetDefault(i int) {
	switch i {
	case 1:
		r.OrdenWH = nil
		return
	case 8:
		r.IDInicial = nil
		return
	case 16:
		r.LineaDeOrden = nil
		return
	case 17:
		r.NumeroDeOleadaWH = nil
		return
	case 25:
		r.NumeroDeAsignacionWH = nil
		return
	case 26:
		r.TipoDeOrigen = nil
		return
	case 28:
		r.FechaFabricacion = nil
		return
	case 29:
		r.FechaVencimiento = nil
		return
	}
	panic("Unknown field index")
}

func (r *Tarea) NullField(i int) {
	switch i {
	case 1:
		r.OrdenWH = nil
		return
	case 8:
		r.IDInicial = nil
		return
	case 16:
		r.LineaDeOrden = nil
		return
	case 17:
		r.NumeroDeOleadaWH = nil
		return
	case 25:
		r.NumeroDeAsignacionWH = nil
		return
	case 26:
		r.TipoDeOrigen = nil
		return
	case 28:
		r.FechaFabricacion = nil
		return
	case 29:
		r.FechaVencimiento = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Tarea) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Tarea) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Tarea) HintSize(int)                     { panic("Unsupported operation") }
func (_ Tarea) Finalize()                        {}

func (_ Tarea) AvroCRC64Fingerprint() []byte {
	return []byte(TareaAvroCRC64Fingerprint)
}

func (r Tarea) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["IDTarea"], err = json.Marshal(r.IDTarea)
	if err != nil {
		return nil, err
	}
	output["OrdenWH"], err = json.Marshal(r.OrdenWH)
	if err != nil {
		return nil, err
	}
	output["TipoDeTarea"], err = json.Marshal(r.TipoDeTarea)
	if err != nil {
		return nil, err
	}
	output["Propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["SKU"], err = json.Marshal(r.SKU)
	if err != nil {
		return nil, err
	}
	output["LoteInternoWH"], err = json.Marshal(r.LoteInternoWH)
	if err != nil {
		return nil, err
	}
	output["CantidadTarea"], err = json.Marshal(r.CantidadTarea)
	if err != nil {
		return nil, err
	}
	output["UbicacionInicial"], err = json.Marshal(r.UbicacionInicial)
	if err != nil {
		return nil, err
	}
	output["IDInicial"], err = json.Marshal(r.IDInicial)
	if err != nil {
		return nil, err
	}
	output["UbicacionFinal"], err = json.Marshal(r.UbicacionFinal)
	if err != nil {
		return nil, err
	}
	output["IDFinal"], err = json.Marshal(r.IDFinal)
	if err != nil {
		return nil, err
	}
	output["Estado"], err = json.Marshal(r.Estado)
	if err != nil {
		return nil, err
	}
	output["Prioridad"], err = json.Marshal(r.Prioridad)
	if err != nil {
		return nil, err
	}
	output["Usuario"], err = json.Marshal(r.Usuario)
	if err != nil {
		return nil, err
	}
	output["FechaInicioTarea"], err = json.Marshal(r.FechaInicioTarea)
	if err != nil {
		return nil, err
	}
	output["FechaFinTarea"], err = json.Marshal(r.FechaFinTarea)
	if err != nil {
		return nil, err
	}
	output["LineaDeOrden"], err = json.Marshal(r.LineaDeOrden)
	if err != nil {
		return nil, err
	}
	output["NumeroDeOleadaWH"], err = json.Marshal(r.NumeroDeOleadaWH)
	if err != nil {
		return nil, err
	}
	output["FechaDeCreacion"], err = json.Marshal(r.FechaDeCreacion)
	if err != nil {
		return nil, err
	}
	output["UsuarioCreacion"], err = json.Marshal(r.UsuarioCreacion)
	if err != nil {
		return nil, err
	}
	output["FechaEdicion"], err = json.Marshal(r.FechaEdicion)
	if err != nil {
		return nil, err
	}
	output["UsuarioEdicion"], err = json.Marshal(r.UsuarioEdicion)
	if err != nil {
		return nil, err
	}
	output["PesoTara"], err = json.Marshal(r.PesoTara)
	if err != nil {
		return nil, err
	}
	output["PesoNeto"], err = json.Marshal(r.PesoNeto)
	if err != nil {
		return nil, err
	}
	output["PesoBruto"], err = json.Marshal(r.PesoBruto)
	if err != nil {
		return nil, err
	}
	output["NumeroDeAsignacionWH"], err = json.Marshal(r.NumeroDeAsignacionWH)
	if err != nil {
		return nil, err
	}
	output["TipoDeOrigen"], err = json.Marshal(r.TipoDeOrigen)
	if err != nil {
		return nil, err
	}
	output["LoteSecundario"], err = json.Marshal(r.LoteSecundario)
	if err != nil {
		return nil, err
	}
	output["FechaFabricacion"], err = json.Marshal(r.FechaFabricacion)
	if err != nil {
		return nil, err
	}
	output["FechaVencimiento"], err = json.Marshal(r.FechaVencimiento)
	if err != nil {
		return nil, err
	}
	output["ProductoTrazable"], err = json.Marshal(r.ProductoTrazable)
	if err != nil {
		return nil, err
	}
	output["AlmacenConsumo"], err = json.Marshal(r.AlmacenConsumo)
	if err != nil {
		return nil, err
	}
	output["EstadoLote"], err = json.Marshal(r.EstadoLote)
	if err != nil {
		return nil, err
	}
	output["BloqueoUbicacion"], err = json.Marshal(r.BloqueoUbicacion)
	if err != nil {
		return nil, err
	}
	output["VidaUtilLote"], err = json.Marshal(r.VidaUtilLote)
	if err != nil {
		return nil, err
	}
	output["EntregaAntesDe"], err = json.Marshal(r.EntregaAntesDe)
	if err != nil {
		return nil, err
	}
	output["ConsumoAntesDe"], err = json.Marshal(r.ConsumoAntesDe)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Tarea) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["IDTarea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IDTarea); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IDTarea")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OrdenWH"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenWH); err != nil {
			return err
		}
	} else {
		r.OrdenWH = NewUnionNullString()

		r.OrdenWH = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoDeTarea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeTarea); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoDeTarea")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Propietario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Propietario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Propietario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SKU"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SKU); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SKU")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LoteInternoWH"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteInternoWH); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LoteInternoWH")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CantidadTarea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CantidadTarea); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CantidadTarea")
	}
	val = func() json.RawMessage {
		if v, ok := fields["UbicacionInicial"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UbicacionInicial); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for UbicacionInicial")
	}
	val = func() json.RawMessage {
		if v, ok := fields["IDInicial"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IDInicial); err != nil {
			return err
		}
	} else {
		r.IDInicial = NewUnionNullString()

		r.IDInicial = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["UbicacionFinal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UbicacionFinal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for UbicacionFinal")
	}
	val = func() json.RawMessage {
		if v, ok := fields["IDFinal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IDFinal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IDFinal")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Estado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Estado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Prioridad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Prioridad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Prioridad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Usuario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Usuario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Usuario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaInicioTarea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaInicioTarea); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaInicioTarea")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaFinTarea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaFinTarea); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaFinTarea")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LineaDeOrden"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LineaDeOrden); err != nil {
			return err
		}
	} else {
		r.LineaDeOrden = NewUnionNullString()

		r.LineaDeOrden = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeOleadaWH"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeOleadaWH); err != nil {
			return err
		}
	} else {
		r.NumeroDeOleadaWH = NewUnionNullString()

		r.NumeroDeOleadaWH = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaDeCreacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaDeCreacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaDeCreacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["UsuarioCreacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UsuarioCreacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for UsuarioCreacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaEdicion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaEdicion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaEdicion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["UsuarioEdicion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UsuarioEdicion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for UsuarioEdicion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PesoTara"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PesoTara); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PesoTara")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PesoNeto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PesoNeto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PesoNeto")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PesoBruto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PesoBruto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PesoBruto")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeAsignacionWH"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeAsignacionWH); err != nil {
			return err
		}
	} else {
		r.NumeroDeAsignacionWH = NewUnionNullString()

		r.NumeroDeAsignacionWH = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoDeOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeOrigen); err != nil {
			return err
		}
	} else {
		r.TipoDeOrigen = NewUnionNullString()

		r.TipoDeOrigen = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["LoteSecundario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteSecundario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LoteSecundario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaFabricacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaFabricacion); err != nil {
			return err
		}
	} else {
		r.FechaFabricacion = NewUnionNullLong()

		r.FechaFabricacion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaVencimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaVencimiento); err != nil {
			return err
		}
	} else {
		r.FechaVencimiento = NewUnionNullLong()

		r.FechaVencimiento = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ProductoTrazable"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ProductoTrazable); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ProductoTrazable")
	}
	val = func() json.RawMessage {
		if v, ok := fields["AlmacenConsumo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AlmacenConsumo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for AlmacenConsumo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EstadoLote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoLote); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EstadoLote")
	}
	val = func() json.RawMessage {
		if v, ok := fields["BloqueoUbicacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BloqueoUbicacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for BloqueoUbicacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["VidaUtilLote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VidaUtilLote); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for VidaUtilLote")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EntregaAntesDe"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EntregaAntesDe); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EntregaAntesDe")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ConsumoAntesDe"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ConsumoAntesDe); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ConsumoAntesDe")
	}
	return nil
}
