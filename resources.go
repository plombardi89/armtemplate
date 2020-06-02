package armtemplate

type DependsOn []string
type Properties map[string]interface{}

type Resource struct {
	APIVersion string            `json:"apiVersion,omitempty"`
	Condition  string            `json:"condition,omitempty"`
	Comment    string            `json:"comments,omitempty"`
	Name       string            `json:"name"`
	Type       string            `json:"type"`
	Location   string            `json:"location"`
	DependsOn  DependsOn         `json:"dependsOn,omitempty"`
	Plan       *ResourcePlan     `json:"plan,omitempty"`
	Properties interface{}       `json:"properties"`
	SKU        *ResourceSKU      `json:"sku,omitempty"`
	Tags       map[string]string `json:"tags,omitempty"`
}

func NewResource(name, resourceType, apiVersion string) Resource {
	return Resource{Name: name, Type: resourceType, APIVersion: apiVersion}
}

type ResourceSKU struct {
	Capacity int    `json:"capacity,omitempty"`
	Family   string `json:"family,omitempty"`
	Name     string `json:"name"`
	Size     string `json:"size,omitempty"`
	Tier     string `json:"tier,omitempty"`
}

type ResourcePlan struct {
	Name          string `json:"name"`
	Product       string `json:"product,omitempty"`
	PromotionCode string `json:"promotionCode,omitempty"`
	Publisher     string `json:"publisher,omitempty"`
	Version       string `json:"version,omitempty"`
}

type Resources interface {
	Add(resources ...Resource)
	Len() int
}

func NewResources() Resources {
	resources := make(resources, 0)
	return &resources
}

type resources []Resource

func (r *resources) Add(resources ...Resource) {
	*r = append(*r, resources...)
}

func (r resources) Len() int {
	return len(r)
}
