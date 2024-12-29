import { get } from 'svelte/store';
import { loadStoredUser, userStore } from './users';
import { API_URL } from '@lib/env';
import { addError, Levels } from '@lib/errorHandler';

interface Stat {
  date_value: Number;
  sum: Number;
  count: Number;
}

export interface RefuelStatistic {
  year: string;
  weekly: Stat[];
  monthly: Stat[];
}

export async function fetchUserRefuelStatistics(year: string): Promise<RefuelStatistic[]> {
  loadStoredUser();
  const user = get(userStore);
  if (!user?.token) return [];

  try {
    const response = await fetch(`${API_URL}/statistics/${user.id}?year=${year}`, {
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