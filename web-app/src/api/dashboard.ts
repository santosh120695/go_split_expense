import api from "../utils/api.ts";


export const fetchDashboardCount =  () => new Promise<{you_owe: {currency: string, amount: number}, you_are_owed: {currency: string, amount: number}, activities: string[]}>((resolve, reject) => {
        api.get("/dashboard").then((response) => {
            resolve(response.data.data)
        }).catch((error) => {
            reject(error)
        })
    }
)