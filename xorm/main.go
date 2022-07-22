package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	xtime "time"
	"xorm.io/xorm"
)

var engine *xorm.Engine

type Goods struct {
	Id int64 `json:"id"`
	SkuSn string `json:"sku_sn"`
	Barcode string `json:"barcode"`
	CreatedAt JsonTime `json:"created_at" xorm:"created"`
}

func (g *Goods) Table()string  {
	return "goods"
}
type JsonTime int64
func (j JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"`+xtime.Unix(int64(j),0).Format("2006-01-02 15:04:05")+`"`), nil
}
func main() {
	e := gin.Default()
	e.GET("/", func(c *gin.Context) {
		var err error
		engine, err = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/local_test?charset=utf8mb4&interpolateParams=true")
		if err != nil {
			log.Fatal(err)
		}
		var goods []Goods
		err = engine.Table(new(Goods).Table()).Find(&goods)
		if err!=nil {
			log.Fatal(err)
		}
		/*for _,v:=range goods{
			bytes, _ := v.CreatedAt.MarshalJSON()
			fmt.Printf("id=%v\tsku_sn=%v\tbarcode=%v\tcreated_at=%v\n",v.Id,v.SkuSn,v.Barcode,string(bytes))
		}*/
		c.JSON(http.StatusOK,gin.H{"goods":goods})
	})
	log.Fatal(e.Run(":8090"))
}
