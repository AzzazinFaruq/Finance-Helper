import { useEffect, useMemo } from "react";
import useTransactionStore from "../stores/transactionStore";
import useBudgetStore from "../stores/budgetStore";
import useAuthStore from "../stores/authStore";

export default function Dashboard() {
  const { user } = useAuthStore();
  const { transactions, fetchTransactions } = useTransactionStore();
  const { budgets, fetchBudgets } = useBudgetStore();

  useEffect(() => {
    fetchTransactions();
    fetchBudgets();
  }, [fetchTransactions, fetchBudgets]);

  const stats = useMemo(() => {
    const now = new Date();
    const currentMonth = `${now.getFullYear()}-${String(now.getMonth() + 1).padStart(2, "0")}`;

    const monthlyTransactions = transactions.filter((t) =>
      t.date?.startsWith(currentMonth)
    );

    const income = monthlyTransactions
      .filter((t) => t.type === "income")
      .reduce((sum, t) => sum + Number(t.amount), 0);

    const expense = monthlyTransactions
      .filter((t) => t.type === "expense")
      .reduce((sum, t) => sum + Number(t.amount), 0);

    const totalBudget = budgets
      .filter((b) => b.month === currentMonth)
      .reduce((sum, b) => sum + Number(b.amount), 0);

    return { income, expense, balance: income - expense, totalBudget };
  }, [transactions, budgets]);

  const recentTransactions = transactions.slice(0, 5);

  const formatCurrency = (amount) =>
    new Intl.NumberFormat("id-ID", { style: "currency", currency: "IDR", minimumFractionDigits: 0 }).format(amount);

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-800">
          Halo, {user?.username || "..."}
        </h1>
        <p className="text-gray-500">Ringkasan keuangan bulan ini</p>
      </div>

      {/* Stats Cards */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
        <div className="bg-white p-5 rounded-xl border border-gray-200">
          <p className="text-sm text-gray-500">Pemasukan</p>
          <p className="text-2xl font-bold text-emerald-600 mt-1">{formatCurrency(stats.income)}</p>
        </div>
        <div className="bg-white p-5 rounded-xl border border-gray-200">
          <p className="text-sm text-gray-500">Pengeluaran</p>
          <p className="text-2xl font-bold text-red-500 mt-1">{formatCurrency(stats.expense)}</p>
        </div>
        <div className="bg-white p-5 rounded-xl border border-gray-200">
          <p className="text-sm text-gray-500">Saldo</p>
          <p className={`text-2xl font-bold mt-1 ${stats.balance >= 0 ? "text-emerald-600" : "text-red-500"}`}>
            {formatCurrency(stats.balance)}
          </p>
        </div>
        <div className="bg-white p-5 rounded-xl border border-gray-200">
          <p className="text-sm text-gray-500">Total Budget</p>
          <p className="text-2xl font-bold text-blue-600 mt-1">{formatCurrency(stats.totalBudget)}</p>
        </div>
      </div>

      {/* Recent Transactions */}
      <div className="bg-white rounded-xl border border-gray-200">
        <div className="p-5 border-b border-gray-200">
          <h2 className="text-lg font-semibold text-gray-800">Transaksi Terakhir</h2>
        </div>
        {recentTransactions.length === 0 ? (
          <div className="p-8 text-center text-gray-400">Belum ada transaksi</div>
        ) : (
          <div className="divide-y divide-gray-100">
            {recentTransactions.map((t) => (
              <div key={t.id} className="flex items-center justify-between px-5 py-3">
                <div>
                  <p className="text-sm font-medium text-gray-800">{t.description || "-"}</p>
                  <p className="text-xs text-gray-400">
                    {t.Category?.category_name || "Tanpa kategori"} &middot; {t.date?.slice(0, 10)}
                  </p>
                </div>
                <span
                  className={`text-sm font-semibold ${
                    t.type === "income" ? "text-emerald-600" : "text-red-500"
                  }`}
                >
                  {t.type === "income" ? "+" : "-"}{formatCurrency(Number(t.amount))}
                </span>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
