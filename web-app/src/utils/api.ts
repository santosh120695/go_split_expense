import axios from "axios";

const HOST_URL = "https://go-split-expense.onrender.com"
const api = axios.create({
  baseURL: `${HOST_URL}/v1/`,
  headers: {
    "Content-Type": "application/json",
    "Authorization": `${localStorage.getItem("authToken")}`,
  },
});

export default api;
