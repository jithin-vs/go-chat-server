import { sendErrorResponse, sendResponse } from "../utils/responseUtils.js";
import userService from "../services/userServices.js";


export const findUser = async (req, res) => {
    try {
      const query  = req.query.query; 
      console.log("Query:", req.query);
  
      if (!query || Object.keys(query).length === 0) {
        return sendErrorResponse(res, 400, "Query parameters are required");
      }
      // Query the user from the database
      const user = await userService.findUserByName(query);
  
      if (!user) {
        return sendErrorResponse(res, 404, "User not found");
      }
        
      console.log("Result:", user);
  
      // Return the found user
      const response = {
        message: "User retrieved successfully",
        results: user,
      };
  
        sendResponse(res, 200, response);
        
    } catch (error) {
      console.error("Error in findUser handler:", error.message);
      sendErrorResponse(res, 500, "Internal server error");
    }
  };
  