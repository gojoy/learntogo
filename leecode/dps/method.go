package dps
/*
import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
	"strings"
	"net/url"
	"bytes"
	"io"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


var(
	v=url.Values{}
)

func encodeData(data interface{}) (*bytes.Buffer,error) {
	params:=bytes.NewBuffer(nil)
	if data != nil {
		if err:=json.NewEncoder(params).Encode(data);err!=nil{
			return nil,err
		}
	}
	return params,nil
}

func call(method,path string,data interface{}, headers map[string][]string)  (io.ReadCloser, int, error) {

	params, err := encodeData(nil)

	if err != nil {
		return nil, -1, err
	}
	body, _, statusCode, err := clientRequest(method, path, params)
	return body,statusCode,err
}

func ps() error {
	fmt.Println("begin ps")

	rdr, _, err := call("GET", "/containers/json?"+v.Encode(), nil,nil)
	if err != nil {
		fmt.Printf("in func ps call err %v\n",err )
		return err
	}


	stripNamePrefix := func(ss []string) []string {
		for i, s := range ss {
			ss[i] = s[1:]
		}
		return ss
	}

	containers:=[]Container{}
	if err := json.NewDecoder(rdr).Decode(&containers); err != nil {
		return err
	}

	//w := tabwriter.NewWriter(os.Stdout, 20, 1, 3, ' ', 0)
	fmt.Fprint(os.Stdout, "CONTAINER ID\tIMAGE\tCOMMAND\tCREATED\tSTATUS\tPORTS\tNAMES")

	for _,container:=range containers{
		ID:=container.ID
		image:=container.Image
		port:="port"
		var(
			names   = stripNamePrefix(container.Names)
			command = strconv.Quote(container.Command)
		)
		fmt.Fprintf(os.Stdout, "%s\t%s\t%s\t%s ago\t%s\t%s\t%s\t", ID, image, command,
			HumanDuration(time.Now().UTC().Sub(time.Unix(int64(container.Created), 0))),
			container.Status, port, strings.Join(names, ","))

	}
	return nil
}



func testmysql()  {
	fmt.Printf("in main begin\n")
	conn,err:=sql.Open("mysql","gg:123456@tcp(123.207.89.211:3306)/ha")
	defer conn.Close()
	if err != nil {
		fmt.Printf("con err is %v\n",err)
		log.Println(err)
		return
	}
	fmt.Printf("after open\n")
	res,err:=conn.Query("select * from st")
	defer res.Close()
	if err != nil {
		panic(err)
	}
	/*
	con,err:=res.Columns()
	if err != nil {
		panic(err)
	}
	fmt.Printf("col is %v\n",con)

	var user interface{}
	var id interface{}
	for res.Next(){
		err:=res.Scan(&id,&user)
		if err!=nil {
			fmt.Printf("scan err %v\n",err)
		}
		//fmt.Printf("is id %v\n",id)
		fmt.Printf("id is %v \t user is %v\n",id,user)
	}
	if err := res.Err(); err != nil {
		log.Fatal(err)
	}
}
*/