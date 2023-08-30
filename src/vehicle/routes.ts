import { Request, Response, Router } from 'express';
import { v4 as uuidv4 } from 'uuid';
import { dbInstance } from '../shared/database';

export const vehicleRouter = Router();

// Refuels //
vehicleRouter.get('/v1/refuels', (_req: Request, res: Response) => {
  /*
  dbInstance.each('SELECT * FROM lorem', (err, row) => {
    console.log(row);
  });*/
  dbInstance.get('SELECT * FROM lorem', (error, row) => {
    console.log(row);
  });
  res.send([]);
});

vehicleRouter.get('/v1/refuels/:id', (req: Request, res: Response) => {
  const refuelId: string = req.params.id!;

  res.send({
    refuelId,
  });
});

vehicleRouter.post('/v1/refuels', (_req: Request, res: Response) => {
  const refuelId = uuidv4();

  res.send({
    refuelId,
  });
});

vehicleRouter.post('/v1/refuels/:id', (req: Request, res: Response) => {
  const vehicleId: string = req.params.id!;
  // replace old
  res.send({
    id: vehicleId,
  });
});

// vehicles //
vehicleRouter.get('/v1/vehicles', (req: Request, res: Response) => {
  res.send([]);
});

vehicleRouter.get('/v1/vehicles/:id', (req: Request, res: Response) => {
  const vehicleId: string = req.params.id!;

  res.send({
    vehicleId,
  });
});

vehicleRouter.post('/v1/vehicles', (req: Request, res: Response) => {
  res.send({});
});

vehicleRouter.post('/v1/vehicles/:id', (req: Request, res: Response) => {
  const vehicleId: string = req.params.id!;

  res.send({
    vehicleId,
  });
});
