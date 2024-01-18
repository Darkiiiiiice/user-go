/**
 * @Author: mariomang
 * @Date: 2024-01-19 00:18:12
 * @File: version/version.go
**/

package version

import "fmt"

var App string
var GoVersion string
var BuildVersion string
var BuildTime string
var BuildTag string
var CommitBranch string
var CommitID string

func Version() string {
	return fmt.Sprintf("App: %s\nGoVersion: %s\nBuildVersion: %s\nBudilTag: %s\nBuildTime: %s\nCommitBranch: %s\nCommitID: %s\n", App, GoVersion, BuildVersion, BuildTag, BuildTime, CommitBranch, CommitID)
}
