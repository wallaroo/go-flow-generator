package flow

// RecordDescriptor describes records
type RecordDescriptor struct {
	Name   string            `yaml:"name"`
	Fields []FieldDescriptor `yaml:"fields"`
}
