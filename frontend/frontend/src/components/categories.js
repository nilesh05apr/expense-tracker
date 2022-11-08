import React from 'react'
import {useEffect, useState} from 'react'
import Category from '../components/Models/Category'
import { getTotalAmount } from './utils';


function Categories() {
    const [categories, setCategories] = useState([]);
    const [total, setTotal] = useState(0);

    useEffect(() => {
        fetch('http://localhost:8081/api/expenseCategories/')
        .then(res => res.json())
        .then(data => {
            setCategories(data['data'])
            console.log(data['data'])
            setTotal(getTotalAmount(data['data'][0]._id))
            console.log(total)
        })
    }, [total])

  return (
    <div>
        {categories.map(category => (
            <Category key={category._id}
                categoryid={category._id}
                name={category.name}
                amount={total}
                maxamount={category.maxAmount} />
            ))}
    </div>
  )
}

export default Categories
