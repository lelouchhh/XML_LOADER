package db

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"log"
)

const (
	Host     = "178.154.254.105"
	Port     = 5432
	User     = "inotech"
	Password = "platex"
	Dbname   = "postgres"
	Size	 = 5000
)


type DbConfig struct {
	Host string
	Port int
	User string
	Password string
	Dbname string
}
func (dc DbConfig) Connect() (*sql.DB,error){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully connected!")
	return db, err
}


type Reader interface {
	Read()
}

func (s *AddHouseTypes) Read(db *sql.DB, r string, f io.ReadCloser, size int) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				for i := 0; i < total; i++ {
					_, err := db.Exec(r, s.Attr[i].ID, s.Attr[i].Name, s.Attr[i].ShortName, s.Attr[i].Desc, s.Attr[i].StartDate, s.Attr[i].IsActive, s.Attr[i].EndDate)
					if err != nil {
						//fmt.Println(err)
					}
				}
				break
			}
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
				for i:=0; i<size; i++ {
					_, err := db.Exec(r, s.Attr[i].ID, s.Attr[i].Name, s.Attr[i].ShortName, s.Attr[i].Desc, s.Attr[i].StartDate, s.Attr[i].IsActive, s.Attr[i].EndDate)
					if err != nil {
						//fmt.Println(err)
					}
					total = 0
				}
				break
			}
		}
	}
}
func (s *AddrObj) Read(db *sql.DB, r string, f io.ReadCloser, size int, code string) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				ctx := context.Background()
				tx, err := db.BeginTx(ctx, nil)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(total)
				for i:=0; i<size; i++ {
					_, err := db.ExecContext(ctx, r, s.Attr[i].ID, s.Attr[i].ObjectID, s.Attr[i].ObjectUID, s.Attr[i].ChangeID,
						s.Attr[i].Name, s.Attr[i].TypeName, s.Attr[i].Level, s.Attr[i].OperTypeID, s.Attr[i].prevID,
						s.Attr[i].nextID, s.Attr[i].UpdateDate, s.Attr[i].StartDate, s.Attr[i].EndDate,
						s.Attr[i].IsActual, s.Attr[i].IsActive, code)
					if err != nil {
						//fmt.Println(err)
					}
				}
				err = tx.Commit()
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
					s.Attr[total] = offer
					total++
				}
			default:
			}

			if total % Size == 0 && total != 0{
				ctx := context.Background()
				tx, err := db.BeginTx(ctx, nil)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(total)
				for i:=0; i<size; i++ {
					_, err := db.ExecContext(ctx, r, s.Attr[i].ID, s.Attr[i].ObjectID, s.Attr[i].ObjectUID, s.Attr[i].ChangeID,
						s.Attr[i].Name, s.Attr[i].TypeName, s.Attr[i].Level, s.Attr[i].OperTypeID, s.Attr[i].prevID,
						s.Attr[i].nextID, s.Attr[i].UpdateDate, s.Attr[i].StartDate, s.Attr[i].EndDate,
						s.Attr[i].IsActual, s.Attr[i].IsActive, code)
					if err != nil {
						//fmt.Println(err)
					}
				}
				total = 0
				err = tx.Commit()
				break
			}
		}
	}
}
func (s *AddrObjsDiv) Read(db *sql.DB, r string, f io.ReadCloser, size int, code string) {
	decoder := xml.NewDecoder(f)
	fmt.Println(r)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	for{
		if t == nil {
			if total != 0 {
				for i := 0; i < total; i++ {
					db.Exec(r, code, s.Attr[i].ID, s.Attr[i].ParentID, s.Attr[i].ChildID, s.Attr[i].ChangeID)
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
					var offer AddrObjsItem
					decoder.DecodeElement(&offer, &se)
					s.Attr[total] = offer
					total++
				}
			default:
			}
			if total % Size == 0 && total != 0{
				for i:=0; i<size; i++ {
					_, err := db.Exec(r, code, s.Attr[i].ID, s.Attr[i].ParentID, s.Attr[i].ChildID, s.Attr[i].ChangeID)
					if err != nil {
						break
					}
					total = 0
				}
				break
			}
		}
	}
}
func (s *AddrObjParams) Read(db *sql.DB, r string, f io.ReadCloser, size int, code string) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				for i := 0; i < total; i++ {
					_, err := db.Exec(r, code, s.Attr[i].ID, s.Attr[i].ObjectID, s.Attr[i].ChangeID, s.Attr[i].ChangeIdEnd,s.Attr[i].TypeID, s.Attr[i].Value,  s.Attr[i].UpdateDate, s.Attr[i].StartDate, s.Attr[i].EndDate )
					if err != nil {
						//fmt.Println(err)
					}
				}
				break
			}
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
					s.Attr[total] = offer
					total++
				}
			default:
			}
			if total % Size == 0 && total != 0{
				for i:=0; i<size; i++ {
					_, err := db.Exec(r, code, s.Attr[i].ID, s.Attr[i].ObjectID, s.Attr[i].ChangeID, s.Attr[i].ChangeIdEnd,s.Attr[i].TypeID, s.Attr[i].Value,  s.Attr[i].UpdateDate, s.Attr[i].StartDate, s.Attr[i].EndDate )
					if err != nil {
						//fmt.Println(err)
					}
					total = 0
				}
				break
			}
		}
	}
}
func (s *MunHierarchy) Read(db *sql.DB, r string, f io.ReadCloser, size int, code string) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	for{
		fmt.Println(t)
		if t == nil {
			if total != 0 {
				for i := 0; i < total; i++ {
					_, err := db.Exec(r, code, s.Attr[i].ID, s.Attr[i].ObjectID, s.Attr[i].ParentObjID, s.Attr[i].ChangeID,s.Attr[i].Oktmo, s.Attr[i].PrevID,  s.Attr[i].NextID, s.Attr[i].UpdateDate, s.Attr[i].StartDate, s.Attr[i].EndDate, s.Attr[i].IsActive )
					if err != nil {
						fmt.Println(err)
					}
				}
			}
			break
		}
		for {
			t, _ = decoder.Token()
			fmt.Println(t)
			if t == nil {
				break
			}
			switch se := t.(type) {
			case xml.StartElement:
				inElement = se.Name.Local
				fmt.Println(inElement)
				if inElement == "ITEM" {
					var offer MunHierarchyItem
					decoder.DecodeElement(&offer, &se)
					s.Attr[total] = offer
					total++
				}
			default:
			}
			if total % Size == 0 && total != 0{
				for i:=0; i<size; i++ {
					_, err := db.Exec(r, code, s.Attr[i].ID, s.Attr[i].ObjectID, s.Attr[i].ParentObjID, s.Attr[i].ChangeID,s.Attr[i].Oktmo, s.Attr[i].PrevID,  s.Attr[i].NextID, s.Attr[i].UpdateDate, s.Attr[i].StartDate, s.Attr[i].EndDate, s.Attr[i].IsActive )
					if err != nil {
						fmt.Println(err)
					}
					total = 0
				}
				break
			}
		}
	}
}
func (s *ChangeHist) Read(db *sql.DB, r string, f io.ReadCloser, size int, code string) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				for i := 0; i < total; i++ {
					_, err := db.Exec(r, code, s.Attr[i].ChangeID, s.Attr[i].ObjectID, s.Attr[i].AdrobjectID,s.Attr[i].OperTypeID, s.Attr[i].NdocID,  s.Attr[i].ChangeDate)
					if err != nil {
						//fmt.Println(err)
					}
				}
				break
			}
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
					s.Attr[total] = offer
					total++
				}
			default:
			}
			if total % Size == 0 && total != 0{
				for i:=0; i<size; i++ {
					_, err := db.Exec(r, code, s.Attr[i].ChangeID, s.Attr[i].ObjectID, s.Attr[i].AdrobjectID,s.Attr[i].OperTypeID, s.Attr[i].NdocID,  s.Attr[i].ChangeDate)
					if err != nil {
						fmt.Println(err)
					}
					total = 0
				}
				break
			}
		}
	}
}
func (s *AddrObjTypes) Read(db *sql.DB, r string, f io.ReadCloser, size int) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				for i := 0; i < total; i++ {
					_, err := db.Exec(r, s.Attr[i].ID, s.Attr[i].Level, s.Attr[i].ShortName, s.Attr[i].Name,s.Attr[i].Desc, s.Attr[i].UpdateDate,  s.Attr[i].StartDate,s.Attr[i].EndDate, s.Attr[i].IsActive)
					if err != nil {
						//fmt.Println(err)
					}
				}
				break
			}
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
				for i:=0; i<size; i++ {
					_, err := db.Exec(r, s.Attr[i].ID, s.Attr[i].Level, s.Attr[i].ShortName, s.Attr[i].Name,s.Attr[i].Desc, s.Attr[i].UpdateDate,  s.Attr[i].StartDate,s.Attr[i].EndDate, s.Attr[i].IsActive)
					if err != nil {
						//fmt.Println(err)
					}
					total = 0
				}
				break
			}
		}
	}
}
func (s *HouseTypes) Read(db *sql.DB, r string, f io.ReadCloser, size int) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				for i := 0; i < total; i++ {
					_, err := db.Exec(r, s.Attr[i].ID, s.Attr[i].Name, s.Attr[i].ShortName, s.Attr[i].Desc, s.Attr[i].StartDate, s.Attr[i].IsActive, s.Attr[i].EndDate)
					if err != nil {
						//fmt.Println(err)
					}
				}
				break
			}
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
				for i:=0; i<size; i++ {
					_, err := db.Exec(r, s.Attr[i].ID, s.Attr[i].Name, s.Attr[i].ShortName, s.Attr[i].Desc, s.Attr[i].StartDate, s.Attr[i].IsActive, s.Attr[i].EndDate)
					if err != nil {
						//fmt.Println(err)
					}
					total = 0
				}
				break
			}
		}
	}
}
func (s *AdmHierarchy) Read(db *sql.DB, r string, f io.ReadCloser, size int) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				for i := 0; i < total; i++ {
					_, err := db.Exec(r, s.Attr[i].ID, s.Attr[i].ObjectID, s.Attr[i].ParentObjID, s.Attr[i].ChangeID,
						s.Attr[i].AreaCode, s.Attr[i].CityCode, s.Attr[i].PlaceCode, s.Attr[i].PlanCode,
						s.Attr[i].StreetCode, s.Attr[i].PrevID, s.Attr[i].NextID, s.Attr[i].UpdateDate,
						s.Attr[i].StartDate, s.Attr[i].EndDate, s.Attr[i].IsActive, s.Attr[i].RegionCode)
					if err != nil {
						//fmt.Println(err)
					}
				}
				break
			}
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
				for i:=0; i<size; i++ {
					_, err := db.Exec(r, s.Attr[i].ID, s.Attr[i].ObjectID, s.Attr[i].ParentObjID, s.Attr[i].ChangeID,
						s.Attr[i].AreaCode, s.Attr[i].CityCode, s.Attr[i].PlaceCode, s.Attr[i].PlanCode,
						s.Attr[i].StreetCode, s.Attr[i].PrevID, s.Attr[i].NextID, s.Attr[i].UpdateDate,
						s.Attr[i].StartDate, s.Attr[i].EndDate, s.Attr[i].IsActive, s.Attr[i].RegionCode)
					if err != nil {
						//fmt.Println(err)
					}
					total = 0
				}
				break
			}
		}
	}
}
func (s *ObjectLevels) Read(db *sql.DB, r string, f io.ReadCloser, size int) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				for i := 0; i < total; i++ {
					_, err := db.Exec(r, s.Attr[i].Level, s.Attr[i].Name, s.Attr[i].UpdateDate, s.Attr[i].StartDate,
						s.Attr[i].EndDate, s.Attr[i].IsActive)
					if err != nil {
						//fmt.Println(err)
					}
				}
				break
			}
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
				for i:=0; i<size; i++ {
					_, err := db.Exec(r, s.Attr[i].Level, s.Attr[i].Name, s.Attr[i].UpdateDate, s.Attr[i].StartDate,
						s.Attr[i].EndDate, s.Attr[i].IsActive)
					if err != nil {
						//fmt.Println(err)
					}
					total = 0
				}
				break
			}
		}
	}
}
func (s *ReestrObj) Read(db *sql.DB, r string, f io.ReadCloser, size int, code string) {
	decoder := xml.NewDecoder(f)
	var inElement string
	total := 0
	t, _ := decoder.Token()
	for{
		//fmt.Println(t)
		if t == nil {
			if total != 0 {
				for i := 0; i < total; i++ {
					_, err := db.Exec(r, code, s.Attr[i].ObjectID, s.Attr[i].ObjectGUID, s.Attr[i].ChangeID,
						s.Attr[i].IsActive, s.Attr[i].LevelID, s.Attr[i].CreateDate, s.Attr[i].UpdateDate)
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
				if inElement == "OBJECT" {
					var offer ReestrObjObj
					decoder.DecodeElement(&offer, &se)
					s.Attr[total] = offer
					total++
				}
			default:
			}
			if total % Size == 0 && total != 0{
				ctx := context.Background()
				tx, err := db.BeginTx(ctx, nil)
				if err != nil {
					log.Fatal(err)
				}
				for i:=0; i<size; i++ {
					_, err := db.ExecContext(ctx,r, code, s.Attr[i].ObjectID, s.Attr[i].ObjectGUID, s.Attr[i].ChangeID,
						s.Attr[i].IsActive, s.Attr[i].LevelID, s.Attr[i].CreateDate, s.Attr[i].UpdateDate)
					if err != nil {
						//fmt.Println(err)
					}
				}
				total = 0
				err = tx.Commit()
				break
			}
		}
	}
}
