// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Ticket.avsc
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

type Ticket struct {
	IdTicket int32 `json:"IdTicket"`

	Grupo string `json:"Grupo"`

	Categoria string `json:"Categoria"`

	Articulo string `json:"Articulo"`

	Solicitante string `json:"Solicitante"`

	Asunto string `json:"Asunto"`

	Descripcion string `json:"Descripcion"`

	Tecnico string `json:"Tecnico"`

	Cliente *UnionNullString `json:"Cliente"`

	Planta *UnionNullString `json:"Planta"`

	Nave *UnionNullString `json:"Nave"`

	Sector *UnionNullString `json:"Sector"`

	Sectores *UnionNullString `json:"Sectores"`

	Proceso *UnionNullString `json:"Proceso"`

	Procesos *UnionNullString `json:"Procesos"`

	Etiqueta *UnionNullString `json:"Etiqueta"`

	Sucursal *UnionNullString `json:"Sucursal"`

	Sucursales *UnionNullString `json:"Sucursales"`

	FechaCreacion int64 `json:"FechaCreacion"`

	FechaVencimiento int64 `json:"FechaVencimiento"`

	FechaResolucion int64 `json:"FechaResolucion"`

	FechaFinalizado int64 `json:"FechaFinalizado"`

	Estado string `json:"Estado"`
}

const TicketAvroCRC64Fingerprint = "\x15ss\x8c\xef\xfc\xe0z"

func NewTicket() Ticket {
	r := Ticket{}
	return r
}

func DeserializeTicket(r io.Reader) (Ticket, error) {
	t := NewTicket()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTicketFromSchema(r io.Reader, schema string) (Ticket, error) {
	t := NewTicket()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTicket(r Ticket, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.IdTicket, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Grupo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Categoria, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Articulo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Solicitante, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Asunto, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Descripcion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Tecnico, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Cliente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Planta, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Nave, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Sector, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Sectores, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Proceso, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Procesos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Etiqueta, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Sucursal, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Sucursales, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.FechaCreacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.FechaVencimiento, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.FechaResolucion, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.FechaFinalizado, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Estado, w)
	if err != nil {
		return err
	}
	return err
}

func (r Ticket) Serialize(w io.Writer) error {
	return writeTicket(r, w)
}

func (r Ticket) Schema() string {
	return "{\"fields\":[{\"name\":\"IdTicket\",\"type\":\"int\"},{\"name\":\"Grupo\",\"type\":\"string\"},{\"name\":\"Categoria\",\"type\":\"string\"},{\"name\":\"Articulo\",\"type\":\"string\"},{\"name\":\"Solicitante\",\"type\":\"string\"},{\"name\":\"Asunto\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"name\":\"Tecnico\",\"type\":\"string\"},{\"name\":\"Cliente\",\"type\":[\"null\",\"string\"]},{\"name\":\"Planta\",\"type\":[\"null\",\"string\"]},{\"name\":\"Nave\",\"type\":[\"null\",\"string\"]},{\"name\":\"Sector\",\"type\":[\"null\",\"string\"]},{\"name\":\"Sectores\",\"type\":[\"null\",\"string\"]},{\"name\":\"Proceso\",\"type\":[\"null\",\"string\"]},{\"name\":\"Procesos\",\"type\":[\"null\",\"string\"]},{\"name\":\"Etiqueta\",\"type\":[\"null\",\"string\"]},{\"name\":\"Sucursal\",\"type\":[\"null\",\"string\"]},{\"name\":\"Sucursales\",\"type\":[\"null\",\"string\"]},{\"name\":\"FechaCreacion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"FechaVencimiento\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"FechaResolucion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"FechaFinalizado\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"Estado\",\"type\":\"string\"}],\"name\":\"Andreani.EAM.Events.MDA.Ticket\",\"type\":\"record\"}"
}

func (r Ticket) SchemaName() string {
	return "Andreani.EAM.Events.MDA.Ticket"
}

func (_ Ticket) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Ticket) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Ticket) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Ticket) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Ticket) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Ticket) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Ticket) SetString(v string)   { panic("Unsupported operation") }
func (_ Ticket) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Ticket) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.IdTicket}

		return w

	case 1:
		w := types.String{Target: &r.Grupo}

		return w

	case 2:
		w := types.String{Target: &r.Categoria}

		return w

	case 3:
		w := types.String{Target: &r.Articulo}

		return w

	case 4:
		w := types.String{Target: &r.Solicitante}

		return w

	case 5:
		w := types.String{Target: &r.Asunto}

		return w

	case 6:
		w := types.String{Target: &r.Descripcion}

		return w

	case 7:
		w := types.String{Target: &r.Tecnico}

		return w

	case 8:
		r.Cliente = NewUnionNullString()

		return r.Cliente
	case 9:
		r.Planta = NewUnionNullString()

		return r.Planta
	case 10:
		r.Nave = NewUnionNullString()

		return r.Nave
	case 11:
		r.Sector = NewUnionNullString()

		return r.Sector
	case 12:
		r.Sectores = NewUnionNullString()

		return r.Sectores
	case 13:
		r.Proceso = NewUnionNullString()

		return r.Proceso
	case 14:
		r.Procesos = NewUnionNullString()

		return r.Procesos
	case 15:
		r.Etiqueta = NewUnionNullString()

		return r.Etiqueta
	case 16:
		r.Sucursal = NewUnionNullString()

		return r.Sucursal
	case 17:
		r.Sucursales = NewUnionNullString()

		return r.Sucursales
	case 18:
		w := types.Long{Target: &r.FechaCreacion}

		return w

	case 19:
		w := types.Long{Target: &r.FechaVencimiento}

		return w

	case 20:
		w := types.Long{Target: &r.FechaResolucion}

		return w

	case 21:
		w := types.Long{Target: &r.FechaFinalizado}

		return w

	case 22:
		w := types.String{Target: &r.Estado}

		return w

	}
	panic("Unknown field index")
}

func (r *Ticket) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Ticket) NullField(i int) {
	switch i {
	case 8:
		r.Cliente = nil
		return
	case 9:
		r.Planta = nil
		return
	case 10:
		r.Nave = nil
		return
	case 11:
		r.Sector = nil
		return
	case 12:
		r.Sectores = nil
		return
	case 13:
		r.Proceso = nil
		return
	case 14:
		r.Procesos = nil
		return
	case 15:
		r.Etiqueta = nil
		return
	case 16:
		r.Sucursal = nil
		return
	case 17:
		r.Sucursales = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Ticket) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Ticket) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Ticket) HintSize(int)                     { panic("Unsupported operation") }
func (_ Ticket) Finalize()                        {}

func (_ Ticket) AvroCRC64Fingerprint() []byte {
	return []byte(TicketAvroCRC64Fingerprint)
}

func (r Ticket) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["IdTicket"], err = json.Marshal(r.IdTicket)
	if err != nil {
		return nil, err
	}
	output["Grupo"], err = json.Marshal(r.Grupo)
	if err != nil {
		return nil, err
	}
	output["Categoria"], err = json.Marshal(r.Categoria)
	if err != nil {
		return nil, err
	}
	output["Articulo"], err = json.Marshal(r.Articulo)
	if err != nil {
		return nil, err
	}
	output["Solicitante"], err = json.Marshal(r.Solicitante)
	if err != nil {
		return nil, err
	}
	output["Asunto"], err = json.Marshal(r.Asunto)
	if err != nil {
		return nil, err
	}
	output["Descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	output["Tecnico"], err = json.Marshal(r.Tecnico)
	if err != nil {
		return nil, err
	}
	output["Cliente"], err = json.Marshal(r.Cliente)
	if err != nil {
		return nil, err
	}
	output["Planta"], err = json.Marshal(r.Planta)
	if err != nil {
		return nil, err
	}
	output["Nave"], err = json.Marshal(r.Nave)
	if err != nil {
		return nil, err
	}
	output["Sector"], err = json.Marshal(r.Sector)
	if err != nil {
		return nil, err
	}
	output["Sectores"], err = json.Marshal(r.Sectores)
	if err != nil {
		return nil, err
	}
	output["Proceso"], err = json.Marshal(r.Proceso)
	if err != nil {
		return nil, err
	}
	output["Procesos"], err = json.Marshal(r.Procesos)
	if err != nil {
		return nil, err
	}
	output["Etiqueta"], err = json.Marshal(r.Etiqueta)
	if err != nil {
		return nil, err
	}
	output["Sucursal"], err = json.Marshal(r.Sucursal)
	if err != nil {
		return nil, err
	}
	output["Sucursales"], err = json.Marshal(r.Sucursales)
	if err != nil {
		return nil, err
	}
	output["FechaCreacion"], err = json.Marshal(r.FechaCreacion)
	if err != nil {
		return nil, err
	}
	output["FechaVencimiento"], err = json.Marshal(r.FechaVencimiento)
	if err != nil {
		return nil, err
	}
	output["FechaResolucion"], err = json.Marshal(r.FechaResolucion)
	if err != nil {
		return nil, err
	}
	output["FechaFinalizado"], err = json.Marshal(r.FechaFinalizado)
	if err != nil {
		return nil, err
	}
	output["Estado"], err = json.Marshal(r.Estado)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Ticket) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["IdTicket"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdTicket); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IdTicket")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Grupo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Grupo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Grupo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Categoria"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Categoria); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Categoria")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Articulo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Articulo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Articulo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Solicitante"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Solicitante); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Solicitante")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Asunto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Asunto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Asunto")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Descripcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Descripcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Descripcion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Tecnico"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tecnico); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Tecnico")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Cliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Cliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Planta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Planta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Planta")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Nave"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Nave); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Nave")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Sector"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sector); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Sector")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Sectores"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sectores); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Sectores")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Proceso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Proceso); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Proceso")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Procesos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Procesos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Procesos")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Etiqueta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Etiqueta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Etiqueta")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Sucursal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sucursal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Sucursal")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Sucursales"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sucursales); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Sucursales")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaCreacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaCreacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaCreacion")
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
		return fmt.Errorf("no value specified for FechaVencimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaResolucion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaResolucion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaResolucion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaFinalizado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaFinalizado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaFinalizado")
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
	return nil
}
