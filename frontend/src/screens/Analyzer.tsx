import { Routes, Route, Outlet } from "react-router-dom";
import DefaultLayout from "../layouts/DefaultLayout";
import { analyzer } from "../../wailsjs/go/models";
import { ListDisks, GetDiskUsage } from "../../wailsjs/go/main/App";
import { useEffect, useState } from "react";
import { bytesToHumanReadable } from "../lib/humanize";
import { Link } from "react-router-dom";
import { ResponsiveSunburst } from "@nivo/sunburst";
import FileBrowser from "../components/FileBrowser";

export const DiskAnalysis = () => {
    const [usage, setUsage] = useState<analyzer.Node | null>(null);
    const [path, setPath] = useState<string | null>(null);

    useEffect(() => {
        if (path) {
            GetDiskUsage(path).then((usage) => {
                setUsage(usage);
            }).catch((err) => {
                console.error(err);
            });
        }
    }, [path]);

    const followPath = (path: string) => {
        setPath(path);
        // GetDiskUsage(path).then((usage) => {
        //     setUsage(usage);
        // }).catch((err) => {
        //     console.error(err);
        // });
    };

    if (!usage) {
        return (
            <>
                <FileBrowser
                    value={path || undefined}
                    onSelected={(path) =>
                        setPath(path)
                    }
                />
                <div>Loading...</div>
            </>
        );
    }

    console.log(usage);

    return (
        <div className="flex w-full h-full">
            <FileBrowser
                value={path || undefined}
                onSelected={(path) =>
                    setPath(path)
                }
            />
            <div className="flex-1">
                <ResponsiveSunburst
                    data={usage}
                    margin={{ top: 10, right: 10, bottom: 10, left: 10 }}
                    id="name"
                    value="loc"
                    cornerRadius={2}
                    borderColor={{ theme: 'background' }}
                    onClick={(node) => {
                        followPath(node.data.name);
                    }}
                    colors={{ scheme: 'nivo' }}
                    // childColor={{
                    //     from: 'color',
                    //     modifiers: [
                    //         [
                    //             'brighter',
                    //             0.1
                    //         ]
                    //     ]
                    // }}
                    enableArcLabels={true}
                    arcLabelsSkipAngle={10}
                    arcLabelsTextColor={{
                        from: 'color',
                        modifiers: [
                            [
                                'darker',
                                1.4
                            ]
                        ]
                    }}
                />
            </div>
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
                <Link key={disk.mount} to={`/analyzer/test`}>
                    <div className="">
                        {disk.device} - {disk.mount} - {bytesToHumanReadable(disk.used)} / {bytesToHumanReadable(disk.total)}
                    </div>
                </Link>
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
