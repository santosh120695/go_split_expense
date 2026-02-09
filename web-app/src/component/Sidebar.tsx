import React from "react"
import { Home, BarChart3, LogOut, Users } from "lucide-react"
import useAuthStore from "../store/useStore.ts";

const Sidebar = () => {
    const clearToken = useAuthStore(state => state.clearToken)
    const menuItems = [
        {
            id: 1,
            icon: Home,
            label: "Home",
            href: "/",
        },        {
            id: 2,
            icon: Users,
            label: "Groups",
            href: "/groups",
        },
        {
            id: 3,
            icon: BarChart3,
            label: "Friends",
            href: "/friends",
        },

    ]

    return (
        <React.Fragment>
            <div className="sidebar h-full bg-(--card) text-card-foreground  border-r border-[#C5C3C3] flex flex-col">
                <div className="border-b border-[#C5C3C3] bg-(--primary) text-white p-4 flex items-center justify-center">
                    <h1 className="text-2xl font-bold text-(--card)">SplitEx</h1>
                </div>

                <nav className="mt-8 px-4 flex-1">
                    <ul className="space-y-2">
                        {menuItems.map((item) => {
                            const Icon = item.icon
                            return (
                                <li key={item.id}>
                                    <a
                                        href={item.href}
                                        className="flex items-center gap-3 px-4 py-3 rounded-lg text-muted-foreground hover:text-foreground hover:bg-accent transition-colors duration-200"
                                    >
                                        <Icon size={20} className="shrink-0" />
                                        <span className="text-sm font-medium">{item.label}</span>
                                    </a>
                                </li>
                            )
                        })}
                    </ul>
                </nav>

                <div className="border-t border-[#C5C3C3] p-4">
                    <button
                        onClick={() => {
                            clearToken()
                        }}
                        className="w-full flex items-center gap-3 px-4 py-3 rounded-lg text-muted-foreground hover:text-foreground hover:bg-red-100 dark:hover:bg-red-900/30 transition-colors duration-200"
                    >
                        <LogOut size={20} className="shrink-0" />
                        <span className="text-sm font-medium">Logout</span>
                    </button>
                </div>
            </div>
        </React.Fragment>
    )
}

export default Sidebar;