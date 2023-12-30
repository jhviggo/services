import { writable } from 'svelte/store';

interface Vehicle {
  id: string;
}

export const vehicleStore = writable<Vehicle>();

export async function fetchVehicles() {}

export async function addVehicle() {}
