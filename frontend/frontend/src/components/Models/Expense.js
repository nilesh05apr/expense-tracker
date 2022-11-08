import React from 'react'
import Stack from 'react-bootstrap/esm/Stack';
//import handleCurrency from '../utils';
import { handleCurrency } from '../utils';

function Expense({amount, name, description, date}) {
  return (
    <div>
      <Stack direction='horizontal' gap={2}>
      <div className="ms-2 me-auto">
          <div className="fw-bold">{name}</div>
          <span className='me-auto'>{date}</span>
        </div>
        <div className="me-auto fs-4">{description}</div>
        <div className='fs-5'>{handleCurrency(amount)}</div>
      </Stack>
    </div>
  )
}

export default Expense
