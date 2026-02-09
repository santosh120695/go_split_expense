import { create } from "zustand";

interface AuthState {
  token: string | null;
  isAuthenticated: boolean;
}

interface AuthActions {
  setToken: (token: string | null) => void;
  clearToken: () => void;
}

const useAuthStore = create<AuthState & AuthActions>((set) => ({
  token: localStorage.getItem("authToken") || null,
  isAuthenticated: !!localStorage.getItem("authToken"),

  setToken: (token) => {
    if (token) {
      localStorage.setItem("authToken", token);
    } else {
      localStorage.removeItem("authToken");
    }
    set({ token, isAuthenticated: !!token });
  },

  clearToken: () => {
    localStorage.removeItem("authToken");
    set({ token: null, isAuthenticated: false });
  },
}));

export default useAuthStore;