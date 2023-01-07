import { IRouter, Router } from "express";
import { loginController } from "../controllers/authentication";

const authentication: IRouter = Router();

authentication.post("/login", loginController);
authentication.post("/signup", loginController);

export default authentication;
