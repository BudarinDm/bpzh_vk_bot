package domain

type Keyboard struct {
	OneTime bool       `json:"one_time,omitempty"`
	Buttons [][]Button `json:"buttons"`
	Inline  bool       `json:"inline"`
}

type Button struct {
	Action Action `json:"action"`
	Color  string `json:"color"`
}

type Action struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
	Label   string `json:"label"`
}
