import { Users, Plus, Trash2 } from "lucide-react"
import { useState } from "react"
import CreateFriendModal from "../component/CreateFriendModal.tsx"

function Friends() {
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [formData, setFormData] = useState({ name: "", email: "" })
  const [friends, setFriends] = useState([
    {
      id: 1,
      name: "John Doe",
      email: "john@example.com",
      phone: "+1 234-567-8900",
      totalSpent: "$450.00",
      balance: "$12.50 owes you",
    },
    {
      id: 2,
      name: "Jane Smith",
      email: "jane@example.com",
      phone: "+1 234-567-8901",
      totalSpent: "$680.75",
      balance: "$25.00 you owe",
    },
    {
      id: 3,
      name: "Mike Johnson",
      email: "mike@example.com",
      phone: "+1 234-567-8902",
      totalSpent: "$320.50",
      balance: "settled",
    },
    {
      id: 4,
      name: "Sarah Williams",
      email: "sarah@example.com",
      phone: "+1 234-567-8903",
      totalSpent: "$540.25",
      balance: "$8.75 owes you",
    },
  ])

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target
    setFormData(prev => ({ ...prev, [name]: value }))
  }

  const handleAddFriend = (e: React.FormEvent) => {
    e.preventDefault()
    if (formData.name.trim() && formData.email.trim()) {
      const newFriend = {
        id: Math.max(...friends.map(f => f.id), 0) + 1,
        name: formData.name,
        email: formData.email,
        phone: "",
        totalSpent: "$0.00",
        balance: "settled",
      }
      setFriends([...friends, newFriend])
      setFormData({ name: "", email: "" })
      setIsModalOpen(false)
      alert("Friend added successfully!")
    }
  }

  const handleDeleteFriend = (id: number) => {
    setFriends(friends.filter(f => f.id !== id))
  }

  return (
    <div className="min-h-screen bg-background p-8">
      <div className="max-w-7xl mx-auto lg:flex lg:justify-center">
        {/* Header Section */}
        <div className="w-full">
          <div className="flex items-center justify-between mb-8">
            <div>
              <h1 className="text-4xl font-bold text-foreground">Friends</h1>
              <p className="text-muted-foreground mt-2">Manage your friends and track payments</p>
            </div>
            <button 
              onClick={() => setIsModalOpen(true)}
              className="flex items-center gap-2 bg-[var(--primary)] text-[var(--card)] px-6 py-3 border border-[#C5C3C3] rounded-lg hover:bg-opacity-90 transition-all duration-200 font-medium"
            >
              <Plus size={20} />
              Add Friend
            </button>
          </div>

          {/* Friends Grid */}
          <div className="bg-card rounded-lg border border-[#C5C3C3] overflow-hidden">
            {/* Table Header */}
            <div className="hidden md:grid md:grid-cols-6 gap-4 bg-background p-4 border-b border-[#C5C3C3] font-semibold text-foreground">
              <div>Name</div>
              <div>Email</div>
              <div>Phone</div>
              <div>Total Spent</div>
              <div>Balance</div>
              <div>Actions</div>
            </div>

            {/* Friends List */}
            <div className="divide-y divide-[#C5C3C3]">
              {friends.map((friend) => (
                <div
                  key={friend.id}
                  className="p-4 hover:bg-background transition-colors duration-200"
                >
                  {/* Mobile View */}
                  <div className="md:hidden space-y-3 mb-4">
                    <div className="flex items-start justify-between">
                      <div>
                        <h3 className="text-lg font-bold text-foreground">{friend.name}</h3>
                        <p className="text-sm text-muted-foreground">{friend.email}</p>
                      </div>
                      <button 
                        onClick={() => handleDeleteFriend(friend.id)}
                        className="p-2 hover:bg-red-100 dark:hover:bg-red-900 rounded-lg transition-colors duration-200"
                      >
                        <Trash2 size={18} className="text-destructive" />
                      </button>
                    </div>
                    {friend.phone && <p className="text-sm text-muted-foreground">{friend.phone}</p>}
                    <div className="flex justify-between items-center text-sm">
                      <span className="text-muted-foreground">Total Spent: {friend.totalSpent}</span>
                      <span className={`font-semibold ${
                        friend.balance === "settled" 
                          ? "text-green-600 dark:text-green-400" 
                          : friend.balance.includes("owes you")
                          ? "text-blue-600 dark:text-blue-400"
                          : "text-orange-600 dark:text-orange-400"
                      }`}>
                        {friend.balance}
                      </span>
                    </div>
                    <div className="flex gap-2">
                      <button className="flex-1 px-3 py-2 bg-primary text-white rounded text-sm hover:bg-opacity-90 transition-colors duration-200 font-medium">
                        Settle
                      </button>
                      <button className="flex-1 px-3 py-2 bg-accent text-foreground rounded text-sm hover:bg-opacity-80 transition-colors duration-200 font-medium">
                        Details
                      </button>
                    </div>
                  </div>

                  {/* Desktop View */}
                  <div className="hidden md:grid md:grid-cols-6 gap-4 items-center">
                    <div>
                      <p className="font-medium text-foreground">{friend.name}</p>
                    </div>
                    <div>
                      <p className="text-sm text-muted-foreground">{friend.email}</p>
                    </div>
                    <div>
                      <p className="text-sm text-muted-foreground">{friend.phone || "-"}</p>
                    </div>
                    <div>
                      <p className="text-sm font-semibold text-foreground">{friend.totalSpent}</p>
                    </div>
                    <div>
                      <span className={`text-sm font-semibold ${
                        friend.balance === "settled" 
                          ? "text-green-600 dark:text-green-400" 
                          : friend.balance.includes("owes you")
                          ? "text-blue-600 dark:text-blue-400"
                          : "text-orange-600 dark:text-orange-400"
                      }`}>
                        {friend.balance}
                      </span>
                    </div>
                    <div className="flex gap-2">
                      <button className="px-3 py-1 bg-primary text-white rounded text-sm hover:bg-opacity-90 transition-colors duration-200 font-medium">
                        Settle
                      </button>
                      <button className="px-3 py-1 bg-accent text-foreground rounded text-sm hover:bg-opacity-80 transition-colors duration-200 font-medium">
                        Details
                      </button>
                      <button 
                        onClick={() => handleDeleteFriend(friend.id)}
                        className="p-1 hover:bg-red-100 dark:hover:bg-red-900 rounded transition-colors duration-200"
                      >
                        <Trash2 size={16} className="text-destructive" />
                      </button>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>

          {/* Create Friend Modal */}
          <CreateFriendModal
            isOpen={isModalOpen}
            onClose={() => setIsModalOpen(false)}
            formData={formData}
            onInputChange={handleInputChange}
            onSubmit={handleAddFriend}
          />

          {/* Empty State */}
          {friends.length === 0 && (
            <div className="flex flex-col items-center justify-center py-20">
              <div className="bg-accent p-6 rounded-full mb-4">
                <Users size={48} className="text-muted-foreground" />
              </div>
              <h3 className="text-lg font-semibold text-foreground mb-2">No Friends Yet</h3>
              <p className="text-muted-foreground mb-6">Add friends to start tracking shared expenses</p>
              <button 
                onClick={() => setIsModalOpen(true)}
                className="flex items-center gap-2 bg-primary text-white px-6 py-3 rounded-lg hover:bg-opacity-90 transition-all duration-200 font-medium"
              >
                <Plus size={20} />
                Add Your First Friend
              </button>
            </div>
          )}
        </div>
      </div>
    </div>
  )
}

export default Friends
