import { apiConfig } from './api.config';

export const BagianService = {
    async getAll(params = {}) {
        // TODO: Uncomment when API ready
        // const queryString = new URLSearchParams(params).toString();
        // return await apiCall(`${apiConfig.endpoints.bagian}?${queryString}`);

        // Mock data for development
        return {
            data: [],
            total: 0,
            page: 1,
            limit: 10
        };
    },

    async create(data) {
        console.log('Create division:', data);
        return { success: true };
    },

    async update(id, data) {
        console.log('Update division:', id, data);
        return { success: true };
    },

    async delete(id) {
        console.log('Delete division:', id);
        return { success: true };
    }
};