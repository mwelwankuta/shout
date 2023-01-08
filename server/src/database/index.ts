import { connect } from "@planetscale/database";
import config from "../config";

const database = {
  host: config.DATABASE_HOST,
  username: config.DATABASE_USER,
  password: config.DATABASE_PASSWORD,
};

const conn = connect(database);
export default conn;
