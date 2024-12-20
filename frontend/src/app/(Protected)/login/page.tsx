"use client"
import SubmitButton from "@/components/Buttons/SubmitButton";
import TextInput from "@/components/Inputs/TextInput";
import { PORT } from "@/config/index";
import { useAuth } from "@/context/AuthContext";
import axios from "axios";
import Link from "next/link";
import { useRouter } from "next/navigation"
import { FormEvent, useEffect, useState } from "react";


export default function Login() {
  
  const {user,setUser}=useAuth()
  const router = useRouter();
  const [errors, setErrors] = useState<string[]>([]);
  useEffect(() => {
    console.log('User state changed:', user)
  }, [user])

  const submitForm = (e:FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    console.log("formData", Object.fromEntries(formData.entries()));
    const email = formData.get("email")?.toString() || "";
    const password = formData.get("password")?.toString() || "";
  
    const apiData = {
      email,
      password,
    }
    // Perform API call to sign up user
    axios.post(`${PORT}/api/login`, apiData, {
      headers: {
          'Content-Type': 'application/json'
      }
  })
      .then((response) => {
        if(response.status === 200) {
          console.log("User logged in successfully:", response.data.data);
          setUser(response.data.data)
          router.push('/home')
          return
        }
        console.log("User signed up successfully:", response.data);
      })
     .catch((error) => {
        console.error("Error signing up user:", error);
        setErrors([error.response.data.error]);
      });
    
  }
  return (
    <div className="mt-12 max-w-lg py-6 rounded-xl mx-auto bg-white shadow-md">
  <div className="flex flex-col justify-center items-center py-6">
    <h1 className="text-3xl font-bold text-center text-gray-800">Login</h1>
  </div>
  <div className="max-w-md mx-auto p-4">
    <form onSubmit={submitForm}>
      <TextInput label ="Username" name="email" type ="text" isRequired = {true} placeholder="Enter your email" />
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