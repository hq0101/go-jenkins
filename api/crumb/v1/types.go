package v1

type CrumbIssuer struct {
	Class             string `json:"_class" xml:"_class,attr"`
	Crumb             string `json:"crumb" xml:"crumb"`
	CrumbRequestField string `json:"crumbRequestField" xml:"crumbRequestField"`
}
