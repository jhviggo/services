import { writable, get } from 'svelte/store';
import { API_URL } from '@lib/env';
import { addError, Levels } from '@lib/errorHandler';

interface User {
  id?: string;
  name: string;
  username: string;
  token: string;
}

export const userStore = writable<User>();

let userHasBeenUpdated = false;

export function loadStoredUser() {
  if (userHasBeenUpdated) return;

  const storedUser = localStorage.getItem('user');
  if (storedUser && storedUser !== "undefined") {
    userStore.set(JSON.parse(storedUser));
    userHasBeenUpdated = true;
  }

  userStore.subscribe((value) => {
    if (value)
      localStorage.setItem('user', JSON.stringify(value));
  });
}

export async function login(username: string, password: string): Promise<boolean> {
  try {
    const response = await fetch(`${API_URL}/v1/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        username,
        password,
      }),
    });
    if (response.status !== 200) {
      throw new Error('unable to login');
    }
    const user = await response.json();
    userStore.set(user);
  } catch (err) {
    addError({
      level: Levels.WARN,
      message: 'Unable to login',
    });
    return false;
  }
  return true;
}

export function logout() {
  userStore.set({} as User);
}

export async function fetchUser(id: string) {
}

export async function addUser() {}
