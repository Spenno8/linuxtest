import { create } from "zustand"
import { persist } from "zustand/middleware"

// User represents the authenticated user's basic identity
// (non-sensitive data only)
type User = {
  id: string
  email: string
  username: string
}

// AuthState defines the shape of the authentication store.
// It tracks the current user, JWT token, auth status,
// and exposes actions to log in and log out.
type AuthState = {
  user: User | null           // Logged-in user details
  token: string | null        // JWT token returned from backend
  isAuthenticated: boolean    // Quick flag for auth checks

  // Action to store user data and token after successful login
  login: (user: User, token: string) => void

  // Action to clear authentication state on logout
  logout: () => void
}

// useAuthStore is a global Zustand store for authentication.
// The persist middleware saves auth state to localStorage
// so the user remains logged in after a page refresh.
export const useAuthStore = create<AuthState>()(
    persist(
    // Initial state and state-modifying actions
    (set) => ({
      // Default state: no user, no token, not authenticated
        user: null,
        token: null,
        isAuthenticated: false,
    
        // Login action:
        // Saves user info and JWT token,
        // and marks the user as authenticated
        login: (user, token) =>
            set({
            user,
            token,
            isAuthenticated: true,
        }),

        // Logout action:
        // Clears user info and token,
        // and resets authentication status
        logout: () =>
            set({
            user: null,
            token: null,
            isAuthenticated: false,
            }),
    }),

    // Persist configuration:
    // Store auth data in localStorage under this key
    {
        name: "auth-storage",
    }
  )
)