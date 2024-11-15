import { useEffect, useState } from "react";

/*
[
  { text: "first", completed: true },
  { text: "second", completed: false },
  { text: "third", completed: false },
]
  */

function useLocalStorage(itemName, initialValue) {
  const [item, setItem] = useState(initialValue);

  const [loading, setLoading] = useState(true);

  const [error, setError] = useState(false);

  useEffect(() => {
    setTimeout(() => {
      try {
        const todosFromStorage = window.localStorage.getItem(itemName);
        if (todosFromStorage) {
          setItem(JSON.parse(todosFromStorage));
        } else {
          localStorage.setItem(itemName, JSON.stringify(initialValue));
          setItem(initialValue);
        }
        setLoading(false);
      } catch (error) {
        console.log(error);
        setError(true);
      }
    }, 1000);
  }, [initialValue, itemName]);

  const saveItem = (newItem) => {
    localStorage.setItem(itemName, JSON.stringify(newItem));
    setItem(newItem);
  };

  return { item, saveItem, loading, error };
}

export { useLocalStorage };
