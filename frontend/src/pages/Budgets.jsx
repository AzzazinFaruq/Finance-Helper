import { useEffect, useState } from "react";
import useBudgetStore from "../stores/budgetStore";
import useCategoryStore from "../stores/categoryStore";
import useAuthStore from "../stores/authStore";

export default function Budgets() {
  const { user } = useAuthStore();
  const { budgets, loading, fetchBudgets, createBudget, updateBudget, deleteBudget } = useBudgetStore();
  const { categories, fetchCategories } = useCategoryStore();

  const [showForm, setShowForm] = useState(false);
  const [editId, setEditId] = useState(null);
  const [form, setForm] = useState({ category_id: "", amount: "", month: "" });

  useEffect(() => {
    fetchBudgets();
    fetchCategories();
  }, [fetchBudgets, fetchCategories]);

  const resetForm = () => {
    setForm({ category_id: "", amount: "", month: "" });
    setEditId(null);
    setShowForm(false);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const data = {
      user_id: user.id,
      category_id: parseInt(form.category_id),
      amount: parseFloat(form.amount),
      month: form.month,
    };

    let success;
    if (editId) {
      success = await updateBudget(editId, data);
    } else {
      success = await createBudget(data);
    }
    if (success) resetForm();
  };

  const handleEdit = (b) => {
    setForm({
      category_id: String(b.category_id),
      amount: String(b.amount),
      month: b.month,
    });
    setEditId(b.id);
    setShowForm(true);
  };

  const handleDelete = async (id) => {
    if (window.confirm("Hapus budget ini?")) {
      await deleteBudget(id);
    }
  };

  const formatCurrency = (amount) =>
    new Intl.NumberFormat("id-ID", { style: "currency", currency: "IDR", minimumFractionDigits: 0 }).format(amount);

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold text-gray-800">Budget</h1>
        <button
          onClick={() => { resetForm(); setShowForm(!showForm); }}
          className="bg-emerald-600 text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-emerald-700 transition cursor-pointer"
        >
          {showForm ? "Batal" : "+ Tambah"}
        </button>
      </div>

      {showForm && (
        <form onSubmit={handleSubmit} className="bg-white p-6 rounded-xl border border-gray-200 grid grid-cols-1 sm:grid-cols-3 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Kategori</label>
            <select
              value={form.category_id}
              onChange={(e) => setForm({ ...form, category_id: e.target.value })}
              className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 outline-none"
              required
            >
              <option value="">Pilih kategori</option>
              {categories.map((c) => (
                <option key={c.id} value={c.id}>{c.category_name}</option>
              ))}
            </select>
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Jumlah Budget</label>
            <input
              type="number"
              step="0.01"
              value={form.amount}
              onChange={(e) => setForm({ ...form, amount: e.target.value })}
              className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 outline-none"
              required
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Bulan</label>
            <input
              type="month"
              value={form.month}
              onChange={(e) => setForm({ ...form, month: e.target.value })}
              className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 outline-none"
              required
            />
          </div>
          <div className="sm:col-span-3">
            <button
              type="submit"
              className="bg-emerald-600 text-white px-6 py-2 rounded-lg text-sm font-medium hover:bg-emerald-700 transition cursor-pointer"
            >
              {editId ? "Update" : "Simpan"}
            </button>
          </div>
        </form>
      )}

      <div className="bg-white rounded-xl border border-gray-200">
        {loading ? (
          <div className="p-8 text-center text-gray-400">Loading...</div>
        ) : budgets.length === 0 ? (
          <div className="p-8 text-center text-gray-400">Belum ada budget</div>
        ) : (
          <table className="w-full text-sm">
            <thead>
              <tr className="border-b border-gray-200 text-left text-gray-500">
                <th className="px-5 py-3 font-medium">Bulan</th>
                <th className="px-5 py-3 font-medium">Kategori</th>
                <th className="px-5 py-3 font-medium text-right">Budget</th>
                <th className="px-5 py-3 font-medium text-right">Aksi</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-100">
              {budgets.map((b) => (
                <tr key={b.id} className="hover:bg-gray-50">
                  <td className="px-5 py-3 text-gray-600">{b.month}</td>
                  <td className="px-5 py-3 text-gray-800">{b.Category?.category_name || "-"}</td>
                  <td className="px-5 py-3 text-right font-medium text-blue-600">{formatCurrency(Number(b.amount))}</td>
                  <td className="px-5 py-3 text-right space-x-2">
                    <button onClick={() => handleEdit(b)} className="text-blue-600 hover:underline cursor-pointer">Edit</button>
                    <button onClick={() => handleDelete(b.id)} className="text-red-500 hover:underline cursor-pointer">Hapus</button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
    </div>
  );
}
