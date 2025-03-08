
export const SkeletonBox = ({ className }: { className: string }) => (
    <div className={`bg-gray-500 animate-pulse rounded-xs ${className}`} style={{ animationDuration: '1.5s' }}></div>
);
