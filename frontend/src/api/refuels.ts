import { API_URL } from '@lib/env';
import { get, writable } from 'svelte/store';
import { loadStoredUser, userStore } from './users';
import { addError, Levels } from '@lib/errorHandler';

export interface Refuel {
  id?: number | string;
  vehicleId: number | string;
  totalKM: number,
  tripKM: number,
  liters: number,
  cost: number,
  currency: string,
  createdAt?: string,
}

export const refuelStore = writable<Refuel>();

export async function fetchRefuels(vehicleId: number | string): Promise<Refuel[]> {
  loadStoredUser();
  const user = get(userStore);
  if (!user?.token) return [];

  try {
    const response = await fetch(`${API_URL}/v1/users/${user.id}/vehicles/${vehicleId}/refuels`, {
      headers: {
        'Authorization': `Bearer ${user.token}`,
        'Content-Type': 'application/json',
      },
    });
    if (response.status !== 200) throw new Error();
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
  const user = get(userStore);
  if (!user?.token) return false;

  try {
    const response = await fetch(`${API_URL}/v1/users/${user.id}/vehicles/${refuel.vehicleId}/refuels`, {
      method: "POST",
      headers: {
        'Authorization': `Bearer ${user.token}`
      },
      body: JSON.stringify({
        userId: Number(user.id),
        ...refuel,
      }),
    });
    if (response.status !== 200) throw new Error();
  } catch (err) {
    addError({
      level: Levels.ERROR,
      message: 'Unable to add refuels',
    });
    return false;
  }
  return true;
}
