package config

type Vertion struct {
	Vs string
	BID string
	SN	string
}
var Path = "com.apple.recovery.boot"
var Select = []Vertion{Last,Mojave,High_Sierra}
var	Last = Vertion{BID:"Mac-E43C1C25D4880AD6",SN:"00000000000000000",Vs:"Last Version(最新版)"}
var Mojave =Vertion{BID:"Mac-7BA5B2DFE22DDD8C",SN:"00000000000KXPG00",Vs:"Mojave"}
var High_Sierra = Vertion{BID:"Mac-7BA5B2D9E42DDD94",SN:"00000000000J80300",Vs:"High Sierra"}