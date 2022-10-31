package model

const (
	// MKConfigKey ...
	MKConfigKey = "mock_config"
)

// MKToggle Config ...
type MKToggle struct {
	Method     int      `json:"method"`
	PFB        string   `json:"pfb"`
	MainAPI    string   `json:"main_api"`
	SubAPIList []string `json:"sub_api_list"`
}

// MKCenter Config ...
type MKCenter struct {
	MockURL   string `json:"mock_url"`
	RecordURL string `json:"record_url"`
}

// SpexMockConfig ...
type SpexMockConfig struct {
	Toggle   []MKToggle `json:"toggle"`
	MKCenter MKCenter   `json:"mock_center"`
}
