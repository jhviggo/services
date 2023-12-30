import { writable } from 'svelte/store';

interface User {
  userId: string;
  name: string;
  username: string;
}

export const userStore = writable<User>();

export async function login(username: string, password: string) {}

export async function fetchUser() {}

export async function addUser() {}
