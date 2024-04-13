package repositories

import (
	"fmt"

	"leanmeal/api/interfaces"
	"leanmeal/api/models"
)

type Accounts struct {
	ConnectionString string
	Storage          interfaces.Storage
}

func (accountService *Accounts) OpenConnection(storage *interfaces.Storage) bool {

	accountService.Storage = *storage
	accountService.Storage.Open(accountService.ConnectionString)
	return true

}

func (accountService *Accounts) UserExists(email string) models.Account {
	var account models.Account
	data := accountService.Storage.Single("select id, email from public.accounts where email = $1", []interface{}{email})

	err := data.Scan(&account.Id, &account.Email)
	if err != nil {
		fmt.Printf("Failed to fetch account with email %v", email)
		fmt.Println(err)
		return account
	}

	fmt.Println(&account)
	return account
}

func (accountService *Accounts) Get() []models.Account {

	rows := accountService.Storage.Where("SELECT * from public.accounts", []interface{}{})

	var accounts []models.Account

	for rows.Next() {
		var account models.Account
		rows.Scan(&account.Id, &account.Email, &account.Name, &account.CreatedAt, &account.Enabled)
		accounts = append(accounts, account)
		fmt.Println(account)
	}

	return accounts
}

func (accountService *Accounts) Close() bool {
	accountService.Storage.Close()
	return true
}
