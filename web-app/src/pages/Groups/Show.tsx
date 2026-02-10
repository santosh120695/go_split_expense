import { TrendingUp, Wallet, UserPlus, PlusCircle } from "lucide-react";
import {useParams} from "react-router-dom";
import {useMutation, useQuery, useQueryClient} from "@tanstack/react-query";
import {expenseRepays, fetchGroup, addUserToGroup} from "../../api/group.ts";
import moment from "moment";
import {useState} from "react";
import AddUserToGroupModal from "./AddUserToGroupModal.tsx";
import AddExpenseModal from "./AddExpenseModal.tsx";
import {toast} from "react-toastify";

function GroupDetails() {
  const {id} = useParams<{ id: string }>();
  const queryClient = useQueryClient();
  const [isAddUserModalOpen, setIsAddUserModalOpen] = useState(false);
  const [isAddExpenseModalOpen, setIsAddExpenseModalOpen] = useState(false);
  const [selectedUserIds, setSelectedUserIds] = useState<number[]>([]);

  const {data: group, isError, isLoading} = useQuery({
    queryKey: ['group_details', id],
    queryFn: () => fetchGroup(Number(id)),
    enabled: !!id,
  })

  const {data: expenseSplit, isLoading: expenseLoading} = useQuery< {From: string, To: string, Amount: number}[]> ({
    queryKey: ['expenseSplits', id],
    queryFn: () => expenseRepays(Number(id)),
    enabled: !!id,
  })

  const {mutate: addUserToGroupMutation} = useMutation({
    mutationFn: async () => {
        if (!selectedUserIds || selectedUserIds.length === 0) {
            toast.error("Please select at least one user to add");
            return;
        }
        const groupId = Number(id);
        await addUserToGroup(groupId, selectedUserIds)
    },
    onSuccess: () => {
        toast.success("Users added to group successfully");
        queryClient.invalidateQueries({queryKey: ['group_details', id]});
        setIsAddUserModalOpen(false);
        setSelectedUserIds([]);
    },
    onError: (error) => {
        toast.error(`Failed to add users to group: ${error.message}`);
    }
  });

  if (isLoading) {
    return <h2>
      Loading Group Details...
    </h2>
  }

  if (isError) {
    return (
        <h1> Error in Fetching Group Details </h1>
    )
  }

  if (!group) {
    return <h2>Group not found</h2>
  }

  const handleAddUserToGroup = (e: React.FormEvent) => {
      e.preventDefault();
      if (selectedUserIds.length > 0) {
          addUserToGroupMutation();
      } else {
          toast.error("Please select at least one user to add");
      }
  }

  return (
      <>
      <div className="mt-12">
        <div className="bg-(--card) rounded-lg border border-[#C5C3C3] p-8 m-5">
          <div className="flex items-center justify-between mb-6">
            <div>
              <h2 className="text-3xl font-bold text-foreground">{group.name}</h2>
              <p className="text-muted-foreground text-sm mt-1">{group.description}</p>
            </div>
            <div className="flex gap-2">
              <button
                  onClick={() => setIsAddExpenseModalOpen(true)}
                  className="p-2  text-green-500 rounded-lg  transition-all duration-200"
              >
                  <PlusCircle size={24} />
              </button>
              <button
                  onClick={() => setIsAddUserModalOpen(true)}
                  className="p-2 bg-primary text-(--primary) rounded-lg hover:bg-opacity-90 transition-all duration-200"
              >
                  <UserPlus size={24} />
              </button>
            </div>
          </div>

          <div className="mb-8">
            <h3 className="text-xl font-bold text-foreground mb-4 flex items-center gap-2">
              <TrendingUp size={20}/>
              Activities
            </h3>
            <div className="space-y-3">
              {(group.expenses || []).map((expense) => {
                return (
                    <div
                        className="border border-[#C5C3C3] rounded-lg p-4 hover:bg-background transition-colors duration-200">
                      <div className="flex items-center justify-between mb-2">
                        <p className="font-medium text-foreground">{expense.title}</p>
                        <span className="text-lg font-bold text-primary">{expense.amount}</span>
                      </div>
                      <div className="flex items-center justify-between text-sm text-muted-foreground">
                        <span>{expense?.paid_by?.user_name || ""}</span>
                        <span>{moment(expense?.created_at).format("DD-MM-YYYY")}</span>
                      </div>
                    </div>
                );
              })}
            </div>
          </div>

          <div>
            <h3 className="text-xl font-bold text-foreground mb-4 flex items-center gap-2">
              <Wallet size={20}/>
              Who Owes What
            </h3>
            {!expenseLoading &&
            <div className="space-y-3">
              <div className="border border-[#C5C3C3] rounded-lg p-4 hover:bg-background transition-colors duration-200">
                {(expenseSplit || []).map((repay_transaction: {From: string, To: string, Amount: number}) => (
                    <div className="flex items-center justify-between mb-2 border-b-2 border-gray-200 ">
                      <div>
                        <p className="text-foreground font-medium">{repay_transaction.From}</p>
                        <p className="text-sm text-muted-foreground">owes</p>
                      </div>
                      <div className="text-right">
                        <p className="text-foreground font-medium">{repay_transaction.To}</p>
                      </div>
                      <div className="ml-4">
                        <span className="text-xl font-bold text-red-600 dark:text-red-400">{repay_transaction.Amount}</span>
                      </div>
                    </div>
                ))}
              </div>
            </div>
            }
          </div>
        </div>
      </div>
    <AddUserToGroupModal
        isOpen={isAddUserModalOpen}
        onClose={() => setIsAddUserModalOpen(false)}
        onSelectUser={setSelectedUserIds}
        onSubmit={handleAddUserToGroup}
    />
    <AddExpenseModal
        isOpen={isAddExpenseModalOpen}
        onClose={() => setIsAddExpenseModalOpen(false)}
        groupId={Number(id)}
    />
    </>
  )
}

export default GroupDetails