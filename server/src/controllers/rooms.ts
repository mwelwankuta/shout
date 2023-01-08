import { Request, Response } from "express";

export async function roomsController(req: Request, res: Response) {
  res.status(200).send("Hello Room");
}
