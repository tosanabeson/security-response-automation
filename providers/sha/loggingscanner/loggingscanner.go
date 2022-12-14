package loggingscanner

import (
	"encoding/json"
	"strings"

	"github.com/googlecloudplatform/security-response-automation/cloudfunctions/iam/enableauditlogs"
	pb "github.com/googlecloudplatform/security-response-automation/compiled/sha/protos"
)

// Automation defines the configuration for this finding.
type Automation struct {
	Action     string
	Target     []string
	Exclude    []string
	Properties struct {
		DryRun bool `yaml:"dry_run"`
	}
}

// Finding represents this finding.
type Finding struct {
	loggingscanner *pb.LoggingScanner
}

// New returns a new finding.
func New(b []byte) (*Finding, error) {
	var f Finding
	if err := json.Unmarshal(b, &f.loggingscanner); err != nil {
		return nil, err
	}
	return &f, nil
}

// Name returns the category of the finding.
func (f *Finding) Name(b []byte) string {
	var finding pb.LoggingScanner
	if err := json.Unmarshal(b, &finding); err != nil {
		return ""
	}
	return strings.ToLower(finding.GetFinding().GetCategory())
}

// EnableAuditLogs return values for the enable audit logs automation.
func (f *Finding) EnableAuditLogs() *enableauditlogs.Values {
	return &enableauditlogs.Values{
		ProjectID: f.loggingscanner.GetFinding().GetSourceProperties().GetProjectID(),
	}
}
