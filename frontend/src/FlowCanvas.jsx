import React, { useCallback, useState, useEffect } from 'react';
import { ReactFlow, Background, Controls, useNodesState, useEdgesState, addEdge } from '@xyflow/react';
import '@xyflow/react/dist/style.css';
import { getService } from './api';
import { resolvePath } from './PathResolver'
import EndpointDropdown from './EndpointDropdown';
import SelectRepoList from './SelectRepo';

const FlowCanvas = () => {
    const [nodes, setNodes, onNodesChange] = useNodesState([]);
    const [edges, setEdges, onEdgesChange] = useEdgesState([]);
    const [loading, setLoading] = useState(false)
    const [selectedRepo, setSelectedRepo] = useState(import.meta.env.VITE_DEFAULT_REPO)
    const [endpoint, setEndpoint] = useState(import.meta.env.VITE_DEFAULT_SERVICE)

    const nprodServicesPath = 'repos/' + selectedRepo + import.meta.env.VITE_NPROD_CLUSTER_PATH

    useEffect(() => {
        const fetchData = async () => {
            setLoading(true);
            const { nodes: fetchedNodes, edges: fetchedEdges } = await getService(resolvePath(nprodServicesPath, endpoint));

            const nodeMap = new Map(); // Store unique nodes and positions
            const positions = new Set(); // Track positions to prevent overlaps

            let maxY = 250; // Keep track of the maximum y value for stacking independent components

            // Create nodes with unique positions
            const formattedNodes = fetchedNodes.reduce((uniqueNodes, node, index) => {
                const isComponent = node.includes('flux/components');
                if (!nodeMap.has(node)) {
                    // Default position for file nodes
                    let x = index * 200; // Horizontal spacing
                    let y = 150; // Default vertical spacing for files

                    // For components, ensure they stack
                    if (isComponent) {
                        y = maxY; // Stack components below the current maxY
                        maxY += 100; // Increment maxY for the next component
                    }

                    // Resolve overlap conflicts, if needed
                    while (positions.has(`${x},${y}`)) {
                        y += 100; // Push down further to avoid overlap
                    }
                    positions.add(`${x},${y}`); // Mark position as occupied
                    nodeMap.set(node, { x, y }); // Save position to nodeMap

                    // Assign a yellow background color to components
                    const style = {
                        backgroundColor: isComponent ? '#FFD700' : '#FFFFFF', // Yellow for components, white otherwise
                        width: 200, // Fixed width for nodes
                        padding: '5px', // Add padding for text
                        overflow: 'hidden', // Ensure no text flows outside the box
                        whiteSpace: 'normal', // Enable text wrapping
                        display: 'flex', // Use flexbox for alignment
                        alignItems: 'center', // Vertically center text
                        justifyContent: 'center', // Horizontally center text
                        flexDirection: 'column', // Arrange text vertically if wrapped
                        border: '1px solid #000', // Optional: add border for clarity
                    };

                    uniqueNodes.push({
                        id: node, // Ensure the ID is unique
                        data: { label: node }, // Display the ID as text
                        position: { x, y },
                        style, // Apply styles to allow text wrapping
                    });
                }
                return uniqueNodes;
            }, []);

            const uniqueEdgeSet = new Set();
            const formattedEdges = fetchedEdges.reduce((edges, edge) => {
                const edgeKey = `e-${edge.from}-${edge.to}`;
                if (!uniqueEdgeSet.has(edgeKey)) {
                    uniqueEdgeSet.add(edgeKey);
                    edges.push({
                        id: edgeKey,
                        source: edge.from,
                        target: edge.to,
                        type: 'smoothstep',
                    });
                }
                return edges;
            }, []);

            setNodes(formattedNodes);
            setEdges(formattedEdges);
            setLoading(false);
        };
        fetchData();
    }, [endpoint]);


    const onConnect = useCallback(
        (params) => setEdges((eds) => addEdge(params, eds)),
        [setEdges],
    );

    if (loading) {
        return <div>Loading...</div>;
    }

    return (
        <div style={styles.container}>
            <SelectRepoList selectedRepo={selectedRepo} setSelectedRepo={setSelectedRepo} />
            <EndpointDropdown repo={selectedRepo} endpoint={endpoint} setEndpoint={setEndpoint} />
            <ReactFlow
                nodes={nodes}
                edges={edges}
                onNodesChange={onNodesChange}
                onEdgesChange={onEdgesChange}
                onConnect={onConnect}
                fitView
                style={styles.flow}
                defaultEdgeOptions={{
                    style: { stroke: '#1D4ED8', strokeWidth: 2 }, // nice green
                    animated: true,
                }}
            >
                <Background variant="dots" gap={12} size={1} />
                <Controls />
            </ReactFlow>
        </div>
    );
};

const styles = {
    container: {
        width: '95vw',
        height: '95vh',
    },
    flow: {
        width: '100%',
        height: '100%',
    },
};

export default FlowCanvas;
