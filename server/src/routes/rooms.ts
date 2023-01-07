import { IRouter, Router } from "express";
import { loginController } from "../controllers/authentication";

const rooms: IRouter = Router();

rooms.get("/", loginController);

export default rooms;
