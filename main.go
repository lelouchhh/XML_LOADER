package main

import (
	"./internal/config"
	"./internal/db"
	"./internal/verification"
	"archive/zip"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

func main() {
	f, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "info", log.LstdFlags)
	var conf config.Conf
	conf.GetConf()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Db.Host, conf.Db.Port, conf.Db.User, conf.Db.Password, conf.Db.Dbname)

	gar, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil{
		logger.Println(err)
		db.ErrorLog(gar, err.Error())
	}else{
		db.MsgLog(gar, "successful connection to the db")
		logger.Println("Successful connection")
	}
	logger.Println("ZIP opening...")
	db.MsgLog(gar, "ZIP opening")
	zipReader, err := zip.OpenReader(conf.FilePath)
	if err != nil{
		logger.Println(err)
		db.ErrorLog(gar, err.Error())
	}else{
		logger.Println("ZIP file opened successfully")
		db.MsgLog(gar, "ZIP file opened successfully")
	}
	regions := db.Regions(gar)
	TotalTime := time.Now()
	logger.Println("Start reading ZIP file")
	db.MsgLog(gar, "Start reading ZIP file")
	for _, zipFile := range zipReader.File {
		var region string
		if len(zipFile.Name) == 3{
			db.MsgLog(gar, "Current dir is: " + zipFile.Name)
			logger.Println("Current dir is " + zipFile.Name)
			continue
		}
		if string(zipFile.Name[2]) == "/"{
			region = zipFile.Name[0:2]
			logger.Println("Current region is ", region)
			db.MsgLog(gar, "Current region is: " + region)
			if !InRegion(region, regions){
				continue
			}

		}
		logger.Println("Current file is " + zipFile.Name)
		db.MsgLog(gar, "Current file is " + zipFile.Name)


		table, tableS := verification.Verification(zipFile.Name)
		if reflect.TypeOf(table) == nil{
			continue
		}
		switch i := table.(type) {
		case db.AddHouseTypes:

			xmlFile, err := zipFile.Open()
			if err != nil {
				logger.Println("Error opening file:", err)
				db.ErrorLog(gar, err.Error())
				return
			}else{
				logger.Println("File successfully opened")
				db.MsgLog(gar, "File successfully opened")

			}
			var table db.AddHouseTypes
			table.Request = `INSERT INTO gar.addhouse_types(id, name, shortname, description, startdate, isactive, enddate)
			VALUES (:id, :name, :shortname, :description, :startdate, :isactive, :enddate)`
			t := time.Now()
			logger.Println("Clearing table...")
			db.MsgLog(gar, "Starting the cleaning function")
			err = db.ClearItable(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			} else{
				db.MsgLog(gar, "Table successfully cleared")
				db.MsgLog(gar, "Clearing time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully cleared ")
				logger.Println("Clearing time is: ", time.Since(t))
			}
			logger.Println("Writing into the database...")
			db.MsgLog(gar, "Writing into the database")
			table.Read(gar, table.Request, xmlFile, db.Size)
			logger.Println("Successful recording")
			fmt.Println(
				"reading and inserting time: ",
				time.Since(t),
			)
			logger.Println("Segmenting table...")
			err = db.CreateIndexSegment(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			}else{
				db.MsgLog(gar, "Table successfully segmented")
				db.MsgLog(gar, "Segmenting time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully segmented ")
				logger.Println("Segmenting time is: ", time.Since(t))
			}
		case db.AddrObj:
			xmlFile, err := zipFile.Open()
			if err != nil {
				logger.Println("Error opening file:", err)
				db.ErrorLog(gar, err.Error())
				return
			}else{
				logger.Println("File successfully opened")
				db.MsgLog(gar, "File successfully opened")

			}
			var table db.AddrObj
			table.Request = `INSERT INTO gar.addr_obj(id, objectid, objectguid, changeid, name, typename, level, opertypeid, previd, nextid, updatedate, startdate, enddate, isactual, isactive, regioncode)
			VALUES (:id, :objectid, :objectguid, :changeid, :name, :typename, :level, :opertypeid, :previd, :nextid, :updatedate, :startdate, :enddate, :isactual, :isactive, :regioncode)`
			t := time.Now()
			logger.Println("Clearing table...")
			db.MsgLog(gar, "Starting the cleaning function")
			err = db.ClearItable(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			} else{
				db.MsgLog(gar, "Table successfully cleared")
				db.MsgLog(gar, "Clearing time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully cleared ")
				logger.Println("Clearing time is: ", time.Since(t))
			}
			logger.Println("Writing into the database...")
			db.MsgLog(gar, "Writing into the database")
			table.Read(gar, table.Request, xmlFile, db.Size, region)
			logger.Println("Successful recording")
			fmt.Println("reading and inserting time: ", time.Since(t))
			logger.Println("Segmenting table...")
			err = db.CreateIndexSegment(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			}else{
				db.MsgLog(gar, "Table successfully segmented")
				db.MsgLog(gar, "Segmenting time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully segmented ")
				logger.Println("Segmenting time is: ", time.Since(t))
			}
		case db.AddrObjParams:
			xmlFile, err := zipFile.Open()
			if err != nil {
				logger.Println("Error opening file:", err)
				db.ErrorLog(gar, err.Error())
				return
			}else{
				logger.Println("File successfully opened")
				db.MsgLog(gar, "File successfully opened")

			}
			var table db.AddrObjParams
			table.Request = `INSERT INTO gar.addr_obj_params(regioncode, id, objectid, changeid, changeidend, typeid, value, updatedate, startdate, enddate)
			VALUES (:regioncode, :id, :objectid, :changeid, :changeidend, :typeid, :value, :updatedate, :startdate, :enddate)`
			t := time.Now()
			logger.Println("Clearing table...")
			db.MsgLog(gar, "Starting the cleaning function")
			err = db.ClearItable(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			} else{
				db.MsgLog(gar, "Table successfully cleared")
				db.MsgLog(gar, "Clearing time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully cleared ")
				logger.Println("Clearing time is: ", time.Since(t))
			}
			logger.Println("Writing into the database...")
			db.MsgLog(gar, "Writing into the database")
			table.Read(gar, table.Request, xmlFile, db.Size, region)
			logger.Println("Successful recording")
			fmt.Println("reading and inserting time: ", time.Since(t))
			logger.Println("Segmenting table...")
			err = db.CreateIndexSegment(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			}else{
				db.MsgLog(gar, "Table successfully segmented")
				db.MsgLog(gar, "Segmenting time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully segmented ")
				logger.Println("Segmenting time is: ", time.Since(t))
			}
		case db.AddrObjTypes:
			xmlFile, err := zipFile.Open()
			if err != nil {
				logger.Println("Error opening file:", err)
				db.ErrorLog(gar, err.Error())
				return
			}else{
				logger.Println("File successfully opened")
				db.MsgLog(gar, "File successfully opened")

			}
			var table db.AddrObjTypes
			table.Request = `INSERT INTO gar.addr_obj_types(id, level, shortname, name, description, updatedate, startdate, enddate, isactive)
			VALUES (:id, :level ,:shortname ,:name ,:description ,:updatedate ,:startdate ,:enddate ,:isactive)`
			t := time.Now()
			logger.Println("Clearing table...")
			db.MsgLog(gar, "Starting the cleaning function")
			err = db.ClearItable(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			} else{
				db.MsgLog(gar, "Table successfully cleared")
				db.MsgLog(gar, "Clearing time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully cleared ")
				logger.Println("Clearing time is: ", time.Since(t))
			}
			logger.Println("Writing into the database...")
			db.MsgLog(gar, "Writing into the database")
			table.Read(gar, table.Request, xmlFile, db.Size)
			logger.Println("Successful recording")
			fmt.Println("reading and inserting time: ", time.Since(t))
			logger.Println("Segmenting table...")
			err = db.CreateIndexSegment(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			}else{
				db.MsgLog(gar, "Table successfully segmented")
				db.MsgLog(gar, "Segmenting time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully segmented ")
				logger.Println("Segmenting time is: ", time.Since(t))
			}

		case db.AdmHierarchy:
			xmlFile, err := zipFile.Open()
			if err != nil {
				logger.Println("Error opening file:", err)
				db.ErrorLog(gar, err.Error())
				return
			}else{
				logger.Println("File successfully opened")
				db.MsgLog(gar, "File successfully opened")

			}
			var table db.AdmHierarchy
			table.Request = `INSERT INTO gar.adm_hierarchy(id, objectid,parentobjid, changeid, areacode, citycode, placecode, plancode, streetcode, previd, nextid, updatedate,startdate,enddate,isactive,regioncode)
			VALUES (:id, :objectid,:parentobjid, :changeid, :areacode, :citycode, :placecode, :plancode, :streetcode, :previd, :nextid, :updatedate,:startdate,:enddate,:isactive,:regioncode)`
			t := time.Now()
			logger.Println("Clearing table...")
			db.MsgLog(gar, "Starting the cleaning function")
			err = db.ClearItable(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			} else{
				db.MsgLog(gar, "Table successfully cleared")
				db.MsgLog(gar, "Clearing time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully cleared ")
				logger.Println("Clearing time is: ", time.Since(t))
			}
			logger.Println("Writing into the database...")
			db.MsgLog(gar, "Writing into the database")
			table.Read(gar, table.Request, xmlFile, db.Size)
			logger.Println("Successful recording")
			fmt.Println("reading and inserting time: ", time.Since(t))
			logger.Println("Segmenting table...")
			err = db.CreateIndexSegment(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			}else{
				db.MsgLog(gar, "Table successfully segmented")
				db.MsgLog(gar, "Segmenting time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully segmented ")
				logger.Println("Segmenting time is: ", time.Since(t))
			}
		case db.ChangeHist:
			xmlFile, err := zipFile.Open()
			if err != nil {
				logger.Println("Error opening file:", err)
				db.ErrorLog(gar, err.Error())
				return
			}else{
				logger.Println("File successfully opened")
				db.MsgLog(gar, "File successfully opened")

			}
			var table db.ChangeHist
			table.Request = `INSERT INTO gar.change_history(regioncode, changeid, objectid, adrobjectid, opertypeid, ndocid, changedate)
			VALUES (:regioncode, :changeid, :objectid, :adrobjectid, :opertypeid, :ndocid, :changedate)`
			t := time.Now()
			logger.Println("Clearing table...")
			db.MsgLog(gar, "Starting the cleaning function")
			err = db.ClearItable(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			} else{
				db.MsgLog(gar, "Table successfully cleared")
				db.MsgLog(gar, "Clearing time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully cleared ")
				logger.Println("Clearing time is: ", time.Since(t))
			}
			logger.Println("Writing into the database...")
			db.MsgLog(gar, "Writing into the database")
			table.Read(gar, table.Request, xmlFile, db.Size, region)
			logger.Println("Successful recording")
			fmt.Println("reading and inserting time: ", time.Since(t))
			logger.Println("Segmenting table...")
			err = db.CreateIndexSegment(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			}else{
				db.MsgLog(gar, "Table successfully segmented")
				db.MsgLog(gar, "Segmenting time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully segmented ")
				logger.Println("Segmenting time is: ", time.Since(t))
			}

		case db.HouseTypes:

			xmlFile, err := zipFile.Open()
			if err != nil {
				logger.Println("Error opening file:", err)
				db.ErrorLog(gar, err.Error())
				return
			}else{
				logger.Println("File successfully opened")
				db.MsgLog(gar, "File successfully opened")

			}
			var table db.HouseTypes
			table.Request = `INSERT INTO gar.house_types(id, name, shortname, description, startdate, isactive, enddate)
			VALUES (:id, :name, :shortname, :description, :startdate, :isactive, :enddate)`
			t := time.Now()
			logger.Println("Clearing table...")
			db.MsgLog(gar, "Starting the cleaning function")
			err = db.ClearItable(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			} else{
				db.MsgLog(gar, "Table successfully cleared")
				db.MsgLog(gar, "Clearing time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully cleared ")
				logger.Println("Clearing time is: ", time.Since(t))
			}
			logger.Println("Writing into the database...")
			db.MsgLog(gar, "Writing into the database")
			table.Read(gar, table.Request, xmlFile, db.Size)
			logger.Println("Successful recording")
			fmt.Println("reading and inserting time: ", time.Since(t))
			logger.Println("Segmenting table...")
			err = db.CreateIndexSegment(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			}else{
				db.MsgLog(gar, "Table successfully segmented")
				db.MsgLog(gar, "Segmenting time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully segmented ")
				logger.Println("Segmenting time is: ", time.Since(t))
			}

		case db.MunHierarchy:
			xmlFile, err := zipFile.Open()
			if err != nil {
				logger.Println("Error opening file:", err)
				db.ErrorLog(gar, err.Error())
				return
			}else{
				logger.Println("File successfully opened")
				db.MsgLog(gar, "File successfully opened")

			}
			var table db.MunHierarchy
			table.Request = `INSERT INTO gar.mun_hierarchy(regioncode, id, objectid, parentobjid, changeid, oktmo, previd, nextid, updatedate, startdate, enddate, isactive)
			VALUES (:regioncode, :id, :objectid, :parentobjid, :changeid, :oktmo, :previd, :nextid, :updatedate, :startdate, :enddate, :isactive)`
			t := time.Now()
			logger.Println("Clearing table...")
			db.MsgLog(gar, "Starting the cleaning function")
			err = db.ClearItable(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			} else{
				db.MsgLog(gar, "Table successfully cleared")
				db.MsgLog(gar, "Clearing time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully cleared ")
				logger.Println("Clearing time is: ", time.Since(t))
			}
			logger.Println("Writing into the database...")
			db.MsgLog(gar, "Writing into the database")
			table.Read(gar, table.Request, xmlFile, db.Size, region)
			logger.Println("Successful recording")
			fmt.Println("reading and inserting time: ", time.Since(t))
			logger.Println("Segmenting table...")
			err = db.CreateIndexSegment(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			}else{
				db.MsgLog(gar, "Table successfully segmented")
				db.MsgLog(gar, "Segmenting time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully segmented ")
				logger.Println("Segmenting time is: ", time.Since(t))
			}

		case db.ObjectLevels:
			xmlFile, err := zipFile.Open()
			if err != nil {
				logger.Println("Error opening file:", err)
				db.ErrorLog(gar, err.Error())
				return
			}else{
				logger.Println("File successfully opened")
				db.MsgLog(gar, "File successfully opened")

			}
			var table db.ObjectLevels
			table.Request = `INSERT INTO gar.object_levels(level, name, updatedate, startdate, enddate, isactive)
			VALUES (:level, :name, :updatedate, :startdate, :enddate, :isactive)`
			t := time.Now()
			logger.Println("Clearing table...")
			db.MsgLog(gar, "Starting the cleaning function")
			err = db.ClearItable(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			} else{
				db.MsgLog(gar, "Table successfully cleared")
				db.MsgLog(gar, "Clearing time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully cleared ")
				logger.Println("Clearing time is: ", time.Since(t))
			}
			logger.Println("Writing into the database...")
			db.MsgLog(gar, "Writing into the database")
			table.Read(gar, table.Request, xmlFile, db.Size)
			logger.Println("Successful recording")
			fmt.Println("reading and inserting time: ", time.Since(t))
			logger.Println("Segmenting table...")
			err = db.CreateIndexSegment(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			}else{
				db.MsgLog(gar, "Table successfully segmented")
				db.MsgLog(gar, "Segmenting time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully segmented ")
				logger.Println("Segmenting time is: ", time.Since(t))
			}

		case db.AddrObjsDiv:
			xmlFile, err := zipFile.Open()
			if err != nil {
				logger.Println("Error opening file:", err)
				db.ErrorLog(gar, err.Error())
				return
			}else{
				logger.Println("File successfully opened")
				db.MsgLog(gar, "File successfully opened")

			}
			var table db.AddrObjsDiv
			table.Request = `INSERT INTO gar.addr_obj_division(regioncode, id, parentid, childid, changeid)
			VALUES (:regioncode, :id, :parentid, :childid, :changeid)`
			t := time.Now()
			logger.Println("Clearing table...")
			db.MsgLog(gar, "Starting the cleaning function")
			err = db.ClearItable(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			} else{
				db.MsgLog(gar, "Table successfully cleared")
				db.MsgLog(gar, "Clearing time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully cleared ")
				logger.Println("Clearing time is: ", time.Since(t))
			}
			logger.Println("Writing into the database...")
			db.MsgLog(gar, "Writing into the database")
			table.Read(gar, table.Request, xmlFile, db.Size, region)
			logger.Println("Successful recording")
			fmt.Println("reading and inserting time: ", time.Since(t))
			logger.Println("Segmenting table...")
			err = db.CreateIndexSegment(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			}else{
				db.MsgLog(gar, "Table successfully segmented")
				db.MsgLog(gar, "Segmenting time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully segmented ")
				logger.Println("Segmenting time is: ", time.Since(t))
			}
		case db.ReestrObj:
			xmlFile, err := zipFile.Open()
			if err != nil {
				logger.Println("Error opening file:", err)
				db.ErrorLog(gar, err.Error())
				return
			}else{
				logger.Println("File successfully opened")
				db.MsgLog(gar, "File successfully opened")

			}
			var table db.ReestrObj
			table.Request = `INSERT INTO gar.reestr_objects(regioncode, objectid, objectguid, changeid, isactive, levelid, createdate, updatedate)
			VALUES (:regioncode, :objectid, :objectguid, :changeid, :isactive, :levelid, :createdate, :updatedate)`
			t := time.Now()
			logger.Println("Clearing table...")
			db.MsgLog(gar, "Starting the cleaning function")
			err = db.ClearItable(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			} else{
				db.MsgLog(gar, "Table successfully cleared")
				db.MsgLog(gar, "Clearing time is: "+strconv.FormatInt(int64(time.Since(t)), 10))
				logger.Println("Table successfully cleared ")
				logger.Println("Clearing time is: ", time.Since(t))
			}
			logger.Println("Writing into the database...")
			db.MsgLog(gar, "Writing into the database")
			table.Read(gar, table.Request, xmlFile, db.Size, region)
			logger.Println("Successful recording")
			fmt.Println("reading and inserting time: ", time.Since(t))
			logger.Println("Segmenting table...")
			err = db.CreateIndexSegment(gar, tableS, region)
			if err != nil{
				logger.Println("Error: "+err.Error())
				db.ErrorLog(gar, err.Error())
			}else{
				db.MsgLog(gar, "Table successfully segmented")
				db.MsgLog(gar, "Segmenting time is: "+ string(time.Since(t)))
				logger.Println("Table successfully segmented ")
				logger.Println("Segmenting time is: ", time.Since(t))
			}

		default:
			fmt.Println(i)
		}
	}
	fmt.Println("Total time is: ", time.Since(TotalTime))
}

func InRegion(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}