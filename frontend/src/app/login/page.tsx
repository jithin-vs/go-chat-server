"use client"
import SubmitButton from "@/components/Buttons/SubmitButton";
import TextInput from "@/components/Inputs/TextInput";
import Link from "next/link";
import { FormEvent } from "react";

const submitForm = (e:FormEvent<HTMLFormElement>) => {
  e.preventDefault();
  const formData = new FormData(e.currentTarget);
  console.log("formData", Object.fromEntries(formData.entries()));
  const username = formData.get("username")?.toString() || "";
  const password = formData.get("password")?.toString() || "";
  
}
export default function Login() {
  return (
    <div className="mt-12 max-w-lg py-6 rounded-xl mx-auto bg-white shadow-md">
  <div className="flex flex-col justify-center items-center py-6">
    <h1 className="text-3xl font-bold text-center text-gray-800">Login</h1>
  </div>
  <div className="max-w-md mx-auto p-4">
    <form onSubmit={submitForm}>
      <TextInput label ="Username" name="username" type ="text" isRequired = {true} placeholder="Enter your email or username" />
      <TextInput label ="Password" name="password" type="password" isRequired = {true} placeholder="Enter your password" />
      <SubmitButton Text="Login" />
    </form>
  </div>
  <div className="flex items-center justify-center mt-4 space-x-1 text-gray-600">
    <span>Don't have an account?</span>
    <Link href="/signup" className="text-blue-600 font-semibold hover:underline">
      Register here
    </Link>
  </div>
</div>
  )
}