import { cookies } from "next/headers"

export const getCookies = async(key: string) => { 
    const cookieStore = cookies()
    const value = cookieStore.get(key)
    if (value)
        return value
    return null
}

export const setCookies = async(key: string, value: string) => { 
    const cookieStore = cookies()
    const newCookie = cookieStore.set(key,value)
    if (newCookie) 
        return newCookie
    return null
}