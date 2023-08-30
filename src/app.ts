import express, { Application } from 'express';
import dotenv from 'dotenv';
//import { authenticationMiddleware } from './shared/auth';
import { initializeDatabase } from './shared/database';
import { vehicleRouter } from './vehicle/routes';
import { userRouter } from './users/routes';

dotenv.config();

const app: Application = express();
const port = process.env.PORT || 1337;

//app.use(authenticationMiddleware);
app.use(vehicleRouter);
app.use(userRouter);

Promise.all([
  initializeDatabase(),
]).then(() => {
  app.listen(port, () => {
    console.log(`⚡ Server running at at http://localhost:${port}`);
  });
});
