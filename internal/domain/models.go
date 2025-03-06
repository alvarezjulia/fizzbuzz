package domain

// Request represents the input parameters for the fizzbuzz endpoint
type Request struct {
	FirstDivisor  int    `json:"firstDivisor"`
	SecondDivisor int    `json:"secondDivisor"`
	Limit         int    `json:"limit"`
	FirstWord     string `json:"firstWord"`
	SecondWord    string `json:"secondWord"`
}

// Response represents the output for the fizzbuzz endpoint
type Response struct {
	Result []string `json:"result"`
}

// Stats represents the statistics for the most frequent request
type Stats struct {
	Parameters Request `json:"parameters"`
	Hits       int     `json:"hits"`
}
