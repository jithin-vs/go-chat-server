"use client";

import SubmitButton from "@/components/Buttons/SubmitButton";
import TextInput from "@/components/Inputs/TextInput";
import Link from "next/link";
import { FormEvent, useState } from "react";
import axios from "axios";
import { useRouter } from "next/navigation"
import { PORT } from "@/config/index";

export default function SignUp() {
  const router = useRouter();
  const [errors, setErrors] = useState<string[]>([]);
  const submitForm = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log("env",PORT,process.env.SERVER_URL)
    const formData = new FormData(e.currentTarget);
    const regex =
      /^[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?$/;
    console.log("formData", Object.fromEntries(formData.entries()));
    const username = formData.get("username")?.toString() || "";
    const email = formData.get("email")?.toString() || "";
    if (!regex.test(email)) {
      setErrors(["Invalid email format"]);
      return;
    }
    const password = formData.get("password")?.toString() || "";
    const confirmPassword = formData.get("re-password")?.toString() || "";
    if (password !== confirmPassword) {
      setErrors(["Passwords do not match"]);
      return;
    }
    const apiData = {
      username,
      email,
      password,
    }
    // Perform API call to sign up user
    axios.post(`${PORT}/api/signup`, apiData, {
      headers: {
          'Content-Type': 'application/json'
      }
  })
    .then((response) => {
        console.log("User signed up successfully:", response.data);
        router.push('/login')
      })
     .catch((error) => {
        console.error("Error signing up user:", error.response.data.error);
        setErrors([error.response.data.error]);
      });
  };

  return (
    <div className="mt-6 max-w-lg py-6 rounded-xl mx-auto bg-white shadow-md">
      <div className="flex flex-col justify-center items-center py-6">
        <h1 className="text-3xl font-bold text-center text-gray-800">
          Sign Up
        </h1>
      </div>
      <div className="max-w-md mx-auto p-4">
        <form onSubmit={submitForm}>
          <TextInput
            label="Username"
            name="username"
            type="text"
            isRequired={true}
            placeholder="Enter your username"
          />
          <TextInput
            label="Email"
            name="email"
            type="email"
            isRequired={true}
            placeholder="Enter your email"
          />
          <TextInput
            label="Password"
            name="password"
            type="password"
            isRequired={true}
            placeholder="Enter your password"
          />
          <TextInput
            label="Confirm Password"
            name="re-password"
            type="password"
            isRequired={true}
            placeholder="Re-enter your password"
          />
          {errors.length > 0 && (
            <div className="flex items-center justify-start mt-2 space-x-1 text-red-700">
              <span>{"* "}{errors}</span>
            </div>
          )}
          <SubmitButton Text="Register" />
        </form>
      </div>
      <div className="flex items-center justify-center mt-4 space-x-1 text-gray-600">
        <span>Already have an account?</span>
        <Link
          href="/login"
          className="text-blue-600 font-semibold hover:underline"
        >
          Log In
        </Link>
      </div>
    </div>
  );
}
