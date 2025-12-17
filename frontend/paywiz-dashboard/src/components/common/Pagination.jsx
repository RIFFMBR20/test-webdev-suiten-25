import { ChevronLeft, ChevronRight } from 'lucide-react';

export const Pagination = ({ onPrevious, onNext, currentPage, totalPages }) => (
    <div className="flex items-center justify-center gap-4 p-4 border-t border-gray-200">
        <button
            onClick={onPrevious}
            disabled={currentPage === 1}
            className="flex items-center gap-1 text-sm text-gray-600 hover:text-gray-800 disabled:opacity-50 disabled:cursor-not-allowed"
        >
            <ChevronLeft className="w-4 h-4" />
            Previous
        </button>
        <span className="text-sm text-gray-600">
      Page {currentPage} of {totalPages || 1}
    </span>
        <button
            onClick={onNext}
            disabled={currentPage === totalPages}
            className="flex items-center gap-1 text-sm text-gray-600 hover:text-gray-800 disabled:opacity-50 disabled:cursor-not-allowed"
        >
            Next
            <ChevronRight className="w-4 h-4" />
        </button>
    </div>
);