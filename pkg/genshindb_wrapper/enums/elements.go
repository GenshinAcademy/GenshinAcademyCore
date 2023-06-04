package enums

type ElementText string //@name ElementText
type ElementType string //@name ElementType

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

const (
	ElementTypeUndefinedElement ElementType = "ELEMENT_NONE"
	ElementTypePyro             ElementType = "ELEMENT_PYRO"
	ElementTypeHydro            ElementType = "ELEMENT_HYDRO"
	ElementTypeGeo              ElementType = "ELEMENT_GEO"
	ElementTypeAnemo            ElementType = "ELEMENT_ANEMO"
	ElementTypeElectro          ElementType = "ELEMENT_ELECTRO"
	ElementTypeCryo             ElementType = "ELEMENT_CRYO"
	ElementTypeDendro           ElementType = "ELEMENT_DENDRO"
)
