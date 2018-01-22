package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    _"time"
)
/*
    create table t_gotest(
        id int primary key auto_increment,
        name varchar(20) not null,
        ts timestamp
    );
*/

func main() {

    db, err := sql.Open("mysql", "huya_order_query:6iU8Aw@173@tcp(10.26.1.18:6302)/db_item_order?charset=utf8")
    if err != nil {
        log.Fatalf("Open database error: %s\n", err)
    }

    defer db.Close()
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    var pid int64
    var paytotal int 
   /*
    var tableName = [7]string {"presentertotal_20170201", "presentertotal_20170202","presentertotal_20170203",
                            "presentertotal_20170204","presentertotal_20170205","presentertotal_20170206","presentertotal_20170207"}
                            */
    
    var dayTotal map[int64]int 
    dayTotal = make(map[int64]int)


    //var n = len(tableName) 
    //for  i := 0 ; i< n; i++ {

        //log.Println(tableName[i])

        rows, err := db.Query("select presenter_uid, pay_total from presentertotal_20170201 where pay_total> 100000")

        //stmtOut, err := db.Prepare("select presenter_uid, pay_total from ? where pay_total>?")
        if err != nil {
            log.Println(err)
        }

        for rows.Next() {

            pid = 0;
            paytotal = 0;
            err := rows.Scan(&pid, &paytotal)
            if err != nil {
                log.Fatal(err)
            }

            paytotal = paytotal / 100 
            dayTotal[pid] +=paytotal 
            //log.Println(pid, paytotal)

            err = rows.Err()
            if err != nil {
                log.Fatal(err)
            }

            rows.Close()
        }

        rows2, err := db.Query("select presenter_uid, pay_total from presentertotal_20170202 where pay_total> 100000")

        for rows2.Next() {

            pid = 0;
            paytotal = 0;
            err := rows2.Scan(&pid, &paytotal)
            if err != nil {
                log.Fatal(err)
            }

            paytotal = paytotal / 100 
            dayTotal[pid] +=paytotal 
            //log.Println(pid, paytotal)

            err = rows2.Err()
            if err != nil {
                log.Fatal(err)
            }

            rows2.Close()
        }

        rows3, err:= db.Query("select presenter_uid, pay_total from presentertotal_20170203 where pay_total> 100000")
        
        for rows3.Next() {

            pid = 0;
            paytotal = 0;
            err := rows3.Scan(&pid, &paytotal)
            if (err != nil) {
                log.Fatal(err)
            }

            paytotal = paytotal / 100 
            dayTotal[pid] +=paytotal 
            //log.Println(pid, paytotal)

            err = rows.Err()
            if err != nil {
                log.Fatal(err)
            }

            rows3.Close()
        }

        rows4, err := db.Query("select presenter_uid, pay_total from presentertotal_20170204 where pay_total> 100000")
        
        for rows4.Next() {

            pid = 0;
            paytotal = 0;
            err := rows4.Scan(&pid, &paytotal)
            if (err != nil) {
                log.Fatal(err)
            }

            paytotal = paytotal / 100 
            dayTotal[pid] +=paytotal 
            //log.Println(pid, paytotal)

            err = rows4.Err()
            if err != nil {
                log.Fatal(err)
            }

            rows4.Close()
        }

        rows5, err := db.Query("select presenter_uid, pay_total from presentertotal_20170205 where pay_total> 100000")
        
        for rows5.Next() {

            pid = 0;
            paytotal = 0;
            err := rows5.Scan(&pid, &paytotal)
            if err != nil {
                log.Fatal(err)
            }

            paytotal = paytotal / 100 
            dayTotal[pid] +=paytotal 
            //log.Println(pid, paytotal)

            err = rows5.Err()
            if err != nil {
                log.Fatal(err)
            }

            rows5.Close()
        }

        rows6, err := db.Query("select presenter_uid, pay_total from presentertotal_20170206 where pay_total> 100000")
        
        for rows6.Next() {

            pid = 0;
            paytotal = 0;
            err6 := rows6.Scan(&pid, &paytotal)
            if err6 != nil {
                log.Fatal(err6)
            }

            paytotal = paytotal / 100 
            dayTotal[pid] +=paytotal 
            //log.Println(pid, paytotal)

            err = rows6.Err()
            if err != nil {
                log.Fatal(err)
            }

            rows6.Close()
        }

        rows7, err := db.Query("select presenter_uid, pay_total from presentertotal_20170207 where pay_total> 100000")
        
        for rows7.Next() {

            pid = 0;
            paytotal = 0;
            err := rows7.Scan(&pid, &paytotal)
            if err != nil {
                log.Fatal(err)
            }

            paytotal = paytotal / 100 
            dayTotal[pid] +=paytotal 
            //log.Println(pid, paytotal)

            err = rows7.Err()
            if err != nil {
                log.Fatal(err)
            }

            rows7.Close()
        }


    //}
      

    for key, value := range  dayTotal { 
        log.Println(key, value)
    }

}