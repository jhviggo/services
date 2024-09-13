import { API_URL } from '@lib/env';
import { addError, Levels } from '@lib/errorHandler';
import { get, writable } from 'svelte/store';
import { userStore } from './users';

export interface IVehicle {
  id: string;
  estimatedKML: number;
  model: string;
}

export const vehicleStore = writable<IVehicle>();

export async function listVehicles(): Promise<IVehicle[]> {
  const user = get(userStore);
  console.log('weh', user);
  if (!user?.token) return [];
  console.log('wah');

  try {
    const response = await fetch(`${API_URL}/v1/users/${user.id}/vehicles`, {
      headers: {
        'Authorization': `Bearer ${user.token}`,
        'Content-Type': 'application/json',
      }
    });
    if (response.status !== 200) throw new Error();
    return await response.json();
  } catch (err) {
    addError({
      level: Levels.ERROR,
      message: 'Unable to fetch vehicles',
    });
  }
  return [];
}

export async function addVehicle() {}
