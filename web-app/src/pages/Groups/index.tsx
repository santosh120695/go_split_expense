import { Users, Plus, Trash2 } from "lucide-react"
import React, {useState} from "react"
import CreateGroupModal from "./CreateGroupModal.tsx"
import AddExpenseModal from "./AddExpenseModal.tsx"
import type {GroupType} from "../../types/Groups.ts";
import {useMutation, useQuery, useQueryClient} from "@tanstack/react-query";
import {createGroup, fetchGroups} from "../../api/group.ts";
import {Link, useNavigate} from "react-router-dom";


function Index() {
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [isExpenseModalOpen, setIsExpenseModalOpen] = useState(false)
  const [formData, setFormData] = useState({ name: "", description: "" })
  const [selectedGroupId, setSelectedGroupId] = useState<number | null>(null)

  const queryClient = useQueryClient()
  const navigate = useNavigate();

  const {data: groups, isError, isLoading} = useQuery<GroupType[]>({
    queryKey: ['group_list'],
    queryFn: fetchGroups,
  })

  const mutation = useMutation<GroupType, Error, GroupType>({
    mutationFn: (group: GroupType) => {
      return createGroup(group)
    },
    onSuccess: (data) => {
      if (data && data.ID) {
        navigate(`/groups/${data.ID}`);
        queryClient.invalidateQueries({queryKey: ['group_list']});
      } else {
        console.error("No group ID found after creation.");
      }
    }
  })


  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target
    setFormData(prev => ({ ...prev, [name]: value }))
  }

  const handleCreateGroup = (e: React.FormEvent) => {
    e.preventDefault()
    if (formData.name.trim()) {
      const newGroup: GroupType = {
        name: formData.name,
        description: formData.description,
        total_expense: 0,
        icon: "Users",
        currency: "USD",
        ID: null,
        members: []
      }
      mutation.mutate({...newGroup})
      setFormData({ name: "", description: "" })
      setIsModalOpen(false)
    }
  }

  if(isLoading) {
    return <h2>
      Loading Groups
    </h2>
  }

  if(isError){
    return (
        <h1> Error in Fetching Groups </h1>
    )
  }

  return (
    <div className="min-h-screen bg-background p-8">
      <div className="max-w-7xl mx-auto lg:flex lg:justify-center">
        <div className="w-full">
          <div className="flex items-center justify-between mb-8">
            <div>
              <h1 className="text-4xl font-bold text-foreground">Groups</h1>
              <p className="text-muted-foreground mt-2">Manage your shared expense groups</p>
            </div>
            <button
              onClick={() =>  setIsModalOpen(true)}
              className="flex items-center gap-2  text-(--primary) px-6 py-3 border border-(--primary) rounded-lg hover:bg-opacity-90 transition-all duration-200 font-medium"
            >
              <Plus size={20} />
              Create Group
            </button>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {(groups || []).map((group: GroupType) => (
              <div
                key={group.ID}
                className="bg-(--card)  rounded-lg border border-[#C5C3C3] p-6 shadow-sm hover:shadow-md transition-all duration-200"
              >
                <div className="flex items-start justify-between mb-4">
                  <div className={`bg-blue-100 dark:bg-blue-900 p-3 rounded-lg`}>
                    <Users size={24} className="text-blue-600 dark:text-blue-400" />
                  </div>
                  <button className="p-2 hover:bg-red-100 dark:hover:bg-red-900 rounded-lg transition-colors duration-200">
                    <Trash2 size={18} className="text-destructive" />
                  </button>
                </div>

                <h3 className="text-lg font-bold text-foreground mb-1">{group.name}</h3>
                <p className="text-sm text-muted-foreground mb-4">{group.description}</p>

                {/* Group Stats */}
                <div className="space-y-3 border-t border-[#C5C3C3] pt-4">
                  <div className="flex justify-between items-center">
                    <span className="text-sm text-muted-foreground">Members</span>
                    <span className="text-sm font-semibold text-foreground">{group.members.length}</span>
                  </div>
                  <div className="flex justify-between items-center">
                    <span className="text-sm text-muted-foreground">Total Expense</span>
                    <span className="text-sm font-semibold text-foreground">{group.currency}{group.total_expense}</span>
                  </div>
                </div>

                <div className="flex gap-3 mt-4">
                  <button 
                    onClick={() => { setSelectedGroupId(group.ID);  setIsExpenseModalOpen(true) }}
                    className="flex-1 px-4 py-2 bg-primary text-foreground rounded-lg hover:bg-opacity-90 transition-colors duration-200 font-medium text-sm flex items-center justify-center gap-2"
                  >
                    <Plus size={18} />
                    Add Expense
                  </button>
                  <Link
                      to={`/groups/${group.ID}`}
                      className="flex-1 px-4 py-2 bg-accent text-foreground rounded-lg hover:bg-opacity-80 transition-colors duration-200 font-medium text-sm text-center"
                  >
                    View Details
                  </Link>
                </div>
              </div>
            ))}
          </div>

          <CreateGroupModal
            isOpen={isModalOpen}
            onClose={() => setIsModalOpen(false)}
            formData={formData}
            onInputChange={handleInputChange}
            onSubmit={handleCreateGroup}
          />
          {selectedGroupId &&
            <AddExpenseModal
              isOpen={isExpenseModalOpen}
              onClose={() => setIsExpenseModalOpen(false)}
              groupId={selectedGroupId}
            />
          }

          {(groups || []).length === 0 && (
            <div className="flex flex-col items-center justify-center py-20">
              <div className="bg-accent p-6 rounded-full mb-4">
                <Users size={48} className="text-muted-foreground" />
              </div>
              <h3 className="text-lg font-semibold text-foreground mb-2">No Groups Yet</h3>
              <p className="text-muted-foreground mb-6">Create your first group to start tracking shared expenses</p>
              <button className="flex items-center gap-2 bg-primary text-white px-6 py-3 rounded-lg hover:bg-opacity-90 transition-all duration-200 font-medium">
                <Plus size={20} />
                Create Your First Group
              </button>
            </div>
          )}
        </div>
      </div>
    </div>
  )
}

export default Index
