import { Schema, model } from "mongoose";

const User = new Schema(
  {
    username: String,
    display_name: String,
    email: String,
    password: String,
    avatar: String,
    facebook: String,
  },
  {
    timestamps: true,
    autoIndex: true,
  }
);

const user = model("users", User);
export { user as User };
