package main

import (
	"./internal/db"
	"./internal/verification"
	"archive/zip"
	"encoding/xml"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
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
	zipReader, _ := zip.OpenReader("gigatest.zip")
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

		gar, _ := garConf.Connect()

		t := verification.Verification(zipFile.Name)
		if reflect.TypeOf(t) == nil{
			continue
		}

		switch i := t.(type) {
		case db.AddHouseTypes:
			var t db.AddHouseTypes
			unzippedFileBytes, err := readZipFile(zipFile)
			t.Request = `INSERT INTO gar.addhouse_types(id, name, shortname, description, startdate, isactive)
			VALUES (:id, :name, :shortname, :description, :startdate, :isactive)`
			if err != nil {
				log.Println(err)
				continue
			}
			err = xml.Unmarshal(unzippedFileBytes, &t)
			if err != nil{
				panic(err)
			}
			db.Insert(gar, t.Request, t.Attr)
			fmt.Println(i)
		case db.AddrObj:
			// v is a float64 here, so e.g. v + 1.0 is possible.
			t := &db.AddrObj{}
			unzippedFileBytes, err := readZipFile(zipFile)
			t.Request = `INSERT INTO gar.addr_obj(id, objectid, objectguid, changeid, name, typename, level, opertypeid, previd, nextid, updatedate, startdate, enddate, isactual, isactive, regioncode)
			VALUES (:id, :objectid, :objectguid, :changeid, :name, :typename, :level, :opertypeid, :previd, :nextid, :updatedate, :startdate, :enddate, :isactual, :isactive, :regioncode)`
			if err != nil {
				log.Println(err)
				continue
			}

			err = xml.Unmarshal(unzippedFileBytes, &t)
			fmt.Println(t.Attr)
			for i:=0;i<=len(t.Attr)-1;i++ {
				t.Attr[i].RegionCode = region
			}
			if err != nil{
				panic(err)
			}
			db.Insert(gar, t.Request, t.Attr)
		case db.AddrObjParams:
			// v is a float64 here, so e.g. v + 1.0 is possible.
			var t db.AddrObjParams
			unzippedFileBytes, err := readZipFile(zipFile)
			if err != nil {
				log.Println(err)
				continue
			}
			//fmt.Println("SUKA",reflect.TypeOf(table))

			err = xml.Unmarshal(unzippedFileBytes, &t)
			for i:=0;i<=len(t.Attr)-1;i++ {
				t.Attr[i].RegionCode = region
			}
			t.Request = `INSERT INTO gar.addr_obj_params(id, objectid, changeid, changeidend, typeid, value, startdate, updatedate, enddate, regioncode)
			VALUES (:id, :objectid, :changeid, :changeidend, :typeid, :value, :startdate, :updatedate, :enddate, :regioncode)`
			//fmt.Println(&t)
			if err != nil{
				panic(err)
			}
			db.Insert(gar, t.Request, t.Attr)
		case db.AddrObjTypes:
			// v is a float64 here, so e.g. v + 1.0 is possible.
			var t db.AddrObjTypes
			unzippedFileBytes, err := readZipFile(zipFile)
			if err != nil {
				log.Println(err)
				continue
			}
			//fmt.Println("SUKA",reflect.TypeOf(table))

			err = xml.Unmarshal(unzippedFileBytes, &t)

			t.Request = `INSERT INTO gar.addr_obj_types(id, level, shortname, name, _desc, updatedate, startdate, enddate, isactive)
			VALUES (:id, :level, :shortname, :name, :_desc, :updatedate, :startdate, :enddate, :isactive)`
			//fmt.Println(&t)
			if err != nil{
				panic(err)
			}
			db.Insert(gar, t.Request, t.Attr)
		case db.AdmHierarchy:
			// v is a float64 here, so e.g. v + 1.0 is possible.
			var t db.AdmHierarchy
			unzippedFileBytes, err := readZipFile(zipFile)
			if err != nil {
				log.Println(err)
				continue
			}
			//fmt.Println("SUKA",reflect.TypeOf(table))

			err = xml.Unmarshal(unzippedFileBytes, &t)
			for i:=0;i<=len(t.Attr)-1;i++ {
				t.Attr[i].RegionCode = region
			}
			t.Request = `INSERT INTO gar.adm_hierarchy(id, objectid,parentobjid, changeid, areacode, citycode, placecode, plancode, streetcode, previd, nextid, updatedate,startdate,enddate,isactive,regioncode)
			VALUES (:id, :objectid,:parentobjid, :changeid, :areacode, :citycode, :placecode, :plancode, :streetcode, :previd, :nextid, :updatedate,:startdate,:enddate,:isactive,:regioncode)`
			fmt.Println(&t)
			if err != nil{
				panic(err)
			}
			db.Insert(gar, t.Request, t.Attr)
		case db.ChangeHist:
			// v is a float64 here, so e.g. v + 1.0 is possible.
			var t db.ChangeHist
			unzippedFileBytes, err := readZipFile(zipFile)

			if err != nil {
				log.Println(err)
				continue
			}
			//fmt.Println("SUKA",reflect.TypeOf(table))

			err = xml.Unmarshal(unzippedFileBytes, &t)
			t.Request = `INSERT INTO gar.change_history(changeid, objectid, adrobjectid, opertypeid, ndocid, changedate, regioncode)
			VALUES (:changeid, :objectid, :adrobjectid, :opertypeid, :ndocid, :changedate, :regioncode)`
			for i:=0;i<=len(t.Attr)-1;i++ {
				t.Attr[i].RegionCode = region
			}
			//fmt.Println(&t)
			if err != nil{
				panic(err)
			}
			db.Insert(gar, t.Request, t.Attr)
		case db.HouseTypes:
			// v is a float64 here, so e.g. v + 1.0 is possible.
			var t db.HouseTypes
			unzippedFileBytes, err := readZipFile(zipFile)
			t.Request = `INSERT INTO gar.house_types(id, name, shortname, _desc, updatedate, startdate, enddate, isactive)
			VALUES (:id, :name, :shortname, :_desc, :updatedate, :startdate, :enddate, :isactive)`
			if err != nil {
				log.Println(err)
				continue
			}
			//fmt.Println("SUKA",reflect.TypeOf(table))
			fmt.Println("-------------------------")
			fmt.Println(reflect.TypeOf(t))
			err = xml.Unmarshal(unzippedFileBytes, &t)

			fmt.Println(&t)
			if err != nil{
				panic(err)
			}
			db.Insert(gar, t.Request, t.Attr)
		case db.MunHierarchy:
			// v is a float64 here, so e.g. v + 1.0 is possible.
			var t db.MunHierarchy
			unzippedFileBytes, err := readZipFile(zipFile)
			t.Request = `INSERT INTO gar.mun_hierarchy(id, objectid, parentobjid, changeid, oktmo, previd, nextid, updatedate, startdate, enddate, isactive, regioncode)
			VALUES (:id, :objectid, :parentobjid, :changeid, :oktmo, :previd, :nextid, :updatedate, :startdate, :enddate, :isactive, :regioncode)`
			if err != nil {
				log.Println(err)
				continue
			}
			//fmt.Println("SUKA",reflect.TypeOf(table))

			err = xml.Unmarshal(unzippedFileBytes, &t)
			for i:=0;i<=len(t.Attr)-1;i++ {
				t.Attr[i].RegionCode = region
			}
			//fmt.Println(&t)
			if err != nil{
				panic(err)
			}
			db.Insert(gar, t.Request, t.Attr)
		case db.ObjectLevels:
			// v is a float64 here, so e.g. v + 1.0 is possible.
			var t db.ObjectLevels
			unzippedFileBytes, err := readZipFile(zipFile)
			t.Request = `INSERT INTO gar.object_levels(level, name, updatedate, startdate, enddate, isactive)
			VALUES (:level, :name, :updatedate, :startdate, :enddate, :isactive)`
			if err != nil {
				log.Println(err)
				continue
			}
			//fmt.Println("SUKA",reflect.TypeOf(table))

			err = xml.Unmarshal(unzippedFileBytes, &t)

			//fmt.Println(&t)
			if err != nil{
				panic(err)
			}
			db.Insert(gar, t.Request, t.Attr)
		case db.AddrObjsDiv:
			// v is a float64 here, so e.g. v + 1.0 is possible.
			var t db.AddrObjsDiv
			unzippedFileBytes, err := readZipFile(zipFile)
			t.Request = `INSERT INTO gar.addr_obj_division(id, praentid, childid, changeid, regioncode)
			VALUES (:id, :praentid, :childid, :changeid, :regioncode)`
			if err != nil {
				log.Println(err)
				continue
			}
			//fmt.Println("SUKA",reflect.TypeOf(table))

			err = xml.Unmarshal(unzippedFileBytes, &t)
			for i:=0;i<=len(t.Attr)-1;i++ {
				t.Attr[i].RegionCode = region
			}
			fmt.Println(t.Attr)
			if err != nil{
				panic(err)
			}
			db.Insert(gar, t.Request, t.Attr)
		default:
			// v is a float64 here, so e.g. v + 1.0 is possible.
			var t db.ReestrObj
			unzippedFileBytes, err := readZipFile(zipFile)
			t.Request = `INSERT INTO gar.reestr_objects(objectid, objectguid, changeid, isactive, levelid, createdate, updatedate, regioncode)
			VALUES (:objectid, :objectguid, :changeid, :isactive, :levelid, :createdate, :updatedate, :regioncode)`
			if err != nil {
				log.Println(err)
				continue
			}
			//fmt.Println("SUKA",reflect.TypeOf(table))

			err = xml.Unmarshal(unzippedFileBytes, &t)
			for i:=0;i<=len(t.Attr)-1;i++ {
				t.Attr[i].RegionCode = region
			}
			//fmt.Println(&t)
			if err != nil{
				panic(err)
			}
			db.Insert(gar, t.Request, t.Attr)
		}
	}
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}