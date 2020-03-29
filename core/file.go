package core

import (
	"fmt"
	"howett.net/plist"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
func ChangeSN(mod int){
	//1:Clover to OC.
	//2:OC to OC
	//Read User config.plist
	Target,err:=os.Open("clover.plist")
	if err != nil {
		panic(err)
	}
	var Targetdata map[string]interface{}
	decoder := plist.NewDecoder(Target)
	derr := decoder.Decode(&Targetdata)
	if derr != nil {
		fmt.Println(derr)
	}
	//Read EFI plist
	OC,err:=os.Open("EFI/OC/config.plist")
	if err != nil {
		panic(err)
	}
	var OCdata map[string]interface{}
	EFIdecoder := plist.NewDecoder(OC)
	EFIderr := EFIdecoder.Decode(&OCdata)
	if EFIderr != nil {
		fmt.Println(derr)
	}
	//Create New config.plist
	newfile,err:=os.Create("EFI/OC/config.plist")
	if err!=nil {
		panic(err)
	}
	//New Encoder
	a:=plist.NewEncoder(newfile)
	a.Indent("    ")
	switch mod {
	case 1:
		OCdata["PlatformInfo"].(map[string]interface{})["Generic"].(map[string]interface{})["MLB"]=Targetdata["SMBIOS"].(map[string]interface{})["BoardSerialNumber"]
		OCdata["PlatformInfo"].(map[string]interface{})["Generic"].(map[string]interface{})["ROM"]=Targetdata["RtVariables"].(map[string]interface{})["ROM"]
		OCdata["PlatformInfo"].(map[string]interface{})["Generic"].(map[string]interface{})["SystemSerialNumber"]=Targetdata["SMBIOS"].(map[string]interface{})["SerialNumber"]
		OCdata["PlatformInfo"].(map[string]interface{})["Generic"].(map[string]interface{})["SystemUUID"]=Targetdata["SystemParameters"].(map[string]interface{})["CustomUUID"]
		a.Encode(OCdata)
		break
	case 2:
		OCdata["PlatformInfo"].(map[string]interface{})["Generic"].(map[string]interface{})["MLB"]=Targetdata["PlatformInfo"].(map[string]interface{})["Generic"].(map[string]interface{})["MLB"]
		OCdata["PlatformInfo"].(map[string]interface{})["Generic"].(map[string]interface{})["ROM"]=Targetdata["PlatformInfo"].(map[string]interface{})["Generic"].(map[string]interface{})["ROM"]
		OCdata["PlatformInfo"].(map[string]interface{})["Generic"].(map[string]interface{})["SystemSerialNumber"]=Targetdata["PlatformInfo"].(map[string]interface{})["Generic"].(map[string]interface{})["SystemSerialNumber"]
		OCdata["PlatformInfo"].(map[string]interface{})["Generic"].(map[string]interface{})["SystemUUID"]=Targetdata["PlatformInfo"].(map[string]interface{})["Generic"].(map[string]interface{})["SystemUUID"]
		a.Encode(OCdata)
		break;
	default:fmt.Println("选择错误");break;
	}
}
