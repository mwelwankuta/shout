import { Request } from "express";
import { JwtPayload } from "jsonwebtoken";

export interface AuthenticatedRequest
  extends Request<never, never, { name: string; surname: string }, never> {
  user: string | JwtPayload;
}
