import React, { useEffect, useState } from 'react';
import { getAllServices } from './api';

const EndpointDropdown = ({ repo, endpoint, setEndpoint }) => {
    const [loading, setLoading] = useState(false);
    const [endpoints, setEndpoints] = useState([]);

    const nprodServicesPath = 'repos/' + repo + import.meta.env.VITE_NPROD_CLUSTER_PATH

    useEffect(() => {
        const fetchData = async () => {
            setLoading(true);
            try {
                const data = await getAllServices(nprodServicesPath);
                data.Resources.sort((a, b) => {
                    let x = a.toLowerCase();
                    let y = b.toLowerCase();
                    return x < y ? -1 : x > y ? 1 : 0;
                });
                setEndpoints(data.Resources);
            } catch (error) {
                console.error('Error fetching endpoints:', error);
            } finally {
                setLoading(false);
            }
        };
        console.log("Running for repo: " + repo)
        fetchData();
    }, [repo]);

    return (
        <div>
            {loading ? (
                <p>Loading...</p>
            ) : (
                <select
                    value={endpoint}
                    onChange={(e) => {
                        setEndpoint(e.target.value)
                    }}
                >
                    <option disabled={true} value="">Select a service</option>
                    {endpoints.map((ep, index) => (
                        <option key={index} value={ep}>
                            {ep}
                        </option>
                    ))}
                </select>
            )}
        </div>
    );
};

export default EndpointDropdown;
