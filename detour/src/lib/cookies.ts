import { writable, type Writable } from 'svelte/store';

export const COOKIE_ACCEPTED = 'accepted';
export const COOKIE_DENIED = 'denied';
export const COOKIE_UNKNOWN = 'unknown';

interface CookieAccepted {
  status: string;
}

export const cookiesStore: Writable<CookieAccepted> = writable({
  status: COOKIE_UNKNOWN,
});

export function checkCookie() {
  const storedStatus = localStorage.getItem('cookies-status');

  if (storedStatus && storedStatus !== '') {
    cookiesStore.set({
      status: storedStatus,
    });
  }
  
  cookiesStore.subscribe((value) => {
    localStorage.setItem('cookies-status', value.status);
  });
}



