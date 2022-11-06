import React from 'react'

function Expense({amount, name, description, date}) {
  return (
    <div>
        <h1>{name}</h1>
        <h3>{description}</h3> 
        <p>{amount}</p>
        <p>{date}</p>
    </div>
  )
}

export default Expense
