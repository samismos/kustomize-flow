export const resolvePath = (basePath, additionalPath) => {
    // Split paths into segments
    let baseSegments = basePath.split('/');
    let additionalSegments = additionalPath.split('/');

    // Process additional path (handling ".." to go back directories)
    for (let segment of additionalSegments) {
        if (segment === '..') {
            // Remove the last directory from basePath
            baseSegments.pop();
        } else if (segment !== '.') {
            // Add valid segments to the path
            baseSegments.push(segment);
        }
    }

    // Join resolved path back into a string
    return baseSegments.join('/');
}