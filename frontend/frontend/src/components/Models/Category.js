import React from 'react'

function Category({name, amount, maxamount}) {
  return (
    <div>
        <h1>{name}</h1>
        <p>{amount}</p>
        <p>{maxamount}</p>
    </div>
  )
}

export default Category
