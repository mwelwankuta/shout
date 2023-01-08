import { IRouter, Router } from "express";
import { roomsController } from "../controllers/rooms";

const rooms: IRouter = Router();

rooms.get("/", roomsController);

export default rooms;
