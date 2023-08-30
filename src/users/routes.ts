import { Request, Response, Router } from 'express';
import { v4 as uuidv4 } from 'uuid';

export const userRouter = Router();

userRouter.get('/v1/users', (_req: Request, res: Response) => {
  res.send([]);
});

userRouter.get('/v1/users/:id', (req: Request, res: Response) => {
  const userId: string = req.params.id!;

  res.send({
    userId,
  });
});

userRouter.post('/v1/users', (_req: Request, res: Response) => {
  const userId = uuidv4();

  res.send({
    userId,
  });
});

userRouter.post('/v1/users/:id', (req: Request, res: Response) => {
  const userId: string = req.params.id!;
  // replace old
  res.send({
    userId,
  });
});
