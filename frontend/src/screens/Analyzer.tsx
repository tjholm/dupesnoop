import { Routes, Route, Outlet } from "react-router-dom";
import DefaultLayout from "../layouts/DefaultLayout";
import { analyzer } from "../../wailsjs/go/models";
import { ListDisks } from "../../wailsjs/go/main/App";
import { useEffect, useState } from "react";
import { bytesToHumanReadable }  from "../lib/humanize";
import { Link } from "react-router-dom";

export const DiskAnalysis = () => {
    return (
        <div>
            Disk Analysis
        </div>
    );
}

export const DiskSelector = () => {
    const [disks, setDisks] = useState<analyzer.DiskInfo[]>([]);

    useEffect(() => {
        ListDisks().then((disks) => {
            setDisks(disks);
        });
    });

    return (
        <div className="flex flex-col gap-3 pt-1 flex-auto overflow-visible">

            {disks.map((disk) => (
                <>
                    <Link to={`/analyzer/test`}>
                        <div key={disk.device} className="">
                            {disk.device} - {disk.mount} - {bytesToHumanReadable(disk.used)} / {bytesToHumanReadable(disk.total)}
                        </div>
                    </Link>
                </>
            ))}
        </div>
    );
}

const Analyzer: React.FunctionComponent = () => {
    return (
        <DefaultLayout>
            <Outlet />
        </DefaultLayout>
    );
};

export default Analyzer;
