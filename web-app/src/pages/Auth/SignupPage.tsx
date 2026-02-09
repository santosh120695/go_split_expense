import React, { useState } from "react";
import { useForm } from "react-hook-form";
import { Link, useNavigate } from "react-router-dom";
import api from "../../utils/api.ts"; // Import the Axios instance

interface SignupFormData {
  user_name: string;
  email: string;
  password: string;
  confirm_password: string;
  contact_no: string;
}

const SignupPage: React.FC = () => {
  const navigate = useNavigate();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<SignupFormData>({
    mode: "onChange",
  });

  const password = watch("password");

  const onSubmit = async (data: SignupFormData) => {
    setLoading(true);
    setError(null);
    setSuccess(null);
    try {
      const response = await api.post("/users/sign_up", {
        user_name: data.user_name,
        email: data.email,
        password: data.password,
        contact_no: data.contact_no,
        confirm_password: data.confirm_password,
      });

      if (response.status === 201 || response.status === 200) {
        setSuccess("Account created successfully! Please log in.");
        setTimeout(() => {
          navigate("/sign_in");
        }, 2000);
      } else {
        setError("Signup failed: Unexpected response from server.");
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
          Sign Up
        </h2>

        {error && (
          <div
            className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative mb-4"
            role="alert"
          >
            <span className="block sm:inline">{error}</span>
          </div>
        )}
        {success && (
          <div
            className="bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded relative mb-4"
            role="alert"
          >
            <span className="block sm:inline">{success}</span>
          </div>
        )}

        <div className="mb-6">
          <label
            htmlFor="user_name"
            className="block text-sm font-medium text-(--card-foreground) mb-2"
          >
            Name
          </label>
          <input
            type="text"
            id="name"
            {...register("user_name", {
              required: "Name is required",
              minLength: {
                value: 2,
                message: "Name must be at least 2 characters",
              },
            })}
            className={`w-full px-4 py-2 border rounded-md focus:ring-2 focus:ring-(--primary) focus:border-transparent outline-none bg-transparent text-[var(--card-foreground)] transition ${
              errors.user_name ? "border-red-500" : "border-(--card-foreground)"
            }`}
            placeholder="Enter your full name"
          />
          {errors.user_name && (
            <p className="text-red-500 text-sm mt-1">
              {errors.user_name.message}
            </p>
          )}
        </div>

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
            htmlFor="contact_no"
            className="block text-sm font-medium text-(--card-foreground) mb-2"
          >
            Contact No
          </label>
          <input
            type="text"
            id="contact_no"
            {...register("contact_no", {
              required: "Contact number is required",
            })}
            className={`w-full px-4 py-2 border rounded-md focus:ring-2 focus:ring-(--primary) focus:border-transparent outline-none bg-transparent text-[var(--card-foreground)] transition ${
              errors.contact_no ? "border-red-500" : "border-(--card-foreground)"
            }`}
            placeholder="Enter your contact number"
          />
          {errors.contact_no && (
            <p className="text-red-500 text-sm mt-1">
              {errors.contact_no.message}
            </p>
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
            className={`w-full px-4 py-2 border rounded-md focus:ring-2 focus:ring-(--primary) focus:border-transparent outline-none bg-transparent text-[var(--card-foreground)] transition ${
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

        <div className="mb-6">
          <label
            htmlFor="confirm_password"
            className="block text-sm font-medium text-(--card-foreground) mb-2"
          >
            Confirm Password
          </label>
          <input
            type="password"
            id="confirm_password"
            {...register("confirm_password", {
              required: "Please confirm your password",
              validate: (value) =>
                value === password || "Passwords do not match",
            })}
            className={`w-full px-4 py-2 border rounded-md focus:ring-2 focus:ring-(--primary) focus:border-transparent outline-none bg-transparent text-[var(--card-foreground)] transition ${
              errors.confirm_password ? "border-red-500" : "border-(--card-foreground)"
            }`}
            placeholder="Confirm your password"
          />
          {errors.confirm_password && (
            <p className="text-red-500 text-sm mt-1">
              {errors.confirm_password.message}
            </p>
          )}
        </div>

        <button
          type="submit"
          className="w-full bg-(--primary) hover:opacity-90 text-white font-semibold py-2 px-4 rounded-md transition duration-200 mb-6"
          disabled={loading}
        >
          {loading ? "Signing Up..." : "Sign Up"}
        </button>

        <div className="text-center text-sm">
          <p className="text-(--card-foreground)">
            Already have an account?{" "}
            <Link
              to="/sign_in"
              className="text-(--primary) hover:underline font-medium transition"
            >
              Login
            </Link>
          </p>
        </div>
      </form>
    </div>
  );
};

export default SignupPage;