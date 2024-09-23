package v1

type Queue struct {
	Class             string     `json:"_class" xml:"_class,attr"`
	DiscoverableItems []struct{} `json:"discoverableItems" xml:"discoverableItems"`
	Items             []struct {
		Class        string     `json:"_class" xml:"_class,attr"`
		Actions      []struct{} `json:"actions" xml:"actions"`
		Blocked      bool       `json:"blocked" xml:"blocked"`
		Buildable    bool       `json:"buildable"`
		ID           int        `json:"id" xml:"id"`
		InQueueSince int64      `json:"inQueueSince" xml:"inQueueSince"`
		Params       string     `json:"params" xml:"params"`
		Stuck        bool       `json:"stuck" xml:"stuck"`
		Task         struct {
			Class string `json:"_class" xml:"_class,attr"`
		} `json:"task" xml:"task"`
		Url                        string `json:"url" xml:"url"`
		Why                        string `json:"why" xml:"why"`
		BuildableStartMilliseconds int64  `json:"buildableStartMilliseconds" xml:"buildableStartMilliseconds"`
		Pending                    bool   `json:"pending" xml:"pending"`
	} `json:"items" xml:"items"`
}
