// lib/api.ts
const API_BASE = '/api';

export const apiClient = {  
  async login(email: string) {
    const response = await fetch(`${API_BASE}/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email }),
    });
    if (!response.ok) throw new Error('Login failed');
    return response.json();
  },

  async register(email:string, name:string){
    const response = await fetch(`${API_BASE}/createUser`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, name }),
    });
    if (!response.ok) throw new Error('Registration failed');
    return response.json();
  }
};