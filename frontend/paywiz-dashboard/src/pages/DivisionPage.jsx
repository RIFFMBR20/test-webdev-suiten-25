import { useState } from 'react';
import { Plus } from 'lucide-react';
import { useDivision } from '../hooks/useDivision';
import { SearchBar } from '../components/common/SearchBar';
import { EmptyState } from '../components/common/EmptyState';
import { LoadingSpinner } from '../components/common/LoadingSpinner';
import { Pagination } from '../components/common/Pagination';

export const DivisionPage = () => {
    const { data, loading, error, pagination, fetchData } = useDivision();
    const [searchQuery, setSearchQuery] = useState('');

    const handleSearch = (query) => {
        setSearchQuery(query);
        fetchData({ search: query });
    };

    const handlePrevious = () => {
        if (pagination.page > 1) {
            fetchData({ page: pagination.page - 1 });
        }
    };

    const handleNext = () => {
        const totalPages = Math.ceil(pagination.total / pagination.limit);
        if (pagination.page < totalPages) {
            fetchData({ page: pagination.page + 1 });
        }
    };

    return (
        <div className="flex-1 bg-gray-50">
            <div className="bg-white border-b border-gray-200 p-6">
                <div className="flex items-center justify-between mb-6">
                    <h1 className="text-2xl font-semibold">Master Data</h1>
                    <div className="flex items-center gap-3">
                        <SearchBar value={searchQuery} onChange={handleSearch} />
                        <button className="flex items-center gap-2 bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600">
                            <Plus className="w-5 h-5" />
                            Data Bagian
                        </button>
                    </div>
                </div>

                <h2 className="text-lg font-medium">Daftar Bagian</h2>
            </div>

            <div className="p-6">
                <div className="bg-white rounded-lg border border-gray-200 overflow-hidden">
                    <table className="w-full">
                        <thead className="bg-gray-50 border-b border-gray-200">
                        <tr>
                            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Nama Bagian</th>
                            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Jumlah Pegawai</th>
                        </tr>
                        </thead>
                    </table>

                    {loading ? (
                        <LoadingSpinner />
                    ) : error ? (
                        <div className="text-center py-16 text-red-500">{error}</div>
                    ) : data.length === 0 ? (
                        <EmptyState />
                    ) : (
                        <table className="w-full">
                            <tbody className="divide-y divide-gray-200">
                            {data.map((item) => (
                                <tr key={item.id} className="hover:bg-gray-50">
                                    <td className="px-6 py-4 text-sm text-gray-700">{item.namaBagian}</td>
                                    <td className="px-6 py-4 text-sm text-gray-700">{item.jumlahPegawai}</td>
                                </tr>
                            ))}
                            </tbody>
                        </table>
                    )}

                    <Pagination
                        onPrevious={handlePrevious}
                        onNext={handleNext}
                        currentPage={pagination.page}
                        totalPages={Math.ceil(pagination.total / pagination.limit)}
                    />
                </div>
            </div>
        </div>
    );
};