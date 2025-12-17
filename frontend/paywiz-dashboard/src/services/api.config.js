export const apiConfig = {
    baseURL: import.meta.env.VITE_API_URL || 'https://api.example.com',
    endpoints: {
        pegawai: '/pegawai',
        bagian: '/bagian',
        absensi: '/absensi',
        auth: '/auth'
    }
};

export const apiCall = async (endpoint, options = {}) => {
    const token = localStorage.getItem('token');

    const response = await fetch(`${apiConfig.baseURL}${endpoint}`, {
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...(token && { 'Authorization': `Bearer ${token}` }),
            ...options.headers
        }
    });

    if (!response.ok) {
        throw new Error(`API Error: ${response.statusText}`);
    }

    return await response.json();
};