package parser

import (
	"github.com/aquasecurity/defsec/pkg/scanners/azure"
	"github.com/aquasecurity/defsec/pkg/scanners/azure/arm/parser/armjson"
	"github.com/aquasecurity/defsec/pkg/types"
)

type Template struct {
	Metadata       *types.Metadata        `json:"-"`
	Schema         azure.Value            `json:"$schema"`
	ContentVersion azure.Value            `json:"contentVersion"`
	APIProfile     azure.Value            `json:"apiProfile"`
	Parameters     map[string]Parameter   `json:"parameters"`
	Variables      map[string]azure.Value `json:"variables"`
	Functions      []Function             `json:"functions"`
	Resources      []Resource             `json:"resources"`
	Outputs        map[string]azure.Value `json:"outputs"`
}

type Parameter struct {
	*types.Metadata
	Type         azure.Value `json:"type"`
	DefaultValue azure.Value `json:"defaultValue"`
	MaxLength    azure.Value `json:"maxLength"`
	MinLength    azure.Value `json:"minLength"`
}

type Function struct{}

type Resource struct {
	*types.Metadata `json:"-"`
	innerResource
}

func (t *Template) SetMetadata(m *types.Metadata) {
	t.Metadata = m
}

func (r *Resource) SetMetadata(m *types.Metadata) {
	r.Metadata = m
}

func (p *Parameter) SetMetadata(m *types.Metadata) {
	p.Metadata = m
}

type innerResource struct {
	APIVersion azure.Value `json:"apiVersion"`
	Type       azure.Value `json:"type"`
	Kind       azure.Value `json:"kind"`
	Name       azure.Value `json:"name"`
	Location   azure.Value `json:"location"`
	Tags       azure.Value `json:"tags"`
	Sku        azure.Value `json:"sku"`
	Properties azure.Value `json:"properties"`
}

func (v *Resource) UnmarshalJSONWithMetadata(node armjson.Node) error {

	if err := node.Decode(&v.innerResource); err != nil {
		return err
	}

	v.Metadata = node.Metadata()

	for _, comment := range node.Comments() {
		var str string
		if err := comment.Decode(&str); err != nil {
			return err
		}
		// TODO
		//v.Metadata.Comments = append(v.Metadata.Comments, str)
	}

	return nil
}
