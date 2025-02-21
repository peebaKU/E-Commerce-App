import React, { useEffect, useState } from "react";
import { fetchProducts } from "../api";
import { Product } from "../types";

const Products = () => {
  const [products, setProducts] = useState<Product[]>([]);

  useEffect(() => {
    fetchProducts().then(response => setProducts(response.data));
  }, []);

  return (
    <div>
      <h2>Products</h2>
      <div style={{ display: "grid", gridTemplateColumns: "repeat(3, 1fr)", gap: "20px" }}>
        {products.map((product) => (
          <div key={product.id} style={{ border: "1px solid #ccc", padding: "10px" }}>
            <h3>{product.name}</h3>
            <p>{product.description}</p>
            <strong>${product.price}</strong>
            <button>Add to Cart</button>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Products;
