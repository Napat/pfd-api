package model

type (
	BeefSummaryResponse struct {
		BeefSummary map[string]map[string]int
	}

	BeeflistAdaptorResponse struct {
		BeefSummary map[string]map[string]int
	}
)
