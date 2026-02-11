package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name         string        `json:"name"`
	CreatedById  float64       `json:"created_by_id"`
	Users        []User        `gorm:"many2many:user_groups"`
	Transactions []Transaction `json:"transactions"`
	Closed       bool          `gorm:"default:false" json:"closed"`
	Currency     string        `json:"currency"`
	Icon         string        `gorm:"default:Users" json:"icon"`
	TotalAmount  float64       `json:"total_amount" gorm:"default:0"`
	Description  string        `json:"description"`
	Activities   []Activity    `json:"activities"`
}

type AmountUser struct {
	UserId   float64 `gorm:"column:user_id"`
	Amount   float64 `gorm:"column:amount"`
	Username string  `gorm:"column:user_name"`
}

type RepayTransaction struct {
	FromId int64
	ToId   int64
	From   string
	To     string
	Amount float64
}

func getMinAmount(userAmount1, userAmount2 float64) float64 {
	if userAmount1 > userAmount2 {
		return userAmount2
	}
	return userAmount1
}

func (group Group) CalculateRepayments(db *gorm.DB) []RepayTransaction {
	var usersWhoPaid []AmountUser
	var usersWhoOwe []AmountUser
	var usersAmount []AmountUser

	db.Raw(`Select users.user_name, user_id, SUM(net_balance) as amount
			FROM user_transactions
			JOIN transactions ON transactions.id = user_transactions.transaction_id JOIN users ON user_transactions.user_id = users.id
			WHERE transactions.group_id = ?
			GROUP BY user_id ORDER BY amount DESC`, group.ID).Scan(&usersAmount)

	if len(usersAmount) == 0 {
		return make([]RepayTransaction, 0)
	}
	index := 0
	for usersAmount[index].Amount > 0 {
		usersWhoPaid = append(usersWhoPaid, usersAmount[index])
		index += 1
	}

	usersWhoOwe = usersAmount[index:]

	i := 0
	j := 0

	var repayTransactions []RepayTransaction
	for len(usersWhoPaid) > i && len(usersWhoOwe) > j {
		minAmount := getMinAmount(usersWhoPaid[i].Amount, -usersWhoOwe[j].Amount)
		if minAmount > 0 {
			usersWhoPaid[i].Amount = usersWhoPaid[i].Amount - minAmount
			usersWhoOwe[i].Amount = (-usersWhoOwe[i].Amount) - minAmount
			repayTransactions = append(repayTransactions, RepayTransaction{
				From:   usersWhoOwe[j].Username,
				To:     usersWhoPaid[i].Username,
				Amount: minAmount,
			})
			if usersWhoPaid[i].Amount == 0 {
				i += 1
			}
			if usersWhoOwe[j].Amount == 0 {
				j += 1
			}
		}
	}

	return repayTransactions
}
