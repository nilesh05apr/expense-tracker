import React from 'react'
import {useEffect, useState} from 'react'
import Category from '../components/Models/Category'
import {ListGroup} from 'react-bootstrap/ListGroup'



function categories() {
    const [categories, setCategories] = useState([]);
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
        <ListGroup>
            <ListGroup.Item>
                {categories.map(category => (
                    <Category key={category.id}
                        name={category.name}
                        amount={category.amount}
                        maxamount={category.maxamount} />
                    ))}
            </ListGroup.Item>
        </ListGroup>
    </div>
  )
}

export default categories
