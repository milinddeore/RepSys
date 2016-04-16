package tests

import (
    "github.com/revel/revel/testing"
)


/* 
 * Various testing commands - Manual CLIs.
 * POST (Add) : curl -iv -H 'Content-Type: application/json' -d '{"UserId":34,"Name":"Milind"}' http://127.0.0.1:9000/GitHub
 * PUT (update) : curl -iv -X PUT -H 'Content-Type: application/json' -d '{"description":"Good job"}' http://127.0.0.1:9000/GitHub/Silver/11
 * Get (id) : curl -iv http://127.0.0.1:9000/GitHub/12
 * Get (Top List) : curl -iv http://127.0.0.1:9000/list/github/
 * Delete : curl -i -X DELETE http://127.0.0.1:9000/admin/35 
 */

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
  
  /*  
    // User Added.
    Dbm.Debug().Save(&models.User{
        Name: "Milind",
        UserId: 12,
    })
   
    // Static Configuration - GitHUb
    Dbm.Debug().Save(&models.CatBadge{
        CatName: "GitHub",
        BadgeName: "Gold",
        CatDesc: "This is GITHUB",
        BadgeDesc: "This is the top most badge",
        Credits: 10,
    })
    
    Dbm.Debug().Save(&models.CatBadge{
        CatName: "GitHub",
        BadgeName: "Silver",
        CatDesc: "This is GITHUB",
        BadgeDesc: "This is the 2nd top badge",
        Credits: 5,
    })
    
    Dbm.Debug().Save(&models.CatBadge{
        CatName: "GitHub",
        BadgeName: "Bronze",
        CatDesc: "This is GITHUB",
        BadgeDesc: "This is the 3rd top badge",
        Credits: 1,
    })
    
    // Static Configuration - Blog
    Dbm.Debug().Save(&models.CatBadge{
        CatName: "Blog",
        BadgeName: "Good",
        CatDesc: "This is Blog",
        BadgeDesc: "This is the good article",
        Credits: 10,
    })
    
    Dbm.Debug().Save(&models.CatBadge{
        CatName: "Blog",
        BadgeName: "Bad",
        CatDesc: "This is Blog",
        BadgeDesc: "This is bad!",
        Credits: -5,
    })
    
    Dbm.Debug().Save(&models.CatBadge{
        CatName: "Blog",
        BadgeName: "Ugly",
        CatDesc: "This is Blog",
        BadgeDesc: "This is the worst blog",
        Credits: -10,
    })
    
    
    
   // User query.
    u := models.User{}
    Dbm.Debug().Where("user_id = ?", "12").First(&u)   
    //Dbm.Debug().First(&u, "12")   
    
    c := models.CatBadge{}
    Dbm.Debug().Where("cat_name = ? AND badge_name = ?", "GitHub", "Silver").First(&c)
    fmt.Printf("Value of Category %d\n", c.CatBadgeId)
   
    
    if (u.UserId != 0) {
        Dbm.Debug().Save(&models.LeaderBoard{
            UserId: u.UserId,
            CatName: c.CatName,
            CatBadgeId: c.CatBadgeId,
            CreditPoints: c.Credits,
        })
    }
    fmt.Println(u) 
    */
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) After() {
	println("Tear down")
}
