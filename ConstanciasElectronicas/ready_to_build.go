// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ReadyToBuild.avsc
 */
package ConstanciasElectronicasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ReadyToBuild struct {
	BuilderNotOk bool `json:"builderNotOk"`

	StorageNotOk bool `json:"storageNotOk"`

	Ready bool `json:"ready"`

	CreatedAt int64 `json:"createdAt"`

	UpdatedAt int64 `json:"updatedAt"`

	Deleted bool `json:"deleted"`
}

const ReadyToBuildAvroCRC64Fingerprint = "\xdeB|knl:\x0e"

func NewReadyToBuild() ReadyToBuild {
	r := ReadyToBuild{}
	return r
}

func DeserializeReadyToBuild(r io.Reader) (ReadyToBuild, error) {
	t := NewReadyToBuild()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeReadyToBuildFromSchema(r io.Reader, schema string) (ReadyToBuild, error) {
	t := NewReadyToBuild()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeReadyToBuild(r ReadyToBuild, w io.Writer) error {
	var err error
	err = vm.WriteBool(r.BuilderNotOk, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.StorageNotOk, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Ready, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.CreatedAt, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.UpdatedAt, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Deleted, w)
	if err != nil {
		return err
	}
	return err
}

func (r ReadyToBuild) Serialize(w io.Writer) error {
	return writeReadyToBuild(r, w)
}

func (r ReadyToBuild) Schema() string {
	return "{\"fields\":[{\"name\":\"builderNotOk\",\"type\":\"boolean\"},{\"name\":\"storageNotOk\",\"type\":\"boolean\"},{\"name\":\"ready\",\"type\":\"boolean\"},{\"name\":\"createdAt\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"updatedAt\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"deleted\",\"type\":\"boolean\"}],\"name\":\"Andreani.ConstanciasElectronicas.Events.Record.ReadyToBuild\",\"type\":\"record\"}"
}

func (r ReadyToBuild) SchemaName() string {
	return "Andreani.ConstanciasElectronicas.Events.Record.ReadyToBuild"
}

func (_ ReadyToBuild) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ReadyToBuild) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ReadyToBuild) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ReadyToBuild) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ReadyToBuild) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ReadyToBuild) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ReadyToBuild) SetString(v string)   { panic("Unsupported operation") }
func (_ ReadyToBuild) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ReadyToBuild) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Boolean{Target: &r.BuilderNotOk}

		return w

	case 1:
		w := types.Boolean{Target: &r.StorageNotOk}

		return w

	case 2:
		w := types.Boolean{Target: &r.Ready}

		return w

	case 3:
		w := types.Long{Target: &r.CreatedAt}

		return w

	case 4:
		w := types.Long{Target: &r.UpdatedAt}

		return w

	case 5:
		w := types.Boolean{Target: &r.Deleted}

		return w

	}
	panic("Unknown field index")
}

func (r *ReadyToBuild) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ReadyToBuild) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ReadyToBuild) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ReadyToBuild) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ReadyToBuild) HintSize(int)                     { panic("Unsupported operation") }
func (_ ReadyToBuild) Finalize()                        {}

func (_ ReadyToBuild) AvroCRC64Fingerprint() []byte {
	return []byte(ReadyToBuildAvroCRC64Fingerprint)
}

func (r ReadyToBuild) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["builderNotOk"], err = json.Marshal(r.BuilderNotOk)
	if err != nil {
		return nil, err
	}
	output["storageNotOk"], err = json.Marshal(r.StorageNotOk)
	if err != nil {
		return nil, err
	}
	output["ready"], err = json.Marshal(r.Ready)
	if err != nil {
		return nil, err
	}
	output["createdAt"], err = json.Marshal(r.CreatedAt)
	if err != nil {
		return nil, err
	}
	output["updatedAt"], err = json.Marshal(r.UpdatedAt)
	if err != nil {
		return nil, err
	}
	output["deleted"], err = json.Marshal(r.Deleted)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ReadyToBuild) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["builderNotOk"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BuilderNotOk); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for builderNotOk")
	}
	val = func() json.RawMessage {
		if v, ok := fields["storageNotOk"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StorageNotOk); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for storageNotOk")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ready"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ready); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ready")
	}
	val = func() json.RawMessage {
		if v, ok := fields["createdAt"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CreatedAt); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for createdAt")
	}
	val = func() json.RawMessage {
		if v, ok := fields["updatedAt"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UpdatedAt); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for updatedAt")
	}
	val = func() json.RawMessage {
		if v, ok := fields["deleted"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Deleted); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for deleted")
	}
	return nil
}
