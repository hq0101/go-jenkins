package v1

type State string

const (
	StateSuccess State = "success"
	StateFailure State = "failure"
)

func (s State) String() string {
	return string(s)
}

type Result struct {
	Status State `json:"status"`
	Data   struct {
		Result string `json:"result"`
		Errors []struct {
			Error string `json:"error"`
		} `json:"errors"`
		Json        interface{} `json:"json"`
		Jenkinsfile string      `json:"jenkinsfile"`
	} `json:"data"`
}
