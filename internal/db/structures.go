package db

import (
	"encoding/xml"
	_ "github.com/creasty/defaults"
	"github.com/google/uuid"
)

type AddHouseTypes struct {
	XMLName xml.Name        `xml:"HOUSETYPES"`
	Attr    []HouseType 	`xml:"HOUSETYPE"`
	Request string
}
type HouseType struct{
	ID				int	   			`xml:"ID,attr" db:"id"`
	Name			string 			`xml:"NAME,attr" db:"name"`
	ShortName		string 			`xml:"SHORTNAME,attr" db:"shortname"`
	Desc			string 			`xml:"DESC,attr" db:"description"`
	IsActive		string 			`xml:"ISACTIVE,attr" db:"isactive"`
	UpdateDate		string 			`xml:"UPDATEDATE,attr" db:"updatedate"`
	StartDate		string 			`xml:"STARTDATE,attr" db:"startdate"`
	EndDate			string 			`xml:"ENDDATE,attr" db:"enddate"`
}

type AddrObj struct{
	XMLName			xml.Name		`xml:"ADDRESSOBJECTS"`
	Attr			[]Object		`xml:"OBJECT"`
	Request			string
}
type Object struct{
	XMLName    xml.Name `xml:"OBJECT"`
	ID         int      `xml:"ID,attr" db:"id"`
	ObjectID   int      `xml:"OBJECTID,attr" db:"objectid"`
	ObjectUID  uuid.UUID   `xml:"OBJECTGUID,attr" db:"objectguid" pg:"type:uuid,default:uuid_generate_v4()"`
	ChangeID   int      `xml:"CHANGEID,attr" db:"changeid"`
	Name       string   `xml:"NAME,attr" db:"name"`
	TypeName   string   `xml:"TYPENAME,attr" db:"typename"`
	Level      int      `xml:"LEVEL,attr" db:"level"`
	OperTypeID int      `xml:"OPERTYPEID,attr" db:"opertypeid"`
	PrevID     int      `xml:"PREVID,attr" db:"previd"`
	NextID     int      `xml:"NEXTID,attr" db:"nextid"`
	UpdateDate string   `xml:"UPDATEDATE,attr" db:"updatedate"`
	StartDate  string   `xml:"STARTDATE,attr" db:"startdate"`
	EndDate    string   `xml:"ENDDATE,attr" db:"enddate"`
	IsActual   string   `xml:"ISACTUAL,attr" db:"isactual"`
	IsActive   string   `xml:"ISACTIVE,attr" db:"isactive"`
	RegionCode string   `db:"regioncode"`

}

type AddrObjsDiv struct{
	XMLName			xml.Name		`xml:"ITEMS"`
	Attr		[]AddrObjsItem		`xml:"ITEM"`
	Request			string
}
type AddrObjsItem struct{
	XMLName			xml.Name		`xml:"ITEM"`
	ID				int				`xml:"ID,attr" db:"id"`
	ParentID		int				`xml:"PARENTID,attr" db:"parentid"`
	ChildID			int				`xml:"CHILDID,attr" db:"childid"`
	ChangeID		int				`xml:"CHANGEID,attr" db:"changeid"`
	RegionCode		string			`db:"regioncode"`
}

type AddrObjParams struct{
	XMLName			xml.Name		`xml:"PARAMS"`
	Attr			[]Params		`xml:"PARAM"`
	Request			string
}
type Params struct{
	XMLName			xml.Name		`xml:"PARAM"`
	ID				int				`xml:"ID,attr" db:"id"`
	ObjectID		int				`xml:"OBJECTID,attr" db:"objectid"`
	ChangeID		int				`xml:"CHANGEID,attr" db:"changeid"`
	ChangeIdEnd		int				`xml:"CHANGEIDEND,attr" db:"changeidend"`
	TypeID			int				`xml:"TYPEID,attr" db:"typeid"`
	Value			string			`xml:"VALUE,attr" db:"value"`
	UpdateDate		string			`xml:"UPDATEDATE,attr" db:"updatedate"`
	StartDate		string			`xml:"STARTDATE,attr" db:"startdate"`
	EndDate			string			`xml:"ENDDATE,attr" db:"enddate"`
	RegionCode		string			`db:"regioncode"`
}

type AddrObjTypes struct{
	XMLName			xml.Name		`xml:"ADDRESSOBJECTTYPES"`
	Attr			[]AddrObjType 	`xml:"ADDRESSOBJECTTYPE"`
	Request			string
}
type AddrObjType struct{
	XMLName			xml.Name	`xml:"ADDRESSOBJECTTYPE"`
	ID				int			`xml:"ID,attr" db:"id"`
	Level			int			`xml:"LEVEL,attr" db:"level"`
	Name			string		`xml:"NAME,attr" db:"name"`
	ShortName		string		`xml:"SHORTNAME,attr" db:"shortname"`
	Desc			string		`xml:"DESC,attr" db:"description"`
	UpdateDate		string		`xml:"UPDATEDATE,attr" db:"updatedate"`
	StartDate		string		`xml:"STARTDATE,attr" db:"startdate"`
	EndDate			string		`xml:"ENDDATE,attr" db:"enddate"`
	IsActive		string		`xml:"ISACTIVE,attr" db:"isactive"`
}

type MunHierarchy struct{
	XMLName			xml.Name		`xml:"ITEMS"`
	Attr			[]MunHierarchyItem 	`xml:"ITEM"`
	Request			string

}
type MunHierarchyItem struct{
	XMLName			xml.Name		`xml:"ITEM"`
	ID				int				`xml:"ID,attr" db:"id"`
	ObjectID		int				`xml:"OBJECTID,attr" db:"objectid"`
	ParentObjID		int				`xml:"PARENTOBJID,attr" db:"parentobjid"`
	ChangeID		int				`xml:"CHANGEID,attr" db:"changeid"`
	Oktmo			string			`xml:"OKTMO,attr" db:"oktmo"`
	PrevID			int				`xml:"PREVID,attr" db:"previd"`
	NextID			int				`xml:"NEXTID,attr" db:"nextid"`
	UpdateDate		string			`xml:"UPDATEDATE,attr" db:"updatedate"`
	StartDate		string			`xml:"STARTDATE,attr" db:"startdate"`
	EndDate			string			`xml:"ENDDATE,attr" db:"enddate"`
	IsActive		string			`xml:"ISACTIVE,attr" db:"isactive"`
	RegionCode		string			`db:"regioncode"`
}

type ChangeHist struct{
	XMLName			xml.Name		`xml:"ITEMS"`
	Attr	[]ChangeHistItem 		`xml:"ITEM"`
	Request			string
}
type ChangeHistItem struct{
	XMLName        	xml.Name		`xml:"ITEM"`
	ChangeID    	int 			`xml:"CHANGEID,attr" db:"changeid"`
	ObjectID    	int 			`xml:"OBJECTID,attr" db:"objectid"`
	AdrobjectID		string 			`xml:"ADROBJECTID,attr" db:"adrobjectid"`
	OperTypeID  	int 			`xml:"OPERTYPEID,attr" db:"opertypeid"`
	NdocID      	int				`xml:"NDOCID,attr" db:"ndocid"`
	ChangeDate		string 			`xml:"CHANGEDATE,attr" db:"changedate"`
	RegionCode		string			`db:"regioncode"`
}

type HouseTypes struct{
	XMLName   		xml.Name		`xml:"HOUSETYPES"`
	Attr			[]HouseType 	`xml:"HOUSETYPE"`
	Request			string

}

type AdmHierarchy struct{
	XMLName			xml.Name		`xml:"ITEMS"`
	Attr	[]AdmHierarchyItem 		`xml:"ITEM"`
	Request			string
}
type AdmHierarchyItem struct{
	XMLName			xml.Name		`xml:"ITEM"`
	ID				int				`xml:"ID,attr" db:"id"`
	ObjectID		int				`xml:"OBJECTID,attr" db:"objectid"`
	ParentObjID		int				`xml:"PARENTOBJID,attr" db:"parentobjid"`
	ChangeID		int				`xml:"CHANGEID,attr" db:"changeid"`
	AreaCode		int				`xml:"AREACODE,attr" db:"areacode"`
	CityCode		int				`xml:"CITYCODE,attr" db:"citycode"`
	PlaceCode		int				`xml:"PLACECODE,attr" db:"placecode"`
	PlanCode		int				`xml:"PLANCODE,attr" db:"plancode"`
	StreetCode		int				`xml:"STREETCODE,attr" db:"streetcode"`
	RegionCode		string			`xml:"REGIONCODE,attr" db:"regioncode"`
	PrevID			int				`xml:"PREVID,attr" db:"previd"`
	NextID			int				`xml:"NEXTID,attr" db:"nextid"`
	UpdateDate		string			`xml:"UPDATEDATE,attr" db:"updatedate"`
	StartDate		string			`xml:"STARTDATE,attr" db:"startdate"`
	EndDate			string			`xml:"ENDDATE,attr" db:"enddate"`
	IsActive		string			`xml:"ISACTIVE,attr" db:"isactive"`

}

type ObjectLevels struct{
	XMLName			xml.Name		`xml:"OBJECTLEVELS"`
	Attr			[]ObjectLevel 	`xml:"OBJECTLEVEL"`
	Request			string
}
type ObjectLevel struct{
	XMLName			xml.Name		`xml:"OBJECTLEVEL"`
	Level			int				`xml:"LEVEL,attr" db:"level"`
	Name			string			`xml:"NAME,attr" db:"name"`
	StartDate		string			`xml:"STARTDATE,attr" db:"startdate"`
	EndDate			string			`xml:"ENDDATE,attr" db:"enddate"`
	UpdateDate		string			`xml:"UPDATEDATE,attr" db:"updatedate"`
	IsActive		string			`xml:"ISACTIVE,attr" db:"isactive"`
}

type ReestrObj struct{
	XMLName			xml.Name		`xml:"REESTR_OBJECTS"`
	Attr			[]ReestrObjObj 	`xml:"OBJECT"`
	Request			string
}
type ReestrObjObj struct {
	XMLName			xml.Name		`xml:"OBJECT"`
	ObjectID		int				`xml:"OBJECTID,attr" db:"objectid"`
	ObjectGUID		uuid.UUID		`xml:"OBJECTGUID,attr" db:"objectguid" pg:"type:uuid,default:uuid_generate_v4()"`
	ChangeID		int				`xml:"CHANGEID,attr" db:"changeid"`
	IsActive		int				`xml:"ISACTIVE,attr" db:"isactive"`
	LevelID			int				`xml:"LEVELID,attr" db:"levelid"`
	CreateDate		string		`xml:"CREATEDATE,attr" db:"createdate"`
	UpdateDate		string		`xml:"UPDATEDATE,attr" db:"updatedate"`
	RegionCode		string			`db:"regioncode"`
}