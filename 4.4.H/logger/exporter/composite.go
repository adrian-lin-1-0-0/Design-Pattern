package exporter

type Composite struct {
	exporters []Exporter
}

func NewComposite(exporters ...Exporter) *Composite {
	c := &Composite{exporters}
	return c
}

func (c *Composite) Export(message string) {
	for _, exporter := range c.exporters {
		exporter.Export(message)
	}
}
