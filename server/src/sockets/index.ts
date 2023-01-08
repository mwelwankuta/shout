import { Socket } from "socket.io";
import { IO } from "../server";

IO.on("connection", (socket: Socket) => {
  socket.emit("hello this message came from the server");
});
