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
