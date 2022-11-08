export function handleCurrency(amount) {
    return amount.toLocaleString('en-US', 
    { style: 'currency', currency: 'USD',   minimumFractionDigits: 0 });
}

export function getTotalAmount(c_id)   {
    let c_amount = 0;
    fetch('http://localhost:8081/api/expenses/category/'+c_id)
    .then(res => res.json())
    .then(data => {
        //console.log(data['data'].reduce((a, b) => a + b.amount, 0))
        c_amount = data['data'].reduce((a, b) => a + b.amount, 0)
    })
    console.log(c_amount)
    return c_amount;
}
