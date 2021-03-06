/*
 * CircleCI API
 *
 * This describes the resources that make up the CircleCI API v2. API v2 is currently in Preview. Additional documentation for this API can be found in the [API Preview Docs](https://github.com/CircleCI-Public/api-preview-docs/tree/master/docs). Breaking changes to the API will be announced in the [Breaking Changes log](https://github.com/CircleCI-Public/api-preview-docs/blob/master/docs/breaking.md).
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package circleci

// The information you can supply when rerunning a workflow.
type RerunWorkflowParameters struct {
	// A list of job IDs to rerun.
	Jobs []string `json:"jobs,omitempty"`
	// Whether to rerun the workflow from the failed job. Mutually exclusive with the jobs parameter.
	FromFailed bool `json:"from_failed,omitempty"`
}
