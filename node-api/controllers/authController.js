import userService from "../services/userServices.js";
import { sendResponse, sendErrorResponse } from "../utils/responseUtils.js";

export const loginHandler = async (req, res) => {
    try {
    const user = req.body;
    console.log("user", user)
    // Call the service to log in the user
    const result = await userService.loginUser(user);

    if (!result) {
      return sendErrorResponse(res, 401, "Invalid credentials");
    }

    // Example response structure
    const response = {
      message: "User logged in successfully",
      data: result,
    };

    sendResponse(res, 200, response);
    } catch (error) {
     console.log("err",error)
    if (error.message === "user not found") {
      sendErrorResponse(res, 404, "User does not exist");
    } else if (error.message === "incorrect password") {
      sendErrorResponse(res, 401, "Invalid credentials");
    } else {
        console.log(error)
      sendErrorResponse(res, 500, "Something went wrong");
    }
  }
};

export const registerHandler = async (req, res) => {

  try {
    const user = req.body;

    // Call the service to register the user
    const result = await userService.registerUser(user);

    const response = {
      message: "User registered successfully",
      data: result,
    };

    sendResponse(res, 201, response);
  } catch (error) {
      console.log("error",error)
    sendErrorResponse(res, 500, "Failed to register user");
  }
};
