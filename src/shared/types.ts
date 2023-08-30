export interface Refuel {
  refuelId: string;
  cost: number;
  totelMilage: number;
  tripMilage: number;
  date: string;
  currency: string;
  location: string;
}

export interface Vehicle {
  vehicleId: string;
  name: string;
  model: string;
  year: number;
}

export interface User {
  displayName: string;
  username: string;
  password: string;
}