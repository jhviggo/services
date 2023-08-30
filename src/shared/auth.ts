import { Request, Response, NextFunction } from 'express';
import { decode, sign, verify } from 'jsonwebtoken';

const PRIVATE_KEY: string = process.env.JWT_PRIVATE_KEY!;
const ISSUER = process.env.JWT_ISSUER;

export function signJWT(user: string) {
  return sign({
    user
  }, PRIVATE_KEY, {
    issuer: ISSUER,
    expiresIn: '3h'
  });
}

export function verifyJWT(token: string) {
  return verify(token, PRIVATE_KEY);
}

export function decodeJWT(token: string) {
  return decode(token);
}

export function authenticationMiddleware(req: Request, res: Response, next: NextFunction) {
  const token: string = req.headers.authorization!;

  if (!token) {
    res.status(400).send({
      message: 'Missing auth token.',
    });
    return;
  }

  if (!verifyJWT(token)) {
    res.status(401).send({
      message: 'Unauthorized token.'
    });
    return;
  }

  next();
}
