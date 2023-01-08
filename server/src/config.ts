import dotenv from "dotenv";
// configure env
const ENVIRONMENT = process.env.NODE_ENV || "development";
dotenv.config({ path: `.env.${ENVIRONMENT}` });

export default {
  DATABASE_HOST: process.env.DATABASE_HOST,
  DATABASE_USER: process.env.DATABASE_USER,
  DATABASE_PASSWORD: process.env.DATABASE_PASSWORD,
  ENVIRONMENT,
  PORT: process.env.PORT,
};
