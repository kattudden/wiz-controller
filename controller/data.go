package controller

type StatusResponse struct {
	Method string       `json:"method"`
	Env    string       `json:"env"`
	Result PilotDetails `json:"result"`
}

type PilotDetails struct {
	Mac     string `json:"mac"`
	RSSI    int    `json:"rssi"`
	Src     string `json:"src"`
	State   bool   `json:"state"`
	SceneId int    `json:"sceneId"`
	Temp    int    `json:"temp"`
	Dimming int    `json:"dimming"`
}
