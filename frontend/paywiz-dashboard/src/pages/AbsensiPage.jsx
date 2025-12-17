import {useState} from 'react';
import {Plus, Calendar, ChevronDown} from 'lucide-react';
import {useAbsensi} from '../hooks/useAbsensi';
import {SearchBar} from '../components/common/SearchBar';
import {EmptyState} from '../components/common/EmptyState';
import {LoadingSpinner} from '../components/common/LoadingSpinner';
import {Pagination} from '../components/common/Pagination';

export const AbsensiPage = () => {
    const {data, loading, error, pagination, fetchData} = useAbsensi();
    const [searchQuery, setSearchQuery] = useState('');
    const [selectedDate] = useState('03/10/2024');
    const [selectedBagian] = useState('Tukang Kayu');
    const handleSearch = (query) => {
        setSearchQuery(query);
        fetchData({search: query, date: selectedDate, bagian: selectedBagian});
    };
    const handlePrevious = () => {
        if (pagination.page > 1) {
            fetchData({page: pagination.page - 1, date: selectedDate, bagian: selectedBagian});
        }
    };
    const handleNext = () => {
        const totalPages = Math.ceil(pagination.total / pagination.limit);
        if (pagination.page < totalPages) {
            fetchData({page: pagination.page + 1, date: selectedDate, bagian: selectedBagian});
        }
    };
    return (
        <div className="flex-1 bg-gray-50">
            <div className="bg-white border-b border-gray-200 p-6">
                <div className="flex items-center justify-between">
                    <h1 className="text-2xl font-semibold">Absensi</h1>
                    <div className="flex items-center gap-3">
                        <SearchBar value={searchQuery} onChange={handleSearch}/>
                        <button
                            className="flex items-center gap-2 bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600">
                            <Plus className="w-5 h-5"/>
                            Input Absen
                        </button>
                    </div>
                </div>
            </div>
            <div className="p-6">
                <div className="bg-white rounded-lg border border-gray-200 overflow-hidden">
                    <div className="p-4 border-b border-gray-200 flex items-center justify-between">
                        <h2 className="text-lg font-medium">Daftar Absensi Pegawai</h2>
                        <div className="flex items-center gap-3">
                            <button
                                className="flex items-center gap-2 text-sm border border-gray-300 px-3 py-1.5 rounded">
                                <Calendar className="w-4 h-4"/>
                                {selectedDate}
                            </button>
                            <button
                                className="flex items-center gap-2 text-sm text-gray-600 border border-gray-300 px-3 py-1.5 rounded">
                                {selectedBagian}
                                <ChevronDown className="w-4 h-4"/>
                            </button>
                        </div>
                    </div>

                    {loading ? (
                        <LoadingSpinner/>
                    ) : error ? (
                        <div className="text-center py-16 text-red-500">{error}</div>
                    ) : data.length === 0 ? (
                        <EmptyState/>
                    ) : (
                        <div className="overflow-x-auto">
                            <table className="w-full">
                                <thead className="bg-gray-50 border-b border-gray-200">
                                <tr>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Tanggal</th>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Nama
                                        Pegawai
                                    </th>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Bagian</th>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Jam
                                        Pulang
                                    </th>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Total
                                        Lembur
                                    </th>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Catatan</th>
                                </tr>
                                </thead>
                                <tbody className="divide-y divide-gray-200">
                                {data.map((item) => (
                                    <tr key={item.id} className="hover:bg-gray-50">
                                        <td className="px-6 py-4 text-sm text-gray-700">{item.tanggal}</td>
                                        <td className="px-6 py-4 text-sm text-gray-700">{item.nama}</td>
                                        <td className="px-6 py-4 text-sm text-gray-700">{item.bagian}</td>
                                        <td className="px-6 py-4 text-sm text-gray-700">{item.jamPulang}</td>
                                        <td className="px-6 py-4 text-sm text-gray-700">{item.totalLembur}</td>
                                        <td className="px-6 py-4 text-sm text-gray-700">{item.catatan}</td>
                                    </tr>
                                ))}
                                </tbody>
                            </table>
                        </div>
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