import api from "../utils/api.ts";
import type {ExpenseType} from "../types/Expense.ts";

export const addExpense = (expense: ExpenseType) => new Promise<ExpenseType>((resolve, reject) => {
    api.post('/transactions', expense).then((res) => {
        resolve(res.data.data)
    }).catch((err) => {
        reject(err)
    })
})

