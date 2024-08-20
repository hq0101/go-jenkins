package v1

type Build struct {
	Class  string `json:"_class" xml:"_class,attr"`
	Number int    `json:"number" xml:"number"`
	Url    string `json:"url" xml:"url"`
}
