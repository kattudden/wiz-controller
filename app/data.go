package main

// BulbStatus fasst den Status zusammen, der im Template dargestellt wird.
type BulbStatus struct {
	IP      string
	Port    string
	Name    string
	IsOn    bool
	Message string // Optional: Zusätzliche Infos, z. B. Fehleranzeige
}

// GroupStatus repräsentiert eine Gruppe von Bulbs für die Anzeige.
type GroupStatus struct {
	Name  string
	Bulbs []BulbStatus
	AllOn bool
}
