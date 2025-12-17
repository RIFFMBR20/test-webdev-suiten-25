import { apiConfig } from './api.config';

export const AuthService = {
    async login(credentials) {
        // TODO: Implement API call
        console.log('Login:', credentials);
        localStorage.setItem('token', 'dummy-token');
        return { success: true, token: 'dummy-token' };
    },

    async logout() {
        // TODO: Implement API call
        console.log('Logout');
        localStorage.removeItem('token');
        return { success: true };
    }
};