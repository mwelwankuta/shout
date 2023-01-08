import express, { Application, Response, json } from "express";
import cors from "cors";
import http from "http";
import morgan from "morgan";
import socketIO from "socket.io";
import config from "./config";

import { authentication, rooms } from "./routes";
import { checkAuthentication } from "./helper";
import { AuthenticatedRequest } from "./types";
import conn from "./database";

const app: Application = express();
const server = http.createServer(app);

// @ts-ignore
const io = socketIO(server);

/**
 * starts the express server after the database connects
 */
// global middleware
app.use(
  morgan(":method :url :status :res[content-length] - :response-time ms")
);
app.use(cors({ origin: "*" }));
app.use(json());

// routes
app.use("/auth", authentication);
app.get(
  "/user",
  checkAuthentication,
  async (req: AuthenticatedRequest, res: Response) => {
    const user = await conn.execute("select * from users where id=?", [
      req.user,
    ]);
    res.send(user);
  }
);
app.use("/rooms", checkAuthentication, rooms);

server.listen(config.PORT, () => console.log("running on port " + config.PORT));

export { io as IO };
