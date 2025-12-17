import { apiConfig } from './api.config';

export const PegawaiService = {
    async getAll(params = {}) {
        // TODO: Uncomment when API ready
        // const queryString = new URLSearchParams(params).toString();
        // return await apiCall(`${apiConfig.endpoints.pegawai}?${queryString}`);

        // Mock data for development
        return {
            data: [],
            total: 0,
            page: 1,
            limit: 10
        };
    },

    async create(data) {
        // TODO: Implement API call
        console.log('Create pegawai:', data);
        return { success: true };
    },

    async update(id, data) {
        // TODO: Implement API call
        console.log('Update pegawai:', id, data);
        return { success: true };
    },

    async delete(id) {
        // TODO: Implement API call
        console.log('Delete pegawai:', id);
        return { success: true };
    }
};