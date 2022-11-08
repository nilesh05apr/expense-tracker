import React from 'react'
import { useEffect, useState } from 'react'
import Expense from '../components/Models/Expense'
import Stack from 'react-bootstrap/esm/Stack';

function Expenses({c_id}) {
    
    const [expenses, setExpenses] = useState([]);

    useEffect(() => {  
        fetch('http://localhost:8081/api/expenses/')
        .then(res => res.json())
        .then(data => {
            setExpenses(data['data'])
            console.log(data['data'])
        })
    }, [])
    
  return (
    <div>
            <Stack direction='vertical' gap={3}>
                {
                    expenses.filter((expense,{c_id})=>{return expense.expenseCategoryId === c_id;}).map(expense => (
                        <Expense key={expense._id}
                            name={expense.name}
                            amount={expense.amount}
                            description={expense.description}
                            date={expense.createdAt} />
                    ))
                }
            </Stack>
    </div>
  )
}



export default Expenses
