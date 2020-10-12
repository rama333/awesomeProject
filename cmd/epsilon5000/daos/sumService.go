package daos

import (
	"awesomeProject/cmd/epsilon5000/models"
	"database/sql"
	"fmt"
	_ "gopkg.in/goracle.v2"
	"log"
)

type sumServiceDAO struct{}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("goracle", "@(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=tcp)(HOST=)(PORT=1521)))(CONNECT_DATA=(SERVICE_NAME=)))")

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

}

func NewSumServiceDAO() *sumServiceDAO {

	return &sumServiceDAO{}
}

func (dao *sumServiceDAO) Get(inc string) ([]models.ServicesCount, error) {

	q := fmt.Sprintf("select t.unit_ip, sum(case when t.TK_TYPE= 122 then 1 else 0 end) iptv, sum(case when t.TK_TYPE in (2, 82, 42, 142, 282, 302,374,375, 462, 323) then 1 else 0 end) spd, sum(case when t.TK_TYPE= 202 then 1 else 0 end) sip, m_ttk.aod_rsc.GET_ADDRESSF(o.obj_adr) from m_ttk.TKSERVICES t, m_ttk.rm_object o where unit_ip in ('%s') and t.OBJECTID=o.obj_id group by t.unit_ip, o.obj_adr", inc)

	value, err := db.Query(q)

	defer value.Close()

	list := []models.ServicesCount{}

	for value.Next() {

		var ser models.ServicesCount

		err := value.Scan(&ser.Ip, &ser.Iptv, &ser.Spd, &ser.Sip, &ser.Addr)

		if err != nil {
			return nil, err
		}

		list = append(list, ser)
		fmt.Println(ser)
	}

	return list, err

}
