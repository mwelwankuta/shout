import dotenv from "dotenv";
// configure env
const ENVIRONMENT = process.env.NODE_ENV || "development";
dotenv.config({ path: `.env.${ENVIRONMENT}` });

export default {
  DATABASE_URI: process.env.DATABASE_URI,
  ENVIRONMENT,
  PORT: process.env.PORT,
};
