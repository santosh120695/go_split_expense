import React from "react";
import { Navigate } from "react-router-dom";
import useAuthStore from "../store/useStore";
import Layout from "./Layout.tsx";

interface ProtectedRouteProps {
  children: React.ReactNode;
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ children }) => {
  const { isAuthenticated } = useAuthStore();

  if (!isAuthenticated) {
    return <Navigate to="/sign_in" replace />;
  }

  return (
    <>
      <Layout>
        {children}
      </Layout>
    </>
  );
};

export default ProtectedRoute;
