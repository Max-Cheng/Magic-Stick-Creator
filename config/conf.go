package config

type Vertion struct {
	Vs string
	BID string
	SN	string
}
type OC struct {
	PlatformInfo
}
type PlatformInfo struct {
	PlatformInfo Generic `plist:"PlatformInfo"`
}
type Generic struct {
	Generic OCValue `plist:"Generic"`
}
type OCValue struct {
	MLB string `plist:"MLB"`
	SystemSerialNumber string `plist:"SystemSerialNumber"`
	SystemUUID string `plist:"SystemUUID"`
	ROM []byte `plist:"ROM"`
}
type Clover struct {
	RtVariables
	SMBIOS
	SystemParameters
}
type SystemParameters struct {
	SPvalue SPvalue `plist:"SystemParameters"`
}
type SPvalue struct {
	CustomUUID string `plist:"CustomUUID"`
}
type RtVariables struct {
	RtValue RtValue `plist:"RtVariables"`
}
type RtValue struct {
	MLB string `plist:"MLB"`
	ROM []byte `plist:"ROM"`
}
type SMBIOS struct {
	SMValue SMValue `plist:SMBIOS`
}
type SMValue struct {
	SerialNumber string `plist:"SerialNumber"`
}
var Path = "com.apple.recovery.boot"
var Select = []Vertion{Last,Mojave,High_Sierra}
var	Last = Vertion{BID:"Mac-E43C1C25D4880AD6",SN:"00000000000000000",Vs:"Last Version(最新版)"}
var Mojave =Vertion{BID:"Mac-7BA5B2DFE22DDD8C",SN:"00000000000KXPG00",Vs:"Mojave"}
var High_Sierra = Vertion{BID:"Mac-7BA5B2D9E42DDD94",SN:"00000000000J80300",Vs:"High Sierra"}