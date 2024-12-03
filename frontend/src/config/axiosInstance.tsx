import axios from "axios"
import { PORT } from "./index";
import { getCookies, setCookies } from "@/hooks/cookies";

const axiosInstance = axios.create({
    baseURL: `${PORT}`, 
    timeout: 1000,
    headers: { 'Content-Type': 'application/json' }
});
  
axiosInstance.interceptors.request.use(
    function (config) {
    const token = getCookies('accessToken');
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
      return config;
    },
    function (error) {
      return Promise.reject(error);
    }
);
  
axiosInstance.interceptors.response.use(
    function (response) {
    if (response.status === 401) {
        const refreshToken = getCookies('refreshToken');
        axios.post(`${PORT}/auth/refresh-token`, {
            headers: {
                'Content-Type': 'application/json'
            }
          })
         .then((resp) => {
             if (resp.status === 200) {
                 console.log("User signed up successfully:", resp.data);
                 const accessToken = resp.data.access_token;
                 const refreshToken = resp.data.refresh_token;
                 setCookies('accessToken', accessToken)
             } else {
                 console.error("request failed")
             }
            })
        .catch((error) => {
              console.error("Error signing up user:", error.resp.data.error);
            });
    }
      console.log('resp:', response);
      return response;
    },
    function (error) {
      // Handle the response error
      if (error.response && error.response.status === 401) {
        // Handle unauthorized error
        console.error('Unauthorized, logging out...');
        // Perform any logout actions or redirect to login page
      }
      return Promise.reject(error);
    }
  );
  
  
  export default axiosInstance;