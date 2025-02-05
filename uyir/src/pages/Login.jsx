import React, { useState } from "react";
import { FormInput } from "../components/FormInput";
import { SubmitButton } from "../components/SubmitButton";
import { useNavigate } from "react-router-dom"; // For navigation after login
import styles from "./SignUp.module.css";

export const Login = () => {
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
        credentials: "include", // Important for cookies!
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
      navigate("/user"); // Redirect after login
    } catch (error) {
      console.error("Login error:", error);
      setErrorMessage(error.message);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <main className={styles.signUpContainer}>
      <form className={styles.formWrapper} onSubmit={handleSubmit} noValidate>
        <h1 className={styles.title}>Login</h1>

        {errorMessage && <p className={styles.error}>{errorMessage}</p>}

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
          autoComplete="new-password"
          required
        />

        <SubmitButton text="Login" isLoading={isLoading} />
      </form>
    </main>
  );
};
