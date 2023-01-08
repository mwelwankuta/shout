import { Request } from "express";
import { JwtPayload } from "jsonwebtoken";

export interface AuthenticatedRequest
  extends Request<never, never, { name: string; surname: string }, never> {
  user: string | JwtPayload;
}

export interface User {
  username: string;
  display_name: string;
  email: string;
  password?: string;
  avatar: string;
  facebook: string;
}

export interface Room {
  name: string;
  description: string;
  host: User;
  participants: User[];
}
