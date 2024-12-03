// AuthContext.tsx
import React, { createContext, useState, useContext, ReactNode } from 'react';

interface AuthContextType {
  token: string | null;
  setToken: (token: string | null) => void;
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
  const value = {
    token,
    setToken,
  }
  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
};


// Export the AuthContext for manual access if needed
export { AuthContext };
