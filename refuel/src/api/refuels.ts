import { API_URL } from '@lib/env';
import { get, writable } from 'svelte/store';
import { loadStoredUser, userStore } from './users';
import { addError, Levels } from '@lib/errorHandler';
import { DefaultApiResponse } from './defaultApiResponse';

export interface Refuel {
  id?: number | string;
  vehicle_id: number | string;
  total_km: number,
  trip_km: number,
  liters: number,
  cost: number,
  currency: string,
  created_at?: string,
}

export const refuelStore = writable<Refuel>();

export async function fetchRefuels(vehicleId: number | string): Promise<Refuel[]> {
  loadStoredUser();
  const user = get(userStore);
  if (!user?.token) return [];

  try {
    const response = await fetch(`${API_URL}/users/${user.id}/vehicles/${vehicleId}/refuels`, {
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
    const response = await fetch(`${API_URL}/users/${user.id}/vehicles/${refuel.vehicle_id}/refuels`, {
      method: "POST",
      headers: {
        'Authorization': `Bearer ${user.token}`
      },
      body: JSON.stringify({
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

export async function deleteRefuel(refuel: Refuel): Promise<void> {
  const user = get(userStore);
  if (!user?.token) return;
  console.log('ref', refuel)

  try {
    const response = await fetch(`${API_URL}/users/${user.id}/vehicles/${refuel.vehicle_id}/refuels/${refuel.id}`, {
      method: "DELETE",
      headers: {
        'Authorization': `Bearer ${user.token}`,
      },
    });
    if (response.status !== 200) {
      const data: DefaultApiResponse = await response.json();
      throw new Error(data.message);
    }
  } catch (err) {
    addError({
      level: Levels.ERROR,
      message: 'Unable to delete refuels',
    });
  }
}
