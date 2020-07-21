// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     SysFlow.avsc
 */
package sfgo

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

type Container struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Image string `json:"image"`

	Imageid string `json:"imageid"`

	Type ContainerType `json:"type"`

	Privileged bool `json:"privileged"`

	Imagerepo string `json:"imagerepo"`
}

const ContainerAvroCRC64Fingerprint = "\xa21MEM\x1d?\x9b"

func NewContainer() *Container {
	return &Container{}
}

func DeserializeContainer(r io.Reader) (*Container, error) {
	t := NewContainer()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeContainerFromSchema(r io.Reader, schema string) (*Container, error) {
	t := NewContainer()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeContainer(r *Container, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Image, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Imageid, w)
	if err != nil {
		return err
	}
	err = writeContainerType(r.Type, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Privileged, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Imagerepo, w)
	if err != nil {
		return err
	}
	return err
}

func (r *Container) Serialize(w io.Writer) error {
	return writeContainer(r, w)
}

func (r *Container) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"image\",\"type\":\"string\"},{\"name\":\"imageid\",\"type\":\"string\"},{\"name\":\"type\",\"type\":{\"name\":\"ContainerType\",\"namespace\":\"sysflow.type\",\"symbols\":[\"CT_DOCKER\",\"CT_LXC\",\"CT_LIBVIRT_LXC\",\"CT_MESOS\",\"CT_RKT\",\"CT_CUSTOM\",\"CT_CRI\",\"CT_CONTAINERD\",\"CT_CRIO\",\"CT_BPM\"],\"type\":\"enum\"}},{\"name\":\"privileged\",\"type\":\"boolean\"},{\"default\":\"NA\",\"name\":\"imagerepo\",\"type\":\"string\"}],\"name\":\"sysflow.entity.Container\",\"type\":\"record\"}"
}

func (r *Container) SchemaName() string {
	return "sysflow.entity.Container"
}

func (_ *Container) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *Container) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *Container) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *Container) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *Container) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *Container) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *Container) SetString(v string)   { panic("Unsupported operation") }
func (_ *Container) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Container) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.Id}
	case 1:
		return &types.String{Target: &r.Name}
	case 2:
		return &types.String{Target: &r.Image}
	case 3:
		return &types.String{Target: &r.Imageid}
	case 4:
		return &ContainerTypeWrapper{Target: &r.Type}
	case 5:
		return &types.Boolean{Target: &r.Privileged}
	case 6:
		return &types.String{Target: &r.Imagerepo}
	}
	panic("Unknown field index")
}

func (r *Container) SetDefault(i int) {
	switch i {
	case 6:
		r.Imagerepo = "NA"
		return
	}
	panic("Unknown field index")
}

func (r *Container) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *Container) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Container) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *Container) Finalize()                        {}

func (_ *Container) AvroCRC64Fingerprint() []byte {
	return []byte(ContainerAvroCRC64Fingerprint)
}
