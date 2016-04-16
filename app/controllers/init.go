/* init.go
 *
 * Copyright (C) 2015-2016 Helium Ink
 * Author : Milind Deore
 *
 * This software may be modified and distributed under the terms
 * of the MIT license.  See the LICENSE file for details.
 */
 
package controllers

import (
    "github.com/revel/revel"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "strings"
    "RepSys/app/models"
)


func init(){
    revel.OnAppStart(InitDb)
    revel.InterceptMethod((*GormController).Begin, revel.BEFORE)
    revel.InterceptMethod((*GormController).Commit, revel.AFTER)
    revel.InterceptMethod((*GormController).Rollback, revel.FINALLY)
}

/* 
 * Get params from confgiration. 
 */
func getParamString(param string, defaultValue string) string {
    p, found := revel.Config.String(param)
    if !found {
        if defaultValue == "" {
            revel.ERROR.Fatal("Cound not find parameter: " + param)
        } else {
            return defaultValue
        }
    }
    return p
}

/* 
 * Structure DB connection string. 
 */
func getConnectionString() string {
    host := getParamString("db.host", "")
    port := getParamString("db.port", "3306")
    user := getParamString("db.user", "")
    pass := getParamString("db.password", "")
    dbname := getParamString("db.name", "heliumink")
    protocol := getParamString("db.protocol", "tcp")
    //dbargs := getParamString("dbargs", "parseTime=true&tls=true&charset=utf8")
    dbargs := getParamString("dbargs", "parseTime=true&charset=utf8")

    if strings.Trim(dbargs, " ") != "" {
        dbargs = "?" + dbargs
    } else {
        dbargs = ""
    }
    return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s", 
        user, pass, protocol, host, port, dbname, dbargs)
}

/* 
 * Init DB
 */
var InitDb func() = func(){
    connectionString := getConnectionString()
    dbm, err := gorm.Open("mysql", connectionString)
    if err != nil {
        revel.ERROR.Fatal(err)
    }
    Dbm = dbm
    // Ping function checks the database connectivity
    if err := Dbm.DB().Ping(); err != nil {
	    revel.ERROR.Fatal(err)
    }
    // Max Idle connections - protection code.
    Dbm.DB().SetMaxIdleConns(10)
    // Max Open connections - protection code. 
    Dbm.DB().SetMaxOpenConns(100)
    // Force GORM to create 'table' and not 'tables'
    Dbm.SingularTable(true)
    
    //Dbm.Set("gorm:table_options", "EXTRA=none").CreateTable(&models.UsersTbl{})
    Dbm.DropTable(&models.User{})
    Dbm.DropTable(&models.CatBadge{})
    Dbm.DropTable(&models.LeaderBoard{})
    
    
    Dbm.CreateTable(&models.User{})
    Dbm.CreateTable(&models.CatBadge{})
    Dbm.CreateTable(&models.LeaderBoard{})
    //Dbm.CreateTable(&models.LeaderBoard{})
   
    
    
}
