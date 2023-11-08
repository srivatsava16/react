import React, { useState, useMemo } from 'react';

function MemoExample() {
    const [number, setNumber] = useState(1);
    const [multiplier, setMultiplier] = useState(2);

    const result = useMemo(() => {
        console.log('Calculating result...');
        return number * multiplier;
    }, [number, multiplier]);

    return (
        <div>
            <h1>useMemo Example</h1>
            <p>Number: {number}</p>
            <p>Multiplier: {multiplier}</p>
            <p>Result: {result}</p>
            <button onClick={() => setNumber(number + 1)}>Increment Number</button>
            <button onClick={() => setMultiplier(multiplier + 1)}>Increment Multiplier</button>
        </div>
    );
}

export default MemoExample;
