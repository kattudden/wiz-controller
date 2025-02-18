package config

// Config bildet die Struktur der YAML-Konfiguration ab.
type Config struct {
	Defaults Defaults            `yaml:"defaults"`
	Groups   map[string]GroupDef `yaml:"groups"`
}

// Defaults enthält Standardwerte.
type Defaults struct {
	Temperatur int `yaml:"temperatur"`
	Dimming    int `yaml:"dimming"`
}

// GroupDef enthält die Details einer Gruppe, in diesem Fall eine Liste von Bulps.
type GroupDef struct {
	Bulps []Bulp `yaml:"bulps"`
}

// Bulp repräsentiert ein einzelnes Gerät.
type Bulp struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}
