import mongoose from "mongoose";

const userSchema = new mongoose.Schema(
  {
    id: {
      type: String,
      required: false,
    },
    name: {
      type: String,
      // required: [true, "Name is required"],
      minlength: [3, "Name must be at least 3 characters"],
      maxlength: [50, "Name must be at most 50 characters"],
    },
    username: {
      type: String,
      required: [true, "Username is required"],
      minlength: [3, "Username must be at least 3 characters"],
      maxlength: [50, "Username must be at most 50 characters"],
    },
    email: {
      type: String,
      required: [true, "Email is required"],
    },
    password: {
      type: String,
      required: [true, "Password is required"],
      minlength: [4, "Password must be at least 4 characters"],
    },
  },
  {
    timestamps: true, // Automatically manage `createdAt` and `updatedAt` fields
  }
);

const User = mongoose.model("User", userSchema);

export default User;
