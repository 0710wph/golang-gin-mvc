package controllers
import (
	"fmt"
	"gin/models"
	"github.com/gin-gonic/gin"
	//"log"
	"net/http"
	"strconv"
)

func  BdlistPost(c *gin.Context) {
	city_id := c.PostForm("city_id")
	utype := c.PostForm("type")
	ubdtype := c.PostForm("bd_type")
	ushi := c.PostForm("shi")
	page := c.PostForm("page")

	onpage, _ := strconv.ParseInt(page, 0, 64)

	uarea_id := c.PostForm("area_id")
	ustreet_id := c.PostForm("street_id")

	wstr := " tmpbd.city_id="+city_id;

	if(utype != ""){
		wstr = wstr + " and tmpbd.type = "+utype
	}

	if(uarea_id != ""){
		wstr = wstr + " and tmpbd.area_id = "+uarea_id
	}

	if(ustreet_id != ""){
		wstr = wstr + " and tmpbd.street = "+ustreet_id
	}

	if(ubdtype != ""){
		wstr = wstr + " and tmpbd.bd_type = "+ubdtype
	}

	if(ushi != ""){
		wstr = wstr + " and tmpbd.shi = "+ushi
	}



	numoffset := (onpage-1) * 10

	sqlstrn := "select a.id,b.title,b.type as types,b.bd_type as bdtypes,a.bddate,c.data_from,d.name,e.street from ( select aa.id,aa.bddate,aa.area_id,aa.street from (select tmpbd.id,tmpbd.date as bddate,tmpbd.area_id,tmpbd.street from fcxlt_bd as tmpbd where %s ) aa order by aa.bddate desc limit %d,10) a left join fcxlt_bd as b on a.id=b.id LEFT JOIN fcxlt_bd_from c ON a.id = c.id LEFT JOIN fcxlt_areas as d on a.area_id=d.id LEFT JOIN fcxlt_area_street as e on a.street=e.id"
	sqlstrn = fmt.Sprintf(sqlstrn, wstr, numoffset)

	bds := models.QueryBdlist(sqlstrn)
	fmt.Println(bds)
	c.JSON(http.StatusOK, gin.H{
		"data": bds,
	})
}