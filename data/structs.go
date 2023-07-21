package data

type SiteData struct {
	ServerInfo *ServerInfo
	Players    []Player
	Maps       []string
	Characters []string
	Mods       []string
	Other      []string
}

type ServerInfo struct {
	MapName *string
	Time    string
}

type Player struct {
	PlayerName  *string
	Skin        *string
	Score       *float32
	IsSpectator bool
}
