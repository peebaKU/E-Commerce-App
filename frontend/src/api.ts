import axios from "axios";

const API_BASE = "http://localhost:8080";

export interface User {
  username?: string;
  email: string;
  password: string;
}

export const register = async (user: User) => {
  return axios.post(`${API_BASE}/register`, user);
};

export const login = async (user: User) => {
  return axios.post(`${API_BASE}/login`, user);
};

export const fetchProducts = async () => {
  return axios.get(`${API_BASE}/products`);
};
