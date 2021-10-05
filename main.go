package main

import (
	"./internal/db"
	"./internal/verification"
	"archive/zip"
	"fmt"
	_ "github.com/lib/pq"
	"reflect"
)


func main() {
	garConf := db.DbConfig{
		Host:     db.Host,
		Port:     db.Port,
		User:     db.User,
		Password: db.Password,
		Dbname:   db.Dbname,
	}
	gar, _ := garConf.Connect()
	zipReader, _ := zip.OpenReader("39.zip")
	for _, zipFile := range zipReader.File {
		var region string
		if len(zipFile.Name) == 3{
			fmt.Println("----------------------streaming directory " + zipFile.Name + "----------------------")
			continue
		}
		if string(zipFile.Name[2]) == "/"{
			region = zipFile.Name[0:2]
		}
		fmt.Println("region is", region)
		fmt.Println("current file is " + zipFile.Name)


		table := verification.Verification(zipFile.Name)
		if reflect.TypeOf(table) == nil{
			continue
		}
		switch i := table.(type) {
		case db.AddHouseTypes:
			xmlFile, err := zipFile.Open()
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			var table db.AddHouseTypes
			table.Request = `INSERT INTO gar.addhouse_types(id, name, shortname, description, startdate, isactive, enddate)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`
			table.Read(gar, table.Request, xmlFile, db.Size)
			fmt.Println(i)
		case db.AddrObj:
			xmlFile, err := zipFile.Open()
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			var table db.AddrObj
			table.Request = `INSERT INTO gar.addr_obj(id, objectid, objectguid, changeid, name, typename, level, opertypeid, previd, nextid, updatedate, startdate, enddate, isactual, isactive, regioncode)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`
			table.Read(gar, table.Request, xmlFile, db.Size, region)
		case db.AddrObjParams:
			xmlFile, err := zipFile.Open()
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			var table db.AddrObjParams
			table.Request = `INSERT INTO gar.addr_obj_params(regioncode, id, objectid, changeid, changeidend, typeid, value, updatedate, startdate, enddate)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
			table.Read(gar, table.Request, xmlFile, db.Size, region)
		case db.AddrObjTypes:
			xmlFile, err := zipFile.Open()
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			var table db.AddrObjTypes
			table.Request = `INSERT INTO gar.addr_obj_types(id, level, shortname, name, description, updatedate, startdate, enddate, isactive)
			VALUES ($1, $2 ,$3 ,$4 ,$5 ,$6 ,$7 ,$8 ,$9)`
			table.Read(gar, table.Request, xmlFile, db.Size)
		case db.AdmHierarchy:
			xmlFile, err := zipFile.Open()
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			var table db.AdmHierarchy
			table.Request = `INSERT INTO gar.adm_hierarchy(id, objectid,parentobjid, changeid, areacode, citycode, placecode, plancode, streetcode, previd, nextid, updatedate,startdate,enddate,isactive,regioncode)
			VALUES ($1, $2,$3, $4, $5, $6, $7, $8, $9, $10, $11, $12,$13,$14,$15,$16)`

			table.Read(gar, table.Request, xmlFile, db.Size)
		case db.ChangeHist:
			xmlFile, err := zipFile.Open()
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			var table db.ChangeHist
			table.Request = `INSERT INTO gar.change_history(regioncode, changeid, objectid, adrobjectid, opertypeid, ndocid, changedate)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`
			table.Read(gar, table.Request, xmlFile, db.Size, region)
		case db.HouseTypes:

			xmlFile, err := zipFile.Open()
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			var table db.HouseTypes
			table.Request = `INSERT INTO gar.house_types(id, name, shortname, description, startdate, isacrive, enddate)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`
			table.Read(gar, table.Request, xmlFile, db.Size)
		case db.MunHierarchy:
			xmlFile, err := zipFile.Open()
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			var table db.MunHierarchy
			table.Request = `INSERT INTO gar.mun_hierarchy(regioncode, id, objectid, parentobjid, changeid, oktmo, previd, nextid, updatedate, startdate, enddate, isactive)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
			table.Read(gar, table.Request, xmlFile, db.Size, region)
		case db.ObjectLevels:
			xmlFile, err := zipFile.Open()
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			var table db.ChangeHist
			table.Request = `INSERT INTO gar.object_levels(level, name, updatedate, startdate, enddate, isactive)
			VALUES (:level, :name, :updatedate, :startdate, :enddate, :isactive)`
			table.Read(gar, table.Request, xmlFile, db.Size, region)
		case db.AddrObjsDiv:
			xmlFile, err := zipFile.Open()
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			var table db.AddrObjsDiv
			table.Request = `INSERT INTO gar.addr_obj_division(regioncode, id, parentid, childid, changeid)
			VALUES ($1, $2, $3, $4, $5)`
			table.Read(gar, table.Request, xmlFile, db.Size, region)
		case db.ReestrObj:
			xmlFile, err := zipFile.Open()
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			var table db.ReestrObj
			table.Request = `INSERT INTO gar.reestr_objects(regioncode, objectid, objectguid, changeid, isactive, levelid, createdate, updatedate)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
			table.Read(gar, table.Request, xmlFile, db.Size, region)
		}
	}
}
