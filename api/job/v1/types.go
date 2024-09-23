package v1

type Build struct {
	Class  string `json:"_class" xml:"_class,attr"`
	Number int    `json:"number" xml:"number"`
	Url    string `json:"url" xml:"url"`
}

const (
	ModeWorkflowMultiBranchProject = "org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject"
	ModeWorkflowJob                = "org.jenkinsci.plugins.workflow.job.WorkflowJob"
	ModeFreeStyleProject           = "hudson.model.FreeStyleProject"
	ModeMatrixProject              = "hudson.matrix.MatrixProject"
	ModeOrganizationFolder         = "jenkins.branch.OrganizationFolder"
	ModeFolder                     = "com.cloudbees.hudson.plugins.folder.Folder"
	ModeCopy                       = "copy"
)

type PipelineRun struct {
	Links               map[string]map[string]string `json:"_links" xml:"_links"`
	ID                  string                       `json:"id" xml:"id"`
	Name                string                       `json:"name" xml:"name"`
	ExecNode            string                       `json:"execNode" xml:"execNode"`
	Status              string                       `json:"status" xml:"status"`
	StartTimeMillis     int64                        `json:"startTimeMillis" xml:"startTimeMillis"`
	DurationMillis      int64                        `json:"durationMillis" xml:"durationMillis"`
	PauseDurationMillis int64                        `json:"pauseDurationMillis" xml:"pauseDurationMillis"`
	Stages              []struct {
		Links               map[string]map[string]string `json:"_links" xml:"_links"`
		ID                  string                       `json:"id" xml:"id"`
		Name                string                       `json:"name" xml:"name"`
		ExecNode            string                       `json:"execNode" xml:"execNode"`
		Status              string                       `json:"status" xml:"status"`
		StartTimeMillis     int64                        `json:"startTimeMillis" xml:"startTimeMillis"`
		DurationMillis      int64                        `json:"durationMillis" xml:"durationMillis"`
		PauseDurationMillis int64                        `json:"pauseDurationMillis" xml:"pauseDurationMillis"`
		ParentNodes         []int64                      `json:"parentNodes" xml:"parentNodes"`
	} `json:"stages"`
}

type ContextMenu struct {
	Class string `json:"_class" xml:"_class,attr"`
	Items []struct {
		Badge                struct{} `json:"badge" xml:"badge"`
		DisplayName          string   `json:"displayName" xml:"displayName"`
		Icon                 string   `json:"icon" xml:"icon"`
		IconXml              string   `json:"iconXml" xml:"iconXml"`
		Message              string   `json:"message" xml:"message"`
		Post                 bool     `json:"post" xml:"post"`
		RequiresConfirmation bool     `json:"requiresConfirmation" xml:"requiresConfirmation"`
		SubMenu              struct{} `json:"subMenu" xml:"subMenu"`
		Type                 string   `json:"type" xml:"type"`
		Url                  string   `json:"url" xml:"url"`
	} `json:"items" xml:"items"`
}
