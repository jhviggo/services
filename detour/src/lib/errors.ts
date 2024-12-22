import { writable } from 'svelte/store';

export enum Levels {
  DEBUG = 'debug',
  INFO = 'info',
  WARN = 'warn',
  ERROR = 'error',
  CRITICAL = 'critical',
}

export interface ErrorType {
  level: Levels;
  message: string;
}

export const errorStore = writable<ErrorType[]>([]);

export function addError(err: ErrorType) {
  errorStore.update((value) => {
    value.push(err)
    setTimeout(() => removeError(err), 2000);
    return value;
  });
}

export function removeError(err: ErrorType) {
  errorStore.update((value) => value.filter((e) => e !== err));
}