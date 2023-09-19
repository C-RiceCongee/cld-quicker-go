// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNameYkFoodnutrition = "yk_foodnutrition"

// YkFoodnutrition mapped from table <yk_foodnutrition>
type YkFoodnutrition struct {
	ID           int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Docv         string  `gorm:"column:docv;not null" json:"docv"`
	Leixing      int32   `gorm:"column:leixing" json:"leixing"`
	Name         string  `gorm:"column:name;not null" json:"name"`
	Xuhao        string  `gorm:"column:xuhao" json:"xuhao"`
	Shibu        int32   `gorm:"column:shibu" json:"shibu"`
	Clcnengliang int32   `gorm:"column:clcnengliang" json:"clcnengliang"`
	Nengliang    int32   `gorm:"column:nengliang" json:"nengliang"`
	Shuifen      float64 `gorm:"column:shuifen" json:"shuifen"`
	Danbai       float64 `gorm:"column:danbai" json:"danbai"`
	Zhifang      float64 `gorm:"column:zhifang" json:"zhifang"`
	Tanshui      float64 `gorm:"column:tanshui" json:"tanshui"`
	Xianwei      float64 `gorm:"column:xianwei" json:"xianwei"`
	Huifen       float64 `gorm:"column:huifen" json:"huifen"`
	VA           float64 `gorm:"column:VA" json:"VA"`
	B1           float64 `gorm:"column:B1" json:"B1"`
	B2           float64 `gorm:"column:B2" json:"B2"`
	B3           float64 `gorm:"column:B3" json:"B3"`
	VC           float64 `gorm:"column:VC" json:"VC"`
	VE           float64 `gorm:"column:VE" json:"VE"`
	Jia          float64 `gorm:"column:jia" json:"jia"`
	Na           float64 `gorm:"column:na" json:"na"`
	Gai          float64 `gorm:"column:gai" json:"gai"`
	Mei          float64 `gorm:"column:mei" json:"mei"`
	Tie          float64 `gorm:"column:tie" json:"tie"`
	Meng         float64 `gorm:"column:meng" json:"meng"`
	Xin          float64 `gorm:"column:xin" json:"xin"`
	Tong         float64 `gorm:"column:tong" json:"tong"`
	Lin          float64 `gorm:"column:lin" json:"lin"`
	Xi           float64 `gorm:"column:xi" json:"xi"`
	Leibie       int32   `gorm:"column:leibie" json:"leibie"`
	Lei          int32   `gorm:"column:lei" json:"lei"`
	Danguchun    float64 `gorm:"column:danguchun" json:"danguchun"`
}

// TableName YkFoodnutrition's table name
func (*YkFoodnutrition) TableName() string {
	return TableNameYkFoodnutrition
}
