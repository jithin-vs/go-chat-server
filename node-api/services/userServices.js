import User from "../models/user.model.js";
import bcrypt from "bcrypt";

const loginUser = async ({ email,password }) => {
  console.log("email", email);
  try {
    const user = await User.findOne({email:email});
    console.log("result",user)
    if (!user) {
      throw new Error("user not found");
    }
  
    const isValid = await bcrypt.compare(password,user.password);
    if (!isValid) {
      throw new Error("incorrect password");
    }
  
    return {
      id: user._id,
      username: user.username,
      email: user.email,
    };
    
  } catch (error) {
    throw new Error("internal error: " + error.message);
  }
};

const registerUser = async ({ name, username, email, password }) => {
  try {   
    const hashedPassword = await bcrypt.hash(password, 14);
  
    const newUser = new User({
      name,
      username,
      email,
      password: hashedPassword,
    });
  
    const result = await newUser.save();
  
    return {
      id: result._id,
      username: result.username,
      email: result.email,
    };
  } catch (error) {
    throw new Error("internal error: " + error.message);
  }
};

export default {
  loginUser,
  registerUser,
};
