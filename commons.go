package common

// GetCommonData a method
func GetCommonData() string {
	return "Common Data"
}

func GetData() string {
	return "Data"
}

type Device struct {
	ID          string     `json:"id"`
	DESCRIPTION string     `json:"description"`
	IP          string     `json:"ip"`
	STATE       string     `json:"state"`
	SCHEDULES   []Schedule `json:"schedules"`
	ONLINE      bool       `json:"online"`
}

type Devices struct {
	DEVICES []Device `json:"devices"`
}

type Schedule struct {
	START string `json:"start"`
	STOP  string `json:"stop"`
}
