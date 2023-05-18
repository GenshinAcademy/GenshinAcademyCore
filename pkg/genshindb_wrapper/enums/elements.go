package enums

type ElementText string

const (
	UndefinedElement ElementText = "None"
	Pyro             ElementText = "Pyro"
	Hydro            ElementText = "Hydro"
	Geo              ElementText = "Geo"
	Anemo            ElementText = "Anemo"
	Electro          ElementText = "Electro"
	Cryo             ElementText = "Cryo"
	Dendro           ElementText = "Dendro"
)

type ElementType string

const (
	ElementTypeUndefinedElement ElementText = "ELEMENT_NONE"
	ElementTypePyro             ElementText = "ELEMENT_PYRO"
	ElementTypeHydro            ElementText = "ELEMENT_HYDRO"
	ElementTypeGeo              ElementText = "ELEMENT_GEO"
	ElementTypeAnemo            ElementText = "ELEMENT_ANEMO"
	ElementTypeElectro          ElementText = "ELEMENT_ELECTRO"
	ElementTypeCryo             ElementText = "ELEMENT_CRYO"
	ElementTypeDendro           ElementText = "ELEMENT_DENDRO"
)
