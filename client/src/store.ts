import type { User } from './domain/user';
import { writable } from 'svelte/store';

export const user = writable<User | null>(null);
