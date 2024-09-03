import { Routes, Route } from "react-router-dom";
import DefaultLayout from "../layouts/DefaultLayout";
import { analyzer } from "../../wailsjs/go/models";
import { ListDisks } from "../../wailsjs/go/main/App";
import { useEffect, useState } from "react";
import { bytesToHumanReadable }  from "../lib/humanize";

const DiskSelector = () => {
    const [disks, setDisks] = useState<analyzer.DiskInfo[]>([]);

    useEffect(() => {
        ListDisks().then((disks) => {
            setDisks(disks);
        });
    });

    return (
        <div className="w-full h-full flex flex-col gap-3 pt-1 flex-auto overflow-visible">
            {disks.map((disk) => (
                <div key={disk.device}>
                    {disk.device} - {disk.mount} - {bytesToHumanReadable(disk.used)} / {bytesToHumanReadable(disk.total)}
                </div>
            ))}
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
