import { Search } from 'lucide-react';

export const SearchBar = ({ value, onChange, placeholder = "Cari data..." }) => (
    <div className="relative">
        <input
            type="text"
            placeholder={placeholder}
            value={value}
            onChange={(e) => onChange(e.target.value)}
            className="pl-10 pr-4 py-2 border border-gray-300 rounded-lg w-80"
        />
        <Search className="w-5 h-5 text-gray-400 absolute left-3 top-2.5" />
    </div>
);