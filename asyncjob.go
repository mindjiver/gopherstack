package gopherstack

import (
	"log"
	"net/url"
)

type QueryAsyncJobResultResponse struct {
	Queryasyncjobresultresponse struct {
		Accountid     string  `json:"accountid"`
		Cmd           string  `json:"cmd"`
		Created       string  `json:"created"`
		Jobid         string  `json:"jobid"`
		Jobprocstatus float64 `json:"jobprocstatus"`
		Jobresultcode float64 `json:"jobresultcode"`
		Jobstatus     float64 `json:"jobstatus"`
		Userid        string  `json:"userid"`
	} `json:"queryasyncjobresultresponse"`
}

// Query CloudStack for the state of a scheduled job
func (c CloudStackClient) QueryAsyncJobResult(jobid string) (float64, error) {
	params := url.Values{}
	params.Set("jobid", jobid)
	response, err := NewRequest(c, "queryAsyncJobResult", params)

	if err != nil {
		return -1, err
	}

	log.Printf("response: %v", response)
	status := response.(QueryAsyncJobResultResponse).Queryasyncjobresultresponse.Jobstatus

	return status, err
}
