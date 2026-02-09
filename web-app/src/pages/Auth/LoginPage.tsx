import React, { useState } from "react";
import { useForm } from "react-hook-form";
import { Link, useNavigate } from "react-router-dom";
import api from "../../utils/api.ts"; // Import the Axios instance
import useAuthStore from "../../store/useStore.ts"; // Import the Zustand store

interface LoginFormData {
  email: string;
  password: string;
}

const LoginPage: React.FC = () => {
  const navigate = useNavigate();
  const setToken = useAuthStore((state) => state.setToken); // Get setToken from Zustand
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormData>({
    mode: "onChange",
  });

  const onSubmit = async (data: LoginFormData) => {
    setLoading(true);
    setError(null); // Clear previous errors
    try {
      const response = await api.post("/users/sign_in", {
        email: data.email,
        password: data.password,
      });

      const token = response.data.token;
      if (token) {
        setToken(token);
        navigate("/");
      } else {
        setError("Login failed: No token received.");
      }
    } catch (err: any) {
      setError(
        err.response?.data?.message ||
          err.message ||
          "An unexpected error occurred.",
      );
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-(--background) flex items-center justify-center px-4 sm:px-6 lg:px-8">
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="w-full max-w-md bg-(--card) rounded-lg shadow-xl p-8"
      >
        <h2 className="text-3xl font-bold text-(--card-foreground) mb-8 text-center">
          Login
        </h2>

        {error && (
          <div
            className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative mb-4"
            role="alert"
          >
            <span className="block sm:inline">{error}</span>
          </div>
        )}

        <div className="mb-6">
          <label
            htmlFor="email"
            className="block text-sm font-medium text-(--card-foreground) mb-2"
          >
            Email
          </label>
          <input
            type="email"
            id="email"
            {...register("email", {
              required: "Email is required",
              pattern: {
                value: /^[^\s@]+@[^\s@]+\.[^\s@]+$/,
                message: "Please enter a valid email",
              },
            })}
            className={`w-full px-4 py-2 border rounded-md focus:ring-2 focus:ring-(--primary) focus:border-transparent outline-none bg-transparent text-[var(--card-foreground)] transition ${
              errors.email ? "border-red-500" : "border-(--card-foreground)"
            }`}
            placeholder="Enter your email"
          />
          {errors.email && (
            <p className="text-red-500 text-sm mt-1">{errors.email.message}</p>
          )}
        </div>

        <div className="mb-6">
          <label
            htmlFor="password"
            className="block text-sm font-medium text-(--card-foreground) mb-2"
          >
            Password
          </label>
          <input
            type="password"
            id="password"
            {...register("password", {
              required: "Password is required",
              minLength: {
                value: 6,
                message: "Password must be at least 6 characters",
              },
            })}
            className={`w-full px-4 py-2 border text-(--card-foreground) rounded-md focus:ring-2 focus:ring-(--primary) focus:border-transparent outline-none bg-transparent transition ${
              errors.password ? "border-red-500" : "border-(--card-foreground)"
            }`}
            placeholder="Enter your password"
          />
          {errors.password && (
            <p className="text-red-500 text-sm mt-1">
              {errors.password.message}
            </p>
          )}
        </div>

        <button
          type="submit"
          className="w-full bg-(--primary) hover:opacity-90 text-white py-2 px-4 rounded-md transition duration-200 mb-6"
          disabled={loading}
        >
          {loading ? "Logging in..." : "Login"}
        </button>

        <div className="flex flex-col gap-2 text-center text-sm">
          <Link
            to="/forgot-password"
            className="text-(--primary) hover:underline font-medium transition"
          >
            Forgot Password?
          </Link>
          <p className="text-(--card-foreground)">
            Don't have an account?{" "}
            <Link
              to="/sign_up"
              className="text-(--primary) hover:underline font-medium transition"
            >
              Sign Up
            </Link>
          </p>
        </div>
      </form>
    </div>
  );
};

export default LoginPage;