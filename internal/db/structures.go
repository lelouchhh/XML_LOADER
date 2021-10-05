package db

import (
	"encoding/xml"
	_ "github.com/creasty/defaults"
)

type AddHouseTypes struct {
	XMLName xml.Name        `xml:"HOUSETYPES"`
	Attr    [Size+1]HouseType 	`xml:"HOUSETYPE"`
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
	Attr			[Size+1]Object		`xml:"OBJECT"`
	Request			string
	RegionCode		string
}
type Object struct{
	XMLName			xml.Name		`xml:"OBJECT"`
	ID				int				`xml:"ID,attr" db:"id"`
	ObjectID		int			`xml:"OBJECTID,attr" db:"objectid"`
	ObjectUID		string			`xml:"OBJECTGUID,attr" db:"objectguid"`
	ChangeID		int			`xml:"CHANGEID,attr" db:"changeid"`
	Name			string			`xml:"NAME,attr" db:"name"`
	TypeName		string			`xml:"TYPENAME,attr" db:"typename"`
	Level			int			`xml:"LEVEL,attr" db:"level"`
	OperTypeID		int			`xml:"OPERTYPEID,attr" db:"opertypeid"`
	prevID			int			`xml:"PREVID,attr" db:"previd"`
	nextID			int			`xml:"NEXTID,attr" db:"nextid"`
	UpdateDate		string			`xml:"UPDATEDATE,attr" db:"updatedate"`
	StartDate		string			`xml:"STARTDATE,attr" db:"startdate"`
	EndDate			string			`xml:"ENDDATE,attr" db:"enddate"`
	IsActual		string			`xml:"ISACTUAL,attr" db:"isactual"`
	IsActive		string			`xml:"ISACTIVE,attr" db:"isactive"`
}

type AddrObjsDiv struct{
	XMLName			xml.Name		`xml:"ITEMS"`
	Attr		[Size+1]AddrObjsItem		`xml:"ITEM"`
	Request			string
	RegionCode		string
}
type AddrObjsItem struct{
	XMLName			xml.Name		`xml:"ITEM"`
	ID				int				`xml:"ID,attr" db:"id"`
	ParentID		int				`xml:"PARENTID,attr" db:"parentid"`
	ChildID			int				`xml:"CHILDID,attr" db:"childid"`
	ChangeID		int				`xml:"CHANGEID,attr" db:"changeid"`
}

type AddrObjParams struct{
	XMLName			xml.Name		`xml:"PARAMS"`
	Attr			[Size+1]Params		`xml:"PARAM"`
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
	Attr			[Size+1]AddrObjType 	`xml:"ADDRESSOBJECTTYPE"`
	Request			string
}
type AddrObjType struct{
	XMLName			xml.Name	`xml:"ADDRESSOBJECTTYPE"`
	ID				int			`xml:"ID,attr" db:"id"`
	Level			int			`xml:"LEVEL,attr" db:"level"`
	Name			string		`xml:"NAME,attr" db:"name"`
	ShortName		string		`xml:"SHORTNAME,attr" db:"shortname"`
	Desc			string		`xml:"DESC,attr" db:"_desc"`
	UpdateDate		string		`xml:"UPDATEDATE,attr" db:"updatedate"`
	StartDate		string		`xml:"STARTDATE,attr" db:"startdate"`
	EndDate			string		`xml:"ENDDATE,attr" db:"enddate"`
	IsActive		string		`xml:"ISACTIVE,attr" db:"isactive"`
}

type MunHierarchy struct{
	XMLName			xml.Name		`xml:"ITEMS"`
	Attr			[Size+1]MunHierarchyItem 	`xml:"ITEM"`
	Request			string
	RegionCode		string			`db:"regioncode"`

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
}

type ChangeHist struct{
	XMLName			xml.Name		`xml:"ITEMS"`
	Attr	[Size+1]ChangeHistItem 		`xml:"ITEM"`
	Request			string
	RegionCode		string			`db:"regioncode"`
}
type ChangeHistItem struct{
	XMLName        	xml.Name		`xml:"ITEM"`
	ChangeID    	int 			`xml:"CHANGEID,attr" db:"changeid"`
	ObjectID    	int 			`xml:"OBJECTID,attr" db:"objectid"`
	AdrobjectID		string 			`xml:"ADROBJECTID,attr" db:"adrobjectid"`
	OperTypeID  	int 			`xml:"OPERTYPEID,attr" db:"opertypeid"`
	NdocID      	int				`xml:"NDOCID,attr" db:"ndocid"`
	ChangeDate		string 			`xml:"CHANGEDATE,attr" db:"changedate"`

}

type HouseTypes struct{
	XMLName   		xml.Name		`xml:"HOUSETYPES"`
	Attr			[4]HouseType 	`xml:"HOUSETYPE"`
	Request			string

}

type AdmHierarchy struct{
	XMLName			xml.Name		`xml:"ITEMS"`
	Attr	[Size+1]AdmHierarchyItem 		`xml:"ITEM"`
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
	Attr			[Size+1]ObjectLevel 	`xml:"OBJECTLEVEL"`
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
	Attr			[Size+1]ReestrObjObj 	`xml:"OBJECT"`
	Request			string
	RegionCode		string			`db:"regioncode"`
}
type ReestrObjObj struct {
	XMLName			xml.Name		`xml:"OBJECT"`
	ObjectID		int				`xml:"OBJECTID,attr" db:"objectid"`
	ObjectGUID		string			`xml:"OBJECTGUID,attr" db:"objectguid"`
	ChangeID		string				`xml:"CHANGEID,attr" db:"changeid"`
	IsActive		string			`xml:"ISACTIVE,attr" db:"isactive"`
	LevelID			string				`xml:"LEVELID,attr" db:"levelid"`
	CreateDate		string			`xml:"CREATEDATE,attr" db:"createdate"`
	UpdateDate		string			`xml:"UPDATEDATE,attr" db:"updatedate"`
}