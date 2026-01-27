package handler

import (
	"net/http"
	"splitwise/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionIndexParams struct {
	GroupId int64 `form:"group_id"`
}

func TransactionIndex(c *gin.Context, db *gorm.DB) {
	var params TransactionIndexParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	var transactions []model.Transaction

	query := db
	if params.GroupId != 0 {
		query = query.Where("group_id = ?", params.GroupId).Find(&transactions)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactions,
	})
}

type CreateTransactionParams struct {
	GroupID  int64   `json:"group_id"`
	PaidById int64   `json:"paid_by_id"`
	Amount   float64 `json:"amount"`
	Title    string  `json:"title"`
	UserIds  []int64 `json:"user_ids"`
}

type AmountUser struct {
	UserId int64   `gorm:"column:user_id"`
	Amount float64 `gorm:"column:amount"`
}
type splitParams struct {
	GroupId int64 `uri:"group_id"`
}

type RepayTransaction struct {
	FromId int64
	ToId   int64
	Amount float64
}

func getMinAmout(user_amount1, user_amount2 float64) float64 {
	if user_amount1 > user_amount2 {
		return user_amount2
	}
	return user_amount1
}

func CalculateSplit(c *gin.Context, db *gorm.DB) {
	var usersWhoPaid []AmountUser
	var usersWhoOwe []AmountUser
	var usersAmount []AmountUser
	var group model.Group
	var params splitParams

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	db.Where("id = ?", params.GroupId).First(&group)

	db.Raw(`Select user_id, SUM(net_balance) as amount
			FROM user_transactions
			JOIN transactions ON transactions.id = user_transactions.transaction_id
			WHERE transactions.group_id = ?
			GROUP BY user_id ORDER BY amount DESC`, group.Id).Scan(&usersAmount)

	index := 0
	for usersAmount[index].Amount > 0 {
		usersWhoPaid = append(usersWhoPaid, usersAmount[index])
		index += 1
	}

	usersWhoOwe = usersAmount[index:]
	i := 0
	j := 0
	repay_transactions := []RepayTransaction{}
	for len(usersWhoPaid) > i && len(usersWhoOwe) > j {
		min_amount := getMinAmout(usersWhoPaid[i].Amount, -usersWhoOwe[j].Amount)
		usersWhoPaid[i].Amount = usersWhoPaid[i].Amount - min_amount
		usersWhoOwe[i].Amount = usersWhoOwe[i].Amount - min_amount
		repay_transactions = append(repay_transactions, RepayTransaction{
			FromId: usersWhoOwe[j].UserId,
			ToId:   usersWhoPaid[i].UserId,
			Amount: min_amount,
		})
		if usersWhoPaid[i].Amount == 0 {
			i += 1
		}
		if usersWhoOwe[j].Amount == 0 {
			j += 1
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"user_amounts": usersAmount,
		"repay":        repay_transactions,
	})
}

func TransactionCreate(c *gin.Context, db *gorm.DB) {
	var params CreateTransactionParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	var users []model.User
	db.Find(&users, params.UserIds)
	transaction := model.Transaction{
		Amount:   params.Amount,
		GroupId:  params.GroupID,
		PaidById: params.PaidById,
		Title:    params.Title,
	}

	db.Save(&transaction)
	share := transaction.Amount / float64(len(params.UserIds))
	var user_transactions []model.UserTransaction
	for _, userId := range params.UserIds {

		user_transactions = append(user_transactions,
			model.UserTransaction{
				UserId:        userId,
				TransactionId: int64(transaction.ID),
				Share:         share,
				NetBalance: func() float64 {
					if transaction.PaidById == userId {
						return transaction.Amount - share
					}
					return -share
				}(),
			},
		)
	}
	db.Save(&user_transactions)
	c.JSON(http.StatusOK, gin.H{
		"data": transaction,
	})
}
