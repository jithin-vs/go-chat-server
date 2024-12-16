import API from "@/config/axiosInstance";

export const acessCreate = async (body:any) => {
    try {
  
      const  data  = await API.post(`/api/chat`, body);
      console.log(data);
      return data;
    } catch (error) {
      console.log('error in access create api');
    }
};
  
export const fetchAllChats = async () => {
    try {
      const { data } = await API.get('/api/chat');
      return data;
    } catch (error) {
      console.log('error in fetch all chats api');
    }
  };