import React, {useState} from 'react'
import Expenses from '../expenses'
import Card from 'react-bootstrap/Card'
import Stack from 'react-bootstrap/Stack'
import Button from 'react-bootstrap/Button'
import ProgressBar from 'react-bootstrap/ProgressBar';
import handleCurrency from '../utils';
import Modal from 'react-bootstrap/Modal'
const AMOUNT = 10000;

function Category({categoryid, name, amount, maxamount}) {

  console.log("category id from Category.js: "+categoryid)

  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  const classNames = []
  if (amount > maxamount) {
    classNames.push("mb-4","bg-danger", "bg-opacity-10")
  } else if(amount === maxamount) {
    classNames.push("mb-4","bg-info", "bg-opacity-10")
  }
  else {
    classNames.push("mb-4","bg-success", "bg-opacity-10")
  }

  return (
    <>
      <Card  className={classNames.join(" ")}>
      <Card.Body> 
        <Card.Title className="d-flex justify-content-between align-items-baseline fw-normal mb-3">
          <div className='me-2'>{name}</div>
          <div className="d-flex align-items-baseline mt-2">
            {handleCurrency(AMOUNT)}  
            <span className="text-muted fs-6 ms-1">/ {handleCurrency(maxamount)}</span>
          </div>
        </Card.Title>
        <ProgressBar className='rounded-pill'
          animated
          variant={getColor(AMOUNT, maxamount)}
          min={0} 
          now={AMOUNT} 
          max={maxamount} 
         />
          <Stack direction="horizontal" gap="2" className="mt-4">
            <Button variant="outline-primary" className="ms-auto">Add Expense</Button>
            <Button onClick={handleShow} variant="outline-secondary">View Expenses</Button>
          </Stack>

          <Modal show={show} onHide={handleClose}>
            <Modal.Header closeButton>
              <Modal.Title>
                <Stack direction='horizontal' gap={2}>
                  <div className="me-auto fs-4">{name}</div>
                  <div className="me-auto fs-4">{handleCurrency(AMOUNT)}</div>
                </Stack>
              </Modal.Title>
            </Modal.Header>
            <Modal.Body>
              <Stack direction='vertical' gap={3}>
                <Expenses c_id={categoryid}/>
              </Stack>
            </Modal.Body>
          </Modal>

      </Card.Body>
      </Card>
    </>
  )
}

function getColor(amount, max) {
  const percent = (amount / max) * 100;
  if (percent < 50) return 'success';
  else if (percent < 75) return 'warning';
  else return 'danger';
}

export default Category
