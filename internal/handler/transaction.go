package handler

import (
	"fmt"
	"net/http"
	"slices"
	"splitwise/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionIndexParams struct {
	GroupId int64 `form:"group_id"`
	UserId  int64 `form:"user_id"`
}

func TransactionIndex(c *gin.Context, db *gorm.DB) {
	var params TransactionIndexParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	if params.GroupId == 0 && params.UserId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "GroupId or UserId is required",
			"success": false,
		})
		return
	}

	var transactions []model.Transaction

	query := db.WithContext(c.Request.Context())
	if params.GroupId != 0 {
		query = query.Where("group_id = ?", params.GroupId)
	}
	if params.UserId != 0 {
		query = query.Where("user_id = ?", params.UserId)
	}
	query.Order("created_at desc").Find(&transactions)

	c.JSON(http.StatusOK, gin.H{
		"data":    transactions,
		"success": true,
	})
}

type CreateTransactionParams struct {
	GroupID float64   `json:"group_id"`
	Amount  float64   `json:"amount"`
	Title   string    `json:"title"`
	UserIds []float64 `json:"user_ids"`
}

type AmountUser struct {
	UserId   float64 `gorm:"column:user_id"`
	Amount   float64 `gorm:"column:amount"`
	Username string  `gorm:"column:user_name"`
}
type splitParams struct {
	GroupId int64 `uri:"group_id"`
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

	db.WithContext(c.Request.Context()).Where("id = ?", params.GroupId).First(&group)

	db.WithContext(c.Request.Context()).Raw(`Select users.user_name, user_id, SUM(net_balance) as amount
			FROM user_transactions
			JOIN transactions ON transactions.id = user_transactions.transaction_id JOIN users ON user_transactions.user_id = users.id
			WHERE transactions.group_id = ?
			GROUP BY user_id ORDER BY amount DESC`, group.ID).Scan(&usersAmount)

	index := 0
	for usersAmount[index].Amount > 0 {
		usersWhoPaid = append(usersWhoPaid, usersAmount[index])
		index += 1
	}

	usersWhoOwe = usersAmount[index:]

	i := 0
	j := 0

	fmt.Println("userAmount", usersAmount)
	fmt.Println("usersWhoOwe ", usersWhoOwe)
	fmt.Println("usersWhoPaid ", usersWhoPaid)
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

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user_amounts": usersAmount,
			"repay":        repayTransactions,
		},
	})
}

func TransactionCreate(c *gin.Context, db *gorm.DB) {
	var params CreateTransactionParams
	currentUserId, _ := c.Get("current_user")
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var users []model.User

	db.WithContext(c.Request.Context()).Find(&users, params.UserIds)
	var transaction model.Transaction
	var userTransactions []model.UserTransaction
	err := db.WithContext(c.Request.Context()).Transaction(func(tx *gorm.DB) error {

		transaction = model.Transaction{
			Amount:   params.Amount,
			GroupId:  params.GroupID,
			Title:    params.Title,
			PaidById: currentUserId.(float64),
		}

		err := db.WithContext(c.Request.Context()).Create(&transaction)
		if err.Error != nil {
			return err.Error
		}

		share := transaction.Amount / float64(len(params.UserIds))

		userTransactions = append(userTransactions, model.UserTransaction{
			UserId:        transaction.PaidById,
			Share:         share,
			TransactionId: int64(transaction.ID),
			NetBalance: func() float64 {
				if slices.Contains(params.UserIds, transaction.PaidById) {
					return transaction.Amount - share
				}
				return transaction.Amount
			}(),
		})

		fmt.Println("userTransactions", userTransactions)

		for _, userId := range params.UserIds {
			if userId != transaction.PaidById {
				userTransactions = append(userTransactions,
					model.UserTransaction{
						UserId:        userId,
						TransactionId: int64(transaction.ID),
						Share:         share,
						NetBalance:    -share,
					})
			}
		}
		err = db.WithContext(c.Request.Context()).Create(&userTransactions)
		if err.Error != nil {
			return err.Error
		}
		return nil

	})

	fmt.Println(err)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userTransactions,
	})
}
