package resource

import (
	"encoding/json"
	"os"
	"time"

	m "github.com/digitalocean/concourse-resource-library/metadata"
)

// Source represents the configuration for the resource
type Source struct {
	// Repository to check, get, put
	Repository string `json:"repository"`
	// Branch to check, get, put
	Branch string `json:"branch,omitempty"`
	// AccessToken for GitHub API with permissions to Repository
	AccessToken string `json:"access_token"`
	// Endpoint for GitHub GraphQL API (leave blank for cloud)
	Endpoint string `json:"endpoint"`
	// Paths of Repository to return versions for
	Paths []string `json:"paths,omitempty"`
	// IgnorePaths of Repository to skip returning versions for
	IgnorePaths []string `json:"ignore_paths,omitempty"`
	// DisableCISkip disables ability to skip CI via PR title / message
	DisableCISkip bool `json:"disable_ci_skip,omitempty"`
	// SkipSSLVerification when executing GitHub API requests
	SkipSSLVerification bool `json:"skip_ssl_verification,omitempty"`
	// GitCryptKey enables GitCrypt unlocking
	GitCryptKey string `json:"git_crypt_key,omitempty"`
	// PreviewSchema enables GraphQL preview schemas, see: https://developer.github.com/v4/previews/
	PreviewSchema string `json:"preview_schema,omitempty"`
}

// Validate ensures that the source configuration is valid
func (s Source) Validate() error {
	return nil
}

// Version contains the version data Concourse uses to determine if a build should run
type Version struct {
	OID        string    `json: "oid"`
	PushedDate time.Time `json: "pushed"`
}

// CheckRequest is the data struct received from Concoruse by the resource check operation
type CheckRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

// Read will read the json response from Concourse via stdin
func (r *CheckRequest) Read(input []byte) error {
	return json.Unmarshal(input, r)
}

// CheckResponse is the data struct returned to Concourse by the resource check operation
type CheckResponse []Version

// Len returns the number of versions in the response
func (r CheckResponse) Len() int {
	return len(r)
}

// Write will write the json response to stdout for Concourse to parse
func (r CheckResponse) Write() error {
	return json.NewEncoder(os.Stdout).Encode(r)
}

// GetParameters is the configuration for a resource step
type GetParameters struct {
	// SkipDownload will skip downloading the code to the volume, used with `put` steps
	SkipDownload bool `json:"skip_download"`
	// IntegrationTool defines the method of checking out the code (checkout [default], merge, rebase)
	IntegrationTool string `json:"integration_tool"`
	// GitDepth sets the number of commits to include in the clone (shallow clone)
	GitDepth int `json:"git_depth"`
	// ListChangedFiles generates a list of changed files in the `.git` directory
	ListChangedFiles bool `json:"list_changed_files"`
}

// GetRequest is the data struct received from Concoruse by the resource get operation
type GetRequest struct {
	Source  Source        `json:"source"`
	Params  GetParameters `json:"params"`
	Version Version       `json:"version"`
}

// GetResponse ...
type GetResponse struct {
	Version  Version    `json:"version"`
	Metadata m.Metadata `json:"metadata,omitempty"`
}

// PutParameters for the resource
type PutParameters struct {
	Path        string `json:"path"`
	BaseContext string `json:"base_context"`
	Context     string `json:"context"`
	TargetURL   string `json:"target_url"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CommentFile string `json:"comment_file"`
	Comment     string `json:"comment"`
}

// PutRequest is the data struct received from Concoruse by the resource put operation
type PutRequest struct {
	Source Source        `json:"source"`
	Params PutParameters `json:"params"`
}
