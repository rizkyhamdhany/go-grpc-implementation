package app

import (
	"github.com/rizkyhamdhany/jobagency-client-service/jobagencyclientpb"
	"context"
)

func (a *App) JobAll(ctx context.Context, req *jobagencyclientpb.JobAllRequest) (*jobagencyclientpb.JobAllResponse, error) {
	rows, err := a.DB.Raw("SELECT id, client_id, uuid, title FROM jobs").Rows()
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var jobs []*jobagencyclientpb.Job
	for rows.Next() {
		var job jobagencyclientpb.Job
		rows.Scan(&job.Id, &job.ClienId, &job.Uuid, &job.Title)
		jobs = append(jobs, &job)
	}

	res := &jobagencyclientpb.JobAllResponse{
		Job: jobs,
	}
	return res, nil
}
