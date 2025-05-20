// src/api/axios.ts
import axios from 'axios';

// ✅ Public instance — no auth token
export const publicAxios = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  // withCredentials: true,
});

// 🔐 Protected instance — adds token automatically
export const privateAxios = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  withCredentials: true,
});

privateAxios.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) config.headers.Authorization = `Bearer ${token}`;
    return config;
  },
  (error) => Promise.reject(error)
);