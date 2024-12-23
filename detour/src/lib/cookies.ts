import { get, writable, type Writable } from 'svelte/store';

const VISITOR_KEY = 'detour-visitor';

export function setVisitorCookie(value: string) {
  if (value != '') {
    localStorage.setItem(VISITOR_KEY, value);
  }
}

export function getVisitorCookie(): string {
  return localStorage.getItem(VISITOR_KEY) || '';
}
