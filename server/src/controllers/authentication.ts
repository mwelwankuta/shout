import { Request, Response } from "express";

export async function loginController(req: Request, res: Response) {
  res.status(200).send("Hello Login");
}

export async function signUpController(req: Request, res: Response) {
  res.status(200).send("Hello Login");
}

export async function refreshController(req: Request, res: Response) {
  res.status(200).send("Hello Login");
}
