// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package howto

import (
	"context"
	"fmt"
	"io"
	"time"

	talent "cloud.google.com/go/talent/apiv4beta1"
	talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"
)

// [START job_search_create_job_custom_attributes]

// createJobWithCustomAttributes creates a job with custom attributes.
func createJobWithCustomAttributes(w io.Writer, projectId string, companyName string, jobTitle string) (*talentpb.Job, error) {
	// requisitionId shoud be the unique ID in your system
	requisitionId := fmt.Sprintf("job-with-custom-attribute-%d", time.Now().UnixNano())

	jobToCreate := &talentpb.Job{
		RequisitionId: requisitionId,
		Title:         jobTitle,
		CompanyName:   companyName,
		ApplicationInfo: &talentpb.Job_ApplicationInfo{
			Uris: []string{"https://googlesample.com/career"},
		},
		Description: "Design, devolop, test, deploy, maintain and improve software.",
		CustomAttributes: map[string]*talentpb.CustomAttribute{
			"someFieldString": {
				Filterable:   true,
				StringValues: []string{"someStrVal"},
			},
			"someFieldLong": {
				Filterable: true,
				LongValues: []int64{900},
			},
		},
	}

	ctx := context.Background()

	// Initialize a job service client.
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		fmt.Printf("talent.NewJobClient: %v", err)
		return nil, err
	}

	// Construct a createJob request.
	req := &talentpb.CreateJobRequest{
		Parent: "projects/" + projectId,
		Job: jobToCreate,
	}

	resp, err := c.CreateJob(ctx, req)
	if err != nil {
		fmt.Printf("Failed to create job with custom attributes: %v", err)
		return nil, err
	}

	fmt.Printf("Created job with custom attributres: %q\n", resp.GetName())

	return resp, nil
}

// [END job_search_create_job_custom_attributes]
