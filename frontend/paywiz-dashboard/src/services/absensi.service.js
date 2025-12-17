import { apiConfig } from './api.config';

export const AbsensiService = {
    async getAll(params = {}) {
        // TODO: Uncomment when API ready
        // const queryString = new URLSearchParams(params).toString();
        // return await apiCall(`${apiConfig.endpoints.absensi}?${queryString}`);

        // Mock data for development
        return {
            data: [
                { id: 1, tanggal: '03, Okt 2024', nama: 'Nurhadi', bagian: 'Tukang Kayu', jamPulang: '-', totalLembur: '-', catatan: '-' },
                { id: 2, tanggal: '03, Okt 2024', nama: 'Ahmad Saroni', bagian: 'Tukang Kayu', jamPulang: '-', totalLembur: '-', catatan: '-' },
                { id: 3, tanggal: '03, Okt 2024', nama: 'Ahmad Syahroni', bagian: 'Tukang Kayu', jamPulang: '-', totalLembur: '-', catatan: '-' },
                { id: 4, tanggal: '03, Okt 2024', nama: 'Syafiq', bagian: 'Tukang Kayu', jamPulang: '-', totalLembur: '-', catatan: '-' },
                { id: 5, tanggal: '03, Okt 2024', nama: 'Agus Joko Nursiyo', bagian: 'Tukang Kayu', jamPulang: '-', totalLembur: '-', catatan: '-' },
                { id: 6, tanggal: '03, Okt 2024', nama: 'Dedi Suryamin', bagian: 'Tukang Kayu', jamPulang: '-', totalLembur: '-', catatan: '-' },
                { id: 7, tanggal: '03, Okt 2024', nama: 'Parmadi', bagian: 'Tukang Kayu', jamPulang: '-', totalLembur: '-', catatan: '-' },
                { id: 8, tanggal: '03, Okt 2024', nama: 'Purnawan', bagian: 'Tukang Kayu', jamPulang: '-', totalLembur: '-', catatan: '-' },
                { id: 9, tanggal: '03, Okt 2024', nama: 'Rudiyanto', bagian: 'Tukang Kayu', jamPulang: '-', totalLembur: '-', catatan: '-' },
                { id: 10, tanggal: '03, Okt 2024', nama: 'Suyanto', bagian: 'Tukang Kayu', jamPulang: '-', totalLembur: '-', catatan: '-' },
            ],
            total: 10,
            page: 1,
            limit: 10
        };
    },

    async create(data) {
        console.log('Create absensi:', data);
        return { success: true };
    },

    async update(id, data) {
        console.log('Update absensi:', id, data);
        return { success: true };
    },

    async delete(id) {
        console.log('Delete absensi:', id);
        return { success: true };
    }
};