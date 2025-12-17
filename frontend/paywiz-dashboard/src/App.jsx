import { useState } from 'react';
import { Sidebar } from './components/layout/Sidebar';
import { EmployeePage } from './pages/EmployeePage';
import { DivisionPage } from './pages/DivisionPage';
import { AbsensiPage } from './pages/AbsensiPage';

function App() {
    const [currentPage, setCurrentPage] = useState('employee');

    return (
        <div className="flex h-screen bg-gray-50">
            <Sidebar currentPage={currentPage} setCurrentPage={setCurrentPage} />
            {currentPage === 'employee' && <EmployeePage />}
            {currentPage === 'division' && <DivisionPage />}
            {currentPage === 'absensi' && <AbsensiPage />}
        </div>
    );
}

export default App;