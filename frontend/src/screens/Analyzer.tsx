import { Routes, Route } from "react-router-dom";
import DefaultLayout from "../layouts/DefaultLayout";

const DiskSelector = () => {
    return (
        <div className="w-full h-full flex flex-col gap-3 pt-1 flex-auto overflow-visible">
            Disk Selector
        </div>
    );
}

const Analyzer: React.FunctionComponent = () => {
    return (
        <DefaultLayout>
            <Routes location="/analyzer">
                <Route path="/" element={<DiskSelector />} />
            </Routes>
        </DefaultLayout>
    );
};

export default Analyzer;
