import { env } from '@config/env';
import axios from 'axios';

export const apiRequest = axios.create({
  baseURL: env.apiUrl,
  headers: {
    Accept: 'application/json'
  }
});
