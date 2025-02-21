import React, { useState } from "react";
import { register } from "../api";
import { useNavigate } from "react-router-dom";

const Register = () => {
  const [user, setUser] = useState({ username: "", email: "", password: "" });
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await register(user);
    navigate("/login");
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type="text" placeholder="Username" onChange={(e) => setUser({ ...user, username: e.target.value })} required />
      <input type="email" placeholder="Email" onChange={(e) => setUser({ ...user, email: e.target.value })} required />
      <input type="password" placeholder="Password" onChange={(e) => setUser({ ...user, password: e.target.value })} required />
      <button type="submit">Register</button>
    </form>
  );
};

export default Register;
