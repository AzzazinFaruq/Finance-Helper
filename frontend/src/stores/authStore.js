import { create } from "zustand";
import api from "../api/axios";

const useAuthStore = create((set) => ({
  user: null,
  token: localStorage.getItem("token") || null,
  loading: false,
  error: null,

  register: async (username, password) => {
    set({ loading: true, error: null });
    try {
      await api.post("/register", { username, password });
      set({ loading: false });
      return true;
    } catch (err) {
      set({ loading: false, error: err.response?.data?.error || "Register gagal" });
      return false;
    }
  },

  login: async (username, password) => {
    set({ loading: true, error: null });
    try {
      const res = await api.post("/login", { username, password });
      const { token, user } = res.data;
      localStorage.setItem("token", token);
      set({ token, user, loading: false });
      return true;
    } catch (err) {
      set({ loading: false, error: err.response?.data?.error || "Login gagal" });
      return false;
    }
  },

  logout: async () => {
    try {
      await api.post("/api/logout");
    } catch {
      // ignore
    }
    localStorage.removeItem("token");
    set({ user: null, token: null });
  },

  fetchUser: async () => {
    try {
      const res = await api.get("/api/users");
      set({ user: res.data.user });
    } catch {
      set({ user: null, token: null });
      localStorage.removeItem("token");
    }
  },

  clearError: () => set({ error: null }),
}));

export default useAuthStore;
