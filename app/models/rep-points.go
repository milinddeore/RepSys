/* rep-points.go
 *
 * Copyright (C) 2015-2016 Helium Ink
 * Author : Milind Deore
 *
 * This software may be modified and distributed under the terms
 * of the MIT license.  See the LICENSE file for details.
 */
 
package models

import
(
    "time"
    "regexp"
    "github.com/revel/revel"
)

/* User Table - Add 'new' user */
type User struct {
    UserId          uint32     `sql:"type:bigint PRIMARY KEY" json:"UserId"` //Hack: God knows why the type and PK together disables auto_increment.
    Name            string     `sql:"type:varchar(25)" json:"Name"`
    CreatedAt       time.Time   
    UpdatedAt       time.Time
    DeletedAt       *time.Time
}

/* 
 * Category Table : Github, Blog, Events, etc  - Populated by Helium Ink 
 * Badges will be populated by Helium Ink based on badge rules
 */
type CatBadge struct {
    CatBadgeId      uint32      `gorm:"primary_key"      json:"CatBadgeId"`     
    CatName         string      `sql:"type:varchar(25)"  json:"CatName"`
    BadgeName       string      `sql:"type:varchar(25)"  json:"BadgeName"`
    CatDesc         string      `sql:"type:varchar(200)" json:"CatDesc"`
    BadgeDesc       string      `sql:"type:varchar(200)" json:"BadgeDesc"`
    Credits         int64       // Badge credits
    CreatedAt       time.Time   
    UpdatedAt       time.Time
    DeletedAt       *time.Time
}

/* Overall Leader Board - 'update' user */
type LeaderBoard struct {
    LeaderBoard         uint32         `gorm:"primary_key"       json:"LeaderBoard"`
    UserId              uint32                                  `json:"UserId"`
    CatName             string         `sql:"type:varchar(25)"   json:"CatName"`
    CatBadgeId          uint32                                  `json:"CatBadgeId"`
    CreditPoints        int64          // User's credits.
    CreatedAt           time.Time   
    UpdatedAt           time.Time
    DeletedAt           *time.Time     
}


/* 
 * Data Validation logic 
 */

func (b *User) Validate(v *revel.Validation) {
    
    v.Check(b.Name,
        revel.ValidRequired(),
        revel.ValidMaxSize(25))
        
}

func (b *CatBadge) Validate(v *revel.Validation) {

    v.Check(b.BadgeName,
        revel.ValidRequired(),
        revel.ValidMaxSize(25))
        
    v.Check(b.CatName,
        revel.ValidRequired(),
        revel.ValidMaxSize(25))

    v.Check(b.BadgeDesc,
        revel.ValidRequired(),
        revel.ValidMaxSize(200))    
    
    v.Check(b.CatDesc,
        revel.ValidRequired(),
        revel.ValidMaxSize(200),
        revel.ValidMatch(
            regexp.MustCompile(
                "^(Github|Blog)$")))        
}

