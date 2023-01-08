import express, { Application, Response, json } from "express";
import cors from "cors";
import mongoose from "mongoose";
import http from "http";
import morgan from "morgan";
import socketIO from "socket.io";
import config from "./config";

import { authentication, rooms } from "./routes";
import { checkAuthentication } from "./helper";
import { User } from "./models";
import { AuthenticatedRequest } from "./types";

mongoose.set("strictQuery", false);
mongoose.connect(config.DATABASE_URI);

mongoose.connection.on("connect", () => {
  console.log("Database Connected...");
  startServer();
});

mongoose.connection.once("error", (err) => {
  console.log(`Could Not Connect to Database...`);
  throw new Error(err);
});

const app: Application = express();
const server = http.createServer(app);

// @ts-ignore
const io = socketIO(server);

/**
 * starts the express server after the database connects
 */
async function startServer() {
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
      const user = await User.findOne({ _id: req.user });
      res.send(user);
    }
  );
  app.use("/rooms", checkAuthentication, rooms);

  server.listen(config.PORT, () =>
    console.log("running on port " + config.PORT)
  );
}

export { io as IO };
