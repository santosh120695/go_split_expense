import { X, Wallet, TrendingUp } from "lucide-react"

interface Activity {
  id: number
  description: string
  amount: string
  user: string
  date: string
}

interface Settlement {
  from: string
  to: string
  amount: string
}

interface GroupDetailsModalProps {
  isOpen: boolean
  onClose: () => void
  group: {
    id: number
    name: string
    description: string
    members: number
    totalExpense: string
  } | null
}

function GroupDetailsModal({ isOpen, onClose, group }: GroupDetailsModalProps) {
  if (!isOpen || !group) return null

  // Mock activities data
  const activities: Activity[] = [
    {
      id: 1,
      description: "Dinner at restaurant",
      amount: "$50.00",
      user: "John Doe",
      date: "2024-02-03",
    },
    {
      id: 2,
      description: "Groceries shopping",
      amount: "$85.30",
      user: "Jane Smith",
      date: "2024-02-02",
    },
    {
      id: 3,
      description: "Movie tickets",
      amount: "$30.00",
      user: "Mike Johnson",
      date: "2024-02-01",
    },
    {
      id: 4,
      description: "Pizza delivery",
      amount: "$45.50",
      user: "Sarah Williams",
      date: "2024-01-31",
    },
  ]

  // Mock settlement data
  const settlements: Settlement[] = [
    { from: "John Doe", to: "Jane Smith", amount: "$12.50" },
    { from: "Mike Johnson", to: "Sarah Williams", amount: "$8.75" },
    { from: "Tom Brown", to: "Jane Smith", amount: "$15.00" },
  ]

  return (
    <>
      {isOpen && (
        <div
          className="fixed inset-0 bg-black bg-opacity-50 z-40 transition-opacity duration-200"
          onClick={onClose}
        />
      )}
      {isOpen && (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4 pointer-events-none">
          <div
            className="bg-card rounded-lg border border-[#C5C3C3] p-8 w-full max-w-2xl shadow-lg max-h-[90vh] overflow-y-auto pointer-events-auto"
            onClick={(e) => e.stopPropagation()}
          >
            {/* Modal Header */}
            <div className="flex items-center justify-between mb-6">
              <div>
                <h2 className="text-3xl font-bold text-foreground">{group.name}</h2>
                <p className="text-muted-foreground text-sm mt-1">{group.description}</p>
              </div>
              <button
                onClick={onClose}
                className="p-1 hover:bg-accent rounded-lg transition-colors duration-200"
              >
                <X size={24} className="text-muted-foreground" />
              </button>
            </div>

            {/* Group Summary */}
            <div className="grid grid-cols-2 gap-4 mb-6">
              <div className="bg-background border border-[#C5C3C3] rounded-lg p-4">
                <p className="text-sm text-muted-foreground mb-1">Total Members</p>
                <p className="text-2xl font-bold text-foreground">{group.members}</p>
              </div>
              <div className="bg-background border border-[#C5C3C3] rounded-lg p-4">
                <p className="text-sm text-muted-foreground mb-1">Total Expenses</p>
                <p className="text-2xl font-bold text-foreground">{group.totalExpense}</p>
              </div>
            </div>

            {/* Activities Section */}
            <div className="mb-8">
              <h3 className="text-xl font-bold text-foreground mb-4 flex items-center gap-2">
                <TrendingUp size={20} />
                Activities
              </h3>
              <div className="space-y-3 max-h-48 overflow-y-auto">
                {activities.map((activity) => (
                  <div
                    key={activity.id}
                    className="border border-[#C5C3C3] rounded-lg p-4 hover:bg-background transition-colors duration-200"
                  >
                    <div className="flex items-center justify-between mb-2">
                      <p className="font-medium text-foreground">{activity.description}</p>
                      <span className="text-lg font-bold text-primary">{activity.amount}</span>
                    </div>
                    <div className="flex items-center justify-between text-sm text-muted-foreground">
                      <span>{activity.user}</span>
                      <span>{activity.date}</span>
                    </div>
                  </div>
                ))}
              </div>
            </div>

            {/* Settlement Summary Section */}
            <div className="mb-8">
              <h3 className="text-xl font-bold text-foreground mb-4 flex items-center gap-2">
                <Wallet size={20} />
                Who Owes What
              </h3>
              <div className="space-y-3 max-h-48 overflow-y-auto">
                {settlements.map((settlement, index) => (
                  <div
                    key={index}
                    className="border border-[#C5C3C3] rounded-lg p-4 hover:bg-background transition-colors duration-200"
                  >
                    <div className="flex items-center justify-between">
                      <div>
                        <p className="text-foreground font-medium">{settlement.from}</p>
                        <p className="text-sm text-muted-foreground">owes</p>
                      </div>
                      <div className="text-right">
                        <p className="text-foreground font-medium">{settlement.to}</p>
                      </div>
                      <div className="ml-4">
                        <span className="text-xl font-bold text-red-600 dark:text-red-400">
                          {settlement.amount}
                        </span>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            </div>

            {/* Footer Actions */}
            <div className="flex gap-3 pt-4 border-t border-[#C5C3C3]">
              <button
                onClick={onClose}
                className="flex-1 px-4 py-2 bg-accent text-foreground rounded-lg hover:bg-opacity-80 transition-colors duration-200 font-medium"
              >
                Close
              </button>
              <button className="flex-1 px-4 py-2 bg-primary text-white rounded-lg hover:bg-opacity-90 transition-all duration-200 font-medium">
                Settle Up
              </button>
            </div>
          </div>
        </div>
      )}
    </>
  )
}

export default GroupDetailsModal
