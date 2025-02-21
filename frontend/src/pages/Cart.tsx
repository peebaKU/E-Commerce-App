import React, { useState } from "react";
import { Product } from "../types";

const Cart = () => {
  const [cart, setCart] = useState<Product[]>([]);

  return (
    <div>
      <h2>Shopping Cart</h2>
      {cart.length === 0 ? <p>Your cart is empty.</p> : (
        cart.map((item, index) => (
          <div key={index}>
            <h3>{item.name}</h3>
            <p>${item.price}</p>
          </div>
        ))
      )}
    </div>
  );
};

export default Cart;
