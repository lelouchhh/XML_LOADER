package verification

import (
	"../db"
)

var TablesDict = map[string]interface{}{
	"AS_ADDHOUSE_TYPES" 	: db.AddHouseTypes{},
	"AS_ADDR_OBJ" 			: db.AddrObj{} ,
	"AS_ADDR_OBJ_DIVISION" 	: db.AddrObjsDiv{},
	"AS_ADDR_OBJ_PARAMS"	: db.AddrObjParams{},
	"AS_ADDR_OBJ_TYPES" 	: db.AddrObjTypes{},
	"AS_ADM_HIERARCHY" 		: db.AdmHierarchy{},
	"AS_CHANGE_HISTORY" 	: db.ChangeHist{},
	"AS_HOUSE_TYPES" 		: db.HouseTypes{},
	"AS_MUN_HIERARCHY" 		: db.MunHierarchy{},
	"AS_OBJECT_LEVELS" 		: db.ObjectLevels{},
	"AS_REESTR_OBJECTS" 	: db.ReestrObj{},
}


func Verification(n string) (t interface{}, s string){
	if string(n[2]) == "/"{
		return TablesDict[n[3:len(n)-50]], n[6:len(n)-50]
	}else {
		return TablesDict[n[0:len(n)-50]], n[3:len(n)-50]
	}
}

