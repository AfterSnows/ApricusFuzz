package risk

type Risk struct {
	IP string `json:"ip"`

	Url  string `json:"url"`
	Port int    `json:"port"`
	Host string `json:"host"`

	Title       string `json:"title"`
	Description string `json:"description"`
	RiskType    string `json:"risk_type"`
	Payload     string `json:"payload"`
	Details     string `json:"details"`
	Severity    string `json:"severity"`
}
