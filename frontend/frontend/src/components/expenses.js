import React from 'react'
import { useEffect, useState } from 'react'
import Expense from '../components/Models/Expense'
import Stack from 'react-bootstrap/esm/Stack';

function Expenses({c_id}) {
    
    const [expenses, setExpenses] = useState([]);
    const [total, setTotal] = useState(0);


    useEffect(() => {  
        fetch('http://localhost:8081/api/expenses/category/'+c_id)
        .then(res => res.json())
        .then(data => {
            setExpenses(data['data'])
            console.log(data['data'])
            setTotal(data['data'].reduce((a, b) => a + b.amount, 0))
        })
    }, [])
    
  return (
    <div>
            <Stack direction='vertical' gap={3}>
                {
                    expenses.map(expense => (
                        <Expense key={expense._id} name={expense.name} amount={expense.amount} description={expense.description} date={expense.createdAt} />))
                }
            </Stack>
    </div>
  )
}



export default Expenses
