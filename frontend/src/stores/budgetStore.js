import { create } from "zustand";
import api from "../api/axios";

const useBudgetStore = create((set) => ({
  budgets: [],
  loading: false,
  error: null,

  fetchBudgets: async () => {
    set({ loading: true });
    try {
      const res = await api.get("/api/get-budget");
      set({ budgets: res.data.data || [], loading: false });
    } catch (err) {
      set({ loading: false, error: err.response?.data?.error || "Gagal memuat budget" });
    }
  },

  createBudget: async (data) => {
    try {
      await api.post("/api/add-budget", data);
      const res = await api.get("/api/get-budget");
      set({ budgets: res.data.data || [] });
      return true;
    } catch (err) {
      set({ error: err.response?.data?.error || "Gagal membuat budget" });
      return false;
    }
  },

  updateBudget: async (id, data) => {
    try {
      await api.put(`/api/update-budget/${id}`, data);
      const res = await api.get("/api/get-budget");
      set({ budgets: res.data.data || [] });
      return true;
    } catch (err) {
      set({ error: err.response?.data?.error || "Gagal update budget" });
      return false;
    }
  },

  deleteBudget: async (id) => {
    try {
      await api.delete(`/api/delete-budget/${id}`);
      set((state) => ({
        budgets: state.budgets.filter((b) => b.id !== id),
      }));
      return true;
    } catch (err) {
      set({ error: err.response?.data?.error || "Gagal hapus budget" });
      return false;
    }
  },

  clearError: () => set({ error: null }),
}));

export default useBudgetStore;
