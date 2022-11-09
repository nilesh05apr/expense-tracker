import React from 'react'
import Modal from 'react-bootstrap/Modal';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import { useRef } from 'react';
import { useCategory } from '../contexts/CategoryContext';


function CategoryForm({show, handleClose}) {
    const nameRef = useRef();
    const maxAmountRef = useRef();

    const {createCategory} = useCategory();

    function handleSubmit(e) {
        e.preventDefault();
        console.log(nameRef.current.value);
        console.log(maxAmountRef.current.value);
        createCategory({name:nameRef.current.value, maxAmount:maxAmountRef.current.value});
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
