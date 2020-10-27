package codegen

// endpointRegistry is a registry of all the endpoints we'd
// like to have generated, along with the type of template
// we should use.
// IMPORTANT NOTE: To support high quality, only add one
// endpoint per PR.
var endpointRegistry = map[string]*additionalInfo{
	"/ad/roles/{name}": {
		Type:                 tfTypeResource,
		AdditionalParameters: []templatableParam{},
	},
	"/ad/rotate-root": {
		Type:                 tfTypeResource,
		AdditionalParameters: []templatableParam{},
	},
}

// tfType is the type of Terraform code to generate.
type tfType int

const (
	tfTypeUnset tfType = iota
	tfTypeDataSource
	tfTypeResource
)

func (t tfType) String() string {
	switch t {
	case tfTypeDataSource:
		return "datasource"
	case tfTypeResource:
		return "resource"
	}
	return "unset"
}

type additionalInfo struct {
	Type                 tfType
	AdditionalParameters []templatableParam
}
