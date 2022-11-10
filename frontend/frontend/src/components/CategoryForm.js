import React from 'react'
import Modal from 'react-bootstrap/Modal';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import { useRef } from 'react';


function CategoryForm({show, handleClose}) {
    const nameRef = useRef();
    const maxAmountRef = useRef();


    function handleSubmit(e) {
        e.preventDefault();
        console.log(nameRef.current.value);
        console.log(maxAmountRef.current.value);
        const data = {
            name: nameRef.current.value,
            amount: 0,
            maxAmount: parseFloat(maxAmountRef.current.value)
        }
        fetch('http://localhost:8081/api/expenseCategories/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
        .then(res => res.json())
        .then(data => {
            console.log(data)
        })
        handleClose();
    }



  return (
    <div>
      <Modal show={show} onHide={handleClose}>
        <Modal.Header closeButton>
            <Modal.Title>Add Category</Modal.Title>
        </Modal.Header>
        <Modal.Body>
            <Form onSubmit={handleSubmit}>
                <Form.Group id="name">
                    <Form.Label>Name</Form.Label>
                    <Form.Control type="text" ref={nameRef} required />
                </Form.Group>
                <Form.Group id="maxAmount">
                    <Form.Label>Max Amount</Form.Label>
                    <Form.Control type="number" ref={maxAmountRef} required />
                </Form.Group>
                <Button className="w-100 my-2" type="submit">Add Category</Button>
            </Form>
        </Modal.Body>
      </Modal>
    </div>



)
}

export default CategoryForm
