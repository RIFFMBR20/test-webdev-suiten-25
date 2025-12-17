import { useState, useEffect } from 'react';
import { PegawaiService } from '../services/employee.service';

export const useEmployee = () => {
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const [pagination, setPagination] = useState({ page: 1, total: 0, limit: 10 });

    const fetchData = async (params = {}) => {
        setLoading(true);
        setError(null);
        try {
            const response = await PegawaiService.getAll(params);
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

    const createPegawai = async (pegawaiData) => {
        setLoading(true);
        try {
            await PegawaiService.create(pegawaiData);
            await fetchData();
        } catch (err) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    };

    const updatePegawai = async (id, pegawaiData) => {
        setLoading(true);
        try {
            await PegawaiService.update(id, pegawaiData);
            await fetchData();
        } catch (err) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    };

    const deletePegawai = async (id) => {
        setLoading(true);
        try {
            await PegawaiService.delete(id);
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

    return { data, loading, error, pagination, fetchData, createPegawai, updatePegawai, deletePegawai };
};