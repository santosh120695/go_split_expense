import axios from "axios";

// const HOST_URL = "https://go-split-expense.onrender.com"
const HOST_URL = "http://192.168.0.149:3000";
const api = axios.create({
  baseURL: `${HOST_URL}/v1/`,
  headers: {
    "Content-Type": "application/json",
  },
});

api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem("authToken");
    if (token) {
      config.headers.Authorization = `${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default api;
