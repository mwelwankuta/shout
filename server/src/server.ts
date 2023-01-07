import express, { Application, Response, json } from "express";
import dotenv from "dotenv";
import cors from "cors";
import mongoose from "mongoose";
import http from "http";
import morgan from "morgan";
import socketIO, { Socket } from "socket.io";

// configure env
const environment = process.env.NODE_ENV || "development";
dotenv.config({ path: `.env.${environment}` });

import { authentication, rooms } from "./routes";
import { checkAuthentication } from "./helper";
import { User } from "./models";
import { AuthenticatedRequest } from "./types";

const DATABASE_URI = process.env.DATABASE_URI;

mongoose.set("strictQuery", false);
mongoose.connect(DATABASE_URI);

mongoose.connection.on("connect", () => {
  console.log("Database Connected...");
  startSocketServer();
});

mongoose.connection.once("error", (err) => {
  console.log(`Could Not Connect to Database... \n ${err}`);
});

async function startSocketServer() {
  const app = express();
  const server = http.createServer(app);

  setUpApp(app);

  const PORT = process.env.PORT;
  server.listen(PORT, () => console.log("running on port " + PORT));

  // @ts-ignore
  const io = socketIO(server);

  io.on("connection", (socket: Socket) => {
    socket.emit("hello this message came from the server");
  });
}

/**
 * starts the express server after the database connects
 */
async function setUpApp(app: Application) {
  // global middleware
  app.use(
    morgan(":method :url :status :res[content-length] - :response-time ms")
  );
  app.use(cors({ origin: "*" }));
  app.use(json());

  // routes
  app.use("/auth", authentication);
  app.get(
    "/auth/user",
    checkAuthentication,
    async (req: AuthenticatedRequest, res: Response) => {
      const user = await User.findOne({ _id: req.user });
      res.send(user);
    }
  );
  app.use("/rooms", checkAuthentication, rooms);
}
