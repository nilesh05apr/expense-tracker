import React, { useState, useEffect, useContext } from 'react';
import useLocalStorage from '../hooks/useLocalStorage';

const CategoryContext = React.createContext();

export function useCategory() {
  return useContext(CategoryContext);
}

export const CategoryProvider = ({ children }) => {
  const [category, setCategory] = useLocalStorage('category', []);
  const [expense, setExpense] = useLocalStorage('expense', []);

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

    const createExpense = (name, amount, description, category) => {
        setExpense(prevExpense => {
            return [...prevExpense, { name, amount, description, category }]
        })
    }

    const deleteExpense = (id) => {
        setExpense(prevExpense => {
            return prevExpense.filter(expense => expense.id !== id)
        })
    }

    const value = {
        category,
        expense,
        createCategory,
        deleteCategory,
        createExpense,
        deleteExpense
    }

  return (
    <CategoryContext.Provider value={value}>
      {children}
    </CategoryContext.Provider>
  );
}