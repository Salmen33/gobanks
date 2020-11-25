package transactionhandler

import (
	"fmt"
	conn "gobanks/config"
	session "gobanks/config"
	us "gobanks/model/user"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var sessions = session.Sessions()

func NewTransfer(c *fiber.Ctx) error {
	var mail string
	email := getEmail(c, mail)
	return c.Render("users/transfer", fiber.Map{
		"email": email,
	}, "layouts/layout")
}

func NewDeposite(c *fiber.Ctx) error {
	var mail string
	email := getEmail(c, mail)
	return c.Render("users/deposite", fiber.Map{
		"email": email,
	}, "layouts/layout")
}

func NewService(c *fiber.Ctx) error {
	var mail string
	email := getEmail(c, mail)
	return c.Render("users/service", fiber.Map{
		"email": email,
	}, "layouts/layout")
}

func Transfer(c *fiber.Ctx) error {
	var user us.User
	var sender us.User

	db := conn.DBConn()
	store := sessions.Get(c)
	defer store.Save()

	nId := store.Get("id").(int64)
	receiver := c.FormValue("receiver")
	amount, err := strconv.Atoi(c.FormValue("amount"))

	if err != nil {
		panic(err.Error())
	}

	db.Where("id = ?", nId).First(&sender)
	db.Where("email = ?", receiver).First(&user)

	if sender.AccountNumber == user.AccountNumber || sender.Balance < amount {
		return c.JSON("No enough money")
	}
	receiverBalance := user.Balance + amount
	senderBalance := sender.Balance - amount

	fmt.Println(sender.Fullname)
	fmt.Println(senderBalance)

	db.Create(&us.Transfer{
		Receiver: receiver,
		UserID:   uint(nId),
		Amount:   amount,
	})

	db.Model(&sender).Where("id = ?", nId).Update("balance", senderBalance)
	db.Model(&user).Where("email = ?", receiver).Update("balance", receiverBalance)

	return c.JSON("Money Transfered")

}

func History(c *fiber.Ctx) error {
	db := conn.DBConn()

	store := sessions.Get(c)
	defer store.Save()

	nID := store.Get("id").(int64)
	email := store.Get("email")

	var transfer []us.Transfer
	db.Where("user_id = ?", uint(nID)).Find(&transfer)

	return c.Render("users/all", fiber.Map{
		"transactions": transfer,
		"email":        email,
	}, "layouts/layout")
}

func Deposite(c *fiber.Ctx) error {
	var user us.User

	db := conn.DBConn()
	store := sessions.Get(c)
	defer store.Save()

	nId := store.Get("id").(int64)
	receiver := c.FormValue("receiver")
	amount, err := strconv.Atoi(c.FormValue("amount"))

	if err != nil {
		panic(err.Error())
	}

	db.Where("account_number = ?", receiver).First(&user)

	receiverBalance := user.Balance + amount

	db.Create(&us.Transfer{
		Receiver: receiver,
		UserID:   uint(nId),
		Amount:   amount,
	})

	db.Model(&user).Where("account_number = ?", receiver).Update("balance", receiverBalance)

	return c.JSON("Money Added")

}

func getEmail(c *fiber.Ctx, email string) string {
	store := sessions.Get(c)
	defer store.Save()

	email = store.Get("email").(string)

	return email
}
