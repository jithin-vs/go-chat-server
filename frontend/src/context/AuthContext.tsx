// AuthContext.tsx
import { User } from '@/interfaces/userInterface';
import React, { createContext, useState, useContext, ReactNode } from 'react';

interface AuthContextType {
  token: string | null;
  setToken: (token: string | null) => void;
  user: User | null; // User details or null if no user is logged in
  setUser: (user: User | null) => void; // Function to update user details
}

// Provide a default value for the context
const AuthContext = createContext<AuthContextType | undefined>(undefined);

interface AuthProviderProps {
  children: ReactNode;
}

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

export const AuthProvider = ({ children }: AuthProviderProps) => {
  const [token, setToken] = useState<string | null>(null);
  const [user,setUser] = useState<User | null>(null);
  const value = {
    token,
    setToken,
    user,
    setUser,
  }
  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
};


// Export the AuthContext for manual access if needed
export { AuthContext };
