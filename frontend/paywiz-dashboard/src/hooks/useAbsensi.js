import { useState, useEffect } from 'react';
import { AbsensiService } from '../services/absensi.service';

export const useAbsensi = () => {
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const [pagination, setPagination] = useState({ page: 1, total: 0, limit: 10 });

    const fetchData = async (params = {}) => {
        setLoading(true);
        setError(null);
        try {
            const response = await AbsensiService.getAll(params);
            setData(response.data);
            setPagination({
                page: response.page,
                total: response.total,
                limit: response.limit
            });
        } catch (err) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    };

    const createAbsensi = async (absensiData) => {
        setLoading(true);
        try {
            await AbsensiService.create(absensiData);
            await fetchData();
        } catch (err) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchData();
    }, []);

    return { data, loading, error, pagination, fetchData, createAbsensi };
};