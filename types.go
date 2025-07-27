package main

import (
	"time"
)

type Node struct {
	Active     bool     `json:"Active"`
	Addrs      any      `json:"Addrs"`
	AllowedIPs []string `json:"AllowedIPs"`
	CapMap     struct {
		SuggestExitNode any `json:"suggest-exit-node"`
	} `json:"CapMap"`
	Created        time.Time `json:"Created,omitzero"`
	CurAddr        string    `json:"CurAddr"`
	DnsName        string    `json:"DNSName,omitzero"`
	ExitNode       bool      `json:"ExitNode"`
	ExitNodeOption bool      `json:"ExitNodeOption,omitzero"`
	HostName       string    `json:"HostName,omitzero"`
	ID             string    `json:"ID,omitzero"`
	InEngine       bool      `json:"InEngine"`
	InMagicSock    bool      `json:"InMagicSock,omitzero"`
	InNetworkMap   bool      `json:"InNetworkMap,omitzero"`
	LastHandshake  time.Time `json:"LastHandshake"`
	LastSeen       time.Time `json:"LastSeen,omitzero"`
	LastWrite      time.Time `json:"LastWrite"`
	Location       struct {
		City        string  `json:"City,omitzero"`
		CityCode    string  `json:"CityCode,omitzero"`
		Country     string  `json:"Country,omitzero"`
		CountryCode string  `json:"CountryCode,omitzero"`
		Latitude    float64 `json:"Latitude,omitzero"`
		Longitude   float64 `json:"Longitude,omitzero"`
		Priority    int     `json:"Priority,omitzero"`
	} `json:"Location"`
	NoFileSharingReason string   `json:"NoFileSharingReason"`
	OS                  string   `json:"OS"`
	Online              bool     `json:"Online,omitzero"`
	PeerApiurl          any      `json:"PeerAPIURL"`
	PublicKey           string   `json:"PublicKey,omitzero"`
	Relay               string   `json:"Relay"`
	RxBytes             int      `json:"RxBytes"`
	Tags                []string `json:"Tags"`
	TaildropTarget      int      `json:"TaildropTarget,omitzero"`
	TailscaleIPs        []string `json:"TailscaleIPs"`
	TxBytes             int      `json:"TxBytes"`
	UserID              int      `json:"UserID,omitzero"`
}

type Status struct {
	AuthURL       string   `json:"AuthURL"`
	BackendState  string   `json:"BackendState,omitzero"`
	CertDomains   []string `json:"CertDomains"`
	ClientVersion struct {
		LatestVersion string `json:"LatestVersion,omitzero"`
	} `json:"ClientVersion"`
	CurrentTailnet struct {
		MagicDnsEnabled bool   `json:"MagicDNSEnabled,omitzero"`
		MagicDnsSuffix  string `json:"MagicDNSSuffix,omitzero"`
		Name            string `json:"Name,omitzero"`
	} `json:"CurrentTailnet"`
	ExitNodeStatus struct {
		ID           string   `json:"ID,omitzero"`
		Online       bool     `json:"Online,omitzero"`
		TailscaleIPs []string `json:"TailscaleIPs"`
	} `json:"ExitNodeStatus"`
	HaveNodeKey    bool            `json:"HaveNodeKey,omitzero"`
	Health         []any           `json:"Health"`
	MagicDnsSuffix string          `json:"MagicDNSSuffix,omitzero"`
	Peer           map[string]Node `json:"Peer,omitzero"`
	Self           struct {
		Active     bool     `json:"Active"`
		Addrs      []string `json:"Addrs"`
		AllowedIPs []string `json:"AllowedIPs"`
		CapMap     struct {
			AutoExitNode                          any   `json:"auto-exit-node"`
			Funnel                                any   `json:"funnel"`
			HTTPS                                 any   `json:"https"`
			Https___tailscale_com_cap_fileSharing any   `json:"https://tailscale.com/cap/file-sharing"`
			Https___tailscale_com_cap_isAdmin     any   `json:"https://tailscale.com/cap/is-admin"`
			HTTPS___Tailscale_Com_Cap_SSH         any   `json:"https://tailscale.com/cap/ssh"`
			Mullvad                               any   `json:"mullvad"`
			ProbeUdpLifetime                      any   `json:"probe-udp-lifetime"`
			SSHBehaviorV1                         any   `json:"ssh-behavior-v1"`
			SSHEnvVars                            any   `json:"ssh-env-vars"`
			StoreAppcRoutes                       any   `json:"store-appc-routes"`
			SuggestExitNodeUi                     any   `json:"suggest-exit-node-ui"`
			Tailnet_MaxKeyDuration                []int `json:"tailnet.maxKeyDuration"`
			// "https://tailscale.com/cap/funnel-ports?ports=443,8443,10000" cannot be unmarshalled into a struct field by encoding/json.
		} `json:"CapMap"`
		Capabilities        []string  `json:"Capabilities"`
		Created             time.Time `json:"Created,omitzero"`
		CurAddr             string    `json:"CurAddr"`
		DnsName             string    `json:"DNSName,omitzero"`
		ExitNode            bool      `json:"ExitNode"`
		ExitNodeOption      bool      `json:"ExitNodeOption"`
		HostName            string    `json:"HostName,omitzero"`
		ID                  string    `json:"ID,omitzero"`
		InEngine            bool      `json:"InEngine"`
		InMagicSock         bool      `json:"InMagicSock"`
		InNetworkMap        bool      `json:"InNetworkMap,omitzero"`
		KeyExpiry           time.Time `json:"KeyExpiry,omitzero"`
		LastHandshake       time.Time `json:"LastHandshake"`
		LastSeen            time.Time `json:"LastSeen"`
		LastWrite           time.Time `json:"LastWrite"`
		NoFileSharingReason string    `json:"NoFileSharingReason"`
		OS                  string    `json:"OS,omitzero"`
		Online              bool      `json:"Online,omitzero"`
		PeerApiurl          []string  `json:"PeerAPIURL"`
		PublicKey           string    `json:"PublicKey,omitzero"`
		Relay               string    `json:"Relay,omitzero"`
		RxBytes             int       `json:"RxBytes"`
		TaildropTarget      int       `json:"TaildropTarget"`
		TailscaleIPs        []string  `json:"TailscaleIPs"`
		TxBytes             int       `json:"TxBytes"`
		UserID              int       `json:"UserID,omitzero"`
	} `json:"Self"`
	Tun          bool     `json:"TUN,omitzero"`
	TailscaleIPs []string `json:"TailscaleIPs"`
	Version      string   `json:"Version,omitzero"`
}
