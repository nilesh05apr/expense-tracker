import React from 'react'
import Card from 'react-bootstrap/Card'

function Expense({amount, name, description, date}) {
  return (
    <div>
      <Card style={{width:"18rem"}}>
        <Card.Body> 
          <Card.Title>{name}</Card.Title>
          <Card.Subtitle className="mb-2 text-muted">{amount}</Card.Subtitle>
          <Card.Subtitle className="mb-2 text-muted">{date}</Card.Subtitle>
          <Card.Text>
            {description}
          </Card.Text>
        </Card.Body>
      </Card>
    </div>
  )
}

export default Expense
