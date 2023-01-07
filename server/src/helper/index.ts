import { Request, Response, NextFunction } from "express";
import { verify } from "jsonwebtoken";
import { AuthenticatedRequest } from "../types";

const secret = process.env.SECRET;

export async function checkAuthentication(
  req: AuthenticatedRequest,
  res: Response,
  next: NextFunction
) {
  const token = req.headers.authorization;
  const user = verify(token, secret);

  if (!user) {
    res.status(403).send({ message: "not authorized" });
    return;
  }

  req.user = user;
  next();
}
