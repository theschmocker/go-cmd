package cmd

type Flags struct {
	String  map[string]string
	Boolean map[string]bool
}

type Flag struct {
	Name         string
	ShortName    string
	DefaultValue string
	Description  string
	IsBoolean    bool
}
