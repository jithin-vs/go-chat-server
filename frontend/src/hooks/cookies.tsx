// Use browser-native methods or libraries like js-cookie
import Cookies from 'js-cookie'


export const getCookies = async(key: string) => { 
    const value = Cookies.get(key)
    if (value)
        return value
    return null
}

export const setCookies = async(key: string, value: string) => { 
    const newCookie = Cookies.set(key,value)
    if (newCookie) 
        return newCookie
    return null
}