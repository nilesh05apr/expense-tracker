import React from 'react'
import {useEffect, useState} from 'react'
import Category from '../components/Models/Category'


function Categories() {
    const [categories, setCategories] = useState([]);
    useEffect(() => {
        fetch('http://localhost:8081/api/expenseCategories/')
        .then(res => res.json())
        .then(data => {
            setCategories(data['data'])
            console.log(data['data'])
//            console.log(data['data'][0]._id)
        })
    }, [])

  return (
    <div>
        {categories.map(category => (
            <Category key={category._id}
                categoryid={category._id}
                name={category.name}
                amount={12000}
                maxamount={category.maxAmount} />
            ))}
    </div>
  )
}

export default Categories
