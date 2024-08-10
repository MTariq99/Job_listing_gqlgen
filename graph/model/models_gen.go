// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateJobListingInput struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Company     string `json:"Company"`
	URL         string `json:"URL"`
}

type DeleteJobListingResponse struct {
	DeleteJobID string `json:"DeleteJobID"`
}

type JobListing struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Company     string `json:"Company"`
	URL         string `json:"URL"`
}

type Mutation struct {
}

type Query struct {
}

type UpdateJobListingInput struct {
	Title       *string `json:"Title,omitempty"`
	Description *string `json:"Description,omitempty"`
	URL         *string `json:"URL,omitempty"`
}
