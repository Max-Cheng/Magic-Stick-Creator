package core

import (
	"Magic-Stick-Creator/config"
	"encoding/hex"
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
func ReadSN(mod int){
	buf,err:=os.Open("config.plist")
	if err != nil {
		panic(err)
	}
	switch mod {
	case 1:var data config.Clover
		decoder := plist.NewDecoder(buf)
		derr := decoder.Decode(&data)
		if derr != nil {
			fmt.Println(derr)
		}
		fmt.Println("MLB:"+data.RtVariables.RtValue.MLB)
		fmt.Println("ROM:"+hex.EncodeToString(data.RtVariables.RtValue.ROM))
		fmt.Println("SystemUUID:"+data.SystemParameters.SPvalue.CustomUUID)
		fmt.Println("SystemSerialNumber:"+data.SMBIOS.SMValue.SerialNumber)
		break
	case 2:var data config.OC
		decoder := plist.NewDecoder(buf)
		derr := decoder.Decode(&data)
		if derr != nil {
			fmt.Println(derr)
		}
		fmt.Println("MLB:"+data.PlatformInfo.PlatformInfo.Generic.MLB)
		fmt.Println("SystemSerialNumber:"+data.PlatformInfo.PlatformInfo.Generic.SystemSerialNumber)
		fmt.Println("SystemUUID:"+data.PlatformInfo.PlatformInfo.Generic.SystemUUID)
		fmt.Println("ROM:"+hex.EncodeToString(data.PlatformInfo.PlatformInfo.Generic.ROM))
		break;
	default:fmt.Println("选择错误");break;
	}
}