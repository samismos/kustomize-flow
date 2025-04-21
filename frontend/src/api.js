// api.js

const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL;
const DEFAULT_TIMEOUT = 8000;

// Build a full URL with query parameters
const buildUrl = (endpoint, params) => {
    const url = new URL(endpoint, BASE_URL);
    url.search = new URLSearchParams(params).toString();
    return url.toString();
};

// Fetch data with timeout and error handling
const fetchWithTimeout = async (url, timeout = DEFAULT_TIMEOUT) => {
    const controller = new AbortController();
    const timer = setTimeout(() => controller.abort(), timeout);

    try {
        const response = await fetch(url, { signal: controller.signal });
        clearTimeout(timer);

        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status} ${response.statusText}`);
        }

        return await response.json();
    } catch (error) {
        console.error(`Error fetching from ${url}:`, error);
        return { nodes: [], edges: [] };
    }
};

// Fetch a single service by entrypoint
export const getService = (path) => {
    const url = buildUrl('/getService', { entrypoint: path });
    return fetchWithTimeout(url);
};

// Fetch all services by entrypoint
export const getAllServices = (path) => {
    const url = buildUrl('/getAllServices', { entrypoint: path });
    return fetchWithTimeout(url);
};