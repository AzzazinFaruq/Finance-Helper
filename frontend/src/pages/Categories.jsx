import { useEffect, useState } from "react";
import useCategoryStore from "../stores/categoryStore";
import useAuthStore from "../stores/authStore";

export default function Categories() {
  const { user } = useAuthStore();
  const { categories, loading, fetchCategories, createCategory, updateCategory, deleteCategory } = useCategoryStore();

  const [showForm, setShowForm] = useState(false);
  const [editId, setEditId] = useState(null);
  const [name, setName] = useState("");

  useEffect(() => {
    fetchCategories();
  }, [fetchCategories]);

  const resetForm = () => {
    setName("");
    setEditId(null);
    setShowForm(false);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const data = { user_id: user.id, name };

    let success;
    if (editId) {
      success = await updateCategory(editId, data);
    } else {
      success = await createCategory(data);
    }
    if (success) resetForm();
  };

  const handleEdit = (c) => {
    setName(c.category_name);
    setEditId(c.id);
    setShowForm(true);
  };

  const handleDelete = async (id) => {
    if (window.confirm("Hapus kategori ini?")) {
      await deleteCategory(id);
    }
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold text-gray-800">Kategori</h1>
        <button
          onClick={() => { resetForm(); setShowForm(!showForm); }}
          className="bg-emerald-600 text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-emerald-700 transition cursor-pointer"
        >
          {showForm ? "Batal" : "+ Tambah"}
        </button>
      </div>

      {showForm && (
        <form onSubmit={handleSubmit} className="bg-white p-6 rounded-xl border border-gray-200 flex gap-4 items-end">
          <div className="flex-1">
            <label className="block text-sm font-medium text-gray-700 mb-1">Nama Kategori</label>
            <input
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
              className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-emerald-500 outline-none"
              required
            />
          </div>
          <button
            type="submit"
            className="bg-emerald-600 text-white px-6 py-2 rounded-lg text-sm font-medium hover:bg-emerald-700 transition cursor-pointer"
          >
            {editId ? "Update" : "Simpan"}
          </button>
        </form>
      )}

      <div className="bg-white rounded-xl border border-gray-200">
        {loading ? (
          <div className="p-8 text-center text-gray-400">Loading...</div>
        ) : categories.length === 0 ? (
          <div className="p-8 text-center text-gray-400">Belum ada kategori</div>
        ) : (
          <div className="divide-y divide-gray-100">
            {categories.map((c) => (
              <div key={c.id} className="flex items-center justify-between px-5 py-4">
                <div>
                  <p className="text-sm font-medium text-gray-800">{c.category_name}</p>
                  <p className="text-xs text-gray-400">Dibuat: {c.created_at?.slice(0, 10)}</p>
                </div>
                <div className="space-x-2">
                  <button onClick={() => handleEdit(c)} className="text-sm text-blue-600 hover:underline cursor-pointer">Edit</button>
                  <button onClick={() => handleDelete(c.id)} className="text-sm text-red-500 hover:underline cursor-pointer">Hapus</button>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
