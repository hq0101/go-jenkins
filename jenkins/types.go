package jenkins

import "encoding/xml"

type CrumbIssuer struct {
	Class             string `json:"_class"`
	Crumb             string `json:"crumb"`
	CrumbRequestField string `json:"crumbRequestField"`
}

type Categories struct {
	Class      string `json:"_class"`
	Categories []struct {
		Description string `json:"description"`
		ID          string `json:"id"`
		Items       []struct {
			DisplayName         string `json:"displayName"`
			IconFilePathPattern string `json:"iconFilePathPattern"`
			IconQualifiedUrl    string `json:"iconQualifiedUrl"`
			Description         string `json:"description"`
			IconClassName       string `json:"iconClassName"`
			Class               string `json:"class"`
			Order               int    `json:"order"`
		}
		MinToShow int    `json:"minToShow"`
		Name      string `json:"name"`
		Order     int    `json:"order"`
	} `json:"categories"`
}

type AutoCompleteCopyNewItemFromResponse struct {
	Class       string `json:"_class" xml:"_class,attr"`
	Suggestions []struct {
		Name string `json:"name" xml:"name"`
	} `json:"suggestions" xml:"suggestions"`
}

type Build struct {
	Class  string `json:"_class" xml:"_class,attr"`
	Number int    `json:"number" xml:"number"`
	Url    string `json:"url" xml:"url"`
}

type JobInfo struct {
	XMLName         xml.Name            `json:"-" xml:"hudson"`
	Class           string              `json:"_class" xml:"_class,attr"`
	Actions         []map[string]string `json:"actions" xml:"actions"`
	Description     string              `json:"description" xml:"description"`
	DisplayName     string              `json:"displayName" xml:"displayName"`
	FullDisplayName string              `json:"fullDisplayName" xml:"fullDisplayName"`
	FullName        string              `json:"fullName" xml:"fullName"`
	Name            string              `json:"name" xml:"name"`
	Url             string              `json:"url" xml:"url"`
	Buildable       bool                `json:"buildable" xml:"buildable"`
	Builds          []Build             `json:"builds" xml:"build"`
	Color           string              `json:"color" xml:"color"`
	FirstBuild      Build               `json:"firstBuild" xml:"firstBuild"`
	HealthReport    []struct {
		Description   string `json:"description" xml:"description"`
		IconClassName string `json:"iconClassName" xml:"iconClassName"`
		IconUrl       string `json:"iconUrl" xml:"iconUrl"`
		Score         int    `json:"score" xml:"score"`
	} `json:"healthReport,omitempty" xml:"healthReport,omitempty"`
	InQueue               bool  `json:"inQueue" xml:"inQueue"`
	KeepDependencies      bool  `json:"keepDependencies,omitempty" xml:"keepDependencies,omitempty"`
	LastBuild             Build `json:"lastBuild,omitempty" xml:"lastBuild,omitempty"`
	LastCompleteBuild     Build `json:"lastCompleteBuild,omitempty" xml:"lastCompleteBuild,omitempty"`
	LastFailedBuild       Build `json:"lastFailedBuild,omitempty" xml:"lastFailedBuild,omitempty"`
	LastStableBuild       Build `json:"lastStableBuild,omitempty" xml:"lastStableBuild,omitempty"`
	LastSuccessfulBuild   Build `json:"lastSuccessfulBuild,omitempty" xml:"lastSuccessfulBuild,omitempty"`
	LastUnstableBuild     Build `json:"lastUnstableBuild,omitempty" xml:"lastUnstableBuild,omitempty"`
	LastUnsuccessfulBuild Build `json:"lastUnsuccessfulBuild,omitempty" xml:"lastUnsuccessfulBuild,omitempty"`
	NextBuildNumber       int   `json:"nextBuildNumber" xml:"nextBuildNumber"`
	Property              struct {
	} `json:"property,omitempty" xml:"property,omitempty"`
	QueueItem struct {
	} `json:"queueItem,omitempty" xml:"queueItem,omitempty"`
	ConCurrentBuild bool       `json:"concurrentBuild" xml:"concurrentBuild"`
	ReSumeBlocked   bool       `json:"resumeBlocked" xml:"resumeBlocked"`
	Jobs            []struct{} `json:"jobs,omitempty" xml:"job,omitempty"`
	PrimaryView     struct {
		Class string `json:"_class" xml:"_class,attr"`
		Name  string `json:"name" xml:"name"`
		Url   string `json:"url" xml:"url"`
	} `json:"primaryView,omitempty" xml:"primaryView,omitempty"`
	View []struct {
		Class string `json:"_class" xml:"_class,attr"`
		Name  string `json:"name" xml:"name"`
		Url   string `json:"url" xml:"url"`
	} `json:"view,omitempty" xml:"view,omitempty"`
	Disabled           bool     `json:"disabled,omitempty" xml:"disabled,omitempty"`
	DownStreamProjects struct{} `json:"downstreamProjects,omitempty" xml:"downstreamProjects,omitempty"`
	LabelExpression    string   `json:"labelExpression,omitempty" xml:"labelExpression,omitempty"`
	Scm                struct {
		Class string `json:"_class" xml:"_class,attr"`
	} `json:"scm,omitempty" xml:"scm,omitempty"`
	UpStreamProjects struct{} `json:"upstreamProjects,omitempty" xml:"upstreamProjects,omitempty"`
}

// ViewConfig 用于创建或复制 View 时发送的配置
type ViewConfig struct {
	XMLName     xml.Name `json:"-" xml:"hudson"`
	Class       string   `json:"_class" xml:"_class,attr"`
	Description string   `json:"description" xml:"description"`
	Name        string   `json:"name" xml:"name"`
	Url         string   `json:"url" xml:"url"`
	Jobs        []Job    `json:"jobs"`
}

type Job struct {
	Class string `json:"_class" xml:"_class,attr"`
	Name  string `json:"name"   xml:"name"`
	URL   string `json:"url"    xml:"url"`
	Color string `json:"color,omitempty" xml:"color,omitempty"`
}
