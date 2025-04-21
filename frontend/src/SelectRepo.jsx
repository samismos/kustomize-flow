import React, { useEffect, useState } from 'react';

const SelectRepoList = ({ selectedRepo, setSelectedRepo }) => {
    const repos = [
        'oi-doc-core-deployment',
        'oi-vmp-deployment',
        'oi-doc-standard-work-deployment',
    ]

    return (
        <div>
                <select
                    value={selectedRepo}
                    onChange={(e) => {
                        setSelectedRepo(e.target.value)
                    }}
                >
                    <option disabled={true} value="">Select a Repo</option>
                    {repos.map((ep, index) => (
                        <option key={index} value={ep}>
                            {ep}
                        </option>
                    ))}
                </select>
        </div>
    );
};

export default SelectRepoList;
