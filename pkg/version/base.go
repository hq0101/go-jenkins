package version

var (
	gitMajor string = "" // major version, always numeric
	gitMinor string = "" // minor version, numeric possibly followed by "+"

	// NOTE: The $Format strings are replaced during 'git archive' thanks to the
	// companion .gitattributes file containing 'export-subst' in this same
	// directory.  See also https://git-scm.com/docs/gitattributes
	gitVersion   string = "v0.0.0-master+$Format:%H$"
	gitCommit    string = "$Format:%H$" // sha1 from git, output of $(git rev-parse HEAD)
	gitTreeState string = ""            // state of git tree, either "clean" or "dirty"

	buildDate string = "1970-01-01T00:00:00Z" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)
