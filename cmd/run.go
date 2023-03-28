package cmdpackage

import (
	"Rpcon/common"
)

func Parse() {
	common.Parse()
	if common.Method != "" {
		switch common.Method {
		case "ssh":
			common.Port = 22
			Sshcmd()
		case "redis":
			common.Port = 6379
			Codes()
		case "mysql":
			common.Port = 3306
			Query()
		case "mssql":
			common.Port = 1433
			Mssquery()
		case "wmi":
			common.Port = 135
			WmiExec()
		}
	}
	if common.Redis != "" {
		switch common.Redis {
		case "rk":
			Wkey()
		case "rs":
			Wshell()
		case "rn":
			ncshell()
		case "ms":
			Mysqlshell()
		case "mg":
			Myconfig()
		case "msg":
			Msscfg()
		case "xcmd":
			Cmdshell()
		case "xol":
			Oleshell()
		}
	} else {
		con()
	}
}
