import { useEffect, useState } from "react";
import useTransactionStore from "../stores/transactionStore";
import useCategoryStore from "../stores/categoryStore";
import useAuthStore from "../stores/authStore";

export default function Transactions() {
  const { user } = useAuthStore();
  const { transactions, loading, fetchTransactions, createTransaction, updateTransaction, deleteTransaction } = useTransactionStore();
  const { categories, fetchCategories } = useCategoryStore();

  const [showForm, setShowForm] = useState(false);
  const [editId, setEditId] = useState(null);
  const [form, setForm] = useState({ amount: "", type: "expense", date: "", description: "", category_id: "" });

  useEffect(() => {
    fetchTransactions();
    fetchCategories();
  }, [fetchTransactions, fetchCategories]);

  const resetForm = () => {
    setForm({ amount: "", type: "expense", date: "", description: "", category_id: "" });
    setEditId(null);
    setShowForm(false);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const data = {
      user_id: user.id,
      amount: parseFloat(form.amount),
      type: form.type,
      date: form.date,
      description: form.description,
      category_id: parseInt(form.category_id),
    };

    let success;
    if (editId) {
      success = await updateTransaction(editId, data);
    } else {
      success = await createTransaction(data);
    }
    if (success) resetForm();
  };

  const handleEdit = (t) => {
    setForm({
      amount: String(t.amount),
      type: t.type,
      date: t.date?.slice(0, 10),
      description: t.description,
      category_id: String(t.category_id),
    });
    setEditId(t.id);
    setShowForm(true);
  };

  const handleDelete = async (id) => {
    if (window.confirm("Hapus transaksi ini?")) {
      await deleteTransaction(id);
    }
  };

  const formatCurrency = (amount) =>
    new Intl.NumberFormat("id-ID", { style: "currency", currency: "IDR", minimumFractionDigits: 0 }).format(amount);

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold text-gray-800">Transaksi</h1>
        <button
          onClick={() => { resetForm(); setShowForm(!showForm); }}
          className="bg-emerald-600 text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-emerald-700 transition cursor-pointer"
        >
          {showForm ? "Batal" : "+ Tambah"}
        </button>
      </div>

      {/* Form */}
      {showForm && (
        <form onSubmit={handleSubmit} className="bg-white p-6 rounded-xl border border-gray-200 grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Jumlah</label>
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
            <label className="block text-sm font-medium text-gray-700 mb-1">Tipe</label>
            <select
              value={form.type}
              onChange={(e) => setForm({ ...form, type: e.target.value })}
              className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 outline-none"
            >
              <option value="expense">Pengeluaran</option>
              <option value="income">Pemasukan</option>
            </select>
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Tanggal</label>
            <input
              type="date"
              value={form.date}
              onChange={(e) => setForm({ ...form, date: e.target.value })}
              className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 outline-none"
              required
            />
          </div>
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
          <div className="sm:col-span-2">
            <label className="block text-sm font-medium text-gray-700 mb-1">Deskripsi</label>
            <input
              type="text"
              value={form.description}
              onChange={(e) => setForm({ ...form, description: e.target.value })}
              className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 outline-none"
            />
          </div>
          <div className="sm:col-span-2">
            <button
              type="submit"
              className="bg-emerald-600 text-white px-6 py-2 rounded-lg text-sm font-medium hover:bg-emerald-700 transition cursor-pointer"
            >
              {editId ? "Update" : "Simpan"}
            </button>
          </div>
        </form>
      )}

      {/* Table */}
      <div className="bg-white rounded-xl border border-gray-200 overflow-x-auto">
        {loading ? (
          <div className="p-8 text-center text-gray-400">Loading...</div>
        ) : transactions.length === 0 ? (
          <div className="p-8 text-center text-gray-400">Belum ada transaksi</div>
        ) : (
          <table className="w-full text-sm">
            <thead>
              <tr className="border-b border-gray-200 text-left text-gray-500">
                <th className="px-5 py-3 font-medium">Tanggal</th>
                <th className="px-5 py-3 font-medium">Deskripsi</th>
                <th className="px-5 py-3 font-medium">Kategori</th>
                <th className="px-5 py-3 font-medium">Tipe</th>
                <th className="px-5 py-3 font-medium text-right">Jumlah</th>
                <th className="px-5 py-3 font-medium text-right">Aksi</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-100">
              {transactions.map((t) => (
                <tr key={t.id} className="hover:bg-gray-50">
                  <td className="px-5 py-3 text-gray-600">{t.date?.slice(0, 10)}</td>
                  <td className="px-5 py-3 text-gray-800">{t.description || "-"}</td>
                  <td className="px-5 py-3 text-gray-600">{t.Category?.category_name || "-"}</td>
                  <td className="px-5 py-3">
                    <span className={`inline-block px-2 py-0.5 rounded text-xs font-medium ${
                      t.type === "income" ? "bg-emerald-100 text-emerald-700" : "bg-red-100 text-red-700"
                    }`}>
                      {t.type === "income" ? "Masuk" : "Keluar"}
                    </span>
                  </td>
                  <td className={`px-5 py-3 text-right font-medium ${t.type === "income" ? "text-emerald-600" : "text-red-500"}`}>
                    {formatCurrency(Number(t.amount))}
                  </td>
                  <td className="px-5 py-3 text-right space-x-2">
                    <button onClick={() => handleEdit(t)} className="text-blue-600 hover:underline cursor-pointer">Edit</button>
                    <button onClick={() => handleDelete(t.id)} className="text-red-500 hover:underline cursor-pointer">Hapus</button>
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
