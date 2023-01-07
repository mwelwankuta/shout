import { Schema, model } from "mongoose";

const User = new Schema(
  {
    username: {
      type: String,
      default: "",
    },
    display_name: {
      type: String,
      default: "",
    },
    email: {
      type: String,
      default: "",
    },
    password: {
      type: String,
      default: "",
    },
    avatar: {
      type: String,
      default: "",
    },
  },
  {
    timestamps: true,
    autoIndex: true,
  }
);

const user = model("users", User);
export { user as User };
