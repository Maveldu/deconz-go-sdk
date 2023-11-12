package sdk

type Config struct {
	ApiVersion          string                   `json:"apiversion"`
	BridgeId            string                   `json:"bridgeid"`
	DeviceName          string                   `json:"devicename"`
	FWVersion           string                   `json:"fwversion"`
	Gateway             string                   `json:"gateway"`
	IPaddress           string                   `json:"ipaddress"`
	Localtime           string                   `json:"localtime"`
	MAC                 string                   `json:"mac"`
	ModelID             string                   `json:"modelid"`
	Name                string                   `json:"name"`
	Netmask             string                   `json:"netmask"`
	NetworkOpenDuration int64                    `json:"networkopenduration"`
	NTP                 string                   `json:"ntp"`
	PanID               int64                    `json:"panid"`
	ProxyAddress        string                   `json:"proxyaddress"`
	ProxyPort           int64                    `json:"proxyport"`
	SWVersion           string                   `json:"swversion"`
	Timeformat          string                   `json:"timeformat"`
	TimeZone            string                   `json:"timezone"`
	UTC                 string                   `json:"UTC"`
	UUID                string                   `json:"uuid"`
	WebsocketPort       int64                    `json:"websocketport"`
	Whitelist           map[string]interface{}   `json:"whitelist"`
	ZigbeeChannel       int64                    `json:"zigbeechannel"`

	DHCP               bool `json:"dhcp"`
	LinkButton         bool `json:"linkbutton"`
	PortalServices     bool `json:"portalservices"`
	RFConnected        bool `json:"rfconnected"`
	WebsocketNotifyAll bool `json:"websocketnotifyall"`
}
