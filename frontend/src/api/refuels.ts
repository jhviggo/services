import { writable } from 'svelte/store';

interface Refuel {
  id: string;
}

export const refuelStore = writable<Refuel>();

export async function fetchRefuels() {}

export async function addRefuel() {}
