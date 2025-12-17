import { useState } from 'react';
import { ChevronDown, ChevronRight } from 'lucide-react';
import { AuthService } from '../../services/auth.service';

export const Sidebar = ({ currentPage, setCurrentPage }) => {
    const [masterDataExpanded, setMasterDataExpanded] = useState(true);

    const handleLogout = () => {
        AuthService.logout();
    };

    return (
        <div className="w-64 bg-white border-r border-gray-200 flex flex-col h-screen">
            <div className="p-6 border-b border-gray-200">
                <div className="flex items-center gap-2">
                    <div className="w-8 h-8 bg-blue-500 rounded flex items-center justify-center text-white font-bold">W</div>
                    <span className="font-semibold text-lg">PayWiz</span>
                </div>
            </div>

            <div className="flex-1 p-4 overflow-y-auto">
                <div className="mb-2">
                    <button
                        onClick={() => setMasterDataExpanded(!masterDataExpanded)}
                        className="w-full flex items-center gap-2 px-3 py-2 text-gray-700 hover:bg-gray-50 rounded"
                    >
                        {masterDataExpanded ? <ChevronDown className="w-4 h-4" /> : <ChevronRight className="w-4 h-4" />}
                        <span className="text-sm">Master Data</span>
                    </button>

                    {masterDataExpanded && (
                        <div className="ml-4 mt-1">
                            <button
                                onClick={() => setCurrentPage('pegawai')}
                                className={`w-full text-left px-3 py-2 text-sm rounded ${
                                    currentPage === 'pegawai' ? 'bg-blue-100 text-blue-600' : 'text-gray-600 hover:bg-gray-50'
                                }`}
                            >
                                Pegawai
                            </button>
                            <button
                                onClick={() => setCurrentPage('bagian')}
                                className={`w-full text-left px-3 py-2 text-sm rounded ${
                                    currentPage === 'bagian' ? 'bg-blue-100 text-blue-600' : 'text-gray-600 hover:bg-gray-50'
                                }`}
                            >
                                Bagian
                            </button>
                            <button className="w-full text-left px-3 py-2 text-sm text-gray-600 hover:bg-gray-50 rounded">
                                Shift
                            </button>
                            <button className="w-full text-left px-3 py-2 text-sm text-gray-600 hover:bg-gray-50 rounded">
                                Bank
                            </button>
                        </div>
                    )}
                </div>

                <button
                    onClick={() => setCurrentPage('absensi')}
                    className={`w-full flex items-center gap-2 px-3 py-2 text-sm rounded ${
                        currentPage === 'absensi' ? 'bg-blue-100 text-blue-600' : 'text-gray-600 hover:bg-gray-50'
                    }`}
                >
                    <ChevronRight className="w-4 h-4" />
                    <span>Absensi</span>
                </button>

                <button className="w-full flex items-center gap-2 px-3 py-2 text-sm text-gray-600 hover:bg-gray-50 rounded">
                    <ChevronRight className="w-4 h-4" />
                    <span>Uang Makan</span>
                </button>

                <button className="w-full flex items-center gap-2 px-3 py-2 text-sm text-gray-600 hover:bg-gray-50 rounded">
                    <ChevronRight className="w-4 h-4" />
                    <span>Hutang</span>
                </button>

                <button className="w-full flex items-center gap-2 px-3 py-2 text-sm text-gray-600 hover:bg-gray-50 rounded">
                    <ChevronDown className="w-4 h-4" />
                    <span>Payroll</span>
                </button>

                <button className="w-full flex items-center gap-2 px-3 py-2 text-sm text-gray-600 hover:bg-gray-50 rounded">
                    <ChevronDown className="w-4 h-4" />
                    <span>Setting</span>
                </button>
            </div>

            <div className="p-4 border-t border-gray-200">
                <div className="mb-2 px-3 py-2">
                    <div className="text-xs text-gray-500">Nungky</div>
                    <div className="text-xs text-gray-400">Admin</div>
                </div>
                <button onClick={handleLogout} className="w-full flex items-center gap-2 px-3 py-2 text-sm text-red-500 hover:bg-red-50 rounded">
                    <span>Logout</span>
                </button>
                <div className="text-xs text-gray-400 px-3 py-2">© RUITEN 2024</div>
            </div>
        </div>
    );
};