// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCES:
 *     DefaultEventCreated.avsc
 *     EventHeader.avsc
 */
package order_orderrepository2

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v9/compiler"
	"github.com/actgardner/gogen-avro/v9/vm"
	"github.com/actgardner/gogen-avro/v9/vm/types"
)

var _ = fmt.Printf

type DefaultEventCreated struct {
	// Example: This is a sample attribute, please edit.
	Sample1 string `json:"sample1"`

	NewField1 *UnionNullString `json:"newField1"`

	NewField2 *UnionNullString `json:"newField2"`
}

const DefaultEventCreatedAvroCRC64Fingerprint = "\xa0\xfa\xa5\xf7\x80\xe2\xb0\xe7"

func NewDefaultEventCreated() DefaultEventCreated {
	r := DefaultEventCreated{}
	r.NewField1 = NewUnionNullString()

	r.NewField1 = nil
	r.NewField2 = NewUnionNullString()

	r.NewField2 = nil
	return r
}

func DeserializeDefaultEventCreated(r io.Reader) (DefaultEventCreated, error) {
	t := NewDefaultEventCreated()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDefaultEventCreatedFromSchema(r io.Reader, schema string) (DefaultEventCreated, error) {
	t := NewDefaultEventCreated()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDefaultEventCreated(r DefaultEventCreated, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Sample1, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NewField1, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NewField2, w)
	if err != nil {
		return err
	}
	return err
}

func (r DefaultEventCreated) Serialize(w io.Writer) error {
	return writeDefaultEventCreated(r, w)
}

func (r DefaultEventCreated) Schema() string {
	return "{\"fields\":[{\"doc\":\"Example: This is a sample attribute, please edit.\",\"name\":\"sample1\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"default\":null,\"doc\":\"\",\"name\":\"newField1\",\"type\":[\"null\",{\"avro.java.string\":\"String\",\"type\":\"string\"}]},{\"aliases\":[\"fred\"],\"default\":null,\"doc\":\"\",\"name\":\"newField2\",\"type\":[\"null\",{\"avro.java.string\":\"String\",\"type\":\"string\"}]}],\"name\":\"com.kroger.desp.events.kcp.order.orderreposixx.xxx.DefaultEventCreated\",\"type\":\"record\"}"
}

func (r DefaultEventCreated) SchemaName() string {
	return "com.kroger.desp.events.kcp.order.orderreposixx.xxx.DefaultEventCreated"
}

func (_ DefaultEventCreated) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DefaultEventCreated) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DefaultEventCreated) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DefaultEventCreated) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DefaultEventCreated) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DefaultEventCreated) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DefaultEventCreated) SetString(v string)   { panic("Unsupported operation") }
func (_ DefaultEventCreated) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DefaultEventCreated) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.Sample1}
	case 1:
		r.NewField1 = NewUnionNullString()

		return r.NewField1
	case 2:
		r.NewField2 = NewUnionNullString()

		return r.NewField2
	}
	panic("Unknown field index")
}

func (r *DefaultEventCreated) SetDefault(i int) {
	switch i {
	case 1:
		r.NewField1 = nil
		return
	case 2:
		r.NewField2 = nil
		return
	}
	panic("Unknown field index")
}

func (r *DefaultEventCreated) NullField(i int) {
	switch i {
	case 1:
		r.NewField1 = nil
		return
	case 2:
		r.NewField2 = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ DefaultEventCreated) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DefaultEventCreated) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DefaultEventCreated) Finalize()                        {}

func (_ DefaultEventCreated) AvroCRC64Fingerprint() []byte {
	return []byte(DefaultEventCreatedAvroCRC64Fingerprint)
}

func (r DefaultEventCreated) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["sample1"], err = json.Marshal(r.Sample1)
	if err != nil {
		return nil, err
	}
	output["newField1"], err = json.Marshal(r.NewField1)
	if err != nil {
		return nil, err
	}
	output["newField2"], err = json.Marshal(r.NewField2)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DefaultEventCreated) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["sample1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sample1); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for sample1")
	}
	val = func() json.RawMessage {
		if v, ok := fields["newField1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NewField1); err != nil {
			return err
		}
	} else {
		r.NewField1 = NewUnionNullString()

		r.NewField1 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["newField2"]; ok {
			return v
		}
		if v, ok := fields["fred"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NewField2); err != nil {
			return err
		}
	} else {
		r.NewField2 = NewUnionNullString()

		r.NewField2 = nil
	}
	return nil
}
