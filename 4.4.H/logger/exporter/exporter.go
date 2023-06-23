package exporter

type Exporter interface {
	Export(message string)
}
