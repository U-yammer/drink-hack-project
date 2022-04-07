package models

import (
    "log"
    "math/rand"
)

type Drink struct {
    ID        int
    UserID    int
    DName     string
    Amount    int
    Total     int
    UserDrink DInfo
}

type DInfo struct {
    DName        string // Drink name
    DDescription string // Drink description
    AoSugar      int    // Amount of sugar
}

var d1 DInfo = DInfo{"water", "Drinking 2L of water a day is good for your health", 0}
var d2 DInfo = DInfo{"tea", "Be careful not to take too much caffeine", 0}
var d3 DInfo = DInfo{"cola", "Cola contains sugar equivalent to 16 sugar cubes", 0}
var d4 DInfo = DInfo{"energy drink", "If you want to break your body, drink a lot", 0}

func GetDrinkMessage() (message string) {
    var d [5]DInfo
    d[0] = DInfo{"water", "1日に2Lの水を飲むと健康に良いですよ", 0}
    d[1] = DInfo{"tea", "カフェインの摂りすぎに注意してくださいね", 0}
    d[2] = DInfo{"cola", "コーラには角砂糖16個分の糖分が含まれているらしいですよ", 0}
    d[3] = DInfo{"energy drink", "体を壊したいときには, エナジードリンクをたくさん飲みましょう", 0}

    num := rand.Intn(4)
    return d[num].DDescription
}

func (u *User) CreateDrink(dName string, amount int) (err error) {
    cmd := `insert into drinks(
        user_id,
        drink_name,
        amount) values(?,?,?)`

    _, err = Db.Exec(cmd, u.ID, dName, amount)

    if err != nil {
        log.Fatalln(err)
    }
    return err
}

func (u *User) GetDrinkSumByCategory() (drinks []Drink, err error) {
    cmd := `select drink_name, SUM(amount) from drinks where user_id = ? group by drink_name`
    rows, err := Db.Query(cmd, u.ID)
    if err != nil {
        log.Fatalln(err)
    }

    for rows.Next() {
        var drink Drink
        err = rows.Scan(
            &drink.DName,
            &drink.Total)

        if err != nil {
            log.Fatalln(err)
        }

        drinks = append(drinks, drink)
    }
    rows.Close()

    return drinks, err
}

func (u *User) GetDrink(id int) (drink Drink, err error) {
    drink = Drink{}
    cmd := `select id,
        DName,
        amount
        from drinks
        where id = ?`
    err = Db.QueryRow(cmd, id).Scan(
        &drink.ID,
        &drink.UserDrink,
        &drink.Amount,
    )
    return drink, err
}

func (d *Drink) UpdateDrink() (err error) {
    cmd := `update drinks set DName = ?,
    amount = ?
    where id = ?`
    _, err = Db.Exec(cmd, d.DName, d.Amount, d.ID)
    if err != nil {
        log.Fatalln(err)
    }
    return err
}

func (u *Drink) DeleteDink() (err error) {
    cmd := `delete from drinks where id = ?`
    _, err = Db.Exec(cmd, u.ID)
    if err != nil {
        log.Fatalln(err)
    }
    return err
}
