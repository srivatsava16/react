import React, { useState } from 'react';

export default function App() {
    const PRODUCTS = [
        { category: 'Smartphones', price: '$599', stocked: true, name: 'iPhone 13' },
        { category: 'Smartphones', price: '$499', stocked: true, name: 'Samsung Galaxy S21' },
        { category: 'Smartphones', price: '$399', stocked: false, name: 'Google Pixel 5' },
        { category: 'Laptops', price: '$999', stocked: true, name: 'MacBook Air' },
        { category: 'Laptops', price: '$849', stocked: false, name: 'Dell XPS 13' },
        { category: 'Laptops', price: '$699', stocked: true, name: 'HP Spectre x360' },
        { category: 'Headphones', price: '$199', stocked: true, name: 'Sony WH-1000XM4' },
        { category: 'Headphones', price: '$349', stocked: true, name: 'Bose Noise Cancelling Headphones 700' },
        { category: 'Headphones', price: '$149', stocked: false, name: 'JBL Quantum 800' },
    ];

    return <Div1 products={PRODUCTS} />;
}

function Div1({ products }) {
    const [Bar, setBar] = useState('');
    const [Check, setCheck] = useState(false);

    return (
        <div>
            <div>
                <Search
                    barValue={Bar}
                    barFunction={setBar}
                    checkValue={Check}
                    checkFunction={setCheck}
                />
            </div>
            <div>
                <Second
                    product={products}
                    checkValue={Check}
                    barValue={Bar}
                />
            </div>
        </div>
    );
}

function Second({ product, checkValue, barValue }) {
    const rows = [];
    let lastCategory = null;

    product.forEach((row) => {
        if (row.stocked && !checkValue) {
            return;
        }
        if (barValue && !row.category.toLowerCase().includes(barValue.toLowerCase())) {
            return;
        }

        if (row.category !== lastCategory) {
            rows.push(
                <Col key={row.category} val={row.category} />
            );
        }

        rows.push(
            <Row key={row.name} valu={row} />
        );

        lastCategory = row.category;
    });

    return (
        <table>
            <thead>
            <tr>
                <th>Name</th>
                <th>Price</th>
            </tr>
            </thead>
            <tbody>{rows}</tbody>
        </table>
    );
}

function Row({ valu }) {
    return (
        <tr>
            <td>{valu.name}</td>
            <td>{valu.price}</td>
        </tr>
    );
}

function Col({ val }) {
    return (
        <tr>
            <th>{val}</th>
        </tr>
    );
}

function Search({ barValue, barFunction, checkValue, checkFunction }) {
    return (
        <div>
            <input
                type="text"
                value={barValue}
                onChange={(e) => barFunction(e.target.value)}
                placeholder="..."
            />
            <label>
                <input
                    type="checkbox"
                    checked={checkValue}
                    onChange={(e) => checkFunction(e.target.checked)}
                />
                {' '}
                Only show products in stock
            </label>
        </div>
    );
}
