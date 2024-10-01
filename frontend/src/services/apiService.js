// src/services/apiService.js
import axios from 'axios';

const api = axios.create({
  baseURL: process.env.REACT_APP_API_URL,
});

export const getData = async () => {
  try {
    const response = await api.get('/your-endpoint');
    return response.data;
  } catch (error) {
    console.error('API call error:', error);
  }
};
