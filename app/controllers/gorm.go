/* gorm.go
 *
 * Copyright (C) 2015-2016 Helium Ink
 * Author : Milind Deore
 *
 * This software may be modified and distributed under the terms
 * of the MIT license.  See the LICENSE file for details.
 */

package controllers

import (
    "github.com/jinzhu/gorm"
    "database/sql"
    "github.com/revel/revel"
)

var (
    Dbm *gorm.DB
)

type GormController struct {
    *revel.Controller
    Txn *gorm.DB
}

// transactions

// This method fills the c.Txn before each transaction
func (c *GormController) Begin() revel.Result {
    txn := Dbm.Begin()
    if txn.Error != nil {
        panic(txn.Error)
    }
    c.Txn = txn
    return nil
}

// This method clears the c.Txn after each transaction
func (c *GormController) Commit() revel.Result {
    if c.Txn == nil {
        return nil
    }
    c.Txn.Commit()
    if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
        panic(err)
    }
    c.Txn = nil
    return nil
}

// This method clears the c.Txn after each transaction, too
func (c *GormController) Rollback() revel.Result {
    if c.Txn == nil {
        return nil
    }
    c.Txn.Rollback()
    if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
        panic(err)
    }
    c.Txn = nil
    return nil
}

