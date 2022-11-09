import React from 'react'
import Modal from 'react-bootstrap/Modal';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import { useRef } from 'react';
import { useCategory } from '../contexts/CategoryContext';


export default function ExpenseForm({show, handleClose}) {
    const nameRef = useRef();
    const amountRef = useRef();
    const descriptionRef = useRef();
    const categoryRef = useRef();

    const {category, createExpense} = useCategory();

    console.log(category);

    function handleSubmit(e) {
        e.preventDefault();
        console.log(nameRef.current.value);
        console.log(amountRef.current.value);
        console.log(descriptionRef.current.value);
        console.log(categoryRef.current.value);
        //createExpense({name:nameRef.current.value, amount:amountRef.current.value, description:descriptionRef.current.value, expenseCategoryId:categoryRef.current.value});
    }
    

  return (
    <div>
        <Modal show={show} onHide={handleClose}>
            <Modal.Header closeButton>
                <Modal.Title>Add Expense</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form onSubmit={handleSubmit}>
                    <Form.Group id="name">
                        <Form.Label>Name</Form.Label>
                        <Form.Control type="text" ref={nameRef} required />
                    </Form.Group>
                    <Form.Group id="amount">
                        <Form.Label>Amount</Form.Label>
                        <Form.Control type="number" ref={amountRef} required />
                    </Form.Group>
                    <Form.Group id="description">
                        <Form.Label>Description</Form.Label>
                        <Form.Control type="text" ref={descriptionRef} required />
                    </Form.Group>
                    <Form.Group id="category">
                        <Form.Label>Category</Form.Label>
                        <Form.Control as="select" ref={categoryRef} required>
                            {category.map(c => {
                                return <option key={c._id} value={c._id}>{c.name}</option>
                            })}
                        </Form.Control>
                    </Form.Group>
                    <Button className="w-100 my-2" type="submit">Add Expense</Button>
                </Form>
            </Modal.Body>
        </Modal>      
    </div>
  )
}
