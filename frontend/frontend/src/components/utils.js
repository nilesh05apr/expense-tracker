export default function handleCurrency(amount) {
    return amount.toLocaleString('en-US', 
    { style: 'currency', currency: 'USD',   minimumFractionDigits: 0 });
}