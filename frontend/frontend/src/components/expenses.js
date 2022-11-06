import React from 'react'
import { useEffect, useState } from 'react'
import Expense from '../components/Models/Expense'



function Expenses() {

    const [expenses, setExpenses] = useState([]);

    useEffect(() => {  
        fetch('http://localhost:8000/api/expenses/')
        .then(res => res.json())
        .then(data => {
            setExpenses(data)
        })
    }, [])
    

  return (
    <div>
        {expenses.map(expense => (
            <Expense key={expense.id} amount={expense.amount} name={expense.name} description={expense.description} date={expense.date} />
        ))} 
    </div>
  )
}

export default Expenses
