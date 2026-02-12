import { create } from "zustand";
import type {UserType} from "../types/user.ts";
import { fetchCurrentUser } from "../api/user.ts";

interface AuthState {
  token: string | null;
  isAuthenticated: boolean;
  user: UserType | null;
}

interface AuthActions {
  setToken: (token: string | null) => void;
  clearToken: () => void;
  setUser: (user: UserType | null) => void;
  initializeAuth: () => Promise<void>;
}

const useAuthStore = create<AuthState & AuthActions>((set, get) => ({
  token: localStorage.getItem("authToken") || null,
  isAuthenticated: !!localStorage.getItem("authToken"),
  user: null,

  setToken: async (token) => {
    if (token) {
      localStorage.setItem("authToken", token);
      let fetchedUser: UserType | null = null;
      try {
        fetchedUser = await fetchCurrentUser();
      } catch (error) {
        console.error("Failed to fetch current user:", error);
      }
      set({ token, isAuthenticated: !!token, user: fetchedUser });
    } else {
      localStorage.removeItem("authToken");
      set({ token: null, isAuthenticated: false, user: null });
    }
  },

  clearToken: () => {
    localStorage.removeItem("authToken");
    set({ token: null, isAuthenticated: false, user: null });
  },
  setUser: (user) => set({ user }),

  initializeAuth: async () => {
    const token = localStorage.getItem("authToken");
    if (token != null && !get().user) { // Only fetch if token exists and user is not already set
      let fetchedUser: UserType | null = null;
      try {
        fetchedUser = await fetchCurrentUser();
      } catch (error) {
        console.error("Failed to fetch current user on initialization:", error);
      }
      set({ token, isAuthenticated: !!token, user: fetchedUser });
    }
  },
}));

useAuthStore.getState().initializeAuth();

interface SidebarState {
    isOpen: boolean;
    toggle: () => void;
}

export const useSidebarStore = create<SidebarState>((set) => ({
    isOpen: false,
    toggle: () => set((state) => ({ isOpen: !state.isOpen })),
}));

export default useAuthStore;