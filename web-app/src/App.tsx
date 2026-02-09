import { Routes, Route } from "react-router-dom";
import Home from "./pages/Home.tsx";
import About from "./component/About";
import Index from "./pages/Groups";
import GroupDetails from "./pages/Groups/Show";
import Friends from "./pages/Friends";
import LoginPage from "./pages/Auth/LoginPage";
import SignupPage from "./pages/Auth/SignupPage";
import ProtectedRoute from "./component/ProtectedRoute";
import "./App.css";

function App() {
  return (
    <div className="App">
      <Routes>
        <Route
          path="/"
          element={
            <ProtectedRoute>
              <Home />
            </ProtectedRoute>
          }
        />
        <Route
          path="/about"
          element={
            <ProtectedRoute>
              <About />
            </ProtectedRoute>
          }
        />
        <Route
          path="/groups"
          element={
            <ProtectedRoute>
              <Index />
            </ProtectedRoute>
          }
        />
        <Route
          path="/groups/:id"
          element={
            <ProtectedRoute>
              <GroupDetails />
            </ProtectedRoute>
          }
        />
        <Route
          path="/friends"
          element={
            <ProtectedRoute>
              <Friends />
            </ProtectedRoute>
          }
        />
        <Route path="/sign_in" element={<LoginPage />} />
        <Route path="/sign_up" element={<SignupPage />} />
      </Routes>
    </div>
  );
}

export default App;

