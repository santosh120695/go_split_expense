import { TrendingDown, Wallet } from "lucide-react"
import {useQuery} from "@tanstack/react-query";
import {fetchDashboardCount} from "../api/dashboard.ts";

function Home() {

  const { data: dashboardCounts } = useQuery<{you_owe: {currency: string, amount: number}, you_are_owed: {currency: string, amount: number}, activities: string[]}>({
        queryKey: ['dashboard_counts'],
        queryFn: fetchDashboardCount
  })

  return (
    <div className="min-h-screen bg-background p-8">
      <div className="max-w-7xl">
        <div className="mb-8">
          <p className="text-muted-foreground mt-2">Welcome back! Here's your expense summary</p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div className="bg-card bg-(--card) rounded-lg border border-[#C5C3C3] p-6 shadow-sm hover:shadow-md transition-shadow duration-200">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-muted-foreground font-medium">You owe</p>
                <p className="text-4xl font-bold text-foreground mt-2">{dashboardCounts?.you_owe?.currency}{dashboardCounts?.you_owe?.amount}</p>
                {/*<p className="text-sm text-green-600 mt-2">↓ 12% from last month</p>*/}
              </div>
              <div className="bg-red-100 dark:bg-red-900 p-4 rounded-lg">
                <TrendingDown size={32} className="text-red-600 dark:text-red-400" />
              </div>
            </div>
          </div>

          <div className="bg-card bg-(--card) rounded-lg border border-[#C5C3C3] p-6 shadow-sm hover:shadow-md transition-shadow duration-200">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-muted-foreground font-medium">You are owed</p>
                <p className="text-4xl font-bold text-foreground mt-2">{dashboardCounts?.you_are_owed?.currency}{dashboardCounts?.you_are_owed?.amount}</p>
                {/*<p className="text-sm text-blue-600 mt-2">↑ 8% from last month</p>*/}
              </div>
              <div className="bg-blue-100 dark:bg-blue-900 p-4 rounded-lg">
                <Wallet size={32} className="text-blue-600 dark:text-blue-400" />
              </div>
            </div>
          </div>
        </div>
      </div>
        <div className="mt-8 w-full bg-(--card)">
          <div className="bg-card rounded-lg border border-[#C5C3C3] p-6 shadow-sm hover:shadow-md transition-shadow duration-200">
            <h2 className="text-xl font-bold text-foreground mb-6">Activities</h2>
            <ul className="space-y-3">
              {(dashboardCounts?.activities || []).map((activity: string) => (
                <li className="flex items-center text-muted-foreground hover:text-foreground transition-colors duration-200 cursor-pointer border-b border-[#C5C3C3] pb-2">
                  <span className="inline-block w-2 h-2 bg-primary rounded-full mr-3"></span>
                  {activity}
                </li>
              ))}
            </ul>
          </div>
        </div>
    </div>
  );
}

export default Home;
