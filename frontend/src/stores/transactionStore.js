import { create } from "zustand";
import api from "../api/axios";

const useTransactionStore = create((set) => ({
  transactions: [],
  loading: false,
  error: null,

  fetchTransactions: async () => {
    set({ loading: true });
    try {
      const res = await api.get("/api/get-transaction");
      set({ transactions: res.data.data || [], loading: false });
    } catch (err) {
      set({ loading: false, error: err.response?.data?.error || "Gagal memuat transaksi" });
    }
  },

  createTransaction: async (data) => {
    try {
      await api.post("/api/add-transaction", data);
      const res = await api.get("/api/get-transaction");
      set({ transactions: res.data.data || [] });
      return true;
    } catch (err) {
      set({ error: err.response?.data?.error || "Gagal membuat transaksi" });
      return false;
    }
  },

  updateTransaction: async (id, data) => {
    try {
      await api.put(`/api/update-transaction/${id}`, data);
      const res = await api.get("/api/get-transaction");
      set({ transactions: res.data.data || [] });
      return true;
    } catch (err) {
      set({ error: err.response?.data?.error || "Gagal update transaksi" });
      return false;
    }
  },

  deleteTransaction: async (id) => {
    try {
      await api.delete(`/api/delete-transaction/${id}`);
      set((state) => ({
        transactions: state.transactions.filter((t) => t.id !== id),
      }));
      return true;
    } catch (err) {
      set({ error: err.response?.data?.error || "Gagal hapus transaksi" });
      return false;
    }
  },

  clearError: () => set({ error: null }),
}));

export default useTransactionStore;
