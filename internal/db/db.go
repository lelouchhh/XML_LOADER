package db

import (
	"../config"
	"database/sql"
	"encoding/xml"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"io"
	"strings"
)

var Size = getBatch()
func getBatch() int{
	var conf config.Conf
	conf.GetConf()
	return conf.Batch
}

func Connect(host, password, dbname, user string, port int) (*sql.DB, error){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully connected!")
	return db, err
}
func Regions(db *sqlx.DB) []string{
	var regincode []string
	err := db.Select(&regincode,"SELECT regioncode FROM gar.t_import_filter")
	if err != nil {
		// handle this error better than this
		fmt.Println(err)
	}
	return regincode
}

type Reader interface {
	Read()
}

func (s *AddHouseTypes) Read(db *sqlx.DB, r string, f io.ReadCloser, size int) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	s.Attr = make([]HouseType, Size+1)
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				_, err := db.NamedExec(s.Request, s.Attr[:total])
				if err != nil {
					fmt.Println(err)
				}
			}
			break
		}
		for {
			t, _ = decoder.Token()
			//fmt.Println(t)
			if t == nil {
				break
			}
			switch se := t.(type) {
			case xml.StartElement:
				inElement = se.Name.Local
				if inElement == "HOUSETYPE" {
					var offer HouseType
					decoder.DecodeElement(&offer, &se)
					s.Attr[total] = offer
					total++
				}
			default:
			}
			if total % Size == 0 && total != 0{
				_, err := db.NamedExec(s.Request, s.Attr[:Size])
				if err != nil {
					fmt.Println(err)
				}
				break
			}
		}
	}
}
func (s *AddrObj) Read(db *sqlx.DB, r string, f io.ReadCloser, size int, code string) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	s.Attr = make([]Object, Size+1)
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				_, err := db.NamedExec(s.Request, s.Attr[0:total])
				if err != nil {
					fmt.Println(err)
				}
			}
			break
		}
		for {
			t, _ = decoder.Token()
			//fmt.Println(t)
			if t == nil {
				break
			}
			switch se := t.(type) {
			case xml.StartElement:
				inElement = se.Name.Local
				if inElement == "OBJECT" {
					var offer Object
					decoder.DecodeElement(&offer, &se)
					offer.RegionCode = code
					s.Attr[total] = offer
					total++
				}
			default:
			}
			if total % Size == 0 && total != 0{
				//fmt.Println(s.Attr)
				_, err := db.NamedExec(s.Request, s.Attr[:Size])
				if err != nil {
					fmt.Println(err)
				}
				total = 0
				break
			}
		}
	}
}
func (s *AddrObjsDiv) Read(db *sqlx.DB, r string, f io.ReadCloser, size int, code string) {
	decoder := xml.NewDecoder(f)
	fmt.Println(r)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	s.Attr = make([]AddrObjsItem, Size+1)
	for {
		if t == nil {
			if total != 0 {
				_, err := db.NamedExec(s.Request, s.Attr[:total])
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		break
		for {
			t, _ = decoder.Token()
			//fmt.Println(t)
			if t == nil {
				break
			}
			switch se := t.(type) {
			case xml.StartElement:
				inElement = se.Name.Local
				if inElement == "ITEM" {
					var offer AddrObjsItem
					offer.RegionCode = code
					decoder.DecodeElement(&offer, &se)
					s.Attr[total] = offer
					total++
				}
			default:
			}
			if total%Size == 0 && total != 0 {
				_, err := db.NamedExec(s.Request, s.Attr[:Size])
				if err != nil {
					fmt.Println(err)
				}
				total = 0
				break
			}
		}
	}
}
func (s *AddrObjParams) Read(db *sqlx.DB, r string, f io.ReadCloser, size int, code string) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	s.Attr = make([]Params, Size+1)
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				_, err := db.NamedExec(s.Request, s.Attr[:total])
				if err != nil {
					fmt.Println(err)
				}
			}
			break
		}
		for {
			t, _ = decoder.Token()
			//fmt.Println(t)
			if t == nil {
				break
			}
			switch se := t.(type) {
			case xml.StartElement:
				inElement = se.Name.Local
				if inElement == "PARAM" {
					var offer Params
					decoder.DecodeElement(&offer, &se)
					offer.RegionCode = code
					s.Attr[total] = offer
					total++
				}
			default:
			}
			if total % Size == 0 && total != 0{
				_, err := db.NamedExec(s.Request, s.Attr[:Size])
				if err != nil {
					fmt.Println(err)
				}
				total = 0
			}
			break
		}
	}
}
func (s *MunHierarchy) Read(db *sqlx.DB, r string, f io.ReadCloser, size int, code string) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	s.Attr = make([]MunHierarchyItem, Size+1)
	for{
		if t == nil {
			if total != 0 {
				_, err := db.NamedExec(s.Request, s.Attr[:total])
				if err != nil {
					fmt.Println(err)
				}
			}
			break
		}
		for {
			t, _ = decoder.Token()
			if t == nil {
				break
			}
			switch se := t.(type) {
			case xml.StartElement:
				inElement = se.Name.Local
				if inElement == "ITEM" {
					var offer MunHierarchyItem
					offer.RegionCode = code
					decoder.DecodeElement(&offer, &se)
					s.Attr[total] = offer
					total++
				}
			default:
			}
			if total % Size == 0 && total != 0{
					_, err := db.NamedExec(s.Request, s.Attr[:Size])
					if err != nil {
						fmt.Println(err)
				}
				total = 0

				break
			}
		}
	}
}
func (s *ChangeHist) Read(db *sqlx.DB, r string, f io.ReadCloser, size int, code string) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	s.Attr = make([]ChangeHistItem, Size+1)
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				_, err := db.NamedExec(s.Request, s.Attr[0:total])
				if err != nil {
					fmt.Println(err)
				}
			}
			break
		}
		for {
			t, _ = decoder.Token()
			//fmt.Println(t)
			if t == nil {
				break
			}
			switch se := t.(type) {
			case xml.StartElement:
				inElement = se.Name.Local
				if inElement == "ITEM" {
					var offer ChangeHistItem
					decoder.DecodeElement(&offer, &se)
					offer.RegionCode = code
					s.Attr[total] = offer
					total++
				}
			default:
			}
			if total % Size == 0 && total != 0{
				//fmt.Println(s.Attr)
				_, err := db.NamedExec(s.Request, s.Attr[:Size])
				if err != nil {
					fmt.Println(err)
				}
				total = 0
				break
			}
		}
	}
}
func (s *AddrObjTypes) Read(db *sqlx.DB, r string, f io.ReadCloser, size int) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	s.Attr = make([]AddrObjType, Size+1)
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				_, err := db.NamedExec(s.Request, s.Attr[:total])
				if err != nil {
					fmt.Println(err)
				}
			}
			break
		}
		for {
			t, _ = decoder.Token()
			//fmt.Println(t)
			if t == nil {
				break
			}
			switch se := t.(type) {
			case xml.StartElement:
				inElement = se.Name.Local
				if inElement == "ADDRESSOBJECTTYPE" {
					var offer AddrObjType
					decoder.DecodeElement(&offer, &se)
					s.Attr[total] = offer
					total++
				}
			default:
			}
			if total % Size == 0 && total != 0{
				_, err := db.NamedExec(s.Request, s.Attr[:Size])
				if err != nil {
					fmt.Println(err)
				}
				total = 0
			}
			break
		}
	}
}
func (s *HouseTypes) Read(db *sqlx.DB, r string, f io.ReadCloser, size int) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	s.Attr = make([]HouseType, Size+1)
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				_, err := db.NamedExec(s.Request, s.Attr[0:total])
				if err != nil {
					fmt.Println(err)
				}
			}
			break
		}
		for {
			t, _ = decoder.Token()
			//fmt.Println(t)
			if t == nil {
				break
			}
			switch se := t.(type) {
			case xml.StartElement:
				inElement = se.Name.Local
				if inElement == "HOUSETYPE" {
					var offer HouseType
					decoder.DecodeElement(&offer, &se)
					s.Attr[total] = offer
					total++
				}
			default:
			}
			if total % Size == 0 && total != 0{
				_, err := db.NamedExec(s.Request, s.Attr[:Size])
				if err != nil {
					fmt.Println(err)
				}
				total = 0
			}
			break
		}
	}
}
func (s *AdmHierarchy) Read(db *sqlx.DB, r string, f io.ReadCloser, size int) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	s.Attr = make([]AdmHierarchyItem, Size+1)
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				_, err := db.NamedExec(s.Request, s.Attr[:total])
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		break
	}
	for {
		t, _ = decoder.Token()
		//fmt.Println(t)
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			inElement = se.Name.Local
			if inElement == "ITEM" {
				var offer AdmHierarchyItem
				decoder.DecodeElement(&offer, &se)
				s.Attr[total] = offer
				total++
			}
		default:
		}
		if total % Size == 0 && total != 0{

			_, err := db.NamedExec(s.Request, s.Attr[:Size])
			if err != nil {
				fmt.Println(err)
			}
			total = 0

			break
		}
	}
}
func (s *ObjectLevels) Read(db *sqlx.DB, r string, f io.ReadCloser, size int) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	s.Attr = make([]ObjectLevel, Size+1)
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				_, err := db.NamedExec(s.Request, s.Attr[:total])
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		break
	}
	for {
		t, _ = decoder.Token()
		//fmt.Println(t)
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			inElement = se.Name.Local
			if inElement == "OBJECTLEVEL" {
				var offer ObjectLevel
				decoder.DecodeElement(&offer, &se)
				s.Attr[total] = offer
				total++
			}
		default:
		}
		if total % Size == 0 && total != 0{

			_, err := db.NamedExec(s.Request, s.Attr[:Size])
			if err != nil {
				fmt.Println(err)
			}
			total = 0

			break
		}
	}
}
func (s *ReestrObj) Read(db *sqlx.DB, r string, f io.ReadCloser, size int, code string) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	s.Attr = make([]ReestrObjObj, Size+1)
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				_, err := db.NamedExec(s.Request, s.Attr[0:total])
				if err != nil {
					fmt.Println(err)
				}
			}
			break
		}
		for {
			t, _ = decoder.Token()
			//fmt.Println(t)
			if t == nil {
				break
			}
			switch se := t.(type) {
			case xml.StartElement:
				inElement = se.Name.Local
				if inElement == "OBJECT" {
					var offer ReestrObjObj
					decoder.DecodeElement(&offer, &se)
					offer.RegionCode = code
					s.Attr[total] = offer
					total++
				}
			default:
			}
			if total % Size == 0 && total != 0{
				//fmt.Println(s.Attr)
				_, err := db.NamedExec(s.Request, s.Attr[:Size])
				if err != nil {
					fmt.Println(err)
				}
				total = 0
				break
			}
		}
	}
}
func ClearItable(db *sqlx.DB, t, c string) error{
	sqlStatement := "call gar.gar_clear_itable ('"+strings.ToLower(t)+"', '"+c+"');"
	_, err := db.Exec(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
func CreateIndexSegment (db *sqlx.DB, t, c string) error{
	sqlStatement := "call gar.gar_create_index_segment ('"+strings.ToLower(t)+"', '"+c+"');"
	_, err := db.Exec(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
func ErrorLog (db *sqlx.DB, msg string) error{
	sqlStatement := "call gar.loader_error_msg('"+msg+"');"
	_, err := db.Exec(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
func MsgLog(db *sqlx.DB, msg string) error{
	sqlStatement := "call gar.loader_notify_msg('"+msg+"');"
	_, err := db.Exec(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
