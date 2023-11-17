package sdk

type Config struct {
	ApiVersion          string                 `json:"apiversion"`
	BridgeId            string                 `json:"bridgeid"`
	DeviceName          string                 `json:"devicename"`
	FWVersion           string                 `json:"fwversion"`
	Gateway             string                 `json:"gateway"`
	IPaddress           string                 `json:"ipaddress"`
	Localtime           string                 `json:"localtime"`
	MAC                 string                 `json:"mac"`
	ModelID             string                 `json:"modelid"`
	Name                string                 `json:"name"`
	Netmask             string                 `json:"netmask"`
	NetworkOpenDuration int64                  `json:"networkopenduration"`
	NTP                 string                 `json:"ntp"`
	PanID               int64                  `json:"panid"`
	ProxyAddress        string                 `json:"proxyaddress"`
	ProxyPort           int64                  `json:"proxyport"`
	SWVersion           string                 `json:"swversion"`
	Timeformat          string                 `json:"timeformat"`
	TimeZone            string                 `json:"timezone"`
	UTC                 string                 `json:"UTC"`
	UUID                string                 `json:"uuid"`
	WebsocketPort       int64                  `json:"websocketport"`
	Whitelist           map[string]interface{} `json:"whitelist"`
	ZigbeeChannel       int64                  `json:"zigbeechannel"`

	DHCP               bool `json:"dhcp"`
	LinkButton         bool `json:"linkbutton"`
	PortalServices     bool `json:"portalservices"`
	RFConnected        bool `json:"rfconnected"`
	WebsocketNotifyAll bool `json:"websocketnotifyall"`
}

type Device struct {
	LastAnnounced    *string `json:"lastannounced"`
	LastSeen         string  `json:"lastseen"`
	ManufacturerName string  `json:"manufacturername"`
	ModelID          string  `json:"modelid"`
	Name             string  `json:"name"`
	ProductID        string  `json:"productid"`
	SWVersion        string  `json:"swversion"`
	UniqueID         string  `json:"uniqueid"`
}

type Group struct {
	Action           GroupAction  `json:"action"`
	DeviceMembership []string     `json:"devicemembership"`
	ETag             string       `json:"etag"`
	ID               string       `json:"id"`
	Lights           []string     `json:"lights"`
	Name             string       `json:"name"`
	Scenes           []GroupScene `json:"scenes"`
	Hidden           bool         `json:"hidden"`
}

type GroupAction struct {
	Brightness       int64      `json:"bri"`
	ColorTemperature int64      `json:"ct"`
	Effect           string     `json:"effect"`
	Hue              int64      `json:"hue"`
	Saturation       int64      `json:"sat"`
	XY               [2]float64 `json:"xy"`
	On               bool       `json:"on"`
}

type GroupScene struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GroupSceneDetailed struct {
	Lights []string `json:"lights"`
	Name   string   `json:"name"`
}

type Scene struct {
	Lights []SceneLight `json:"lights"`
	Name   string       `json:"name"`
	State  int64        `json:"state"`
}

type SceneLight struct {
	Brightness       int64  `json:"bri"`
	ColorTemperature int64  `json:"ct"`
	Hue              int64  `json:"hue"`
	ID               string `json:"id"`
	Saturation       int64  `json:"sat"`
	TransitionTime   int64  `json:"transitiontime"`
	X                int64  `json:"x"`
	Y                int64  `json:"y"`
	On               bool   `json:"on"`
}

type Light struct {
	ColorCapabilities int64      `json:"colorcapabilities"`
	CtMax             int64      `json:"ctmax"`
	CtMin             int64      `json:"ctmin"`
	ETag              string     `json:"etag"`
	LastAnnounced     *string    `json:"lastannounced"`
	LastSeen          string     `json:"lastseen"`
	ManufacturerName  string     `json:"manufacturername"`
	ModelID           string     `json:"modelid"`
	Name              string     `json:"name"`
	State             LightState `json:"state"`
	SWVersion         string     `json:"swversion"`
	Type              string     `json:"type"`
	UniqueID          string     `json:"uniqueid"`
	HasColor          bool       `json:"hascolor"`
}

type LightState struct {
	Alert            string     `json:"alert"`
	Brightness       int64      `json:"bri"`
	ColorMode        string     `json:"colormode"`
	ColorTemperature int64      `json:"ct"`
	Effect           string     `json:"effect"`
	Hue              int64      `json:"hue"`
	Saturation       int64      `json:"sat"`
	XY               [2]float64 `json:"xy"`
	On               bool       `json:"on"`
	Reachable        bool       `json:"reachable"`
}

type Rule struct {
	Actions        []Action    `json:"actions"`
	Conditions     []Condition `json:"conditions"`
	Created        string      `json:"created"`
	ETag           string      `json:"etag"`
	LastTriggered  string      `json:"lasttriggered"`
	Name           string      `json:"name"`
	Owner          string      `json:"owner"`
	Periodic       int64       `json:"periodic"`
	Status         string      `json:"status"`
	TimesTriggered int64       `json:"timestriggered"`
}

type Action struct {
	Address string                 `json:"address"`
	Body    map[string]interface{} `json:"body"`
	Method  string                 `json:"method"`
}

type Condition struct {
	Address  string `json:"address"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type Sensor struct {
	Config           SensorConfig `json:"config"`
	Ep               int64        `json:"ep"`
	ETag             string       `json:"etag"`
	LastSeen         string       `json:"lastseen"`
	ManufacturerName string       `json:"manufacturername"`
	Mode             int64        `json:"mode"`
	ModelID          string       `json:"modelid"`
	Name             string       `json:"name"`
	State            SensorState  `json:"state"`
	SWVersion        string       `json:"swversion"`
	Type             string       `json:"type"`
	UniqueID         string       `json:"uniqueid"`
}

type SensorConfig struct {
	Battery   int64 `json:"battery"`
	On        bool  `json:"on"`
	Reachable bool  `json:"reachable"`
}

type SensorState struct {
	ButtonEvent uint32 `json:"buttonevent"`
	LastUpdated string `json:"lastupdated"`
}
