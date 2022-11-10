import React from 'react'
import {useEffect, useState} from 'react'
import Category from '../components/Models/Category'
import CategoryForm from './CategoryForm';
import Button from 'react-bootstrap/esm/Button';
import ExpenseForm from './ExpenseForm';


function Categories() {
    const [categories, setCategories] = useState([]);
    const [showCategory, setShowCategory] = useState(false);
    const [showExpense, setShowExpense] = useState(false);

    useEffect(() => {
        fetch('http://localhost:8081/api/expenseCategories/')
        .then(res => res.json())
        .then(data => {
            setCategories(data['data'])
            console.log(data['data'])
        })
    }, [])


  return (
    <div>
        <div className="d-flex justify-content-between align-items-baseline">
            <h1 className="fw-normal">Categories</h1>
            <div className="d-flex align-items-baseline">
                <h1 className="fw-normal me-2">
                    <Button variant='outline-primary' onClick={() => setShowCategory(true) }>Add Category</Button>
                </h1>
                <h1 className="fw-normal me-2">
                    <Button variant='outline-primary' onClick={() => setShowExpense(true) }>Add Expense</Button>
                </h1>
            </div>
        </div>
        {categories.map(category => (
            <Category key={category._id}
                categoryid={category._id}
                name={category.name}
                amount={category.amount}
                maxamount={category.maxAmount} />
            ))}

         <CategoryForm show={showCategory} handleClose={() => setShowCategory(false)}/>
         <ExpenseForm show={showExpense} handleClose={() => setShowExpense(false)}/>   
    </div>
  )
}

export default Categories
