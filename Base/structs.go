package Base

type Accesses struct {
	Certification bool         `json:"Certification"`
	PermitTime    string       `json:"PermitTime"`
	Level         int          `json:"Level"`
	PermitLog     []PermitLogs `json:"Permit_log"`
}

type AccountHas struct {
	Status StatusData `json:"Status"`
	Has    bool       `json:"Has"`
}

type AccountPW struct {
	Password string `json:"Password"`
}

type CarData struct {
	Status      StatusData `json:"Status"`
	CarID       string     `json:"CarID"`
	CarName     string     `json:"CarName"`
	RefreshTime string     `json:"RefreshTime"`
	CreateTime  string     `json:"CreateTime"`
}

type CarIDReturn struct {
	Status StatusData `json:"Status"`
	ID     string     `json:"ID"`
	CarID  string     `json:"CarID"`
	Token  string     `json:"Token"`
}

type CarNews struct {
	ID               string `json:"ID"`
	CarID            string `json:"CarID"`
	CarName          string `json:"CarName"`
	TemporarilyToken string `json:"TemporarilyToken"`
}

type CreateReturn struct {
	Status StatusData `json:"Status"`
	ID     string     `json:"ID"`
}

type Historys struct {
	Times    string `json:"Times"`
	UseToken string `json:"UseToken"`
	Types    int    `json:"Types"`
	Device   string `json:"Device"`
}

type InputToken struct {
	Token string `json:"Token"`
}

type LogInToken struct {
	Status       StatusData `json:"Status"`
	GetTimes     string     `json:"GetTimes"`
	AccountToken string     `json:"AccountToken"`
	AccountID    string     `json:"AccountID"`
}

type MonitorData struct {
	Status         StatusData        `json:"Status"`
	WaterStatus    MonitorStatusData `json:"WaterStatus"`
	GasolineStatus MonitorStatusData `json:"GasolineStatus"`
	BatteryStatus  MonitorStatusData `json:"BatteryStatus"`
}

type MonitorStatus struct {
	ID           string `json:"ID"`
	Token        string `json:"Token"`
	CarID        string `json:"CarID"`
	SelectObject string `json:"SelectObject"`
	StatusCode   int    `json:"StatusCode"`
}

type MonitorStatusData struct {
	StatusCode  int    `json:"StatusCode"`
	RefreshTime string `json:"RefreshTime"`
}

type NewAccountIDPW struct {
	Account  string `json:"Account"`
	Password string `json:"Password"`
}

type NewAccountUser struct {
	Name    string `json:"Name"`
	Gender  int    `json:"Gender"`
	Country string `json:"Country"`
	Number  string `json:"Number"`
}

type NewCarName struct {
	ID      string `json:"ID"`
	CarID   string `json:"CarID"`
	CarName string `json:"CarName"`
}

type PermitLogs struct {
	Level     int    `json:"Level"`
	Times     string `json:"Times"`
	Authority string `json:"Authority"`
}

type Phones struct {
	Country string `json:"Country"`
	Number  string `json:"Number"`
}

type Profiles struct {
	Name   string `json:"Name"`
	Gender int    `json:"Gender"`
	Phone  Phones `json:"Phone"`
}

type SecurityData struct {
	Status       StatusData           `json:"Status"`
	DoorStatus   []SecurityStatusData `json:"DoorStatus"`
	WindowStatus []SecurityStatusData `json:"WindowStatus"`
	LightStatus  []SecurityStatusData `json:"LightStatus"`
}

type SecurityStatus struct {
	ID           string `json:"ID"`
	Token        string `json:"Token"`
	CarID        string `json:"CarID"`
	Name         string `json:"Name"`
	SelectObject string `json:"SelectObject"`
	StatusCode   int    `json:"StatusCode"`
}

type SecurityStatusData struct {
	Name        string `json:"Name"`
	StatusCode  int    `json:"StatusCode"`
	RefreshTime string `json:"RefreshTime"`
}

type StatusData struct {
	StatusCode  int    `json:"StatusCode"`
	Description string `json:"Description"`
}

type TemporarilyTokenData struct {
	Status   StatusData `json:"Status"`
	Token    string     `json:"Token"`
	GetTimes string     `json:"GetTimes"`
}

type Users struct {
	Status        StatusData `json:"Status"`
	Car           []CarData  `json:"Car"`
	Profile       Profiles   `json:"Profile"`
	Accesse       Accesses   `json:"Accesse"`
	SiginHistory  Historys   `json:"SiginHistory"`
	LogoutHistory Historys   `json:"LogoutHistory"`
}
