import { create } from "zustand";
import api from "../api/axios";

const useCategoryStore = create((set) => ({
  categories: [],
  loading: false,
  error: null,

  fetchCategories: async () => {
    set({ loading: true });
    try {
      const res = await api.get("/api/get-category");
      set({ categories: res.data.data || [], loading: false });
    } catch (err) {
      set({ loading: false, error: err.response?.data?.error || "Gagal memuat kategori" });
    }
  },

  createCategory: async (data) => {
    try {
      await api.post("/api/add-category", data);
      const res = await api.get("/api/get-category");
      set({ categories: res.data.data || [] });
      return true;
    } catch (err) {
      set({ error: err.response?.data?.error || "Gagal membuat kategori" });
      return false;
    }
  },

  updateCategory: async (id, data) => {
    try {
      await api.put(`/api/update-category/${id}`, data);
      const res = await api.get("/api/get-category");
      set({ categories: res.data.data || [] });
      return true;
    } catch (err) {
      set({ error: err.response?.data?.error || "Gagal update kategori" });
      return false;
    }
  },

  deleteCategory: async (id) => {
    try {
      await api.delete(`/api/delete-category/${id}`);
      set((state) => ({
        categories: state.categories.filter((c) => c.id !== id),
      }));
      return true;
    } catch (err) {
      set({ error: err.response?.data?.error || "Gagal hapus kategori" });
      return false;
    }
  },

  clearError: () => set({ error: null }),
}));

export default useCategoryStore;
