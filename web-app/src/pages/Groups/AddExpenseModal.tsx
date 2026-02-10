import { X } from "lucide-react"
import React, {useState} from "react"
import type {GroupType} from "../../types/Groups.ts";
import type {ExpenseType} from "../../types/Expense.ts";
import type {UserType} from "../../types/user.ts";
import {useMutation, useQuery, useQueryClient} from "@tanstack/react-query";
import {addExpense} from "../../api/expense.ts";
import {toast} from "react-toastify";
import {fetchGroup} from "../../api/group.ts";


interface AddExpenseModalProps {
  isOpen: boolean
  onClose: () => void
  groupId: number | null
}

function AddExpenseModal({
  isOpen,
  onClose,
  groupId,
}: AddExpenseModalProps) {

  const queryClient = useQueryClient();
  const [expenseFormData, setExpenseFormData] = useState<ExpenseType>({
      title: "",
      amount: 0,
      group_id: groupId,
      user_ids: [],
      paid_by: {
          ID: 0,
          user_name: "",
          email: "",
          contact_no: "",
      },
  });

  const {data: group} = useQuery<GroupType>({
    queryKey: ['group_details', groupId],
    queryFn: () => fetchGroup(Number(groupId)),
  })

  const {mutate: addExpenseMutation} = useMutation({
      mutationFn: (expense: ExpenseType) => addExpense(expense),
      onSuccess: () => {
          toast.success("Expense added successfully");
          queryClient.invalidateQueries({queryKey: ['group_details', groupId]});
          queryClient.invalidateQueries({queryKey: ['expenseSplits', groupId]});
          onClose();
          setExpenseFormData({
              title: "",
              amount: 0,
              group_id: groupId,
              user_ids: [],
              paid_by: {
                  ID: 0,
                  user_name: "",
                  email: "",
                  contact_no: "",
              },
          });
      },
      onError: (error) => {
          toast.error(`Failed to add expense: ${error.message}`);
      }
  });

  if (!isOpen) return null


  const users: UserType[] = group?.members || []

  const onInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
        const { name, value } = e.target;
        setExpenseFormData((prev) => ({
            ...prev,
            [name]: name === "amount" ? Number(value) : value,
        }));
    };

    const onUserSelect = (userId: number) => {
        setExpenseFormData((prev) => {
            const userIds = prev.user_ids.includes(userId)
                ? prev.user_ids.filter((id) => id !== userId)
                : [...prev.user_ids, userId];
            return { ...prev, user_ids: userIds };
        });
    };

    const handleAddExpenseSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        addExpenseMutation(expenseFormData);
    };

  return (
    <>
        {isOpen && (
              <div
                className="fixed inset-0 bg-(--card) bg-opacity-50 z-40 transition-opacity duration-200"
                onClick={onClose}
              />
            )}
            {isOpen && (
              <div className="fixed inset-0 z-50 flex items-center justify-center p-4 pointer-events-none">
                <div
                  className="bg-card rounded-lg border border-[#C5C3C3] p-8 w-full max-w-md shadow-lg max-h-[90vh] overflow-y-auto pointer-events-auto"            onClick={(e) => e.stopPropagation()}
          >
        <div className="flex items-center justify-between mb-6">
          <h2 className="text-2xl font-bold text-foreground">Add Expense</h2>
          <button
            onClick={onClose}
            className="p-1 hover:bg-accent rounded-lg transition-colors duration-200"
          >
            <X size={24} className="text-muted-foreground" />
          </button>
        </div>

        <form onSubmit={handleAddExpenseSubmit} className="space-y-4">
          <div>
            <label htmlFor="title" className="block text-sm font-medium text-foreground mb-2">
              Expense Title *
            </label>
            <input
              type="text"
              id="title"
              name="title"
              value={expenseFormData.title}
              onChange={onInputChange}
              placeholder="e.g., Dinner, Groceries"
              className="w-full px-4 py-2 bg-background border border-[#C5C3C3] rounded-lg text-foreground placeholder-muted-foreground focus:outline-none focus:ring-2 focus:ring-primary transition-all duration-200"
              required
            />
          </div>

          <div>
            <label htmlFor="amount" className="block text-sm font-medium text-foreground mb-2">
              Amount ($) *
            </label>
            <input
              type="number"
              id="amount"
              name="amount"
              value={expenseFormData.amount}
              onChange={onInputChange}
              placeholder="0.00"
              step="0.01"
              min="0"
              className="w-full px-4 py-2 bg-background border border-[#C5C3C3] rounded-lg text-foreground placeholder-muted-foreground focus:outline-none focus:ring-2 focus:ring-primary transition-all duration-200"
              required
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-foreground mb-2">
              Select Users *
            </label>
            <div className="border border-[#C5C3C3] rounded-lg p-3 space-y-2 bg-background max-h-32 overflow-y-auto">
              {users.map((user: UserType) => (
                <label key={user.ID} className="flex items-center gap-2 cursor-pointer hover:text-primary transition-colors">
                  <input
                    type="checkbox"
                    checked={expenseFormData.user_ids.includes(user.ID as number)}
                    onChange={() => onUserSelect(user.ID as number)}
                    className="w-4 h-4 rounded border-[#C5C3C3] text-primary focus:ring-2 focus:ring-primary cursor-pointer"
                  />
                  <span className="text-sm text-foreground">{user.user_name || user.email}</span>
                </label>
              ))}
            </div>
            {expenseFormData.user_ids.length === 0 && (
              <p className="text-xs text-muted-foreground mt-1">Please select at least one user</p>
            )}
          </div>

          {/* Form Actions */}
          <div className="flex gap-3 pt-4">
            <button
              type="button"
              onClick={onClose}
              className="flex-1 px-4 py-2 bg-accent text-foreground rounded-lg hover:bg-opacity-80 transition-colors duration-200 font-medium"
            >
              Cancel
            </button>
            <button
            type="submit"
            className="flex-1 px-4 py-2 bg-(--card) text-(--primary) border border-(--primary) rounded-lg hover:bg-opacity-90 transition-all duration-200 font-medium"
            >
              Add Expense
            </button>
          </div>
        </form>
          </div>
        </div>
      )}
    </>
  )
}

export default AddExpenseModal
