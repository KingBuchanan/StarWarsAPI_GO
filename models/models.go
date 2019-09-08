package models

type Planets struct {
	Climate        string   `json:"climate"`
	Diameter       string   `json:"diameter"`
	Gravity        string   `json:"gravity"`
	Name           string   `json:"name"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Population     string   `json:"population"`
	Residents      []string `json:"residents"`
	RotationPeriod string   `json:"rotation_period"`
	SurfaceWater   string   `json:"surface_water"`
	Terrain        string   `json:"terrain"`
	URL            string   `json:"url"`
}