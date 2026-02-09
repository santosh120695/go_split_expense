import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:3000/v1/",
  headers: {
    "Content-Type": "application/json",
    "Authorization": `${localStorage.getItem("authToken")}`,
  },
});

export default api;
