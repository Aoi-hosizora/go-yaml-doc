package goapidoc

// Model definitions
type Definition struct {
	Name        string
	Description string

	Generics   []string
	Properties []*Property
}

func NewDefinition(title string, desc string) *Definition {
	return &Definition{Name: title, Description: desc}
}

func (d *Definition) WithGenerics(generics ...string) *Definition {
	d.Generics = generics
	return d
}

func (d *Definition) WithProperties(properties ...*Property) *Definition {
	d.Properties = append(d.Properties, properties...)
	return d
}

// Model property
type Property struct {
	Name        string
	Type        string
	Required    bool
	Description string

	AllowEmptyValue bool
	Default         interface{}
	Enum            []interface{}
}

func NewProperty(name string, t string, req bool, desc string) *Property {
	return &Property{Name: name, Type: t, Required: req, Description: desc}
}

func (p *Property) WithAllowEmptyValue(allow bool) *Property {
	p.AllowEmptyValue = allow
	return p
}

func (p *Property) WithDefault(def interface{}) *Property {
	p.Default = def
	return p
}

func (p *Property) WithEnum(enum ...interface{}) *Property {
	p.Enum = enum
	return p
}
