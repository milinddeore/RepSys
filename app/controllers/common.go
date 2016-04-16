/* common.go
 *
 * Copyright (C) 2015-2016 Helium Ink
 * Author : Milind Deore
 *
 * This software may be modified and distributed under the terms
 * of the MIT license.  See the LICENSE file for details.
 */
 
package controllers


import (
    "strconv"
    "RepSys/app/models"
    "github.com/revel/revel"
    "encoding/json"
)

type RepSysCtrl struct {
    GormController
}


/************************* 
 * Utility functions.
 *************************/
 
func (c RepSysCtrl) parseUser() (models.User, error) {
    user := models.User{}
    decoder := json.NewDecoder(c.Request.Body)
    err := decoder.Decode(&user)
    return user, err
}

/* 
 * For `ParseUint`, the `0` means infer the base from
 * the string. `64` requires that the result fit in 64 bits.
 */
func parseUintOrDefault(intStr string, _default uint64) uint64 {
    if value, err := strconv.ParseUint(intStr, 0, 64); err != nil {
        return _default
    } else {
        return value
    }
}

/* 
 * For `ParseInt`, the `0` means infer the base from
 * the string. `64` requires that the result fit in 64 bits.
 */
func parseIntOrDefault(intStr string, _default int64) int64 {
    if value, err := strconv.ParseInt(intStr, 0, 64); err != nil {
        return _default
    } else {
        return value
    }
}



/*************************
 * REST APIs
 *************************/
 
/* 
 * Add - POST  /:category
 * Access - Only Admin or trusted peers.
 */
func (c RepSysCtrl) Add() revel.Result {
    if user, err := c.parseUser(); err != nil {
        return c.RenderText("Unable to parse the User from JSON.")
    } else {
        // Validate data
        if user.UserId <= 0 {
            return c.RenderText("Invalid UserId!")
        }
        user.Validate(c.Validation)
        if c.Validation.HasErrors() {
            return c.RenderText("You have error in your User.")
        } else {
            // Finally, insert the data.
            if err := Dbm.Debug().Create(&user).Error; err != nil {
                return c.RenderJson(err)
            } else {
                return c.RenderJson(user)
            }
        }
    }
}

/*
 * Struct to store the calculated data.
 * Note: Make sure the field name has to be same as "model's" field name. 
 */
type Result struct {
    UserId              uint64
    CatName             string
    CreditPoints        int64           // Keep it integer, as the credit points can be negative
}
/* 
 * Get - GET  /:category/:id
 * Access - Anyone
 */
func (c RepSysCtrl) Get(category string, id uint64) revel.Result {
    
    result := Result{}
    err := Dbm.Debug().Table("leader_board").Select("leader_board.user_id, cat_badge.cat_name, sum(leader_board.credit_points) as credit_points").Joins("inner join cat_badge on leader_board.cat_badge_id = cat_badge.cat_badge_id").Group("leader_board.user_id, cat_badge.cat_name").Where("leader_board.user_id = ? AND cat_badge.cat_name = ?", id, category).Scan(&result).Error
    if err != nil {
        return c.RenderJson(err)
    }
   
    return c.RenderJson(result)
}


/*
 * Update - PUT  /:category/:etype/:id
 * Access - only Admin or trusted peers.
 */
func (c RepSysCtrl) Update(category string, etype string, id uint64) revel.Result {
        
    // Lookup category and badge strings for validation then insert userID and update the requied badge.
    u := models.User{}
    if err := Dbm.Debug().Where("user_id = ?", id).First(&u).Error; err != nil {
        return c.RenderJson(err)
    }   
    if u.UserId != 0 {
        cb := models.CatBadge{}
        if err:= Dbm.Debug().Where("cat_name = ? AND badge_name = ?", category, etype).First(&cb).Error; err != nil {
            return c.RenderJson(err)
        }
        // Validation 
        if cb.CatBadgeId == 0 {
            return c.RenderText("Invalid Category or Badge!")
        }
        
        lb := models.LeaderBoard{}
        Dbm.Debug().Where("user_id = ? AND cat_badge_id = ?", id, cb.CatBadgeId).First(&lb)
        
        lb.UserId = u.UserId
        lb.CatName = cb.CatName
        lb.CatBadgeId = cb.CatBadgeId
        lb.CreditPoints += cb.Credits
        
        if err := Dbm.Debug().Save(&lb).Error; err != nil {
            return c.RenderJson(err)
        }
        return c.RenderJson(lb)
    } else {
        return c.RenderText("Unable to update user's leader board, Invalid user!")
    }
}

/*
 * Delete - DELETE  /:category/:id
 * Access - Only Admin or trusted peers.
 * NOTE: we always do soft delete. 
 */
func (c RepSysCtrl) Delete(id uint32) revel.Result {
    u := models.User{}
    if err := Dbm.Debug().Where("user_id = ?", id).Delete(&u).Error; err != nil {
        return c.RenderJson(err)        
    }
    
    return c.RenderText("Deleted user %v", id)
}


/* 
 * List - GET  /:category
 * Access - Anyone
 */
func (c RepSysCtrl) List(category string) revel.Result {
    
    result := []Result{}
    // List of all users under specific category.
    err := Dbm.Debug().Table("leader_board").Select("leader_board.user_id, cat_badge.cat_name, sum(leader_board.credit_points) as credit_points").Joins("inner join cat_badge on leader_board.cat_badge_id = cat_badge.cat_badge_id").Group("leader_board.user_id, cat_badge.cat_name").Where("cat_badge.cat_name = ?", category).Scan(&result).Error
    if err != nil {
        return c.RenderJson(err)
    }
    
    return c.RenderJson(result)
}


/* 
 * List - GET  /:top-users-category 
 * Access - Anyone
 */
func (c RepSysCtrl) ListOfTopUsers(category string, cnt int) revel.Result {
    
    result := []Result{}
    
    // List of Top 10 reputations under a specific category.
    err := Dbm.Debug().Table("leader_board").Select("leader_board.user_id, cat_badge.cat_name, sum(leader_board.credit_points) as credit_points").Joins("inner join cat_badge on leader_board.cat_badge_id = cat_badge.cat_badge_id").Group("leader_board.user_id, cat_badge.cat_name").Where("cat_badge.cat_name = ?", category).Limit(cnt).Order("sum(leader_board.credit_points) desc").Scan(&result).Error
    if err != nil {
        return c.RenderJson(err)
    }
    
    //return c.RenderText("LIST is excuted")
    return c.RenderJson(result)
}
