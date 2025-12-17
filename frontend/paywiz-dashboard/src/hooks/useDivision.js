import { useState, useEffect } from 'react';
import { BagianService } from '../services/division.service';

export const useDivision = () => {
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const [pagination, setPagination] = useState({ page: 1, total: 0, limit: 10 });

    const fetchData = async (params = {}) => {
        setLoading(true);
        setError(null);
        try {
            const response = await BagianService.getAll(params);
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

    const createBagian = async (bagianData) => {
        setLoading(true);
        try {
            await BagianService.create(bagianData);
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

    return { data, loading, error, pagination, fetchData, createBagian };
};