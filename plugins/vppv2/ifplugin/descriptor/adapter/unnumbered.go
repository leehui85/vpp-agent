// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"github.com/gogo/protobuf/proto"
	. "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	"github.com/ligato/vpp-agent/api/models/vpp/interfaces"
)

////////// type-safe key-value pair with metadata //////////

type UnnumberedKVWithMetadata struct {
	Key      string
	Value    *vpp_interfaces.Interface_Unnumbered
	Metadata interface{}
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type UnnumberedDescriptor struct {
	Name               string
	KeySelector        KeySelector
	ValueTypeName      string
	KeyLabel           func(key string) string
	ValueComparator    func(key string, oldValue, newValue *vpp_interfaces.Interface_Unnumbered) bool
	NBKeyPrefix        string
	WithMetadata       bool
	MetadataMapFactory MetadataMapFactory
	Validate           func(key string, value *vpp_interfaces.Interface_Unnumbered) error
	Add                func(key string, value *vpp_interfaces.Interface_Unnumbered) (metadata interface{}, err error)
	Delete             func(key string, value *vpp_interfaces.Interface_Unnumbered, metadata interface{}) error
	Modify             func(key string, oldValue, newValue *vpp_interfaces.Interface_Unnumbered, oldMetadata interface{}) (newMetadata interface{}, err error)
	ModifyWithRecreate func(key string, oldValue, newValue *vpp_interfaces.Interface_Unnumbered, metadata interface{}) bool
	IsRetriableFailure func(err error) bool
	Dependencies       func(key string, value *vpp_interfaces.Interface_Unnumbered) []Dependency
	DerivedValues      func(key string, value *vpp_interfaces.Interface_Unnumbered) []KeyValuePair
	Dump               func(correlate []UnnumberedKVWithMetadata) ([]UnnumberedKVWithMetadata, error)
	DumpDependencies   []string /* descriptor name */
}

////////// Descriptor adapter //////////

type UnnumberedDescriptorAdapter struct {
	descriptor *UnnumberedDescriptor
}

func NewUnnumberedDescriptor(typedDescriptor *UnnumberedDescriptor) *KVDescriptor {
	adapter := &UnnumberedDescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:               typedDescriptor.Name,
		KeySelector:        typedDescriptor.KeySelector,
		ValueTypeName:      typedDescriptor.ValueTypeName,
		KeyLabel:           typedDescriptor.KeyLabel,
		NBKeyPrefix:        typedDescriptor.NBKeyPrefix,
		WithMetadata:       typedDescriptor.WithMetadata,
		MetadataMapFactory: typedDescriptor.MetadataMapFactory,
		IsRetriableFailure: typedDescriptor.IsRetriableFailure,
		DumpDependencies:   typedDescriptor.DumpDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Validate != nil {
		descriptor.Validate = adapter.Validate
	}
	if typedDescriptor.Add != nil {
		descriptor.Add = adapter.Add
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Modify != nil {
		descriptor.Modify = adapter.Modify
	}
	if typedDescriptor.ModifyWithRecreate != nil {
		descriptor.ModifyWithRecreate = adapter.ModifyWithRecreate
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	if typedDescriptor.Dump != nil {
		descriptor.Dump = adapter.Dump
	}
	return descriptor
}

func (da *UnnumberedDescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castUnnumberedValue(key, oldValue)
	typedNewValue, err2 := castUnnumberedValue(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *UnnumberedDescriptorAdapter) Validate(key string, value proto.Message) (err error) {
	typedValue, err := castUnnumberedValue(key, value)
	if err != nil {
		return err
	}
	return da.descriptor.Validate(key, typedValue)
}

func (da *UnnumberedDescriptorAdapter) Add(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castUnnumberedValue(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Add(key, typedValue)
}

func (da *UnnumberedDescriptorAdapter) Modify(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castUnnumberedValue(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castUnnumberedValue(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castUnnumberedMetadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Modify(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *UnnumberedDescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castUnnumberedValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castUnnumberedMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *UnnumberedDescriptorAdapter) ModifyWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castUnnumberedValue(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castUnnumberedValue(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castUnnumberedMetadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.ModifyWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *UnnumberedDescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castUnnumberedValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

func (da *UnnumberedDescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castUnnumberedValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *UnnumberedDescriptorAdapter) Dump(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []UnnumberedKVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castUnnumberedValue(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castUnnumberedMetadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			UnnumberedKVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedDump, err := da.descriptor.Dump(correlateWithType)
	if err != nil {
		return nil, err
	}
	var dump []KVWithMetadata
	for _, typedKVWithMetadata := range typedDump {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		dump = append(dump, kvWithMetadata)
	}
	return dump, err
}

////////// Helper methods //////////

func castUnnumberedValue(key string, value proto.Message) (*vpp_interfaces.Interface_Unnumbered, error) {
	typedValue, ok := value.(*vpp_interfaces.Interface_Unnumbered)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castUnnumberedMetadata(key string, metadata Metadata) (interface{}, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(interface{})
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}
