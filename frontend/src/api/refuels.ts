import { API_URL } from '@lib/env';
import { get, writable } from 'svelte/store';
import { loadStoredUser, userStore } from './users';
import { addError, Levels } from '@lib/errorHandler';

export interface Refuel {
  id?: number;
  vehicleId: number;
  totalKM: number,
  tripKM: number,
  liters: number,
  cost: number,
  currency: string,
}

export const refuelStore = writable<Refuel>();

export async function fetchRefuels(): Promise<Refuel[]> {
  loadStoredUser();
  const user = get(userStore);
  if (!user?.token) return [];

  const vehicleId = 18;
  try {
    const response = await fetch(`${API_URL}/v1/users/${user.id}/vehicles/${vehicleId}/refuels`, {
      headers: {
        'Authorization': `Bearer ${user.token}`,
        'Content-Type': 'application/json',
      },
    });
    const data = await response.json();
    return data;
  } catch (err) {
    addError({
      level: Levels.ERROR,
      message: 'Unable to fetch refuels',
    });
    return [];
  }
}

export async function addRefuel(refuel: Refuel): Promise<boolean> {
  console.log('lets gooo');
  const user = get(userStore);
  if (!user?.token) return false;

  const vehicleId = 18;

  try {
    const response = await fetch(`${API_URL}/v1/users/${user.id}/vehicles/${vehicleId}/refuels`, {
      method: "POST",
      headers: {
        'Authorization': `Bearer ${user.token}`
      },
      body: JSON.stringify({
        userId: Number(user.id),
        ...refuel,
      }),
    });
    console.log('res', response);
  } catch (err) {
    addError({
      level: Levels.ERROR,
      message: 'Unable to add refuels',
    });
    return false;
  }
  return true;
}
