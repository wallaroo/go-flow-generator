package flow

// FieldDescriptor descriptor of field
type FieldDescriptor struct {
	Name   string   `yaml:"name"`
	Start  int      `yaml:"start"`
	End    int      `yaml:"end"`
	Type   string   `yaml:"type"`
	Format string   `yaml:"format,omitempty"`
	Const  string   `yaml:"const,omitempty"`
	Enum   []string `yaml:"enum,omitempty"`
	Fill   string   `yaml:"fill,omitempty"`
}

// Size gets the field size in chars
func (d *FieldDescriptor) Size() int {
	return d.End - d.Start
}
