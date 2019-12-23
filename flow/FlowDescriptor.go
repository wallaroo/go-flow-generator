package flow

// FlowDescriptor describes Flows
type FlowDescriptor struct {
	Name     string             `yaml:"name"`
	Version  string             `yaml:"version"`
	Headers  []RecordDescriptor `yaml:"headers"`
	Trailers []RecordDescriptor `yaml:"trailers"`
	Content  RecordDescriptor   `yaml:"content"`
}
