import React, { useState } from "react";
import { FormInput } from "../components/FormInput";
import { SubmitButton } from "../components/SubmitButton";
import styles from "./SignUp.module.css";

export const SignUp = () => {
  const [formData, setFormData] = useState({
    username: "",
    password: "",
  });
  const [isLoading, setIsLoading] = useState(false);

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

    try {
      await new Promise((resolve) => setTimeout(resolve, 1000));
      console.log("Form submitted:", formData);
    } catch (error) {
      console.error("Submission error:", error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <main className={styles.signUpContainer}>
      <form className={styles.formWrapper} onSubmit={handleSubmit} noValidate>
        <h1 className={styles.title}>SignUp</h1>

        <FormInput
          label="Username"
          type="text"
          id="username"
          icon="https://cdn.builder.io/api/v1/image/assets/TEMP/787c978d49e724e5265f4004c67c48feb92e5323a668be0a95af380324d24578?placeholderIfAbsent=true&apiKey=2fc17400dcd74914b50bcc9d036de5cf"
          value={formData.username}
          onChange={handleChange}
          autoComplete="username"
        />

        <FormInput
          label="Password"
          type="password"
          id="password"
          icon="https://cdn.builder.io/api/v1/image/assets/TEMP/787c978d49e724e5265f4004c67c48feb92e5323a668be0a95af380324d24578?placeholderIfAbsent=true&apiKey=2fc17400dcd74914b50bcc9d036de5cf"
          value={formData.password}
          onChange={handleChange}
          autoComplete="new-password"
        />

        <SubmitButton
          text="Register"
          backgroundImage="https://cdn.builder.io/api/v1/image/assets/TEMP/0c35efe22f1408ff2bb603ae6f3950018e7d92be6a49165c0ce4635b4425268b?placeholderIfAbsent=true&apiKey=2fc17400dcd74914b50bcc9d036de5cf"
          isLoading={isLoading}
        />
      </form>
    </main>
  );
};
