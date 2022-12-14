import React, { useContext, useState } from 'react';
import useLocalStorage from '../hooks/useLocalStorage';

const CategoryContext = React.createContext();

export function useCategory() {
  return useContext(CategoryContext);
}

export const CategoryProvider = ({ children }) => {
  // const [category, setCategory] = useLocalStorage('category', []);
  // const [expense, setExpense] = useLocalStorage('expense', []);

  const [category, setCategory] = useState([]);
  const [expense, setExpense] = useState([]);


    const createCategory = (name, maxAmount) => {
        setCategory(prevCategory => {
            return [...prevCategory, { name, maxAmount }]
        })
    }

    const deleteCategory = (id) => {
        setCategory(prevCategory => {
            return prevCategory.filter(category => category.id !== id)
        })
    }

    const createExpense = (name, amount, description, expenseCategoryId) => {
        setExpense(prevExpense => {
            return [...prevExpense, { name, amount, description, expenseCategoryId }]
        })
    }

    const deleteExpense = (id) => {
        setExpense(prevExpense => {
            return prevExpense.filter(expense => expense.id !== id)
        })
    }

    const getCategoryExpenses = (id) => {
        return expense.filter(expense => expense.expenseCategoryId === id)
    }

    const value = {
        category,
        expense,
        createCategory,
        deleteCategory,
        createExpense,
        deleteExpense,
        getCategoryExpenses
    }

  return (
    <CategoryContext.Provider value={value}>
      {children}
    </CategoryContext.Provider>
  );
}