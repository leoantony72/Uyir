import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { FormInput } from "../components/FormInput";
import { SubmitButton } from "../components/SubmitButton";

const Login = () => {
  const [formData, setFormData] = useState({
    username: "",
    password: "",
  });
  const [isLoading, setIsLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const navigate = useNavigate();

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    setErrorMessage("");

    try {
      const response = await fetch("http://localhost:6969/login", {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        const result = await response.json();
        throw new Error(result.error || "Login failed");
      }

      console.log("Login successful");
      navigate("/user");
    } catch (error) {
      console.error("Login error:", error);
      setErrorMessage("Enter your Username and Password");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <main className="min-h-screen flex items-center justify-center bg-[var(--secondary-color)] p-4">
      <div className="card w-full max-w-md">
        <div className="mb-7 text-center"> 
          <h1 className="text-4xl font-bold mt-5 text-3d">
            <span className="text-[var(--primary-color)]">Uyir</span>
            <span className="text-[var(--red-color)]">Safe</span>
          </h1>
          <h2 className="text-3xl font-semibold text-[var(--primary-color)] mt-5">
            Login
          </h2>
        </div>

        {errorMessage && (
          <p className="text-red-500 text-sm text-center mb-6" role="alert">
            {errorMessage}
          </p>
        )}

        <form onSubmit={handleSubmit} noValidate className="space-y-6">
          <FormInput
            label="Username"
            type="text"
            id="username"
            name="username"
            value={formData.username}
            onChange={handleChange}
            autoComplete="username"
            required
          />
          <FormInput
            label="Password"
            type="password"
            id="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            autoComplete="current-password"
            required
          />
          <SubmitButton text="Login" isLoading={isLoading} />
        </form>

        <p className="mt-4 text-center text-sm text-gray-600">
           Don't have an account?{" "}
            <Link to="/signup" className="text-red-500 hover:underline">
               Sign up
            </Link>
        </p>

        <p className="mt-6 text-center text-sm text-gray-500">
          Coimbatore City | UyirSafe © 2025
        </p>
      </div>
    </main>
  );
};

export default Login;