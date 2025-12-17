import { FileX } from 'lucide-react';

export const EmptyState = ({ message = "Data belum tersedia" }) => (
    <div className="flex flex-col items-center justify-center py-16">
        <FileX className="w-16 h-16 text-blue-500 mb-4" />
        <p className="text-gray-500">{message}</p>
    </div>
);