package config

type Vertion struct {
	BID string
	SN  string
	Vs  string
}

//Vs:OS name
var Path = "com.apple.recovery.boot"
var Select = []Vertion{Default, Last, Catelina, Mojave, High_Sierra}
var Last = Vertion{BID: "Mac-E43C1C25D4880AD6", SN: "00000000000000000", Vs: "Last Version(最新版)"}
var Default = Vertion{BID: "Mac-E43C1C25D4880AD6", SN: "00000000000GDVQ00", Vs: "Default(默认)"}
var Catelina = Vertion{BID: "Mac-00BE6ED71E35EB8", SN: "00000000000000000", Vs: "Catelina"}
var Mojave = Vertion{BID: "Mac-7BA5B2DFE22DDD8C", SN: "00000000000KXPG00", Vs: "Mojave"}
var High_Sierra = Vertion{BID: "Mac-7BA5B2D9E42DDD94", SN: "00000000000J80300", Vs: "High Sierra"}
