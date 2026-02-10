# 💸 ExpenseSplitter PWA

A high-performance, offline-first React application for managing group expenses. Built with TypeScript for type safety and Tailwind CSS for a sleek, mobile-first UI.

---

## 🏗️ Technical Stack

* **Framework:** React 18+ (Vite-powered)
* **Language:** TypeScript
* **Styling:** Tailwind CSS
* **State Management:** Context API / Zustand (Local-first)
* **PWA:** vite-plugin-pwa (Service Workers + Manifest)

---

## 🧬 Core Logic: The Balancing Algorithm

To minimize the number of transactions between friends, the app uses a **Net Debt** algorithm. 

1.  **Calculate Net Balance:** For every person $i$, calculate $B_i$:
    $$B_i = \sum \text{Paid}_i - \sum \text{Owed}_i$$
2.  **Settle Up:** * People with $B_i > 0$ are **creditors**.
    * People with $B_i < 0$ are **debtors**.
    * The app greedily matches the largest debtor with the largest creditor until all balances are near zero.



---

## 🛠️ Project Structure

```text
src/
├── components/        # Reusable UI (Buttons, Inputs, Cards)
├── hooks/             # useLocalStorage, useSettlementLogic
├── types/             # Expense, Friend, Group interfaces
├── store/             # Global state (Friends & Transactions)
├── utils/             # Math helpers for splitting logic
└── serviceWorker.ts   # PWA offline configurations
