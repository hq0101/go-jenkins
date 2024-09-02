package v1

const (
	ModeMyView string = "hudson.model.MyView"
)

type Categories struct {
	Class      string `json:"_class" xml:"_class,attr"`
	Categories []struct {
		ID          string `json:"id" xml:"id"`
		Name        string `json:"name" xml:"name"`
		Description string `json:"description" xml:"description"`
		Order       int    `json:"order" xml:"order"`
		MinToShow   int    `json:"minToShow" xml:"minToShow"`
		Items       []struct {
			Class               string `json:"_class" xml:"_class,attr"`
			DisplayName         string `json:"displayName" xml:"displayName"`
			Description         string `json:"description" xml:"description"`
			Order               int    `json:"order" xml:"order"`
			IconFilePathPattern string `json:"iconFilePathPattern" xml:"iconFilePathPattern"`
			IconQualifiedUrl    string `json:"iconQualifiedUrl" xml:"iconQualifiedUrl"`
			IconClassName       string `json:"iconClassName" xml:"iconClassName"`
		}
	} `json:"categories" xml:"categories"`
}
