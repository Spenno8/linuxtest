import { create } from "zustand"
import { persist } from "zustand/middleware"

// User has a specific email
type User = {
  email: string
}

// User that has passed login has their token saved and is Authenticated
type AuthState = {
  user: User | null
  token: string | null
  isAuthenticated: boolean
  login: (user: User, token: string) => void
  logout: () => void
}

export const useAuthStore = create<AuthState>()(
    persist(
    // Initial State, no user, no token and not authenticated
    (set) => ({
        user: null,
        token: null,
        isAuthenticated: false,
    
      // Login Action uses the user and token to save the state in local storage
        login: (user, token) =>
            set({
            user,
            token,
            isAuthenticated: true,
        }),

        // Logout Action to null user and token
        logout: () =>
            set({
            user: null,
            token: null,
            isAuthenticated: false,
            }),
    }),
    {
        name: "auth-storage",
    }
  )
)