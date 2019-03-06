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

  talent "cloud.google.com/go/talent/apiv4beta1"
	"google.golang.org/api/iterator"
	talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"
)

// [START job_search_custom_ranking_search]

// customRankingSearch searches for jobs based on custom ranking.
func customRankingSearch(w io.Writer, projectId, companyName string) error {
	ctx := context.Background()

	// Initialize a jobService client.
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		fmt.Printf("talent.NewJobClient: %v", err)
		return err
	}

	// Construct a searchJobs request.
	req := &talentpb.SearchJobsRequest{
		Parent: "projects/" + projectId,
		// Make sure to set the RequestMetadata the same as the associated
		// search request.
		RequestMetadata: &talentpb.RequestMetadata{
			// Make sure to hash your userID.
			UserId: "HashedUsrId",
			// Make sure to hash the sessionID.
			SessionId: "HashedSessionId",
			// Domain of the website where the search is conducted.
			Domain: "www.googlesample.com",
		},
		JobQuery: &talentpb.JobQuery{
			CompanyNames: []string{companyName},
		},
    // More info on customRankingInfo.
		// https://godoc.org/google.golang.org/genproto/googleapis/cloud/talent/v4beta1#SearchJobsRequest_CustomRankingInfo
    CustomRankingInfo: &talentpb.SearchJobsRequest_CustomRankingInfo{
      ImportanceLevel: 6,
      RankingExpression: "(someFieldLong + 25) * 0.25 - anotherFieldLong",
    },
    OrderBy: "custom_ranking desc",
	}

	it := c.SearchJobs(ctx, req)

	for {
		resp, err := it.Next()
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			fmt.Printf("it.Next: %v", err)
			return err
		}
		fmt.Fprintf(w, "Job: %q\n", resp.Job.GetName())
	}
}

// [END job_search_custom_ranking_search]
